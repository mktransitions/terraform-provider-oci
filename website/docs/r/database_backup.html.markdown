---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_backup"
sidebar_current: "docs-oci-resource-database-backup"
description: |-
  Provides the Backup resource in Oracle Cloud Infrastructure Database service
---

# oci_database_backup
This resource provides the Backup resource in Oracle Cloud Infrastructure Database service.

Creates a new backup in the specified database based on the request parameters you provide. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.


## Example Usage

```hcl
resource "oci_database_backup" "test_backup" {
	#Required
	database_id = "${oci_database_database.test_database.id}"
	display_name = "${var.backup_display_name}"
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `display_name` - (Required) The user-friendly name for the backup. The name does not have to be unique.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain where the database backup is stored.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_edition` - The Oracle Database edition of the DB system from which the database backup was taken. 
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `database_size_in_gbs` - The size of the database in gigabytes at the time the backup was taken. 
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup was completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.

## Import

Backups can be imported using the `id`, e.g.

```
$ terraform import oci_database_backup.test_backup "id"
```

