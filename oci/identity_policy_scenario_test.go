// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/stretchr/testify/suite"
)

type ResourceIdentityPolicyTestSuite struct {
	suite.Suite
	Providers      map[string]terraform.ResourceProvider
	Config         string
	ResourceName   string
	DataSourceName string
	Token          string
	TokenFn        func(string, map[string]string) string
}

func (s *ResourceIdentityPolicyTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig() + s.TokenFn(`
	resource "oci_identity_group" "t" {
		name = "{{.token}}"
		description = "automated test group"
		compartment_id = "${var.tenancy_ocid}"
	}`, nil)
	s.ResourceName = "oci_identity_policy.p"
	s.DataSourceName = "data.oci_identity_policies.p"
}

func (s *ResourceIdentityPolicyTestSuite) TestAccResourceIdentityPolicy_basic() {
	var policyHash string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: s.Config + s.TokenFn(`
				resource "oci_identity_policy" "p" {
					compartment_id = "${var.tenancy_ocid}"
					name = "p1-{{.token}}"
					description = "automated test policy"
					version_date = "2018-04-17"
					statements = ["Allow group ${oci_identity_group.t.name} to read instances in tenancy"]
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ETag"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "lastUpdateETag"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "policyHash"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", "p1-"+s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test policy"),
					resource.TestCheckResourceAttr(s.ResourceName, "statements.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.PolicyLifecycleStateActive)),
					resource.TestCheckResourceAttr(s.ResourceName, "version_date", "2018-04-17"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					func(s *terraform.State) (err error) {
						policyHash, err = fromInstanceState(s, "oci_identity_policy.p", "policyHash")
						return err
					},
				),
			},
			// verify update
			{
				Config: s.Config + s.TokenFn(`
				resource "oci_identity_policy" "p" {
					compartment_id = "${var.tenancy_ocid}"
					name = "{{.token}}"
					description = "automated test policy (updated)"
					version_date = "2018-04-18"
					statements = [
						"Allow group ${oci_identity_group.t.name} to inspect instances in tenancy",
						"Allow group ${oci_identity_group.t.name} to read instances in tenancy"
					]
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test policy (updated)"),
					resource.TestCheckResourceAttr(s.ResourceName, "version_date", "2018-04-18"),
					resource.TestCheckResourceAttr(s.ResourceName, "statements.#", "2"),
					func(s *terraform.State) (err error) {
						newHash, err := fromInstanceState(s, "oci_identity_policy.p", "policyHash")
						if policyHash == newHash {
							return fmt.Errorf("Expected new hash, got same")
						}
						return err
					},
				),
			},
			// verify data source, + filtering against array of items
			{
				Config: s.Config + s.TokenFn(`
				resource "oci_identity_policy" "p" {
					compartment_id = "${var.tenancy_ocid}"
					name = "{{.token}}"
					description = "automated test policy (updated)"
					version_date = "2018-04-18"
					statements = [
						"Allow group ${oci_identity_group.t.name} to inspect instances in tenancy",
						"Allow group ${oci_identity_group.t.name} to read instances in tenancy"
					]
				}
				data "oci_identity_policies" "p" {
					compartment_id = "${var.tenancy_ocid}"
					filter {
						name   = "statements"
						values = ["Allow group ${oci_identity_group.t.name} to inspect instances in tenancy"]
					}
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.DataSourceName, "policies.#", "1"),
					resource.TestCheckResourceAttrSet(s.DataSourceName, "policies.0.id"),
					resource.TestCheckResourceAttr(s.DataSourceName, "policies.0.name", s.Token),
					resource.TestCheckResourceAttr(s.DataSourceName, "policies.0.description", "automated test policy (updated)"),
					resource.TestCheckResourceAttr(s.DataSourceName, "policies.0.state", string(identity.PolicyLifecycleStateActive)),
					// TODO: This field is not being returned by the service call but is still showing up in the datasource
					// resource.TestCheckNoResourceAttr(s.ResourceName, "policies.0.inactive_state"),
					resource.TestCheckResourceAttr(s.DataSourceName, "policies.0.statements.#", "2"),
					resource.TestCheckResourceAttrSet(s.DataSourceName, "policies.0.time_created"),
					resource.TestCheckResourceAttr(s.DataSourceName, "policies.0.version_date", "2018-04-18"),
				),
			},
		},
	},
	)
}

func (s *ResourceIdentityPolicyTestSuite) TestAccResourceIdentityPolicy_emptyStatement() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: s.Config + s.TokenFn(`
				resource "oci_identity_policy" "p" {
					compartment_id = "${var.tenancy_ocid}"
					name = "p1-{{.token}}"
					description = "automated test policy"
					version_date = "2018-04-17"
					statements = [
"Allow group ${oci_identity_group.t.name} to inspect instances in tenancy",
"",
"Allow group ${oci_identity_group.t.name} to inspect instances in tenancy"]
				}`, nil),
				ExpectError: regexp.MustCompile("Service error:InvalidParameter"),
			},
		},
	},
	)
}

func (s *ResourceIdentityPolicyTestSuite) TestAccResourceIdentityPolicy_formattingDiff() {
	var lastUpdateETag, policyHash string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// create policy with bad formatting
			{
				Config: s.Config + s.TokenFn(`
				resource "oci_identity_policy" "p" {
					compartment_id = "${var.tenancy_ocid}"
					name = "{{.token}}"
					description = "automated test policy"
					statements = ["Allow group ${oci_identity_group.t.name} to read instances in >> tenancy"]
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					// policy statements may or may not have invalid characters stripped (">>" above), accommodate this uncertainty as specifically as possible
					resource.TestMatchResourceAttr(s.ResourceName, "statements.0",
						regexp.MustCompile(`Allow group `+s.Token+` to read instances in (>> )?tenancy`)),
					func(s *terraform.State) (err error) {
						if policyHash, err = fromInstanceState(s, "oci_identity_policy.p", "policyHash"); err == nil {
							lastUpdateETag, err = fromInstanceState(s, "oci_identity_policy.p", "lastUpdateETag")
						}
						return err
					},
				),
			},
			// verify update does not change the hash and ETag value
			{
				Config: s.Config + s.TokenFn(`
				resource "oci_identity_policy" "p" {
					compartment_id = "${var.tenancy_ocid}"
					name = "{{.token}}"
					description = "automated test policy"
					statements = ["Allow group ${oci_identity_group.t.name} to read instances in >> tenancy"]
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					func(s *terraform.State) (err error) {
						resource.TestCheckResourceAttr("oci_identity_policy.p", "policyHash", policyHash)
						resource.TestCheckResourceAttr("oci_identity_policy.p", "lastUpdateETag", lastUpdateETag)
						return err
					},
				),
			},
		},
	},
	)
}

func TestResourceIdentityPolicyTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityPolicyTestSuite))
}
