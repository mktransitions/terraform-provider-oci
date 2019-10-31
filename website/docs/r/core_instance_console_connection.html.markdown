---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_console_connection"
sidebar_current: "docs-oci-resource-core-instance_console_connection"
description: |-
  Provides the Instance Console Connection resource in Oracle Cloud Infrastructure Core service
---

# oci_core_instance_console_connection
This resource provides the Instance Console Connection resource in Oracle Cloud Infrastructure Core service.

Creates a new console connection to the specified instance.
Once the console connection has been created and is available,
you connect to the console using SSH.

For more information about console access, see [Accessing the Console](https://docs.cloud.oracle.com/iaas/Content/Compute/References/serialconsole.htm).


## Example Usage

```hcl
resource "oci_core_instance_console_connection" "test_instance_console_connection" {
	#Required
	instance_id = "${oci_core_instance.test_instance.id}"
	public_key = "${var.instance_console_connection_public_key}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `instance_id` - (Required) The OCID of the instance to create the console connection to.
* `public_key` - (Required) The SSH public key used to authenticate the console connection.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment to contain the console connection.
* `connection_string` - The SSH connection string for the console connection.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `fingerprint` - The SSH public key fingerprint for the console connection.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the console connection.
* `instance_id` - The OCID of the instance the console connection connects to.
* `state` - The current state of the console connection.
* `vnc_connection_string` - The SSH connection string for the SSH tunnel used to connect to the console connection over VNC. 

## Import

InstanceConsoleConnections can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance_console_connection.test_instance_console_connection "id"
```

