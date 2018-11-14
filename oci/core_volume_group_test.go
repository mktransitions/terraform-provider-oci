// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

var (
	VolumeGroupRequiredOnlyResource = VolumeGroupRequiredOnlyResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Required, Create, volumeGroupRepresentation)

	VolumeGroupResourceConfig = VolumeGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Optional, Create, volumeGroupRepresentation)

	volumeGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":               Representation{repType: Optional, create: `AVAILABLE`},
		"filter":              RepresentationGroup{Required, volumeGroupDataSourceFilterRepresentation}}
	volumeGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_volume_group.test_volume_group.id}`}},
	}

	volumeGroupRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"source_details":      RepresentationGroup{Required, volumeGroupSourceDetailsRepresentation},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}
	volumeGroupSourceDetailsRepresentation = map[string]interface{}{
		"type":       Representation{repType: Required, create: `volumeIds`},
		"volume_ids": Representation{repType: Required, create: []string{`${oci_core_volume.source_volume_list.*.id}`}},
	}
	sourceDetailsJumbledVolumeIdsRepresentation = map[string]interface{}{
		"type":       Representation{repType: Required, create: `volumeIds`},
		"volume_ids": Representation{repType: Required, create: []string{`${oci_core_volume.source_volume_list.*.id[1]}`, `${oci_core_volume.source_volume_list.*.id[0]}`}},
	}
	sourceDetailsSingleVolumeIdSourceDetailsRepresentation = map[string]interface{}{
		"type":       Representation{repType: Required, create: `volumeIds`},
		"volume_ids": Representation{repType: Required, create: []string{`${oci_core_volume.source_volume_list.*.id[1]}`}},
	}

	VolumeGroupResourceConfigJumbledVolumeIds = VolumeGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Required, Create,
			getUpdatedRepresentationCopy("source_details", RepresentationGroup{Required, sourceDetailsJumbledVolumeIdsRepresentation}, volumeGroupRepresentation))

	VolumeGroupResourceConfigSingleVolumeId = VolumeGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Required, Create,
			getUpdatedRepresentationCopy("source_details", RepresentationGroup{Required, sourceDetailsSingleVolumeIdSourceDetailsRepresentation}, volumeGroupRepresentation))

	VolumeGroupResourceDependencies             = DefinedTagsDependencies + VolumeGroupRequiredOnlyResourceDependencies
	VolumeGroupRequiredOnlyResourceDependencies = AvailabilityDomainConfig + SourceVolumeListDependency
	VolumeGroupAsDependency                     = generateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Required, Create, volumeGroupRepresentation) + SourceVolumeListDependency
	SourceVolumeListDependency                  = `
resource "oci_core_volume" "source_volume_list" {
	count = 2
	display_name = "${format("source-volume-%d", count.index + 1)}"

	#Required
	availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
}
`
)

func TestCoreVolumeGroupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume_group.test_volume_group"
	datasourceName := "data.oci_core_volume_groups.test_volume_groups"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeGroupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VolumeGroupRequiredOnlyResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Required, Create, volumeGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// We need to assert that the volume group created above did cause the source volume to have the volume
			// group id property populated correctly. Since the TF framework doesn't have a RefreshOnly directive, we are
			// using PlanOnly to trigger a refresh, and then assert on the value
			{
				Config:   config + compartmentIdVariableStr + VolumeGroupRequiredOnlyResource,
				PlanOnly: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("oci_core_volume.source_volume_list.0", "volume_group_id"),
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeGroupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + VolumeGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Optional, Create, volumeGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + VolumeGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Optional, Update, volumeGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "size_in_mbs"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.type", "volumeIds"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "volume_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify that the change in order of the volume ids doesn't result in a new resource
			{
				Config:             config + compartmentIdVariableStr + VolumeGroupResourceConfigJumbledVolumeIds,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// verify that the change in list of volume ids does cause a change in the plan
			{
				Config:             config + compartmentIdVariableStr + VolumeGroupResourceConfigSingleVolumeId,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_volume_groups", "test_volume_groups", Optional, Update, volumeGroupDataSourceRepresentation) +
					compartmentIdVariableStr + VolumeGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", Optional, Update, volumeGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "volume_groups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.size_in_mbs"),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.source_details.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.source_details.0.type", "volumeIds"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_groups.0.volume_ids.#"),
					resource.TestCheckResourceAttr(datasourceName, "volume_groups.0.is_hydrated", "false"),
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

func testAccCheckCoreVolumeGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockstorageClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_group" {
			noResourceFound = false
			request := oci_core.GetVolumeGroupRequest{}

			tmp := rs.Primary.ID
			request.VolumeGroupId = &tmp

			response, err := client.GetVolumeGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VolumeGroupLifecycleStateTerminated): true,
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
