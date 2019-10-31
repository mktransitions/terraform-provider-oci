---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_network"
sidebar_current: "docs-oci-resource-database-vm_cluster_network"
description: |-
  Provides the Vm Cluster Network resource in Oracle Cloud Infrastructure Database service
---

# oci_database_vm_cluster_network
This resource provides the Vm Cluster Network resource in Oracle Cloud Infrastructure Database service.

Creates the VM cluster network.


## Example Usage

```hcl
resource "oci_database_vm_cluster_network" "test_vm_cluster_network" {
	#Required
	compartment_id = "${var.compartment_id}"
	display_name = "${var.vm_cluster_network_display_name}"
	exadata_infrastructure_id = "${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}"
	scans {
		#Required
		hostname = "${var.vm_cluster_network_scans_hostname}"
		ips = "${var.vm_cluster_network_scans_ips}"
		port = "${var.vm_cluster_network_scans_port}"
	}
	vm_networks {
		#Required
		domain_name = "${var.vm_cluster_network_vm_networks_domain_name}"
		gateway = "${var.vm_cluster_network_vm_networks_gateway}"
		netmask = "${var.vm_cluster_network_vm_networks_netmask}"
		network_type = "${var.vm_cluster_network_vm_networks_network_type}"
		nodes {
			#Required
			hostname = "${var.vm_cluster_network_vm_networks_nodes_hostname}"
			ip = "${var.vm_cluster_network_vm_networks_nodes_ip}"

			#Optional
			vip = "${var.vm_cluster_network_vm_networks_nodes_vip}"
			vip_hostname = "${var.vm_cluster_network_vm_networks_nodes_vip_hostname}"
		}
		vlan_id = "${var.vm_cluster_network_vm_networks_vlan_id}"
	}

	#Optional
	defined_tags = "${var.vm_cluster_network_defined_tags}"
	dns = "${var.vm_cluster_network_dns}"
	freeform_tags = {"Department"= "Finance"}
	ntp = "${var.vm_cluster_network_ntp}"
	validate_vm_cluster_network = "${var.vm_cluster_network_validate_vm_cluster_network}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) The user-friendly name for the VM cluster network. The name does not need to be unique.
* `dns` - (Optional) (Updatable) The list of DNS server IP addresses. Maximum of 3 allowed.
* `exadata_infrastructure_id` - (Required) The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `ntp` - (Optional) (Updatable) The list of NTP server IP addresses. Maximum of 3 allowed.
* `scans` - (Required) (Updatable) The SCAN details.
	* `hostname` - (Required) (Updatable) The SCAN hostname.
	* `ips` - (Required) (Updatable) The list of SCAN IP addresses. Three addresses should be provided.
	* `port` - (Required) (Updatable) The SCAN port. Default is 1521.
* `vm_networks` - (Required) (Updatable) Details of the client and backup networks.
	* `domain_name` - (Required) (Updatable) The network domain name.
	* `gateway` - (Required) (Updatable) The network gateway.
	* `netmask` - (Required) (Updatable) The network netmask.
	* `network_type` - (Required) (Updatable) The network type.
	* `nodes` - (Required) (Updatable) The list of node details.
		* `hostname` - (Required) (Updatable) The node host name.
		* `ip` - (Required) (Updatable) The node IP address.
		* `vip` - (Optional) (Updatable) The node virtual IP (VIP) address.
		* `vip_hostname` - (Optional) (Updatable) The node virtual IP (VIP) host name.
	* `vlan_id` - (Required) (Updatable) The network VLAN ID.
* `validate_vm_cluster_network` - A boolean flag indicating whether or not to validate VM cluster network after creation. Updates are not allowed on validated exadata VM cluster network.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the VM cluster network. The name does not need to be unique.
* `dns` - The list of DNS server IP addresses. Maximum of 3 allowed.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `ntp` - The list of NTP server IP addresses. Maximum of 3 allowed.
* `scans` - The SCAN details.
	* `hostname` - The SCAN hostname.
	* `ips` - The list of SCAN IP addresses. Three addresses should be provided.
	* `port` - The SCAN port. Default is 1521.
* `state` - The current state of the VM cluster network.
* `time_created` - The date and time when the VM cluster network was created.
* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated VM Cluster.
* `vm_networks` - Details of the client and backup networks.
	* `domain_name` - The network domain name.
	* `gateway` - The network gateway.
	* `netmask` - The network netmask.
	* `network_type` - The network type.
	* `nodes` - The list of node details.
		* `hostname` - The node host name.
		* `ip` - The node IP address.
		* `vip` - The node virtual IP (VIP) address.
		* `vip_hostname` - The node virtual IP (VIP) host name.
	* `vlan_id` - The network VLAN ID.

## Import

Import is not supported for this resource.

