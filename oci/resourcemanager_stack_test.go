// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/v25/common"
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v25/resourcemanager"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	stackSingularDataSourceRepresentation = map[string]interface{}{
		"stack_id": Representation{repType: Required, create: `${var.resource_manager_stack_id}`},
	}

	stackDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `TestResourcemanagerStackResource_basic`, update: `TestResourcemanagerStackResource_basic`},
		"id":             Representation{repType: Optional, create: `${oci_resourcemanager_stack.test_stack.id}`},
		"state":          Representation{repType: Required, create: `ACTIVE`}, // make `required` here so it can be asserted against in step 0
	}

	StackResourceConfig = DefinedTagsDependencies
)

func TestResourcemanagerStackResource_basic(t *testing.T) {
	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "TestResourcemanagerStackResource_basic") {
		t.Skip("Skipping suppressed TestResourcemanagerStackResource_basic")
	}

	httpreplay.SetScenario("TestResourcemanagerStackResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	client := GetTestClients(&schema.ResourceData{}).resourceManagerClient()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceManagerStackId, err := createResourceManagerStack(*client, "TestResourcemanagerStackResource_basic", compartmentId)
	if err != nil {
		t.Errorf("cannot create resource manager stack for the test run: %v", err)
	}

	datasourceName := "data.oci_resourcemanager_stacks.test_stacks"
	singularDatasourceName := "data.oci_resourcemanager_stack.test_stack"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		CheckDestroy: func(s *terraform.State) error {
			return destroyResourceManagerStack(*client, resourceManagerStackId)
		},
		PreventPostDestroyRefresh: true,
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `
					variable "resource_manager_stack_id" { default = "` + resourceManagerStackId + `" }
					` +
					generateDataSourceFromRepresentationMap("oci_resourcemanager_stacks", "test_stacks", Required, Create, stackDataSourceRepresentation) +
					compartmentIdVariableStr + StackResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttrSet(datasourceName, "stacks.#"),
					resource.TestCheckResourceAttr(datasourceName, "stacks.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "stacks.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "stacks.0.description"),
					resource.TestCheckResourceAttr(datasourceName, "stacks.0.display_name", "TestResourcemanagerStackResource_basic"),
					resource.TestCheckResourceAttr(datasourceName, "stacks.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "stacks.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "stacks.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "stacks.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config + `
					variable "resource_manager_stack_id" { default = "` + resourceManagerStackId + `" }
					` +
					generateDataSourceFromRepresentationMap("oci_resourcemanager_stacks", "test_stacks", Required, Create, stackDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_resourcemanager_stack", "test_stack", Required, Create, stackSingularDataSourceRepresentation) +
					compartmentIdVariableStr + StackResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "stack_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "config_source.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "config_source.0.config_source_type", "ZIP_UPLOAD"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TestResourcemanagerStackResource_basic"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "variables.%", "3"),
				),
			},
		},
	})
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("ResourcemanagerStack") {
		resource.AddTestSweepers("ResourcemanagerStack", &resource.Sweeper{
			Name:         "ResourcemanagerStack",
			Dependencies: DependencyGraph["stack"],
			F:            sweepResourcemanagerStackResource,
		})
	}
}

func sweepResourcemanagerStackResource(compartment string) error {
	resourceManagerClient := GetTestClients(&schema.ResourceData{}).resourceManagerClient()
	stackIds, err := getStackIds(compartment)
	if err != nil {
		return err
	}
	for _, stackId := range stackIds {
		if ok := SweeperDefaultResourceId[stackId]; !ok {
			deleteStackRequest := oci_resourcemanager.DeleteStackRequest{}

			deleteStackRequest.StackId = &stackId

			deleteStackRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "resourcemanager")
			_, error := resourceManagerClient.DeleteStack(context.Background(), deleteStackRequest)
			if error != nil {
				fmt.Printf("Error deleting Stack %s %s, It is possible that the resource is already deleted. Please verify manually \n", stackId, error)
				continue
			}
			waitTillCondition(testAccProvider, &stackId, stackSweepWaitCondition, time.Duration(3*time.Minute),
				stackSweepResponseFetchOperation, "resourcemanager", true)
		}
	}
	return nil
}

func getStackIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "StackId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	resourceManagerClient := GetTestClients(&schema.ResourceData{}).resourceManagerClient()

	listStacksRequest := oci_resourcemanager.ListStacksRequest{}
	listStacksRequest.CompartmentId = &compartmentId
	listStacksRequest.LifecycleState = oci_resourcemanager.StackLifecycleStateActive
	listStacksResponse, err := resourceManagerClient.ListStacks(context.Background(), listStacksRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Stack list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, stack := range listStacksResponse.Items {
		id := *stack.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "StackId", id)
	}
	return resourceIds, nil
}

func stackSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if stackResponse, ok := response.Response.(oci_resourcemanager.GetStackResponse); ok {
		return stackResponse.LifecycleState != oci_resourcemanager.StackLifecycleStateDeleted
	}
	return false
}

func stackSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.resourceManagerClient().GetStack(context.Background(), oci_resourcemanager.GetStackRequest{
		StackId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
