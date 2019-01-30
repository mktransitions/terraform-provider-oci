// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	AutonomousDatabaseBackupResourceConfig = AutonomousDatabaseBackupResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", Optional, Update, autonomousDatabaseBackupRepresentation)

	autonomousDatabaseBackupSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_backup_id": Representation{repType: Required, create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}`},
	}

	autonomousDatabaseBackupDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Optional, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"display_name":           Representation{repType: Optional, create: `Monthly Backup`},
		"state":                  Representation{repType: Optional, create: `ACTIVE`},
		"filter":                 RepresentationGroup{Required, autonomousDatabaseBackupDataSourceFilterRepresentation}}
	autonomousDatabaseBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}`}},
	}

	autonomousDatabaseBackupRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"display_name":           Representation{repType: Required, create: `Monthly Backup`},
	}

	AutonomousDatabaseBackupResourceDependencies = AutonomousDatabaseResourceConfig
)

func TestDatabaseAutonomousDatabaseBackupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_backup.test_autonomous_database_backup"
	datasourceName := "data.oci_database_autonomous_database_backups.test_autonomous_database_backups"
	singularDatasourceName := "data.oci_database_autonomous_database_backup.test_autonomous_database_backup"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", Required, Create, autonomousDatabaseBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Monthly Backup"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database_backups", "test_autonomous_database_backups", Optional, Update, autonomousDatabaseBackupDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", Optional, Update, autonomousDatabaseBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "Monthly Backup"),

					resource.TestCheckResourceAttr(datasourceName, "autonomous_database_backups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.autonomous_database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_database_backups.0.display_name", "Monthly Backup"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.is_automatic"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.type"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", Required, Create, autonomousDatabaseBackupSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_backup_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Monthly Backup"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_automatic"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseBackupResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
				ExpectNonEmptyPlan:      true,
			},
		},
	})
}
