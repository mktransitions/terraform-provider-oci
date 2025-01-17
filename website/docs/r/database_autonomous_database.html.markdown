---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database"
sidebar_current: "docs-oci-resource-database-autonomous_database"
description: |-
  Provides the Autonomous Database resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_database
This resource provides the Autonomous Database resource in Oracle Cloud Infrastructure Database service.

Creates a new Autonomous Database.

This API must be called on the remote region where the peer needs to be created.

## Example Usage

```hcl
resource "oci_database_autonomous_database" "test_autonomous_database" {
	#Required
	admin_password = var.autonomous_database_admin_password
	compartment_id = var.compartment_id
	db_name = var.autonomous_database_db_name

	#Optional
	are_primary_whitelisted_ips_used = var.autonomous_database_are_primary_whitelisted_ips_used
	auto_refresh_frequency_in_seconds = var.autonomous_database_auto_refresh_frequency_in_seconds
	auto_refresh_point_lag_in_seconds = var.autonomous_database_auto_refresh_point_lag_in_seconds
	autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
	autonomous_database_backup_id = oci_database_autonomous_database_backup.test_autonomous_database_backup.id
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
	autonomous_maintenance_schedule_type = var.autonomous_database_autonomous_maintenance_schedule_type
    backup_retention_period_in_days = var.autonomous_database_backup_retention_period_in_days
	character_set = var.autonomous_database_character_set
	clone_type = var.autonomous_database_clone_type
	compute_count = var.autonomous_database_compute_count
	compute_model = var.autonomous_database_compute_model
	cpu_core_count = var.autonomous_database_cpu_core_count
	customer_contacts {

		#Optional
		email = var.autonomous_database_customer_contacts_email
	}
	data_safe_status = var.autonomous_database_data_safe_status
	data_storage_size_in_gb = var.autonomous_database_data_storage_size_in_gb
	data_storage_size_in_tbs = var.autonomous_database_data_storage_size_in_tbs
	database_edition = var.autonomous_database_database_edition
	db_name = var.autonomous_database_db_name
	db_tools_details {
		#Required
		name = var.autonomous_database_db_tools_details_name

		#Optional
		compute_count = var.autonomous_database_db_tools_details_compute_count
		is_enabled = var.autonomous_database_db_tools_details_is_enabled
		max_idle_time_in_minutes = var.autonomous_database_db_tools_details_max_idle_time_in_minutes
	}
	db_version = var.autonomous_database_db_version
	db_workload = var.autonomous_database_db_workload
	defined_tags = var.autonomous_database_defined_tags
	disaster_recovery_type = var.autonomous_database_disaster_recovery_type
	display_name = var.autonomous_database_display_name
	encryption_key {

		#Optional
		arn_role = var.autonomous_database_encryption_key_arn_role
		autonomous_database_provider = var.autonomous_database_encryption_key_autonomous_database_provider
		certificate_directory_name = var.autonomous_database_encryption_key_certificate_directory_name
		certificate_id = oci_apigateway_certificate.test_certificate.id
		directory_name = var.autonomous_database_encryption_key_directory_name
		external_id = oci_database_external.test_external.id
		key_arn = var.autonomous_database_encryption_key_key_arn
		key_name = oci_kms_key.test_key.name
		kms_key_id = oci_kms_key.test_key.id
		okv_kms_key = var.autonomous_database_encryption_key_okv_kms_key
		okv_uri = var.autonomous_database_encryption_key_okv_uri
		service_endpoint_uri = var.autonomous_database_encryption_key_service_endpoint_uri
		vault_id = oci_kms_vault.test_vault.id
		vault_uri = var.autonomous_database_encryption_key_vault_uri
	}
	freeform_tags = {"Department"= "Finance"}
    in_memory_percentage = var.autonomous_database_in_memory_percentage
	is_access_control_enabled = var.autonomous_database_is_access_control_enabled
	is_auto_scaling_enabled = var.autonomous_database_is_auto_scaling_enabled
	is_auto_scaling_for_storage_enabled = var.autonomous_database_is_auto_scaling_for_storage_enabled
	is_backup_retention_locked = var.autonomous_database_is_backup_retention_locked
	is_data_guard_enabled = var.autonomous_database_is_data_guard_enabled
	is_dedicated = var.autonomous_database_is_dedicated
	is_dev_tier = var.autonomous_database_is_dev_tier
	is_free_tier = var.autonomous_database_is_free_tier
	is_local_data_guard_enabled = var.autonomous_database_is_local_data_guard_enabled
	is_mtls_connection_required = var.autonomous_database_is_mtls_connection_required
	is_preview_version_with_service_terms_accepted = var.autonomous_database_is_preview_version_with_service_terms_accepted
	is_replicate_automatic_backups = var.autonomous_database_is_replicate_automatic_backups
	kms_key_id = oci_kms_key.test_key.id
	license_model = var.autonomous_database_license_model
	max_cpu_core_count = var.autonomous_database_max_cpu_core_count
	ncharacter_set = var.autonomous_database_ncharacter_set
	nsg_ids = var.autonomous_database_nsg_ids
	ocpu_count = var.autonomous_database_ocpu_count
        private_endpoint_ip = var.autonomous_database_private_endpoint_ip
	private_endpoint_label = var.autonomous_database_private_endpoint_label
	refreshable_mode = var.autonomous_database_refreshable_mode
	resource_pool_leader_id = oci_database_resource_pool_leader.test_resource_pool_leader.id
	resource_pool_summary {
		#Optional
		is_disabled = var.autonomous_database_resource_pool_summary_is_disabled
		pool_size = var.autonomous_database_resource_pool_summary_pool_size
	}
	scheduled_operations {
		#Required
		day_of_week {
			#Required
			name = var.autonomous_database_scheduled_operations_day_of_week_name
		}

		#Optional
		scheduled_start_time = var.autonomous_database_scheduled_operations_scheduled_start_time
		scheduled_stop_time = var.autonomous_database_scheduled_operations_scheduled_stop_time
	}
	secret_id = oci_vault_secret.test_secret.id
	secret_version_number = var.autonomous_database_secret_version_number
	security_attributes = var.autonomous_database_security_attributes
	source = var.autonomous_database_source
	source_id = oci_database_source.test_source.id
	standby_whitelisted_ips = var.autonomous_database_standby_whitelisted_ips
	subnet_id = oci_core_subnet.test_subnet.id
	subscription_id = oci_onesubscription_subscription.test_subscription.id
	time_of_auto_refresh_start = var.autonomous_database_time_of_auto_refresh_start
	timestamp = var.autonomous_database_timestamp
	use_latest_available_backup_time_stamp = var.autonomous_database_use_latest_available_backup_time_stamp
	vault_id = oci_kms_vault.test_vault.id
	whitelisted_ips = var.autonomous_database_whitelisted_ips
}
```

## Argument Reference

The following arguments are supported:

