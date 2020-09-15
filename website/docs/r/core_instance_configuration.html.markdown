---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_configuration"
sidebar_current: "docs-oci-resource-core-instance_configuration"
description: |-
  Provides the Instance Configuration resource in Oracle Cloud Infrastructure Core service
---

# oci_core_instance_configuration
This resource provides the Instance Configuration resource in Oracle Cloud Infrastructure Core service.

Creates an instance configuration. An instance configuration is a template that defines the
settings to use when creating Compute instances.


## Example Usage

```hcl
resource "oci_core_instance_configuration" "test_instance_configuration" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.instance_configuration_display_name
	freeform_tags = {"Department"= "Finance"}
	instance_details {
		#Required
		instance_type = var.instance_configuration_instance_details_instance_type

		#Optional
		block_volumes {

			#Optional
			attach_details {
				#Required
				type = var.instance_configuration_instance_details_block_volumes_attach_details_type

				#Optional
				device = var.instance_configuration_instance_details_block_volumes_attach_details_device
				display_name = var.instance_configuration_instance_details_block_volumes_attach_details_display_name
				is_pv_encryption_in_transit_enabled = var.instance_configuration_instance_details_block_volumes_attach_details_is_pv_encryption_in_transit_enabled
				is_read_only = var.instance_configuration_instance_details_block_volumes_attach_details_is_read_only
				is_shareable = var.instance_configuration_instance_details_block_volumes_attach_details_is_shareable
				use_chap = var.instance_configuration_instance_details_block_volumes_attach_details_use_chap
			}
			create_details {

				#Optional
				availability_domain = var.instance_configuration_instance_details_block_volumes_create_details_availability_domain
				backup_policy_id = data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id
				compartment_id = var.compartment_id
				defined_tags = {"Operations.CostCenter"= "42"}
				display_name = var.instance_configuration_instance_details_block_volumes_create_details_display_name
				freeform_tags = {"Department"= "Finance"}
				kms_key_id = oci_kms_key.test_key.id
				size_in_gbs = var.instance_configuration_instance_details_block_volumes_create_details_size_in_gbs
				source_details {
					#Required
					type = var.instance_configuration_instance_details_block_volumes_create_details_source_details_type

					#Optional
					id = var.instance_configuration_instance_details_block_volumes_create_details_source_details_id
				}
				vpus_per_gb = var.instance_configuration_instance_details_block_volumes_create_details_vpus_per_gb
			}
			volume_id = oci_core_volume.test_volume.id
		}
		launch_details {

			#Optional
			agent_config {

				#Optional
				is_management_disabled = var.instance_configuration_instance_details_launch_details_agent_config_is_management_disabled
				is_monitoring_disabled = var.instance_configuration_instance_details_launch_details_agent_config_is_monitoring_disabled
			}
			availability_config {

				#Optional
				recovery_action = var.instance_configuration_instance_details_launch_details_availability_config_recovery_action
			}
			availability_domain = var.instance_configuration_instance_details_launch_details_availability_domain
			compartment_id = var.compartment_id
			create_vnic_details {

				#Optional
				assign_public_ip = var.instance_configuration_instance_details_launch_details_create_vnic_details_assign_public_ip
				defined_tags = {"Operations.CostCenter"= "42"}
				display_name = var.instance_configuration_instance_details_launch_details_create_vnic_details_display_name
				freeform_tags = {"Department"= "Finance"}
				hostname_label = var.instance_configuration_instance_details_launch_details_create_vnic_details_hostname_label
				nsg_ids = var.instance_configuration_instance_details_launch_details_create_vnic_details_nsg_ids
				private_ip = var.instance_configuration_instance_details_launch_details_create_vnic_details_private_ip
				skip_source_dest_check = var.instance_configuration_instance_details_launch_details_create_vnic_details_skip_source_dest_check
				subnet_id = oci_core_subnet.test_subnet.id
			}
			dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id
			defined_tags = {"Operations.CostCenter"= "42"}
			display_name = var.instance_configuration_instance_details_launch_details_display_name
			extended_metadata = var.instance_configuration_instance_details_launch_details_extended_metadata
			fault_domain = var.instance_configuration_instance_details_launch_details_fault_domain
			freeform_tags = {"Department"= "Finance"}
			ipxe_script = var.instance_configuration_instance_details_launch_details_ipxe_script
			is_pv_encryption_in_transit_enabled = var.instance_configuration_instance_details_launch_details_is_pv_encryption_in_transit_enabled
			launch_mode = var.instance_configuration_instance_details_launch_details_launch_mode
			launch_options {

				#Optional
				boot_volume_type = var.instance_configuration_instance_details_launch_details_launch_options_boot_volume_type
				firmware = var.instance_configuration_instance_details_launch_details_launch_options_firmware
				is_consistent_volume_naming_enabled = var.instance_configuration_instance_details_launch_details_launch_options_is_consistent_volume_naming_enabled
				is_pv_encryption_in_transit_enabled = var.instance_configuration_instance_details_launch_details_launch_options_is_pv_encryption_in_transit_enabled
				network_type = var.instance_configuration_instance_details_launch_details_launch_options_network_type
				remote_data_volume_type = var.instance_configuration_instance_details_launch_details_launch_options_remote_data_volume_type
			}
			metadata = var.instance_configuration_instance_details_launch_details_metadata
			preferred_maintenance_action = var.instance_configuration_instance_details_launch_details_preferred_maintenance_action
			shape = var.instance_configuration_instance_details_launch_details_shape
			shape_config {

				#Optional
				ocpus = var.instance_configuration_instance_details_launch_details_shape_config_ocpus
			}
			source_details {
				#Required
				source_type = var.instance_configuration_instance_details_launch_details_source_details_source_type

				#Optional
				boot_volume_id = oci_core_boot_volume.test_boot_volume.id
				boot_volume_size_in_gbs = var.instance_configuration_instance_details_launch_details_source_details_boot_volume_size_in_gbs
				image_id = oci_core_image.test_image.id
			}
		}
		secondary_vnics {

			#Optional
			create_vnic_details {

				#Optional
				assign_public_ip = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_assign_public_ip
				defined_tags = {"Operations.CostCenter"= "42"}
				display_name = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_display_name
				freeform_tags = {"Department"= "Finance"}
				hostname_label = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_hostname_label
				nsg_ids = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_nsg_ids
				private_ip = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_private_ip
				skip_source_dest_check = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_skip_source_dest_check
				subnet_id = oci_core_subnet.test_subnet.id
			}
			display_name = var.instance_configuration_instance_details_secondary_vnics_display_name
			nic_index = var.instance_configuration_instance_details_secondary_vnics_nic_index
		}
	}
	instance_id = oci_core_instance.test_instance.id
	source = var.instance_configuration_source
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the instance configuration. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name for the instance configuration.  Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `instance_details` - (Required when source=NONE) 
	* `block_volumes` - (Optional) 
		* `attach_details` - (Optional) 
			* `device` - (Applicable when instance_type=compute) The device name.
			* `display_name` - (Applicable when instance_type=compute) A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information. 
			* `is_pv_encryption_in_transit_enabled` - (Applicable when type=paravirtualized) Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
			* `is_read_only` - (Applicable when instance_type=compute) Whether the attachment should be created in read-only mode.
			* `is_shareable` - (Applicable when instance_type=compute) Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified. 
			* `type` - (Required) The type of volume. The only supported values are "iscsi" and "paravirtualized".
			* `use_chap` - (Applicable when type=iscsi) Whether to use CHAP authentication for the volume attachment. Defaults to false.
		* `create_details` - (Optional) 
			* `availability_domain` - (Optional) The availability domain of the volume.  Example: `Uocm:PHX-AD-1` 
			* `backup_policy_id` - (Optional) If provided, specifies the ID of the volume backup policy to assign to the newly created volume. If omitted, no policy will be assigned. 
			* `compartment_id` - (Optional) The OCID of the compartment that contains the volume.
			* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `kms_key_id` - (Optional) The OCID of the Key Management key to assign as the master encryption key for the volume. 
			* `size_in_gbs` - (Optional) The size of the volume in GBs.
			* `source_details` - (Optional) Specifies the volume source details for a new Block volume. The volume source is either another Block volume in the same availability domain or a Block volume backup. This is an optional field. If not specified or set to null, the new Block volume will be empty. When specified, the new Block volume will contain data from the source volume or backup. 
				* `id` - (Optional) The OCID of the volume backup.
				* `type` - (Required) The type can be one of these values: `volume`, `volumeBackup`
			* `vpus_per_gb` - (Optional) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.

				Allowed values:
				* `0`: Represents Lower Cost option.
				* `10`: Represents Balanced option.
				* `20`: Represents Higher Performance option. 
		* `volume_id` - (Optional) The OCID of the volume.
	* `instance_type` - (Required) The type of instance details. Supported instanceType is compute 
	* `launch_details` - (Optional) 
		* `agent_config` - (Optional) 
			* `is_management_disabled` - (Optional) Whether the agent running on the instance can run all the available management plugins. Default value is false. 
			* `is_monitoring_disabled` - (Optional) Whether the agent running on the instance can gather performance metrics and monitor the instance. Default value is false. 
		* `availability_config` - (Optional) 
			* `recovery_action` - (Optional) Actions customers can specify that would be applied to their instances after scheduled or unexpected host maintenance.
				* `RESTORE_INSTANCE` - This would be the default action if recoveryAction is not set. VM instances will be restored to the power state it was in before maintenance.
				* `STOP_INSTANCE` - This action allow customers to have their VM instances be stopped after maintenance. 
		* `availability_domain` - (Optional) The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
		* `compartment_id` - (Optional) The OCID of the compartment.
		* `create_vnic_details` - (Optional) Details for the primary VNIC, which is automatically created and attached when the instance is launched. 
			* `assign_public_ip` - (Optional) Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - (Optional) A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
			* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `hostname_label` - (Optional) The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `nsg_ids` - (Optional) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
			* `private_ip` - (Optional) A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `skip_source_dest_check` - (Optional) Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `subnet_id` - (Optional) The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
		* `dedicated_vm_host_id` - (Optional) The OCID of dedicated VM host.

			Dedicated VM hosts can be used when launching individual instances from an instance configuration. They cannot be used to launch instance pools. 
		* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
		* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My bare metal instance` 
		* `extended_metadata` - (Optional) Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

			They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only).

			The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes. 
		* `fault_domain` - (Optional) A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

			If you do not specify the fault domain, the system selects one for you.

			 To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

			Example: `FAULT-DOMAIN-1` 
		* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
		* `ipxe_script` - (Optional) This is an advanced option.

			When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

			If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

			The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

			For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

			For more information about iPXE, see http://ipxe.org. 
		* `is_pv_encryption_in_transit_enabled` - (Optional) Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
		* `launch_mode` - (Optional) Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
			* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for Oracle-provided images.
			* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
			* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
			* `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter. 
		* `launch_options` - (Optional) Options for tuning the compatibility and performance of VM shapes. The values that you specify override any default values. 
			* `boot_volume_type` - (Optional) Emulation type for the boot volume.
				* `ISCSI` - ISCSI attached block storage device.
				* `SCSI` - Emulated SCSI disk.
				* `IDE` - Emulated IDE disk.
				* `VFIO` - Direct attached Virtual Function storage.  This is the default option for local data volumes on Oracle provided images.
				* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on Oracle-provided images. 
			* `firmware` - (Optional) Firmware used to boot VM.  Select the option that matches your operating system.
				* `BIOS` - Boot VM using BIOS style firmware.  This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
				* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems.  This is the default for Oracle-provided images. 
			* `is_consistent_volume_naming_enabled` - (Optional) Whether to enable consistent volume naming feature. Defaults to false.
			* `is_pv_encryption_in_transit_enabled` - (Optional) Deprecated. Instead use `isPvEncryptionInTransitEnabled` in [InstanceConfigurationLaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/datatypes/InstanceConfigurationLaunchInstanceDetails). 
			* `network_type` - (Optional) Emulation type for the physical network interface card (NIC).
				* `E1000` - Emulated Gigabit ethernet controller.  Compatible with Linux e1000 network driver.
				* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
				* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers. 
			* `remote_data_volume_type` - (Optional) Emulation type for volume.
				* `ISCSI` - ISCSI attached block storage device.
				* `SCSI` - Emulated SCSI disk.
				* `IDE` - Emulated IDE disk.
				* `VFIO` - Direct attached Virtual Function storage.  This is the default option for local data volumes on Oracle provided images.
				* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on Oracle-provided images. 
		* `metadata` - (Optional) Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.

			A metadata service runs on every launched instance. The service is an HTTP endpoint listening on 169.254.169.254. You can use the service to:
			* Provide information to [Cloud-Init](https://cloudinit.readthedocs.org/en/latest/) to be used for various system initialization tasks.
			* Get information about the instance, including the custom metadata that you provide when you launch the instance.

			**Providing Cloud-Init Metadata**

			You can use the following metadata key names to provide information to Cloud-Init:

			**"ssh_authorized_keys"** - Provide one or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on the instance. Use a newline character to separate multiple keys. The SSH keys must be in the format necessary for the `authorized_keys` file, as shown in the example below.

			**"user_data"** - Provide your own base64-encoded data to be used by Cloud-Init to run custom scripts or provide custom Cloud-Init configuration. For information about how to take advantage of user data, see the [Cloud-Init Documentation](http://cloudinit.readthedocs.org/en/latest/topics/format.html).

			**Metadata Example**

			"metadata" : { "quake_bot_level" : "Severe", "ssh_authorized_keys" : "ssh-rsa <your_public_SSH_key>== rsa-key-20160227", "user_data" : "<your_public_SSH_key>==" } **Getting Metadata on the Instance**

			To get information about your instance, connect to the instance using SSH and issue any of the following GET requests:

			curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/<any-key-name>

			You'll get back a response that includes all the instance information; only the metadata information; or the metadata information for the specified key name, respectively.

			The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes. 
		* `preferred_maintenance_action` - (Optional) The preferred maintenance action for an instance. The default is LIVE_MIGRATE, if live migration is supported.
			* `LIVE_MIGRATE` - Run maintenance using a live migration.
			* `REBOOT` - Run maintenance using a reboot. 
		* `shape` - (Optional) The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

			You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Shape/ListShapes). 
		* `shape_config` - (Optional) 
			* `ocpus` - (Optional) The total number of OCPUs available to the instance. 
		* `source_details` - (Optional) Details for creating an instance. Use this parameter to specify whether a boot volume or an image should be used to launch a new instance. 
			* `boot_volume_id` - (Applicable when source_type=bootVolume) The OCID of the boot volume used to boot the instance.
			* `boot_volume_size_in_gbs` - (Applicable when source_type=image) The size of the boot volume in GBs. The minimum value is 50 GB and the maximum value is 16384 GB (16TB).
			* `image_id` - (Applicable when source_type=image) The OCID of the image used to boot the instance.
			* `source_type` - (Required) The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
	* `secondary_vnics` - (Optional) 
		* `create_vnic_details` - (Optional) Details for creating a new VNIC. 
			* `assign_public_ip` - (Optional) Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - (Optional) A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
			* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `hostname_label` - (Optional) The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `nsg_ids` - (Optional) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
			* `private_ip` - (Optional) A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `skip_source_dest_check` - (Optional) Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `subnet_id` - (Optional) The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
		* `display_name` - (Optional) A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed. 
		* `nic_index` - (Optional) Which physical network interface card (NIC) the VNIC will use. Defaults to 0. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
* `instance_id` - (Required when source=INSTANCE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance to use to create the instance configuration. 
* `source` - (Optional) The source of the instance configuration. An instance configuration defines the settings to use when creating Compute instances, including details such as the base image, shape, and metadata. You can also specify the associated resources for the instance, such as block volume attachments and network configuration.

	The following values are supported:
	* `NONE`: Creates an instance configuration using the list of settings that you specify.
	* `INSTANCE`: Creates an instance configuration using an existing instance as a template. The instance configuration uses the same settings as the instance. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the instance configuration. 
* `deferred_fields` - Parameters that were not specified when the instance configuration was created, but that are required to launch an instance from the instance configuration. See the [LaunchInstanceConfiguration](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Instance/LaunchInstanceConfiguration) operation. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the instance configuration. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance configuration.
* `instance_details` - 
	* `block_volumes` - 
		* `attach_details` - 
			* `device` - The device name.
			* `display_name` - A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information. 
			* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
			* `is_read_only` - Whether the attachment should be created in read-only mode.
			* `is_shareable` - Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified. 
			* `type` - The type of volume. The only supported values are "iscsi" and "paravirtualized".
			* `use_chap` - Whether to use CHAP authentication for the volume attachment. Defaults to false.
		* `create_details` - 
			* `availability_domain` - The availability domain of the volume.  Example: `Uocm:PHX-AD-1` 
			* `backup_policy_id` - If provided, specifies the ID of the volume backup policy to assign to the newly created volume. If omitted, no policy will be assigned. 
			* `compartment_id` - The OCID of the compartment that contains the volume.
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `kms_key_id` - The OCID of the Key Management key to assign as the master encryption key for the volume. 
			* `size_in_gbs` - The size of the volume in GBs.
			* `source_details` - Specifies the volume source details for a new Block volume. The volume source is either another Block volume in the same availability domain or a Block volume backup. This is an optional field. If not specified or set to null, the new Block volume will be empty. When specified, the new Block volume will contain data from the source volume or backup. 
				* `id` - The OCID of the volume backup.
				* `type` - The type can be one of these values: `volume`, `volumeBackup`
			* `vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.

				Allowed values:
				* `0`: Represents Lower Cost option.
				* `10`: Represents Balanced option.
				* `20`: Represents Higher Performance option. 
		* `volume_id` - The OCID of the volume.
	* `instance_type` - The type of instance details. Supported instanceType is compute 
	* `launch_details` - 
		* `agent_config` - 
			* `is_management_disabled` - Whether the agent running on the instance can run all the available management plugins. Default value is false. 
			* `is_monitoring_disabled` - Whether the agent running on the instance can gather performance metrics and monitor the instance. Default value is false. 
		* `availability_config` - 
			* `recovery_action` - Actions customers can specify that would be applied to their instances after scheduled or unexpected host maintenance.
				* `RESTORE_INSTANCE` - This would be the default action if recoveryAction is not set. VM instances will be restored to the power state it was in before maintenance.
				* `STOP_INSTANCE` - This action allow customers to have their VM instances be stopped after maintenance. 
		* `availability_domain` - The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
		* `compartment_id` - The OCID of the compartment.
		* `create_vnic_details` - Details for the primary VNIC, which is automatically created and attached when the instance is launched. 
			* `assign_public_ip` - Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `hostname_label` - The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
			* `private_ip` - A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `subnet_id` - The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
		* `dedicated_vm_host_id` - The OCID of dedicated VM host.

			Dedicated VM hosts can be used when launching individual instances from an instance configuration. They cannot be used to launch instance pools. 
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My bare metal instance` 
		* `extended_metadata` - Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

			They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only).

			The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes. 
		* `fault_domain` - A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

			If you do not specify the fault domain, the system selects one for you.

			 To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

			Example: `FAULT-DOMAIN-1` 
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
		* `ipxe_script` - This is an advanced option.

			When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

			If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

			The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

			For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

			For more information about iPXE, see http://ipxe.org. 
		* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
		* `launch_mode` - Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
			* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for Oracle-provided images.
			* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
			* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
			* `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter. 
		* `launch_options` - Options for tuning the compatibility and performance of VM shapes. The values that you specify override any default values. 
			* `boot_volume_type` - Emulation type for the boot volume.
				* `ISCSI` - ISCSI attached block storage device.
				* `SCSI` - Emulated SCSI disk.
				* `IDE` - Emulated IDE disk.
				* `VFIO` - Direct attached Virtual Function storage.  This is the default option for local data volumes on Oracle provided images.
				* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on Oracle-provided images. 
			* `firmware` - Firmware used to boot VM.  Select the option that matches your operating system.
				* `BIOS` - Boot VM using BIOS style firmware.  This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
				* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems.  This is the default for Oracle-provided images. 
			* `is_consistent_volume_naming_enabled` - Whether to enable consistent volume naming feature. Defaults to false.
			* `is_pv_encryption_in_transit_enabled` - Deprecated. Instead use `isPvEncryptionInTransitEnabled` in [InstanceConfigurationLaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/datatypes/InstanceConfigurationLaunchInstanceDetails). 
			* `network_type` - Emulation type for the physical network interface card (NIC).
				* `E1000` - Emulated Gigabit ethernet controller.  Compatible with Linux e1000 network driver.
				* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
				* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers. 
			* `remote_data_volume_type` - Emulation type for volume.
				* `ISCSI` - ISCSI attached block storage device.
				* `SCSI` - Emulated SCSI disk.
				* `IDE` - Emulated IDE disk.
				* `VFIO` - Direct attached Virtual Function storage.  This is the default option for local data volumes on Oracle provided images.
				* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on Oracle-provided images. 
		* `metadata` - Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.

			A metadata service runs on every launched instance. The service is an HTTP endpoint listening on 169.254.169.254. You can use the service to:
			* Provide information to [Cloud-Init](https://cloudinit.readthedocs.org/en/latest/) to be used for various system initialization tasks.
			* Get information about the instance, including the custom metadata that you provide when you launch the instance.

			**Providing Cloud-Init Metadata**

			You can use the following metadata key names to provide information to Cloud-Init:

			**"ssh_authorized_keys"** - Provide one or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on the instance. Use a newline character to separate multiple keys. The SSH keys must be in the format necessary for the `authorized_keys` file, as shown in the example below.

			**"user_data"** - Provide your own base64-encoded data to be used by Cloud-Init to run custom scripts or provide custom Cloud-Init configuration. For information about how to take advantage of user data, see the [Cloud-Init Documentation](http://cloudinit.readthedocs.org/en/latest/topics/format.html).

			**Metadata Example**

			"metadata" : { "quake_bot_level" : "Severe", "ssh_authorized_keys" : "ssh-rsa <your_public_SSH_key>== rsa-key-20160227", "user_data" : "<your_public_SSH_key>==" } **Getting Metadata on the Instance**

			To get information about your instance, connect to the instance using SSH and issue any of the following GET requests:

			curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/<any-key-name>

			You'll get back a response that includes all the instance information; only the metadata information; or the metadata information for the specified key name, respectively.

			The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes. 
		* `preferred_maintenance_action` - The preferred maintenance action for an instance. The default is LIVE_MIGRATE, if live migration is supported.
			* `LIVE_MIGRATE` - Run maintenance using a live migration.
			* `REBOOT` - Run maintenance using a reboot. 
		* `shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

			You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Shape/ListShapes). 
		* `shape_config` - 
			* `ocpus` - The total number of OCPUs available to the instance. 
		* `source_details` - Details for creating an instance. Use this parameter to specify whether a boot volume or an image should be used to launch a new instance. 
			* `boot_volume_id` - The OCID of the boot volume used to boot the instance.
			* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. The minimum value is 50 GB and the maximum value is 16384 GB (16TB).
			* `image_id` - The OCID of the image used to boot the instance.
			* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
	* `secondary_vnics` - 
		* `create_vnic_details` - Details for creating a new VNIC. 
			* `assign_public_ip` - Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `hostname_label` - The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
			* `private_ip` - A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
			* `subnet_id` - The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information. 
		* `display_name` - A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed. 
		* `nic_index` - Which physical network interface card (NIC) the VNIC will use. Defaults to 0. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
* `time_created` - The date and time the instance configuration was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Import

InstanceConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance_configuration.test_instance_configuration "id"
```

