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

type ResourceCoreRouteTableTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.RouteTable
	DeletedRes   *baremetal.RouteTable
}

func (s *ResourceCoreRouteTableTestSuite) SetupTest() {
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
resource "baremetal_core_internet_gateway" "CompleteIG" {
    compartment_id = "${var.compartment_id}"
    display_name = "CompleteIG"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
}
resource "baremetal_core_route_table" "t" {
	compartment_id = "${var.compartment_id}"
	display_name = "display_name"
	route_rules {
		cidr_block = "0.0.0.0/0"
		network_entity_id = "${baremetal_core_internet_gateway.CompleteIG.id}"
	}
	vcn_id = "${baremetal_core_virtual_network.t.id}"
}
	`
	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_route_table.t"

	routeRules := []baremetal.RouteRule{
		{
			CidrBlock:       "cidr_block",
			NetworkEntityID: "network_entity_id",
		},
		{
			CidrBlock:       "cidr_block",
			NetworkEntityID: "network_entity_id",
		},
	}

	s.Res = &baremetal.RouteTable{
		CompartmentID: "compartment_id",
		DisplayName:   "display_name",
		ID:            "id",
		RouteRules:    routeRules,
		TimeModified:  s.TimeCreated,
		State:         baremetal.ResourceAvailable,
		TimeCreated:   s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	s.DeletedRes = &baremetal.RouteTable{
		CompartmentID: "compartment_id",
		DisplayName:   "display_name",
		ID:            "id",
		RouteRules:    routeRules,
		TimeModified:  s.TimeCreated,
		State:         baremetal.ResourceTerminated,
		TimeCreated:   s.TimeCreated,
	}
	s.DeletedRes.ETag = "etag"
	s.DeletedRes.RequestID = "opcrequestid"

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"
	s.Client.On(
		"CreateRouteTable",
		"compartment_id",
		"vcn_id",
		routeRules,
		opts).Return(s.Res, nil)
	s.Client.On("DeleteRouteTable", "id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreRouteTableTestSuite) TestCreateResourceCoreRouteTable() {
	s.Client.On("GetRouteTable", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetRouteTable", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "route_rules.0.network_entity_id"),
				),
			},
		},
	})
}

func (s ResourceCoreRouteTableTestSuite) TestUpdateRouteTable() {
	if IsAccTest() {
		s.T().Skip()
	}
	s.Client.On("GetRouteTable", "id").Return(s.Res, nil).Times(3)

	config := `
		resource "baremetal_core_route_table" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
      route_rules {
				cidr_block = "new_cidr_block"
				network_entity_id = "network_entity_id"
			}
			vcn_id = "vcn_id"
		}
	`
	config += testProviderConfig()

	routeRules := []baremetal.RouteRule{
		{
			CidrBlock:       "new_cidr_block",
			NetworkEntityID: "network_entity_id",
		},
	}

	res := &baremetal.RouteTable{
		CompartmentID: "compartment_id",
		DisplayName:   "display_name",
		ID:            "id",
		RouteRules:    routeRules,
		TimeModified:  s.TimeCreated,
		State:         baremetal.ResourceAvailable,
		TimeCreated:   s.TimeCreated,
	}
	res.ETag = "etag"
	res.RequestID = "opcrequestid"

	opts := &baremetal.UpdateRouteTableOptions{}
	opts.RouteRules = routeRules
	s.Client.On("UpdateRouteTable", "id", opts).Return(res, nil)
	s.Client.On("GetRouteTable", "id").Return(res, nil).Times(2)
	s.Client.On("GetRouteTable", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check:  resource.TestCheckResourceAttr(s.ResourceName, "route_rules.0.cidr_block", "new_cidr_block"),
			},
		},
	})
}

func (s *ResourceCoreRouteTableTestSuite) TestDeleteRouteTable() {
	s.Client.On("GetRouteTable", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetRouteTable", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
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

	s.Client.AssertCalled(s.T(), "DeleteRouteTable", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreRouteTableTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreRouteTableTestSuite))
}
