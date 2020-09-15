// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/v25/identity"
	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUIPasswordTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentityUIPasswordTestSuite) SetupTest() {
	_, tokenFn := tokenizeWithHttpReplay("ui_pass_resource")
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig() + tokenFn(`
	resource "oci_identity_user" "t" {
		name = "-tf-user"
		description = "tf test user"
		compartment_id = "${var.tenancy_ocid}"
	}`, nil)

	s.ResourceName = "oci_identity_ui_password.t"
}

func (s *ResourceIdentityUIPasswordTestSuite) TestAccIdentityUIPassword_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: s.Config + `
				resource "oci_identity_ui_password" "t" {
					user_id = "${oci_identity_user.t.id}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "password"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.UiPasswordLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_status"),
				),
			},
		},
	})
}

func TestResourceIdentityUIPasswordTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceIdentityUIPasswordTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceIdentityUIPasswordTestSuite))
}