* `admin_password` - (Optional) (Updatable) The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing. The password is mandatory if source value is "BACKUP_FROM_ID", "BACKUP_FROM_TIMESTAMP", "DATABASE" or "NONE".
* `are_primary_whitelisted_ips_used` - (Optional) (Updatable) This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled. It's value would be `TRUE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses primary IP access control list (ACL) for standby. It's value would be `FALSE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses different IP access control list (ACL) for standby compared to primary.
* `auto_refresh_frequency_in_seconds` - (Applicable when source=CLONE_TO_REFRESHABLE) (Updatable) The frequency a refreshable clone is refreshed after auto-refresh is enabled. The minimum is 1 hour. The maximum is 7 days. The date and time that auto-refresh is enabled is controlled by the `timeOfAutoRefreshStart` parameter.
* `auto_refresh_point_lag_in_seconds` - (Applicable when source=CLONE_TO_REFRESHABLE) (Updatable) The time, in seconds, the data of the refreshable clone lags the primary database at the point of refresh. The minimum is 0 minutes (0 mins means refresh to the latest available timestamp). The maximum is 7 days. The lag time increases after refreshing until the next data refresh happens.
* `autonomous_container_database_id` - (Optional) The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `autonomous_database_backup_id` - (Required when source=BACKUP_FROM_ID) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database Backup that you will clone to create a new Autonomous Database.
* `autonomous_database_id` - (Required when source=BACKUP_FROM_TIMESTAMP) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that you will clone to create a new Autonomous Database.
* `autonomous_maintenance_schedule_type` - (Optional) The maintenance schedule type of the Autonomous Database Serverless instances. The EARLY maintenance schedule of this Autonomous Database follows a schedule that applies patches prior to the REGULAR schedule.The REGULAR maintenance schedule of this Autonomous Database follows the normal cycle.
* `backup_retention_period_in_days` - (Optional) (Updatable) Retention period, in days, for backups.
* `character_set` - (Optional) The character set for the autonomous database.  The default is AL32UTF8. Allowed values for an Autonomous Database on Serverless infrastructure as returned by [List Autonomous Database Character Sets](/autonomousDatabaseCharacterSets)

  For an Autonomous Database on dedicated infrastructure, the allowed values are:

  AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS
* `clone_type` - (Required when source=BACKUP_FROM_ID | BACKUP_FROM_TIMESTAMP | DATABASE) The Autonomous Database clone type. This parameter is not used to create a refreshable clone type, and for refreshable clones one must use the (source=CLONE_TO_REFRESHABLE) parameter.
	* `FULL` - This option creates a new database that includes all source database data.
	* `METADATA` - This option creates a new database that includes the source database schema and select metadata, but not the source database data.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of the Autonomous Database.
* `compute_count` - (Optional) (Updatable) The compute amount available to the database. Minimum and maximum values depend on the compute model and whether the database is an Autonomous Database Serverless instance or an Autonomous Database on Dedicated Exadata Infrastructure. For an Autonomous Database Serverless instance, the 'ECPU' compute model requires a minimum value of one, for databases in the elastic resource pool and minimum value of two, otherwise. Required when using the `computeModel` parameter. When using `cpuCoreCount` parameter, it is an error to specify computeCount to a non-null value. Providing `computeModel` and `computeCount` is the preferred method for both OCPU and ECPU.
* `compute_model` - (Optional) (Updatable) The compute model of the Autonomous Database. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. ECPU compute model is the recommended model and OCPU compute model is legacy.
* `cpu_core_count` - (Optional) (Updatable) The number of CPU cores to be made available to the database. For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
  **Note:** This parameter cannot be used with the `ocpuCount` parameter. This input is ignored for Always Free resources.
    * The data type must be an *integer*.
    * The minimum number of cores for all types of autonomous database is *1*
    * The maximum number of cores is as follows:
        * Autonomous Database Serverless instances: The maximum number of cores is *128*.
        * Autonomous Databases on dedicated Exadata infrastructure: The maximum number of cores is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
        * Providing `computeModel` and `computeCount` is the preferred method for setting CPUs for both OCPU and ECPU.
* `customer_contacts` - (Optional) (Updatable) Customer Contacts.
	* `email` - (Optional) (Updatable) The email address used by Oracle to send notifications regarding databases and infrastructure.
* `data_safe_status` - (Optional) (Updatable) Status of the Data Safe registration for this Autonomous Database. Could be REGISTERED or NOT_REGISTERED.
* `data_storage_size_in_gb` - (Optional) (Updatable) The size, in gigabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. The maximum storage value is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.

  **Notes**
	* This parameter is only supported for dedicated Exadata infrastructure.
	* This parameter cannot be used with the `dataStorageSizeInTBs` parameter. 
* `data_storage_size_in_tbs` - (Optional) (Updatable) The size, in terabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. For Autonomous Databases on dedicated Exadata infrastructure, the maximum storage value is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.  A full Exadata service is allocated when the Autonomous Database size is set to the upper limit (384 TB).

	**Note:** This parameter cannot be used with the `dataStorageSizeInGBs` parameter. This input is ignored for Always Free resources.
* `database_edition` - (Optional) (Updatable) The Oracle Database Edition that applies to the Autonomous databases. It can be set to `ENTERPRISE_EDITION` or `STANDARD_EDITION`.
* `db_name` - (Optional) The database name. The name must begin with an alphabetic character and can contain a maximum of 14 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy. It is required in all cases except when creating a cross-region Autonomous Data Guard standby instance or a cross-region disaster recovery standby instance.
* `db_tools_details` - (Optional) (Updatable) The list of database tools details.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, isLocalDataGuardEnabled, or isFreeTier. 
	* `compute_count` - (Optional) (Updatable) Compute used by database tools.
	* `is_enabled` - (Optional) (Updatable) Indicates whether tool is enabled.
	* `max_idle_time_in_minutes` - (Optional) (Updatable) The max idle time, in minutes, after which the VM used by database tools will be terminated.
	* `name` - (Required) (Updatable) Name of database tool.
* `db_version` - (Optional) (Updatable) A valid Oracle Database version for Autonomous Database.`db_workload` AJD and APEX are only supported for `db_version` `19c` and above.
* `db_workload` - (Optional) (Updatable) The Autonomous Database workload type. The following values are valid:
	* OLTP - indicates an Autonomous Transaction Processing database
	* DW - indicates an Autonomous Data Warehouse database
	* AJD - indicates an Autonomous JSON Database
	* APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier. *Note: `db_workload` can only be updated from AJD to OLTP or from a free OLTP to AJD.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `disaster_recovery_type` - (Required when source=CROSS_TENANCY_DISASTER_RECOVERY) Indicates the disaster recovery (DR) type of the standby Autonomous Database Serverless instance. Autonomous Data Guard (ADG) DR type provides business critical DR with a faster recovery time objective (RTO) during failover or switchover. Backup-based DR type provides lower cost DR with a slower RTO during failover or switchover. 
* `display_name` - (Optional) (Updatable) The user-friendly name for the Autonomous Database. The name does not have to be unique.
* `encryption_key` - (Optional) (Updatable) Details of the Autonomous Database encryption key.
	* `arn_role` - (Applicable when provider=AWS) (Updatable) AWS ARN role
	* `autonomous_database_provider` - (Optional) (Updatable) The provider for the Autonomous Database encryption key.
	* `certificate_directory_name` - (Required when provider=OKV) (Updatable) OKV certificate directory name
	* `certificate_id` - (Applicable when provider=OKV) (Updatable) OKV certificate id
	* `directory_name` - (Required when provider=OKV) (Updatable) OKV wallet directory name
	* `external_id` - (Applicable when provider=AWS) (Updatable) AWS external ID
	* `key_arn` - (Required when provider=AWS) (Updatable) AWS key ARN
	* `key_name` - (Required when provider=AZURE) (Updatable) Azure key name
	* `kms_key_id` - (Required when provider=OCI) (Updatable) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `okv_kms_key` - (Required when provider=OKV) (Updatable) UUID of OKV KMS Key
	* `okv_uri` - (Required when provider=OKV) (Updatable) URI of OKV server
	* `service_endpoint_uri` - (Required when provider=AWS) (Updatable) AWS key service endpoint URI
	* `vault_id` - (Required when provider=OCI) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	* `vault_uri` - (Required when provider=AZURE) (Updatable) Azure vault URI
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `in_memory_percentage` - (Optional) (Updatable) The percentage of the System Global Area(SGA) assigned to In-Memory tables in Autonomous Database. This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform.
* `is_access_control_enabled` - (Optional) (Updatable) Indicates if the database-level access control is enabled. If disabled, database access is defined by the network security rules. If enabled, database access is restricted to the IP addresses defined by the rules specified with the `whitelistedIps` property. While specifying `whitelistedIps` rules is optional, if database-level access control is enabled and no rules are specified, the database will become inaccessible. The rules can be added later using the `UpdateAutonomousDatabase` API operation or edit option in console. When creating a database clone, the desired access control setting should be specified. By default, database-level access control will be disabled for the clone.

	This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform. For Autonomous Database Serverless instances, `whitelistedIps` is used. 
* `is_auto_scaling_enabled` - (Optional) (Updatable) Indicates if auto scaling is enabled for the Autonomous Database CPU core count. The default value is `TRUE`. 
* `is_auto_scaling_for_storage_enabled` - (Optional) (Updatable) Indicates if auto scaling is enabled for the Autonomous Database storage. The default value is `FALSE`. 
* `is_backup_retention_locked` - (Optional) (Updatable) True if the Autonomous Database is backup retention locked.
* `is_data_guard_enabled` - (Optional) (Updatable) **Deprecated.** Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure. 
* `is_dedicated` - (Optional) True if the database is on [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html). 
* `is_dev_tier` - (Optional) (Updatable) Autonomous Database for Developers are free Autonomous Databases that developers can use to build and test new applications.With Autonomous these database instancess instances, you can try new Autonomous Database features for free and apply them to ongoing or new development projects. Developer database comes with limited resources and is, therefore, not suitable for large-scale testing and production deployments. When you need more compute or storage resources, you can transition to a paid database licensing by cloning your developer database into a regular Autonomous Database. See [Autonomous Database documentation](https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/eddjo/index.html) for more details.         
* `is_free_tier` - (Optional) (Updatable) Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isLocalDataGuardEnabled When `db_workload` is `AJD` it cannot be `true`.
* `is_local_data_guard_enabled` - (Optional) (Updatable) Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. It takes boolean values. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
* `is_mtls_connection_required` - (Optional) (Updatable) Specifies if the Autonomous Database requires mTLS connections.

	This may not be updated in parallel with any of the following: licenseModel, databaseEdition, cpuCoreCount, computeCount, dataStorageSizeInTBs, whitelistedIps, openMode, permissionLevel, db-workload, privateEndpointLabel, nsgIds, customerContacts, dbVersion, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.

	Service Change: The default value of the isMTLSConnectionRequired attribute will change from true to false on July 1, 2023 in the following APIs:
	* CreateAutonomousDatabase
	* GetAutonomousDatabase
	* UpdateAutonomousDatabase Details: Prior to the July 1, 2023 change, the isMTLSConnectionRequired attribute default value was true. This applies to Autonomous Database Serverless. Does this impact me? If you use or maintain custom scripts or Terraform scripts referencing the CreateAutonomousDatabase, GetAutonomousDatabase, or UpdateAutonomousDatabase APIs, you want to check, and possibly modify, the scripts for the changed default value of the attribute. Should you choose not to leave your scripts unchanged, the API calls containing this attribute will continue to work, but the default value will switch from true to false. How do I make this change? Using either Oracle Cloud Infrastructure SDKs or command line tools, update your custom scripts to explicitly set the isMTLSConnectionRequired attribute to true. 
* `is_preview_version_with_service_terms_accepted` - (Optional) If set to `TRUE`, indicates that an Autonomous Database preview version is being provisioned, and that the preview version's terms of service have been accepted. Note that preview version software is only available for Autonomous Database Serverless instances (https://docs.oracle.com/en/cloud/paas/autonomous-database/serverless/). 
* `is_replicate_automatic_backups` - (Applicable when source=CROSS_REGION_DISASTER_RECOVERY) If true, 7 days worth of backups are replicated across regions for Cross-Region ADB or Backup-Based DR between Primary and Standby. If false, the backups taken on the Primary are not replicated to the Standby database.
* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `key_version_id` - (Optional) The OCID of the key version that is used in rotate key operations.
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle services in the cloud. License Included allows you to subscribe to new Oracle Database software licenses and the Oracle Database service. Note that when provisioning an [Autonomous Database on dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null. It is already set at the Autonomous Exadata Infrastructure level. When provisioning an [Autonomous Database Serverless] (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) database, if a value is not specified, the system defaults the value to `BRING_YOUR_OWN_LICENSE`. Bring your own license (BYOL) also allows you to select the DB edition using the optional parameter.
* `is_auto_scaling_for_storage_enabled` - (Optional) (Updatable) Indicates if auto scaling is enabled for the Autonomous Database storage. The default value is `FALSE`.
* `is_dedicated` - (Optional) True if the database is on [dedicated Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm).
* `is_free_tier` - (Optional) (Updatable) Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled. When `db_workload` is `AJD` or `APEX` it cannot be `true`.
* `is_mtls_connection_required` - (Optional) (Updatable) Indicates whether the Autonomous Database requires mTLS connections.
* `is_preview_version_with_service_terms_accepted` - (Optional) If set to `TRUE`, indicates that an Autonomous Database preview version is being provisioned, and that the preview version's terms of service have been accepted. Note that preview version software is only available for databases on [Autonomous Database Serverless](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).
* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle services in the cloud. License Included allows you to subscribe to new Oracle Database software licenses and the Oracle Database service. Note that when provisioning an [Autonomous Database on dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null. It is already set at the Autonomous Exadata Infrastructure level. When provisioning an [Autonomous Database Serverless] (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) database, if a value is not specified, the system defaults the value to `BRING_YOUR_OWN_LICENSE`. Bring your own license (BYOL) also allows you to select the DB edition using the optional parameter.

	This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, dataStorageSizeInTBs, adminPassword, isMTLSConnectionRequired, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, or isFreeTier. It is a required field when `db_workload` is AJD and needs to be set to `LICENSE_INCLUDED` as AJD does not support default `license_model` value `BRING_YOUR_OWN_LICENSE`.
* `max_cpu_core_count` - (Optional) (Updatable) **Deprecated.** The number of Max OCPU cores to be made available to the autonomous database with auto scaling of cpu enabled. 
* `ncharacter_set` - (Optional) The character set for the Autonomous Database.  The default is AL32UTF8. Use [List Autonomous Database Character Sets](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/autonomousDatabaseCharacterSets/ListAutonomousDatabaseCharacterSets) to list the allowed values for an Autonomous Database Serverless. For an Autonomous Database on dedicated Exadata infrastructure, the allowed values are: AL16UTF16 or UTF8.
* `nsg_ids` - (Optional) (Updatable) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
* `is_refreshable_clone` - (Applicable when source=CLONE_TO_REFRESHABLE) (Updatable) True for creating a refreshable clone and False for detaching the clone from source Autonomous Database. Detaching is one time operation and clone becomes a regular Autonomous Database.
* `is_remote_data_guard_enabled` - Indicates whether the Autonomous Database has Cross Region Data Guard enabled. It takes boolean values. Not applicable to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud. License Included allows you to subscribe to new Oracle Database software licenses and the Database service. Note that when provisioning an Autonomous Database on [dedicated Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the Autonomous Exadata Infrastructure level. When using [shared Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`. It is a required field when `db_workload` is AJD and needs to be set to `LICENSE_INCLUDED` as AJD does not support default `license_model` value `BRING_YOUR_OWN_LICENSE`.
* `ncharacter_set` - (Optional) The national character set for the autonomous database.  The default is AL16UTF16. Allowed values are: AL16UTF16 or UTF8.
* `ocpu_count` - (Optional) (Updatable) The number of OCPU cores to be made available to the database.


    * Providing `computeModel` and `computeCount` is the preferred method for setting CPUs for both OCPU and ECPU.
  The following points apply:
	* For Autonomous Databases on dedicated Exadata infrastructure, to provision less than 1 core, enter a fractional value in an increment of 0.1. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. (Note that fractional OCPU values are not supported for Autonomous Databasese on shared Exadata infrastructure.)
	* To provision 1 or more cores, you must enter an integer between 1 and the maximum number of cores available for the infrastructure shape. For example, you can provision 2 cores or 3 cores, but not 2.5 cores. This applies to Autonomous Databases on both shared and dedicated Exadata infrastructure.

  For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.

  **Note:** This parameter cannot be used with the `cpuCoreCount` parameter.
