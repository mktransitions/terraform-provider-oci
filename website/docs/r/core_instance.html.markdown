---
layout: "oci"
page_title: "OCI: oci_core_instance"
sidebar_current: "docs-oci-resource-core-instance"
description: |-
  Creates and manages an OCI Instance
---

# oci_core_instance
The `oci_core_instance` resource creates and manages an OCI Instance

Creates a new instance in the specified compartment and the specified availability domain.
For general information about instances, see
[Overview of the Compute Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Concepts/computeoverview.htm).

For information about access control and compartments, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).

For information about availability domains, see
[Regions and Availability Domains](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/regions.htm).
To get a list of availability domains, use the `ListAvailabilityDomains` operation
in the Identity and Access Management Service API.

All Oracle Cloud Infrastructure resources, including instances, get an Oracle-assigned,
unique ID called an Oracle Cloud Identifier (OCID).
When you create a resource, you can find its OCID in the response. You can
also retrieve a resource's OCID by using a List API operation
on that resource type, or by viewing the resource in the Console.

To launch an instance using an image or a boot volume use the `sourceDetails` parameter in [LaunchInstanceDetails](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/LaunchInstanceDetails).

When you launch an instance, it is automatically attached to a virtual
network interface card (VNIC), called the *primary VNIC*. The VNIC
has a private IP address from the subnet's CIDR. You can either assign a
private IP address of your choice or let Oracle automatically assign one.
You can choose whether the instance has a public IP address. To retrieve the
addresses, use the [ListVnicAttachments](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VnicAttachment/ListVnicAttachments)
operation to get the VNIC ID for the instance, and then call
[GetVnic](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vnic/GetVnic) with the VNIC ID.

You can later add secondary VNICs to an instance. For more information, see
[Virtual Network Interface Cards (VNICs)](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingVNICs.htm).


## Example Usage

```hcl
resource "oci_core_instance" "test_instance" {
	#Required
	availability_domain = "${var.instance_availability_domain}"
	compartment_id = "${var.compartment_id}"
	shape = "${var.instance_shape}"

	#Optional
	create_vnic_details {
		#Required
		subnet_id = "${oci_core_subnet.test_subnet.id}"

		#Optional
		assign_public_ip = "${var.instance_create_vnic_details_assign_public_ip}"
		defined_tags = {"Operations.CostCenter"= "42"}
		display_name = "${var.instance_create_vnic_details_display_name}"
		freeform_tags = {"Department"= "Finance"}
		hostname_label = "${var.instance_create_vnic_details_hostname_label}"
		private_ip = "${var.instance_create_vnic_details_private_ip}"
		skip_source_dest_check = "${var.instance_create_vnic_details_skip_source_dest_check}"
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.instance_display_name}"
	extended_metadata {
		some_string = "stringA"
		nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
	}
	fault_domain = "${var.instance_fault_domain}"
	freeform_tags = {"Department"= "Finance"}
	hostname_label = "${var.instance_hostname_label}"
	ipxe_script = "${var.instance_ipxe_script}"
	metadata {
		ssh_authorized_keys = "${var.ssh_public_key}"
		user_data = "${base64encode(file(var.custom_bootstrap_file_name))}"
	}
	source_details {
		#Required
		source_id = "${oci_core_image.test_image.id}"
		source_type = "image"

		#Optional
		boot_volume_size_in_gbs = "60"
	}
	preserve_boot_volume = false
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `create_vnic_details` - (Optional) Details for the primary VNIC, which is automatically created and attached when the instance is launched. 
	* `assign_public_ip` - (Optional) Whether the VNIC should be assigned a public IP address. Defaults to whether the subnet is public or private. If not set and the VNIC is being created in a private subnet (that is, where `prohibitPublicIpOnVnic` = true in the [Subnet](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Subnet/)), then no public IP address is assigned. If not set and the subnet is public (`prohibitPublicIpOnVnic` = false), then a public IP address is assigned. If set to true and `prohibitPublicIpOnVnic` = true, an error is returned.

		**Note:** This public IP address is associated with the primary private IP on the VNIC. For more information, see [IP Addresses](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingIPaddresses.htm).

		**Note:** There's a limit to the number of [public IPs](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PublicIp/) a VNIC or instance can have. If you try to create a secondary VNIC with an assigned public IP for an instance that has already reached its public IP limit, an error is returned. For information about the public IP limits, see [Public IP Addresses](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingpublicIPs.htm).

		Example: `false` 
	* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - (Optional) (Updatable) A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information.
	* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `hostname_label` - (Optional) (Updatable) The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123). The value appears in the [Vnic](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vnic/) object and also the [PrivateIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/) object returned by [ListPrivateIps](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/GetPrivateIp).

		For more information, see [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).

		When launching an instance, use this `hostnameLabel` instead of the deprecated `hostnameLabel` in [LaunchInstanceDetails](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/requests/LaunchInstanceDetails). If you provide both, the values must match.

		Example: `bminstance-1` 
	* `private_ip` - (Optional) A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This is the VNIC's *primary* private IP address. The value appears in the [Vnic](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vnic/) object and also the [PrivateIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/) object returned by [ListPrivateIps](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/GetPrivateIp).  Example: `10.0.3.3` 
	* `skip_source_dest_check` - (Optional) (Updatable) Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. For information about why you would skip the source/destination check, see [Using a Private IP as a Route Target](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm#privateip).  Example: `true` 
	* `subnet_id` - (Required) The OCID of the subnet to create the VNIC in. When launching an instance, use this `subnetId` instead of the deprecated `subnetId` in [LaunchInstanceDetails](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/requests/LaunchInstanceDetails). At least one of them is required; if you provide both, the values must match. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My bare metal instance` 
