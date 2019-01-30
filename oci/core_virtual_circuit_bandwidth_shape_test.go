// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	virtualCircuitBandwidthShapeDataSourceRepresentation = map[string]interface{}{
		"provider_service_id": Representation{repType: Required, create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}`},
	}

	// @CODEGEN 07/2018: ProviderService is actually FastConnectProviderService
	VirtualCircuitBandwidthShapeResourceDependencies = FastConnectProviderServiceResourceConfig
)

func TestCoreVirtualCircuitBandwidthShapeResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_virtual_circuit_bandwidth_shapes.test_virtual_circuit_bandwidth_shapes"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `

data "oci_core_fast_connect_provider_services" "test_fast_connect_provider_services" {
	#Required
	compartment_id = "${var.compartment_id}"

}` +
					generateDataSourceFromRepresentationMap("oci_core_virtual_circuit_bandwidth_shapes", "test_virtual_circuit_bandwidth_shapes", Required, Create, virtualCircuitBandwidthShapeDataSourceRepresentation) +
					compartmentIdVariableStr + VirtualCircuitBandwidthShapeResourceDependencies,
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_bandwidth_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_bandwidth_shapes.0.name"),
				),
			},
		},
	})
}
