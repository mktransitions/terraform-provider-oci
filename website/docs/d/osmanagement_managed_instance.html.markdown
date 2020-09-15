---
subcategory: "Osmanagement"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_managed_instance"
sidebar_current: "docs-oci-datasource-osmanagement-managed_instance"
description: |-
  Provides details about a specific Managed Instance in Oracle Cloud Infrastructure Osmanagement service
---

# Data Source: oci_osmanagement_managed_instance
This data source provides details about a specific Managed Instance resource in Oracle Cloud Infrastructure Osmanagement service.

Returns a specific Managed Instance.


## Example Usage

```hcl
data "oci_osmanagement_managed_instance" "test_managed_instance" {
	#Required
	managed_instance_id = oci_osmanagement_managed_instance.test_managed_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_id` - (Required) OCID for the managed instance


## Attributes Reference

The following attributes are exported:

* `child_software_sources` - list of child Software Sources attached to the Managed Instance
	* `id` - software source identifier
	* `name` - software source name
* `compartment_id` - OCID for the Compartment
* `description` - Information specified by the user about the managed instance
* `display_name` - Managed Instance identifier
* `id` - OCID for the managed instance
* `is_reboot_required` - Indicates whether a reboot is required to complete installation of updates.
* `last_boot` - Time at which the instance last booted
* `last_checkin` - Time at which the instance last checked in
* `managed_instance_groups` - The ids of the managed instance groups of which this instance is a member. 
	* `display_name` - User friendly name
	* `id` - unique identifier that is immutable on creation
* `os_family` - The Operating System type of the managed instance.
* `os_kernel_version` - Operating System Kernel Version
* `os_name` - Operating System Name
* `os_version` - Operating System Version
* `parent_software_source` - the parent (base) Software Source attached to the Managed Instance
	* `id` - software source identifier
	* `name` - software source name
* `status` - status of the managed instance.
* `updates_available` - Number of updates available to be installed

