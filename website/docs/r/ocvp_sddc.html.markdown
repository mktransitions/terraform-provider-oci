---
subcategory: "Ocvp"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_sddc"
sidebar_current: "docs-oci-resource-ocvp-sddc"
description: |-
  Provides the Sddc resource in Oracle Cloud Infrastructure Ocvp service
---

# oci_ocvp_sddc
This resource provides the Sddc resource in Oracle Cloud Infrastructure Ocvp service.

Creates a software-defined data center (SDDC).

Use the [WorkRequest](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/WorkRequest/) operations to track the
creation of the SDDC.


## Example Usage

```hcl
resource "oci_ocvp_sddc" "test_sddc" {
	#Required
	compartment_id = var.compartment_id
	compute_availability_domain = var.sddc_compute_availability_domain
	esxi_hosts_count = var.sddc_esxi_hosts_count
	nsx_edge_uplink1vlan_id = oci_core_vlan.test_nsx_edge_uplink1vlan.id
	nsx_edge_uplink2vlan_id = oci_core_vlan.test_nsx_edge_uplink2vlan.id
	nsx_edge_vtep_vlan_id = oci_core_vlan.test_nsx_edge_vtep_vlan.id
	nsx_vtep_vlan_id = oci_core_vlan.test_nsx_vtep_vlan.id
	provisioning_subnet_id = oci_core_subnet.test_subnet.id
	ssh_authorized_keys = var.sddc_ssh_authorized_keys
	vmotion_vlan_id = oci_core_vlan.test_vmotion_vlan.id
	vmware_software_version = var.sddc_vmware_software_version
	vsan_vlan_id = oci_core_vlan.test_vsan_vlan.id
	vsphere_vlan_id = oci_core_vlan.test_vsphere_vlan.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.sddc_display_name
	freeform_tags = {"Department"= "Finance"}
	instance_display_name_prefix = var.sddc_instance_display_name_prefix
	workload_network_cidr = var.sddc_workload_network_cidr
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the SDDC. 
* `compute_availability_domain` - (Required) The availability domain to create the SDDC's ESXi hosts in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A descriptive name for the SDDC. It must be unique, start with a letter, and contain only letters, digits, whitespaces, dashes and underscores. Avoid entering confidential information. 
* `esxi_hosts_count` - (Required) The number of ESXi hosts to create in the SDDC. You can add more hosts later (see [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/EsxiHost/CreateEsxiHost)).

	**Note:** If you later delete EXSi hosts from the SDDC to total less than 3, you are still billed for the 3 minimum recommended EXSi hosts. Also, you cannot add more VMware workloads to the SDDC until it again has at least 3 ESXi hosts. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `instance_display_name_prefix` - (Optional) A prefix used in the name of each ESXi host and Compute instance in the SDDC. If this isn't set, the SDDC's `displayName` is used as the prefix.

	For example, if the value is `mySDDC`, the ESXi hosts are named `mySDDC-1`, `mySDDC-2`, and so on. 
* `nsx_edge_uplink1vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the NSX Edge Uplink 1 component of the VMware environment. 
* `nsx_edge_uplink2vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the NSX Edge Uplink 2 component of the VMware environment. 
* `nsx_edge_vtep_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the NSX Edge VTEP component of the VMware environment. 
* `nsx_vtep_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the NSX VTEP component of the VMware environment. 
* `provisioning_subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management subnet to use for provisioning the SDDC. 
* `ssh_authorized_keys` - (Required) (Updatable) One or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on each ESXi host. Use a newline character to separate multiple keys. The SSH keys must be in the format required for the `authorized_keys` file 
* `vmotion_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the vMotion component of the VMware environment. 
* `vmware_software_version` - (Required) (Updatable) The VMware software bundle to install on the ESXi hosts in the SDDC. To get a list of the available versions, use [ListSupportedVmwareSoftwareVersions](#/en/ocvs/20200501/SupportedVmwareSoftwareVersionSummary/ ListSupportedVmwareSoftwareVersions). 
* `vsan_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the vSAN component of the VMware environment. 
* `vsphere_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the vSphere component of the VMware environment. 
* `workload_network_cidr` - (Optional) The CIDR block for the IP addresses that VMware VMs in the SDDC use to run application workloads. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the SDDC. 
* `compute_availability_domain` - The availability domain the ESXi hosts are running in.  Example: `Uocm:PHX-AD-1` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A descriptive name for the SDDC. It must be unique, start with a letter, and contain only letters, digits, whitespaces, dashes and underscores. Avoid entering confidential information. 
* `esxi_hosts_count` - The number of ESXi hosts in the SDDC.
* `actual_esxi_hosts_count` - The number of actual ESXi hosts in the SDDC on the cloud. This attribute will be different when esxi Host is added to an existing SDDC.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC. 
* `instance_display_name_prefix` - A prefix used in the name of each ESXi host and Compute instance in the SDDC. If this isn't set, the SDDC's `displayName` is used as the prefix.

	For example, if the value is `MySDDC`, the ESXi hosts are named `MySDDC-1`, `MySDDC-2`, and so on. 
* `nsx_edge_uplink1vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the NSX Edge Uplink 1 component of the VMware environment.

	This attribute is not guaranteed to reflect the NSX Edge Uplink 1 VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the NSX Edge Uplink 1 VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the NSX Edge Uplink 1 component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/Sddc/UpdateSddc) to update the SDDC's `nsxEdgeUplink1VlanId` with that new VLAN's OCID. 
