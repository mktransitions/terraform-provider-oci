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
	smtpCredentialDataSourceRepresentation = map[string]interface{}{
		"user_id": Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
		"filter":  RepresentationGroup{Required, smtpCredentialDataSourceFilterRepresentation}}
	smtpCredentialDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_smtp_credential.test_smtp_credential.id}`}},
	}

	smtpCredentialRepresentation = map[string]interface{}{
		"description": Representation{repType: Required, create: `description`, update: `description2`},
		"user_id":     Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
	}

	SmtpCredentialResourceDependencies = generateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)
)

func TestIdentitySmtpCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentitySmtpCredentialResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_smtp_credential.test_smtp_credential"
	datasourceName := "data.oci_identity_smtp_credentials.test_smtp_credentials"

	var resId, resId2 string
	var compositeId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentitySmtpCredentialDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + SmtpCredentialResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_smtp_credential", "test_smtp_credential", Required, Create, smtpCredentialRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						userId, _ := fromInstanceState(s, resourceName, "user_id")
						compositeId = "users/" + userId + "/smtpCredentials/" + resId
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
				Config: config + compartmentIdVariableStr + SmtpCredentialResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_smtp_credential", "test_smtp_credential", Optional, Update, smtpCredentialRepresentation),
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
					generateDataSourceFromRepresentationMap("oci_identity_smtp_credentials", "test_smtp_credentials", Optional, Update, smtpCredentialDataSourceRepresentation) +
					compartmentIdVariableStr + SmtpCredentialResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_smtp_credential", "test_smtp_credential", Optional, Update, smtpCredentialRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

					resource.TestCheckResourceAttr(datasourceName, "smtp_credentials.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "smtp_credentials.0.description", "description2"),
					resource.TestCheckResourceAttrSet(datasourceName, "smtp_credentials.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "smtp_credentials.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "smtp_credentials.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "smtp_credentials.0.user_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "smtp_credentials.0.username"),
				),
			},
			// verify resource import
			//{
			//	Config:            config,
			//	ImportState:       true,
			//	ImportStateVerify: true,
			//	ImportStateVerifyIgnore: []string{
			//		"password",
			//	},
			//	ResourceName: resourceName,
			//},
		},
	})
}

func testAccCheckIdentitySmtpCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_smtp_credential" {
			noResourceFound = false
			request := oci_identity.ListSmtpCredentialsRequest{}

			if value, ok := rs.Primary.Attributes["user_id"]; ok {
				request.UserId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")
			response, err := client.ListSmtpCredentials(context.Background(), request)

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
