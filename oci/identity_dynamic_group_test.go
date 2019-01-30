// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

var (
	dynamicGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"filter":         RepresentationGroup{Required, dynamicGroupDataSourceFilterRepresentation}}
	dynamicGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_dynamic_group.test_dynamic_group.id}`}},
	}

	dynamicGroupRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"description":    Representation{repType: Required, create: `Instance group for dev compartment`, update: `description2`},
		"matching_rule":  Representation{repType: Required, create: `${var.dynamic_group_matching_rule}`, update: `${var.dynamic_group_matching_rule}`},
		"name":           Representation{repType: Required, create: `DevCompartmentDynamicGroup`},
	}

	DynamicGroupResourceDependencies = ""
)

func TestIdentityDynamicGroupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	matchingRuleValueStr := fmt.Sprintf("instance.compartment_id='%s'", compartmentId)
	matchingRuleVariableStr := fmt.Sprintf("variable \"dynamic_group_matching_rule\" {default = \"%s\" }\n", matchingRuleValueStr)

	matchingRule2ValueStr := fmt.Sprintf("instance.compartment_id='%s'", compartmentId)
	matchingRule2VariableStr := fmt.Sprintf("variable \"dynamic_group_matching_rule\" {default = \"%s\" }\n", matchingRule2ValueStr)
	resourceName := "oci_identity_dynamic_group.test_dynamic_group"
	datasourceName := "data.oci_identity_dynamic_groups.test_dynamic_groups"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityDynamicGroupDestroy,
		Steps: []resource.TestStep{
			// verify matching rule syntax
			{
				Config: config + `
variable "dynamic_group_description" { default = "description2" }
variable "dynamic_group_matching_rule" { default = "bad_matching_rule" }
variable "dynamic_group_name" { default = "DevCompartmentDynamicGroup" }` + compartmentIdVariableStr + generateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", Required, Create, dynamicGroupRepresentation),
				ExpectError: regexp.MustCompile("Unable to parse matching rule"),
			},
			// verify create
			{
				Config: config + compartmentIdVariableStr + matchingRuleVariableStr + DynamicGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", Required, Create, dynamicGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "description", "Instance group for dev compartment"),
					resource.TestCheckResourceAttr(resourceName, "matching_rule", matchingRuleValueStr),
					resource.TestCheckResourceAttr(resourceName, "name", "DevCompartmentDynamicGroup"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + matchingRule2VariableStr + DynamicGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", Optional, Update, dynamicGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "matching_rule", matchingRule2ValueStr),
					resource.TestCheckResourceAttr(resourceName, "name", "DevCompartmentDynamicGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				Config: config + matchingRule2VariableStr +
					generateDataSourceFromRepresentationMap("oci_identity_dynamic_groups", "test_dynamic_groups", Optional, Update, dynamicGroupDataSourceRepresentation) +
					compartmentIdVariableStr + DynamicGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", Optional, Update, dynamicGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

					resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.description", "description2"),
					resource.TestCheckResourceAttrSet(datasourceName, "dynamic_groups.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.matching_rule", matchingRule2ValueStr),
					resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.name", "DevCompartmentDynamicGroup"),
					resource.TestCheckResourceAttrSet(datasourceName, "dynamic_groups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "dynamic_groups.0.time_created"),
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

func testAccCheckIdentityDynamicGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_dynamic_group" {
			noResourceFound = false
			request := oci_identity.GetDynamicGroupRequest{}

			tmp := rs.Primary.ID
			request.DynamicGroupId = &tmp

			response, err := client.GetDynamicGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.DynamicGroupLifecycleStateDeleted): true,
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
	resource.AddTestSweepers("IdentityDynamicGroup", &resource.Sweeper{
		Name:         "IdentityDynamicGroup",
		Dependencies: DependencyGraph["dynamicGroup"],
		F:            sweepIdentityDynamicGroupResource,
	})
}

func sweepIdentityDynamicGroupResource(compartment string) error {
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient
	dynamicGroupIds, err := getDynamicGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, dynamicGroupId := range dynamicGroupIds {
		if ok := SweeperDefaultResourceId[dynamicGroupId]; !ok {
			deleteDynamicGroupRequest := oci_identity.DeleteDynamicGroupRequest{}

			deleteDynamicGroupRequest.DynamicGroupId = &dynamicGroupId

			deleteDynamicGroupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")
			_, error := identityClient.DeleteDynamicGroup(context.Background(), deleteDynamicGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting DynamicGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", dynamicGroupId, error)
				continue
			}
			waitTillCondition(testAccProvider, &dynamicGroupId, dynamicGroupSweepWaitCondition, time.Duration(3*time.Minute),
				dynamicGroupSweepResponseFetchOperation, "identity", true)
		}
	}
	return nil
}

func getDynamicGroupIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DynamicGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient

	listDynamicGroupsRequest := oci_identity.ListDynamicGroupsRequest{}
	listDynamicGroupsRequest.CompartmentId = &compartmentId
	listDynamicGroupsResponse, err := identityClient.ListDynamicGroups(context.Background(), listDynamicGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DynamicGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dynamicGroup := range listDynamicGroupsResponse.Items {
		id := *dynamicGroup.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "DynamicGroupId", id)
	}
	return resourceIds, nil
}

func dynamicGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dynamicGroupResponse, ok := response.Response.(oci_identity.GetDynamicGroupResponse); ok {
		return dynamicGroupResponse.LifecycleState == oci_identity.DynamicGroupLifecycleStateDeleted
	}
	return false
}

func dynamicGroupSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.identityClient.GetDynamicGroup(context.Background(), oci_identity.GetDynamicGroupRequest{
		DynamicGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
