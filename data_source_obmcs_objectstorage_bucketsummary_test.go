// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ObjectstorageBucketSummaryTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	TimeCreated  time.Time
}

func (s *ObjectstorageBucketSummaryTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}
	s.Config = `
	resource "oci_objectstorage_bucket" "t" {
		compartment_id = "${var.compartment_id}"
		name = "bucketID"
		namespace = "${var.namespace}"
		metadata = {
			"foo" = "bar"
		}
	}
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.oci_objectstorage_bucket_summaries.t"
	s.TimeCreated = time.Now()
}

func (s *ObjectstorageBucketSummaryTestSuite) TestReadBucketSummaries() {
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
					data "oci_objectstorage_bucket_summaries" "t" {
						compartment_id = "${var.compartment_id}"
						namespace = "${var.namespace}"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "bucket_summaries.0.name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "bucket_summaries.#"),
				),
			},
		},
	},
	)
}

func TestObjectstorageBucketSummaryTestSuite(t *testing.T) {
	suite.Run(t, new(ObjectstorageBucketSummaryTestSuite))
}
