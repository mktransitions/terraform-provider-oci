// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcediscovery

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-exec/tfexec"

	"github.com/hashicorp/hcl2/hclwrite"
)

var isInitDone bool
var initLock sync.Mutex

/*  ctxLock is the common lock for the whole struct
WARN: Make sure NOT to pass resourceDiscoveryContext as value,
as that would copy the struct and locks should not be copied
*/

// Resource discovery Exit status
type Status int

const (
	// Exit statuses
	StatusSuccess Status = iota
	StatusFail
	StatusPartialSuccess Status = 64

	// parallelism configs
	ChunkSize int = 10
)

var isAllDataSourceLock sync.Mutex

func createResourceDiscoveryContext(clients *tf_client.OracleClients, args *tf_export.ExportCommandArgs, tenancyOcid string) (*tf_export.ResourceDiscoveryContext, error) {

	result := &tf_export.ResourceDiscoveryContext{
		Clients:             clients,
		ExportCommandArgs:   args,
		TenancyOcid:         tenancyOcid,
		DiscoveredResources: []*tf_export.OCIResource{},
		SummaryStatements:   []string{},
		ErrorList: tf_export.ErrorList{
			Errors: []*tf_export.ResourceDiscoveryError{},
		},
		TargetSpecificResources: false,
		ResourceHintsLookup:     createResourceHintsLookupMap(),
	}
	// Use user provided terraform-provider-oci executable
	if pluginDir := utils.GetEnvSettingWithBlankDefault("provider_bin_path"); pluginDir != "" {
		result.TerraformProviderBinaryPath = pluginDir
		utils.Logf("[INFO] terraform provider binary path (pluginDir) set using `provider_bin_path`: '%s'", result.TerraformProviderBinaryPath)
	}

	if *result.CompartmentId == "" {
		*result.CompartmentId = tenancyOcid
		tf_export.Vars["tenancy_ocid"] = fmt.Sprintf("\"%s\"", tenancyOcid)
		tf_export.ReferenceMap[tenancyOcid] = tf_export.TfHclVersionvar.GetVarHclString("tenancy_ocid")
	} else {
		tf_export.Vars["compartment_ocid"] = fmt.Sprintf("\"%s\"", *result.CompartmentId)
		tf_export.ReferenceMap[*result.CompartmentId] = tf_export.TfHclVersionvar.GetVarHclString("compartment_ocid")
	}

	result.ExpectedResourceIds = tf_export.ConvertStringSliceToSet(args.IDs, true)

	re := regexp.MustCompile(`oci_([^:]+):(.+$)`)

	for id := range result.ExpectedResourceIds {
		subMatchAll := re.FindStringSubmatch(id)
		if subMatchAll != nil && len(subMatchAll) == 3 {
			result.TargetSpecificResources = true
			break
		}
	}
	// validate terraform version and initialize terraform for import - only required if generating state file
	if args.GenerateState {
		if tf, terraformCLIPath, err := createTerraformStruct(args); err != nil {
			return result, err
		} else {
			result.Terraform = tf
			result.TerraformCLIPath = terraformCLIPath
		}
	}
	return result, nil
}

type resourceDiscoveryStep interface {
	discover() error
	getOmittedResources() []*tf_export.OCIResource
	writeTmpConfigurationForImport() error
	writeConfiguration() error
	writeTmpState() error
	getBaseStep() *resourceDiscoveryBaseStep
	mergeTempStateFiles(tmpStateOutputDir string) error
	mergeGeneratedStateFile() error
	getDiscoveredResources() []*tf_export.OCIResource
	updateTimeTakenForDiscovery(timeTaken time.Duration)
	updateTimeTakenForGeneratingState(timeTaken time.Duration)
}

type resourceDiscoveryBaseStep struct {
	ctx                         *tf_export.ResourceDiscoveryContext
	name                        string
	discoveredResources         []*tf_export.OCIResource
	omittedResources            []*tf_export.OCIResource
	tempState                   interface{}
	discoveryParallelism        bool
	timeTakenForDiscovery       time.Duration
	timeTakenForGeneratingState time.Duration
}

