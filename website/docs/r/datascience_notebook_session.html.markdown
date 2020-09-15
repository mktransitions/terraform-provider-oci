---
subcategory: "Datascience"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_notebook_session"
sidebar_current: "docs-oci-resource-datascience-notebook_session"
description: |-
  Provides the Notebook Session resource in Oracle Cloud Infrastructure Datascience service
---

# oci_datascience_notebook_session
This resource provides the Notebook Session resource in Oracle Cloud Infrastructure Datascience service.

Creates a new notebook session.

## Example Usage

```hcl
resource "oci_datascience_notebook_session" "test_notebook_session" {
	#Required
	compartment_id = var.compartment_id
	notebook_session_configuration_details {
		#Required
		shape = var.notebook_session_notebook_session_configuration_details_shape
		subnet_id = oci_core_subnet.test_subnet.id

		#Optional
		block_storage_size_in_gbs = var.notebook_session_notebook_session_configuration_details_block_storage_size_in_gbs
	}
	project_id = oci_datascience_project.test_project.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.notebook_session_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the compartment where you want to create the notebook session.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information. Example: `My NotebookSession` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `notebook_session_configuration_details` - (Required) (Updatable) 
	* `block_storage_size_in_gbs` - (Optional) (Updatable) A notebook session instance is provided with a block storage volume. This specifies the size of the volume in GBs. 
	* `shape` - (Required) (Updatable) The shape used to launch the notebook session compute instance.  The list of available shapes in a given compartment can be retrieved from the `ListNotebookSessionShapes` endpoint. 
	* `subnet_id` - (Required) (Updatable) A notebook session instance is provided with a VNIC for network access.  This specifies the [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT gateway for egress to the internet. 
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the project to associate with the notebook session.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the notebook session's compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the user who created the notebook session.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information. Example: `My NotebookSession` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the notebook session.
* `lifecycle_details` - Details about the state of the notebook session.
* `notebook_session_configuration_details` - 
	* `block_storage_size_in_gbs` - A notebook session instance is provided with a block storage volume. This specifies the size of the volume in GBs. 
	* `shape` - The shape used to launch the notebook session compute instance.  The list of available shapes in a given compartment can be retrieved from the `ListNotebookSessionShapes` endpoint. 
	* `subnet_id` - A notebook session instance is provided with a VNIC for network access.  This specifies the [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT gateway for egress to the internet. 
* `notebook_session_url` - The URL to interact with the notebook session.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the project associated with the notebook session.
* `state` - The state of the notebook session.
* `time_created` - The date and time the resource was created, in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 

## Import

NotebookSessions can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_notebook_session.test_notebook_session "id"
```

