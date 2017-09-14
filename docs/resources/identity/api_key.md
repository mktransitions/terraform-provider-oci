# oci\_identity\_api\_key

Provides an api key resource.

## Example Usage

```
resource "oci_identity_api_key" "t" {
    user_id = "user_id"
    key_value = "1"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.
* `key_value` - (Required) The public key. Must be an RSA key in PEM format.

## Attributes Reference
* `key_id` - An Oracle-assigned identifier for the key, in this format: TENANCY_OCID/USER_OCID/KEY_FINGERPRINT.
* `key_value` - The key's value.
* `fingerprint` - The key's fingerprint (e.g., 12:34:56:78:90:ab:cd:ef:12:34:56:78:90:ab:cd:ef).
* `user_id` - The OCID of the user the key belongs to.
* `time_created` - Date and time the ApiKey was created.
* `state` - The compartment's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
