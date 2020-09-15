---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_gi_versions"
sidebar_current: "docs-oci-datasource-database-gi_versions"
description: |-
  Provides the list of Gi Versions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_gi_versions
This data source provides the list of Gi Versions in Oracle Cloud Infrastructure Database service.

Gets a list of supported GI versions for VM Cluster.

## Example Usage

```hcl
data "oci_database_gi_versions" "test_gi_versions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	shape = var.gi_version_shape
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `shape` - (Optional) If provided, filters the results for the given shape.


## Attributes Reference

The following attributes are exported:

* `gi_versions` - The list of gi_versions.

### GiVersion Reference

The following attributes are exported:

* `version` - A valid Oracle Grid Infrastructure (GI) software version.

