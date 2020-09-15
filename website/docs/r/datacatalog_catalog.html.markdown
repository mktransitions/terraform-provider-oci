---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_catalog"
sidebar_current: "docs-oci-resource-datacatalog-catalog"
description: |-
  Provides the Catalog resource in Oracle Cloud Infrastructure Data Catalog service
---

# oci_datacatalog_catalog
This resource provides the Catalog resource in Oracle Cloud Infrastructure Data Catalog service.

Creates a new data catalog instance that includes a console and an API URL for managing metadata operations.
For more information, please see the documentation.


## Example Usage

```hcl
resource "oci_datacatalog_catalog" "test_catalog" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.catalog_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment identifier.
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) Data catalog identifier.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `attached_catalog_private_endpoints` - (Optional) (Updatable) The list of private reverse connection endpoints attached to the catalog


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `attached_catalog_private_endpoints` - The list of private reverse connection endpoints attached to the catalog
* `compartment_id` - Compartment identifier.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Data catalog identifier, which can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - An message describing the current state in more detail.  For example, it can be used to provide actionable information for a resource in 'Failed' state. 
* `number_of_objects` - The number of data objects added to the data catalog. Please see the data catalog documentation for further information on how this is calculated. 
* `service_api_url` - The REST front endpoint URL to the data catalog instance.
* `service_console_url` - The console front endpoint URL to the data catalog instance.
* `state` - The current state of the data catalog resource.
* `time_created` - The time the data catalog was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - The time the data catalog was updated. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.

## Import

Catalogs can be imported using the `id`, e.g.

```
$ terraform import oci_datacatalog_catalog.test_catalog "id"
```

