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
	//	"errors"
	"errors"
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
	s.Res = &baremetal.InternetGateway{
		CompartmentID: "compartment_id",
		DisplayName:   "display_name",
		ID:            "id",
		IsEnabled:     true,
		State:         baremetal.ResourceAvailable,
		ModifiedTime:  s.TimeCreated,
		TimeCreated:   s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "requestid"

	s.DeletedRes = &baremetal.InternetGateway{}
	*s.DeletedRes = *s.Res
	s.DeletedRes.State = baremetal.ResourceTerminated

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"
	s.Client.On(
		"CreateInternetGateway",
		s.Res.CompartmentID,
		"vcnid",
		s.Res.IsEnabled,
		opts).Return(s.Res, nil)
	s.Client.On("DeleteInternetGateway", s.Res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreInternetGatewayTestSuite) TestCreateResourceCoreInternetGateway() {
	s.Client.On("GetInternetGateway", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetInternetGateway", "id").Return(s.DeletedRes, nil).Times(2)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
		},
	})
}

func (s ResourceCoreInternetGatewayTestSuite) TestUpdateCompartmentIDForcesNewInternetGateway() {
	s.Client.On("GetInternetGateway", s.Res.ID).Return(s.Res, nil).Times(2)
	s.Client.On("GetInternetGateway", s.Res.ID).Return(s.DeletedRes, nil).Times(2)

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

	res := &baremetal.InternetGateway{
		CompartmentID: "new_compartment_id",
		DisplayName:   "display_name",
		ID:            "id",
		IsEnabled:     true,
		State:         baremetal.ResourceAvailable,
		ModifiedTime:  s.TimeCreated,
		TimeCreated:   s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "requestid"

	delRes := &baremetal.InternetGateway{}
	*delRes = *res
	delRes.State = baremetal.ResourceTerminated

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"
	s.Client.On(
		"CreateInternetGateway",
		res.CompartmentID,
		"vcnid",
		res.IsEnabled,
		opts).Return(res, nil)

	s.Client.On("DeleteInternetGateway", res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)

	s.Client.On("GetInternetGateway", res.ID).Return(res, nil).Times(2)
	s.Client.On("GetInternetGateway", res.ID).Return(delRes, nil).Once()
	s.Client.On("GetInternetGateway", res.ID).Return(nil, errors.New("blah does not exist"))

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
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "CompleteIG2"),
				),
			},
		},
	})
}

func (s *ResourceCoreInternetGatewayTestSuite) TestDeleteInternetGateway() {
	s.Client.On("GetInternetGateway", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetInternetGateway", "id").Return(s.DeletedRes, nil)

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

	s.Client.AssertCalled(s.T(), "DeleteInternetGateway", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreInternetGatewayTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreInternetGatewayTestSuite))
}
