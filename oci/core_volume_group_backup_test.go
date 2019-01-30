// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

var (
	VolumeGroupBackupRequiredOnlyResource = VolumeGroupBackupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup", Required, Create, volumeGroupBackupRepresentation)

	volumeGroupBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":    Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"volume_group_id": Representation{repType: Optional, create: `${oci_core_volume_group.test_volume_group.id}`},
		"filter":          RepresentationGroup{Required, volumeGroupBackupDataSourceFilterRepresentation}}
	volumeGroupBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_volume_group_backup.test_volume_group_backup.id}`}},
	}

	volumeGroupBackupRepresentation = map[string]interface{}{
		"volume_group_id": Representation{repType: Required, create: `${oci_core_volume_group.test_volume_group.id}`},
		"compartment_id":  Representation{repType: Optional, create: `${var.compartment_id}`},
		"defined_tags":    Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":    Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":   Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"type":            Representation{repType: Optional, create: `INCREMENTAL`},
	}

	VolumeGroupBackupResourceDependencies = VolumeGroupResourceConfig
)

func TestCoreVolumeGroupBackupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume_group_backup.test_volume_group_backup"
	datasourceName := "data.oci_core_volume_group_backups.test_volume_group_backups"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeGroupBackupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VolumeGroupBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup", Required, Create, volumeGroupBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "volume_group_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeGroupBackupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + VolumeGroupBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup", Optional, Create, volumeGroupBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "type", "INCREMENTAL"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_backup_ids.#"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_group_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + VolumeGroupBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup", Optional, Update, volumeGroupBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "type", "INCREMENTAL"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_backup_ids.#"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_group_id"),

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
					generateDataSourceFromRepresentationMap("oci_core_volume_group_backups", "test_volume_group_backups", Optional, Update, volumeGroupBackupDataSourceRepresentation) +
					compartmentIdVariableStr + VolumeGroupBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup", Optional, Update, volumeGroupBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_id"),

					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.type", "INCREMENTAL"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.volume_backup_ids.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.volume_group_id"),
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

func testAccCheckCoreVolumeGroupBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockstorageClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_group_backup" {
			noResourceFound = false
			request := oci_core.GetVolumeGroupBackupRequest{}

			tmp := rs.Primary.ID
			request.VolumeGroupBackupId = &tmp

			response, err := client.GetVolumeGroupBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VolumeGroupBackupLifecycleStateTerminated): true,
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
	resource.AddTestSweepers("CoreVolumeGroupBackup", &resource.Sweeper{
		Name:         "CoreVolumeGroupBackup",
		Dependencies: DependencyGraph["volumeGroupBackup"],
		F:            sweepCoreVolumeGroupBackupResource,
	})
}

func sweepCoreVolumeGroupBackupResource(compartment string) error {
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient
	volumeGroupBackupIds, err := getVolumeGroupBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, volumeGroupBackupId := range volumeGroupBackupIds {
		if ok := SweeperDefaultResourceId[volumeGroupBackupId]; !ok {
			deleteVolumeGroupBackupRequest := oci_core.DeleteVolumeGroupBackupRequest{}

			deleteVolumeGroupBackupRequest.VolumeGroupBackupId = &volumeGroupBackupId

			deleteVolumeGroupBackupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteVolumeGroupBackup(context.Background(), deleteVolumeGroupBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting VolumeGroupBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", volumeGroupBackupId, error)
				continue
			}
			waitTillCondition(testAccProvider, &volumeGroupBackupId, volumeGroupBackupSweepWaitCondition, time.Duration(3*time.Minute),
				volumeGroupBackupSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVolumeGroupBackupIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "VolumeGroupBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient

	listVolumeGroupBackupsRequest := oci_core.ListVolumeGroupBackupsRequest{}
	listVolumeGroupBackupsRequest.CompartmentId = &compartmentId
	listVolumeGroupBackupsResponse, err := blockstorageClient.ListVolumeGroupBackups(context.Background(), listVolumeGroupBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VolumeGroupBackup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, volumeGroupBackup := range listVolumeGroupBackupsResponse.Items {
		id := *volumeGroupBackup.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "VolumeGroupBackupId", id)
	}
	return resourceIds, nil
}

func volumeGroupBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if volumeGroupBackupResponse, ok := response.Response.(oci_core.GetVolumeGroupBackupResponse); ok {
		return volumeGroupBackupResponse.LifecycleState == oci_core.VolumeGroupBackupLifecycleStateTerminated
	}
	return false
}

func volumeGroupBackupSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.blockstorageClient.GetVolumeGroupBackup(context.Background(), oci_core.GetVolumeGroupBackupRequest{
		VolumeGroupBackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
