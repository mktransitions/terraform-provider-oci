---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_swift_passwords"
sidebar_current: "docs-oci-datasource-identity-swift_passwords"
description: |-
  Provides the list of Swift Passwords in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_swift_passwords
This data source provides the list of Swift Passwords in Oracle Cloud Infrastructure Identity service.

**Deprecated. Use [ListAuthTokens](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/AuthToken/ListAuthTokens) instead.**

Lists the Swift passwords for the specified user. The returned object contains the password's OCID, but not
the password itself. The actual password is returned only upon creation.


## Example Usage

```hcl
data "oci_identity_swift_passwords" "test_swift_passwords" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.


## Attributes Reference

The following attributes are exported:

* `passwords` - The list of passwords.

### SwiftPassword Reference

The following attributes are exported:

* `description` - The description you assign to the Swift password. Does not have to be unique, and it's changeable.
* `expires_on` - Date and time when this password will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `id` - The OCID of the Swift password.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `password` - The Swift password. The value is available only in the response for `CreateSwiftPassword`, and not for `ListSwiftPasswords` or `UpdateSwiftPassword`. 
* `state` - The password's current state.
* `time_created` - Date and time the `SwiftPassword` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the password belongs to.