* `nsx_edge_uplink2vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the NSX Edge Uplink 2 component of the VMware environment.

	This attribute is not guaranteed to reflect the NSX Edge Uplink 2 VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the NSX Edge Uplink 2 VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the NSX Edge Uplink 2 component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/Sddc/UpdateSddc) to update the SDDC's `nsxEdgeUplink2VlanId` with that new VLAN's OCID. 
* `nsx_edge_uplink_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is the virtual IP (VIP) for the NSX Edge Uplink. Use this OCID as the route target for route table rules when setting up connectivity between the SDDC and other networks. For information about `PrivateIp` objects, see the Core Services API. 
* `nsx_edge_vtep_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the NSX Edge VTEP component of the VMware environment.

	This attribute is not guaranteed to reflect the NSX Edge VTEP VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the NSX Edge VTEP VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the NSX Edge VTEP component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/Sddc/UpdateSddc) to update the SDDC's `nsxEdgeVTepVlanId` with that new VLAN's OCID. 
* `nsx_manager_fqdn` - FQDN for NSX Manager  Example: `nsx-my-sddc.sddc.us-phoenix-1.oraclecloud.com` 
* `nsx_manager_initial_password` - The SDDC includes an administrator username and initial password for NSX Manager. Make sure to change this initial NSX Manager password to a different value. 
* `nsx_manager_private_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is the virtual IP (VIP) for NSX Manager. For information about `PrivateIp` objects, see the Core Services API. 
* `nsx_manager_username` - The SDDC includes an administrator username and initial password for NSX Manager. You can change this initial username to a different value in NSX Manager. 
* `nsx_overlay_segment_name` - The VMware NSX overlay workload segment to host your application. Connect to workload portgroup in vCenter to access this overlay segment. 
* `nsx_vtep_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the NSX VTEP component of the VMware environment.

	This attribute is not guaranteed to reflect the NSX VTEP VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the NSX VTEP VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the NSX VTEP component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/Sddc/UpdateSddc) to update the SDDC's `nsxVTepVlanId` with that new VLAN's OCID. 
* `provisioning_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management subnet used to provision the SDDC. 
* `ssh_authorized_keys` - One or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on each ESXi host. Use a newline character to separate multiple keys. The SSH keys must be in the format required for the `authorized_keys` file.

	This attribute is not guaranteed to reflect the public SSH keys currently installed on the ESXi hosts in the SDDC. The purpose of this attribute is to show the public SSH keys that Oracle Cloud VMware Solution will install on any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you upgrade the existing ESXi hosts in the SDDC to use different SSH keys, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/Sddc/UpdateSddc) to update the SDDC's `sshAuthorizedKeys` with the new public keys. 
* `state` - The current state of the SDDC.
* `time_created` - The date and time the SDDC was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the SDDC was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `vcenter_fqdn` - FQDN for vCenter  Example: `vcenter-my-sddc.sddc.us-phoenix-1.oraclecloud.com` 
* `vcenter_initial_password` - The SDDC includes an administrator username and initial password for vCenter. Make sure to change this initial vCenter password to a different value. 
* `vcenter_private_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is the virtual IP (VIP) for vCenter. For information about `PrivateIp` objects, see the Core Services API. 
* `vcenter_username` - The SDDC includes an administrator username and initial password for vCenter. You can change this initial username to a different value in vCenter. 
* `vmotion_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the vMotion component of the VMware environment.

	This attribute is not guaranteed to reflect the vMotion VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the vMotion VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the vMotion component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/Sddc/UpdateSddc) to update the SDDC's `vmotionVlanId` with that new VLAN's OCID. 
* `vmware_software_version` - In general, this is a specific version of bundled VMware software supported by Oracle Cloud VMware Solution (see [ListSupportedVmwareSoftwareVersions](#/en/ocvs/20200501/SupportedVmwareSoftwareVersionSummary/ ListSupportedVmwareSoftwareVersions)).

	This attribute is not guaranteed to reflect the version of software currently installed on the ESXi hosts in the SDDC. The purpose of this attribute is to show the version of software that the Oracle Cloud VMware Solution will install on any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you upgrade the existing ESXi hosts in the SDDC to use a newer version of bundled VMware software supported by the Oracle Cloud VMware Solution, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/Sddc/UpdateSddc) to update the SDDC's `vmwareSoftwareVersion` with that new version. 
* `vsan_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the vSAN component of the VMware environment.

	This attribute is not guaranteed to reflect the vSAN VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the vSAN VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the vSAN component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/Sddc/UpdateSddc) to update the SDDC's `vsanVlanId` with that new VLAN's OCID. 
* `vsphere_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the vSphere component of the VMware environment.

	This attribute is not guaranteed to reflect the vSphere VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the vSphere VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the vSphere component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/ocvs/20200501/Sddc/UpdateSddc) to update the SDDC's `vsphereVlanId` with that new VLAN's OCID. 
* `workload_network_cidr` - The CIDR block for the IP addresses that VMware VMs in the SDDC use to run application workloads. 

## Import

Sddcs can be imported using the `id`, e.g.

```
$ terraform import oci_ocvp_sddc.test_sddc "id"
```

