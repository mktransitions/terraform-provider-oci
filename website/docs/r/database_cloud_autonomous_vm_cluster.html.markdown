---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_autonomous_vm_cluster"
sidebar_current: "docs-oci-resource-database-cloud_autonomous_vm_cluster"
description: |-
  Provides the Cloud Autonomous Vm Cluster resource in Oracle Cloud Infrastructure Database service
---

# oci_database_cloud_autonomous_vm_cluster
This resource provides the Cloud Autonomous Vm Cluster resource in Oracle Cloud Infrastructure Database service.

Creates an Autonomous Exadata VM cluster in the Oracle cloud. For Exadata Cloud@Customer systems, see [CreateAutonomousVmCluster](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/AutonomousVmCluster/CreateAutonomousVmCluster).


## Example Usage

```hcl
resource "oci_database_cloud_autonomous_vm_cluster" "test_cloud_autonomous_vm_cluster" {
	#Required
	cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id
	compartment_id = var.compartment_id
	display_name = var.cloud_autonomous_vm_cluster_display_name
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	autonomous_data_storage_size_in_tbs = var.cloud_autonomous_vm_cluster_autonomous_data_storage_size_in_tbs
	cluster_time_zone = var.cloud_autonomous_vm_cluster_cluster_time_zone
	compute_model = var.cloud_autonomous_vm_cluster_compute_model
	cpu_core_count_per_node = var.cloud_autonomous_vm_cluster_cpu_core_count_per_node
	db_servers = var.cloud_autonomous_vm_cluster_db_servers
	defined_tags = var.cloud_autonomous_vm_cluster_defined_tags
	description = var.cloud_autonomous_vm_cluster_description
	freeform_tags = {"Department"= "Finance"}
	license_model = var.cloud_autonomous_vm_cluster_license_model
	maintenance_window_details {

		#Optional
		custom_action_timeout_in_mins = var.cloud_autonomous_vm_cluster_maintenance_window_details_custom_action_timeout_in_mins
		days_of_week {
			#Required
			name = var.cloud_autonomous_vm_cluster_maintenance_window_details_days_of_week_name
		}
		hours_of_day = var.cloud_autonomous_vm_cluster_maintenance_window_details_hours_of_day
		is_custom_action_timeout_enabled = var.cloud_autonomous_vm_cluster_maintenance_window_details_is_custom_action_timeout_enabled
		is_monthly_patching_enabled = var.cloud_autonomous_vm_cluster_maintenance_window_details_is_monthly_patching_enabled
		lead_time_in_weeks = var.cloud_autonomous_vm_cluster_maintenance_window_details_lead_time_in_weeks
		months {
			#Required
			name = var.cloud_autonomous_vm_cluster_maintenance_window_details_months_name
		}
		patching_mode = var.cloud_autonomous_vm_cluster_maintenance_window_details_patching_mode
		preference = var.cloud_autonomous_vm_cluster_maintenance_window_details_preference
		weeks_of_month = var.cloud_autonomous_vm_cluster_maintenance_window_details_weeks_of_month
	}
	memory_per_oracle_compute_unit_in_gbs = var.cloud_autonomous_vm_cluster_memory_per_oracle_compute_unit_in_gbs
	nsg_ids = var.cloud_autonomous_vm_cluster_nsg_ids
	total_container_databases = var.cloud_autonomous_vm_cluster_total_container_databases
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_data_storage_size_in_tbs` - (Optional) The data disk group size to be allocated for Autonomous Databases, in TBs.
* `cloud_exadata_infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
* `cluster_time_zone` - (Optional) The time zone to use for the Cloud Autonomous VM cluster. For details, see [DB System Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpu_core_count_per_node` - (Optional) The number of OCPU cores to be enabled per VM cluster node.
* `db_servers` - (Optional) The list of database servers.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `compute_model` - (Optional) The compute model of the Cloud Autonomous VM Cluster.
* `description` - (Optional) (Updatable) User defined description of the cloud Autonomous VM cluster.
* `display_name` - (Required) (Updatable) The user-friendly name for the cloud Autonomous VM cluster. The name does not need to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud. License Included allows you to subscribe to new Oracle Database software licenses and the Database service. Note that when provisioning an Autonomous Database on [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null because the attribute is already set at the Autonomous Exadata Infrastructure level. When using [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`. 
* `maintenance_window_details` - (Optional) (Updatable) The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
    * `custom_action_timeout_in_mins` - (Optional) (Updatable) Determines the amount of time the system will wait before the start of each database server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
    * `days_of_week` - (Optional) (Updatable) Days during the week when maintenance should be performed.
        * `name` - (Required) (Updatable) Name of the day of the week.
    * `hours_of_day` - (Optional) (Updatable) The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
        * 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
    * `is_custom_action_timeout_enabled` - (Optional) (Updatable) If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
    * `is_monthly_patching_enabled` - (Optional) (Updatable) If true, enables the monthly patching option.
    * `lead_time_in_weeks` - (Optional) (Updatable) Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
    * `months` - (Optional) (Updatable) Months during the year when maintenance should be performed.
        * `name` - (Required) (Updatable) Name of the month of the year.
    * `patching_mode` - (Optional) (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `preference` - (Optional) (Updatable) The maintenance window scheduling preference.
	* `weeks_of_month` - (Optional) (Updatable) Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `memory_per_oracle_compute_unit_in_gbs` - (Optional) The amount of memory (in GBs) to be enabled per each OCPU core.
* `nsg_ids` - (Optional) (Updatable) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
    * A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty. 
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the cloud Autonomous VM Cluster is associated with. 
* `total_container_databases` - (Optional) The total number of Autonomous Container Databases that can be created.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `autonomous_data_storage_size_in_tbs` - The data disk group size allocated for Autonomous Databases, in TBs.
* `availability_domain` - The name of the availability domain that the cloud Autonomous VM cluster is located in.
* `available_autonomous_data_storage_size_in_tbs` - The data disk group size available for Autonomous Databases, in TBs.
* `available_container_databases` - The number of Autonomous Container Databases that can be created with the currently available local storage.
* `available_cpus` - CPU cores available for allocation to Autonomous Databases.
* `cloud_exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
* `cluster_time_zone` - The time zone of the Cloud Autonomous VM Cluster.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_model` - The compute model of the Cloud Autonomous VM Cluster.
* `cpu_core_count` - The number of CPU cores on the cloud Autonomous VM cluster.
* `cpu_core_count_per_node` - The number of OCPU cores enabled per VM cluster node.
* `data_storage_size_in_gb` - The total data storage allocated, in gigabytes (GB).
* `data_storage_size_in_tbs` - The total data storage allocated, in terabytes (TB).
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `description` - User defined description of the cloud Autonomous VM cluster.
* `display_name` - The user-friendly name for the cloud Autonomous VM cluster. The name does not need to be unique.
* `domain` - The domain name for the cloud Autonomous VM cluster.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `hostname` - The hostname for the cloud Autonomous VM cluster.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud Autonomous VM cluster.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `last_update_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance update history. This value is updated when a maintenance update starts.
* `license_model` - The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud. License Included allows you to subscribe to new Oracle Database software licenses and the Database service. Note that when provisioning an Autonomous Database on [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null because the attribute is already set at the Autonomous Exadata Infrastructure level. When using [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `maintenance_window` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `custom_action_timeout_in_mins` - Determines the amount of time the system will wait before the start of each database server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `is_custom_action_timeout_enabled` - If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
	* `is_monthly_patching_enabled` - If true, enables the monthly patching option.
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `patching_mode` - Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `preference` - The maintenance window scheduling preference.
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `memory_per_oracle_compute_unit_in_gbs` - The amount of memory (in GBs) enabled per each OCPU core.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `node_count` - The number of database servers in the cloud VM cluster.
* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty. 
* `ocpu_count` - The number of CPU cores on the cloud Autonomous VM cluster. Only 1 decimal place is allowed for the fractional part.
* `reclaimable_cpus` - CPU cores that continue to be included in the count of OCPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available OCPUs at its parent AVMC level by restarting the Autonomous Container Database.
* `shape` - The model name of the Exadata hardware running the cloud Autonomous VM cluster.
* `state` - The current state of the cloud Autonomous VM cluster.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the cloud Autonomous VM Cluster is associated with.

  **Subnet Restrictions:**
	* For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.

  These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet.
* `time_created` - The date and time that the cloud Autonomous VM cluster was created.
* `time_updated` - The last date and time that the cloud Autonomous VM cluster was updated.
* `total_container_databases` - The total number of Autonomous Container Databases that can be created with the allocated local storage.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cloud Autonomous Vm Cluster
	* `update` - (Defaults to 20 minutes), when updating the Cloud Autonomous Vm Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Cloud Autonomous Vm Cluster


## Import

CloudAutonomousVmClusters can be imported using the `id`, e.g.

```
$ terraform import oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster "id"
```
