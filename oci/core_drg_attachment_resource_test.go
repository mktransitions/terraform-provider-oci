// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"fmt"

	"github.com/oracle/oci-go-sdk/v25/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreDrgAttachmentTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreDrgAttachmentTestSuite) SetupTest() {
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig() + `
		resource "oci_core_virtual_network" "t" {
			cidr_block = "10.0.0.0/16"
			compartment_id = "${var.compartment_id}"
			display_name = "-tf-vcn"
		}
		resource "oci_core_drg" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "-tf-drg"
		}`

	s.ResourceName = "oci_core_drg_attachment.t"
}

func (s *ResourceCoreDrgAttachmentTestSuite) TestAccResourceCoreDrgAttachment_basic() {
	var resId, resId2 string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify a drg attachment can be created
			{
				Config: s.Config + `
				resource "oci_core_drg_attachment" "t" {
					drg_id = "${oci_core_drg.t.id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.DrgAttachmentLifecycleStateAttached)),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_core_drg.t", "id")
						return err
					},
				),
			},
			// verify drg attachment update
			{
				Config: s.Config + `
				resource "oci_core_drg_attachment" "t" {
					drg_id = "${oci_core_drg.t.id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					display_name = "-tf-drg-attachment"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-drg-attachment"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.DrgAttachmentLifecycleStateAttached)),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_core_drg.t", "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceCoreDrgAttachmentTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreDrgAttachmentTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceCoreDrgAttachmentTestSuite))
}