* `operations_insights_status` - (Optional) (Updatable) Status of Operations Insights for this Autonomous Database. Values supported are `ENABLED` and `NOT_ENABLED`
* `private_endpoint_label` - (Optional) (Updatable) (Optional) (Updatable) The resource's private endpoint label.
	* Setting the endpoint label to a non-empty string creates a private endpoint database.
	* Resetting the endpoint label to an empty string, after the creation of the private endpoint database, changes the private endpoint database to a public endpoint database.
	* Setting the endpoint label to a non-empty string value, updates to a new private endpoint database, when the database is disabled and re-enabled.
  
* `refreshable_mode` - (Applicable when source=CLONE_TO_REFRESHABLE) (Updatable) The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
* `remote_disaster_recovery_type` - (Required when source=CROSS_REGION_DISASTER_RECOVERY) Indicates the cross-region disaster recovery (DR) type of the standby Autonomous Database Serverless instance. Autonomous Data Guard (ADG) DR type provides business critical DR with a faster recovery time objective (RTO) during failover or switchover. Backup-based DR type provides lower cost DR with a slower RTO during failover or switchover. 
* `resource_pool_leader_id` - (Optional) (Updatable) The unique identifier for leader autonomous database OCID [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `resource_pool_summary` - (Optional) (Updatable) The configuration details for resource pool
	* `is_disabled` - (Optional) (Updatable) Indicates if the resource pool should be deleted for the Autonomous Database.  
	* `pool_size` - (Optional) (Updatable) Resource pool size.
* `scheduled_operations` - (Optional) (Updatable) The list of scheduled operations. Consists of values such as dayOfWeek, scheduledStartTime, scheduledStopTime.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier. 
	* `day_of_week` - (Required) (Updatable) Day of the week.
		* `name` - (Required) (Updatable) Name of the day of the week.
	* `scheduled_start_time` - (Optional) (Updatable) auto start time. value must be of ISO-8601 format "HH:mm"
	* `scheduled_stop_time` - (Optional) (Updatable) auto stop time. value must be of ISO-8601 format "HH:mm"
* `secret_id` - (Optional) (Updatable) The Oracle Cloud Infrastructure vault secret [/Content/General/Concepts/identifiers.htm]OCID.

	This cannot be used in conjunction with adminPassword. 
* `secret_version_number` - (Optional) (Updatable) The version of the vault secret. If no version is specified, the latest version will be used.
* `security_attributes` - (Optional) (Updatable) Security Attributes for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}` 
* `source` - (Optional) The source of the database:
	* Use `NONE` for creating a new Autonomous Database.
	* Use `DATABASE` for creating a new Autonomous Database by cloning an existing running Autonomous Database from the latest timestamp, also provide the source database OCID in the `source_id` parameter.
	* Use `CROSS_REGION_DATAGUARD` to create a standby Data Guard database in another region, also provide the remote primary database OCID in the `source_id` parameter.
	* Use `CLONE_TO_REFRESHABLE` for creating a refreshable clone.
	
  For [Autonomous Database Serverless](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) instances, the following cloning options are available:
	* Use `BACKUP_FROM_ID` for creating a new Autonomous Database by cloning from a specified backup. Also provide the backup OCID in the `autonomous_database_backup_id` parameter.
	* Use `BACKUP_FROM_TIMESTAMP` for creating a point-in-time Autonomous Database clone using backups. Also provide the backup timestamp in the `timestamp` parameter. For more information, see [Cloning and Moving an Autonomous Database](https://docs.oracle.com/en/cloud/paas/autonomous-database/adbsa/clone-autonomous-database.html#GUID-D771796F-5081-4CFB-A7FF-0F893EABD7BC).
* `source_id` - (Required when source=CLONE_TO_REFRESHABLE | CROSS_REGION_DATAGUARD | CROSS_REGION_DISASTER_RECOVERY | DATABASE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that will be used to create a new standby database for the Data Guard association.
* `standby_whitelisted_ips` - (Optional) (Updatable) The client IP access control list (ACL). This feature is available for [Autonomous Database Serverless] (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. If `arePrimaryWhitelistedIpsUsed` is 'TRUE' then Autonomous Database uses this primary's IP access control list (ACL) for the disaster recovery peer called `standbywhitelistedips`.

  For Autonomous Database Serverless instances, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]` For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations. Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`

	For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.

    If `arePrimaryWhitelistedIpsUsed` is 'TRUE' then Autonomous Database uses `whitelisted_ips` primary's IP access control list (ACL) as `standbywhitelistedips` for the disaster recovery peer.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
* `subnet_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with. This the only parameter to configure private endpoint, VCN details are obtained from the `subnet_id`.

  **Subnet Restrictions:**
	* For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	* For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	* For Autonomous Database, setting this will disable public secure access to the database.

