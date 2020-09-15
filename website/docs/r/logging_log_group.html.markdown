---
subcategory: "Logging"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_logging_log_group"
sidebar_current: "docs-oci-resource-logging-log_group"
description: |-
  Provides the Log Group resource in Oracle Cloud Infrastructure Logging service
---

# oci_logging_log_group
This resource provides the Log Group resource in Oracle Cloud Infrastructure Logging service.

Create new log group with unique display name. This call fails
if log group is already created with same displayName in the compartment.


## Example Usage

```hcl
resource "oci_logging_log_group" "test_log_group" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.log_group_display_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.log_group_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment that the resource belongs to.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description for this resource.
* `display_name` - (Required) (Updatable) The display name of a user-friendly name. It has to be unique within enclosing resource, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that the resource belongs to.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description for this resource.
* `display_name` - The display name of a user-friendly name. It has to be unique within enclosing resource, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the resource.
* `state` - The state of the log group object.
* `time_created` - Time the resource was created.
* `time_last_modified` - Time the resource was last modified.

## Import

LogGroups can be imported using the `id`, e.g.

```
$ terraform import oci_logging_log_group.test_log_group "id"
```

