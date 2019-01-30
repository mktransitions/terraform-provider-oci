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
	DrgRequiredOnlyResource = generateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation)

	drgDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"filter":         RepresentationGroup{Required, drgDataSourceFilterRepresentation}}
	drgDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_drg.test_drg.id}`}},
	}

	drgRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `MyDrg`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	DrgResourceDependencies = DefinedTagsDependencies
)

func TestCoreDrgResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_drg.test_drg"
	datasourceName := "data.oci_core_drgs.test_drgs"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreDrgDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DrgResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DrgResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DrgResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg", "test_drg", Optional, Create, drgRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyDrg"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DrgResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg", "test_drg", Optional, Update, drgRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

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
					generateDataSourceFromRepresentationMap("oci_core_drgs", "test_drgs", Optional, Update, drgDataSourceRepresentation) +
					compartmentIdVariableStr + DrgResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_drg", "test_drg", Optional, Update, drgRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "drgs.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "drgs.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "drgs.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "drgs.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "drgs.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "drgs.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "drgs.0.state"),
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

func testAccCheckCoreDrgDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_drg" {
			noResourceFound = false
			request := oci_core.GetDrgRequest{}

			tmp := rs.Primary.ID
			request.DrgId = &tmp

			response, err := client.GetDrg(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.DrgLifecycleStateTerminated): true,
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
	resource.AddTestSweepers("CoreDrg", &resource.Sweeper{
		Name:         "CoreDrg",
		Dependencies: DependencyGraph["drg"],
		F:            sweepCoreDrgResource,
	})
}

func sweepCoreDrgResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient
	drgIds, err := getDrgIds(compartment)
	if err != nil {
		return err
	}
	for _, drgId := range drgIds {
		if ok := SweeperDefaultResourceId[drgId]; !ok {
			deleteDrgRequest := oci_core.DeleteDrgRequest{}

			deleteDrgRequest.DrgId = &drgId

			deleteDrgRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteDrg(context.Background(), deleteDrgRequest)
			if error != nil {
				fmt.Printf("Error deleting Drg %s %s, It is possible that the resource is already deleted. Please verify manually \n", drgId, error)
				continue
			}
			waitTillCondition(testAccProvider, &drgId, drgSweepWaitCondition, time.Duration(3*time.Minute),
				drgSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getDrgIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DrgId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient

	listDrgsRequest := oci_core.ListDrgsRequest{}
	listDrgsRequest.CompartmentId = &compartmentId
	listDrgsResponse, err := virtualNetworkClient.ListDrgs(context.Background(), listDrgsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Drg list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, drg := range listDrgsResponse.Items {
		id := *drg.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "DrgId", id)
	}
	return resourceIds, nil
}

func drgSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if drgResponse, ok := response.Response.(oci_core.GetDrgResponse); ok {
		return drgResponse.LifecycleState == oci_core.DrgLifecycleStateTerminated
	}
	return false
}

func drgSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient.GetDrg(context.Background(), oci_core.GetDrgRequest{
		DrgId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
