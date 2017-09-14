// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreIPSecConnectionsTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreIPSecConnectionsTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	resource "oci_core_drg" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-drg"
	}
	resource "oci_core_cpe" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-cpe"
		ip_address = "123.123.123.123"
	}
	resource "oci_core_ipsec" "t" {
		compartment_id = "${var.compartment_id}"
		cpe_id = "${oci_core_cpe.t.id}"
		drg_id = "${oci_core_drg.t.id}"
		display_name = "-tf-ipsec"
		static_routes = ["10.0.0.0/16"]
	}`
	s.ResourceName = "data.oci_core_ipsec_connections.s"
}

func (s *DatasourceCoreIPSecConnectionsTestSuite) TestAccDatasourceCoreIPConnections_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			// todo: investigate--there's some kind of consistency/sync issue here. If this extra step isn't
			// here the connections length asserts with 0, though the data is properly pulled in Get() at 
			// it first availability state and correctly synced in SetData(). Adding this multistep somehow 
			// overcomes it.
			{	
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config +
					`data "oci_core_ipsec_connections" "s" {
						compartment_id = "${var.compartment_id}"
						drg_id = "${oci_core_drg.t.id}"
						cpe_id = "${oci_core_cpe.t.id}"
						limit = 1
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "cpe_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "connections.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "connections.0.drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "connections.0.cpe_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "connections.0.compartment_id"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreIPSecConnectionsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreIPSecConnectionsTestSuite))
}