* `extended_metadata` - (Optional) (Updatable) Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the 'metadata' object.

	They are distinguished from 'metadata' fields in that these can be nested JSON objects (whereas 'metadata' fields are string/string maps only).

	If you don't need nested metadata values, it is strongly advised to avoid using this object and use the Metadata object instead.

	Input in terraform is the same as metadata but allows nested metadata if you pass a valid JSON string as a value. See the example above.
* `fault_domain` - (Optional) A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

	If you do not specify the fault domain, the system selects one for you. To change the fault domain for an instance, terminate it and launch a new instance in the preferred fault domain.

	Example: `FAULT-DOMAIN-1`
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - (Optional) Deprecated. Instead use `hostnameLabel` in [CreateVnicDetails](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/CreateVnicDetails/). If you provide both, the values must match. 
* `image` - (Optional) Deprecated. Use `sourceDetails` with [InstanceSourceViaImageDetails](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/latest/requests/InstanceSourceViaImageDetails) source type instead. If you specify values for both, the values must match. 
* `ipxe_script` - (Optional) This is an advanced option.

	When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

	If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

	The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

	For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/References/bringyourownimage.htm).

	For more information about iPXE, see http://ipxe.org. 
* `metadata` - (Optional) (Updatable) Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.

	A metadata service runs on every launched instance. The service is an HTTP endpoint listening on 169.254.169.254. You can use the service to:
	* Provide information to [Cloud-Init](https://cloudinit.readthedocs.org/en/latest/) to be used for various system initialization tasks.
	* Get information about the instance, including the custom metadata that you provide when you launch the instance.

	**Providing Cloud-Init Metadata**

	You can use the following metadata key names to provide information to Cloud-Init:

	**"ssh_authorized_keys"** - Provide one or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on the instance. Use a newline character to separate multiple keys. The SSH keys must be in the format necessary for the `authorized_keys` file, as shown in the example below.

	**"user_data"** - Provide your own base64-encoded data to be used by Cloud-Init to run custom scripts or provide custom Cloud-Init configuration. For information about how to take advantage of user data, see the [Cloud-Init Documentation](http://cloudinit.readthedocs.org/en/latest/topics/format.html).

	**Note:** Cloud-Init does not pull this data from the `http://169.254.169.254/opc/v1/instance/metadata/` path. When the instance launches and either of these keys are provided, the key values are formatted as OpenStack metadata and copied to the following locations, which are recognized by Cloud-Init:

	`http://169.254.169.254/openstack/latest/meta_data.json` - This JSON blob contains, among other things, the SSH keys that you provided for **"ssh_authorized_keys"**.

	`http://169.254.169.254/openstack/latest/user_data` - Contains the base64-decoded data that you provided for **"user_data"**.

	**Metadata Example**

	"metadata" : { "quake_bot_level" : "Severe", "ssh_authorized_keys" : "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCZ06fccNTQfq+xubFlJ5ZR3kt+uzspdH9tXL+lAejSM1NXM+CFZev7MIxfEjas06y80ZBZ7DUTQO0GxJPeD8NCOb1VorF8M4xuLwrmzRtkoZzU16umt4y1W0Q4ifdp3IiiU0U8/WxczSXcUVZOLqkz5dc6oMHdMVpkimietWzGZ4LBBsH/LjEVY7E0V+a0sNchlVDIZcm7ErReBLcdTGDq0uLBiuChyl6RUkX1PNhusquTGwK7zc8OBXkRuubn5UKXhI3Ul9Nyk4XESkVWIGNKmw8mSpoJSjR8P9ZjRmcZVo8S+x4KVPMZKQEor== ryan.smith@company.com ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAQEAzJSAtwEPoB3Jmr58IXrDGzLuDYkWAYg8AsLYlo6JZvKpjY1xednIcfEVQJm4T2DhVmdWhRrwQ8DmayVZvBkLt+zs2LdoAJEVimKwXcJFD/7wtH8Lnk17HiglbbbNXsemjDY0hea4JUE5CfvkIdZBITuMrfqSmA4n3VNoorXYdvtTMoGG8fxMub46RPtuxtqi9bG9Zqenordkg5FJt2mVNfQRqf83CWojcOkklUWq4CjyxaeLf5i9gv1fRoBo4QhiA8I6NCSppO8GnoV/6Ox6TNoh9BiifqGKC9VGYuC89RvUajRBTZSK2TK4DPfaT+2R+slPsFrwiT/oPEhhEK1S5Q== rsa-key-20160227", "user_data" : "SWYgeW91IGNhbiBzZWUgdGhpcywgdGhlbiBpdCB3b3JrZWQgbWF5YmUuCg==" } **Getting Metadata on the Instance**

	To get information about your instance, connect to the instance using SSH and issue any of the following GET requests:

	curl http://169.254.169.254/opc/v1/instance/ curl http://169.254.169.254/opc/v1/instance/metadata/ curl http://169.254.169.254/opc/v1/instance/metadata/<any-key-name>

	You'll get back a response that includes all the instance information; only the metadata information; or the metadata information for the specified key name, respectively. 
* `preserve_boot_volume` - (Optional) Specifies whether to delete or preserve the boot volume when terminating an instance. The default value is false. Note: This value only applies to destroy operations initiated by Terraform.
* `shape` - (Required) The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

	You can enumerate all available shapes by calling [ListShapes](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Shape/ListShapes). 
* `source_details` - (Optional) Details for creating an instance. Use this parameter to specify whether a boot volume or an image should be used to launch a new instance. 
	* `boot_volume_size_in_gbs` - (Applicable when source_type=image) The size of the boot volume in GBs. Minimum value is 50 GB and maximum value is 16384 GB (16TB).
	* `source_id` - (Required) The OCID of an image or a boot volume to use, depending on the value of `source_type`.
	* `source_type` - (Required) The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
* `subnet_id` - (Optional) Deprecated. Instead use `subnetId` in [CreateVnicDetails](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/CreateVnicDetails/). At least one of them is required; if you provide both, the values must match. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the instance is running in.  Example: `Uocm:PHX-AD-1` 
* `boot_volume_id` - The OCID of the attached boot volume. If the `source_type` is `bootVolume`, this will be the same OCID as the `source_id`.
* `compartment_id` - The OCID of the compartment that contains the instance.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My bare metal instance` 
* `extended_metadata` - Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the 'metadata' object.

	They are distinguished from 'metadata' fields in that these can be nested JSON objects (whereas 'metadata' fields are string/string maps only).

	If you don't need nested metadata values, it is strongly advised to avoid using this object and use the Metadata object instead.

	Input in terraform is the same as metadata but allows nested metadata if you pass a valid JSON string as a value. See the example below.
* `fault_domain` - The name of the fault domain the instance is running in.

	A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

	If you do not specify the fault domain, the system selects one for you. To change the fault domain for an instance, terminate it and launch a new instance in the preferred fault domain.

	Example: `FAULT-DOMAIN-1` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the instance.
* `image` - Deprecated. Use `sourceDetails` instead. 
* `ipxe_script` - When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

	If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

	The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

	For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/References/bringyourownimage.htm).

	For more information about iPXE, see http://ipxe.org. 
* `launch_mode` - Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
	* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for Oracle-provided images.
	* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
	* `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter. 
* `launch_options` - 
	* `boot_volume_type` - Emulation type for volume.
		* `ISCSI` - ISCSI attached block storage device. This is the default for Boot Volumes and Remote Block Storage volumes on Oracle provided images.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data volumes on Oracle provided images.
		* `PARAVIRTUALIZED` - Paravirtualized disk. 
	* `firmware` - Firmware used to boot VM.  Select the option that matches your operating system.
		* `BIOS` - Boot VM using BIOS style firmware.  This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
		* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems.  This is the default for Oracle provided images. 
	* `network_type` - Emulation type for NIC.
		* `E1000` - Emulated Gigabit ethernet controller.  Compatible with Linux e1000 network driver.
		* `VFIO` - Direct attached Virtual Function network controller.  Default for Oracle provided images. 
	* `remote_data_volume_type` - Emulation type for volume.
		* `ISCSI` - ISCSI attached block storage device. This is the default for Boot Volumes and Remote Block Storage volumes on Oracle provided images.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data volumes on Oracle provided images.
		* `PARAVIRTUALIZED` - Paravirtualized disk. 
* `metadata` - Custom metadata that you provide.
* `preserve_boot_volume` - Specifies whether to delete or preserve the boot volume when terminating an instance. The default value is false. Note: This value only applies to destroy operations initiated by Terraform.
* `private_ip` - The private IP address of instance VNIC. To set the private IP address, use the `private_ip` argument in create_vnic_details.
* `public_ip` - The public IP address of instance VNIC (if enabled).
* `region` - The region that contains the availability domain the instance is running in.  Example: `phx` 
* `shape` - The shape of the instance. The shape determines the number of CPUs and the amount of memory allocated to the instance. You can enumerate all available shapes by calling [ListShapes](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Shape/ListShapes). 
* `source_details` - Details for creating an instance
	* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. Minimum value is 50 GB and maximum value is 16384 GB (16TB). This should only be specified when `source_type` is `image`.
	* `source_id` - The OCID of an image or a boot volume to use, depending on the value of `source_type`.
	* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
* `state` - The current state of the instance.
* `time_created` - The date and time the instance was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

## Import

Instances can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance.test_instance "id"
```
