// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	dbSystemShapeDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
	}

	DbSystemShapeResourceConfig = AvailabilityDomainConfig
)

func TestDatabaseDbSystemShapeResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_system_shapes.test_db_system_shapes"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_db_system_shapes", "test_db_system_shapes", Required, Create, dbSystemShapeDataSourceRepresentation) +
					compartmentIdVariableStr + DbSystemShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.available_core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.core_count_increment"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.maximum_node_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.minimum_core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.minimum_node_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.name"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.shape"),
				),
			},
		},
	})
}
