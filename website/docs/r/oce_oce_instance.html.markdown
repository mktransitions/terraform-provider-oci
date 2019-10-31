---
subcategory: "Oce"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oce_oce_instance"
sidebar_current: "docs-oci-resource-oce-oce_instance"
description: |-
  Provides the Oce Instance resource in Oracle Cloud Infrastructure Oce service
---

# oci_oce_oce_instance
This resource provides the Oce Instance resource in Oracle Cloud Infrastructure Oce service.

Creates a new OceInstance.


## Example Usage

```hcl
resource "oci_oce_oce_instance" "test_oce_instance" {
	#Required
	admin_email = "${var.oce_instance_admin_email}"
	compartment_id = "${var.compartment_id}"
	idcs_access_token = "${var.oce_instance_idcs_access_token}"
	name = "${var.oce_instance_name}"
	object_storage_namespace = "${var.oce_instance_object_storage_namespace}"
	tenancy_id = "${oci_identity_tenancy.test_tenancy.id}"
	tenancy_name = "${oci_identity_tenancy.test_tenancy.name}"

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = "${var.oce_instance_description}"
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `admin_email` - (Required) Admin Email for Notification
* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) OceInstance description
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `idcs_access_token` - (Required) Identity Cloud Service access token identifying a stripe and service administrator user. 
        **Note:** The `idcs_access_token` is stored in the Terraform state file.
* `name` - (Required) OceInstance Name
* `object_storage_namespace` - (Required) Object Storage Namespace of Tenancy
* `tenancy_id` - (Required) Tenancy Identifier
* `tenancy_name` - (Required) Tenancy Name


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `admin_email` - Admin Email for Notification
* `compartment_id` - Compartment Identifier
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - OceInstance description, can be updated
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `guid` - Unique GUID identifier that is immutable on creation
* `id` - Unique identifier that is immutable on creation
* `idcs_tenancy` - IDCS Tenancy Identifier
* `name` - OceInstance Name
* `object_storage_namespace` - Object Storage Namespace of tenancy
* `service` - SERVICE data. Example: `{"service": {"IDCS": "value"}}` 
* `state` - The current state of the file system.
* `state_message` - An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `tenancy_id` - Tenancy Identifier
* `tenancy_name` - Tenancy Name
* `time_created` - The time the the OceInstance was created. An RFC3339 formatted datetime string
* `time_updated` - The time the OceInstance was updated. An RFC3339 formatted datetime string

## Import

OceInstances can be imported using the `id`, e.g.

```
$ terraform import oci_oce_oce_instance.test_oce_instance "id"
```

