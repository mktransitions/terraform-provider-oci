---
subcategory: "Network Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_load_balancer_backend"
sidebar_current: "docs-oci-resource-network_load_balancer-backend"
description: |-
  Provides the Backend resource in Oracle Cloud Infrastructure Network Load Balancer service
---

# oci_network_load_balancer_backend
This resource provides the Backend resource in Oracle Cloud Infrastructure Network Load Balancer service.

Adds a backend server to a backend set.

## Example Usage

```hcl
resource "oci_network_load_balancer_backend" "test_backend" {
	#Required
	backend_set_name = oci_network_load_balancer_backend_set.test_backend_set.name
	network_load_balancer_id = oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id
	port = var.backend_port

	#Optional
	ip_address = var.backend_ip_address
	is_backup = var.backend_is_backup
	is_drain = var.backend_is_drain
	is_offline = var.backend_is_offline
	name = var.backend_name
	target_id = oci_cloud_guard_target.test_target.id
	weight = var.backend_weight
}
```

## Argument Reference

The following arguments are supported:

* `backend_set_name` - (Required) The name of the backend set to which to add the backend server.  Example: `example_backend_set` 
* `name` - (Required) Unique name identifying the backend within the backend set. Example: `webServer1` 
* `ip_address` - (Optional) The IP address of the backend server. Example: `10.0.0.3` 
* `is_backup` - (Optional) (Updatable) Whether the network load balancer should treat this server as a backup unit. If `true`, then the network load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "isBackup" fail the health check policy.  Example: `false` 
* `is_drain` - (Optional) (Updatable) Whether the network load balancer should drain this server. Servers marked "isDrain" receive no incoming traffic.  Example: `false` 
* `is_offline` - (Optional) (Updatable) Whether the network load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
* `network_load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
* `port` - (Required) The communication port for the backend server.  Example: `8080` 
* `target_id` - (Optional) The IP OCID/Instance OCID associated with the backend server. Example: `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>` 
* `weight` - (Optional) (Updatable) The network load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives three times the number of new connections as a server weighted '1'. For more information about load balancing policies, see [How Network Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/NetworkLoadBalancer/introducton.htm#Policies).  Example: `3` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `ip_address` - The IP address of the backend server. Example: `10.0.0.3` 
* `is_backup` - Whether the network load balancer should treat this server as a backup unit. If `true`, then the network load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "isBackup" fail the health check policy.  Example: `false` 
* `is_drain` - Whether the network load balancer should drain this server. Servers marked "isDrain" receive no incoming traffic.  Example: `false` 
* `is_offline` - Whether the network load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
* `name` - A read-only field showing the IP address/IP OCID and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080`, or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:443` or `10.0.0.3:0` 
* `port` - The communication port for the backend server.  Example: `8080` 
* `target_id` - The IP OCID/Instance OCID associated with the backend server. Example: `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>` 
* `weight` - The network load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives three times the number of new connections as a server weighted '1'. For more information about load balancing policies, see [How Network Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/NetworkLoadBalancer/introducton.htm#Policies).  Example: `3` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Backend
	* `update` - (Defaults to 20 minutes), when updating the Backend
	* `delete` - (Defaults to 20 minutes), when destroying the Backend


## Import

Backends can be imported using the `id`, e.g.

```
$ terraform import oci_network_load_balancer_backend.test_backend "networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}/backends/{backendName}" 
```

