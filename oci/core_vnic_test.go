// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	vnicSingularDataSourceRepresentation = map[string]interface{}{
		"vnic_id": Representation{repType: Required, create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}`},
	}

	VnicResourceConfig = ``
)

func TestCoreVnicResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_vnic.test_vnic"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + instanceDnsConfig + `

data "oci_core_vnic_attachments" "t" {
	compartment_id = "${var.compartment_id}"
	instance_id = "${oci_core_instance.t.id}"
}` +
					generateDataSourceFromRepresentationMap("oci_core_vnic", "test_vnic", Required, Create, vnicSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VnicResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vnic_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "-tf-instance-vnic"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "hostname_label", "testinstance"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_primary", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "mac_address"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip_address"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip_address"),
					resource.TestCheckResourceAttr(singularDatasourceName, "skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
		},
	})
}
