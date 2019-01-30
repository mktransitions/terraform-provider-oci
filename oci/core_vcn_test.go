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
	VcnRequiredOnlyResource = VcnRequiredOnlyResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation)

	VcnResourceConfig = generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, vcnRepresentation)

	vcnDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, vcnDataSourceFilterRepresentation}}
	vcnDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_vcn.test_vcn.id}`}},
	}

	vcnRepresentation = map[string]interface{}{
		"cidr_block":     Representation{repType: Required, create: `10.0.0.0/16`},
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"dns_label":      Representation{repType: Optional, create: `dnslabel`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	VcnRequiredOnlyResourceDependencies = ``
	VcnResourceDependencies             = DefinedTagsDependencies
)

func TestCoreVcnResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_vcn.test_vcn"
	datasourceName := "data.oci_core_vcns.test_vcns"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVcnDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VcnResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VcnResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + VcnResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, vcnRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
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
				Config: config + compartmentIdVariableStr + VcnResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Update, vcnRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
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
					generateDataSourceFromRepresentationMap("oci_core_vcns", "test_vcns", Optional, Update, vcnDataSourceRepresentation) +
					compartmentIdVariableStr + VcnResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Update, vcnRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "virtual_networks.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.dns_label", "dnslabel"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.state"),
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

func testAccCheckCoreVcnDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_vcn" {
			noResourceFound = false
			request := oci_core.GetVcnRequest{}

			tmp := rs.Primary.ID
			request.VcnId = &tmp

			response, err := client.GetVcn(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VcnLifecycleStateTerminated): true,
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
	resource.AddTestSweepers("CoreVcn", &resource.Sweeper{
		Name:         "CoreVcn",
		Dependencies: DependencyGraph["vcn"],
		F:            sweepCoreVcnResource,
	})
}

func sweepCoreVcnResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient
	vcnIds, err := getVcnIds(compartment)
	if err != nil {
		return err
	}
	for _, vcnId := range vcnIds {
		if ok := SweeperDefaultResourceId[vcnId]; !ok {
			deleteVcnRequest := oci_core.DeleteVcnRequest{}

			deleteVcnRequest.VcnId = &vcnId

			deleteVcnRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteVcn(context.Background(), deleteVcnRequest)
			if error != nil {
				fmt.Printf("Error deleting Vcn %s %s, It is possible that the resource is already deleted. Please verify manually \n", vcnId, error)
				continue
			}
			waitTillCondition(testAccProvider, &vcnId, vcnSweepWaitCondition, time.Duration(3*time.Minute),
				vcnSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVcnIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "VcnId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient

	listVcnsRequest := oci_core.ListVcnsRequest{}
	listVcnsRequest.CompartmentId = &compartmentId
	listVcnsRequest.LifecycleState = oci_core.VcnLifecycleStateAvailable
	listVcnsResponse, err := virtualNetworkClient.ListVcns(context.Background(), listVcnsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Vcn list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vcn := range listVcnsResponse.Items {
		id := *vcn.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "VcnId", id)
		SweeperDefaultResourceId[*vcn.DefaultDhcpOptionsId] = true
		SweeperDefaultResourceId[*vcn.DefaultRouteTableId] = true
		SweeperDefaultResourceId[*vcn.DefaultSecurityListId] = true

	}
	return resourceIds, nil
}

func vcnSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vcnResponse, ok := response.Response.(oci_core.GetVcnResponse); ok {
		return vcnResponse.LifecycleState == oci_core.VcnLifecycleStateTerminated
	}
	return false
}

func vcnSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient.GetVcn(context.Background(), oci_core.GetVcnRequest{
		VcnId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
