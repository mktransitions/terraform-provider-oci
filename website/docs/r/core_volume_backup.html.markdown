---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_backup"
sidebar_current: "docs-oci-resource-core-volume_backup"
description: |-
  Provides the Volume Backup resource in Oracle Cloud Infrastructure Core service
---

# oci_core_volume_backup
This resource provides the Volume Backup resource in Oracle Cloud Infrastructure Core service.

Creates a new backup of the specified volume. For general information about volume backups,
see [Overview of Block Volume Service Backups](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumebackups.htm)

When the request is received, the backup object is in a REQUEST_RECEIVED state.
When the data is imaged, it goes into a CREATING state.
After the backup is fully uploaded to the cloud, it goes into an AVAILABLE state.


## Example Usage

```hcl
resource "oci_core_volume_backup" "test_volume_backup" {
	#Required
	volume_id = "${oci_core_volume.test_volume.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.volume_backup_display_name}"
	freeform_tags = {"Department"= "Finance"}
	type = "${var.volume_backup_type}"
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name for the volume backup. Does not have to be unique and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `type` - (Optional) The type of backup to create. If omitted, defaults to INCREMENTAL.
* `volume_id` - (Required) The OCID of the volume that needs to be backed up.
* `source_details` - (Optional) Details of the volume backup source in the cloud. 
    * `region` - The region of the volume backup source.
    * `volume_backup_id` - The OCID of the source volume backup.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the volume backup.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the volume backup. Does not have to be unique and it's changeable. Avoid entering confidential information. 
* `expiration_time` - The date and time the volume backup will expire and be automatically deleted. Format defined by RFC3339. This parameter will always be present for backups that were created automatically by a scheduled-backup policy. For manually created backups, it will be absent, signifying that there is no expiration time and the backup will last forever until manually deleted. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the volume backup.
* `size_in_gbs` - The size of the volume, in GBs. 
* `size_in_mbs` - The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use `size_in_gbs`. 
* `source_type` - Specifies whether the backup was created manually, or via scheduled backup policy.
* `source_volume_backup_id` - The OCID of the source volume backup.
* `state` - The current state of a volume backup.
* `time_created` - The date and time the volume backup was created. This is the time the actual point-in-time image of the volume data was taken. Format defined by RFC3339. 
* `time_request_received` - The date and time the request to create the volume backup was received. Format defined by RFC3339. 
* `type` - The type of a volume backup. Supported values are 'FULL' or 'INCREMENTAL'.
* `unique_size_in_gbs` - The size used by the backup, in GBs. It is typically smaller than `size_in_gbs`, depending on the space consumed on the volume and whether the backup is full or incremental. 
* `unique_size_in_mbs` - The size used by the backup, in MBs. It is typically smaller than `size_in_mbs`, depending on the space consumed on the volume and whether the backup is full or incremental. This field is deprecated. Please use `unique_size_in_gbs`. 
* `volume_id` - The OCID of the volume.

## Import

VolumeBackups can be imported using the `id`, e.g.

```
$ terraform import oci_core_volume_backup.test_volume_backup "id"
```

