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
	DhcpOptionsRequiredOnlyResource = DhcpOptionsResourceDependencies + `
resource "oci_core_dhcp_options" "test_dhcp_options" {
	#Required
	compartment_id = "${var.compartment_id}"
	options {
		#Required
		server_type = "VcnLocalPlusInternet"
		type = "${var.dhcp_options_options_type}"
	}
	options {
		search_domain_names = [ "test.com" ]
		type = "SearchDomain"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
`

	DhcpOptionsResourceConfig = DhcpOptionsResourceDependencies + DhcpOptionsResourceConfigOnly

	DhcpOptionsResourceConfigOnly = `
resource "oci_core_dhcp_options" "test_dhcp_options" {
	#Required
	compartment_id = "${var.compartment_id}"
	options {
		#Required
		server_type = "VcnLocalPlusInternet"
		type = "${var.dhcp_options_options_type}"
	}
	options {
		search_domain_names = [ "test.com" ]
		type = "SearchDomain"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.dhcp_options_defined_tags_value}")}"
	display_name = "${var.dhcp_options_display_name}"
	freeform_tags = "${var.dhcp_options_freeform_tags}"
}
`
	DhcpOptionsPropertyVariables = `
variable "dhcp_options_defined_tags_value" { default = "value" }
variable "dhcp_options_display_name" { default = "MyDhcpOptions" }
variable "dhcp_options_freeform_tags" { default = {"Department"= "Finance"} }
variable "dhcp_options_options_type" { default = "DomainNameServer" }
variable "dhcp_options_state" { default = "AVAILABLE" }

`
	DhcpOptionsResourceDependencies = VcnPropertyVariables + VcnResourceConfig
)

func TestCoreDhcpOptionsResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_dhcp_options.test_dhcp_options"
	datasourceName := "data.oci_core_dhcp_options.test_dhcp_options"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreDhcpOptionsDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + DhcpOptionsPropertyVariables + compartmentIdVariableStr + DhcpOptionsRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DhcpOptionsResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + DhcpOptionsPropertyVariables + compartmentIdVariableStr + DhcpOptionsResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyDhcpOptions"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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
variable "dhcp_options_defined_tags_value" { default = "updatedValue" }
variable "dhcp_options_display_name" { default = "displayName2" }
variable "dhcp_options_freeform_tags" { default = {"Department"= "Accounting"} }
variable "dhcp_options_options_type" { default = "DomainNameServer" }
variable "dhcp_options_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + DhcpOptionsResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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
variable "dhcp_options_defined_tags_value" { default = "updatedValue" }
variable "dhcp_options_display_name" { default = "displayName2" }
variable "dhcp_options_freeform_tags" { default = {"Department"= "Accounting"} }
variable "dhcp_options_options_type" { default = "DomainNameServer" }
variable "dhcp_options_state" { default = "AVAILABLE" }

data "oci_core_dhcp_options" "test_dhcp_options" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.dhcp_options_display_name}"
	state = "${var.dhcp_options_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_dhcp_options.test_dhcp_options.id}"]
    }
}
                ` + compartmentIdVariableStr + DhcpOptionsResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "options.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "options.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "options.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "options.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "options.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "options.0.options.#", "2"),
					resource.TestCheckResourceAttr(datasourceName, "options.0.options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttrSet(datasourceName, "options.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "options.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "options.0.vcn_id"),
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

func testAccCheckCoreDhcpOptionsDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_dhcp_options" {
			noResourceFound = false
			request := oci_core.GetDhcpOptionsRequest{}

			tmp := rs.Primary.ID
			request.DhcpId = &tmp

			response, err := client.GetDhcpOptions(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.DhcpOptionsLifecycleStateTerminated): true,
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
