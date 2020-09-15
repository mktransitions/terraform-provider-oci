---
subcategory: "Logging"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_logging_log"
sidebar_current: "docs-oci-datasource-logging-log"
description: |-
  Provides details about a specific Log in Oracle Cloud Infrastructure Logging service
---

# Data Source: oci_logging_log
This data source provides details about a specific Log resource in Oracle Cloud Infrastructure Logging service.

Gets the log object config for log object OCID.


## Example Usage

```hcl
data "oci_logging_log" "test_log" {
	#Required
	log_group_id = oci_logging_log_group.test_log_group.id
	log_id = oci_logging_log.test_log.id
}
```

## Argument Reference

The following arguments are supported:

* `log_group_id` - (Required) OCID of a log group to work with.
* `log_id` - (Required) OCID of a log to work with.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that the resource belongs to.
* `configuration` - 
	* `compartment_id` - The OCID of the compartment that the resource belongs to.
	* `source` - 
		* `category` - Log object category.
		* `resource` - The unique identifier of the resource emitting the log.
		* `service` - Service generating log.
		* `source_type` - The source of the log.
			* **OCISERVICE:** Oracle Service. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of a user-friendly name. It has to be unique within enclosing resource, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the resource.
* `is_enabled` - Whether or not this resource is currently enabled.
* `log_group_id` - Log group OCID.
* `log_type` - The logType that the log object is for, custom or service.
* `retention_duration` - Log retention duration in days.
* `state` - The state of an pipeline.
* `tenancy_id` - The OCID of the tenancy.
* `time_created` - Time the resource was created.
* `time_last_modified` - Time the resource was last modified.

