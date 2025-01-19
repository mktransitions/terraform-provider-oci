---
subcategory: "Certificates"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_certificates_certificate_authority_bundle"
sidebar_current: "docs-oci-datasource-certificates-certificate_authority_bundle"
description: |-
Provides details about a certificate authority bundle in Oracle Cloud Infrastructure Certificates Retrieval service
---

# Data Source: oci_certificates_certificate_bundle
This data source provides details about a specific certificate authority bundle in Oracle Cloud Infrastructure Certificates Retrieval service.

Gets details about the specified certificate authority bundle.

## Example Usage

```hcl
data "oci_certificates_certificate_authority_bundle" "test_certificate_authority_bundle" {
	#Required
	certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
  
    #Optional
    certificate_version_name = oci_certificates_management_certificate_authority.test_certificate_authority.current_version.version_name
    stage = "CURRENT"
    version_number = oci_certificates_management_certificate_authority.test_certificate_authority.current_version.version_number
}
```

## Argument Reference

The following arguments are supported:

* `certificate_authority_id` - (Required) The OCID of the certificate authority (CA).
* `certificate_version_name` - (Optional) The name of the certificate authority (CA). (This might be referred to as the
name of the CA version, as every CA consists of at least one version.) Names are unique across versions of a given CA.
* `stage` - (Optional) The rotation state of the certificate authority version. Valid values are: `CURRENT`, `PENDING`,
`LATEST`, `PREVIOUS` or `DEPRECATED`.
* `version_number` - (Optional) The version number of the certificate authority (CA).

## Attributes Reference

The following attributes are exported:

* `cert_chain_pem` - The certificate chain (in PEM format) for this CA version.
* `certificate_authority_id` - The OCID of the certificate authority (CA).
* `certificate_authority_name` - The name of the CA.
* `certificate_pem` - The certificate (in PEM format) for this CA version.
* `revocation_status` - The revocation status of the certificate.
* `serial_number` - A unique certificate identifier used in certificate revocation tracking, formatted as octets.
* `stages` - A list of rotation states for this CA.
* `time_created` - An optional property indicating when the certificate version was created, expressed in RFC 3339
timestamp format.
* `validity` - The validity of the certificate.
* `version_name` - The name of the CA version.
* `version_number` - The version number of the CA.

### Revocation Status Reference

The following attributes are exported:

* `revocation_reason` - The reason that the CA was revoked.
* `time_revoked` - The time when the CA was revoked.

### Validity Reference

The following attributes are exported:

* `time_of_validity_not_after` - The date on which the CA validity period ends, expressed in RFC 3339 timestamp
format.
* `time_of_validity_not_before` - The date on which the CA validity period begins, expressed in RFC 3339
timestamp format.
