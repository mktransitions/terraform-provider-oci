---
subcategory: "Vault"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_vault_secret"
sidebar_current: "docs-oci-datasource-vault-secret"
description: |-
  Provides details about a specific Secret in Oracle Cloud Infrastructure Vault service
---

# Data Source: oci_vault_secret
This data source provides details about a specific Secret resource in Oracle Cloud Infrastructure Vault service.

Gets information about the specified secret.

## Example Usage

```hcl
data "oci_vault_secret" "test_secret" {
	#Required
	secret_id = oci_vault_secret.test_secret.id
}
```

## Argument Reference

The following arguments are supported:

* `secret_id` - (Required) The OCID of the secret.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment where you want to create the secret.
* `current_version_number` - The version number of the secret version that's currently in use.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A brief description of the secret. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the secret.
* `key_id` - The OCID of the master encryption key that is used to encrypt the secret.
* `lifecycle_details` - Additional information about the current lifecycle state of the secret.
* `metadata` - Additional metadata that you can use to provide context about how to use the secret or during rotation or other administrative tasks. For example, for a secret that you use to connect to a database, the additional metadata might specify the connection endpoint and the connection string. Provide additional metadata as key-value pairs. 
* `secret_name` - The user-friendly name of the secret. Avoid entering confidential information.
* `secret_rules` - A list of rules that control how the secret is used and managed.
	* `is_enforced_on_deleted_secret_versions` - A property indicating whether the rule is applied even if the secret version with the content you are trying to reuse was deleted. 
	* `is_secret_content_retrieval_blocked_on_expiry` - A property indicating whether to block retrieval of the secret content, on expiry. The default is false. If the secret has already expired and you would like to retrieve the secret contents, you need to edit the secret rule to disable this property, to allow reading the secret content. 
	* `rule_type` - The type of rule, which either controls when the secret contents expire or whether they can be reused.
	* `secret_version_expiry_interval` - A property indicating how long the secret contents will be considered valid, expressed in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format. The secret needs to be updated when the secret content expires. No enforcement mechanism exists at this time, but audit logs record the expiration on the appropriate date, according to the time interval specified in the rule. The timer resets after you update the secret contents. The minimum value is 1 day and the maximum value is 90 days for this property. Currently, only intervals expressed in days are supported. For example, pass `P3D` to have the secret version expire every 3 days. 
	* `time_of_absolute_expiry` - An optional property indicating the absolute time when this secret will expire, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. The minimum number of days from current time is 1 day and the maximum number of days from current time is 365 days. Example: `2019-04-03T21:10:29.600Z` 
* `state` - The current lifecycle state of the secret.
* `time_created` - A property indicating when the secret was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `time_of_current_version_expiry` - An optional property indicating when the current secret version will expire, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property indicating when to delete the secret, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `vault_id` - The OCID of the Vault in which the secret exists

