# oci\_core\_subnets

**API:** [Subnet Reference][c512efd1]

  [c512efd1]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Subnet/ "SubnetReference"

Gets a list of subnets.

## Example Usage

```
    data "oci_core_subnets" "s" {
      compartment_id = "compartmentid"
      vcn_id = "vcnid"
    }
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `vcn_id` - (Required) The OCID of the VCN.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) Length of the snapshot data to retrieve.

## Attributes Reference

The following attributes are exported:

* `subnets` - The list of subnets.

## Subnet reference
* `availability_domain` - The subnet's Availability Domain.
* `cidr_block` - The CIDR IP address block of the VCN.
* `compartment_id` - The OCID of the compartment containing the VCN.
* `dhcp_options_id` - The OCID for the VCN's default set of DHCP options.
* `route_table_id` - The OCID for the VCN's default route table.
* `security_list_ids` - OCIDs for the security lists to use for VNICs in this subnet.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The subnet's Oracle ID (OCID).
* `prohibit_public_ip_on_vnic` - Whether VNICs within this subnet can have public IP addresses. Example: `false`
* `vcn_id` - The OCID of the VCN the subnet is in.
* `state` - The VCN's current state. Allowed values are: [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `time_created` - The date and time the subnet was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `virtual_router_ip` - The IP address of the virtual router.
* `virtual_router_mac` - The MAC address of the virtual router.
