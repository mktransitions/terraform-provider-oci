// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	ExportSetRequiredOnlyResource = ExportSetResourceDependencies +
		generateResourceFromRepresentationMap("oci_file_storage_export_set", "test_export_set", Required, Create, exportSetRepresentation)

	exportSetDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":        Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"id":                  Representation{repType: Optional, create: `${oci_file_storage_export_set.test_export_set.id}`},
		"state":               Representation{repType: Optional, create: `ACTIVE`},
		"filter":              RepresentationGroup{Required, exportSetDataSourceFilterRepresentation}}
	exportSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_file_storage_export_set.test_export_set.id}`}},
	}

	exportSetRepresentation = map[string]interface{}{
		"mount_target_id":   Representation{repType: Required, create: `${oci_file_storage_mount_target.test_mount_target.id}`},
		"display_name":      Representation{repType: Optional, create: `export set display name`},
		"max_fs_stat_bytes": Representation{repType: Optional, create: `23843202333`},
		"max_fs_stat_files": Representation{repType: Optional, create: `223442`},
	}

	ExportSetResourceDependencies = MountTargetRequiredOnlyResource
)

func TestFileStorageExportSetResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_export_set.test_export_set"
	datasourceName := "data.oci_file_storage_export_sets.test_export_sets"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ExportSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_file_storage_export_set", "test_export_set", Required, Create, exportSetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "max_fs_stat_bytes"),
					resource.TestCheckResourceAttrSet(resourceName, "max_fs_stat_files"),
					resource.TestCheckResourceAttrSet(resourceName, "mount_target_id"),
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ExportSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_file_storage_export_set", "test_export_set", Optional, Update, exportSetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "export set display name"),
					resource.TestCheckResourceAttr(resourceName, "max_fs_stat_bytes", "23843202333"),
					resource.TestCheckResourceAttr(resourceName, "max_fs_stat_files", "223442"),
					resource.TestCheckResourceAttrSet(resourceName, "mount_target_id"),
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
					generateDataSourceFromRepresentationMap("oci_file_storage_export_sets", "test_export_sets", Optional, Update, exportSetDataSourceRepresentation) +
					compartmentIdVariableStr + ExportSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_file_storage_export_set", "test_export_set", Optional, Update, exportSetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "export_sets.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "export_sets.0.state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(datasourceName, "export_sets.0.time_created"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"mount_target_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}
