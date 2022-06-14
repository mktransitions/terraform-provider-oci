// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-oci/httpreplay"
	"terraform-provider-oci/internal/acctest"

	"terraform-provider-oci/internal/utils"
)

var (
	managedDatabasesAsmPropertySingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	managedDatabasesAsmPropertyDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	ManagedDatabasesAsmPropertyResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, managedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabasesAsmPropertyResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabasesAsmPropertyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_databases_asm_properties.test_managed_databases_asm_properties"
	singularDatasourceName := "data.oci_database_management_managed_databases_asm_property.test_managed_databases_asm_property"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_asm_properties", "test_managed_databases_asm_properties", acctest.Required, acctest.Create, managedDatabasesAsmPropertyDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabasesAsmPropertyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "asm_property_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "asm_property_collection.0.items.#", "2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_asm_property", "test_managed_databases_asm_property", acctest.Required, acctest.Create, managedDatabasesAsmPropertySingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabasesAsmPropertyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "2"),
			),
		},
	})
}
