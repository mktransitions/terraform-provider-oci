---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_key_version"
sidebar_current: "docs-oci-datasource-kms-key_version"
description: |-
  Provides details about a specific Key Version in Oracle Cloud Infrastructure Kms service
---

# Data Source: oci_kms_key_version
This data source provides details about a specific Key Version resource in Oracle Cloud Infrastructure Kms service.

Gets information about the specified key version.


## Example Usage

```hcl
data "oci_kms_key_version" "test_key_version" {
	#Required
	key_id = "${oci_kms_key.test_key.id}"
	key_version_id = "${oci_kms_key_version.test_key_version.id}"
	management_endpoint = "${var.key_version_management_endpoint}"
}
```

## Argument Reference

The following arguments are supported:

* `key_id` - (Required) The OCID of the key.
* `key_version_id` - (Required) The OCID of the key version.
* `management_endpoint` - (Required) The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains this key version.
* `key_version_id` - The OCID of the key version.
* `key_id` - The OCID of the key associated with this key version.
* `time_created` - The date and time this key version was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: "2018-04-03T21:10:29.600Z" 
* `vault_id` - The OCID of the vault that contains this key version.

