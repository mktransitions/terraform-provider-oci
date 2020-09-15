---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_group_backup"
sidebar_current: "docs-oci-resource-core-volume_group_backup"
description: |-
  Provides the Volume Group Backup resource in Oracle Cloud Infrastructure Core service
---

# oci_core_volume_group_backup
This resource provides the Volume Group Backup resource in Oracle Cloud Infrastructure Core service.

Creates a new backup volume group of the specified volume group.
For more information, see [Volume Groups](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).


## Example Usage

```hcl
resource "oci_core_volume_group_backup" "test_volume_group_backup" {
	#Required
	volume_group_id = oci_core_volume_group.test_volume_group.id

	#Optional
	compartment_id = var.compartment_id
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.volume_group_backup_display_name
	freeform_tags = {"Department"= "Finance"}
	type = var.volume_group_backup_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment that will contain the volume group backup. This parameter is optional, by default backup will be created in the same compartment and source volume group.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name for the volume group backup. Does not have to be unique and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `type` - (Optional) The type of backup to create. If omitted, defaults to incremental.
	* Allowed values are :
		* FULL
		* INCREMENTAL
* `volume_group_id` - (Required) The OCID of the volume group that needs to be backed up.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the volume group backup.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the volume group backup. Does not have to be unique and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the volume group backup.
* `size_in_gbs` - The aggregate size of the volume group backup, in GBs. 
* `size_in_mbs` - The aggregate size of the volume group backup, in MBs. 
* `state` - The current state of a volume group backup.
* `time_created` - The date and time the volume group backup was created. This is the time the actual point-in-time image of the volume group data was taken. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_request_received` - The date and time the request to create the volume group backup was received. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `type` - The type of backup.
* `unique_size_in_gbs` - The aggregate size used by the volume group backup, in GBs.  It is typically smaller than `size_in_gbs`, depending on the space consumed on the volume group and whether the volume backup is full or incremental. 
* `unique_size_in_mbs` - The aggregate size used by the volume group backup, in MBs.  It is typically smaller than `size_in_mbs`, depending on the space consumed on the volume group and whether the volume backup is full or incremental. 
* `volume_backup_ids` - OCIDs for the volume backups in this volume group backup.
* `volume_group_id` - The OCID of the source volume group.

## Import

VolumeGroupBackups can be imported using the `id`, e.g.

```
$ terraform import oci_core_volume_group_backup.test_volume_group_backup "id"
```

