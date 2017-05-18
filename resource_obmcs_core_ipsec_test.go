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

type ResourceCoreIPSecTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.IPSecConnection
	DeletedRes   *baremetal.IPSecConnection
}

func (s *ResourceCoreIPSecTestSuite) SetupTest() {
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
		resource "baremetal_core_drg" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
		}
		resource "baremetal_core_cpe" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "displayname"
      			ip_address = "123.123.123.123"
		}
		resource "baremetal_core_ipsec" "t" {
			compartment_id = "${var.compartment_id}"
      			cpe_id = "${baremetal_core_cpe.t.id}"
      			drg_id = "${baremetal_core_drg.t.id}"
			display_name = "display_name"
      			static_routes = ["route1","route2"]
		}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_ipsec.t"
	s.Res = &baremetal.IPSecConnection{
		CompartmentID: "compartmentid",
		DisplayName:   "display_name",
		ID:            "id",
		DrgID:         "drgid",
		CpeID:         "cpeid",
		StaticRoutes:  []string{"route1", "route2"},
		TimeCreated:   s.TimeCreated,
		State:         baremetal.ResourceUp,
	}

	s.DeletedRes = s.Res
	s.DeletedRes.State = baremetal.ResourceDown

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"
	s.Client.On(
		"CreateIPSecConnection",
		s.Res.CompartmentID,
		s.Res.CpeID,
		s.Res.DrgID,
		s.Res.StaticRoutes,
		opts).Return(s.Res, nil)
	s.Client.On("DeleteIPSecConnection", s.Res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreIPSecTestSuite) TestCreateResourceCoreSubnet() {
	s.Client.On("GetIPSecConnection", "id").Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "drg_id", s.Res.DrgID),

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
				),
			},
		},
	})
}

func (s *ResourceCoreIPSecTestSuite) TestCreateResourceCoreSubnetWithoutDisplayName() {
	s.Client.On("GetIPSecConnection", "id").Return(s.Res, nil)

	s.Config = `

resource "baremetal_core_drg" "t" {
	compartment_id = "${var.compartment_id}"
	display_name = "display_name"
}
resource "baremetal_core_cpe" "t" {
	compartment_id = "${var.compartment_id}"
	display_name = "displayname"
      ip_address = "123.123.123.123"
}
  resource "baremetal_core_ipsec" "t" {
    compartment_id = "${var.compartment_id}"
    cpe_id = "${baremetal_core_cpe.t.id}"
    drg_id = "${baremetal_core_drg.t.id}"
    static_routes = ["route1","route2"]
  }
	`

	s.Config += testProviderConfig()

	opts := &baremetal.CreateOptions{}
	s.Res.DisplayName = ""

	s.Client.On(
		"CreateIPSecConnection",
		s.Res.CompartmentID,
		s.Res.CpeID,
		s.Res.DrgID,
		s.Res.StaticRoutes,
		opts).Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
				),
			},
		},
	})
}

func (s ResourceCoreIPSecTestSuite) TestUpdateCompartmentIDForcesNewIPSec() {
	s.Client.On("GetIPSecConnection", "id").Return(s.Res, nil)

	config := `
		resource "baremetal_core_ipsec" "t" {
    compartment_id = "new_compartment_id"
    cpe_id = "cpeid"
    drg_id = "drgid"
    display_name = "display_name"
    static_routes = ["route1","route2"]
		}
	`

	config += testProviderConfig()

	res := &baremetal.IPSecConnection{
		CompartmentID: "new_compartment_id",
		DisplayName:   "display_name",
		ID:            "new_id",
		DrgID:         "drgid",
		CpeID:         "cpeid",
		StaticRoutes:  []string{"route1", "route2"},
		TimeCreated:   s.TimeCreated,
		State:         baremetal.ResourceUp,
	}

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"
	s.Client.On(
		"CreateIPSecConnection",
		res.CompartmentID,
		res.CpeID,
		res.DrgID,
		res.StaticRoutes,
		opts).Return(res, nil).Once()

	s.Client.On("GetIPSecConnection", res.ID).Return(res, nil)
	s.Client.On("DeleteIPSecConnection", res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)

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

				),
			},
		},
	})
}

func (s *ResourceCoreIPSecTestSuite) TestTerminateIPSec() {
	s.Client.On("GetIPSecConnection", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetIPSecConnection", "id").Return(s.DeletedRes, nil)

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

	s.Client.On("DeleteIPSecConnection", s.Res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)

}

func TestResourceCoreIPSecTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreIPSecTestSuite))
}
