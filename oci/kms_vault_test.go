// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

var (
	VaultResourceConfig = VaultResourceDependencies +
		generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Required, Create, vaultRepresentation)

	vaultSingularDataSourceRepresentation = map[string]interface{}{
		"vault_id": Representation{repType: Required, create: `{}`},
	}

	vaultDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"filter":         RepresentationGroup{Required, vaultDataSourceFilterRepresentation}}
	vaultDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_kms_vault.test_vault.id}`}},
	}

	vaultRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `Vault 1`, update: `displayName2`},
		"vault_type":     Representation{repType: Required, create: `VIRTUAL_PRIVATE`},
	}

	VaultResourceDependencies = ""
)

func TestKmsVaultResource_basic(t *testing.T) {
	t.Skip("Skip this test till KMS provides a better way of testing this.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_vault.test_vault"
	datasourceName := "data.oci_kms_vaults.test_vaults"
	singularDatasourceName := "data.oci_kms_vault.test_vault"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckKMSVaultDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Required, Create, vaultRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Vault 1"),
					resource.TestCheckResourceAttr(resourceName, "vault_type", "VIRTUAL_PRIVATE"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + VaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Update, vaultRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "vault_type", "VIRTUAL_PRIVATE"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_vaults", "test_vaults", Optional, Update, vaultDataSourceRepresentation) +
					compartmentIdVariableStr + VaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Update, vaultRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "vaults.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "vaults.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.crypto_endpoint"),
					resource.TestCheckResourceAttr(datasourceName, "vaults.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.management_endpoint"),
					resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "vaults.0.vault_type", "VIRTUAL_PRIVATE"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_vault", "test_vault", Required, Create, vaultSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VaultResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "crypto_endpoint", "cryptoEndpoint"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "id", "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "management_endpoint", "managementEndpoint"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "time_created", "timeCreated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "time_of_deletion", "timeOfDeletion"),
					resource.TestCheckResourceAttr(singularDatasourceName, "vault_type", "VIRTUAL_PRIVATE"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + VaultResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckKMSVaultDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).kmsVaultClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_kms_vault" {
			noResourceFound = false
			request := oci_kms.GetVaultRequest{}

			tmp := rs.Primary.ID
			request.VaultId = &tmp

			response, err := client.GetVault(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_kms.VaultLifecycleStatePendingDeletion): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
