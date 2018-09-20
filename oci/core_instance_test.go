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
	InstanceRequiredOnlyResource = InstanceResourceDependenciesRequiredOnly + `
resource "oci_core_instance" "test_instance" {
	#Required
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	compartment_id = "${var.compartment_id}"
	shape = "${var.instance_shape}"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
	image = "${var.InstanceImageOCID[var.region]}"
}
`
	InstanceResourceAsDependencyConfig = InstanceResourceDependenciesRequiredOnly + `
resource "oci_core_instance" "test_instance" {
	#Required
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	compartment_id = "${var.compartment_id}"
	shape = "${var.instance_shape}"
	image = "${var.InstanceImageOCID[var.region]}"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
	metadata {
		ssh_authorized_keys = "${var.ssh_public_key}"
	}

	timeouts {
		create = "15m"
	}
}
`

	InstanceResourceConfig = InstanceResourceDependencies + `
resource "oci_core_instance" "test_instance" {
	#Required
	availability_domain = "${oci_core_subnet.test_subnet.availability_domain}"
	compartment_id = "${var.compartment_id}"
	shape = "${var.instance_shape}"

	#Optional
	create_vnic_details {
		#Required
		subnet_id = "${oci_core_subnet.test_subnet.id}"

		#Optional
		assign_public_ip = "${var.instance_create_vnic_details_assign_public_ip}"
		defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.instance_create_vnic_details_defined_tags_value}")}"
		display_name = "${var.instance_create_vnic_details_display_name}"
		freeform_tags = "${var.instance_create_vnic_details_freeform_tags}"
		private_ip = "${var.instance_create_vnic_details_private_ip}"
		skip_source_dest_check = "${var.instance_create_vnic_details_skip_source_dest_check}"
	}
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.instance_defined_tags_value}")}"
	display_name = "${var.instance_display_name}"
	extended_metadata = "${var.instance_extended_metadata}"
	fault_domain = "${var.instance_fault_domain}"
	freeform_tags = "${var.instance_freeform_tags}"
	ipxe_script = "${var.instance_ipxe_script}"
	metadata = "${var.instance_metadata}"
	source_details {
		#Required
		source_type = "image"
		source_id = "${var.InstanceImageOCID[var.region]}"
	}
	subnet_id = "${oci_core_subnet.test_subnet.id}"
}
`
	InstanceCommonVariables = `
variable "InstanceImageOCID" {
	  type = "map"
	  default = {
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-7.4-2018.02.21-1"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaupbfz5f5hdvejulmalhyb6goieolullgkpumorbvxlwkaowglslq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlw3xfie2t5t52uegyhiq2npx7bqyu4uvi2zyu3w3mqayc2bxmaa"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaa7d3fsb6272srnftyi4dphdgfjf6gurxqhmv6ileds7ba3m2gltxq"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaa6h6gj6v4n56mqrbgnosskq63blyv2752g36zerymy63cfkojiiq"
	  }
}

`
	InstancePropertyVariables = InstanceCommonVariables + `
variable "instance_availability_domain" { default = "availabilityDomain" }
variable "instance_create_vnic_details_assign_public_ip" { default = false }
variable "instance_create_vnic_details_defined_tags_value" { default = "definedTags" }
variable "instance_create_vnic_details_display_name" { default = "displayName" }
variable "instance_create_vnic_details_freeform_tags" { default = {"Department"= "Accounting"} }
variable "instance_create_vnic_details_private_ip" { default = "10.0.0.5" }
variable "instance_create_vnic_details_skip_source_dest_check" { default = false }
variable "instance_defined_tags_value" { default = "value" }
variable "instance_display_name" { default = "displayName" }
variable "instance_extended_metadata" { default = {
		some_string = "stringA"
		nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
	} }
variable "instance_fault_domain" { default = "FAULT-DOMAIN-2" }
variable "instance_freeform_tags" { default = {"Department"= "Finance"} }
variable "instance_hostname_label" { default = "hostnameLabel" }
variable "instance_image" { default = "image" }
variable "instance_ipxe_script" { default = "ipxeScript" }
variable "instance_metadata" { default = { userdata = "abcd" } }
variable "instance_shape" { default = "VM.Standard1.8" }
variable "instance_source_details_source_type" { default = "image" }
variable "instance_state" { default = "RUNNING" }

`
	InstanceResourceDependenciesRequiredOnly = SubnetPropertyVariables + SubnetRequiredOnlyResource
	InstanceResourceDependencies             = DefinedTagsDependencies + InstanceResourceDependenciesRequiredOnly
)

func TestCoreInstanceResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance.test_instance"
	datasourceName := "data.oci_core_instances.test_instances"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + InstancePropertyVariables + compartmentIdVariableStr + InstanceRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard1.8"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + InstanceResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + InstancePropertyVariables + compartmentIdVariableStr + InstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard1.8"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + InstanceCommonVariables + `
variable "instance_availability_domain" { default = "availabilityDomain" }
variable "instance_create_vnic_details_assign_public_ip" { default = false }
variable "instance_create_vnic_details_defined_tags_value" { default = "definedTags" }
variable "instance_create_vnic_details_display_name" { default = "displayName" }
variable "instance_create_vnic_details_freeform_tags" { default = {"Department"= "Accounting"} }
variable "instance_create_vnic_details_private_ip" { default = "10.0.0.5" }
variable "instance_create_vnic_details_skip_source_dest_check" { default = false }
variable "instance_defined_tags_value" { default = "updatedValue" }
variable "instance_display_name" { default = "displayName2" }
variable "instance_extended_metadata" { default = {
		some_string = "stringA"
		nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
	} }
variable "instance_fault_domain" { default = "FAULT-DOMAIN-2" }
variable "instance_freeform_tags" { default = {"Department"= "Accounting"} }
variable "instance_hostname_label" { default = "hostnameLabel" }
variable "instance_image" { default = "image" }
variable "instance_ipxe_script" { default = "ipxeScript" }
variable "instance_metadata" { default = { userdata = "abcd" } }
variable "instance_shape" { default = "VM.Standard1.8" }
variable "instance_source_details_source_type" { default = "image" }
variable "instance_state" { default = "RUNNING" }

                ` + compartmentIdVariableStr + InstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "extended_metadata.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "image"),
					resource.TestCheckResourceAttr(resourceName, "ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "region"),
					resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard1.8"),
					resource.TestCheckResourceAttr(resourceName, "source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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
				Config: config + InstanceCommonVariables + `
variable "instance_availability_domain" { default = "availabilityDomain" }
variable "instance_create_vnic_details_assign_public_ip" { default = false }
variable "instance_create_vnic_details_defined_tags_value" { default = "definedTags" }
variable "instance_create_vnic_details_display_name" { default = "displayName" }
variable "instance_create_vnic_details_freeform_tags" { default = {"Department"= "Accounting"} }
variable "instance_create_vnic_details_private_ip" { default = "10.0.0.5" }
variable "instance_create_vnic_details_skip_source_dest_check" { default = false }
variable "instance_defined_tags_value" { default = "updatedValue" }
variable "instance_display_name" { default = "displayName2" }
variable "instance_extended_metadata" { default = {
		some_string = "stringA"
		nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
	} }
variable "instance_fault_domain" { default = "FAULT-DOMAIN-2" }
variable "instance_freeform_tags" { default = {"Department"= "Accounting"} }
variable "instance_hostname_label" { default = "hostnameLabel" }
variable "instance_image" { default = "image" }
variable "instance_ipxe_script" { default = "ipxeScript" }
variable "instance_metadata" { default = { userdata = "abcd" } }
variable "instance_shape" { default = "VM.Standard1.8" }
variable "instance_source_details_source_type" { default = "image" }
variable "instance_state" { default = "RUNNING" }

data "oci_core_instances" "test_instances" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${oci_core_instance.test_instance.availability_domain}"
	display_name = "${var.instance_display_name}"
	state = "${var.instance_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_instance.test_instance.id}"]
    }
}
                ` + compartmentIdVariableStr + InstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "RUNNING"),

					resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.extended_metadata.%", "2"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.fault_domain", "FAULT-DOMAIN-2"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.image"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.ipxe_script", "ipxeScript"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.metadata.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "VM.Standard1.8"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.0.source_type", "image"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.time_created"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"create_vnic_details",
					// TODO: extended_metadata intentionally not set in resource Gets, even though supported
					// by GetInstance calls. Remove this when the issue is resolved.
					"extended_metadata",
					"hostname_label",
					"subnet_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckCoreInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_instance" {
			noResourceFound = false
			request := oci_core.GetInstanceRequest{}

			tmp := rs.Primary.ID
			request.InstanceId = &tmp

			response, err := client.GetInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.InstanceLifecycleStateTerminated): true,
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
