// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	workRequestLogEntryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  Representation{repType: Required, create: `${var.compartment_id}`},
		"work_request_id": Representation{repType: Required, create: `${lookup(data.oci_containerengine_work_requests.test_work_requests.work_requests[0], "id")}`},
	}

	WorkRequestLogEntryResourceConfig = WorkRequestResourceConfig +
		generateDataSourceFromRepresentationMap("oci_containerengine_work_requests", "test_work_requests", Optional, Create, workRequestDataSourceRepresentation)
)

func TestContainerengineWorkRequestLogEntryResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_containerengine_work_request_log_entries.test_work_request_log_entries"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_containerengine_work_request_log_entries", "test_work_request_log_entries", Required, Create, workRequestLogEntryDataSourceRepresentation) +
					compartmentIdVariableStr + WorkRequestLogEntryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "work_request_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "work_request_log_entries.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "work_request_log_entries.0.message"),
					resource.TestCheckResourceAttrSet(datasourceName, "work_request_log_entries.0.timestamp"),
				),
			},
		},
	})
}
