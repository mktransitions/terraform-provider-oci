// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	//"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_license_manager "github.com/oracle/oci-go-sdk/v65/licensemanager"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ProductLicenseRequiredOnlyResource = ProductLicenseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_license_manager_product_license", "test_product_license", acctest.Required, acctest.Create, productLicenseRepresentation)

	ProductLicenseResourceConfig = ProductLicenseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_license_manager_product_license", "test_product_license", acctest.Optional, acctest.Update, productLicenseRepresentation)

	productLicenseSingularDataSourceRepresentation = map[string]interface{}{
		"product_license_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_license_manager_product_license.test_product_license.id}`},
	}

	productLicenseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"is_compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: productLicenseDataSourceFilterRepresentation}}
	productLicenseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_license_manager_product_license.test_product_license.id}`}},
	}

	productLicenseRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `displayName`},
		"is_vendor_oracle": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"license_unit":     acctest.Representation{RepType: acctest.Required, Create: `OCPU`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"images":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: productLicenseImagesRepresentation},
		"vendor_name":      acctest.Representation{RepType: acctest.Required, Create: `vendorName`},
	}
	productLicenseImagesRepresentation = map[string]interface{}{
		"listing_id":      acctest.Representation{RepType: acctest.Required, Create: `101747862`},
		"package_version": acctest.Representation{RepType: acctest.Required, Create: `2019.8.9`, Update: `2019.8.10`},
	}

	ProductLicenseResourceDependencies = ""
)

// issue-routing-tag: license_manager/default
func TestLicenseManagerProductLicenseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLicenseManagerProductLicenseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_license_manager_product_license.test_product_license"
	datasourceName := "data.oci_license_manager_product_licenses.test_product_licenses"
	singularDatasourceName := "data.oci_license_manager_product_license.test_product_license"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ProductLicenseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_license_manager_product_license", "test_product_license", acctest.Optional, acctest.Create, productLicenseRepresentation), "licensemanager", "productLicense", t)

	acctest.ResourceTest(t, testAccCheckLicenseManagerProductLicenseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ProductLicenseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_license_manager_product_license", "test_product_license", acctest.Required, acctest.Create, productLicenseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "is_vendor_oracle", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_unit", "OCPU"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ProductLicenseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ProductLicenseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_license_manager_product_license", "test_product_license", acctest.Optional, acctest.Create, productLicenseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "images.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "images.0.listing_id"),
				resource.TestCheckResourceAttr(resourceName, "images.0.package_version", "2019.8.9"),
				resource.TestCheckResourceAttr(resourceName, "is_vendor_oracle", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_unit", "OCPU"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "vendorName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ProductLicenseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_license_manager_product_license", "test_product_license", acctest.Optional, acctest.Update, productLicenseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "images.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "images.0.listing_id"),
				resource.TestCheckResourceAttr(resourceName, "images.0.package_version", "2019.8.10"),
				resource.TestCheckResourceAttr(resourceName, "is_vendor_oracle", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_unit", "OCPU"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "vendorName"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_license_manager_product_licenses", "test_product_licenses", acctest.Optional, acctest.Update, productLicenseDataSourceRepresentation) +
				compartmentIdVariableStr + ProductLicenseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_license_manager_product_license", "test_product_license", acctest.Optional, acctest.Update, productLicenseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "is_compartment_id_in_subtree", "false"),

				resource.TestCheckResourceAttr(datasourceName, "product_license_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "product_license_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_license_manager_product_license", "test_product_license", acctest.Required, acctest.Create, productLicenseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ProductLicenseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "product_license_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "active_license_record_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "images.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "images.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "images.0.listing_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "images.0.package_version", "2019.8.10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "images.0.publisher"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_over_subscribed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_unlimited"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_vendor_oracle", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_unit", "OCPU"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_active_license_unit_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_license_record_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_license_units_consumed"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vendor_name", "vendorName"),
			),
		},
		// verify resource import
		{
			Config:                  config + ProductLicenseRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLicenseManagerProductLicenseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LicenseManagerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_license_manager_product_license" {
			noResourceFound = false
			request := oci_license_manager.GetProductLicenseRequest{}

			tmp := rs.Primary.ID
			request.ProductLicenseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "license_manager")

			response, err := client.GetProductLicense(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_license_manager.LifeCycleStateDeleted): true,
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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("LicenseManagerProductLicense") {
		resource.AddTestSweepers("LicenseManagerProductLicense", &resource.Sweeper{
			Name:         "LicenseManagerProductLicense",
			Dependencies: acctest.DependencyGraph["productLicense"],
			F:            sweepLicenseManagerProductLicenseResource,
		})
	}
}

func sweepLicenseManagerProductLicenseResource(compartment string) error {
	licenseManagerClient := acctest.GetTestClients(&schema.ResourceData{}).LicenseManagerClient()
	productLicenseIds, err := getProductLicenseIds(compartment)
	if err != nil {
		return err
	}
	for _, productLicenseId := range productLicenseIds {
		if ok := acctest.SweeperDefaultResourceId[productLicenseId]; !ok {
			deleteProductLicenseRequest := oci_license_manager.DeleteProductLicenseRequest{}

			deleteProductLicenseRequest.ProductLicenseId = &productLicenseId

			deleteProductLicenseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "license_manager")
			_, error := licenseManagerClient.DeleteProductLicense(context.Background(), deleteProductLicenseRequest)
			if error != nil {
				fmt.Printf("Error deleting ProductLicense %s %s, It is possible that the resource is already deleted. Please verify manually \n", productLicenseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &productLicenseId, productLicenseSweepWaitCondition, time.Duration(3*time.Minute),
				productLicenseSweepResponseFetchOperation, "license_manager", true)
		}
	}
	return nil
}

func getProductLicenseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ProductLicenseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	licenseManagerClient := acctest.GetTestClients(&schema.ResourceData{}).LicenseManagerClient()

	listProductLicensesRequest := oci_license_manager.ListProductLicensesRequest{}
	listProductLicensesRequest.CompartmentId = &compartmentId
	listProductLicensesResponse, err := licenseManagerClient.ListProductLicenses(context.Background(), listProductLicensesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ProductLicense list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, productLicense := range listProductLicensesResponse.Items {
		id := *productLicense.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ProductLicenseId", id)
	}
	return resourceIds, nil
}

func productLicenseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if productLicenseResponse, ok := response.Response.(oci_license_manager.GetProductLicenseResponse); ok {
		return productLicenseResponse.LifecycleState != oci_license_manager.LifeCycleStateDeleted
	}
	return false
}

func productLicenseSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.LicenseManagerClient().GetProductLicense(context.Background(), oci_license_manager.GetProductLicenseRequest{
		ProductLicenseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}