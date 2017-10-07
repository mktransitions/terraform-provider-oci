// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVolumeBackupTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListVolumeBackups
}

func (s *DatasourceCoreVolumeBackupTestSuite) SetupTest() {
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
	}
	resource "oci_core_volume_backup" "t" {
		volume_id = "${oci_core_volume.t.id}"
		display_name = "-tf-volume-backup"
	}`
	s.ResourceName = "data.oci_core_volume_backups.t"
}

func (s *DatasourceCoreVolumeBackupTestSuite) TestAccDatasourceCoreVolumeBackup_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				data "oci_core_volume_backups" "t" {
					compartment_id = "${var.compartment_id}"
					volume_id = "${oci_core_volume.t.id}"
					limit = 1
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_backups.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_backups.0.volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_backups.0.time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.display_name", "-tf-volume-backup"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.state", baremetal.ResourceAvailable),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_backups.0.unique_size_in_mbs"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_backups.0.unique_size_in_gbs"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreVolumeBackupTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreVolumeBackupTestSuite))
}
