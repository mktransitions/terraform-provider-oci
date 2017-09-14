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
	Client       *baremetal.Client
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
		"oci": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = `
		resource "oci_core_drg" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
		}
		resource "oci_core_cpe" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "displayname"
      			ip_address = "123.123.123.123"
      			depends_on = ["oci_core_drg.t"]
		}
		resource "oci_core_ipsec" "t" {
			compartment_id = "${var.compartment_id}"
      			cpe_id = "${oci_core_cpe.t.id}"
      			drg_id = "${oci_core_drg.t.id}"
			display_name = "display_name"
      			static_routes = ["10.0.0.0/16"]
		}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "oci_core_ipsec.t"

}

func (s *ResourceCoreIPSecTestSuite) TestCreateResourceCoreIpsec() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_id"),

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
		},
	})
}

func (s *ResourceCoreIPSecTestSuite) TestTerminateIPSec() {

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

}

func TestResourceCoreIPSecTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreIPSecTestSuite))
}
