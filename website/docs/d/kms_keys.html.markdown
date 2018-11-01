---
layout: "oci"
page_title: "OCI: oci_kms_keys"
sidebar_current: "docs-oci-datasource-kms-keys"
description: |-
  Provides a list of Keys
---

# Data Source: oci_kms_keys
The `oci_kms_keys` data source allows access to the list of OCI keys

Lists the keys in the specified vault and compartment.


## Example Usage

```hcl
data "oci_kms_keys" "test_keys" {
	#Required
	compartment_id = "${var.compartment_id}"
	management_endpoint = "${var.key_management_endpoint}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `management_endpoint` - (Required) The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.


## Attributes Reference

The following attributes are exported:

* `keys` - The list of keys.

### Key Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains this key.
* `current_key_version` - The OCID of the KeyVersion resource used in cryptographic operations. During key rotation, service may be in transitional state where this or a newer KeyVersion are used intermittently, and currentKeyVersion field is updated once service is guaranteed to use new KeyVersion for all consequent encrypt operations. 
* `display_name` - A user-friendly name for the key. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `id` - The OCID of the key.
* `key_shape` - 
	* `algorithm` - The algorithm used by a key's KeyVersions to encrypt or decrypt.
	* `length` - The length of the key, expressed as an integer. Values of 16, 24, or 32 are supported. 
* `state` - The key's current state.  Example: `ENABLED` 
* `time_created` - The date and time the key was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z` 
* `vault_id` - The OCID of the vault that contains this key.