func (r *resourceDiscoveryBaseStep) mergeTempStateFiles(tmpStateOutputDir string) error {
	defer elapsed(fmt.Sprintf("merging temp state files for %v", r.name), nil, 0)()
	files, err := ioutil.ReadDir(tmpStateOutputDir)
	if err != nil {
		return err
	}
	// loop over tmp state files for each chunk and merge all to form temp State for the service
	for _, file := range files {
		var tempState interface{}
		tmpFilePath := filepath.Join(tmpStateOutputDir, file.Name())
		if !strings.HasSuffix(file.Name(), ".backup") { // ignore the backup file created by terraform
			if jsonState, err := ioutil.ReadFile(tmpFilePath); err != nil {
				return err
			} else {
				if err := json.Unmarshal(jsonState, &tempState); err != nil {
					return err
				}
			}
			if r.tempState == nil {
				r.tempState = tempState
			} else {
				r.tempState, _ = mergeState(r.tempState, tempState)
			}
		}
	}
	return nil
}

var terraformInitMockVar = func(r *resourceDiscoveryBaseStep, backgroundCtx context.Context, initArgs []tfexec.InitOption) error {
	return r.ctx.Terraform.Init(backgroundCtx, initArgs...)
}

func (r *resourceDiscoveryBaseStep) writeTmpState() error {
	defer elapsed(fmt.Sprintf("writing temp state for %d '%s' resources", len(r.getDiscoveredResources()), r.name), nil, 0)()
	// Run terraform init if not already done
	if !isInitDone {
		utils.Debugf("[DEBUG] acquiring lock to run terraform init")
		initLock.Lock()
		defer func() {
			utils.Debugf("[DEBUG] releasing lock")
			initLock.Unlock()
		}()
		// Check for existence of .terraform folder to make sure init is not run already by another thread
		if _, err := os.Stat(fmt.Sprintf("%s%s.terraform", *r.ctx.OutputDir, string(os.PathSeparator))); os.IsNotExist(err) {
			// Run init command if not already run
			utils.Debugf("[DEBUG] writeTmpState: running init")
			backgroundCtx := context.Background()

			var initArgs []tfexec.InitOption

			if r.ctx.TerraformProviderBinaryPath != "" {
				utils.Logf("[INFO] plugin dir set to: '%s'", r.ctx.TerraformProviderBinaryPath)
				initArgs = append(initArgs, tfexec.PluginDir(r.ctx.TerraformProviderBinaryPath))
			}

			if err := terraformInitMockVar(r, backgroundCtx, initArgs); err != nil {
				utils.Debugf("[ERROR] error occured while terraform init: %s", err.Error())
				return err
			}
			isInitDone = true
		}
	}
	tmpStateOutputDir := filepath.Join(*r.ctx.OutputDir, "tmp", r.name)
	tmpStateOutputFilePrefix := filepath.Join(tmpStateOutputDir, globalvar.DefaultTmpStateFile)

	if err := os.RemoveAll(tmpStateOutputDir); err != nil {
		utils.Logf("[WARN] unable to delete existing tmp state directory %s", tmpStateOutputDir)
		return err
	}

	isAllDataSources := true
	totalResources := len(r.discoveredResources)
	// divide list of discovered resources which is a slice into chunks
	// process each chunk in parallel
	chunkSize := ChunkSize // chunk size defines number of resources in each chunk.
	// if there are additional chunks required for left over resources.
	// For example, if chunk size is 5 and we have 8 resources then 8/5 gives int output as 1.
	// So chunk 1 will occupy 5 resources and 8 % 5 = 3. For remaining 3 resources we need additional chunk.
	additionalChunks := 1 // no. of additional chunks required to process (totalResources % chunkSize) resources.
	if totalResources%chunkSize == 0 {
		additionalChunks = 0
	}
	totalChunks := totalResources/chunkSize + additionalChunks
	var importWg sync.WaitGroup
	importWg.Add(totalChunks) // we need to wait for all chunks to finish importing resources
	// we Create buffered channel to control max parallel chunks that can be executed in parallel
	semImport := make(chan struct{}, MaxParallelChunks)
	// loop over chunks
	for chunkIdx := 0; chunkIdx < totalResources; chunkIdx += chunkSize {
		// position of last element of chunk
		lastPos := chunkIdx + chunkSize
		semImport <- struct{}{}
		// in case last chunk isn't full set the lastPos accordingly
		if lastPos > totalResources {
			lastPos = totalResources
		}
		go func(resources []*tf_export.OCIResource, chunkIndex int) {
			for _, res := range resources {
				fileName := tmpStateOutputFilePrefix + fmt.Sprint(chunkIndex)
				importResource(r.ctx, res, fileName)
				if res.TerraformTypeInfo != nil && !res.TerraformTypeInfo.IsDataSource {
					isAllDataSourceLock.Lock()
					isAllDataSources = false
					isAllDataSourceLock.Unlock()
				}
			}
			<-semImport
			importWg.Done()
			// take resources beginning at chunkIdx upto and excluding lastPos
		}(r.discoveredResources[chunkIdx:lastPos], chunkIdx)
	}
	// wait for all chunks to finish importing resources
	importWg.Wait()
	// The found resource only include the data sources (ADs and namespaces) that resource discovery adds
	if isAllDataSources {
		return nil
	}
	err := r.mergeTempStateFiles(tmpStateOutputDir)
	if err != nil {
		return err
	}
	return nil
}

