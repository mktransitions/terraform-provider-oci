// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumeBackupTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Res          *baremetal.VolumeBackup
	DeletedRes   *baremetal.VolumeBackup
}

func (s *ResourceCoreVolumeBackupTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
		data "oci_identity_availability_domains" "ADs" {
  			compartment_id = "${var.compartment_id}"
		}
		resource "oci_core_volume" "t" {
			availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
			compartment_id = "${var.compartment_id}"
			display_name = "-tf-volume"
			size_in_mbs = 51200
		}`
	s.ResourceName = "oci_core_volume_backup.t"
}

func (s *ResourceCoreVolumeBackupTestSuite) TestCreateVolumeBackup() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_volume_backup" "t" {
						volume_id = "${oci_core_volume.t.id}"
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
				),
			},
			// verify update
			{
				Config: s.Config + `
					resource "oci_core_volume_backup" "t" {
						volume_id = "${oci_core_volume.t.id}"
						display_name = "-tf-volume-backup"
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-volume-backup"),
				),
			},
			// verify restore
			{
				Config: s.Config + `
					resource "oci_core_volume_backup" "t" {
						volume_id = "${oci_core_volume.t.id}"
						display_name = "-tf-volume-backup"
					}
					resource "oci_core_volume" "t2" {
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-volume-restored"
						size_in_mbs = 51200
						volume_backup_id = "${oci_core_volume_backup.t.id}"
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("oci_core_volume.t2", "display_name", "-tf-volume-restored"),
				),
			},
		},
	})
}

func TestResourceCoreVolumeBackupTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeBackupTestSuite))
}
