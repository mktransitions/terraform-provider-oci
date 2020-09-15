---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_image"
sidebar_current: "docs-oci-datasource-core-image"
description: |-
  Provides details about a specific Image in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_image
This data source provides details about a specific Image resource in Oracle Cloud Infrastructure Core service.

Gets the specified image.

## Example Usage

```hcl
data "oci_core_image" "test_image" {
	#Required
	image_id = oci_core_image.test_image.id
}
```

## Argument Reference

The following arguments are supported:

* `image_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the image.


## Attributes Reference

The following attributes are exported:

* `agent_features` - 
	* `is_management_supported` - Whether the agent running on the instance can run all the available management plugins 
	* `is_monitoring_supported` - Whether the agent running on the instance can gather performance metrics and monitor the instance. 
* `base_image_id` - The OCID of the image originally used to launch the instance.
* `compartment_id` - The OCID of the compartment containing the instance you want to use as the basis for the image. 
* `create_image_allowed` - Whether instances launched with this image can be used to create new images. For example, you cannot create an image of an Oracle Database instance.  Example: `true` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the image. It does not have to be unique, and it's changeable. Avoid entering confidential information.

	You cannot use an Oracle-provided image name as a custom image name.

	Example: `My custom Oracle Linux image` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the image.
* `launch_mode` - Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
	* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for Oracle-provided images.
	* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
	* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
	* `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter. 
* `launch_options` - 
	* `boot_volume_type` - Emulation type for the boot volume.
		* `ISCSI` - ISCSI attached block storage device.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage.  This is the default option for local data volumes on Oracle-provided images.
		* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on Oracle-provided images. 
	* `firmware` - Firmware used to boot VM.  Select the option that matches your operating system.
		* `BIOS` - Boot VM using BIOS style firmware.  This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
		* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems.  This is the default for Oracle-provided images. 
	* `is_consistent_volume_naming_enabled` - Whether to enable consistent volume naming feature. Defaults to false.
	* `is_pv_encryption_in_transit_enabled` - Deprecated. Instead use `isPvEncryptionInTransitEnabled` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/datatypes/LaunchInstanceDetails). 
	* `network_type` - Emulation type for the physical network interface card (NIC).
		* `E1000` - Emulated Gigabit ethernet controller.  Compatible with Linux e1000 network driver.
		* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
		* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers. 
	* `remote_data_volume_type` - Emulation type for volume.
		* `ISCSI` - ISCSI attached block storage device.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage.  This is the default option for local data volumes on Oracle-provided images.
		* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on Oracle-provided images. 
* `operating_system` - The image's operating system.  Example: `Oracle Linux` 
* `operating_system_version` - The image's operating system version.  Example: `7.2` 
* `size_in_mbs` - The boot volume size for an instance launched from this image, (1 MB = 1048576 bytes). Note this is not the same as the size of the image when it was exported or the actual size of the image.  Example: `47694` 
* `state` - The current state of the image.
* `time_created` - The date and time the image was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

