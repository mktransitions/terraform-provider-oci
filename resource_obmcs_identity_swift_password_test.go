// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type ResourceIdentitySwiftPasswordTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentitySwiftPasswordTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	resource "oci_identity_user" "t" {
		name = "tf-user"
		description = "tf test user"
	}`

	s.ResourceName = "oci_identity_swift_password.t"
}

func (s *ResourceIdentitySwiftPasswordTestSuite) TestAccResourceIdentitySwiftPassword_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_identity_swift_password" "t" {
					user_id = "${oci_identity_user.t.id}"
					description = "tf test swift password"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "password"),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "tf test swift password"),
				),
			},
			// verify update
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_identity_swift_password" "t" {
					user_id = "${oci_identity_user.t.id}"
					description = "tf test swift password (updated)"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "tf test swift password (updated)"),
				),
			},
		},
	})
}

func TestResourceIdentitySwiftPasswordTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentitySwiftPasswordTestSuite))
}
