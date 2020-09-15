// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v25/ocvp"
)

func init() {
	RegisterDatasource("oci_ocvp_sddc", OcvpSddcDataSource())
}

func OcvpSddcDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["sddc_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(OcvpSddcResource(), fieldMap, readSingularOcvpSddc)
}

func readSingularOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).sddcClient()

	return ReadResource(sync)
}

type OcvpSddcDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.SddcClient
	Res    *oci_ocvp.GetSddcResponse
}

func (s *OcvpSddcDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpSddcDataSourceCrud) Get() error {
	request := oci_ocvp.GetSddcRequest{}

	if sddcId, ok := s.D.GetOkExists("sddc_id"); ok {
		tmp := sddcId.(string)
		request.SddcId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "ocvp")

	response, err := s.Client.GetSddc(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OcvpSddcDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeAvailabilityDomain != nil {
		s.D.Set("compute_availability_domain", *s.Res.ComputeAvailabilityDomain)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EsxiHostsCount != nil {
		s.D.Set("esxi_hosts_count", *s.Res.EsxiHostsCount)
		s.D.Set("actual_esxi_hosts_count", *s.Res.EsxiHostsCount)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceDisplayNamePrefix != nil {
		s.D.Set("instance_display_name_prefix", *s.Res.InstanceDisplayNamePrefix)
	}

	if s.Res.NsxEdgeUplink1VlanId != nil {
		s.D.Set("nsx_edge_uplink1vlan_id", *s.Res.NsxEdgeUplink1VlanId)
	}

	if s.Res.NsxEdgeUplink2VlanId != nil {
		s.D.Set("nsx_edge_uplink2vlan_id", *s.Res.NsxEdgeUplink2VlanId)
	}

	if s.Res.NsxEdgeUplinkIpId != nil {
		s.D.Set("nsx_edge_uplink_ip_id", *s.Res.NsxEdgeUplinkIpId)
	}

	if s.Res.NsxEdgeVTepVlanId != nil {
		s.D.Set("nsx_edge_vtep_vlan_id", *s.Res.NsxEdgeVTepVlanId)
	}

	if s.Res.NsxManagerFqdn != nil {
		s.D.Set("nsx_manager_fqdn", *s.Res.NsxManagerFqdn)
	}

	if s.Res.NsxManagerInitialPassword != nil {
		s.D.Set("nsx_manager_initial_password", *s.Res.NsxManagerInitialPassword)
	}

	if s.Res.NsxManagerPrivateIpId != nil {
		s.D.Set("nsx_manager_private_ip_id", *s.Res.NsxManagerPrivateIpId)
	}

	if s.Res.NsxManagerUsername != nil {
		s.D.Set("nsx_manager_username", *s.Res.NsxManagerUsername)
	}

	if s.Res.NsxOverlaySegmentName != nil {
		s.D.Set("nsx_overlay_segment_name", *s.Res.NsxOverlaySegmentName)
	}

	if s.Res.NsxVTepVlanId != nil {
		s.D.Set("nsx_vtep_vlan_id", *s.Res.NsxVTepVlanId)
	}

	if s.Res.ProvisioningSubnetId != nil {
		s.D.Set("provisioning_subnet_id", *s.Res.ProvisioningSubnetId)
	}

	if s.Res.SshAuthorizedKeys != nil {
		s.D.Set("ssh_authorized_keys", *s.Res.SshAuthorizedKeys)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VcenterFqdn != nil {
		s.D.Set("vcenter_fqdn", *s.Res.VcenterFqdn)
	}

	if s.Res.VcenterInitialPassword != nil {
		s.D.Set("vcenter_initial_password", *s.Res.VcenterInitialPassword)
	}

	if s.Res.VcenterPrivateIpId != nil {
		s.D.Set("vcenter_private_ip_id", *s.Res.VcenterPrivateIpId)
	}

	if s.Res.VcenterUsername != nil {
		s.D.Set("vcenter_username", *s.Res.VcenterUsername)
	}

	if s.Res.VmotionVlanId != nil {
		s.D.Set("vmotion_vlan_id", *s.Res.VmotionVlanId)
	}

	if s.Res.VmwareSoftwareVersion != nil {
		s.D.Set("vmware_software_version", *s.Res.VmwareSoftwareVersion)
	}

	if s.Res.VsanVlanId != nil {
		s.D.Set("vsan_vlan_id", *s.Res.VsanVlanId)
	}

	if s.Res.VsphereVlanId != nil {
		s.D.Set("vsphere_vlan_id", *s.Res.VsphereVlanId)
	}

	if s.Res.WorkloadNetworkCidr != nil {
		s.D.Set("workload_network_cidr", *s.Res.WorkloadNetworkCidr)
	}

	return nil
}
