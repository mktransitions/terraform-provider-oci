// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/v25/common"
	oci_identity "github.com/oracle/oci-go-sdk/v25/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	authTokenDataSourceRepresentation = map[string]interface{}{
		"user_id": Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
		"filter":  RepresentationGroup{Required, authTokenDataSourceFilterRepresentation}}
	authTokenDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_auth_token.test_auth_token.id}`}},
	}

	authTokenRepresentation = map[string]interface{}{
		"description": Representation{repType: Required, create: `description`, update: `description2`},
		"user_id":     Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
	}

	AuthTokenResourceDependencies = generateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)
)

func TestIdentityAuthTokenResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityAuthTokenResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_auth_token.test_auth_token"
	datasourceName := "data.oci_identity_auth_tokens.test_auth_tokens"

	var resId, resId2 string
	var compositeId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityAuthTokenDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AuthTokenResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_auth_token", "test_auth_token", Required, Create, authTokenRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						userId, _ := fromInstanceState(s, resourceName, "user_id")
						compositeId = "users/" + userId + "/authTokens/" + resId
						log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AuthTokenResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_auth_token", "test_auth_token", Optional, Update, authTokenRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),

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
					generateDataSourceFromRepresentationMap("oci_identity_auth_tokens", "test_auth_tokens", Optional, Update, authTokenDataSourceRepresentation) +
					compartmentIdVariableStr + AuthTokenResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_auth_token", "test_auth_token", Optional, Update, authTokenRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

					resource.TestCheckResourceAttr(datasourceName, "tokens.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tokens.0.description", "description2"),
					resource.TestCheckResourceAttrSet(datasourceName, "tokens.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "tokens.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "tokens.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "tokens.0.user_id"),
				),
			},
			// verify resource import
			//{
			//	Config:                  config,
			//	ImportState:             true,
			//	ImportStateVerify:       true,
			//	ImportStateVerifyIgnore: []string{},
			//	ResourceName:            resourceName,
			//},
		},
	})
}

func testAccCheckIdentityAuthTokenDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_auth_token" {
			noResourceFound = false
			request := oci_identity.ListAuthTokensRequest{}

			if value, ok := rs.Primary.Attributes["user_id"]; ok {
				request.UserId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")
			response, err := client.ListAuthTokens(context.Background(), request)

			if err == nil {
				id := rs.Primary.Attributes["id"]
				for _, item := range response.Items {
					if *item.Id == id {
						return fmt.Errorf("item still exists")
					}
				}
				// no error and item not found, that means item is deleted. continue checking next one
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
