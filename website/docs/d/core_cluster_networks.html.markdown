---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cluster_networks"
sidebar_current: "docs-oci-datasource-core-cluster_networks"
description: |-
  Provides the list of Cluster Networks in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_cluster_networks
This data source provides the list of Cluster Networks in Oracle Cloud Infrastructure Core service.

Lists the cluster networks in the specified compartment.

## Example Usage

```hcl
data "oci_core_cluster_networks" "test_cluster_networks" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.cluster_network_display_name
	state = var.cluster_network_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `cluster_networks` - The list of cluster_networks.

### ClusterNetwork Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the cluster netowrk. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the cluster network. Does not have to be unique, and it's changeable. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster network.
* `instance_pools` - The instance pools in the cluster network.

	Each cluster network can have one instance pool. 
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the instance pool. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - The user-friendly name.  Does not have to be unique.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.
	* `instance_configuration_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance configuration associated with the instance pool. 
	* `load_balancers` - The load balancers attached to the instance pool. 
		* `backend_set_name` - The name of the backend set on the load balancer.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer attachment.
		* `instance_pool_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool of the load balancer attachment. 
		* `load_balancer_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer attached to the instance pool. 
		* `port` - The port value used for the backends.
		* `state` - The status of the interaction between the instance pool and the load balancer.
		* `vnic_selection` - Indicates which VNIC on each instance in the instance pool should be used to associate with the load balancer. Possible values are "PrimaryVnic" or the displayName of one of the secondary VNICs on the instance configuration that is associated with the instance pool.
	* `placement_configurations` - The placement configurations for the instance pool.
		* `availability_domain` - The availability domain to place instances.  Example: `Uocm:PHX-AD-1` 
		* `fault_domains` - The fault domains to place instances.

			If you don't provide any values, the system makes a best effort to distribute instances across all fault domains based on capacity.

			To distribute the instances evenly across selected fault domains, provide a set of fault domains. For example, you might want instances to be evenly distributed if your applications require high availability.

			To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

			Example: `[FAULT-DOMAIN-1, FAULT-DOMAIN-2, FAULT-DOMAIN-3]` 
		* `primary_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the primary subnet to place instances. 
		* `secondary_vnic_subnets` - The set of secondary VNIC data for instances in the pool.
			* `display_name` - The display name of the VNIC. This is also use to match against the instance configuration defined secondary VNIC. 
			* `subnet_id` - The subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
	* `size` - The number of instances that should be in the instance pool.
	* `state` - The current state of the instance pool.
	* `time_created` - The date and time the instance pool was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `placement_configuration` - The placement configuration for the instance pools in the cluster network. 
	* `availability_domain` - The availability domain to place instances.  Example: `Uocm:PHX-AD-1` 
	* `primary_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the primary subnet to place instances. 
	* `secondary_vnic_subnets` - The set of secondary VNIC data for instances in the pool.
		* `display_name` - The display name of the VNIC. This is also use to match against the instance configuration defined secondary VNIC. 
		* `subnet_id` - The subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
* `state` - The current state of the cluster network.
* `time_created` - The date and time the resource was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the resource was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

