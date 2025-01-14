// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcediscovery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-version"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"golang.org/x/mod/semver"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	tf_provider "github.com/oracle/terraform-provider-oci/internal/provider"
	utils "github.com/oracle/terraform-provider-oci/internal/utils"

	hcinstall "github.com/hashicorp/hc-install"
	"github.com/hashicorp/hc-install/fs"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/src"
)

type ResourceDiscoveryStage int

const (
	Discovery       ResourceDiscoveryStage = 1
	GeneratingState                        = 2
)

var (
	referenceResourceNameSet map[string]bool // this set contains terraform resource names for the references in referenceMap

	missingAttributesPerResourceLock sync.Mutex
	sem                              chan struct{}
	exportConfigProvider             oci_common.ConfigurationProvider
	MaxParallelFindResource          int
	MaxParallelChunks                int

	tfexecConfigVar                     = tfexec.Config
	tfexecStateVar                      = tfexec.State
	getProviderEnvSettingWithDefaultVar = utils.GetProviderEnvSettingWithDefault
	getExportConfigVar                  = getExportConfig
	osStatvar                           = os.Stat
	hcInstallerEnsureVar                = func(installer *hcinstall.Installer, ctx context.Context, sources []src.Source) (string, error) {
		return installer.Ensure(ctx, sources)
	}
	isDirVar = func(file os.FileInfo) bool {
		return file.IsDir()
	}
	tfVersionVar = func(tf *tfexec.Terraform, backgroundCtx context.Context) (*version.Version, map[string]*version.Version, error) {
		return tf.Version(backgroundCtx, true)
	}
	terraformInitVar = func(ctx *tf_export.ResourceDiscoveryContext, backgroundCtx context.Context, initArgs []tfexec.InitOption) error {
		return ctx.Terraform.Init(backgroundCtx, initArgs...)
	}
	tfProviderGetSdkConfigProvider  = tf_provider.GetSdkConfigProvider
	sdkConfigProviderTenancyOCIDVar = func(sdkConfigProvider oci_common.ConfigurationProvider) (string, error) {
		return sdkConfigProvider.TenancyOCID()
	}
	tfProviderBuildConfigureClientFn  = tf_provider.BuildConfigureClientFn
	createSDKClientsVar               = tf_client.CreateSDKClients
	identityClientListCompartmentsVar = func(clients *tf_client.OracleClients, req oci_identity.ListCompartmentsRequest) (oci_identity.ListCompartmentsResponse, error) {
		return clients.IdentityClient().ListCompartments(context.Background(), req)
	}
	identityClientGetCompartmentVar = func(clients *tf_client.OracleClients, getCompartmentRequest oci_identity.GetCompartmentRequest) (oci_identity.GetCompartmentResponse, error) {
		return clients.IdentityClient().GetCompartment(context.Background(), getCompartmentRequest)
	}
	ctxTerraformImportVar = func(ctx *tf_export.ResourceDiscoveryContext, ctxBackground context.Context, address, id string, importArgs ...tfexec.ImportOption) error {
		return ctx.Terraform.Import(ctxBackground, address, id, importArgs...)
	}
)

func elapsed(what string, step *resourceDiscoveryBaseStep, stage ResourceDiscoveryStage) func() {
	start := time.Now()
	return func() {
		totalTime := time.Since(start)
		utils.Debugf("[DEBUG] %s took %v\n", what, totalTime)
		if step != nil {
			switch stage {
			case Discovery:
				step.updateTimeTakenForDiscovery(totalTime)
			case GeneratingState:
				step.updateTimeTakenForGeneratingState(totalTime)
			}
		}
	}
}

func init() {
	tf_export.ResourceNameCount = map[string]int{}
	tf_export.Vars = map[string]string{}
	tf_export.ReferenceMap = map[string]string{}
	tf_export.VarsExportForResourceLevel = map[string][]string{}
	tf_export.VarsExportForGlobalLevel = []string{}

	tf_export.CompartmentScopeServices = make([]string, len(tf_export.CompartmentResourceGraphs))
	idx := 0
	for mode := range tf_export.CompartmentResourceGraphs {
		tf_export.CompartmentScopeServices[idx] = mode
		idx++
	}

	tf_export.TenancyScopeServices = make([]string, len(tf_export.TenancyResourceGraphs))
	idx = 0
	for mode := range tf_export.TenancyResourceGraphs {
		tf_export.TenancyScopeServices[idx] = mode
		idx++
	}

	tf_export.IsMissingRequiredAttributes = false
}

func printResourceGraphResources(resourceGraphs map[string]tf_export.TerraformResourceGraph, scope string) error {
	for graphName, resourceGraph := range resourceGraphs {
		// Need a set here because the same resource type may have multiple associations in the same graph
		// This avoids adding duplicates of those resource types
		resourceSet := map[string]bool{}
		for _, association := range resourceGraph {
			for _, hint := range association {
				if _, isResource := tf_export.ResourcesMap[hint.ResourceClass]; isResource {
					resourceSet[hint.ResourceClass] = true
				}
			}
		}

		if len(resourceSet) > 0 {
			supportedResources := make([]string, len(resourceSet))
			idx := 0
			for resourceClass := range resourceSet {
				supportedResources[idx] = resourceClass
				idx++
			}

			sort.Strings(supportedResources)
			utils.Logf("%s (%s-scope resources)", graphName, scope)
			utils.Log("===========")
			for _, resourceClass := range supportedResources {
				utils.Logf("- %s", resourceClass)
			}
			utils.Logln("")
		}
	}
	return nil
}

func RunListExportableResourcesCommand() error {
	tf_export.ResourcesMap = tf_provider.ResourcesMap()
	tf_export.DatasourcesMap = tf_provider.DataSourcesMap()

	utils.Logln("List of Discoverable Oracle Cloud Infrastructure Resources")

	if err := printResourceGraphResources(tf_export.TenancyResourceGraphs, "tenancy"); err != nil {
		return err
	}

	if err := printResourceGraphResources(tf_export.CompartmentResourceGraphs, "compartment"); err != nil {
		return err
	}
	return nil
}

type ExportService struct {
	Name  string
	Scope string
}

const (
	TenancyScope     = "tenancy"
	CompartmentScope = "compartment"
)

