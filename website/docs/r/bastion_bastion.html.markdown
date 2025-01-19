---
subcategory: "Bastion"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bastion_bastion"
sidebar_current: "docs-oci-resource-bastion-bastion"
description: |-
  Provides the Bastion resource in Oracle Cloud Infrastructure Bastion service
---

# oci_bastion_bastion
This resource provides the Bastion resource in Oracle Cloud Infrastructure Bastion service.

Creates a new bastion. A bastion provides secured, public access to target resources in the cloud that you cannot otherwise reach from the internet. A bastion resides in a public subnet and establishes the network infrastructure needed to connect a user to a target resource in a private subnet.


## Example Usage

```hcl
resource "oci_bastion_bastion" "test_bastion" {
	#Required
	bastion_type = var.bastion_bastion_type
	compartment_id = var.compartment_id
	target_subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	client_cidr_block_allow_list = var.bastion_client_cidr_block_allow_list
	defined_tags = {"foo-namespace.bar-key"= "value"}
	dns_proxy_status = var.bastion_dns_proxy_status
	freeform_tags = {"bar-key"= "value"}
	max_session_ttl_in_seconds = var.bastion_max_session_ttl_in_seconds
	name = var.bastion_name
	phone_book_entry = var.bastion_phone_book_entry
	static_jump_host_ip_addresses = var.bastion_static_jump_host_ip_addresses
}
```

## Argument Reference

The following arguments are supported:

* `bastion_type` - (Required) The type of bastion. Use `standard`. 
* `client_cidr_block_allow_list` - (Required) (Updatable) A list of address ranges in CIDR notation that you want to allow to connect to sessions hosted by this bastion.
* `compartment_id` - (Required) (Updatable) The unique identifier (OCID) of the compartment where the bastion is located.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `dns_proxy_status` - (Optional) Flag to enable FQDN and SOCKS5 Proxy Support. Example: `ENABLED`, `DISABLED`
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `max_session_ttl_in_seconds` - (Optional) (Updatable) The maximum amount of time that any session on the bastion can remain active.
* `name` - (Optional) The name of the bastion, which can't be changed after creation.
* `phone_book_entry` - (Optional) The phonebook entry of the customer's team, which can't be changed after creation. Not applicable to `standard` bastions. 
* `static_jump_host_ip_addresses` - (Optional) (Updatable) A list of IP addresses of the hosts that the bastion has access to. Not applicable to `standard` bastions. 
* `target_subnet_id` - (Required) The unique identifier (OCID) of the subnet that the bastion connects to.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bastion_type` - The type of bastion.
* `client_cidr_block_allow_list` - A list of address ranges in CIDR notation that you want to allow to connect to sessions hosted by this bastion.
* `compartment_id` - The unique identifier (OCID) of the compartment where the bastion is located.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `dns_proxy_status` - Flag to enable FQDN and SOCKS5 Proxy Support. Example: `ENABLED`, `DISABLED`
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The unique identifier (OCID) of the bastion, which can't be changed after creation.
* `lifecycle_details` - A message describing the current state in more detail.
* `max_session_ttl_in_seconds` - The maximum amount of time that any session on the bastion can remain active.
* `max_sessions_allowed` - The maximum number of active sessions allowed on the bastion.
* `name` - The name of the bastion, which can't be changed after creation.
* `phone_book_entry` - The phonebook entry of the customer's team, which can't be changed after creation. Not applicable to `standard` bastions. 
* `private_endpoint_ip_address` - The private IP address of the created private endpoint.
* `state` - The current state of the bastion.
* `static_jump_host_ip_addresses` - A list of IP addresses of the hosts that the bastion has access to. Not applicable to `standard` bastions. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_subnet_id` - The unique identifier (OCID) of the subnet that the bastion connects to.
* `target_vcn_id` - The unique identifier (OCID) of the virtual cloud network (VCN) that the bastion connects to.
* `time_created` - The time the bastion was created. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2020-01-25T21:10:29.600Z` 
* `time_updated` - The time the bastion was updated. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2020-01-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bastion
	* `update` - (Defaults to 20 minutes), when updating the Bastion
	* `delete` - (Defaults to 20 minutes), when destroying the Bastion


## Import

Bastions can be imported using the `id`, e.g.

```
$ terraform import oci_bastion_bastion.test_bastion "id"
```

