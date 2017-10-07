// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVolumeAttachmentTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreVolumeAttachmentTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + instanceConfig + `
	resource "oci_core_volume" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-volume"
	}
	resource "oci_core_volume_attachment" "t" {
		attachment_type = "iscsi"
		compartment_id = "${var.compartment_id}"
		instance_id = "${oci_core_instance.t.id}"
		volume_id = "${oci_core_volume.t.id}"
	}`
	s.ResourceName = "data.oci_core_volume_attachments.t"
}

func (s *DatasourceCoreVolumeAttachmentTestSuite) TestAccDatasourceCoreVolumeAttachment_basic() {
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
				data "oci_core_volume_attachments" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					limit = 1
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.#", "1"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreVolumeAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreVolumeAttachmentTestSuite))
}
