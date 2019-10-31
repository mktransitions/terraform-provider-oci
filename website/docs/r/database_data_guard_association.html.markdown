---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_data_guard_association"
sidebar_current: "docs-oci-resource-database-data_guard_association"
description: |-
  Provides the Data Guard Association resource in Oracle Cloud Infrastructure Database service
---

# oci_database_data_guard_association
This resource provides the Data Guard Association resource in Oracle Cloud Infrastructure Database service.

Creates a new Data Guard association.  A Data Guard association represents the replication relationship between the
specified database and a peer database. For more information, see [Using Oracle Data Guard](https://docs.cloud.oracle.com/iaas/Content/Database/Tasks/usingdataguard.htm).

All Oracle Cloud Infrastructure resources, including Data Guard associations, get an Oracle-assigned, unique ID
called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response.
You can also retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the
resource in the Console. For more information, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
resource "oci_database_data_guard_association" "test_data_guard_association" {
	#Required
	creation_type = "${var.data_guard_association_creation_type}"
	database_admin_password = "${var.data_guard_association_database_admin_password}"
	database_id = "${oci_database_database.test_database.id}"
	protection_mode = "${var.data_guard_association_protection_mode}"
	transport_type = "${var.data_guard_association_transport_type}"

	#Optional
	availability_domain = "${var.data_guard_association_availability_domain}"
	backup_network_nsg_ids = "${var.data_guard_association_backup_network_nsg_ids}"
	display_name = "${var.data_guard_association_display_name}"
	hostname = "${var.data_guard_association_hostname}"
	nsg_ids = "${var.data_guard_association_nsg_ids}"
	peer_db_system_id = "${oci_database_db_system.test_db_system.id}"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Applicable when creation_type=NewDbSystem) The name of the availability domain that the standby database DB system will be located in. For example- "Uocm:PHX-AD-1".
* `backup_network_nsg_ids` - (Applicable when creation_type=NewDbSystem) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata DB systems. 
* `creation_type` - (Required) Specifies whether to create the peer database in an existing DB system or in a new DB system. `ExistingDbSystem` is not supported for creating Data Guard associations for virtual machine DB system databases. 
* `database_admin_password` - (Required) A strong password for the `SYS`, `SYSTEM`, and `PDB Admin` users to apply during standby creation.

	The password must contain no fewer than nine characters and include:
	* At least two uppercase characters.
	* At least two lowercase characters.
	* At least two numeric characters.
	* At least two special characters. Valid special characters include "_", "#", and "-" only.

	**The password MUST be the same as the primary admin password.** 
* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `delete_standby_db_home_on_delete` - (Required) (Updatable) if set to true the destroy operation will destroy the standby dbHome/dbSystem that is referenced in the Data Guard Association. The Data Guard Association gets destroyed when standby dbHome/dbSystem is terminated. Only `true` is supported at this time. If you change an argument that is used during the delete operation you must run `terraform apply` first so that that the change in the value is registered in the statefile before running `terraform destroy`. `terraform destroy` only looks at what is currently on the statefile and ignores the terraform configuration files. 
* `display_name` - (Applicable when creation_type=NewDbSystem) The user-friendly name of the DB system that will contain the the standby database. The display name does not have to be unique.
* `hostname` - (Applicable when creation_type=NewDbSystem) The hostname for the DB node.
* `nsg_ids` - (Applicable when creation_type=NewDbSystem) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
* `peer_db_system_id` - (Applicable when creation_type=ExistingDbSystem) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system in which to create the standby database. You must supply this value if creationType is `ExistingDbSystem`. 
* `protection_mode` - (Required) The protection mode to set up between the primary and standby databases. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation.

	**IMPORTANT** - The only protection mode currently supported by the Database service is MAXIMUM_PERFORMANCE. 
* `subnet_id` - (Applicable when creation_type=NewDbSystem) The OCID of the subnet the DB system is associated with. **Subnet Restrictions:**
	* For 1- and 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.16.16/28

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `transport_type` - (Required) The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
	* MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
	* MAXIMUM_PERFORMANCE - ASYNC
	* MAXIMUM_PROTECTION - SYNC

	For more information, see [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400) in the Oracle Data Guard documentation.

	**IMPORTANT** - The only transport type currently supported by the Database service is ASYNC. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `apply_lag` - The lag time between updates to the primary database and application of the redo data on the standby database, as computed by the reporting database.  Example: `9 seconds` 
* `apply_rate` - The rate at which redo logs are synced between the associated databases.  Example: `180 Mb per second` 
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the reporting database.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Data Guard association.
* `lifecycle_details` - Additional information about the current lifecycleState, if available. 
* `peer_data_guard_association_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer database's Data Guard association.
* `peer_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated peer database.
* `peer_db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database home containing the associated peer database. 
* `peer_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system containing the associated peer database. 
* `peer_role` - The role of the peer database in this Data Guard association.
* `protection_mode` - The protection mode of this Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `role` - The role of the reporting database in this Data Guard association.
* `state` - The current state of the Data Guard association.
* `time_created` - The date and time the Data Guard association was created.
* `transport_type` - The redo transport type used by this Data Guard association.  For more information, see [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400) in the Oracle Data Guard documentation. 

## Import

Import is not supported for this resource.

