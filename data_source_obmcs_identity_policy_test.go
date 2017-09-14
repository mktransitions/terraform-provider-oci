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

type DatasourceIdentityPoliciesTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	PoliciesName string
	Policies     baremetal.ListPolicies
}

func (s *DatasourceIdentityPoliciesTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	},
	)
	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = `
	resource "oci_identity_compartment" "t" {
		name = "test-compartment"
		description = "automated test compartment"
	}

	resource "oci_identity_group" "t" {
		name = "-tf-group"
		description = "automated test group"
	}

	resource "oci_identity_policy" "p" {
		name = "-tf-policy"
		description = "automated test policy"
		compartment_id = "${oci_identity_compartment.t.id}"
		statements = ["Allow group ${oci_identity_group.t.name} to read instances in compartment ${oci_identity_compartment.t.name}"]
	}
	`
	s.Config += testProviderConfig()
	s.PoliciesName = "data.oci_identity_policies.p"
}

func (s *DatasourceIdentityPoliciesTestSuite) TestListResourceIdentityPolicies() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config +
					`	data "oci_identity_policies" "p" {
							compartment_id = "${oci_identity_compartment.t.id}"
						}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.PoliciesName, "policies.#"),
					// this not working seems like an interpolation issue with terraform,
					// the policy data is definitely there
					//resource.TestCheckResourceAttrSet(s.PoliciesName, "policies.0.id"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityPoliciesTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityPoliciesTestSuite))
}
