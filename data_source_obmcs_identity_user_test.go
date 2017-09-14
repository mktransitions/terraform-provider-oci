// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityUsersTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListUsers
}

func (s *DatasourceIdentityUsersTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}
	s.Config = `
		resource "oci_identity_user" "t" {
			name = "-tf-user"
			description = "automated test user"
		}
	`
	s.Config += testProviderConfig()
	s.ResourceName = "data.oci_identity_users.t"
}

func (s *DatasourceIdentityUsersTestSuite) TestReadUsers() {

	resource.UnitTest(s.T(), resource.TestCase{
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
				data "oci_identity_users" "t" {
					compartment_id = "${var.compartment_id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "users.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "users.#"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityUsersTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityUsersTestSuite))
}
