---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_authentication_policy"
sidebar_current: "docs-oci-resource-identity-authentication_policy"
description: |-
  Provides the Authentication Policy resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_authentication_policy
This resource provides the Authentication Policy resource in Oracle Cloud Infrastructure Identity service.

Updates authentication policy for the specified tenancy


## Example Usage

```hcl
resource "oci_identity_authentication_policy" "test_authentication_policy" {
	#Required
	compartment_id = "${var.tenancy_ocid}"

	#Optional
	password_policy {

		#Optional
		is_lowercase_characters_required = "${var.authentication_policy_password_policy_is_lowercase_characters_required}"
		is_numeric_characters_required = "${var.authentication_policy_password_policy_is_numeric_characters_required}"
		is_special_characters_required = "${var.authentication_policy_password_policy_is_special_characters_required}"
		is_uppercase_characters_required = "${var.authentication_policy_password_policy_is_uppercase_characters_required}"
		is_username_containment_allowed = "${var.authentication_policy_password_policy_is_username_containment_allowed}"
		minimum_password_length = "${var.authentication_policy_password_policy_minimum_password_length}"
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `password_policy` - (Optional) (Updatable) Password policy.
	* `is_lowercase_characters_required` - (Optional) (Updatable) At least one lower case character required.
	* `is_numeric_characters_required` - (Optional) (Updatable) At least one numeric character required.
	* `is_special_characters_required` - (Optional) (Updatable) At least one special character required.
	* `is_uppercase_characters_required` - (Optional) (Updatable) At least one uppercase character required.
	* `is_username_containment_allowed` - (Optional) (Updatable) User name is allowed to be part of the password.
	* `minimum_password_length` - (Optional) (Updatable) Minimum password length required.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID.
* `password_policy` - Password policy.
	* `is_lowercase_characters_required` - At least one lower case character required.
	* `is_numeric_characters_required` - At least one numeric character required.
	* `is_special_characters_required` - At least one special character required.
	* `is_uppercase_characters_required` - At least one uppercase character required.
	* `is_username_containment_allowed` - User name is allowed to be part of the password.
	* `minimum_password_length` - Minimum password length required.

## Import

AuthenticationPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_identity_authentication_policy.test_authentication_policy "authenticationPolicies/{compartmentId}" 
```

