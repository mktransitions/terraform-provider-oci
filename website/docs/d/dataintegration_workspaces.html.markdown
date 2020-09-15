---
subcategory: "Dataintegration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspaces"
sidebar_current: "docs-oci-datasource-dataintegration-workspaces"
description: |-
  Provides the list of Workspaces in Oracle Cloud Infrastructure Dataintegration service
---

# Data Source: oci_dataintegration_workspaces
This data source provides the list of Workspaces in Oracle Cloud Infrastructure Dataintegration service.

Returns a list of Data Integration Workspaces.


## Example Usage

```hcl
data "oci_dataintegration_workspaces" "test_workspaces" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.workspace_name
	state = var.workspace_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `name` - (Optional) This filter parameter can be used to filter by the name of the object.
* `state` - (Optional) Lifecycle state of the resource.


## Attributes Reference

The following attributes are exported:

* `workspaces` - The list of workspaces.

### Workspace Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the workspace.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A detailed description for the workspace.
* `display_name` - A user-friendly display name for the workspace. Does not have to be unique, and can be modified. Avoid entering confidential information.
* `dns_server_ip` - The IP of the custom DNS.
* `dns_server_zone` - The DNS zone of the custom DNS to use to resolve names.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Unique identifier that is immutable on creation
* `is_private_network_enabled` - Whether the private network connection is enabled or disabled.
* `state` - Lifecycle states for workspaces in Data Integration Service CREATING - The resource is being created and may not be usable until the entire metadata is defined UPDATING - The resource is being updated and may not be usable until all changes are commited DELETING - The resource is being deleted and might require deep cleanup of children. ACTIVE   - The resource is valid and available for access INACTIVE - The resource might be incomplete in its definition or might have been made unavailable for administrative reasons DELETED  - The resource has been deleted and isn't available FAILED   - The resource is in a failed state due to validation or other errors STARTING - The resource is being started and may not be usable until becomes ACTIVE again STOPPING - The resource is in the process of Stopping and may not be usable until it Stops or fails STOPPED  - The resource is in Stopped state due to stop operation. 
* `state_message` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `subnet_id` - The OCID of the subnet for customer connected databases.
* `time_created` - The date and time the workspace was created, in the timestamp format defined by RFC3339. 
* `time_updated` - The date and time the workspace was updated, in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `vcn_id` - The OCID of the VCN the subnet is in.

