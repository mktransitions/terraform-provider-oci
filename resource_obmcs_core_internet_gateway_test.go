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

type ResourceCoreInternetGatewayTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.InternetGateway
	DeletedRes   *baremetal.InternetGateway
}

func (s *ResourceCoreInternetGatewayTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = `
resource "baremetal_core_virtual_network" "t" {
	cidr_block = "10.0.0.0/16"
	compartment_id = "${var.compartment_id}"
	display_name = "display_name"
}

resource "baremetal_core_internet_gateway" "t" {
    compartment_id = "${var.compartment_id}"
    display_name = "display_name"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_internet_gateway.t"
}

func (s *ResourceCoreInternetGatewayTestSuite) TestCreateResourceCoreInternetGateway() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
		},
	})
}

func (s ResourceCoreInternetGatewayTestSuite) TestUpdateCompartmentIDForcesNewInternetGateway() {

	config := `
resource "baremetal_core_virtual_network" "t" {
	cidr_block = "10.0.0.0/16"
	compartment_id = "${var.compartment_id}"
	display_name = "display_name"
}

resource "baremetal_core_internet_gateway" "t" {
    compartment_id = "${var.compartment_id}"
    display_name = "CompleteIG2"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
}
	`

	config += testProviderConfig()

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "CompleteIG2"),
				),
			},
		},
	})
}

func (s *ResourceCoreInternetGatewayTestSuite) TestDeleteInternetGateway() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

}

func TestResourceCoreInternetGatewayTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreInternetGatewayTestSuite))
}
