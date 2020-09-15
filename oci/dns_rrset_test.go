// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/oci-go-sdk/v25/common"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	oci_dns "github.com/oracle/oci-go-sdk/v25/dns"
)

var (
	RrsetRequiredOnlyResource = RrsetResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Required, Create, rrsetRepresentation)

	RrsetResourceConfig = RrsetResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Update, rrsetRepresentation)

	rrsetSingularDataSourceRepresentation = map[string]interface{}{
		"domain":          Representation{repType: Required, create: dnsDomainName},
		"rtype":           Representation{repType: Required, create: `A`},
		"zone_name_or_id": Representation{repType: Required, create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  Representation{repType: Optional, create: `${var.compartment_id}`},
	}

	dnsDomainName       = randomString(5, charsetWithoutDigits) + ".token.oci-record-test"
	rrsetRepresentation = map[string]interface{}{
		"domain":          Representation{repType: Required, create: dnsDomainName},
		"rtype":           Representation{repType: Required, create: `A`},
		"zone_name_or_id": Representation{repType: Required, create: `${oci_dns_zone.test_zone.id}`},
		"compartment_id":  Representation{repType: Optional, create: `${var.compartment_id}`},
		"items":           RepresentationGroup{Optional, rrsetItemsRepresentation},
	}
	rrsetItemsRepresentation = map[string]interface{}{
		"domain": Representation{repType: Required, create: dnsDomainName},
		"rdata":  Representation{repType: Required, create: `192.168.0.1`, update: `77.77.77.77`},
		"rtype":  Representation{repType: Required, create: `A`},
		"ttl":    Representation{repType: Required, create: `3600`, update: `1000`},
	}

	RrsetResourceDependencies = `
data "oci_identity_tenancy" "test_tenancy" {
	tenancy_id = "${var.tenancy_ocid}"
}

resource "oci_dns_zone" "test_zone" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "` + dnsDomainName + `"
	zone_type = "PRIMARY"
}
`
)

func TestDnsRrsetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsRrsetResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_rrset.test_rrset"

	singularDatasourceName := "data.oci_dns_rrset.test_rrset"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDnsRrsetDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Required, Create, rrsetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Create, rrsetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "192.168.0.1",
						"rtype":  "A",
						"ttl":    "3600",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + RrsetResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Optional, Update, rrsetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "77.77.77.77",
						"rtype":  "A",
						"ttl":    "1000",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_dns_rrset", "test_rrset", Required, Create, rrsetSingularDataSourceRepresentation) +
					compartmentIdVariableStr + RrsetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "domain", dnsDomainName),
					resource.TestCheckResourceAttr(singularDatasourceName, "rtype", "A"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "zone_name_or_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{
						"domain": dnsDomainName,
						"rdata":  "77.77.77.77",
						"rtype":  "A",
						"ttl":    "1000",
					},
						[]string{
							"is_protected",
							"record_hash",
							"rrset_version",
						}),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + RrsetResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"compartment_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckDnsRrsetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dnsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_rrset" {
			noResourceFound = false
			request := oci_dns.GetRRSetRequest{}

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}

			if value, ok := rs.Primary.Attributes["domain"]; ok {
				request.Domain = &value
			}

			if value, ok := rs.Primary.Attributes["rtype"]; ok {
				request.Rtype = &value
			}

			if value, ok := rs.Primary.Attributes["zone_name_or_id"]; ok {
				request.ZoneNameOrId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "dns")

			_, err := client.GetRRSet(context.Background(), request)

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