func RunListExportableServicesCommand(listExportServicesPath string) error {

	utils.Logln("List Discoverable Oracle Cloud Infrastructure Services")

	services := make([]*ExportService, 0)
	for _, service := range tf_export.TenancyScopeServices {
		services = append(services, &ExportService{
			Name:  service,
			Scope: TenancyScope,
		})
	}

	for _, service := range tf_export.CompartmentScopeServices {
		services = append(services, &ExportService{
			Name:  service,
			Scope: CompartmentScope,
		})
	}

	servicesJson, err := json.MarshalIndent(services, "", "")
	if err != nil {
		return fmt.Errorf("[ERROR] Error marshalling services to JSON: %v", err)
	}

	if listExportServicesPath != "" {
		if err := ioutil.WriteFile(listExportServicesPath, servicesJson, 0644); err != nil {
			return err
		} else {
			utils.Logf("[INFO] Services written to json file at: %s", listExportServicesPath)
		}
	}
	utils.Logf(string(servicesJson))
	return nil
}

func RunExportCommand(args *tf_export.ExportCommandArgs) (err error, status Status) {
	defer func() {
		if r := recover(); r != nil {
			utils.Logf("[ERROR] panic in RunExportCommand, exiting with status %v", StatusFail)
			debug.PrintStack()
			err = errors.New("[ERROR] panic in RunExportCommand: unknown error occurred in export")
			status = StatusFail
		}
	}()
	tf_export.ResourcesMap = tf_provider.ResourcesMap()
	tf_export.DatasourcesMap = tf_provider.DataSourcesMap()

	if err := args.Validate(); err != nil {
		return err, StatusFail
	}

	tf_export.TfHclVersionvar = *args.TFVersion

	r := &schema.Resource{
		Schema: tf_provider.SchemaMap(),
	}
	d := r.Data(nil)

	err = readEnvironmentVars(d)
	if err != nil {
		utils.Logln(err.Error())
		return err, StatusFail
	}

	clients, err := getExportConfigVar(d)
	if err != nil {
		utils.Logln(err.Error())
		return err, StatusFail
	}

	if args.CompartmentName != nil && *args.CompartmentName != "" {
		var err error
		args.CompartmentId, err = resolveCompartmentId(clients.(*tf_client.OracleClients), args.CompartmentName)
		if err != nil {
			utils.Logln(err.Error())
			return err, StatusFail
		}
	}

	/* Getting the tenancy ocid from env var export_tenancy_id, if not specified get customer tenancy ocid from configuration provider */
	tenancyOcid := getEnvSettingWithBlankDefaultVar("export_tenancy_id")

	if tenancyOcid == "" {
		/* Keep the tenancy lookup for backward compatibility */
		if args.CompartmentId != nil && *args.CompartmentId != "" {
			tenancyOcid, err = getTenancyOcidFromCompartment(clients.(*tf_client.OracleClients), *args.CompartmentId)
			if err != nil {
				utils.Logln(err.Error())
				return err, StatusFail
			}
		} else {
			// If compartment ocid not provided in arguments, get it from configuration provider
			tenancyId, exists := clients.(*tf_client.OracleClients).Configuration["tenancy_ocid"]
			if !exists {
				return fmt.Errorf("[ERROR] could not get a tenancy OCID during initialization"), StatusFail
			}
			tenancyOcid = tenancyId
		}

	}

	sem = make(chan struct{}, args.Parallelism)

	// maximum goroutines for finding resources in a step at base level.
	// Value of MaxParallelFindResources is decided based on perf analysis and experiments.
	// Note: SubResources will be discovered sequentially.
	numCPU := runtime.NumCPU()
	MaxParallelFindResource = numCPU * 4
	// max parallel chunks for state genation that can be executed in parallel
	MaxParallelChunks = numCPU
	utils.Debugf("[INFO] Setting MaxParalleFindResources=%d, MaxParallelChunks=%d", MaxParallelFindResource, MaxParallelChunks)

	ctx, err := createResourceDiscoveryContext(clients.(*tf_client.OracleClients), args, tenancyOcid)
	ctx.Filters = args.Filters

	if err != nil {
		utils.Logln(err.Error())
		return err, StatusFail
	}
	args.FinalizeServices(ctx)

	/*
		Setting retry timeout to a lower value for resource discovery
		This is done to handle the 404 and 500 errors in case
		any resource is unavailable in a region or in case the service is down
		The time out value is configurable from export command
	*/
	if args.RetryTimeout != nil && *args.RetryTimeout != "" {
		tfresource.LongRetryTime = *tfresource.GetTimeoutDuration(*args.RetryTimeout)
		tfresource.ShortRetryTime = tfresource.LongRetryTime
	}

	utils.Logf("[INFO] resource discovery retry timeout duration set to %v", tfresource.ShortRetryTime)

	if err := runExportCommand(ctx); err != nil {
		utils.Logln(err.Error())
		return err, StatusFail
	}
	if len(ctx.ErrorList.Errors) > 0 {
		// If the errors were from discovery of resources return partial success status
		error, status := getListOfNotDiscoveredResources(ctx)
		return error, status
	}
	return nil, StatusSuccess
}

func getListOfNotDiscoveredResources(ctx *tf_export.ResourceDiscoveryContext) (error, Status) {
	notDiscoveredParentResources, notDiscoveredChildResources := ctx.PrintErrors()
	var notDiscoveredResources []string
	var notDiscoveredError tf_export.ResourceDiscoveryCustomError

	notDiscoveredResources = append(notDiscoveredParentResources, notDiscoveredChildResources...) // Not discovered resources eg.Parent resources + Child Resources

	notDiscoveredError = tf_export.ResourceDiscoveryCustomError{
		TypeOfError: tf_export.PartiallyResourcesDiscoveredError,
		Message:     errors.New(strings.Join(notDiscoveredResources, ",")),
	}

	return notDiscoveredError.Error(), StatusPartialSuccess
}

