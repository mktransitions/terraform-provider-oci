---
subcategory: "Dns"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_rrset"
sidebar_current: "docs-oci-resource-dns-rrset"
description: |-
  Provides the Rrset resource in Oracle Cloud Infrastructure Dns service
---

# oci_dns_rrset
This resource provides the Rrset resource in Oracle Cloud Infrastructure Dns service.

Replaces records in the specified RRSet. RRSet with a `domain` and `rtype` is unique within a zone.

## Example Usage

```hcl
resource "oci_dns_rrset" "test_rrset" {
	#Required
	domain = var.rrset_domain
	rtype = var.rrset_rtype
	zone_name_or_id = oci_dns_zone.test_zone.id

	#Optional
	compartment_id = var.compartment_id
	items {
		#Required
		domain = var.rrset_items_domain
		rdata = var.rrset_items_rdata
		rtype = var.rrset_items_rtype
		ttl = var.rrset_items_ttl
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment the resource belongs to.
* `domain` - (Required) The target fully-qualified domain name (FQDN) within the target zone.
* `items` - (Optional) (Updatable) 
    **NOTE** Omitting `items` at time of create, will delete any existing records in the RRSet
	* `domain` - (Required) The fully qualified domain name where the record can be located. 
	* `rdata` - (Required) (Updatable) The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see [Supported DNS Resource Record Types](https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm)  
	* `rtype` - (Required) The canonical name for the record's type, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
	* `ttl` - (Required) (Updatable) The Time To Live for the record, in seconds.
* `rtype` - (Required) The type of the target RRSet within the target zone.
* `zone_name_or_id` - (Required) The name or OCID of the target zone.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - 
	* `domain` - The fully qualified domain name where the record can be located. 
	* `is_protected` - A Boolean flag indicating whether or not parts of the record are unable to be explicitly managed. 
	* `rdata` - The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see [Supported DNS Resource Record Types](https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm) 
	* `record_hash` - A unique identifier for the record within its zone. 
	* `rrset_version` - The latest version of the record's zone in which its RRSet differs from the preceding version. 
	* `rtype` - The canonical name for the record's type, such as A or CNAME. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
	* `ttl` - The Time To Live for the record, in seconds.

## Import

Rrsets can be imported using the `id`, e.g.

```
$ terraform import oci_dns_rrset.test_rrset "zoneNameOrId/{zoneNameOrId}/domain/{domain}/rtype/{rtype}" 
```


