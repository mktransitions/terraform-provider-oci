// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"




	"github.com/stretchr/testify/suite"
)

type CoreInternetGatewayDatasourceTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *CoreInternetGatewayDatasourceTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_internet_gateways" "s" {
      compartment_id = "${var.compartment_id}"
      vcn_id = "vcnid"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_internet_gateways.s"

}

func (s *CoreInternetGatewayDatasourceTestSuite) TestResourceListInternetGateways() {

	s.Client.On(
		"ListInternetGateways",
		"compartmentid",
		"vcnid",
		&baremetal.ListOptions{},
	).Return(
		&baremetal.ListInternetGateways{
			Gateways: []baremetal.InternetGateway{
				{
					CompartmentID: "compartmentid",
					DisplayName:   "display_name",
					ID:            "id1",
					State:         baremetal.ResourceAvailable,
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
					ModifiedTime: baremetal.Time{
						Time: time.Now(),
					},
				},
				{
					CompartmentID: "compartmentid",
					DisplayName:   "display_name",
					ID:            "id2",
					State:         baremetal.ResourceAvailable,
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
					ModifiedTime: baremetal.Time{
						Time: time.Now(),
					},
				},
			},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "vcn_id", "vcnid"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.0.compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.#", "2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListInternetGateways", "compartmentid", "vcnid", &baremetal.ListOptions{})

}

func (s *CoreInternetGatewayDatasourceTestSuite) TestResourceListInternetGatewaysPaged() {

	res := &baremetal.ListInternetGateways{}
	res.NextPage = "nextpage"
	res.Gateways = []baremetal.InternetGateway{
		{
			CompartmentID: "compartmentid",
			DisplayName:   "display_name",
			ID:            "id1",
			State:         baremetal.ResourceAvailable,
			TimeCreated: baremetal.Time{
				Time: time.Now(),
			},
			ModifiedTime: baremetal.Time{
				Time: time.Now(),
			},
		},
		{
			CompartmentID: "compartmentid",
			DisplayName:   "display_name",
			ID:            "id2",
			State:         baremetal.ResourceAvailable,
			TimeCreated: baremetal.Time{
				Time: time.Now(),
			},
			ModifiedTime: baremetal.Time{
				Time: time.Now(),
			},
		},
	}

	s.Client.On(
		"ListInternetGateways",
		"compartmentid",
		"vcnid",
		&baremetal.ListOptions{},
	).Return(res, nil)

	opts := &baremetal.ListOptions{}
	opts.Page = "nextpage"
	s.Client.On(
		"ListInternetGateways",
		"compartmentid",
		"vcnid",
		opts,
	).Return(
		&baremetal.ListInternetGateways{
			Gateways: []baremetal.InternetGateway{
				{
					CompartmentID: "compartmentid",
					DisplayName:   "display_name",
					ID:            "id3",
					State:         baremetal.ResourceAvailable,
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
					ModifiedTime: baremetal.Time{
						Time: time.Now(),
					},
				},
				{
					CompartmentID: "compartmentid",
					DisplayName:   "display_name",
					ID:            "id4",
					State:         baremetal.ResourceAvailable,
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
					ModifiedTime: baremetal.Time{
						Time: time.Now(),
					},
				},
			},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "vcn_id", "vcnid"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.0.compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListInternetGateways", "compartmentid", "vcnid", opts)
}

func TestCoreInternetGatewayDatasource(t *testing.T) {
	suite.Run(t, new(CoreInternetGatewayDatasourceTestSuite))
}