func getExportConfig(d *schema.ResourceData) (interface{}, error) {
	clients := &tf_client.OracleClients{
		SdkClientMap:  make(map[string]interface{}, len(tf_client.OracleClientRegistrationsVar.RegisteredClients)),
		Configuration: make(map[string]string),
	}

	userAgentString := fmt.Sprintf(globalvar.ExportUserAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, globalvar.Version)
	httpClient := tf_provider.BuildHttpClient()

	sdkConfigProvider, err := tfProviderGetSdkConfigProvider(d, clients)
	if err != nil {
		return nil, err
	}
	exportConfigProvider = sdkConfigProvider

	// Note: In case of Instance Principal auth, the TenancyOCID will return
	// the ocid for the tenancy for the compute instance and not the one for the customer
	clients.Configuration["tenancy_ocid"], err = sdkConfigProviderTenancyOCIDVar(sdkConfigProvider)
	if err != nil {
		return nil, err
	}

	// beware: global variable `configureClient` set here--used elsewhere outside this execution path
	configureClientLocal, err := tfProviderBuildConfigureClientFn(sdkConfigProvider, httpClient)
	if err != nil {
		return nil, err
	}

	configureClientWithUserAgent := func(client *oci_common.BaseClient) error {
		if err := configureClientLocal(client); err != nil {
			return err
		}
		client.UserAgent = userAgentString
		return nil
	}
	// beware: global variable `configureClient` set here--used elsewhere outside this execution path
	tf_client.ConfigureClientVar = configureClientWithUserAgent
	err = createSDKClientsVar(clients, sdkConfigProvider, configureClientWithUserAgent)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func runExportCommand(ctx *tf_export.ResourceDiscoveryContext) error {
	utils.Logf("[INFO] Running export command\n")
	utils.Logf("[INFO] parallelism: %d", ctx.Parallelism)
	defer ctx.PrintSummary()
	exportStart := time.Now()
	defer elapsed("entire export command", nil, 0)()
	steps, err := getDiscoverResourceSteps(ctx)
	if err != nil {
		return err
	}
	discoveryStart := time.Now()
	var discoverWg sync.WaitGroup
	discoverWg.Add(len(steps))
	for i, step := range steps {

		sem <- struct{}{}

		go func(i int, step resourceDiscoveryStep) {
			utils.Debugf("[DEBUG] discover: Running step %d, %s", i, step.getBaseStep().name)
			defer elapsed(fmt.Sprintf("time taken in discovering resources for step %d", i), step.getBaseStep(), Discovery)()
			defer func() {
				<-sem
				if r := recover(); r != nil {
					utils.Logf("[ERROR] panic in discover goroutine")
					utils.Logf("[ERROR] panic in discover goroutine")
					debug.PrintStack()
				}
				utils.Debugf("[DEBUG] discoverWg done: step %d", i)
				discoverWg.Done()
			}()

			err := step.discover()
			if err != nil {
				// All errors in discover are added to the ctx.errorList
				utils.Debugf("[ERROR] error occurred while discovering resources for step %d", i)
				utils.Logf("[ERROR] error occurred while discovering resources: %s", err.Error())
				return
			}
			// Cull any references from the ref map that contain omitted resources
			// This is to avoid omitted resources from being referenced in generated configs
			for _, omittedResource := range step.getOmittedResources() {
				for key, reference := range tf_export.ReferenceMap {
					if strings.Contains(reference, omittedResource.GetTerraformReference()) {
						// refactor referenceMap to data structure with lock and methods to modify
						tf_export.RefMapLock.Lock()
						delete(tf_export.ReferenceMap, key)
						tf_export.RefMapLock.Unlock()
					}
				}
			}

			utils.Debugf("[DEBUG] discover: Completed step %d, %s", i, step.getBaseStep().name)
			utils.Debugf("[DEBUG] discovered %d resources for step %d", len(step.getDiscoveredResources()), i)
		}(i, step)

	}

	// Wait for all steps to complete discovery
	discoverWg.Wait()
	totalDiscoveryTime := time.Since(discoveryStart)
	utils.Debugf("discovering resources for all services took %v\n", totalDiscoveryTime)
	ctx.TimeTakenToDiscover = totalDiscoveryTime
	utils.Debug("[DEBUG] ~~~~~~ discover steps completed ~~~~~~")

	if ctx.GenerateState {
		stateStart := time.Now()
		// Run import commands
		if ctx.Parallelism > 1 {
			utils.Debug("[DEBUG] Generating state in parallel")
			if err := generateStateParallel(ctx, steps); err != nil {
				return err
			}
		} else {
			utils.Debug("[DEBUG] Generating state sequentially")
			if err := generateState(ctx, steps); err != nil {
				return err
			}
		}

		// remove invalid references from referenceMap for the resources with import error
		if ctx.IsImportError {
			// lock not required for referenceMap as only 1 thread is running at this point
			deleteInvalidReferences(tf_export.ReferenceMap, ctx.DiscoveredResources)
		}
		timeForStateGeneration := time.Since(stateStart)
		utils.Debugf("[DEBUG] state generation took %v\n", timeForStateGeneration)
		ctx.TimeTakenToGenerateState = timeForStateGeneration
	}

	// Reset discovered resources if already set by writeTmpConfigurationForImport
	ctx.DiscoveredResources = make([]*tf_export.OCIResource, 0)

	/*
		sem allows number of steps equals to arg.Parallelism to execute in parallel.
		arg.Parallelism is very less compare to total number of steps
		So, errorChannel should be atleast equals to number of steps to allow all steps getting chance to execute.
		in case of multiple steps fails, steps will be blocked to write errors and pending steps will wait to acquire sem which leads to deadlock situation
	*/
	errorChannel := make(chan error, len(steps))
	var configWg sync.WaitGroup
	configWg.Add(len(steps))

	wgDone := make(chan bool)

	// Write configuration for imported resources
	configStart := time.Now()
	for i, step := range steps {

		sem <- struct{}{}
		go func(i int, step resourceDiscoveryStep) {
			utils.Debugf("[DEBUG] writeConfiguration: Running step %d", i)
			defer func() {
				<-sem
				if r := recover(); r != nil {
					utils.Logf("[ERROR] panic in writeConfiguration goroutine")
					debug.PrintStack()
				}
				utils.Debugf("[DEBUG] configWg done: step %d", i)
				configWg.Done()
			}()
			if err := step.writeConfiguration(); err != nil {
				errorChannel <- fmt.Errorf("[ERROR] error writing final configuration for resources found: %s", err.Error())
			}

			utils.Debugf("[DEBUG] writeConfiguration: Completed step %d", i)
		}(i, step)
	}

	// goroutine to wait until configWg is done
	go func() {
		utils.Debugf("[DEBUG] waiting for all configWg goroutines to finish...")
		configWg.Wait()
		close(wgDone)
	}()

	// Wait until either configWg is done or an error is received through the errorChannel
	select {
	case <-wgDone:
		utils.Debugf("[DEBUG] ~~~~~~ writeConfiguration steps completed ~~~~~~")
		utils.Debugf("[DEBUG] writing config took %v\n", time.Since(configStart))
		break
	case errs := <-errorChannel:
		close(errorChannel)
		for i := 0; i < len(errorChannel); i++ {
			errs = multierror.Append(errs, <-errorChannel)
		}
		utils.Logf("[ERROR] error writing final configuration for resources found: %s", errs.Error())
		return errs
	}

	region, err := exportConfigProvider.Region()
	if err != nil {
		return err
	}
	tf_export.Vars["region"] = fmt.Sprintf("\"%s\"", region)

	if err := generateProviderFile(ctx.OutputDir); err != nil {
		return err
	}

	if err := generateVarsFile(tf_export.Vars, ctx.OutputDir); err != nil {
		return err
	}

	if tf_export.IsMissingRequiredAttributes {
		ctx.SummaryStatements = append(ctx.SummaryStatements, "")
		ctx.SummaryStatements = append(ctx.SummaryStatements, globalvar.MissingRequiredAttributeWarning)
		ctx.SummaryStatements = append(ctx.SummaryStatements, "Missing required attributes:")
		for key, value := range ctx.MissingAttributesPerResource {
			ctx.SummaryStatements = append(ctx.SummaryStatements, fmt.Sprintf("%s: %s", key, strings.Join(value, ",")))
		}
	}
	ctx.TimeTakenForEntireExport = time.Since(exportStart)
	ctx.PostValidate()
	return nil
}

/*
generateStateParallel is used if value of parallelism arg > 1
- writes temp config for the discovered resources e.g. `resource_type resource_name {}` in order to run import
- writes temp state file for each service in parallel by running import for each of the found resources
- finally it merges all the state files generated into one state file using json merge
*/
func generateStateParallel(ctx *tf_export.ResourceDiscoveryContext, steps []resourceDiscoveryStep) error {

	// isInitDone is to make sure that multiple threads do not call terraform init
	utils.Debugf("[DEBUG] Reset isInitDone")
	isInitDone = false
	// Cleanup the temporary state files created for each input service
	defer cleanupTempStateFiles(ctx)
	defer elapsed("generating state in parallel", nil, 0)()

	/*
		sem allows number of steps equals to arg.Parallelism to execute in parallel.
		arg.Parallelism is very less compare to total number of steps
		So, errorChannel should be atleast equals to number of steps to allow all steps getting chance to execute.
		in case of multiple steps fails, steps will be blocked to write errors and pending steps will wait to acquire sem which leads to deadlock situation
	*/
	errorChannel := make(chan error, len(steps))
	var stateWg sync.WaitGroup
	wgDone := make(chan bool)

	stateWg.Add(len(steps))

	for i, step := range steps {
		if len(step.getDiscoveredResources()) == 0 {
			stateWg.Done()
			continue
		}

		sem <- struct{}{}

		go func(i int, step resourceDiscoveryStep) {
			utils.Debugf("[DEBUG] writing temp config and state: Running step %d-%s", i, step.getBaseStep().name)
			defer elapsed(fmt.Sprintf("time taken by step %s to generate state", fmt.Sprint(i)), step.getBaseStep(), GeneratingState)()
			defer func() {
				<-sem
				if r := recover(); r != nil {
					utils.Logf("[ERROR] panic in writing temp config and state goroutine")
					debug.PrintStack()
				}
				utils.Debugf("[DEBUG] stateWg done: step %d", i)
				stateWg.Done()
			}()

			/* Generate temporary HCL configs from all discovered resources to run import
			   Final configuration will be generated after import so that we can exclude the resources for which import failed
			   and also remove the references to failed resources from the referenceMap */
			var err error = nil
			if err = step.writeTmpConfigurationForImport(); err != nil {
				errorChannel <- fmt.Errorf("[ERROR] error while writing temp config for resources found in step %d: %s", i, err.Error())
			}

			// Write temp state file for each service, this step will import resources into a separate state file for each service in parallel
			// If writing temp configuration thrown error, won't attempt to write temp state
			if err == nil {
				if err = step.writeTmpState(); err != nil {
					tfError := tf_export.ResourceDiscoveryCustomError{
						TypeOfError: tf_export.WriteTmpStateError,
						Message:     errors.New(tf_export.WriteTmpStateErrorMessage),
						Suggestion:  tf_export.WriteTmpStateErrorSuggestion}

					errorChannel <- tfError.Error()
				}
			}

			utils.Debugf("writing temp config and state: Completed step %d", i)
		}(i, step)
	}

	// goroutine to wait until stateWg is done
	go func() {
		utils.Debugf("[DEBUG] waiting for all stateWg threads to finish...")
		stateWg.Wait()
		close(wgDone)
	}()

	// Wait until either stateWg is done or an error is received through the errorChannel
	select {
	case <-wgDone:
		utils.Debugf("[DEBUG] ~~~~~~ writing temp config and state steps completed ~~~~~~")
		break
	case errs := <-errorChannel:
		close(errorChannel)
		for i := 0; i < len(errorChannel); i++ {
			errs = multierror.Append(errs, <-errorChannel)
		}
		utils.Logf("[ERROR] error writing temp config and state: %s", errs.Error())
		return errs
	}

	// Generate final state by merging state json generated for all services
	for _, step := range steps {
		if err := step.mergeGeneratedStateFile(); err != nil {
			return fmt.Errorf("[ERROR] Resource discovery failed. Error merging generated states: %s", err.Error())
		}
	}

	if ctx.State == nil {
		utils.Logf("[INFO] ~~~~~~ no resources were imported to the state file ~~~~~~")
		return nil
	}

	// Create final state file to write state
	stateOutputFile := fmt.Sprintf("%s%s%s", *ctx.OutputDir, string(os.PathSeparator), globalvar.DefaultStateFilename)

	f, err := os.OpenFile(stateOutputFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to Create state file %s: %s", stateOutputFile, err.Error())
	}

	// write generate state to file
	stateBytes, _ := json.MarshalIndent(ctx.State, "", "\t")
	if err := ioutil.WriteFile(stateOutputFile, stateBytes, 0644); err != nil {
		return fmt.Errorf("[ERROR] error writing state file at %s: %s", stateOutputFile, err.Error())
	} else {
		utils.Logf("[INFO] state written to file at: %s", stateOutputFile)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("[WARN] unable to close state file: %s", err.Error())
	}

	// Run init in output_path to validate generated state
	backgroundCtx := context.Background()

	var initArgs []tfexec.InitOption

	if ctx.TerraformProviderBinaryPath != "" {
		utils.Logf("[INFO] plugin dir set to: '%s'", ctx.TerraformProviderBinaryPath)
		initArgs = append(initArgs, tfexec.PluginDir(ctx.TerraformProviderBinaryPath))
	}

	if err := ctx.Terraform.Init(backgroundCtx, initArgs...); err != nil {
		return err
	}
	return nil
}

/*
generateState is used if value of parallelism arg == 1
- writes temp config for the discovered resources e.g. `resource_type resource_name {}` in order to run import
- writes state file by running import for each of the found resources
*/
func generateState(ctx *tf_export.ResourceDiscoveryContext, steps []resourceDiscoveryStep) error {

	utils.Debugln("[DEBUG] writing temp config files for import")
	for _, step := range steps {

		/* Generate temporary HCL configs from all discovered resources to run import
		   Final configuration will be generated after import so that we can exclude the resources for which import failed
		   and also remove the references to failed resources from the referenceMap */
		if err := step.writeTmpConfigurationForImport(); err != nil {
			return err
		}
	}

	// Run init command
	backgroundCtx := context.Background()

	var initArgs []tfexec.InitOption

	if ctx.TerraformProviderBinaryPath != "" {
		utils.Logf("[INFO] plugin dir set to: '%s'", ctx.TerraformProviderBinaryPath)
		initArgs = append(initArgs, tfexec.PluginDir(ctx.TerraformProviderBinaryPath))
	}
	if err := terraformInitVar(ctx, backgroundCtx, initArgs); err != nil {
		return err
	}

	stateOutputFile := fmt.Sprintf("%s%s%s", *ctx.OutputDir, string(os.PathSeparator), globalvar.DefaultStateFilename)
	tmpStateOutputFile := fmt.Sprintf("%s%s%s", *ctx.OutputDir, string(os.PathSeparator), globalvar.DefaultTmpStateFile)
	if err := os.RemoveAll(tmpStateOutputFile); err != nil {
		utils.Logf("[WARN] unable to delete existing tmp state file %s", tmpStateOutputFile)
		return err
	}

	// Run import for all resources
	for _, resource := range ctx.DiscoveredResources {
		importResource(ctx, resource, tmpStateOutputFile)
	}

	if _, err := os.Stat(tmpStateOutputFile); !os.IsNotExist(err) {
		if err := os.Rename(tmpStateOutputFile, stateOutputFile); err != nil {
			return err
		}
	}

	return nil
}

/*
importResource runs terraform import for a given resource using Terraform exec and writes to the given state file
*/
func importResource(ctx *tf_export.ResourceDiscoveryContext, resource *tf_export.OCIResource, tmpStateOutputFile string) {
	utils.Logf("[INFO] ===> Importing resource '%s'", resource.GetTerraformReference())
	utils.Debugf("[DEBUG] ===> Importing resource '%s'", resource.GetTerraformReference())

	resourceDefinition, exists := tf_export.ResourcesMap[resource.TerraformClass]
	if !exists {
		utils.Logf("[INFO] skip importing '%s' since it is not a Terraform OCI resource", resource.GetTerraformReference())
		utils.Debugf("[DEBUG] skip importing '%s' since it is not a Terraform OCI resource", resource.GetTerraformReference())
		return
	}

	if resourceDefinition.Importer == nil {
		utils.Logf("[WARN] unable to import '%s' because import is not supported for '%s'", resource.GetTerraformReference(), resource.TerraformClass)
		return
	}

	importId := resource.ImportId
	if len(importId) == 0 {
		importId = resource.Id
	}

	importArgs := []tfexec.ImportOption{
		tfexecConfigVar(*ctx.OutputDir),
		tfexecStateVar(tmpStateOutputFile),
	}
	if importErr := ctxTerraformImportVar(ctx, context.Background(), resource.GetTerraformReference(), importId, importArgs...); importErr != nil {
		utils.Logf("[ERROR] terraform import command failed for resource '%s' at id '%s': %s", resource.GetTerraformReference(), importId, importErr.Error())
		if strings.Contains(importErr.Error(), "Missing source file for snapshot") {
			utils.ListCurrentDirectory()
		}
		// mark resource as errored so that it can be skipped while writing configurations
		resource.IsErrorResource = true

		ctx.CtxLock.Lock()
		ctx.IsImportError = true
		ctx.CtxLock.Unlock()

		err := fmt.Errorf("[ERROR] terraform import command failed for resource '%s' at id '%s': %s Any references to this resource have been replaced with hard coded values in generated configurations", resource.GetTerraformReference(), importId, importErr.Error())

		var rdError *tf_export.ResourceDiscoveryError
		if ctx.TargetSpecificResources {
			rdError = &tf_export.ResourceDiscoveryError{
				ResourceType:   resource.TerraformClass,
				ParentResource: "",
				Error:          err,
				ResourceGraph:  nil}
		} else {
			rdError = &tf_export.ResourceDiscoveryError{
				ResourceType:   resource.TerraformClass,
				ParentResource: resource.Parent.TerraformName,
				Error:          err,
				ResourceGraph:  nil,
			}
		}
		ctx.AddErrorToList(rdError)

	}
}

func getDiscoverResourceSteps(ctx *tf_export.ResourceDiscoveryContext) ([]resourceDiscoveryStep, error) {
	if !ctx.TargetSpecificResources {
		return getDiscoverResourceWithGraphSteps(ctx)
	}

	result := make([]resourceDiscoveryStep, 1)
	result[0] = &resourceDiscoveryWithTargetIds{
		resourceDiscoveryBaseStep: resourceDiscoveryBaseStep{
			ctx:                 ctx,
			name:                "resources",
			discoveredResources: []*tf_export.OCIResource{},
			omittedResources:    []*tf_export.OCIResource{},
		},
	}
	return result, nil
}

func getDiscoverResourceWithGraphSteps(ctx *tf_export.ResourceDiscoveryContext) ([]resourceDiscoveryStep, error) {
	defer elapsed("Building resource discovery graph", nil, 0)()
	if ctx.CompartmentId == nil || *ctx.CompartmentId == "" {
		*ctx.CompartmentId = ctx.TenancyOcid
	}
	var result []resourceDiscoveryStep

	// Discover tenancy scope resources only if compartmentId is tenancy ocid
	if *ctx.CompartmentId == ctx.TenancyOcid {
		tenancyResource := &tf_export.OCIResource{
			CompartmentId: ctx.TenancyOcid,
			TerraformResource: tf_export.TerraformResource{
				Id:             ctx.TenancyOcid,
				TerraformClass: "oci_identity_tenancy",
				TerraformName:  "export",
			},
		}

		for _, mode := range ctx.Services {
			if resourceGraph, exists := tf_export.TenancyResourceGraphs[mode]; exists {
				result = append(result, &resourceDiscoveryWithGraph{
					root:                      tenancyResource,
					resourceGraph:             resourceGraph,
					resourceDiscoveryBaseStep: resourceDiscoveryBaseStep{name: mode, ctx: ctx},
				})

				tf_export.Vars["tenancy_ocid"] = fmt.Sprintf("\"%s\"", ctx.TenancyOcid)
				tf_export.ReferenceMap[ctx.TenancyOcid] = tf_export.TfHclVersionvar.GetVarHclString("tenancy_ocid")
			}
		}
	}

	compartmentResource := &tf_export.OCIResource{
		CompartmentId: *ctx.CompartmentId,
		TerraformResource: tf_export.TerraformResource{
			Id:             *ctx.CompartmentId,
			TerraformClass: "oci_identity_compartment",
			TerraformName:  "export",
		},
	}

	for _, mode := range ctx.Services {
		if resourceGraph, exists := tf_export.CompartmentResourceGraphs[mode]; exists {
			result = append(result, &resourceDiscoveryWithGraph{
				root:                      compartmentResource,
				resourceGraph:             resourceGraph,
				resourceDiscoveryBaseStep: resourceDiscoveryBaseStep{name: mode, ctx: ctx},
			})

			tf_export.Vars["compartment_ocid"] = fmt.Sprintf("\"%s\"", *ctx.CompartmentId)
			tf_export.ReferenceMap[*ctx.CompartmentId] = tf_export.TfHclVersionvar.GetVarHclString("compartment_ocid")
		}
	}

	return result, nil
}

func runFilters(resources []*tf_export.OCIResource, filters []tf_export.ResourceFilter) ([]*tf_export.OCIResource, error) {

	if resources == nil || filters == nil || len(filters) == 0 {
		return resources, nil
	}

	results := make([]*tf_export.OCIResource, 0)
	for _, resource := range resources {
		includeResource := true
		for _, filter := range filters {
			if !filter.Filter(resource) {
				includeResource = false
				break
			}
		}
		// include resource if it satisfied filter criteria
		if includeResource {
			results = append(results, resource)
		}
	}
	return results, nil
}

func findResources(ctx *tf_export.ResourceDiscoveryContext, root *tf_export.OCIResource, resourceGraph tf_export.TerraformResourceGraph, discoveryParallelism bool) (foundResources []*tf_export.OCIResource, err error) {
	// findResources will never return error, it will add the errors encountered to the errorList and print those after the discovery finishes
	// If find resources needs to fail in some scenario, this func needs to be modified to return error instead of continuing discovery
	// Errors so far are API errors or the errors when service/feature is not available
	foundResources = []*tf_export.OCIResource{}
	var foundResourcesLock sync.Mutex

	childResourceTypes, exists := resourceGraph[root.TerraformClass]
	if !exists {
		return foundResources, nil
	}

	utils.Logf("[INFO] resource discovery: visiting %s\n", root.GetTerraformReference())
	utils.Debugf("[DEBUG] resource discovery: visiting %s\n", root.GetTerraformReference())

	utils.Logf("[INFO] number of child resource types for %s: %d\n", root.GetTerraformReference(), len(childResourceTypes))

	findResourcesStart := time.Now()
	var findResourcesWg sync.WaitGroup
	findResourcesWg.Add(len(childResourceTypes))

	// setting parallelism argument false for subResources to control concurrency and thrashing
	var ch chan struct{}
	if discoveryParallelism == true {
		ch = make(chan struct{}, MaxParallelFindResource)
		// setting parallelism argument false for subResources to control concurrency and thrashing
		discoveryParallelism = false
	} else {
		ch = make(chan struct{}, 1)
	}

	for i, childType := range childResourceTypes {

		ch <- struct{}{}

		go func(i int, childType tf_export.TerraformResourceAssociation) {
			utils.Debugf("[DEBUG] findResources: Starting to find resources for resource type index: %d, type: %s", i, childType.ResourceClass)

			defer func(tfMeta *tf_export.TerraformResourceAssociation, err *error) {
				<-ch
				if r := recover(); r != nil {
					utils.Logf("[WARN] recovered from panic in findResourcesGeneric for resource: %s \n continuing discovery...", tfMeta.ResourceClass)
					returnErr := fmt.Errorf("panic in findResourcesGeneric for resource %s", tfMeta.ResourceClass)
					*err = returnErr
					debug.PrintStack()
				}
				utils.Debugf("[DEBUG] findResourcesWg done, resource type index: %d, type %s", i, childType.ResourceClass)
				findResourcesWg.Done()
			}(&childType, &err)

			findResourceFn := tf_export.FindResourcesGeneric
			if childType.FindResourcesOverrideFn != nil {
				findResourceFn = childType.FindResourcesOverrideFn
			}
			results, err := findResourceFn(ctx, &childType, root, &resourceGraph)
			if err != nil {
				// add error to the errorList and continue discovering rest of the resources
				rdError := &tf_export.ResourceDiscoveryError{
					ResourceType:   childType.ResourceClass,
					ParentResource: root.TerraformName,
					Error:          err,
					ResourceGraph:  &resourceGraph}
				ctx.AddErrorToList(rdError)
				return
			}

			if childType.ProcessDiscoveredResourcesFn != nil {
				results, err = childType.ProcessDiscoveredResourcesFn(ctx, results)
				if err != nil {
					// add error to the errorList and continue discovering rest of the resources
					rdError := &tf_export.ResourceDiscoveryError{
						ResourceType:   childType.ResourceClass,
						ParentResource: root.TerraformName,
						Error:          err,
						ResourceGraph:  &resourceGraph}
					ctx.AddErrorToList(rdError)
					return
				}
			}
			foundResources = append(foundResources, results...)

			for _, resource := range results {
				//referenceMap[resource.id] = resource.getHclReferenceIdString()
				if ctx.ExpectedResourceIds != nil && len(ctx.ExpectedResourceIds) > 0 {
					if _, shouldExport := ctx.ExpectedResourceIds[resource.Id]; shouldExport {
						resource.OmitFromExport = false
						ctx.CtxLock.Lock()
						ctx.ExpectedResourceIds[resource.Id] = true
						ctx.CtxLock.Unlock()
					} else {
						if resource.Parent != nil && ctx.ExpectedResourceIds[resource.Parent.Id] {
							ctx.ExpectedResourceIds[resource.Id] = true
						} else {
							utils.Debugf("[DEBUG] resource %s is not exportable %v, childType %v", resource.GetTerraformReference(), childType.AlwaysExportable, childType.ResourceClass)
							resource.OmitFromExport = !childType.AlwaysExportable
						}
					}

				}

				subResources, err := findResources(ctx, resource, resourceGraph, discoveryParallelism)
				if err != nil {
					continue
				}
				foundResourcesLock.Lock()
				foundResources = append(foundResources, subResources...)
				foundResourcesLock.Unlock()
			}
			utils.Debugf("[DEBUG] findResources: Completed for resource type index: %d, childType: %s", i, childType.ResourceClass)
		}(i, childType)
	}

	// Wait for all steps to complete findResources
	findResourcesWg.Wait()
	totalFindResourcesTime := time.Since(findResourcesStart)
	utils.Debugf("finding resources for %s took %v\n", root.GetTerraformReference(), totalFindResourcesTime)

	// create copies of filters so that in each thread, they are not shared
	// since number of filters is expected to be less than the number of threads running in parallel, this is a cost effective approach than locking
	filtersCopies := make([]tf_export.ResourceFilter, 0)
	if ctx != nil && ctx.ExportCommandArgs != nil && ctx.Filters != nil {
		filtersCopies, err = tf_export.GetFiltersDeepCopy(ctx.Filters)
		if err != nil {
			return nil, err
		}
	}

	// run filters on output of foundResources
	// we can later optimize to run certain filters even before discovery but for now run filter upon discovery
	foundResources, err = runFilters(foundResources, filtersCopies)
	if err != nil {
		return nil, err
	}
	return foundResources, nil
}

func generateVarsFile(vars map[string]string, outputDir *string) error {
	varsTmpFile := fmt.Sprintf("%s%s%s.tmp", *outputDir, string(os.PathSeparator), globalvar.VarsFile)
	varsOutputFile := fmt.Sprintf("%s%s%s", *outputDir, string(os.PathSeparator), globalvar.VarsFile)
	file, err := os.OpenFile(varsTmpFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	for variable, defaultVal := range vars {
		if defaultVal != "" {
			_, _ = file.WriteString(fmt.Sprintf("variable %s { default = %s }\n", variable, defaultVal))
		} else {
			_, _ = file.WriteString(fmt.Sprintf("variable %s {}\n", variable))
		}
	}

	if err := file.Close(); err != nil {
		return err
	}

	if err := os.Rename(varsTmpFile, varsOutputFile); err != nil {
		return err
	}

	return nil
}

func generateProviderFile(outputDir *string) error {
	providerTmpFile := fmt.Sprintf("%s%s%s.tmp", *outputDir, string(os.PathSeparator), globalvar.ProviderFile)
	providerOutputFile := fmt.Sprintf("%s%s%s", *outputDir, string(os.PathSeparator), globalvar.ProviderFile)
	file, err := os.OpenFile(providerTmpFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	_, err = file.WriteString(fmt.Sprintf("provider oci {\n\tregion = %s\n}\n", tf_export.TfHclVersionvar.GetVarHclString("region")))
	if err != nil {
		_ = file.Close()
		return err
	}

	if fErr := file.Close(); fErr != nil {
		return fErr
	}

	if err := os.Rename(providerTmpFile, providerOutputFile); err != nil {
		return err
	}

	return nil
}

//func getOciResource(d *schema.ResourceData, resourceSchema map[string]*schema.Schema, compartmentId string, resourceHint *tf_export.TerraformResourceHints, resourceId string) (*tf_export.OCIResource, error) {
//	resourceMap, err := tf_export.ConvertDatasourceItemToMap(d, "", resourceSchema)
//	if err != nil {
//		return nil, err
//	}
//
//	resource := &tf_export.OCIResource{
//		CompartmentId:    compartmentId,
//		SourceAttributes: resourceMap,
//		RawResource:      d,
//		TerraformResource: tf_export.TerraformResource{
//			TerraformClass:    resourceHint.ResourceClass,
//			TerraformTypeInfo: resourceHint,
//		},
//		GetHclStringFn: tf_export.GetHclStringFromGenericMap,
//	}
//
//	if resourceId != "" {
//		resource.Id = resourceId
//	}
//
//	if resource.Id == "" {
//		resource.Id = d.Id()
//	}
//
//	return resource, nil
//}

func resolveCompartmentId(clients *tf_client.OracleClients, compartmentName *string) (*string, error) {
	req := oci_identity.ListCompartmentsRequest{}

	rootCompartment, err := exportConfigProvider.TenancyOCID()
	if err != nil {
		return nil, err
	}
	req.CompartmentId = &rootCompartment

	recursiveSearch := true
	req.CompartmentIdInSubtree = &recursiveSearch

	for {
		resp, err := identityClientListCompartmentsVar(clients, req)
		if err != nil {
			return nil, err
		}

		for _, compartment := range resp.Items {
			if compartment.Name != nil && *compartment.Name == *compartmentName {
				utils.Logf("[INFO] resolved compartment name '%s' to compartment id '%s'", *compartmentName, *compartment.Id)
				return compartment.Id, nil
			}
		}

		if resp.OpcNextPage == nil {
			break
		}
		req.Page = resp.OpcNextPage
	}

	return nil, fmt.Errorf("[ERROR] Could not find a compartment named '%s' in your tenancy", *compartmentName)
}

func readEnvironmentVars(d *schema.ResourceData) error {

	if err := d.Set(globalvar.AuthAttrName, getProviderEnvSettingWithDefaultVar(globalvar.AuthAttrName, globalvar.AuthAPIKeySetting)); err != nil {
		return err
	}
	if err := d.Set(globalvar.ConfigFileProfileAttrName, getProviderEnvSettingWithDefaultVar(globalvar.ConfigFileProfileAttrName, "")); err != nil {
		return err
	}
	if region := utils.GetProviderEnvSettingWithDefault(globalvar.RegionAttrName, ""); region != "" {
		if err := d.Set(globalvar.RegionAttrName, region); err != nil {
			return err
		}
	}

	if tenancyOcid := getProviderEnvSettingWithDefaultVar(globalvar.TenancyOcidAttrName, ""); tenancyOcid != "" {
		if err := d.Set(globalvar.TenancyOcidAttrName, tenancyOcid); err != nil {
			return err
		}
	}

	if userOcid := getProviderEnvSettingWithDefaultVar(globalvar.UserOcidAttrName, ""); userOcid != "" {
		if err := d.Set(globalvar.UserOcidAttrName, userOcid); err != nil {
			return err
		}
	}
	if fingerprint := getProviderEnvSettingWithDefaultVar(globalvar.FingerprintAttrName, ""); fingerprint != "" {
		if err := d.Set(globalvar.FingerprintAttrName, fingerprint); err != nil {
			return err
		}
	}
	if privateKey := getProviderEnvSettingWithDefaultVar(globalvar.PrivateKeyAttrName, ""); privateKey != "" {
		if err := d.Set(globalvar.PrivateKeyAttrName, privateKey); err != nil {
			return err
		}
	}
	if privateKeyPath := getProviderEnvSettingWithDefaultVar(globalvar.PrivateKeyPathAttrName, ""); privateKeyPath != "" {
		if err := d.Set(globalvar.PrivateKeyPathAttrName, privateKeyPath); err != nil {
			return err
		}
	}
	if privateKeyPassword := getProviderEnvSettingWithDefaultVar(globalvar.PrivateKeyPasswordAttrName, ""); privateKeyPassword != "" {
		if err := d.Set(globalvar.PrivateKeyPasswordAttrName, privateKeyPassword); err != nil {
			return err
		}
	}
	return nil
}

func getTenancyOcidFromCompartment(clients *tf_client.OracleClients, compartmentId string) (string, error) {

	for true {
		response, err := identityClientGetCompartmentVar(clients, oci_identity.GetCompartmentRequest{
			CompartmentId: &compartmentId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(true, "identity"),
			},
		})
		if err != nil {
			return "", fmt.Errorf("[ERROR] could not get tenancy ocid from compartment ocid %v", err)
		}
		if response.CompartmentId == nil {
			utils.Logf("[INFO] root compartment found %v", compartmentId)
			return *response.Id, nil
		}
		compartmentId = *response.CompartmentId
	}

	return "", fmt.Errorf("[ERROR] could not get tenancy ocid from compartment ocid")
}

func deleteInvalidReferences(referenceMap map[string]string, discoveredResources []*tf_export.OCIResource) {
	// intialize referenceResourceNameSet
	// This set contains unique terraform names for resource references
	if referenceResourceNameSet == nil {
		referenceResourceNameSet = make(map[string]bool)
		for _, value := range referenceMap {
			valueParts := strings.Split(value, ".")
			if len(valueParts) < 3 {
				continue
			}
			referenceResourceNameSet[valueParts[1]] = true
		}
	}
	if tf_export.FailedResourceReferenceSet == nil {
		tf_export.FailedResourceReferenceSet = make(map[string]bool)
	}

	for _, resource := range discoveredResources {

		// delete the entry if key is an OCID for a failed resource
		if resource.IsErrorResource {
			// store failed resource reference, will be used later to remove InterpolationString type values when generating config
			tf_export.FailedResourceReferenceSet[resource.GetTerraformReference()] = true
			if _, ok := referenceMap[resource.Id]; ok {
				delete(referenceMap, resource.Id)
			}

			// delete any entries that have references to a failed resource
			// e.g. oci_core_instance.instance1.volume_id should be replaced by volume ocid if instance import failed
			if ok := referenceResourceNameSet[resource.TerraformName]; ok {
				for key, value := range referenceMap {
					valueParts := strings.Split(value, ".")
					if len(valueParts) < 3 {
						continue
					}
					if valueParts[1] == resource.TerraformName {
						delete(referenceMap, key)
					}
				}
			}
		}
	}
}

/*
Initialize Terraform struct from executable provided
Terraform struct will later be copied to each resource discovery step for parallel runs
*/
func createTerraformStruct(args *tf_export.ExportCommandArgs) (*tfexec.Terraform, string, error) {

	utils.Logln("[INFO] validating Terraform CLI")
	var err error
	terraformBinPath := getEnvSettingWithBlankDefaultVar(globalvar.TerraformBinPathName)
	utils.Logf("terraform bin path --- %s", terraformBinPath)
	if terraformBinPath == "" {
		utils.Logf("terraform bin path --- %s", terraformBinPath)
		terraformBinPath, err = hcInstallerEnsureVar(hcinstall.NewInstaller(), context.Background(),
			[]src.Source{src.Findable(&fs.AnyVersion{Product: &product.Terraform})})
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] error finding terraform CLI, either specify the path to terraform CLI "+
				"including name using env var 'terraform_bin_path' or add terraform CLI to your system path: %s", err.Error())
		}
	} else {
		// Validate the path provided
		file, err := osStatvar(terraformBinPath)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR]: error verifying the terraform binary path provided %s: %s", terraformBinPath, err)
		}
		if isDirVar(file) {
			return nil, "", fmt.Errorf("[ERROR]: terraform CLI path provided is a directory: %s", terraformBinPath)
		}

	}

	// Initialize Terraform struct from executable provided
	// Setting the global var 'tf' here, will be later using while running terraform init and import
	tf, err := tfexec.NewTerraform(*args.OutputDir, terraformBinPath)
	if err != nil {
		return nil, terraformBinPath, err
	}

	// Set log path for TF Exec
	logPath := os.Getenv(globalvar.EnvLogFile)
	if logPath != "" {
		utils.Logf("[INFO] setting log path for Terraform exec to '%s'", logPath)
		if err := tf.SetLogPath(logPath); err != nil {
			return nil, terraformBinPath, err
		}
	}

	backgroundCtx := context.Background()

	// discard stdout to avoid showing tf version results in logs
	tf.SetStdout(ioutil.Discard)

	// validate Terraform CLI
	if tfVersion, _, err := tfVersionVar(tf, backgroundCtx); err != nil {
		return nil, terraformBinPath, fmt.Errorf("[ERROR] error verifying the terraform binary provided: %s", err)
	} else {
		utils.Debugf("[DEBUG] version %v", tfVersion)

		// check for tf_version and terraform CLI version so as to avoid
		// scenarios where config is not compatible with TF version of state file
		// version should be >= 0.12.0
		inputTfVersion := "v" + tfVersion.String()

		if semver.MajorMinor(inputTfVersion) == semver.MajorMinor("v"+string(tf_export.TfVersion11)) {
			return nil, terraformBinPath, fmt.Errorf("[ERROR] resource discovery does not support v0.11.* CLI, "+
				"please specify terraform CLI with version v0.12.*, terraform version provided: %s", tfVersion.String())
		}

		executableVersion := semver.MajorMinor(inputTfVersion)
		configVersion := semver.MajorMinor("v" + tf_export.TfHclVersionvar.ToString())

		if executableVersion < configVersion {
			return nil, terraformBinPath, fmt.Errorf("[ERROR] major and minor version of terraform CLI provided is not same as the generated configuration version, "+
				"configuration version: %s, terraform CLI version: %s, please provide CLI version >= %s ", tf_export.TfHclVersionvar.ToString(), tfVersion.String(), tf_export.TfHclVersionvar.ToString())
		}
	}
	// enable stdout again to show init and import output in logs
	tf.SetStdout(os.Stdout)

	return tf, terraformBinPath, nil
}

func cleanupTempStateFiles(ctx *tf_export.ResourceDiscoveryContext) {

	/* Clean up temp state files for individual services */
	if err := os.RemoveAll(fmt.Sprintf("%s%stmp%s", *ctx.OutputDir, string(os.PathSeparator), string(os.PathSeparator))); err != nil {
		utils.Logf("[ERROR] Error removing tmp state files: %s", err.Error())
	}
}

var (
	//tmpl template.Template = *template.New("tmpl")

	lineSeparator                    = "\n"
	getEnvSettingWithBlankDefaultVar = utils.GetEnvSettingWithBlankDefault
	getEnvSettingWithDefaultVar      = utils.GetEnvSettingWithDefault
	tfProviderConfigVar              = tf_provider.ProviderConfig
)
