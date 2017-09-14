// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreInstanceCredentialTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreInstanceCredentialTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}
	s.Config = instanceConfig + `
    data "oci_core_instance_credentials" "s" {
      instance_id = "${oci_core_instance.t.id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.oci_core_instance_credentials.s"

}

func (s *ResourceCoreInstanceCredentialTestSuite) TestResourceReadCoreInstanceCredential() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "username"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "password"),
				),
			},
		},
	},
	)

}

func TestResourceCoreInstanceCredentialTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreInstanceCredentialTestSuite))
}
