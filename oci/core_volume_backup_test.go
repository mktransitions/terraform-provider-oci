// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/v25/common"

	oci_core "github.com/oracle/oci-go-sdk/v25/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VolumeBackupRequiredOnlyResource = VolumeBackupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", Required, Create, volumeBackupRepresentation)

	volumeBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"volume_id":      Representation{repType: Optional, create: `${oci_core_volume.test_volume.id}`},
		"filter":         RepresentationGroup{Required, volumeBackupDataSourceFilterRepresentation}}
	volumeBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_volume_backup.test_volume_backup.id}`}},
	}

	volumeBackupFromSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":            Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"source_volume_backup_id": Representation{repType: Optional, create: `${oci_core_volume_backup.test_volume_backup_copy.source_volume_backup_id}`},
		"state":                   Representation{repType: Optional, create: `AVAILABLE`},
		"filter":                  RepresentationGroup{Required, volumeBackupFromSourceDataSourceFilterRepresentation}}
	volumeBackupFromSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_volume_backup.test_volume_backup_copy.id}`}},
	}

	volumeBackupRepresentation = map[string]interface{}{
		"volume_id":     Representation{repType: Required, create: `${oci_core_volume.test_volume.id}`},
		"defined_tags":  Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags": Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"type":          Representation{repType: Optional, create: `FULL`},
	}
	volumeBackupWithSourceDetailsRepresentation = map[string]interface{}{
		"source_details": RepresentationGroup{Required, volumeBackupSourceDetailsRepresentation},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
	}

	volumeBackupId, volumeId                string
	volumeBackupSourceDetailsRepresentation = map[string]interface{}{}

	VolumeBackupResourceDependencies = VolumeResourceConfig
)

func TestCoreVolumeBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeBackupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_volume_backup.test_volume_backup"
	datasourceName := "data.oci_core_volume_backups.test_volume_backups"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeBackupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", Required, Create, volumeBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeBackupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", Optional, Create,
						representationCopyWithNewProperties(volumeBackupRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "type", "FULL"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentIdU, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + VolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", Optional, Update,
						representationCopyWithNewProperties(volumeBackupRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "type", "FULL"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_volume_backups", "test_volume_backups", Optional, Update, volumeBackupDataSourceRepresentation) +
					compartmentIdVariableStr + VolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", Optional, Update, volumeBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_id"),

					resource.TestCheckResourceAttr(datasourceName, "volume_backups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.kms_key_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.size_in_mbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.source_type"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.time_request_received"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.type", "FULL"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.unique_size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.unique_size_in_mbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.volume_id"),
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreVolumeBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockstorageClient()

	if volumeBackupId != "" || volumeId != "" {
		deleteSourceVolumeBackupToCopy()
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_backup" {
			noResourceFound = false
			request := oci_core.GetVolumeBackupRequest{}

			tmp := rs.Primary.ID
			request.VolumeBackupId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			response, err := client.GetVolumeBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VolumeBackupLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreVolumeBackup") {
		resource.AddTestSweepers("CoreVolumeBackup", &resource.Sweeper{
			Name:         "CoreVolumeBackup",
			Dependencies: DependencyGraph["volumeBackup"],
			F:            sweepCoreVolumeBackupResource,
		})
	}
}

func sweepCoreVolumeBackupResource(compartment string) error {
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient()
	volumeBackupIds, err := getVolumeBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, volumeBackupId := range volumeBackupIds {
		if ok := SweeperDefaultResourceId[volumeBackupId]; !ok {
			deleteVolumeBackupRequest := oci_core.DeleteVolumeBackupRequest{}

			deleteVolumeBackupRequest.VolumeBackupId = &volumeBackupId

			deleteVolumeBackupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteVolumeBackup(context.Background(), deleteVolumeBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting VolumeBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", volumeBackupId, error)
				continue
			}
			waitTillCondition(testAccProvider, &volumeBackupId, volumeBackupSweepWaitCondition, time.Duration(3*time.Minute),
				volumeBackupSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVolumeBackupIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "VolumeBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient()

	listVolumeBackupsRequest := oci_core.ListVolumeBackupsRequest{}
	listVolumeBackupsRequest.CompartmentId = &compartmentId
	listVolumeBackupsRequest.LifecycleState = oci_core.VolumeBackupLifecycleStateAvailable
	listVolumeBackupsResponse, err := blockstorageClient.ListVolumeBackups(context.Background(), listVolumeBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VolumeBackup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, volumeBackup := range listVolumeBackupsResponse.Items {
		id := *volumeBackup.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "VolumeBackupId", id)
	}
	return resourceIds, nil
}

func volumeBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if volumeBackupResponse, ok := response.Response.(oci_core.GetVolumeBackupResponse); ok {
		return volumeBackupResponse.LifecycleState != oci_core.VolumeBackupLifecycleStateTerminated
	}
	return false
}

func volumeBackupSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.blockstorageClient().GetVolumeBackup(context.Background(), oci_core.GetVolumeBackupRequest{
		VolumeBackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
