// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUserCapabilitiesManagementTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentityUserCapabilitiesManagementTestSuite) SetupTest() {
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig()

	s.ResourceName = "oci_identity_user_capabilities_management.t"
}

func (s *ResourceIdentityUserCapabilitiesManagementTestSuite) TestAccResourceIdentityUserCapabilitiesManagement_basic() {
	_, tokenFn := tokenize()
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create with capabilities all as false
			{
				Config: s.Config +
					tokenFn(
						`resource "oci_identity_user" "user" {
							name = "{{.token}}"
							description = "{{.description}}"
							compartment_id = "${var.tenancy_ocid}"
						}
						resource "oci_identity_user_capabilities_management" "t" {
							user_id = "${oci_identity_user.user.id}"
							can_use_api_keys = false,
							can_use_auth_tokens = false,
							can_use_console_password = false,
							can_use_customer_secret_keys = false,
							can_use_smtp_credentials = false,
						}`,
						map[string]string{"description": "automated test user"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "can_use_api_keys", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "can_use_auth_tokens", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "can_use_console_password", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "can_use_customer_secret_keys", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "can_use_smtp_credentials", "false"),
				),
			},
			// verify update with capabilities all as true
			{
				Config: s.Config +
					tokenFn(
						`resource "oci_identity_user" "user" {
							name = "{{.token}}"
							description = "{{.description}}"
							compartment_id = "${var.tenancy_ocid}"
						}
						resource "oci_identity_user_capabilities_management" "t" {
							user_id = "${oci_identity_user.user.id}"
							can_use_api_keys = true,
							can_use_auth_tokens = true,
							can_use_console_password = true,
							can_use_customer_secret_keys = true,
							can_use_smtp_credentials = true,
						}`,
						map[string]string{"description": "automated test user"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "can_use_api_keys", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "can_use_auth_tokens", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "can_use_console_password", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "can_use_customer_secret_keys", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "can_use_smtp_credentials", "true"),
				),
			},
		},
	})
}

func TestResourceIdentityUserCapabilitiesManagementTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserCapabilitiesManagementTestSuite))
}
