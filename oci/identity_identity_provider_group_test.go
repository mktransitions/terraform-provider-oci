// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	identityProviderGroupDataSourceRepresentation = map[string]interface{}{
		"identity_provider_id": Representation{repType: Required, create: `${oci_identity_identity_provider.test_identity_provider.id}`},
	}

	IdentityProviderGroupResourceConfig = IdentityProviderRequiredOnlyResource
)

func TestIdentityIdentityProviderGroupResource_basic(t *testing.T) {
	metadataFile := getEnvSettingWithBlankDefault("identity_provider_metadata_file")
	if metadataFile == "" {
		t.Skip("Skipping generated test for now as it has a dependency on federation metadata file")
	}

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_identity_provider_groups.test_identity_provider_groups"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_identity_provider_groups", "test_identity_provider_groups", Required, Create, identityProviderGroupDataSourceRepresentation) +
					compartmentIdVariableStr + IdentityProviderGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "identity_provider_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "identity_provider_groups.#"),
				),
			},
		},
	})
}
