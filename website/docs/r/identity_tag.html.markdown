---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tag"
sidebar_current: "docs-oci-resource-identity-tag"
description: |-
  Provides the Tag resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_tag
This resource provides the Tag resource in Oracle Cloud Infrastructure Identity service.

Creates a new tag in the specified tag namespace.

You must specify either the OCID or the name of the tag namespace that will contain this tag definition.

You must also specify a *name* for the tag, which must be unique across all tags in the tag namespace
and cannot be changed. The name can contain any ASCII character except the space (_) or period (.) characters.
Names are case insensitive. That means, for example, "myTag" and "mytag" are not allowed in the same namespace.
If you specify a name that's already in use in the tag namespace, a 409 error is returned.

You must also specify a *description* for the tag.
It does not have to be unique, and you can change it with
[UpdateTag](https://docs.cloud.oracle.com/iaas/api/#/en/tagging/20170901/Tag/UpdateTag).


## Example Usage

```hcl
resource "oci_identity_tag" "test_tag" {
	#Required
	description = "${var.tag_description}"
	name = "${var.tag_name}"
	tag_namespace_id = "${oci_identity_tag_namespace.test_tag_namespace.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_cost_tracking = "${var.tag_is_cost_tracking}"
    is_retired = false
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) The description you assign to the tag during creation.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_cost_tracking` - (Optional) (Updatable) Indicates whether the tag is enabled for cost tracking. 
* `name` - (Required) The name you assign to the tag during creation. The name must be unique within the tag namespace and cannot be changed. 
* `tag_namespace_id` - (Required) The OCID of the tag namespace. 
* `is_retired` - (Optional) (Updatable) Indicates whether the tag is retired. See [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`` 
* `description` - The description you assign to the tag.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the tag definition.
* `is_cost_tracking` - Indicates whether the tag is enabled for cost tracking. 
* `is_retired` - Indicates whether the tag is retired. See [Retiring Key Definitions and Namespace Definitions](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/taggingoverview.htm#Retiring). 
* `name` - The name of the tag. The name must be unique across all tags in the tag namespace and can't be changed. 
* `tag_namespace_id` - The OCID of the namespace that contains the tag definition.
* `time_created` - Date and time the tag was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

