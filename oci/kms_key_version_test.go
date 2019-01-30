// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	KeyVersionResourceConfig = KeyVersionResourceDependencies +
		generateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Required, Create, keyVersionRepresentation)

	keyVersionSingularDataSourceRepresentation = map[string]interface{}{
		"key_id":              Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"key_version_id":      Representation{repType: Required, create: `${oci_kms_key_version.test_key_version.key_version_id}`},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
	}

	keyVersionDataSourceRepresentation = map[string]interface{}{
		"key_id":              Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"filter":              RepresentationGroup{Required, keyVersionDataSourceFilterRepresentation}}
	keyVersionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `key_version_id`},
		"values": Representation{repType: Required, create: []string{`${oci_kms_key_version.test_key_version.key_version_id}`}},
	}

	keyVersionRepresentation = map[string]interface{}{
		"key_id":              Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
	}

	KeyVersionResourceDependencies = KeyResourceDependencyConfig
)

func TestKmsKeyVersionResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_kms_key_version.test_key_version"
	datasourceName := "data.oci_kms_key_versions.test_key_versions"
	singularDatasourceName := "data.oci_kms_key_version.test_key_version"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + KeyVersionResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Required, Create, keyVersionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_key_versions", "test_key_versions", Optional, Update, keyVersionDataSourceRepresentation) +
					compartmentIdVariableStr + KeyVersionResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Optional, Update, keyVersionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "key_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_endpoint"),

					resource.TestCheckResourceAttr(datasourceName, "key_versions.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.key_version_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.key_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.vault_id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Required, Create, keyVersionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + KeyVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_version_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "management_endpoint"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + KeyVersionResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: keyVersionImportId,
				ImportStateVerifyIgnore: []string{
					"state",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func keyVersionImportId(state *terraform.State) (string, error) {
	for _, rs := range state.RootModule().Resources {
		if rs.Type == "oci_kms_key_version" {
			return fmt.Sprintf("managementEndpoint/%s/%s", rs.Primary.Attributes["management_endpoint"], rs.Primary.ID), nil
		}
	}

	return "", fmt.Errorf("unable to create import id as no resource of type oci_kms_key_version in state")
}
