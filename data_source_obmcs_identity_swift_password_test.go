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

type ResourceIdentitySwiftPasswordsTestSuite struct {
	suite.Suite
	Client        *baremetal.Client
	Provider      terraform.ResourceProvider
	Providers     map[string]terraform.ResourceProvider
	TimeCreated   time.Time
	Config        string
	PasswordsName string
	PasswordList  baremetal.ListSwiftPasswords
}

func (s *ResourceIdentitySwiftPasswordsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	},
	)
	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = `
		resource "baremetal_identity_user" "t" {
			name = "name1"
			description = "desc!"
		}
		resource "baremetal_identity_swift_password" "t" {
			user_id = "${baremetal_identity_user.t.id}"
			description = "desc"
		}
	`
	s.Config += testProviderConfig()
	s.PasswordsName = "data.baremetal_identity_swift_passwords.p"

}

func (s *ResourceIdentitySwiftPasswordsTestSuite) TestListResourceIdentitySwiftPasswords() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				 data "baremetal_identity_swift_passwords" "p" {
				    user_id = "${baremetal_identity_user.t.id}"
				  }`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.PasswordsName, "passwords.0.id"),
					resource.TestCheckResourceAttr(s.PasswordsName, "passwords.#", "1"),
				),
			},
		},
	},
	)
}

func TestResourceIdentitySwiftPasswordsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentitySwiftPasswordsTestSuite))
}
