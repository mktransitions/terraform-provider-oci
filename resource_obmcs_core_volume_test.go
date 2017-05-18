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
	//"strconv"
)

type ResourceCoreVolumeTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Volume
	DeletedRes   *baremetal.Volume
}

func (s *ResourceCoreVolumeTestSuite) SetupTest() {
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
		data "baremetal_identity_availability_domains" "ADs" {
  			compartment_id = "${var.compartment_id}"
		}
		resource "baremetal_core_volume" "t" {
			availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
			size_in_mbs = 262144
		}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_volume.t"
	s.Res = &baremetal.Volume{
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartment_id",
		DisplayName:        "display_name",
		ID:                 "id",
		SizeInMBs:          123,
		State:              baremetal.ResourceAvailable,
		TimeCreated:        s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	s.DeletedRes = &baremetal.Volume{
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartment_id",
		DisplayName:        "display_name",
		ID:                 "id",
		SizeInMBs:          123,
		State:              baremetal.ResourceTerminated,
		TimeCreated:        s.TimeCreated,
	}
	s.DeletedRes.ETag = "etag"
	s.DeletedRes.RequestID = "opcrequestid"

	opts := &baremetal.CreateVolumeOptions{}
	opts.DisplayName = "display_name"
	opts.SizeInMBs = 123
	s.Client.On(
		"CreateVolume",
		"availability_domain",
		"compartment_id",
		opts).Return(s.Res, nil)
	s.Client.On("DeleteVolume", "id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreVolumeTestSuite) TestCreateResourceCoreVolume() {
	s.Client.On("GetVolume", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetVolume", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", s.Res.AvailabilityDomain),

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					//resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", strconv.Itoa(s.Res.SizeInMBs)),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
				),
			},
		},
	})
}

func (s *ResourceCoreVolumeTestSuite) TestCreateResourceCoreVolumeWithoutDisplayName() {
	s.Client.On("GetVolume", "id").Return(s.Res, nil)

	s.Config = `
		resource "baremetal_core_volume" "t" {
			availability_domain = "availability_domain"
			compartment_id = "${var.compartment_id}"
			size_in_mbs = 123
		}
	`
	s.Config += testProviderConfig()

	opts := &baremetal.CreateVolumeOptions{SizeInMBs: 123}
	s.Client.On(
		"CreateVolume",
		"availability_domain",
		"compartment_id", opts).Return(s.Res, nil)

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

func (s ResourceCoreVolumeTestSuite) TestUpdateVolumeDisplayName() {
	s.Client.On("GetVolume", "id").Return(s.Res, nil).Times(3)

	config := `
		resource "baremetal_core_volume" "t" {
			availability_domain = "availability_domain"
			compartment_id = "${var.compartment_id}"
			display_name = "new_display_name"
			size_in_mbs = 123
		}
	`
	config += testProviderConfig()

	res := &baremetal.Volume{
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartment_id",
		DisplayName:        "new_display_name",
		ID:                 "id",
		SizeInMBs:          123,
		State:              baremetal.ResourceAvailable,
		TimeCreated:        s.TimeCreated,
	}
	res.ETag = "etag"
	res.RequestID = "opcrequestid"

	opts := &baremetal.UpdateOptions{}
	opts.DisplayName = "new_display_name"
	s.Client.On("UpdateVolume", "id", opts).Return(res, nil)
	s.Client.On("GetVolume", "id").Return(res, nil)

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
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", res.DisplayName),
				),
			},
		},
	})
}

func (s ResourceCoreVolumeTestSuite) TestUpdateAvailabilityDomainForcesNewVolume() {
	s.Client.On("GetVolume", "id").Return(s.Res, nil)

	config := `
		resource "baremetal_core_volume" "t" {
			availability_domain = "new_availability_domain"
			compartment_id = "${var.compartment_id}"
			size_in_mbs = 123
		}
  `
	config += testProviderConfig()

	res := &baremetal.Volume{
		AvailabilityDomain: "new_availability_domain",
		CompartmentID:      "compartment_id",
		DisplayName:        "display_name",
		ID:                 "new_id",
		SizeInMBs:          123,
		State:              baremetal.ResourceAvailable,
		TimeCreated:        s.TimeCreated,
	}
	res.ETag = "etag"
	res.RequestID = "opcrequestid"

	opts := &baremetal.CreateVolumeOptions{SizeInMBs: 123}
	s.Client.On(
		"CreateVolume",
		res.AvailabilityDomain,
		res.CompartmentID, opts).Return(res, nil)

	s.Client.On("GetVolume", res.ID).Return(res, nil)
	s.Client.On("DeleteVolume", res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)

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
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", res.AvailabilityDomain),
				),
			},
		},
	})
}

func (s ResourceCoreVolumeTestSuite) TestUpdateCompartmentIdForcesNewVolume() {
	s.Client.On("GetVolume", "id").Return(s.Res, nil)

	config := `
		resource "baremetal_core_volume" "t" {
			availability_domain = "availability_domain"
			compartment_id = "new_compartment_id"
			size_in_mbs = 123
		}
  `
	config += testProviderConfig()

	res := &baremetal.Volume{
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "new_compartment_id",
		DisplayName:        "display_name",
		ID:                 "new_id",
		SizeInMBs:          123,
		State:              baremetal.ResourceAvailable,
		TimeCreated:        s.TimeCreated,
	}
	res.ETag = "etag"
	res.RequestID = "opcrequestid"

	opts := &baremetal.CreateVolumeOptions{SizeInMBs: 123}
	s.Client.On(
		"CreateVolume",
		res.AvailabilityDomain,
		res.CompartmentID, opts).Return(res, nil)

	s.Client.On("GetVolume", res.ID).Return(res, nil)
	s.Client.On("DeleteVolume", res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)

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

func (s *ResourceCoreVolumeTestSuite) TestDeleteVolume() {
	s.Client.On("GetVolume", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetVolume", "id").Return(s.DeletedRes, nil)

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

	s.Client.AssertCalled(s.T(), "DeleteVolume", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreVolumeTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeTestSuite))
}
