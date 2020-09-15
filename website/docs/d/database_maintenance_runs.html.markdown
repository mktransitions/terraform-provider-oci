---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_maintenance_runs"
sidebar_current: "docs-oci-datasource-database-maintenance_runs"
description: |-
  Provides the list of Maintenance Runs in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_maintenance_runs
This data source provides the list of Maintenance Runs in Oracle Cloud Infrastructure Database service.

Gets a list of the maintenance runs in the specified compartment.


## Example Usage

```hcl
data "oci_database_maintenance_runs" "test_maintenance_runs" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.maintenance_run_availability_domain
	maintenance_type = var.maintenance_run_maintenance_type
	state = var.maintenance_run_state
	target_resource_id = oci_database_target_resource.test_target_resource.id
	target_resource_type = var.maintenance_run_target_resource_type
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) A filter to return only resources that match the given availability domain exactly.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `maintenance_type` - (Optional) The maintenance type.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.
* `target_resource_id` - (Optional) The target resource ID.
* `target_resource_type` - (Optional) The type of the target resource. Accepted values are: AUTONOMOUS_CONTAINER_DATABASE, AUTONOMOUS_EXADATA_INFRASTRUCTURE, EXADATA_DB_SYSTEM


## Attributes Reference

The following attributes are exported:

* `maintenance_runs` - The list of maintenance_runs.

### MaintenanceRun Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `description` - Description of the maintenance run.
* `display_name` - The user-friendly name for the maintenance run.
* `id` - The OCID of the maintenance run.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `maintenance_subtype` - Maintenance sub-type.
* `maintenance_type` - Maintenance type.
* `patch_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
* `state` - The current state of the maintenance run.
* `target_resource_id` - The ID of the target resource on which the maintenance run occurs.
* `target_resource_type` - The type of the target resource on which the maintenance run occurs.
* `time_ended` - The date and time the maintenance run was completed.
* `time_scheduled` - The date and time the maintenance run is scheduled to occur.
* `time_started` - The date and time the maintenance run starts.

