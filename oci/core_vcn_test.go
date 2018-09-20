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
	VcnRequiredOnlyResource = VcnRequiredOnlyResourceDependencies + `
resource "oci_core_vcn" "test_vcn" {
	#Required
	cidr_block = "${var.vcn_cidr_block}"
	compartment_id = "${var.compartment_id}"
}
`

	VcnResourceConfig = VcnResourceDependencies + `
resource "oci_core_vcn" "test_vcn" {
	#Required
	cidr_block = "${var.vcn_cidr_block}"
	compartment_id = "${var.compartment_id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.vcn_defined_tags_value}")}"
	display_name = "${var.vcn_display_name}"
	dns_label = "${var.vcn_dns_label}"
	freeform_tags = "${var.vcn_freeform_tags}"
}
`
	VcnPropertyVariables = `
variable "vcn_cidr_block" { default = "10.0.0.0/16" }
variable "vcn_defined_tags_value" { default = "value" }
variable "vcn_display_name" { default = "displayName" }
variable "vcn_dns_label" { default = "dnslabel" }
variable "vcn_freeform_tags" { default = {"Department"= "Finance"} }
variable "vcn_state" { default = "AVAILABLE" }

`
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
				Config: config + VcnPropertyVariables + compartmentIdVariableStr + VcnRequiredOnlyResource,
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
				Config: config + VcnPropertyVariables + compartmentIdVariableStr + VcnResourceConfig,
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
				Config: config + `
variable "vcn_cidr_block" { default = "10.0.0.0/16" }
variable "vcn_defined_tags_value" { default = "updatedValue" }
variable "vcn_display_name" { default = "displayName2" }
variable "vcn_dns_label" { default = "dnslabel" }
variable "vcn_freeform_tags" { default = {"Department"= "Accounting"} }
variable "vcn_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + VcnResourceConfig,
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
				Config: config + `
variable "vcn_cidr_block" { default = "10.0.0.0/16" }
variable "vcn_defined_tags_value" { default = "updatedValue" }
variable "vcn_display_name" { default = "displayName2" }
variable "vcn_dns_label" { default = "dnslabel" }
variable "vcn_freeform_tags" { default = {"Department"= "Accounting"} }
variable "vcn_state" { default = "AVAILABLE" }

data "oci_core_vcns" "test_vcns" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.vcn_display_name}"
	state = "${var.vcn_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_vcn.test_vcn.id}"]
    }
}
                ` + compartmentIdVariableStr + VcnResourceConfig,
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
