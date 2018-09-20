// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ServiceResourceConfig = ServiceResourceDependencies + `

`
	ServicePropertyVariables = `

`
	ServiceResourceDependencies = ""
)

func TestCoreServiceResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_services.test_services"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `

data "oci_core_services" "test_services" {
}
                ` + compartmentIdVariableStr + ServiceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceName, "services.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "services.0.cidr_block"),
					resource.TestCheckResourceAttrSet(datasourceName, "services.0.description"),
					resource.TestCheckResourceAttrSet(datasourceName, "services.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "services.0.name"),
				),
			},
		},
	})
}
