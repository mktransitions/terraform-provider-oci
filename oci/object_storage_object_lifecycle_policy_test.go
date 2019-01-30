// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

var (
	ObjectLifecyclePolicyRequiredOnlyResource = ObjectLifecyclePolicyResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Required, Create, objectLifecyclePolicyRepresentation)

	ObjectLifecyclePolicyResourceConfig = ObjectLifecyclePolicyResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Update, objectLifecyclePolicyRepresentation)

	objectLifecyclePolicySingularDataSourceRepresentation = map[string]interface{}{
		"bucket":    Representation{repType: Required, create: `lifecycle-policy-test-1`},
		"namespace": Representation{repType: Required, create: `${data.oci_objectstorage_namespace.t.namespace}`},
	}

	objectLifecyclePolicyRepresentation = map[string]interface{}{
		"bucket":    Representation{repType: Required, create: `lifecycle-policy-test-1`},
		"namespace": Representation{repType: Required, create: `${data.oci_objectstorage_namespace.t.namespace}`},
		"rules":     RepresentationGroup{Optional, objectLifecyclePolicyRulesRepresentation},
	}
	objectLifecyclePolicyRulesRepresentation = map[string]interface{}{
		"action":             Representation{repType: Required, create: `ARCHIVE`, update: `DELETE`},
		"is_enabled":         Representation{repType: Required, create: `false`, update: `true`},
		"name":               Representation{repType: Required, create: `sampleRule`, update: `name2`},
		"time_amount":        Representation{repType: Required, create: `10`, update: `11`},
		"time_unit":          Representation{repType: Required, create: `DAYS`, update: `YEARS`},
		"object_name_filter": RepresentationGroup{Optional, objectLifecyclePolicyRulesObjectNameFilterRepresentation},
	}
	objectLifecyclePolicyRulesObjectNameFilterRepresentation = map[string]interface{}{
		"inclusion_prefixes": Representation{repType: Optional, create: []string{`lifecycle-policy-test-1`, `lifecycle-policy-test-2`}, update: []string{`lifecycle-policy-test-1`, `lifecycle-policy-test-2`, `lifecycle-policy-test-3`}},
	}

	ObjectLifecyclePolicyResourceDependencies = BucketResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create,
			getUpdatedRepresentationCopy("name", Representation{repType: Required, create: `lifecycle-policy-test-1`}, bucketRepresentation))
)

func TestObjectStorageObjectLifecyclePolicyResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	singularDatasourceName := "data.oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectLifecyclePolicyDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Required, Create, objectLifecyclePolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", "lifecycle-policy-test-1"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create, objectLifecyclePolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", "lifecycle-policy-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "rules", map[string]string{
						"action":               "ARCHIVE",
						"is_enabled":           "false",
						"name":                 "sampleRule",
						"object_name_filter.#": "1",
						"object_name_filter.0.inclusion_prefixes.#": "2",
						"time_amount": "10",
						"time_unit":   "DAYS",
					},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Update, objectLifecyclePolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", "lifecycle-policy-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "rules", map[string]string{
						"action":               "DELETE",
						"is_enabled":           "true",
						"name":                 "name2",
						"object_name_filter.#": "1",
						"object_name_filter.0.inclusion_prefixes.#": "3",
						"time_amount": "11",
						"time_unit":   "YEARS",
					},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Required, Create, objectLifecyclePolicySingularDataSourceRepresentation) +
					compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "bucket", "lifecycle-policy-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "rules", map[string]string{},
						[]string{}),

					resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "rules", map[string]string{
						"action":               "DELETE",
						"is_enabled":           "true",
						"name":                 "name2",
						"object_name_filter.#": "1",
						"object_name_filter.0.inclusion_prefixes.#": "3",
						"time_amount": "11",
						"time_unit":   "YEARS",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig,
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

func TestObjectStorageObjectLifecyclePolicyResource_validations(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectLifecyclePolicyDestroy,
		Steps: []resource.TestStep{
			// verify baseline create
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create, objectLifecyclePolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", "lifecycle-policy-test-1"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// change order of inclusion prefixes
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
						getUpdatedRepresentationCopy("rules.object_name_filter.inclusion_prefixes", Representation{repType: Optional, create: []string{`lifecycle-policy-test-2`, `lifecycle-policy-test-1`}}, objectLifecyclePolicyRepresentation)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Remove inclusion prefixes to see plan has changed
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
						getUpdatedRepresentationCopy("rules.object_name_filter.inclusion_prefixes", Representation{repType: Optional, create: []string{}}, objectLifecyclePolicyRepresentation)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckObjectStorageObjectLifecyclePolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).objectStorageClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_objectstorage_object_lifecycle_policy" {
			noResourceFound = false
			request := oci_object_storage.GetObjectLifecyclePolicyRequest{}

			if value, ok := rs.Primary.Attributes["bucket"]; ok {
				request.BucketName = &value
			}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			_, err := client.GetObjectLifecyclePolicy(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