* `subscription_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
  These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and the backup subnet.
* `timestamp` - (Applicable when source=BACKUP_FROM_TIMESTAMP) The timestamp specified for the point-in-time clone of the source Autonomous Database. The timestamp must be in the past.
* `use_latest_available_backup_time_stamp` - (Applicable when source=BACKUP_FROM_TIMESTAMP) Clone from latest available backup timestamp.
* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `whitelisted_ips` - (Optional) (Updatable) The client IP access control list (ACL). This feature is available for autonomous databases on [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.

  For Autonomous Database Serverless instances, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]` For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations. Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`

	For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier. To remove all whitelisted IPs, set the field to a list with an empty string `[""]`.

* `switchover_to` - (Optional) It is applicable only when `is_local_data_guard_enabled` is true. Could be set to `PRIMARY` or `STANDBY`. Default value is `PRIMARY`.
* `switchover_to_remote_peer_id` - (Optional) (Updatable) It is applicable only when `dataguard_region_type` and `role` are set, and `is_dedicated` is false. For Autonomous Database Serverless instances, Data Guard associations have designated primary and standby regions, and these region types do not change when the database changes roles. It takes the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the remote peer to switchover to and the API is called from the remote region.
* `rotate_key_trigger` - (Optional) (Updatable) An optional property when flipped triggers rotation of KMS key. It is only applicable on dedicated databases i.e. where `is_dedicated` is true.
* `is_shrink_only` - (Optional) (Updatable) An optional property when enabled triggers the Shrinking of Autonomous Database once. To trigger Shrinking of ADB once again, this flag needs to be disabled and re-enabled again. It should not be passed during create database operation. It is only applicable on Serverless databases i.e. where `is_dedicated` is false.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `actual_used_data_storage_size_in_tbs` - The current amount of storage in use for user and system data, in terabytes (TB). 
* `allocated_storage_size_in_tbs` - The amount of storage currently allocated for the database tables and billed for, rounded up. When auto-scaling is not enabled, this value is equal to the `dataStorageSizeInTBs` value. You can compare this value to the `actualUsedDataStorageSizeInTBs` value to determine if a manual shrink operation is appropriate for your allocated storage.

  **Note:** Auto-scaling does not automatically decrease allocated storage when data is deleted from the database.
* `apex_details` - Information about Oracle APEX Application Development.
	* `apex_version` - The Oracle APEX Application Development version.
	* `ords_version` - The Oracle REST Data Services (ORDS) version.
* `are_primary_whitelisted_ips_used` - This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled. It's value would be `TRUE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses primary IP access control list (ACL) for standby. It's value would be `FALSE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses different IP access control list (ACL) for standby compared to primary. 
* `autonomous_container_database_id` - The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Used only by Autonomous Database on Dedicated Exadata Infrastructure.
* `availability_domain` - The availability domain where the Autonomous Database Serverless instance is located.
* `autonomous_maintenance_schedule_type` - The maintenance schedule type of the Autonomous Database Serverless. An EARLY maintenance schedule follows a schedule applying patches prior to the REGULAR schedule. A REGULAR maintenance schedule follows the normal cycle
* `available_upgrade_versions` - List of Oracle Database versions available for a database upgrade. If there are no version upgrades available, this list is empty.
* `backup_config` - Autonomous Database configuration details for storing [manual backups](https://docs.oracle.com/en/cloud/paas/autonomous-database/adbsa/backup-restore.html#GUID-9035DFB8-4702-4CEB-8281-C2A303820809) in the [Object Storage](https://docs.cloud.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) service.
	* `manual_backup_bucket_name` - Name of [Object Storage](https://docs.cloud.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) bucket to use for storing manual backups.
	* `manual_backup_type` - The manual backup destination type.
* `backup_retention_period_in_days` - Retention period, in days, for backups.
* `character_set` - The character set for the autonomous database.  The default is AL32UTF8. Allowed values are:

  AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS
* `cluster_placement_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group of the Autonomous Serverless Database.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_count` - The compute amount (CPUs) available to the database. Minimum and maximum values depend on the compute model and whether the database is an Autonomous Database Serverless instance or an Autonomous Database on Dedicated Exadata Infrastructure.  For an Autonomous Database Serverless instance, the 'ECPU' compute model requires a minimum value of one, for databases in the elastic resource pool and minimum value of two, otherwise. Required when using the `computeModel` parameter. When using `cpuCoreCount` parameter, it is an error to specify computeCount to a non-null value. Providing `computeModel` and `computeCount` is the preferred method for both OCPU and ECPU. 
* `compute_model` - The compute model of the Autonomous Database. This is required if using the `computeCount` parameter. If using `cpuCoreCount` then it is an error to specify `computeModel` to a non-null value. ECPU compute model is the recommended model and OCPU compute model is legacy.
* `connection_strings` - The connection string used to connect to the Autonomous Database. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Database for the password value.
	* `all_connection_strings` - Returns all connection strings that can be used to connect to the Autonomous Database. For more information, please see [Predefined Database Service Names for Autonomous Transaction Processing](https://docs.oracle.com/en/cloud/paas/atp-cloud/atpug/connect-predefined.html#GUID-9747539B-FD46-44F1-8FF8-F5AC650F15BE) 
	* `dedicated` - The database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	* `high` - The High database service provides the highest level of resources to each SQL statement resulting in the highest performance, but supports the fewest number of concurrent SQL statements.
	* `low` - The Low database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	* `medium` - The Medium database service provides a lower level of resources to each SQL statement potentially resulting a lower level of performance, but supports more concurrent SQL statements.
	* `profiles` - A list of connection string profiles to allow clients to group, filter and select connection string values based on structured metadata. 
		* `consumer_group` - Consumer group used by the connection.
		* `display_name` - A user-friendly name for the connection.
		* `host_format` - Host format used in connection string.
		* `is_regional` - True for a regional connection string, applicable to cross-region DG only.
		* `protocol` - Protocol used by the connection.
		* `session_mode` - Specifies whether the listener performs a direct hand-off of the session, or redirects the session. In RAC deployments where SCAN is used, sessions are redirected to a Node VIP. Use `DIRECT` for direct hand-offs. Use `REDIRECT` to redirect the session.
		* `syntax_format` - Specifies whether the connection string is using the long (`LONG`), Easy Connect (`EZCONNECT`), or Easy Connect Plus (`EZCONNECTPLUS`) format. Autonomous Database Serverless instances always use the long format. 
		* `tls_authentication` - Specifies whether the TLS handshake is using one-way (`SERVER`) or mutual (`MUTUAL`) authentication.
		* `value` - Connection string value.
* `connection_urls` - The URLs for accessing Oracle Application Express (APEX) and SQL Developer Web with a browser from a Compute instance within your VCN or that has a direct connection to your VCN. Note that these URLs are provided by the console only for databases on [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).  Example: `{"sqlDevWebUrl": "https://<hostname>/ords...", "apexUrl", "https://<hostname>/ords..."}` 
	* `apex_url` - Oracle Application Express (APEX) URL.
	* `database_transforms_url` - The URL of the Database Transforms for the Autonomous Database.
	* `graph_studio_url` - The URL of the Graph Studio for the Autonomous Database.
	* `machine_learning_notebook_url` - The URL of the Oracle Machine Learning (OML) Notebook for the Autonomous Database.
	* `machine_learning_user_management_url` - Oracle Machine Learning user management URL.
	* `mongo_db_url` - The URL of the MongoDB API for the Autonomous Database.
	* `ords_url` - The Oracle REST Data Services (ORDS) URL of the Web Access for the Autonomous Database.
	* `sql_dev_web_url` - Oracle SQL Developer Web URL.
* `cpu_core_count` - The number of CPU cores to be made available to the database. When the ECPU is selected, the value for cpuCoreCount is 0. For Autonomous Database on Dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	
    **Note:** This parameter cannot be used with the `ocpuCount` parameter.
    * The data type must be an *integer*.
    * The minimum number of cores for all types of autonomous database is *1*
    * The maximum number of cores is as follows:
        * Autonomous Database Serverless instances: The maximum number of cores is *128*.
        * Autonomous Databases on dedicated Exadata infrastructure: The maximum number of cores is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
* `customer_contacts` - Customer Contacts.
    * `email` - The email address used by Oracle to send notifications regarding databases and infrastructure.
* `data_safe_status` - Status of the Data Safe registration for this Autonomous Database. Could be REGISTERED or NOT_REGISTERED.
* `data_storage_size_in_gb` - The quantity of data in the database, in gigabytes.
* `data_storage_size_in_tbs` - The quantity of data in the database, in terabytes.
* `database_edition` - The Oracle Database Edition that applies to the Autonomous databases. It can be set to `ENTERPRISE_EDITION` or `STANDARD_EDITION`.
* `database_management_status` - Status of Database Management for this Autonomous Database.
* `dataguard_region_type` - **Deprecated** (Optional) The Autonomous Data Guard region type of the Autonomous Database. For Autonomous Database Serverless instances, Data Guard associations have designated primary (`PRIMARY_DG_REGION`) and standby (`REMOTE_STANDBY_DG_REGION`) regions, and these region types do not change when the database changes roles. The standby regions in Data Guard associations can be the same region designated as the primary region, or they can be remote regions. Certain database administrative operations may be available only in the primary region of the Data Guard association, and cannot be performed when the database using the "primary" role is operating in a remote Data Guard standby region.
* `db_name` - The database name.
* `db_tools_details` - The list of database tools details.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, isLocalDataGuardEnabled, or isFreeTier. 
	* `compute_count` - Compute used by database tools.
	* `is_enabled` - Indicates whether tool is enabled.
	* `max_idle_time_in_minutes` - The max idle time, in minutes, after which the VM used by database tools will be terminated.
	* `name` - Name of database tool.
* `db_version` - A valid Oracle Database version for Autonomous Database.
* `db_workload` - The Autonomous Database workload type. The following values are valid:
	* OLTP - indicates an Autonomous Transaction Processing database
	* DW - indicates an Autonomous Data Warehouse database
	* AJD - indicates an Autonomous JSON Database
	* APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `disaster_recovery_region_type` - **Deprecated** The disaster recovery (DR) region type of the Autonomous Database. For Serverless Autonomous Databases, DR associations have designated primary (`PRIMARY`) and standby (`REMOTE`) regions. These region types do not change when the database changes roles. The standby region in DR associations can be the same region as the primary region, or they can be in a remote regions. Some database administration operations may be available only in the primary region of the DR association, and cannot be performed when the database using the primary role is operating in a remote region.
* `display_name` - The user-friendly name for the Autonomous Database. The name does not have to be unique.
* `encryption_key` - Details of the Autonomous Database encryption key.
	* `arn_role` - AWS ARN role
	* `autonomous_database_provider` - The provider for the Autonomous Database encryption key.
	* `certificate_directory_name` - OKV certificate directory name
	* `certificate_id` - OKV certificate id
	* `directory_name` - OKV wallet directory name
	* `external_id` - AWS external ID
	* `key_arn` - AWS key ARN
	* `key_name` - Azure key name
	* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `okv_kms_key` - UUID of OKV KMS Key
	* `okv_uri` - URI of OKV server
	* `service_endpoint_uri` - AWS key service endpoint URI
	* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	* `vault_uri` - Azure vault URI
* `encryption_key_history_entry` - Key History Entry.
	* `encryption_key` - Details of the Autonomous Database encryption key.
		* `arn_role` - AWS ARN role
		* `autonomous_database_provider` - The provider for the Autonomous Database encryption key.
		* `certificate_directory_name` - OKV certificate directory name
		* `certificate_id` - OKV certificate id
		* `directory_name` - OKV wallet directory name
		* `external_id` - AWS external ID
		* `key_arn` - AWS key ARN
		* `key_name` - Azure key name
		* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
		* `okv_kms_key` - UUID of OKV KMS Key
		* `okv_uri` - URI of OKV server
		* `service_endpoint_uri` - AWS key service endpoint URI
		* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
		* `vault_uri` - Azure vault URI
	* `time_activated` - The date and time the encryption key was activated.
* `failed_data_recovery_in_seconds` - Indicates the number of seconds of data loss for a Data Guard failover.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
* `in_memory_area_in_gbs` - The area assigned to In-Memory tables in Autonomous Database.
* `in_memory_percentage` - The percentage of the System Global Area(SGA) assigned to In-Memory tables in Autonomous Database. This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform.
* `infrastructure_type` - The infrastructure type this resource belongs to.
* `is_access_control_enabled` - Indicates if the database-level access control is enabled. If disabled, database access is defined by the network security rules. If enabled, database access is restricted to the IP addresses defined by the rules specified with the `whitelistedIps` property. While specifying `whitelistedIps` rules is optional, if database-level access control is enabled and no rules are specified, the database will become inaccessible. The rules can be added later using the `UpdateAutonomousDatabase` API operation or edit option in console. When creating a database clone, the desired access control setting should be specified. By default, database-level access control will be disabled for the clone.

	This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform. For Autonomous Database Serverless instances, `whitelistedIps` is used. 
* `is_auto_scaling_enabled` - Indicates if auto scaling is enabled for the Autonomous Database CPU core count. The default value is `TRUE`. 
* `is_auto_scaling_for_storage_enabled` - Indicates if auto scaling is enabled for the Autonomous Database storage. The default value is `FALSE`. 
* `is_backup_retention_locked` - Indicates if the Autonomous Database is backup retention locked.
* `is_data_guard_enabled` - **Deprecated.** Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure. 
* `is_dedicated` - True if the database uses [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html). 
* `is_dev_tier` - Autonomous Database for Developers are free Autonomous Databases that developers can use to build and test new applications.With Autonomous these database instancess instances, you can try new Autonomous Database features for free and apply them to ongoing or new development projects. Developer database comes with limited resources and is, therefore, not suitable for large-scale testing and production deployments. When you need more compute or storage resources, you can transition to a paid database licensing by cloning your developer database into a regular Autonomous Database. See [Autonomous Database documentation](https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/eddjo/index.html) for more details.         
* `is_free_tier` - Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isLocalDataGuardEnabled
* `is_local_data_guard_enabled` - Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
* `is_mtls_connection_required` - Specifies if the Autonomous Database requires mTLS connections.

	This may not be updated in parallel with any of the following: licenseModel, databaseEdition, cpuCoreCount, computeCount, dataStorageSizeInTBs, whitelistedIps, openMode, permissionLevel, db-workload, privateEndpointLabel, nsgIds, customerContacts, dbVersion, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.

	Service Change: The default value of the isMTLSConnectionRequired attribute will change from true to false on July 1, 2023 in the following APIs:
	* CreateAutonomousDatabase
	* GetAutonomousDatabase
	* UpdateAutonomousDatabase Details: Prior to the July 1, 2023 change, the isMTLSConnectionRequired attribute default value was true. This applies to Autonomous Database Serverless instances. Does this impact me? If you use or maintain custom scripts or Terraform scripts referencing the CreateAutonomousDatabase, GetAutonomousDatabase, or UpdateAutonomousDatabase APIs, you want to check, and possibly modify, the scripts for the changed default value of the attribute. Should you choose not to leave your scripts unchanged, the API calls containing this attribute will continue to work, but the default value will switch from true to false. How do I make this change? Using either Oracle Cloud Infrastructure SDKs or command line tools, update your custom scripts to explicitly set the isMTLSConnectionRequired attribute to true. 
* `is_preview` - Indicates if the Autonomous Database version is a preview version.
* `is_reconnect_clone_enabled` - Indicates if the refreshable clone can be reconnected to its source database.
* `is_refreshable_clone` - Indicates if the Autonomous Database is a refreshable clone.

	This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier. 
* `is_remote_data_guard_enabled` - Indicates whether the Autonomous Database has Cross Region Data Guard enabled. It takes boolean values. Not applicable to Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
* `key_history_entry` - Key History Entry.
    * `id` - The id of the Autonomous Database [Vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts) service key management history entry.
    * `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
    * `time_activated` - The date and time the kms key activated.
    * `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault. This is used in Autonomous Databases on Serverless instances and dedicated Exadata infrastructure.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_lifecycle_details` - KMS key lifecycle details.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
* `license_model` - The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle services in the cloud. License Included allows you to subscribe to new Oracle Database software licenses and the Oracle Database service. Note that when provisioning an [Autonomous Database on dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null. It is already set at the Autonomous Exadata Infrastructure level. When provisioning an [Autonomous Database Serverless] (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) database, if a value is not specified, the system defaults the value to `BRING_YOUR_OWN_LICENSE`. Bring your own license (BYOL) also allows you to select the DB edition using the optional parameter.

	This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, dataStorageSizeInTBs, adminPassword, isMTLSConnectionRequired, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
* `lifecycle_details` - Information about the current lifecycle state.
* `local_adg_auto_failover_max_data_loss_limit` - Parameter that allows users to select an acceptable maximum data loss limit in seconds, up to which Automatic Failover will be triggered when necessary for a Local Autonomous Data Guard
* `local_disaster_recovery_type` - Indicates the local disaster recovery (DR) type of the Autonomous Database Serverless instance. Autonomous Data Guard (ADG) DR type provides business critical DR with a faster recovery time objective (RTO) during failover or switchover. Backup-based DR type provides lower cost DR with a slower RTO during failover or switchover. 
* `local_standby_db` - Autonomous Data Guard standby database details. 
	* `availability_domain` - The availability domain of a local Autonomous Data Guard standby database of an Autonomous Database Serverless instance.
	* `lag_time_in_seconds` - The amount of time, in seconds, that the data of the standby database lags the data of the primary database. Can be used to determine the potential data loss in the event of a failover.
	* `lifecycle_details` - Additional information about the current lifecycle state.
	* `state` - The current state of the Autonomous Database.
	* `time_data_guard_role_changed` - The date and time the Autonomous Data Guard role was switched for the standby Autonomous Database.
	* `time_disaster_recovery_role_changed` - The date and time the Disaster Recovery role was switched for the standby Autonomous Database.
* `long_term_backup_schedule` - Details for the long-term backup schedule.
	* `is_disabled` - Indicates if the long-term backup schedule should be deleted. The default value is `FALSE`. 
	* `repeat_cadence` - The frequency of the long-term backup schedule
	* `retention_period_in_days` - Retention period, in days, for long-term backups
	* `time_of_backup` - The timestamp for the long-term backup schedule. For a MONTHLY cadence, months having fewer days than the provided date will have the backup taken on the last day of that month.
* `max_cpu_core_count` - **Deprecated.** The number of Max OCPU cores to be made available to the autonomous database with auto scaling of cpu enabled. 
* `memory_per_oracle_compute_unit_in_gbs` - The amount of memory (in GBs) enabled per OCPU or ECPU. See [Compute Models in Autonomous Database on Dedicated Exadata Infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbak) for more details. This parameter is not used for Autonomous database Serverless.
* `local_disaster_recovery_type` - Indicates the local disaster recovery (DR) type of the Serverless Autonomous Database. Autonomous Data Guard (`ADG`) DR type provides business critical DR with a faster recovery time objective (RTO) during failover or switchover. Backup-based (`BACKUP_BASED`) DR type provides lower cost DR with a slower RTO during failover or switchover. 
* `local_standby_db` - Autonomous Data Guard local (same region) standby database details.
    * `lag_time_in_seconds` - The amount of time, in seconds, that the data of the standby database lags the data of the primary database. Can be used to determine the potential data loss in the event of a failover.
    * `lifecycle_details` - Additional information about the current lifecycle state.
    * `state` - The current state of the Autonomous Database.
    * `time_data_guard_role_changed` - The date and time the Autonomous Data Guard role was switched for the standby Autonomous Database.
    * `time_disaster_recovery_role_changed` - The date and time the Disaster Recovery role was switched for the standby Autonomous Database.
* `ncharacter_set` - The national character set for the autonomous database.  The default is AL16UTF16. Allowed values are: AL16UTF16 or UTF8. 
* `net_services_architecture` - Enabling SHARED server architecture enables a database server to allow many client processes to share very few server processes, thereby increasing the number of supported users.
* `next_long_term_backup_time_stamp` - The date and time when the next long-term backup would be created.
* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
    * A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
* `ocpu_count` - The number of OCPU cores to be made available to the database.

	The following points apply:
	* For Autonomous Databases on Dedicated Exadata Infrastructure, to provision less than 1 core, enter a fractional value in an increment of 0.1. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. (Note that fractional OCPU values are not supported for Autonomous Database Serverless instances.)
	* To provision cores, enter an integer between 1 and the maximum number of cores available for the infrastructure shape. For example, you can provision 2 cores or 3 cores, but not 2.5 cores. This applies to Autonomous Databases on both serverless and dedicated Exadata infrastructure.
	* For Autonomous Database Serverless instances, this parameter is not used.
    * Providing `computeModel` and `computeCount` is the preferred method for setting CPUs for both OCPU and ECPU.

  For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.


	**Note:** This parameter cannot be used with the `cpuCoreCount` parameter. 
* `open_mode` - Indicates the Autonomous Database mode. The database can be opened in `READ_ONLY` or `READ_WRITE` mode.

	This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier. 
* `operations_insights_status` - Status of Operations Insights for this Autonomous Database.
* `peer_db_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of standby databases located in Autonomous Data Guard remote regions that are associated with the source database. Note that for Autonomous Database Serverless instances, standby databases located in the same region as the source primary database do not have OCIDs.
* `permission_level` - The Autonomous Database permission level. Restricted mode allows access only by admin users.

	This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, or isFreeTier. 
* `private_endpoint` - The private endpoint for the resource. This parameter is not used in Autonomous Databases using Serverless infrastructure and Exadata Cloud@Customer infrastructure.
* `private_endpoint_ip` - The private endpoint Ip address for the resource.
* `private_endpoint_label` - The resource's private endpoint label.
	* Setting the endpoint label to a non-empty string creates a private endpoint database.
	* Resetting the endpoint label to an empty string, after the creation of the private endpoint database, changes the private endpoint database to a public endpoint database.
	* Setting the endpoint label to a non-empty string value, updates to a new private endpoint database, when the database is disabled and re-enabled.
* `provisionable_cpus` - An array of CPU values that an Autonomous Database can be scaled to.
* `public_connection_urls` - The Public URLs of Private Endpoint database for accessing Oracle Application Express (APEX) and SQL Developer Web with a browser from a Compute instance within your VCN or that has a direct connection to your VCN.
	* `apex_url` - Oracle Application Express (APEX) URL.
	* `database_transforms_url` - The URL of the Database Transforms for the Autonomous Database.
	* `graph_studio_url` - The URL of the Graph Studio for the Autonomous Database.
	* `machine_learning_notebook_url` - The URL of the Oracle Machine Learning (OML) Notebook for the Autonomous Database.
	* `machine_learning_user_management_url` - Oracle Machine Learning user management URL.
	* `mongo_db_url` - The URL of the MongoDB API for the Autonomous Database.
	* `ords_url` - The Oracle REST Data Services (ORDS) URL of the Web Access for the Autonomous Database.
	* `sql_dev_web_url` - Oracle SQL Developer Web URL.
* `public_endpoint` - The public endpoint for the private endpoint enabled resource.
* `refreshable_mode` - The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
* `refreshable_status` - The refresh status of the clone. REFRESHING indicates that the clone is currently being refreshed with data from the source Autonomous Database.
* `remote_disaster_recovery_configuration` - Configurations of a Disaster Recovery.
    * `disaster_recovery_type` - Indicates the disaster recovery (DR) type of the Shared Autonomous Database. Autonomous Data Guard (ADG) DR type provides business critical DR with a faster recovery time objective (RTO) during failover or switchover. Backup-based DR type provides lower cost DR with a slower RTO during failover or switchover. 
    * `is_replicate_automatic_backups` - If true, 7 days worth of backups are replicated across regions for Cross-Region ADB or Backup-Based DR between Primary and Standby. If false, the backups taken on the Primary are not replicated to the Standby database.
    * `is_snapshot_standby` - Indicates if user wants to convert to a snapshot standby. For example, true would set a standby database to snapshot standby database. False would set a snapshot standby database back to regular standby database. 
    * `time_snapshot_standby_enabled_till` - Time and date stored as an RFC 3339 formatted timestamp string. For example, 2022-01-01T12:00:00.000Z would set a limit for the snapshot standby to be converted back to a cross-region standby database.
    * `is_snapshot_standby` - Indicates if user wants to convert to a snapshot standby. For example, true would set a standby database to snapshot standby database. False would set a snapshot standby database back to regular standby database. 
    * `time_snapshot_standby_enabled_till` - Time and date stored as an RFC 3339 formatted timestamp string. For example, 2022-01-01T12:00:00.000Z would set a limit for the snapshot standby to be converted back to a cross-region standby database.
* `role` - The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled. 
* `scheduled_operations` - The list of scheduled operations. Consists of values such as dayOfWeek, scheduledStartTime, scheduledStopTime.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier. 
	* `day_of_week` - Day of the week.
		* `name` - Name of the day of the week.
	* `scheduled_start_time` - auto start time. value must be of ISO-8601 format "HH:mm"
	* `scheduled_stop_time` - auto stop time. value must be of ISO-8601 format "HH:mm"
* `security_attributes` - Security Attributes for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}` 
* `service_console_url` - The URL of the Service Console for the Autonomous Database.
* `source_id` - (Required when source=CLONE_TO_REFRESHABLE | CROSS_REGION_DATAGUARD | DATABASE | CROSS_REGION_DISASTER_RECOVERY) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that will be used to create a new standby database for the Data Guard association.
* `standby_db` - **Deprecated** Autonomous Data Guard standby database details. 
	* `availability_domain` - The availability domain of a local Autonomous Data Guard standby database of an Autonomous Database Serverless instance.
	* `lag_time_in_seconds` - The amount of time, in seconds, that the data of the standby database lags the data of the primary database. Can be used to determine the potential data loss in the event of a failover.
	* `lifecycle_details` - Additional information about the current lifecycle state.
	* `state` - The current state of the Autonomous Database.
	* `time_data_guard_role_changed` - The date and time the Autonomous Data Guard role was switched for the standby Autonomous Database.
	* `time_disaster_recovery_role_changed` - The date and time the Disaster Recovery role was switched for the standby Autonomous Database.
* `standby_whitelisted_ips` - The client IP access control list (ACL). This feature is available for [Autonomous Database Serverless] (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. If `arePrimaryWhitelistedIpsUsed` is 'TRUE' then Autonomous Database uses this primary's IP access control list (ACL) for the disaster recovery peer called `standbywhitelistedips`.

  For Autonomous Database Serverless instances, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]` For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations. Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`

	For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
* `state` - The current state of the Autonomous Database.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.

  **Subnet Restrictions:**
    * For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
    * For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
    * For Autonomous Database, setting this will disable public secure access to the database.

* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
  These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and the backup subnet.
* `supported_regions_to_clone_to` - The list of regions that support the creation of an Autonomous Database clone or an Autonomous Data Guard standby database.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `time_created` - The date and time the Autonomous Database was created.
* `time_data_guard_role_changed` - The date and time the Autonomous Data Guard role was switched for the Autonomous Database. For databases that have standbys in both the primary Data Guard region and a remote Data Guard standby region, this is the latest timestamp of either the database using the "primary" role in the primary Data Guard region, or database located in the remote Data Guard standby region.
* `time_deletion_of_free_autonomous_database` - The date and time the Always Free database will be automatically deleted because of inactivity. If the database is in the STOPPED state and without activity until this time, it will be deleted.
* `time_disaster_recovery_role_changed` - The date and time the Disaster Recovery role was switched for the standby Autonomous Database.
* `time_local_data_guard_enabled` - The date and time that Autonomous Data Guard was enabled for an Autonomous Database where the standby was provisioned in the same region as the primary database.
* `time_maintenance_begin` - The date and time when maintenance will begin.
* `time_maintenance_end` - The date and time when maintenance will end.
* `time_of_auto_refresh_start` - The the date and time that auto-refreshing will begin for an Autonomous Database refreshable clone. This value controls only the start time for the first refresh operation. Subsequent (ongoing) refresh operations have start times controlled by the value of the `autoRefreshFrequencyInSeconds` parameter.
* `time_of_last_failover` - The timestamp of the last failover operation.
* `time_of_last_refresh` - The date and time when last refresh happened.
* `time_of_last_refresh_point` - The refresh point timestamp (UTC). The refresh point is the time to which the database was most recently refreshed. Data created after the refresh point is not included in the refresh.
* `time_of_last_switchover` - The timestamp of the last switchover operation for the Autonomous Database.
* `time_of_next_refresh` - The date and time of next refresh.
* `time_reclamation_of_free_autonomous_database` - The date and time the Always Free database will be stopped because of inactivity. If this time is reached without any database activity, the database will automatically be put into the STOPPED state.
* `time_until_reconnect_clone_enabled` - The time and date as an RFC3339 formatted string, e.g., 2022-01-01T12:00:00.000Z, to set the limit for a refreshable clone to be reconnected to its source database.
* `total_backup_storage_size_in_gbs` - The backup storage to the database.
* `used_data_storage_size_in_gbs` - The storage space consumed by Autonomous Database in GBs.
* `used_data_storage_size_in_tbs` - The amount of storage that has been used, in terabytes.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
* `whitelisted_ips` - The client IP access control list (ACL). This feature is available for [Autonomous Database Serverless] (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. If `arePrimaryWhitelistedIpsUsed` is 'TRUE' then Autonomous Database uses this primary's IP access control list (ACL) for the disaster recovery peer called `standbywhitelistedips`.

	For Autonomous Database Serverless instances, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]` For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations. Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`

	For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.

  For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.
## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 12 hours), when creating the Autonomous Database
	* `update` - (Defaults to 12 hours), when updating the Autonomous Database
	* `delete` - (Defaults to 12 hours), when destroying the Autonomous Database


## Import

AutonomousDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_database.test_autonomous_database "id"
```
