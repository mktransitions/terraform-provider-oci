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

const (
	RouteTableRequiredOnlyResource = RouteTableResourceDependencies + `
resource "oci_core_route_table" "test_route_table" {
	#Required
	compartment_id = "${var.compartment_id}"
	route_rules {
		#Required
		cidr_block = "${var.route_table_route_rules_cidr_block}"
		network_entity_id = "${oci_core_internet_gateway.test_network_entity.id}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
`
	RouteTableRequiredOnlyResourceWithSecondNetworkEntity = RouteTableResourceDependencies + `
resource "oci_core_route_table" "test_route_table" {
	#Required
	compartment_id = "${var.compartment_id}"
	route_rules {
		#Required
		cidr_block = "${var.route_table_route_rules_cidr_block}"
		network_entity_id = "${oci_core_drg.test_drg.id}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
`
	RouteTableResourceConfig = RouteTableResourceDependencies + `
resource "oci_core_route_table" "test_route_table" {
	#Required
	compartment_id = "${var.compartment_id}"
	route_rules {
		#Required
		network_entity_id = "${oci_core_internet_gateway.test_network_entity.id}"

		#Optional
		destination = "${var.route_table_route_rules_destination}"
		destination_type = "${var.route_table_route_rules_destination_type}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.route_table_defined_tags_value}")}"
	display_name = "${var.route_table_display_name}"
	freeform_tags = "${var.route_table_freeform_tags}"
}
`
	RouteTableResourceConfigWithServiceCidr = RouteTableResourceDependencies + `
resource "oci_core_route_table" "test_route_table" {
	#Required
	compartment_id = "${var.compartment_id}"
	route_rules {
		#Required
		network_entity_id = "${oci_core_service_gateway.test_service_gateway.id}"

		#Optional
		destination = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"
		destination_type = "${var.route_table_route_rules_destination_type}"
	}
	route_rules {
		#Required
		destination = "${var.route_table_route_rules_destination}"
		network_entity_id = "${oci_core_drg.test_drg.id}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.route_table_defined_tags_value}")}"
	display_name = "${var.route_table_display_name}"
	freeform_tags = "${var.route_table_freeform_tags}"
}

resource "oci_core_service_gateway" "test_service_gateway" {
    #Required
    compartment_id = "${var.compartment_id}"
    services {
        service_id = "${lookup(data.oci_core_services.test_services.services[0], "id")}"
    }
    vcn_id = "${oci_core_vcn.test_vcn.id}"
}

data "oci_core_services" "test_services" {
}
`
	RouteTableResourceConfigWithServiceCidrAddingCidrBlock = RouteTableResourceDependencies + `
resource "oci_core_route_table" "test_route_table" {
	#Required
	compartment_id = "${var.compartment_id}"
	route_rules {
		#Required
		network_entity_id = "${oci_core_service_gateway.test_service_gateway.id}"

		#Optional
		cidr_block = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"
		destination = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"
		destination_type = "${var.route_table_route_rules_destination_type}"
	}
	route_rules {
		#Required
		destination = "${var.route_table_route_rules_destination}"
		network_entity_id = "${oci_core_drg.test_drg.id}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.route_table_defined_tags_value}")}"
	display_name = "${var.route_table_display_name}"
	freeform_tags = "${var.route_table_freeform_tags}"
}

resource "oci_core_service_gateway" "test_service_gateway" {
    #Required
    compartment_id = "${var.compartment_id}"
    services {
        service_id = "${lookup(data.oci_core_services.test_services.services[0], "id")}"
    }
    vcn_id = "${oci_core_vcn.test_vcn.id}"
}

data "oci_core_services" "test_services" {
}
`

	RouteTableResourceConfigWithSecondNetworkEntity = RouteTableResourceDependencies + `
resource "oci_core_route_table" "test_route_table" {
	#Required
	compartment_id = "${var.compartment_id}"
	route_rules {
		#Required
		network_entity_id = "${oci_core_drg.test_drg.id}"

		#Optional
		destination = "${var.route_table_route_rules_destination}"
		destination_type = "${var.route_table_route_rules_destination_type}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.route_table_defined_tags_value}")}"
	display_name = "${var.route_table_display_name}"
	freeform_tags = "${var.route_table_freeform_tags}"
}
`
	RouteTablePropertyVariables = `
variable "route_table_defined_tags_value" { default = "value" }
variable "route_table_display_name" { default = "MyRouteTable" }
variable "route_table_freeform_tags" { default = {"Department"= "Finance"} }
variable "route_table_route_rules_cidr_block" { default = "0.0.0.0/0" }
variable "route_table_route_rules_destination" { default = "0.0.0.0/0" }
variable "route_table_route_rules_destination_type" { default = "CIDR_BLOCK" }
variable "route_table_state" { default = "AVAILABLE" }

`
	RouteTableResourceDependencies = VcnPropertyVariables + VcnResourceConfig + `
	resource "oci_core_internet_gateway" "test_network_entity" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_vcn.test_vcn.id}"
		display_name = "-tf-internet-gateway"
	}

	resource "oci_core_drg" "test_drg" {
		#Required
		compartment_id = "${var.compartment_id}"
	}
	`
)

func TestCoreRouteTableResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_route_table.test_route_table"
	datasourceName := "data.oci_core_route_tables.test_route_tables"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreRouteTableDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + RouteTablePropertyVariables + compartmentIdVariableStr + RouteTableRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"cidr_block": "0.0.0.0/0",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify update to deprecated cidr_block
			{
				Config: config + `
variable "route_table_defined_tags_value" { default = "value" }
variable "route_table_display_name" { default = "MyRouteTable" }
variable "route_table_freeform_tags" { default = {"Department"= "Finance"} }
variable "route_table_route_rules_cidr_block" { default = "10.0.0.0/8" }
variable "route_table_route_rules_destination" { default = "0.0.0.0/0" }
variable "route_table_route_rules_destination_type" { default = "CIDR_BLOCK" }
variable "route_table_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + RouteTableRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"cidr_block": "10.0.0.0/8"}, []string{"network_entity_id"}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify update to network_id
			{
				Config: config + `
variable "route_table_defined_tags_value" { default = "value" }
variable "route_table_display_name" { default = "MyRouteTable" }
variable "route_table_freeform_tags" { default = {"Department"= "Finance"} }
variable "route_table_route_rules_cidr_block" { default = "10.0.0.0/8" }
variable "route_table_route_rules_destination" { default = "0.0.0.0/0" }
variable "route_table_route_rules_destination_type" { default = "CIDR_BLOCK" }
variable "route_table_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + RouteTableRequiredOnlyResourceWithSecondNetworkEntity,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"cidr_block": "10.0.0.0/8"}, []string{"network_entity_id"}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify create with destination_type
			{
				Config: config + `
variable "route_table_defined_tags_value" { default = "value" }
variable "route_table_display_name" { default = "MyRouteTable" }
variable "route_table_freeform_tags" { default = {"Department"= "Finance"} }
variable "route_table_route_rules_cidr_block" { default = "10.0.0.0/8" }
variable "route_table_route_rules_destination" { default = "0.0.0.0/0" }
variable "route_table_route_rules_destination_type" { default = "SERVICE_CIDR_BLOCK" }
variable "route_table_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + RouteTableResourceConfigWithServiceCidr,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "SERVICE_CIDR_BLOCK"}, []string{"network_entity_id", "destination"}),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "CIDR_BLOCK", "destination": "0.0.0.0/0"}, []string{"network_entity_id"}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				),
			},
			// verify update after having a destination_type rule
			{
				Config: config + `
variable "route_table_defined_tags_value" { default = "value" }
variable "route_table_display_name" { default = "MyRouteTable" }
variable "route_table_freeform_tags" { default = {"Department"= "Finance"} }
variable "route_table_route_rules_cidr_block" { default = "10.0.0.0/8" }
variable "route_table_route_rules_destination" { default = "0.0.0.0/1" }
variable "route_table_route_rules_destination_type" { default = "SERVICE_CIDR_BLOCK" }
variable "route_table_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + RouteTableResourceConfigWithServiceCidr,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "SERVICE_CIDR_BLOCK"}, []string{"network_entity_id", "destination"}),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "CIDR_BLOCK", "destination": "0.0.0.0/1"}, []string{"network_entity_id"}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				),
			},
			// verify adding cidr_block to a rule that has destination already
			{
				Config: config + `
variable "route_table_defined_tags_value" { default = "value" }
variable "route_table_display_name" { default = "MyRouteTable" }
variable "route_table_freeform_tags" { default = {"Department"= "Finance"} }
variable "route_table_route_rules_cidr_block" { default = "10.0.0.0/8" }
variable "route_table_route_rules_destination" { default = "0.0.0.0/1" }
variable "route_table_route_rules_destination_type" { default = "SERVICE_CIDR_BLOCK" }
variable "route_table_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + RouteTableResourceConfigWithServiceCidrAddingCidrBlock,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "SERVICE_CIDR_BLOCK"}, []string{"network_entity_id", "destination"}),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "CIDR_BLOCK", "destination": "0.0.0.0/1"}, []string{"network_entity_id"}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + RouteTableResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + RouteTablePropertyVariables + compartmentIdVariableStr + RouteTableResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyRouteTable"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"cidr_block":       "0.0.0.0/0",
						"destination":      "0.0.0.0/0",
						"destination_type": "CIDR_BLOCK",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "route_table_defined_tags_value" { default = "updatedValue" }
variable "route_table_display_name" { default = "displayName2" }
variable "route_table_freeform_tags" { default = {"Department"= "Accounting"} }
variable "route_table_route_rules_cidr_block" { default = "10.0.0.0/8" }
variable "route_table_route_rules_destination" { default = "10.0.0.0/8" }
variable "route_table_route_rules_destination_type" { default = "CIDR_BLOCK" }
variable "route_table_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + RouteTableResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"cidr_block":       "10.0.0.0/8",
						"destination":      "10.0.0.0/8",
						"destination_type": "CIDR_BLOCK",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify updates to network entity
			{
				Config: config + `
variable "route_table_defined_tags_value" { default = "updatedValue" }
variable "route_table_display_name" { default = "displayName2" }
variable "route_table_freeform_tags" { default = {"Department"= "Accounting"} }
variable "route_table_route_rules_destination" { default = "10.0.0.0/8" }
variable "route_table_route_rules_destination_type" { default = "CIDR_BLOCK" }
variable "route_table_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + RouteTableResourceConfigWithSecondNetworkEntity,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"cidr_block": "10.0.0.0/8", "destination_type": "CIDR_BLOCK"}, []string{"network_entity_id"}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				Config: config + `
variable "route_table_defined_tags_value" { default = "updatedValue" }
variable "route_table_display_name" { default = "displayName2" }
variable "route_table_freeform_tags" { default = {"Department"= "Accounting"} }
variable "route_table_route_rules_cidr_block" { default = "10.0.0.0/8" }
variable "route_table_route_rules_destination" { default = "10.0.0.0/8" }
variable "route_table_route_rules_destination_type" { default = "CIDR_BLOCK" }
variable "route_table_state" { default = "AVAILABLE" }

data "oci_core_route_tables" "test_route_tables" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.route_table_display_name}"
	state = "${var.route_table_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_route_table.test_route_table.id}"]
    }
}
                ` + compartmentIdVariableStr + RouteTableResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "route_tables.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "route_tables.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(datasourceName, "route_tables.0.route_rules", map[string]string{
						"cidr_block":       "10.0.0.0/8",
						"destination":      "10.0.0.0/8",
						"destination_type": "CIDR_BLOCK",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttrSet(datasourceName, "route_tables.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "route_tables.0.vcn_id"),
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

func testAccCheckCoreRouteTableDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_route_table" {
			noResourceFound = false
			request := oci_core.GetRouteTableRequest{}

			tmp := rs.Primary.ID
			request.RtId = &tmp

			response, err := client.GetRouteTable(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.RouteTableLifecycleStateTerminated): true,
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
