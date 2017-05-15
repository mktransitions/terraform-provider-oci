// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreDHCPOptionsDatasourceTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListDHCPOptions
}

func (s *ResourceCoreDHCPOptionsDatasourceTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_dhcp_options" "t" {
      compartment_id = "${var.compartment_id}"
      limit = 1
      page = "page"
      vcn_id = "vcn_id"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_dhcp_options.t"
}

func (s *ResourceCoreDHCPOptionsDatasourceTestSuite) TestReadDHCPOptions() {
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "vcn_id", "vcn_id"),
				),
			},
		},
	},
	)

}

func TestResourceCoreDHCPOptionsDatasourceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDHCPOptionsDatasourceTestSuite))
}
