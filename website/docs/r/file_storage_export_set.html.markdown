---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_export_set"
sidebar_current: "docs-oci-resource-file_storage-export_set"
description: |-
  Provides the Export Set resource in Oracle Cloud Infrastructure File Storage service
---

# oci_file_storage_export_set
This resource provides the Export Set resource in Oracle Cloud Infrastructure File Storage service.

The export set resource can neither be directly created, nor destroyed.

An export set is created by the service automatically when a mount target is created.
When a mount target is deleted, the export set associated with it is also deleted automatically.

However, export sets expose a few attributes that can be updated.

Hence we provide this resource for managing the already created export set from within Terraform.
Only one export set resource should be created per mount target.

## Example Usage

```hcl
resource "oci_file_storage_export_set" "test_export_set" {
    #Required
    mount_target_id = "${oci_file_storage_mount_target.test_mount_target.id}"
  
    #Optional
    display_name = "${var.export_set_name}"
    max_fs_stat_bytes = 23843202333
    max_fs_stat_files = 223442
}
```

## Argument Reference

The following arguments are supported:

* `mount_target_id` - (Required) (Updatable) The OCID of the mount target that the export set is associated with
* `display_name` - (Optional) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My export set` 
* `max_fs_stat_bytes` - (Optional) (Updatable) Controls the maximum `tbytes`, `fbytes`, and `abytes`, values reported by `NFS FSSTAT` calls through any associated mount targets. This is an advanced feature. For most applications, use the default value. The `tbytes` value reported by `FSSTAT` will be `maxFsStatBytes`. The value of `fbytes` and `abytes` will be `maxFsStatBytes` minus the metered size of the file system. If the metered size is larger than `maxFsStatBytes`, then `fbytes` and `abytes` will both be '0'. 
* `max_fs_stat_files` - (Optional) (Updatable) Controls the maximum `tfiles`, `ffiles`, and `afiles` values reported by `NFS FSSTAT` calls through any associated mount targets. This is an advanced feature. For most applications, use the default value. The `tfiles` value reported by `FSSTAT` will be `maxFsStatFiles`. The value of `ffiles` and `afiles` will be `maxFsStatFiles` minus the metered size of the file system. If the metered size is larger than `maxFsStatFiles`, then `ffiles` and `afiles` will both be '0'. 

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the export set is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the export set.
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My export set` 
* `id` - The OCID of the export set.
* `max_fs_stat_bytes` - Controls the maximum `tbytes`, `fbytes`, and `abytes`, values reported by `NFS FSSTAT` calls through any associated mount targets. This is an advanced feature. For most applications, use the default value. The `tbytes` value reported by `FSSTAT` will be `maxFsStatBytes`. The value of `fbytes` and `abytes` will be `maxFsStatBytes` minus the metered size of the file system. If the metered size is larger than `maxFsStatBytes`, then `fbytes` and `abytes` will both be '0'. 
* `max_fs_stat_files` - Controls the maximum `tfiles`, `ffiles`, and `afiles` values reported by `NFS FSSTAT` calls through any associated mount targets. This is an advanced feature. For most applications, use the default value. The `tfiles` value reported by `FSSTAT` will be `maxFsStatFiles`. The value of `ffiles` and `afiles` will be `maxFsStatFiles` minus the metered size of the file system. If the metered size is larger than `maxFsStatFiles`, then `ffiles` and `afiles` will both be '0'. 
* `state` - The current state of the export set.
* `time_created` - The date and time the export set was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the virtual cloud network (VCN) the export set is in.

## Import

ExportSets can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_export_set.test_export_set "id"
```

