// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"regexp"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/v25/identity"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityAvailabilityDomainsTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         identity.ListAvailabilityDomainsResponse
}

func (s *DatasourceIdentityAvailabilityDomainsTestSuite) SetupTest() {
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = testProviderConfig()
	s.ResourceName = "data.oci_identity_availability_domains.t"
}

func (s *DatasourceIdentityAvailabilityDomainsTestSuite) TestAccIdentityAvailabilityDomains_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			// Verify expected number of ADs in expected order. Expect this to fail in single AD regions
			{
				Config: s.Config + `
				data "oci_identity_availability_domains" "t" {
					compartment_id = "${var.tenancy_ocid}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domains.#", "3"),
					resource.TestMatchResourceAttr(s.ResourceName, "availability_domains.0.name", regexp.MustCompile(`\w*-AD-1`)),
					resource.TestMatchResourceAttr(s.ResourceName, "availability_domains.1.name", regexp.MustCompile(`\w*-AD-2`)),
					resource.TestMatchResourceAttr(s.ResourceName, "availability_domains.2.name", regexp.MustCompile(`\w*-AD-3`)),
				),
			},
			// Verify regex filtering
			{
				Config: s.Config + `
				data "oci_identity_availability_domains" "t" {
					compartment_id = "${var.tenancy_ocid}"
					filter {
						name = "name"
						values = ["\\w*-AD-2"]
						regex = true
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domains.#", "1"),
					resource.TestMatchResourceAttr(s.ResourceName, "availability_domains.0.name", regexp.MustCompile(".*AD-2")),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityAvailabilityDomainsTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceIdentityAvailabilityDomainsTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceIdentityAvailabilityDomainsTestSuite))
}
