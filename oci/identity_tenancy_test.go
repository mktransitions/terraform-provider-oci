// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	tenancySingularDataSourceRepresentation = map[string]interface{}{
		"tenancy_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
	}
)

func TestIdentityTenancyResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	singularDatasourceName := "data.oci_identity_tenancy.test_tenancy"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_tenancy", "test_tenancy", Required, Create, tenancySingularDataSourceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "home_region_key"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				),
			},
		},
	})
}
