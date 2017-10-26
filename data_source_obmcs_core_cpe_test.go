// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreCpeTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	Token        string
	TokenFn      TokenFn
}

func (s *DatasourceCoreCpeTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + s.TokenFn(`
	resource "oci_core_cpe" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "{{.token}}"
		ip_address = "142.10.10.1"
	}`, nil)
	s.ResourceName = "data.oci_core_cpes.s"
}

func (s *DatasourceCoreCpeTestSuite) TestAccDatasourceCoreCpe_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + s.TokenFn(`
				data "oci_core_cpes" "s" {
					compartment_id = "${oci_core_cpe.t.compartment_id}"
					filter {
						name   = "display_name"
						values = ["{{.token}}"]
					}
				}`, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.0.display_name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.0.ip_address", "142.10.10.1"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreCpeTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreCpeTestSuite))
}
