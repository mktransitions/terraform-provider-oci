---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_user"
sidebar_current: "docs-oci-resource-identity-user"
description: |-
  Provides the User resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_user
This resource provides the User resource in Oracle Cloud Infrastructure Identity service.

Creates a new user in your tenancy. For conceptual information about users, your tenancy, and other
IAM Service components, see [Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

You must specify your tenancy's OCID as the compartment ID in the request object (remember that the
tenancy is simply the root compartment). Notice that IAM resources (users, groups, compartments, and
some policies) reside within the tenancy itself, unlike cloud resources such as compute instances,
which typically reside within compartments inside the tenancy. For information about OCIDs, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You must also specify a *name* for the user, which must be unique across all users in your tenancy
and cannot be changed. Allowed characters: No spaces. Only letters, numerals, hyphens, periods,
underscores, +, and @. If you specify a name that's already in use, you'll get a 409 error.
This name will be the user's login to the Console. You might want to pick a
name that your company's own identity system (e.g., Active Directory, LDAP, etc.) already uses.
If you delete a user and then create a new user with the same name, they'll be considered different
users because they have different OCIDs.

You must also specify a *description* for the user (although it can be an empty string).
It does not have to be unique, and you can change it anytime with
[UpdateUser](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/User/UpdateUser). You can use the field to provide the user's
full name, a description, a nickname, or other information to generally identify the user.
A new user has no permissions until you place the user in one or more groups (see
[AddUserToGroup](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/UserGroupMembership/AddUserToGroup)). If the user needs to
access the Console, you need to provide the user a password (see
[CreateOrResetUIPassword](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/UIPassword/CreateOrResetUIPassword)).
If the user needs to access the Oracle Cloud Infrastructure REST API, you need to upload a
public API signing key for that user (see
[Required Keys and OCIDs](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm) and also
[UploadApiKey](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/ApiKey/UploadApiKey)).

**Important:** Make sure to inform the new user which compartment(s) they have access to.


## Example Usage

```hcl
resource "oci_identity_user" "test_user" {
	#Required
	compartment_id = var.tenancy_ocid
	description = var.user_description
	name = var.user_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	email = var.user_email
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy containing the user.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) The description you assign to the user during creation. Does not have to be unique, and it's changeable.
* `email` - (Required for tenancies that support identity domains, Optional for those that do not) (Updatable) The email you assign to the user. Has to be unique across the tenancy.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) The name you assign to the user during creation. This is the user's login for the Console. The name must be unique across all users in the tenancy and cannot be changed. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `capabilities` - Properties indicating how the user is allowed to authenticate.
	* `can_use_api_keys` - Indicates if the user can use API keys.
	* `can_use_auth_tokens` - Indicates if the user can use SWIFT passwords / auth tokens.
	* `can_use_console_password` - Indicates if the user can log in to the console.
	* `can_use_customer_secret_keys` - Indicates if the user can use SigV4 symmetric keys.
	* `can_use_db_credentials` - Indicates if the user can use DB passwords. 
	* `can_use_oauth2client_credentials` - Indicates if the user can use OAuth2 credentials and tokens. 
	* `can_use_smtp_credentials` - Indicates if the user can use SMTP passwords.
* `compartment_id` - The OCID of the tenancy containing the user.
* `db_user_name` - DB username of the DB credential. Has to be unique across the tenancy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the user. Does not have to be unique, and it's changeable.
* `email` - The email address you assign to the user. The email address must be unique across all users in the tenancy. 
* `email_verified` - Whether the email address has been validated.
* `external_identifier` - Identifier of the user in the identity provider
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the user.
* `identity_provider_id` - The OCID of the `IdentityProvider` this user belongs to.
* `inactive_state` - Returned only if the user's `lifecycleState` is INACTIVE. A 16-bit value showing the reason why the user is inactive:
	* bit 0: SUSPENDED (reserved for future use)
	* bit 1: DISABLED (reserved for future use)
	* bit 2: BLOCKED (the user has exceeded the maximum number of failed login attempts for the Console) 
* `last_successful_login_time` - The date and time of when the user most recently logged in the format defined by RFC3339 (ex. `2016-08-25T21:10:29.600Z`). If there is no login history, this field is null.

	For illustrative purposes, suppose we have a user who has logged in at July 1st, 2020 at 1200 PST and logged out 30 minutes later. They then login again on July 2nd, 2020 at 1500 PST.

	Their previousSuccessfulLoginTime would be `2020-07-01:19:00.000Z`.

	Their lastSuccessfulLoginTime would be `2020-07-02:22:00.000Z`. 
* `name` - The name you assign to the user during creation. This is the user's login for the Console. The name must be unique across all users in the tenancy and cannot be changed. 
* `previous_successful_login_time` - The date and time of when the user most recently logged in the format defined by RFC3339 (ex. `2016-08-25T21:10:29.600Z`). If there is no login history, this field is null.

	For illustrative purposes, suppose we have a user who has logged in at July 1st, 2020 at 1200 PST and logged out 30 minutes later. They then login again on July 2nd, 2020 at 1500 PST.

	Their previousSuccessfulLoginTime would be `2020-07-01:19:00.000Z`.

	Their lastSuccessfulLoginTime would be `2020-07-02:22:00.000Z`. 
* `state` - The user's current state.
* `time_created` - Date and time the user was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the User
	* `update` - (Defaults to 20 minutes), when updating the User
	* `delete` - (Defaults to 20 minutes), when destroying the User


## Import

Users can be imported using the `id`, e.g.

```
$ terraform import oci_identity_user.test_user "id"
```

