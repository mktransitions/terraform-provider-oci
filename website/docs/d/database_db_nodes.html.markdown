---
layout: "oci"
page_title: "OCI: oci_database_db_nodes"
sidebar_current: "docs-oci-datasource-database-db_nodes"
description: |-
  Provides a list of DbNodes
---

# Data Source: oci_database_db_nodes
The `oci_database_db_nodes` data source allows access to the list of OCI db_nodes

Gets a list of database nodes in the specified DB system and compartment. A database node is a server running database software.


## Example Usage

```hcl
data "oci_database_db_nodes" "test_db_nodes" {
	#Required
	compartment_id = "${var.compartment_id}"
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DB system.


## Attributes Reference

The following attributes are exported:

* `db_nodes` - The list of db_nodes.

### DbNode Reference

The following attributes are exported:

* `backup_vnic_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the backup VNIC.
* `db_system_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DB system.
* `hostname` - The host name for the database node.
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the database node.
* `software_storage_size_in_gb` - The size (in GB) of the block storage volume allocation for the DB system. This attribute applies only for virtual machine DB systems. 
* `state` - The current state of the database node.
* `time_created` - The date and time that the database node was created.
* `vnic_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the VNIC.

