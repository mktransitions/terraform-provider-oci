---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_stream_cdn_config"
sidebar_current: "docs-oci-datasource-media_services-stream_cdn_config"
description: |-
  Provides details about a specific Stream Cdn Config in Oracle Cloud Infrastructure Media Services service
---

# Data Source: oci_media_services_stream_cdn_config
This data source provides details about a specific Stream Cdn Config resource in Oracle Cloud Infrastructure Media Services service.

Gets a StreamCdnConfig by identifier.

## Example Usage

```hcl
data "oci_media_services_stream_cdn_config" "test_stream_cdn_config" {
	#Required
	stream_cdn_config_id = oci_media_services_stream_cdn_config.test_stream_cdn_config.id
}
```

## Argument Reference

The following arguments are supported:

* `stream_cdn_config_id` - (Required) Unique StreamCdnConfig identifier.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier.
* `config` - Base fields of the StreamCdnConfig configuration object.
	* `edge_hostname` - The hostname of the CDN edge server to use when building CDN URLs.
	* `edge_path_prefix` - The path to prepend when building CDN URLs.
	* `edge_token_key` - The encryption key to use for edge token authentication.
	* `edge_token_salt` - Salt to use when encrypting authentication token.
	* `is_edge_token_auth` - Whether token authentication should be used at the CDN edge.
	* `origin_auth_secret_key_a` - The shared secret key A, two for errorless key rotation.
	* `origin_auth_secret_key_b` - The shared secret key B, two for errorless key rotation.
	* `origin_auth_secret_key_nonce_a` - Nonce identifier for originAuthSecretKeyA (used to determine key used to sign).
	* `origin_auth_secret_key_nonce_b` - Nonce identifier for originAuthSecretKeyB (used to determine key used to sign).
	* `origin_auth_sign_encryption` - The type of encryption used to compute the signature.
	* `origin_auth_sign_type` - The type of data used to compute the signature.
	* `type` - The name of the CDN configuration type.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The CDN Configuration identifier or display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `distribution_channel_id` - Distribution Channel Identifier.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `is_enabled` - Whether publishing to CDN is enabled.
* `lifecyle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `state` - The current state of the CDN Configuration.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the CDN Config was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the CDN Config was updated. An RFC3339 formatted datetime string.

