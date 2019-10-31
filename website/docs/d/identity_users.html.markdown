---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_users"
sidebar_current: "docs-oci-datasource-identity-users"
description: |-
  Provides the list of Users in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_users
This data source provides the list of Users in Oracle Cloud Infrastructure Identity service.

Lists the users in your tenancy. You must specify your tenancy's OCID as the value for the
compartment ID (remember that the tenancy is simply the root compartment).
See [Where to Get the Tenancy's OCID and User's OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm#five).


## Example Usage

```hcl
data "oci_identity_users" "test_users" {
	#Required
	compartment_id = "${var.tenancy_ocid}"

	#Optional
	external_identifier = "${var.user_external_identifier}"
	identity_provider_id = "${oci_identity_identity_provider.test_identity_provider.id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 
* `external_identifier` - (Optional) The id of a user in the identity provider. 
* `identity_provider_id` - (Optional) The id of the identity provider. 


## Attributes Reference

The following attributes are exported:

* `users` - The list of users.

### User Reference

The following attributes are exported:

* `capabilities` - Properties indicating how the user is allowed to authenticate.
	* `can_use_api_keys` - Indicates if the user can use API keys.
	* `can_use_auth_tokens` - Indicates if the user can use SWIFT passwords / auth tokens.
	* `can_use_console_password` - Indicates if the user can log in to the console.
	* `can_use_customer_secret_keys` - Indicates if the user can use SigV4 symmetric keys.
	* `can_use_smtp_credentials` - Indicates if the user can use SMTP passwords.
* `compartment_id` - The OCID of the tenancy containing the user.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the user. Does not have to be unique, and it's changeable.
* `email` - The email address you assign to the user. The email address must be unique across all users in the tenancy. 
* `external_identifier` - Identifier of the user in the identity provider
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the user.
* `identity_provider_id` - The OCID of the `IdentityProvider` this user belongs to.
* `inactive_state` - Returned only if the user's `lifecycleState` is INACTIVE. A 16-bit value showing the reason why the user is inactive:
	* bit 0: SUSPENDED (reserved for future use)
	* bit 1: DISABLED (reserved for future use)
	* bit 2: BLOCKED (the user has exceeded the maximum number of failed login attempts for the Console) 
* `name` - The name you assign to the user during creation. This is the user's login for the Console. The name must be unique across all users in the tenancy and cannot be changed. 
* `state` - The user's current state.
* `time_created` - Date and time the user was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

