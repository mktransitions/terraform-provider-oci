# oci\_core\_security\_lists

Gets a list of security lists

## Example Usage

```
    data "oci_core_security_lists" "t" {
      compartment_id = "compartment_id"
      limit = 1
      page = "page"
      vcn_id = "vcn_id"
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

* `security_lists` - The list of security lists.

## Security List reference
* `compartment_id` - The OCID of the compartment containing the security list.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `egress_security_rules` - Rules for allowing egress IP packets.
* `id` - The security list's Oracle Cloud ID (OCID).
* `ingress_security_rules` - Rules for allowing ingress IP packets.
* `state` - The security list's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `time_created` - The date and time the security list was created.
* `vcn_id` - The OCID of the VCN the security list belongs to.
