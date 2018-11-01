---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volumes"
sidebar_current: "docs-oci-datasource-core-volumes"
description: |-
  Provides the list of Volumes in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volumes
This data source provides the list of Volumes in Oracle Cloud Infrastructure Core service.

Lists the volumes in the specified compartment and availability domain.


## Example Usage

```hcl
data "oci_core_volumes" "test_volumes" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.volume_availability_domain}"
	display_name = "${var.volume_display_name}"
	state = "${var.volume_state}"
	volume_group_id = "${oci_core_volume_group.test_volume_group.id}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 
* `volume_group_id` - (Optional) The OCID of the volume group.


## Attributes Reference

The following attributes are exported:

* `volumes` - The list of volumes.

### Volume Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the volume.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the volume.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the volume.
* `is_hydrated` - Specifies whether the cloned volume's data has finished copying from the source volume or backup.
* `kms_key_id` - The OCID of the KMS key which is the master encryption key for the volume.
* `size_in_gbs` - The size of the volume in GBs.
* `size_in_mbs` - The size of the volume in MBs. This field is deprecated. Use `size_in_gbs` instead.
* `source_details` - The volume source, either an existing volume in the same availability domain or a volume backup. If null, an empty volume is created. 
	* `id` - The OCID of the volume or volume backup.
	* `type` - The type can be one of these values: `volume`, `volumeBackup`
* `state` - The current state of a volume.
* `time_created` - The date and time the volume was created. Format defined by RFC3339.
* `volume_group_id` - The OCID of the source volume group.

