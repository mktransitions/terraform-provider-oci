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

type ResourceIdentityAPIKeysTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListAPIKeyResponses
}

func (s *ResourceIdentityAPIKeysTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_identity_api_keys" "t" {
      user_id = "user_id"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_identity_api_keys.t"

	b1 := baremetal.APIKey{
		Fingerprint: "fingerprint",
		KeyID:       "id1",
		KeyValue:    "key_value",
		State:       baremetal.ResourceAvailable,
		TimeCreated: time.Now(),
		UserID:      "user_id",
	}

	b2 := b1
	b2.KeyID = "id2"

	s.List = &baremetal.ListAPIKeyResponses{
		Keys: []baremetal.APIKey{b1, b2},
	}
}

func (s *ResourceIdentityAPIKeysTestSuite) TestReadAPIKeys() {
	s.Client.On("ListAPIKeys", "user_id").Return(s.List, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "user_id", "user_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "api_keys.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "api_keys.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "api_keys.#", "2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListAPIKeys", "user_id")
}

func TestResourceIdentityAPIKeysTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityAPIKeysTestSuite))
}
