// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreImageTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreImageTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig()
	s.ResourceName = "data.oci_core_images.t"
}

func (s *DatasourceCoreImageTestSuite) TestAccImage_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_core_images" "t" {
					compartment_id = "${var.compartment_id}"
					operating_system = "Oracle Linux"
					operating_system_version = "7.3"
				
					filter {
						name = "display_name"
						values = [".*2017.07.17-1"]
						regex = true
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "images.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.create_image_allowed", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.display_name", "Oracle-Linux-7.3-2017.07.17-1"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.state", "AVAILABLE"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.operating_system", "Oracle Linux"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.operating_system_version", "7.3"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.time_created"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreImageTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreImageTestSuite))
}
