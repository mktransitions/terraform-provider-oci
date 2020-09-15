---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_home"
sidebar_current: "docs-oci-resource-database-db_home"
description: |-
  Provides the Db Home resource in Oracle Cloud Infrastructure Database service
---

# oci_database_db_home
This resource provides the Db Home resource in Oracle Cloud Infrastructure Database service.

Creates a new Database Home in the specified DB system based on the request parameters you provide. Applies only to bare metal and Exadata DB systems.


## Example Usage

```hcl
resource "oci_database_db_home" "test_db_home" {
	#Required
	database {
		#Required
		admin_password = var.db_home_database_admin_password

		#Optional
		backup_id = oci_database_backup.test_backup.id
		backup_tde_password = var.db_home_database_backup_tde_password
		character_set = var.db_home_database_character_set
		database_id = oci_database_database.test_database.id
		database_software_image_id = oci_database_database_software_image.test_database_software_image.id
		db_backup_config {

			#Optional
			auto_backup_enabled = var.db_home_database_db_backup_config_auto_backup_enabled
			auto_backup_window = var.db_home_database_db_backup_config_auto_backup_window
			backup_destination_details {

				#Optional
				id = var.db_home_database_db_backup_config_backup_destination_details_id
				type = var.db_home_database_db_backup_config_backup_destination_details_type
			}
			recovery_window_in_days = var.db_home_database_db_backup_config_recovery_window_in_days
		}
		db_name = var.db_home_database_db_name
		db_workload = var.db_home_database_db_workload
		defined_tags = var.db_home_database_defined_tags
		freeform_tags = var.db_home_database_freeform_tags
		ncharacter_set = var.db_home_database_ncharacter_set
		pdb_name = var.db_home_database_pdb_name
		time_stamp_for_point_in_time_recovery = var.db_home_database_time_stamp_for_point_in_time_recovery
	}

	#Optional
	database_software_image_id = oci_database_database_software_image.test_database_software_image.id
	db_system_id = oci_database_db_system.test_db_system.id
	db_version {
	}
	defined_tags = var.db_home_defined_tags
	display_name = var.db_home_display_name
	freeform_tags = {"Department"= "Finance"}
	source = var.db_home_source
	vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `database` - (Required) (Updatable) 
	* `admin_password` - (Required) A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	* `backup_id` - (Required when source=DB_BACKUP) The backup [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `backup_tde_password` - (Required when source=DATABASE | DB_BACKUP) The password to open the TDE wallet.
	* `character_set` - (Applicable when source=NONE | VM_CLUSTER_NEW) The character set for the database.  The default is AL32UTF8. Allowed values are:

		AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS 
	* `database_id` - (Required when source=DATABASE) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `database_software_image_id` - (Applicable when source=NONE | VM_CLUSTER_NEW) The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	* `db_backup_config` - (Applicable when source=NONE | VM_CLUSTER_NEW) (Updatable) 
		* `auto_backup_enabled` - (Applicable when source=NONE | VM_CLUSTER_NEW) (Updatable) If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
		* `auto_backup_window` - (Applicable when source=NONE | VM_CLUSTER_NEW) (Updatable) Time window selected for initiating automatic backup for the database system. There are twelve available two-hour time windows. If no option is selected, a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).  Example: `SLOT_TWO` 
		* `backup_destination_details` - (Applicable when source=NONE | VM_CLUSTER_NEW) Backup destination details.
			* `id` - (Applicable when source=NONE | VM_CLUSTER_NEW) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
			* `type` - (Applicable when source=NONE | VM_CLUSTER_NEW) Type of the database backup destination. Supported values: `NFS`.
		* `recovery_window_in_days` - (Applicable when source=NONE | VM_CLUSTER_NEW) (Updatable) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
	* `db_name` - (Optional) The display name of the database to be created from the backup. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	* `db_workload` - (Applicable when source=NONE | VM_CLUSTER_NEW) The database workload type.
	* `defined_tags` - (Applicable when source=NONE | VM_CLUSTER_NEW) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
	* `freeform_tags` - (Applicable when source=NONE | VM_CLUSTER_NEW) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `ncharacter_set` - (Applicable when source=NONE | VM_CLUSTER_NEW) The national character set for the database.  The default is AL16UTF16. Allowed values are: AL16UTF16 or UTF8. 
	* `pdb_name` - (Applicable when source=NONE | VM_CLUSTER_NEW) The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
	* `time_stamp_for_point_in_time_recovery` - (Applicable when source=DATABASE) The point in time of the original database from which the new database is created. If not specifed, the latest backup is used to create the database.
* `database_software_image_id` - (Optional) The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `db_system_id` - (Required when source=DATABASE | DB_BACKUP | NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `db_version` - (Required when source=NONE | VM_CLUSTER_NEW) A valid Oracle Database version. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/20160918/DbVersionSummary/ListDbVersions) operation.
* `defined_tags` - (Optional Applicable when source=VM_CLUSTER_NEW) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Optional) The user-provided name of the Database Home.
* `freeform_tags` - (Optional Applicable when source=VM_CLUSTER_NEW) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `source` - (Optional) The source of database: NONE for creating a new database. DB_BACKUP for creating a new database by restoring from a database backup. VM_CLUSTER_NEW for creating a database for VM Cluster.
* `vm_cluster_id` - (Required when source=VM_CLUSTER_NEW) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_software_image_id` - The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `db_home_location` - The location of the Oracle Database Home.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `db_version` - The Oracle Database version.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-provided name for the Database Home. The name does not need to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
* `last_patch_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation is started.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `one_off_patches` - List of one-off patches for Database Homes.
* `state` - The current state of the Database Home.
* `time_created` - The date and time the Database Home was created.
* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.

## Import

DbHomes can be imported using the `id`, e.g.

```
$ terraform import oci_database_db_home.test_db_home "id"
```

Import is only supported for source=NONE

database.0.admin_password is not returned by the service for security reasons. Add the following to the resource:

```
    lifecycle {
        ignore_changes = ["database.0.admin_password"]
    }
```

The creation of an oci_database_db_system requires that it be created with exactly one oci_database_db_home. Therefore the first db home will have to be a property of the db system resource and any further db homes to be added to the db system will have to be added as first class resources using "oci_database_db_home".
