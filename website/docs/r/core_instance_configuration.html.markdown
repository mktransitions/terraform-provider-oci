---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_configuration"
sidebar_current: "docs-oci-resource-core-instance_configuration"
description: |-
  Provides the Instance Configuration resource in Oracle Cloud Infrastructure Core service
---

# oci_core_instance_configuration
This resource provides the Instance Configuration resource in Oracle Cloud Infrastructure Core service.

Creates an instance configuration


## Example Usage

```hcl
resource "oci_core_instance_configuration" "test_instance_configuration" {
	#Required
	compartment_id = "${var.compartment_id}"
	instance_details {
		#Required
		instance_type = "${var.instance_configuration_instance_details_instance_type}"

		#Optional
		block_volumes {

			#Optional
			attach_details {
				#Required
				type = "${var.instance_configuration_instance_details_block_volumes_attach_details_type}"

				#Optional
				display_name = "${var.instance_configuration_instance_details_block_volumes_attach_details_display_name}"
				is_read_only = "${var.instance_configuration_instance_details_block_volumes_attach_details_is_read_only}"
				use_chap = "${var.instance_configuration_instance_details_block_volumes_attach_details_use_chap}"
			}
			create_details {

				#Optional
				availability_domain = "${var.instance_configuration_instance_details_block_volumes_create_details_availability_domain}"
				backup_policy_id = "${oci_core_backup_policy.test_backup_policy.id}"
				compartment_id = "${var.compartment_id}"
				defined_tags = {"Operations.CostCenter"= "42"}
				display_name = "${var.instance_configuration_instance_details_block_volumes_create_details_display_name}"
				freeform_tags = {"Department"= "Finance"}
				size_in_gbs = "${var.instance_configuration_instance_details_block_volumes_create_details_size_in_gbs}"
				source_details {
					#Required
					type = "${var.instance_configuration_instance_details_block_volumes_create_details_source_details_type}"

					#Optional
					id = "${var.instance_configuration_instance_details_block_volumes_create_details_source_details_id}"
				}
			}
			volume_id = "${oci_core_volume.test_volume.id}"
		}
		launch_details {

			#Optional
			availability_domain = "${var.instance_configuration_instance_details_launch_details_availability_domain}"
			compartment_id = "${var.compartment_id}"
			create_vnic_details {

				#Optional
				assign_public_ip = "${var.instance_configuration_instance_details_launch_details_create_vnic_details_assign_public_ip}"
				display_name = "${var.instance_configuration_instance_details_launch_details_create_vnic_details_display_name}"
				hostname_label = "${var.instance_configuration_instance_details_launch_details_create_vnic_details_hostname_label}"
				private_ip = "${var.instance_configuration_instance_details_launch_details_create_vnic_details_private_ip}"
				skip_source_dest_check = "${var.instance_configuration_instance_details_launch_details_create_vnic_details_skip_source_dest_check}"
				subnet_id = "${oci_core_subnet.test_subnet.id}"
			}
			defined_tags = {"Operations.CostCenter"= "42"}
			display_name = "${var.instance_configuration_instance_details_launch_details_display_name}"
			extended_metadata = "${var.instance_configuration_instance_details_launch_details_extended_metadata}"
			freeform_tags = {"Department"= "Finance"}
			ipxe_script = "${var.instance_configuration_instance_details_launch_details_ipxe_script}"
			metadata = "${var.instance_configuration_instance_details_launch_details_metadata}"
			shape = "${var.instance_configuration_instance_details_launch_details_shape}"
			source_details {
				#Required
				source_type = "${var.instance_configuration_instance_details_launch_details_source_details_source_type}"

				#Optional
				boot_volume_id = "${oci_core_boot_volume.test_boot_volume.id}"
				image_id = "${oci_core_image.test_image.id}"
			}
		}
		secondary_vnics {

			#Optional
			create_vnic_details {

				#Optional
				assign_public_ip = "${var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_assign_public_ip}"
				display_name = "${var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_display_name}"
				hostname_label = "${var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_hostname_label}"
				private_ip = "${var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_private_ip}"
				skip_source_dest_check = "${var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_skip_source_dest_check}"
				subnet_id = "${oci_core_subnet.test_subnet.id}"
			}
			display_name = "${var.instance_configuration_instance_details_secondary_vnics_display_name}"
			nic_index = "${var.instance_configuration_instance_details_secondary_vnics_nic_index}"
		}
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.instance_configuration_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the instance configuration. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name for the instance configuration 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `instance_details` - (Required) 
	* `block_volumes` - (Optional) 
		* `attach_details` - (Optional) 
			* `display_name` - (Applicable when instance_type=compute) A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information. 
			* `is_read_only` - (Applicable when instance_type=compute) Whether the attachment should be created in read-only mode.
			* `type` - (Required) The type of volume. The only supported values are "iscsi" and "paravirtualized".
			* `use_chap` - (Applicable when type=iscsi) Whether to use CHAP authentication for the volume attachment. Defaults to false.
		* `create_details` - (Optional) 
			* `availability_domain` - (Optional) The availability domain of the volume.  Example: `Uocm:PHX-AD-1` 
			* `backup_policy_id` - (Optional) If provided, specifies the ID of the volume backup policy to assign to the newly created volume. If omitted, no policy will be assigned. 
			* `compartment_id` - (Optional) The OCID of the compartment that contains the volume.
			* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `size_in_gbs` - (Optional) The size of the volume in GBs.
			* `source_details` - (Optional) Specifies the volume source details for a new Block volume. The volume source is either another Block volume in the same availability domain or a Block volume backup. This is an optional field. If not specified or set to null, the new Block volume will be empty. When specified, the new Block volume will contain data from the source volume or backup. 
				* `id` - (Optional) The OCID of the volume backup.
				* `type` - (Required) The type can be one of these values: `volume`, `volumeBackup`
		* `volume_id` - (Optional) The OCID of the volume.
	* `instance_type` - (Required) The type of instance details. Supported instanceType is compute 
	* `launch_details` - (Optional) 
		* `availability_domain` - (Optional) The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
		* `compartment_id` - (Optional) The OCID of the compartment.
		* `create_vnic_details` - (Optional) Details for the primary VNIC, which is automatically created and attached when the instance is launched. 
			* `assign_public_ip` - (Optional) 
			* `display_name` - (Optional) A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
			* `hostname_label` - (Optional) 
			* `private_ip` - (Optional) 
			* `skip_source_dest_check` - (Optional) 
			* `subnet_id` - (Optional) 
		* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
		* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My bare metal instance` 
		* `extended_metadata` - (Optional) Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the 'metadata' object.

			They are distinguished from 'metadata' fields in that these can be nested JSON objects (whereas 'metadata' fields are string/string maps only). 
		* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
		* `ipxe_script` - (Optional) This is an advanced option.

			When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

			If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

			The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

			For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

			For more information about iPXE, see http://ipxe.org. 
		* `metadata` - (Optional) Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.

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
		* `shape` - (Optional) The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

			You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Shape/ListShapes). 
		* `source_details` - (Optional) Details for creating an instance. Use this parameter to specify whether a boot volume or an image should be used to launch a new instance. 
			* `boot_volume_id` - (Applicable when source_type=bootVolume) The OCID of the boot volume used to boot the instance.
			* `image_id` - (Applicable when source_type=image) The OCID of the image used to boot the instance.
			* `source_type` - (Required) The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
	* `secondary_vnics` - (Optional) 
		* `create_vnic_details` - (Optional) Details for creating a new VNIC. 
			* `assign_public_ip` - (Optional) 
			* `display_name` - (Optional) A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
			* `hostname_label` - (Optional) 
			* `private_ip` - (Optional) 
			* `skip_source_dest_check` - (Optional) 
			* `subnet_id` - (Optional) 
		* `display_name` - (Optional) A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed. 
		* `nic_index` - (Optional) Which physical network interface card (NIC) the VNIC will use. Defaults to 0. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the instance configuration. 
