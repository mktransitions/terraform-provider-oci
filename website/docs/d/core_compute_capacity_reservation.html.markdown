---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_capacity_reservation"
sidebar_current: "docs-oci-datasource-core-compute_capacity_reservation"
description: |-
  Provides details about a specific Compute Capacity Reservation in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_compute_capacity_reservation
This data source provides details about a specific Compute Capacity Reservation resource in Oracle Cloud Infrastructure Core service.

Gets information about the specified compute capacity reservation.

## Example Usage

```hcl
data "oci_core_compute_capacity_reservation" "test_compute_capacity_reservation" {
	#Required
	capacity_reservation_id = oci_core_capacity_reservation.test_capacity_reservation.id
}
```

## Argument Reference

The following arguments are supported:

* `capacity_reservation_id` - (Required) The OCID of the compute capacity reservation.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the compute capacity reservation.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the compute capacity reservation. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute capacity reservation.
* `instance_reservation_configs` - The capacity configurations for the capacity reservation.

	To use the reservation for the desired shape, specify the shape, count, and optionally the fault domain where you want this configuration. 
	* `cluster_config` - The HPC cluster configuration requested when launching instances in a compute capacity reservation.

		If the parameter is provided, the reservation is created with the HPC island and a list of HPC blocks that you specify. If a list of HPC blocks are missing or not provided, the reservation is created with any HPC blocks in the HPC island that you specify. If the values of HPC island or HPC block that you provide are not valid, an error is returned.
		* `hpc_island_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the HPC island. 
		* `network_block_ids` - The list of OCIDs of the network blocks.
	* `cluster_placement_group_id` - The OCID of the cluster placement group for this instance reservation capacity configuration.
	* `fault_domain` - The fault domain of this capacity configuration. If a value is not supplied, this capacity configuration is applicable to all fault domains in the specified availability domain. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm).
	* `instance_shape` - The shape to use when launching instances using compute capacity reservations. The shape determines the number of CPUs, the amount of memory, and other resources allocated to the instance. You can list all available shapes by calling [ListComputeCapacityReservationInstanceShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/computeCapacityReservationInstanceShapes/ListComputeCapacityReservationInstanceShapes). 
	* `instance_shape_config` - The shape configuration requested when launching instances in a compute capacity reservation.

		If the parameter is provided, the reservation is created with the resources that you specify. If some properties are missing or the parameter is not provided, the reservation is created with the default configuration values for the `shape` that you specify.

		Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned.

		For more information about customizing the resources that are allocated to flexible shapes, see [Flexible Shapes](https://docs.cloud.oracle.com/iaas/Content/Compute/References/computeshapes.htm#flexible). 
		* `memory_in_gbs` - The total amount of memory available to the instance, in gigabytes. 
		* `ocpus` - The total number of OCPUs available to the instance. 
	* `reserved_count` - The total number of instances that can be launched from the capacity configuration.
	* `used_count` - The amount of capacity in use out of the total capacity reserved in this capacity configuration.
* `is_default_reservation` - Whether this capacity reservation is the default. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default). 
* `reserved_instance_count` - The number of instances for which capacity will be held with this compute capacity reservation. This number is the sum of the values of the `reservedCount` fields for all of the instance capacity configurations under this reservation. The purpose of this field is to calculate the percentage usage of the reservation. 
* `state` - The current state of the compute capacity reservation.
* `time_created` - The date and time the compute capacity reservation was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the compute capacity reservation was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `used_instance_count` - The total number of instances currently consuming space in this compute capacity reservation. This number is the sum of the values of the `usedCount` fields for all of the instance capacity configurations under this reservation. The purpose of this field is to calculate the percentage usage of the reservation. 

