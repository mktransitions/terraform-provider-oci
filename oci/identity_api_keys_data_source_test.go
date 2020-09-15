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

type DatasourceIdentityAPIKeysTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         identity.ListApiKeysResponse
}

func (s *DatasourceIdentityAPIKeysTestSuite) SetupTest() {
	_, tokenFn := tokenizeWithHttpReplay("api_data_source")
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig() + tokenFn(`
	resource "oci_identity_user" "t" {
		name = "{{.userName}}"
		description = "automated test user"
		compartment_id = "${var.tenancy_ocid}"
	}
	
	resource "oci_identity_api_key" "t" {
		user_id = "${oci_identity_user.t.id}"
		key_value = <<EOF
`+apiKey+`
EOF
	}
	
	resource "oci_identity_api_key" "u" {
		user_id = "${oci_identity_user.t.id}"
		key_value = <<EOF
`+apiKey2+`
EOF
	}`, map[string]string{"userName": "user_" + timestamp()})
	s.ResourceName = "data.oci_identity_api_keys.t"
}

func (s *DatasourceIdentityAPIKeysTestSuite) TestAccDatasourceIdentityAPIKeys_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config: s.Config + `
				data "oci_identity_api_keys" "t" {
					user_id = "${oci_identity_user.t.id}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "api_keys.#", "2"),
				),
			},
			// Client-side filtering.
			{
				Config: s.Config + `
				data "oci_identity_api_keys" "t" {
					user_id = "${oci_identity_user.t.id}"
					filter {
						name = "id"
						values = ["${oci_identity_api_key.t.id}"]
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "api_keys.#", "1"),
					TestCheckResourceAttributesEqual(s.ResourceName, "api_keys.0.id", "oci_identity_api_key.t", "id"),
					TestCheckResourceAttributesEqual(s.ResourceName, "api_keys.0.fingerprint", "oci_identity_api_key.t", "fingerprint"),
					resource.TestCheckResourceAttr(s.ResourceName, "api_keys.0.key_value", apiKey),
					TestCheckResourceAttributesEqual(s.ResourceName, "api_keys.0.time_created", "oci_identity_api_key.t", "time_created"),
					// TODO: This field is not being returned by the service call but is showing up in the datasource
					//resource.TestCheckNoResourceAttr(s.ResourceName, "api_keys.0.inactive_status"),
					resource.TestCheckResourceAttr(s.ResourceName, "api_keys.0.state", string(identity.ApiKeyLifecycleStateActive)),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityAPIKeysTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceIdentityAPIKeysTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceIdentityAPIKeysTestSuite))
}
