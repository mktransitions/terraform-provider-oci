---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_data_warehouses"
sidebar_current: "docs-oci-datasource-database-autonomous_data_warehouses"
description: |-
  Provides the list of Autonomous Data Warehouses in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_data_warehouses
This data source provides the list of Autonomous Data Warehouses in Oracle Cloud Infrastructure Database service.

Gets a list of Autonomous Data Warehouses.

**IMPORTANT:** This data source is being **deprecated**, use `oci_database_autonomous_databases` instead.

## Example Usage

```hcl
data "oci_database_autonomous_data_warehouses" "test_autonomous_data_warehouses" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.autonomous_data_warehouse_display_name}"
	state = "${var.autonomous_data_warehouse_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `autonomous_data_warehouses` - The list of autonomous_data_warehouses.

### AutonomousDataWarehouse Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_strings` - The connection string used to connect to the Data Warehouse. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Data Warehouse for the password value.
	* `all_connection_strings` - Returns all connection strings that can be used to connect to the Autonomous Data Warehouse. For more information, please see [Predefined Database Service Names for Autonomous Transaction Processing](https://docs.oracle.com/en/cloud/paas/atp-cloud/atpug/connect-predefined.html#GUID-9747539B-FD46-44F1-8FF8-F5AC650F15BE) 
	* `high` - The High database service provides the highest level of resources to each SQL statement resulting in the highest performance, but supports the fewest number of concurrent SQL statements.
	* `low` - The Low database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	* `medium` - The Medium database service provides a lower level of resources to each SQL statement potentially resulting a lower level of performance, but supports more concurrent SQL statements.
* `cpu_core_count` - The number of CPU cores to be made available to the database.
* `data_storage_size_in_tbs` - The quantity of data in the database, in terabytes.
* `db_name` - The database name.
* `db_version` - A valid Oracle Database version for Autonomous Data Warehouse.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Autonomous Data Warehouse. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse.
* `license_model` - The Oracle license model that applies to the Oracle Autonomous Data Warehouse. The default is BRING_YOUR_OWN_LICENSE. 
* `lifecycle_details` - Information about the current lifecycle state.
* `service_console_url` - The URL of the Service Console for the Data Warehouse.
* `state` - The current state of the database.
* `time_created` - The date and time the database was created.