* `deferred_fields` - 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the instance configuration 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the instance configuration
* `instance_details` - 
	* `block_volumes` - 
		* `attach_details` - 
			* `display_name` - A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information. 
			* `is_read_only` - Whether the attachment should be created in read-only mode.
			* `type` - The type of volume. The only supported values are "iscsi" and "paravirtualized".
			* `use_chap` - Whether to use CHAP authentication for the volume attachment. Defaults to false.
		* `create_details` - 
			* `availability_domain` - The availability domain of the volume.  Example: `Uocm:PHX-AD-1` 
			* `backup_policy_id` - If provided, specifies the ID of the volume backup policy to assign to the newly created volume. If omitted, no policy will be assigned. 
			* `compartment_id` - The OCID of the compartment that contains the volume.
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `size_in_gbs` - The size of the volume in GBs.
			* `source_details` - Specifies the volume source details for a new Block volume. The volume source is either another Block volume in the same availability domain or a Block volume backup. This is an optional field. If not specified or set to null, the new Block volume will be empty. When specified, the new Block volume will contain data from the source volume or backup. 
				* `id` - The OCID of the volume backup.
				* `type` - The type can be one of these values: `volume`, `volumeBackup`
		* `volume_id` - The OCID of the volume.
	* `instance_type` - The type of instance details. Supported instanceType is compute 
	* `launch_details` - 
		* `availability_domain` - The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
		* `compartment_id` - The OCID of the compartment.
		* `create_vnic_details` - Details for the primary VNIC, which is automatically created and attached when the instance is launched. 
			* `assign_public_ip` - 
			* `display_name` - A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
			* `hostname_label` - 
			* `private_ip` - 
			* `skip_source_dest_check` - 
			* `subnet_id` - 
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My bare metal instance` 
		* `extended_metadata` - Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the 'metadata' object.

			They are distinguished from 'metadata' fields in that these can be nested JSON objects (whereas 'metadata' fields are string/string maps only). 
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
		* `ipxe_script` - This is an advanced option.

			When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

			If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

			The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

			For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

			For more information about iPXE, see http://ipxe.org. 
		* `metadata` - Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.

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
		* `shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

			You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Shape/ListShapes). 
		* `source_details` - Details for creating an instance. Use this parameter to specify whether a boot volume or an image should be used to launch a new instance. 
			* `boot_volume_id` - The OCID of the boot volume used to boot the instance.
			* `image_id` - The OCID of the image used to boot the instance.
			* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
	* `secondary_vnics` - 
		* `create_vnic_details` - Details for creating a new VNIC. 
			* `assign_public_ip` - 
			* `display_name` - A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
			* `hostname_label` - 
			* `private_ip` - 
			* `skip_source_dest_check` - 
			* `subnet_id` - 
		* `display_name` - A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed. 
		* `nic_index` - Which physical network interface card (NIC) the VNIC will use. Defaults to 0. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
* `time_created` - The date and time the instance configuration was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

## Import

InstanceConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance_configuration.test_instance_configuration "id"
```

