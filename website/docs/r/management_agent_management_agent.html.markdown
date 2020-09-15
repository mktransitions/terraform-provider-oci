---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent"
sidebar_current: "docs-oci-resource-management_agent-management_agent"
description: |-
  Provides the Management Agent resource in Oracle Cloud Infrastructure Management Agent service
---

# oci_management_agent_management_agent
This resource provides the Management Agent resource in Oracle Cloud Infrastructure Management Agent service.



## Example Usage

```hcl
resource "oci_management_agent_management_agent" "test_management_agent" {
	#Required
	managed_agent_id = oci_management_agent_managed_agent.test_managed_agent.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_agent_id` - (Required) Unique Management Agent identifier
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) New displayName of Agent.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_agent_auto_upgradable` - (Optional) (Updatable) if set to true then, agent can be upgraded automatically else needs to be upgraded manually.
* `deploy_plugins_id` - (Optional) (Updatable) Plugin Id list

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Management Agent Name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `host` - Management Agent host machine name
* `id` - agent identifier
* `install_key_id` - agent install key identifier
* `install_path` - Path where Management Agent is installed
* `is_agent_auto_upgradable` - true if the agent can be upgraded automatically; false if it must be upgraded manually. true is currently unsupported.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `platform_name` - Platform Name
* `platform_type` - Platform Type
* `platform_version` - Platform Version
* `plugin_list` - list of managementAgentPlugins associated with the agent
	* `plugin_display_name` - Management Agent Plugin Identifier, can be renamed
	* `plugin_id` - Plugin Id
	* `plugin_name` - Management Agent Plugin Name
	* `plugin_version` - Plugin Version
* `state` - The current state of managementAgent
* `time_created` - The time the Management Agent was created. An RFC3339 formatted datetime string
* `time_last_heartbeat` - The time the Management Agent has last recorded its health status in telemetry. This value will be null if the agent has not recorded its health status in last 7 days. An RFC3339 formatted datetime string
* `time_updated` - The time the Management Agent was updated. An RFC3339 formatted datetime string
* `version` - Management Agent Version

## Import

ManagementAgents can be imported using the `id`, e.g.

```
$ terraform import oci_management_agent_management_agent.test_management_agent "id"
```

