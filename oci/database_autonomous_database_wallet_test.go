// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	autonomousDatabaseWalletSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"password":               Representation{repType: Required, create: `BEstrO0ng_#11`},
	}

	AutonomousDatabaseWalletResourceConfig = AutonomousDatabaseResourceConfig
)

func TestDatabaseAutonomousDatabaseWalletResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_autonomous_database_wallet.test_autonomous_database_wallet"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Required, Create, autonomousDatabaseWalletSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseWalletResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				),
			},
		},
	})
}
