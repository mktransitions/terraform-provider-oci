// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

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
	PrivateIpRequiredOnlyResource = PrivateIpResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Required, Create, privateIpRepresentation)

	privateIpDataSourceRepresentation = map[string]interface{}{
		"vnic_id": Representation{repType: Optional, create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}`},
		"filter":  RepresentationGroup{Required, privateIpDataSourceFilterRepresentation}}
	privateIpDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_private_ip.test_private_ip.id}`}},
	}

	privateIpRepresentation = map[string]interface{}{
		"vnic_id":        Representation{repType: Required, create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"hostname_label": Representation{repType: Optional, create: `privateiptestinstance`, update: `privateiptestinstance2`},
		"ip_address":     Representation{repType: Optional, create: `10.0.1.5`},
	}

	PrivateIpResourceDependencies = instanceDnsConfig + `
	data "oci_core_vnic_attachments" "t" {
		availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		instance_id = "${oci_core_instance.t.id}"
	}

` + AvailabilityDomainConfig
)

func TestCorePrivateIpResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_private_ip.test_private_ip"
	datasourceName := "data.oci_core_private_ips.test_private_ips"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCorePrivateIpDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + PrivateIpResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Required, Create, privateIpRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + PrivateIpResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + PrivateIpResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Optional, Create, privateIpRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "privateiptestinstance"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.1.5"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + PrivateIpResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Optional, Update, privateIpRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "privateiptestinstance2"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.1.5"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

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
					generateDataSourceFromRepresentationMap("oci_core_private_ips", "test_private_ips", Optional, Update, privateIpDataSourceRepresentation) +
					compartmentIdVariableStr + PrivateIpResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_private_ip", "test_private_ip", Optional, Update, privateIpRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_id"),

					resource.TestCheckResourceAttr(datasourceName, "private_ips.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.hostname_label", "privateiptestinstance2"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.ip_address", "10.0.1.5"),
					resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.vnic_id"),
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

func testAccCheckCorePrivateIpDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_private_ip" {
			noResourceFound = false
			request := oci_core.GetPrivateIpRequest{}

			tmp := rs.Primary.ID
			request.PrivateIpId = &tmp

			_, err := client.GetPrivateIp(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