// writeTmpConfigurationForImport writes temporary configuration to run terraform import on the discovered resources
// It only writes the resource block and skips the resource fields
// The configuration will be discarded and written again after import is completed for all resources
func (r *resourceDiscoveryBaseStep) writeTmpConfigurationForImport() error {
	defer elapsed(fmt.Sprintf("writing temp configuration for %d %s resources", len(r.getDiscoveredResources()), r.name), nil, 0)()
	configOutputFile := fmt.Sprintf("%s%s%s.tf", *r.ctx.OutputDir, string(os.PathSeparator), r.name)
	tmpConfigOutputFile := fmt.Sprintf("%s%s%s.tf.tmp", *r.ctx.OutputDir, string(os.PathSeparator), r.name)

	file, err := os.OpenFile(tmpConfigOutputFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	// Build the HCL config
	builder := &strings.Builder{}
	builder.WriteString("## This is tmp config to run import for resources\n\n")
	for _, resource := range r.discoveredResources {
		if resource.TerraformTypeInfo != nil && resource.TerraformTypeInfo.IsDataSource {
			utils.Logf("[INFO]Skipping the data source as we are just writing temp configuration to import the resource", resource.TerraformClass, resource.TerraformName)
		} else {
			utils.Debugf("Writting temp config for resource '%s.%s'", resource.TerraformClass, resource.TerraformName)
			builder.WriteString(fmt.Sprintf("resource %s %s {}\n\n", resource.TerraformClass, resource.TerraformName))
		}

		r.ctx.CtxLock.Lock()
		r.ctx.DiscoveredResources = append(r.ctx.DiscoveredResources, resource)
		r.ctx.CtxLock.Unlock()
	}
	utils.Debugf("[DEBUG] dump temp configuration into '%s'", file.Name())
	_, err = file.WriteString(string(builder.String()))
	if err != nil {
		_ = file.Close()
		return err
	}

	if fErr := file.Close(); fErr != nil {
		return fErr
	}

	if err := os.Rename(tmpConfigOutputFile, configOutputFile); err != nil {
		return err
	}
	return nil
}

func (r *resourceDiscoveryBaseStep) writeConfiguration() error {
	defer elapsed(fmt.Sprintf("writing actual configuration for %d %s resources", len(r.getDiscoveredResources()), r.name), nil, 0)()

	//Do not generate empty terraform configuration file
	if len(r.getDiscoveredResources()) == 0 {
		return nil
	}
	configOutputFile := fmt.Sprintf("%s%s%s.tf", *r.ctx.OutputDir, string(os.PathSeparator), r.name)
	tmpConfigOutputFile := fmt.Sprintf("%s%s%s.tf.tmp", *r.ctx.OutputDir, string(os.PathSeparator), r.name)
	file, err := os.OpenFile(tmpConfigOutputFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	// Build the HCL config
	// Note that we still build a TF file even if no resources were discovered for this TF file.
	// A user may run this command multiple times and may see stale resources if we don't overwrite the file with
	// an empty one.
	builder := &strings.Builder{}
	builder.WriteString("## This configuration was generated by terraform-provider-oci\n\n")

	exportedResourceCount := 0
	for _, resource := range r.discoveredResources {

		// Skip writing the config for resources for which import command failed
		if !resource.IsErrorResource {
			utils.Logf("[INFO] ===> Generating resource '%s'", resource.GetTerraformReference())
			if err := resource.GetHCLString(builder, tf_export.ReferenceMap); err != nil {
				_ = file.Close()
				return err
			}

			if resource.TerraformTypeInfo != nil && len(resource.TerraformTypeInfo.IgnorableRequiredMissingAttributes) > 0 {
				attributes := make([]string, 0, len(resource.TerraformTypeInfo.IgnorableRequiredMissingAttributes))
				for attribute := range resource.TerraformTypeInfo.IgnorableRequiredMissingAttributes {
					attributes = append(attributes, attribute)
				}
				missingAttributesPerResourceLock.Lock()
				if r.ctx.MissingAttributesPerResource == nil {
					r.ctx.MissingAttributesPerResource = make(map[string][]string)
				}
				r.ctx.MissingAttributesPerResource[resource.GetTerraformReference()] = attributes
				missingAttributesPerResourceLock.Unlock()
			}

			r.ctx.DiscoveredResources = append(r.ctx.DiscoveredResources, resource)
			exportedResourceCount++
		} else {
			utils.Debugf("[DEBUG] ===> SKIP Generating resource '%s' .", resource.GetTerraformReference())
			// remove missing attributes info if present for a failed resource
			missingAttributesPerResourceLock.Lock()
			if _, ok := r.ctx.MissingAttributesPerResource[resource.GetTerraformReference()]; ok {
				delete(r.ctx.MissingAttributesPerResource, resource.GetTerraformReference())
			}
			missingAttributesPerResourceLock.Unlock()
		}
	}

	// Format the HCL config
	formattedString := hclwrite.Format([]byte(builder.String()))

	_, err = file.WriteString(string(formattedString))
	if err != nil {
		_ = file.Close()
		return err
	}

	if fErr := file.Close(); fErr != nil {
		return fErr
	}

	if err := os.Rename(tmpConfigOutputFile, configOutputFile); err != nil {
		return err
	}

	if r.ctx.TargetSpecificResources {
		r.ctx.SummaryStatements = append(r.ctx.SummaryStatements, fmt.Sprintf("Found %d resources. Generated under '%s'", exportedResourceCount, configOutputFile))
	} else {
		r.ctx.SummaryStatements = append(r.ctx.SummaryStatements, fmt.Sprintf("Found %d '%s' resources. Generated under '%s'", exportedResourceCount, r.name, configOutputFile))
	}
	r.ctx.SummaryStatements = append(r.ctx.SummaryStatements, fmt.Sprintf("Time taken for discovery: %v, generating state: %v", r.timeTakenForDiscovery, r.timeTakenForGeneratingState))
	return nil
}

func (r *resourceDiscoveryBaseStep) getOmittedResources() []*tf_export.OCIResource {
	return r.omittedResources
}

func (r *resourceDiscoveryBaseStep) getDiscoveredResources() []*tf_export.OCIResource {
	return r.discoveredResources
}

func (r *resourceDiscoveryBaseStep) updateTimeTakenForDiscovery(timeTaken time.Duration) {
	r.timeTakenForDiscovery = timeTaken
}
func (r *resourceDiscoveryBaseStep) updateTimeTakenForGeneratingState(timeTaken time.Duration) {
	r.timeTakenForGeneratingState = timeTaken
}

func (r *resourceDiscoveryBaseStep) getBaseStep() *resourceDiscoveryBaseStep {
	return r
}

type resourceDiscoveryWithGraph struct {
	resourceDiscoveryBaseStep
	root          *tf_export.OCIResource
	resourceGraph tf_export.TerraformResourceGraph
}

func (r *resourceDiscoveryWithGraph) discover() error {
	var err error
	var ociResources []*tf_export.OCIResource

	// for root step, setting discovery parallelism on
	// turn it off for further levels
	r.discoveryParallelism = true

	ociResources, err = findResources(r.ctx, r.root, r.resourceGraph, r.discoveryParallelism)
	if err != nil {
		return err
	}

	// Filter out omitted resources from export
	r.discoveredResources = []*tf_export.OCIResource{}
	r.omittedResources = []*tf_export.OCIResource{}
	for _, resource := range ociResources {
		if !resource.OmitFromExport {

			tf_export.RefMapLock.Lock()
			tf_export.ReferenceMap[resource.Id] = resource.GetHclReferenceIdString()
			tf_export.RefMapLock.Unlock()

			r.discoveredResources = append(r.discoveredResources, resource)
		} else {
			utils.Debugf("[DEBUG] resource %s is omitted", resource.GetTerraformReference())
			r.omittedResources = append(r.omittedResources, resource)
		}
	}
	utils.Logf("[INFO] Discovery complete for step root %s", r.name)
	return nil
}

type resourceDiscoveryWithTargetIds struct {
	resourceDiscoveryBaseStep
	exportIds map[string]string // map of IDs and their respective resource types
}

func createResourceHintsLookupMap() map[string]*tf_export.TerraformResourceHints {
	result := map[string]*tf_export.TerraformResourceHints{}

	for _, graphCollection := range []map[string]tf_export.TerraformResourceGraph{tf_export.CompartmentResourceGraphs, tf_export.TenancyResourceGraphs} {
		for _, graph := range graphCollection {
			for _, associations := range graph {
				for _, assoc := range associations {
					result[assoc.ResourceClass] = assoc.TerraformResourceHints
				}
			}
		}
	}
	return result
}

func (r *resourceDiscoveryWithTargetIds) discover() error {
	sortedIds := make([]string, len(r.ctx.ExpectedResourceIds))
	idx := 0
	for id, _ := range r.ctx.ExpectedResourceIds {
		sortedIds[idx] = id
		idx++
	}
	sort.Strings(sortedIds)

	re := regexp.MustCompile(`(oci_[^:]+):(.+$)`)

	for _, id := range sortedIds {
		subMatchAll := re.FindStringSubmatch(id)
		if len(subMatchAll) != 3 {
			utils.Logf("[WARN] Encountered invalid ID tuple '%s'", id)
			continue
		}

		resourceClass := subMatchAll[1]
		resourceId, _ := url.PathUnescape(subMatchAll[2])

		utils.Logf("===> Finding resource with ID '%s' and type '%s'", resourceId, resourceClass)
		resourceSchema, exists := tf_export.ResourcesMap[resourceClass]
		if !exists || resourceSchema.Read == nil {
			utils.Logf("[WARN] No valid resource schema could be found. Skipping.")
			continue
		}

		d := resourceSchema.Data(nil)
		d.SetId(resourceId)
		if err := resourceSchema.Read(d, r.ctx.Clients); err != nil {
			utils.Logf("[WARN] Unable to read resource due to error: %v", err)
			continue
		}

		if d.Id() == "" {
			utils.Logf("[WARN] Resource ID was voided because resource could not be found. Skipping.")
			continue
		}

		resourceHint, err := r.ctx.GetResourceHint(resourceClass)
		if err != nil {
			continue
		}
		ociResource, err := tf_export.GetOciResource(d, resourceSchema.Schema, *r.ctx.CompartmentId, resourceHint, resourceId)
		if err != nil {
			return err
		}

		if resourceHint.ProcessDiscoveredResourcesFn != nil {
			processResults, err := resourceHint.ProcessDiscoveredResourcesFn(r.ctx, []*tf_export.OCIResource{ociResource})
			if err != nil {
				return err
			}

			if len(processResults) != 1 {
				utils.Logf("[WARN] processing of single resource resulted in %v resources being returned", len(processResults))
				continue
			}
			ociResource = processResults[0]
		}

		if ociResource.TerraformName, err = tf_export.GenerateTerraformNameFromResource(ociResource.SourceAttributes, resourceSchema.Schema); err != nil {
			ociResource.TerraformName = fmt.Sprintf("export_%s", resourceHint.ResourceAbbreviation)
			ociResource.TerraformName = tf_export.CheckDuplicateResourceName(ociResource.TerraformName)
		}

		r.discoveredResources = append(r.discoveredResources, ociResource)

		r.ctx.ExpectedResourceIds[id] = true
		// expectedResourceIds contains tuples in case of export using ids and for related resources the ids will not be a tuple
		//delete(r.ctx.expectedResourceIds, id)
		//r.ctx.expectedResourceIds[ociResource.id] = true

		if _, hasRelatedResources := tf_export.ExportRelatedResourcesGraph[resourceHint.ResourceClass]; hasRelatedResources && r.ctx.IsExportWithRelatedResources {
			utils.Logf("[INFO] resource discovery: finding related resources for %s\n", resourceHint.ResourceClass)
			ociResources, err := findResources(r.ctx, ociResource, tf_export.ExportRelatedResourcesGraph, r.discoveryParallelism)
			if err != nil {
				return err
			}
			/*
				 1. Current closure graph generates only related resources but we may need to filter resources in future as the graph grows
					Because hints use datasources and if data source does not take parent param then it may generate unrelated resources
				 2. With current implementation, resource.omitFromExport will be true for child resources but we do not filter resources. If we add filtering to handle #1,
				 	then logic to set resource.omitFromExport will also need Update to handle related resources
			*/
			r.discoveredResources = append(r.discoveredResources, ociResources...)
		}
		// Add resource reference to referenceMap for discovered resources
		// If there are more than 1 resources found, this will help generate the possible references if the resources are linked
		for _, resource := range r.discoveredResources {
			tf_export.ReferenceMap[resource.Id] = resource.GetHclReferenceIdString()
		}
	}
	return nil
}

func init() {
	// TODO: The following changes to resource hints are deviations from what can currently be handled by the core resource discovery/generation logic
	// We should strive to eliminate these deviations by either improving the core logic or code generator

}

/*
mergeState merges 2 json state files
*/
var mergeState = func(state1 interface{}, state2 interface{}) (interface{}, error) {

	state1Bytes, _ := json.MarshalIndent(state1, "", "\t")
	state2Bytes, _ := json.MarshalIndent(state2, "", "\t")

	out1 := map[string]interface{}{}
	if err := json.Unmarshal(state1Bytes, &out1); err != nil {
		return out1, fmt.Errorf("[ERROR] error occurred while generating state file for resource discovery: %s", err.Error())
	}

	out2 := map[string]interface{}{}
	if err := json.Unmarshal(state2Bytes, &out2); err != nil {
		return out1, fmt.Errorf("[ERROR] error occurred while generating state file for resource discovery: %s", err.Error())
	}

	state1resources, _ := out1["resources"].([]interface{})
	state2resources, _ := out2["resources"].([]interface{})

	out1["resources"] = append(state1resources, state2resources...)

	return out1, nil

}

func (r *resourceDiscoveryBaseStep) mergeGeneratedStateFile() error {
	if r.tempState == nil {
		return nil
	}
	utils.Debugf("[DEBUG] merging state file for %s", r.name)
	defer elapsed(fmt.Sprintf("[DEBUG] merging state file for %s", r.name), nil, 0)()
	if r.ctx.State == nil {
		// if state exists for the step, initialize the final state
		r.ctx.State = r.tempState
	} else {
		// merge the state for step to final state
		r.ctx.State, _ = mergeState(r.ctx.State, r.tempState)
	}

	return nil

}
