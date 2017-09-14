// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func TestLoadBalancerBackendsDatasource(t *testing.T) {
	client := GetTestProvider()
	providers := map[string]terraform.ResourceProvider{
		"oci": Provider(func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		}),
	}
	resourceName := "data.oci_load_balancer_backends.t"
	config := `
data "oci_load_balancer_backends" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
  backendset_name  = "stub_backendset_name"
}
`
	config += testProviderConfig()

	loadbalancerID := "ocid1.loadbalancer.stub_id"

	resource.UnitTest(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "load_balancer_id", loadbalancerID),
					resource.TestCheckResourceAttr(resourceName, "backends.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "backends.0.name", "123.123.123.123"),
					resource.TestCheckResourceAttr(resourceName, "backends.1.name", "122.122.122.122"),
				),
			},
		},
	})
}
