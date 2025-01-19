---
subcategory: "Certificates"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_certificates_certificate_bundle"
sidebar_current: "docs-oci-datasource-certificates-certificate_bundle"
description: |-
Provides details about a certificate bundle in Oracle Cloud Infrastructure Certificates Retrieval service
---

# Data Source: oci_certificates_certificate_bundle
This data source provides details about a specific certificate bundle in Oracle Cloud Infrastructure Certificates Retrieval service.

Gets details about the specified certificate bundle.

## Example Usage

```hcl
data "oci_certificates_certificate_bundle" "test_certificate_bundle" {
	#Required
	certificate_id = oci_certificates_management_certificate.test_certificate.id
  
    #Optional
    certificate_bundle_type = "CERTIFICATE_CONTENT_WITH_PRIVATE_KEY"
    certificate_version_name = oci_certificates_management_certificate.test_certificate.current_version.version_name
    stage = "CURRENT"
    version_number = oci_certificates_management_certificate.test_certificate.current_version.version_number
}
```

## Argument Reference

The following arguments are supported:

* `certificate_id` - (Required) The OCID of the certificate.
* `certificate_bundle_type` - (Optional) The type of certificate bundle. By default, the private key fields are not
returned. When querying for certificate bundles, to return results with certificate contents, the private key in PEM
format, and the private key passphrase, specify the value of this parameter as CERTIFICATE_CONTENT_WITH_PRIVATE_KEY.
Valid values are: `CERTIFICATE_CONTENT_PUBLIC_ONLY` or `CERTIFICATE_CONTENT_WITH_PRIVATE_KEY`.
* `certificate_version_name` - (Optional) The name of the certificate. (This might be referred to as the name of the
certificate version, as every certificate consists of at least one version.) Names are unique across versions of a
given certificate.
* `stage` - (Optional) The rotation state of the certificate version. Valid values are: `CURRENT`, `PENDING`, `LATEST`,
`PREVIOUS` or `DEPRECATED`.
* `version_number` - (Optional) The version number of the certificate.

## Attributes Reference

The following attributes are exported:

* `cert_chain_pem` - The certificate chain (in PEM format) for the certificate bundle.
* `certificate_bundle_type` - The type of certificate bundle, which indicates whether the private key fields are included.
* `certificate_id` - The OCID of the certificate.
* `certificate_name` - The name of the certificate.
* `certificate_pem` - The certificate (in PEM format) for the certificate bundle.
* `private_key_pem` - The private key (in PEM format) for the certificate. This is only set if `certificate_bundle_type`
is set to `CERTIFICATE_CONTENT_WITH_PRIVATE_KEY`.
* `private_key_pem_passphrase` - The passphrase for the private key. This is only set if `certificate_bundle_type`
is set to `CERTIFICATE_CONTENT_WITH_PRIVATE_KEY`.
* `revocation_status` - The revocation status of the certificate.
* `serial_number` - A unique certificate identifier used in certificate revocation tracking, formatted as octets.
* `stages` - A list of rotation states for the certificate bundle.
* `time_created` - An optional property indicating when the certificate version was created, expressed in RFC 3339
timestamp format.
* `validity` - The validity of the certificate.
* `version_name` - The name of the certificate version.
* `version_number` - The version number of the certificate.

### Revocation Status Reference

The following attributes are exported:

* `revocation_reason` - The reason that the certificate was revoked.
* `time_revoked` - The time when the certificate was revoked.

### Validity Reference

The following attributes are exported:

* `time_of_validity_not_after` - The date on which the certificate validity period ends, expressed in RFC 3339 timestamp
format.
* `time_of_validity_not_before` - The date on which the certificate validity period begins, expressed in RFC 3339
timestamp format.
