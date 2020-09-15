---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_export"
sidebar_current: "docs-oci-resource-file_storage-export"
description: |-
  Provides the Export resource in Oracle Cloud Infrastructure File Storage service
---

# oci_file_storage_export
This resource provides the Export resource in Oracle Cloud Infrastructure File Storage service.

Creates a new export in the specified export set, path, and
file system.


## Example Usage

```hcl
resource "oci_file_storage_export" "test_export" {
	#Required
	export_set_id = oci_file_storage_export_set.test_export_set.id
	file_system_id = oci_file_storage_file_system.test_file_system.id
	path = var.export_path

	#Optional
	export_options {
		#Required
		source = var.export_export_options_source

		#Optional
		access = var.export_export_options_access
		anonymous_gid = var.export_export_options_anonymous_gid
		anonymous_uid = var.export_export_options_anonymous_uid
		identity_squash = var.export_export_options_identity_squash
		require_privileged_source_port = var.export_export_options_require_privileged_source_port
	}
}
```

## Argument Reference

The following arguments are supported:

* `export_options` - (Optional) (Updatable) Export options for the new export. If left unspecified, defaults to:

	[]

	**Note:** Mount targets do not have Internet-routable IP addresses.  Therefore they will not be reachable from the Internet, even if an associated `ClientOptions` item has a source of `0.0.0.0/0`.

	**If set to the empty array then the export will not be visible to any clients.**

	The export's `exportOptions` can be changed after creation using the `UpdateExport` operation. 
	* `access` - (Optional) (Updatable) Type of access to grant clients using the file system through this export. If unspecified defaults to `READ_ONLY`. 
	* `anonymous_gid` - (Optional) (Updatable) GID value to remap to when squashing a client GID (see identitySquash for more details.) If unspecified defaults to `65534`. 
	* `anonymous_uid` - (Optional) (Updatable) UID value to remap to when squashing a client UID (see identitySquash for more details.) If unspecified, defaults to `65534`. 
	* `identity_squash` - (Optional) (Updatable) Used when clients accessing the file system through this export have their UID and GID remapped to 'anonymousUid' and 'anonymousGid'. If `ALL`, all users and groups are remapped; if `ROOT`, only the root user and group (UID/GID 0) are remapped; if `NONE`, no remapping is done. If unspecified, defaults to `ROOT`. 
	* `require_privileged_source_port` - (Optional) (Updatable) If `true`, clients accessing the file system through this export must connect from a privileged source port. If unspecified, defaults to `true`. 
	* `source` - (Required) (Updatable) Clients these options should apply to. Must be a either single IPv4 address or single IPv4 CIDR block.

		**Note:** Access will also be limited by any applicable VCN security rules and the ability to route IP packets to the mount target. Mount targets do not have Internet-routable IP addresses. 
* `export_set_id` - (Required) The OCID of this export's export set.
* `file_system_id` - (Required) The OCID of this export's file system.
* `path` - (Required) Path used to access the associated file system.

	Avoid entering confidential information.

	Example: `/mediafiles` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `export_options` - Policies that apply to NFS requests made through this export. `exportOptions` contains a sequential list of `ClientOptions`. Each `ClientOptions` item defines the export options that are applied to a specified set of clients.

	For each NFS request, the first `ClientOptions` option in the list whose `source` attribute matches the source IP address of the request is applied.

	If a client source IP address does not match the `source` property of any `ClientOptions` in the list, then the export will be invisible to that client. This export will not be returned by `MOUNTPROC_EXPORT` calls made by the client and any attempt to mount or access the file system through this export will result in an error.

	**Exports without defined `ClientOptions` are invisible to all clients.**

	If one export is invisible to a particular client, associated file systems may still be accessible through other exports on the same or different mount targets. To completely deny client access to a file system, be sure that the client source IP address is not included in any export for any mount target associated with the file system. 
	* `access` - Type of access to grant clients using the file system through this export. If unspecified defaults to `READ_ONLY`. 
	* `anonymous_gid` - GID value to remap to when squashing a client GID (see identitySquash for more details.) If unspecified defaults to `65534`. 
	* `anonymous_uid` - UID value to remap to when squashing a client UID (see identitySquash for more details.) If unspecified, defaults to `65534`. 
	* `identity_squash` - Used when clients accessing the file system through this export have their UID and GID remapped to 'anonymousUid' and 'anonymousGid'. If `ALL`, all users and groups are remapped; if `ROOT`, only the root user and group (UID/GID 0) are remapped; if `NONE`, no remapping is done. If unspecified, defaults to `ROOT`. 
	* `require_privileged_source_port` - If `true`, clients accessing the file system through this export must connect from a privileged source port. If unspecified, defaults to `true`. 
	* `source` - Clients these options should apply to. Must be a either single IPv4 address or single IPv4 CIDR block.

		**Note:** Access will also be limited by any applicable VCN security rules and the ability to route IP packets to the mount target. Mount targets do not have Internet-routable IP addresses. 
* `export_set_id` - The OCID of this export's export set.
* `file_system_id` - The OCID of this export's file system.
* `id` - The OCID of this export.
* `path` - Path used to access the associated file system.

	Avoid entering confidential information.

	Example: `/accounting` 
* `state` - The current state of this export.
* `time_created` - The date and time the export was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

## Import

Exports can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_export.test_export "id"
```

