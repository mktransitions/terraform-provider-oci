// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreDrgsTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreDrgsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}
	s.Config = `
	resource "oci_core_drg" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
	}
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.oci_core_drgs.t"
}

func (s *ResourceCoreDrgsTestSuite) TestReadDrgs() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				data "oci_core_drgs" "t" {
					compartment_id = "${var.compartment_id}"
					limit = 1
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "drgs.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drgs.#"),
				),
			},
		},
	},
	)

}

func TestDatasourceCoreDrgsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDrgsTestSuite))
}
