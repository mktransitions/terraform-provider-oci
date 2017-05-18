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

type ResourceCoreVolumeBackupTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.VolumeBackup
	DeletedRes   *baremetal.VolumeBackup
}

func (s *ResourceCoreVolumeBackupTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)
	s.Providers = map[string]terraform.ResourceProvider{"baremetal": s.Provider}

	s.ResourceName = "baremetal_core_volume_backup.t"
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
		resource "baremetal_core_volume_backup" "t" {
			volume_id = "${baremetal_core_volume.t.id}"
			display_name = "display_name"
		}
	`
	s.Config += testProviderConfig()

	s.TimeCreated = baremetal.Time{Time: time.Now()}
	s.Res = &baremetal.VolumeBackup{
		CompartmentID:       "compartment_id",
		DisplayName:         "display_name",
		ID:                  "id",
		State:               baremetal.ResourceAvailable,
		SizeInMBs:           1,
		TimeCreated:         s.TimeCreated,
		TimeRequestReceived: s.TimeCreated,
		UniqueSizeInMBs:     1,
		VolumeID:            "volume_id",
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	deletedRes := *s.Res
	s.DeletedRes = &deletedRes
	s.DeletedRes.State = baremetal.ResourceTerminated

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"
	s.Client.On("CreateVolumeBackup", "volume_id", opts).Return(s.Res, nil)
	s.Client.On("DeleteVolumeBackup", "id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreVolumeBackupTestSuite) TestCreateVolumeBackup() {
	s.Client.On("GetVolumeBackup", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetVolumeBackup", "id").Return(s.DeletedRes, nil)

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
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
				),
			},
		},
	})
}

func (s *ResourceCoreVolumeBackupTestSuite) TestCreateVolumeBackupWithoutDisplayName() {
	if IsAccTest() {
		s.T().Skip()
	}
	s.Client.On("GetVolumeBackup", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetVolumeBackup", "id").Return(s.DeletedRes, nil)

	s.Config = `
		resource "baremetal_core_volume_backup" "t" {
			volume_id = "volume_id"
		}
	`
	s.Config += testProviderConfig()

	opts := &baremetal.CreateOptions{}
	s.Client.On("CreateVolumeBackup", "volume_id", opts).
		Return(s.Res, nil)

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

func (s ResourceCoreVolumeBackupTestSuite) TestUpdateVolumeBackupDisplayName() {
	if IsAccTest() {
		s.T().Skip()
	}
	s.Client.On("GetVolumeBackup", "id").Return(s.Res, nil).Times(3)

	config := `
		resource "baremetal_core_volume_backup" "t" {
			volume_id = "volume_id"
			display_name = "new_display_name"
		}
	`
	config += testProviderConfig()

	resVal := *s.Res
	res := &resVal
	res.DisplayName = "new_display_name"

	deletedResVal := *res
	deletedRes := &deletedResVal
	deletedRes.State = baremetal.ResourceTerminated

	opts := &baremetal.IfMatchDisplayNameOptions{}
	opts.DisplayName = "new_display_name"
	s.Client.On("UpdateVolumeBackup", "id", opts).Return(res, nil)
	s.Client.On("GetVolumeBackup", "id").Return(res, nil).Times(2)
	s.Client.On("GetVolumeBackup", "id").Return(deletedRes, nil)

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

func (s ResourceCoreVolumeBackupTestSuite) TestUpdateVolumeIDForcesNewVolumeBackup() {
	if IsAccTest() {
		s.T().Skip()
	}
	s.Client.On("GetVolumeBackup", "id").Return(s.Res, nil).Times(3)
	s.Client.On("GetVolumeBackup", "id").Return(s.DeletedRes, nil).Once()

	config := `
		resource "baremetal_core_volume_backup" "t" {
			volume_id = "new_volume_id"
		}
  `
	config += testProviderConfig()

	resVal := *s.Res
	res := &resVal
	res.VolumeID = "new_volume_id"

	deletedResVal := *res
	deletedRes := &deletedResVal
	deletedRes.State = baremetal.ResourceTerminated

	opts := &baremetal.CreateOptions{}
	s.Client.On("CreateVolumeBackup", res.VolumeID, opts).Return(res, nil)
	s.Client.On("GetVolumeBackup", "id").Return(res, nil).Times(2)
	s.Client.On("GetVolumeBackup", "id").Return(deletedRes, nil)

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
					resource.TestCheckResourceAttr(s.ResourceName, "volume_id", res.VolumeID),
				),
			},
		},
	})
}

func (s *ResourceCoreVolumeBackupTestSuite) TestDeleteVolumeBackup() {
	s.Client.On("GetVolumeBackup", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetVolumeBackup", "id").Return(s.DeletedRes, nil)

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

	s.Client.AssertCalled(s.T(), "DeleteVolumeBackup", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreVolumeBackupTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeBackupTestSuite))
}
