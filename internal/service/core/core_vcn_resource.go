// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreVcnResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreVcn,
		Read:     readCoreVcn,
		Update:   updateCoreVcn,
		Delete:   deleteCoreVcn,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"byoipv6cidr_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				// ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"byoipv6range_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"ipv6cidr_block": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"cidr_block": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cidr_blocks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_label": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"ipv6private_cidr_blocks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				// ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_ipv6enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_oracle_gua_allocation_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				// ForceNew: true,
			},
			"security_attributes": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"byoipv6cidr_blocks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"default_dhcp_options_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_route_table_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_security_list_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6cidr_blocks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcn_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreVcn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreVcn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreVcn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreVcn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreVcnResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.Vcn
	DisableNotFoundRetries bool
}

func (s *CoreVcnResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreVcnResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VcnLifecycleStateProvisioning),
	}
}

func (s *CoreVcnResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VcnLifecycleStateAvailable),
	}
}

func (s *CoreVcnResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VcnLifecycleStateTerminating),
	}
}

func (s *CoreVcnResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VcnLifecycleStateTerminated),
	}
}

func (s *CoreVcnResourceCrud) Create() error {
	request := oci_core.CreateVcnRequest{}

	if byoipv6CidrDetails, ok := s.D.GetOkExists("byoipv6cidr_details"); ok {
		interfaces := byoipv6CidrDetails.([]interface{})
		tmp := make([]oci_core.Byoipv6CidrDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "byoipv6cidr_details", stateDataIndex)
			converted, err := s.mapToByoipv6CidrDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("byoipv6cidr_details") {
			request.Byoipv6CidrDetails = tmp
		}
	}

	if cidrBlock, ok := s.D.GetOkExists("cidr_block"); ok {
		tmp := cidrBlock.(string)
		request.CidrBlock = &tmp
	}

	if cidrBlocks, ok := s.D.GetOkExists("cidr_blocks"); ok {
		interfaces := cidrBlocks.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("cidr_blocks") {
			request.CidrBlocks = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if dnsLabel, ok := s.D.GetOkExists("dns_label"); ok {
		tmp := dnsLabel.(string)
		request.DnsLabel = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if ipv6PrivateCidrBlocks, ok := s.D.GetOkExists("ipv6private_cidr_blocks"); ok {
		interfaces := ipv6PrivateCidrBlocks.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ipv6private_cidr_blocks") {
			request.Ipv6PrivateCidrBlocks = tmp
		}
	}

	if isIpv6Enabled, ok := s.D.GetOkExists("is_ipv6enabled"); ok {
		tmp := isIpv6Enabled.(bool)
		request.IsIpv6Enabled = &tmp
	}

	if isOracleGuaAllocationEnabled, ok := s.D.GetOkExists("is_oracle_gua_allocation_enabled"); ok {
		tmp := isOracleGuaAllocationEnabled.(bool)
		request.IsOracleGuaAllocationEnabled = &tmp
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		convertedAttributes := tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		request.SecurityAttributes = convertedAttributes
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVcn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vcn
	return nil
}

func (s *CoreVcnResourceCrud) Get() error {
	request := oci_core.GetVcnRequest{}

	tmp := s.D.Id()
	request.VcnId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVcn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vcn
	return nil
}

func (s *CoreVcnResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	if _, ok := s.D.GetOkExists("is_ipv6enabled"); ok && s.D.HasChange("is_ipv6enabled") {
		enableIPv6Request := oci_core.AddIpv6VcnCidrRequest{}
		tmp := s.D.Id()
		enableIPv6Request.VcnId = &tmp
		enableIPv6Request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

		_, err := s.Client.AddIpv6VcnCidr(context.Background(), enableIPv6Request)
		if err != nil {
			return err
		}
	}

	if byoipv6CidrDetails, ok := s.D.GetOkExists("byoipv6Cidr_details"); ok && s.D.HasChange("byoipv6Cidr_details") {
		err := s.addByoIpv6CidrBlocks(byoipv6CidrDetails)
		if err != nil {
			return err
		}
	}

	if _, ok := s.D.GetOkExists("ipv6private_cidr_blocks"); ok && s.D.HasChange("ipv6private_cidr_blocks") {
		oldRaw, newRaw := s.D.GetChange("ipv6private_cidr_blocks")
		if newRaw != "" && oldRaw != "" {
			err := s.updateIpv6CidrBlocks(oldRaw, newRaw)
			if err != nil {
				return err
			}
		}
	}

	if _, ok := s.D.GetOkExists("byoipv6cidr_blocks"); ok && s.D.HasChange("byoipv6cidr_blocks") {
		oldRaw, newRaw := s.D.GetChange("byoipv6cidr_blocks")
		if newRaw != "" && oldRaw != "" {
			err := s.updateIpv6CidrBlocks(oldRaw, newRaw)
			if err != nil {
				return err
			}
		}
	}

	if enableOracleGuaAllocation, ok := s.D.GetOkExists("is_oracle_gua_allocation_enabled"); ok && s.D.HasChange("is_oracle_gua_allocation_enabled") {
		enableIPv6Request := oci_core.AddIpv6VcnCidrRequest{}
		addVcnIpv6CidrDetails := oci_core.AddVcnIpv6CidrDetails{}
		tmp := s.D.Id()
		enableIPv6Request.VcnId = &tmp
		enableIPv6Request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
		isOracleGuaAllocationEnabled := enableOracleGuaAllocation.(bool)
		addVcnIpv6CidrDetails.IsOracleGuaAllocationEnabled = &isOracleGuaAllocationEnabled
		enableIPv6Request.AddVcnIpv6CidrDetails = addVcnIpv6CidrDetails
		_, err := s.Client.AddIpv6VcnCidr(context.Background(), enableIPv6Request)
		if err != nil {
			return err
		}
	}

	request := oci_core.UpdateVcnRequest{}

	if _, ok := s.D.GetOkExists("cidr_blocks"); ok && s.D.HasChange("cidr_blocks") {
		oldRaw, newRaw := s.D.GetChange("cidr_blocks")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCidrBlocks(oldRaw, newRaw)
			if err != nil {
				return err
			}
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		convertedAttributes := tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		request.SecurityAttributes = convertedAttributes
	}

	tmp := s.D.Id()
	request.VcnId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateVcn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vcn
	return nil
}

func (s *CoreVcnResourceCrud) addByoIpv6CidrBlocks(byoipv6CidrDetails interface{}) error {
	request := oci_core.AddIpv6VcnCidrRequest{}
	addVcnIpv6CidrDetails := oci_core.AddVcnIpv6CidrDetails{}
	fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "byoipv6cidr_details", byoipv6CidrDetails)
	converted, err := s.mapToByoipv6CidrDetails(fieldKeyFormat)
	if err != nil {
		return err
	}

	addVcnIpv6CidrDetails.Byoipv6CidrDetail = &converted
	idTmp := s.D.Id()
	request.VcnId = &idTmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
	request.AddVcnIpv6CidrDetails = addVcnIpv6CidrDetails
	_, err = s.Client.AddIpv6VcnCidr(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (s *CoreVcnResourceCrud) updateIpv6CidrBlocks(oldRaw interface{}, newRaw interface{}) error {
	interfaces := oldRaw.([]interface{})
	oldBlocks := make([]string, len(interfaces))
	for i := range interfaces {
		if interfaces[i] != nil {
			oldBlocks[i] = interfaces[i].(string)
		}
	}
	interfaces = newRaw.([]interface{})
	newBlocks := make([]string, len(interfaces))
	for i := range interfaces {
		if interfaces[i] != nil {
			newBlocks[i] = interfaces[i].(string)
		}
	}
	canEdit, operation, oldCidr, newCidr := oneEditAway(oldBlocks, newBlocks)
	if !canEdit {
		return fmt.Errorf("only one add/remove is allowed at once, new ipv6_cidr_block must be added at the end of list")
	}
	// add modify error

	if operation == "add" {
		addIpv6VcnCidrRequest := oci_core.AddIpv6VcnCidrRequest{}
		addVcnIpv6CidrDetails := oci_core.AddVcnIpv6CidrDetails{}
		idTmp := s.D.Id()
		addIpv6VcnCidrRequest.VcnId = &idTmp
		addIpv6VcnCidrRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
		addVcnIpv6CidrDetails.Ipv6PrivateCidrBlock = &newCidr
		addIpv6VcnCidrRequest.AddVcnIpv6CidrDetails = addVcnIpv6CidrDetails
		_, err := s.Client.AddIpv6VcnCidr(context.Background(), addIpv6VcnCidrRequest)
		if err != nil {
			return err
		}
	}
	if operation == "remove" {
		removeIpv6VcnCidrRequest := oci_core.RemoveIpv6VcnCidrRequest{}
		removeVcnIpv6CidrDetails := oci_core.RemoveVcnIpv6CidrDetails{}
		idTmp := s.D.Id()
		removeIpv6VcnCidrRequest.VcnId = &idTmp
		removeIpv6VcnCidrRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
		removeVcnIpv6CidrDetails.Ipv6CidrBlock = &oldCidr
		removeIpv6VcnCidrRequest.RemoveVcnIpv6CidrDetails = removeVcnIpv6CidrDetails
		_, err := s.Client.RemoveIpv6VcnCidr(context.Background(), removeIpv6VcnCidrRequest)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *CoreVcnResourceCrud) Delete() error {
	request := oci_core.DeleteVcnRequest{}

	tmp := s.D.Id()
	request.VcnId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVcn(context.Background(), request)
	return err
}

func (s *CoreVcnResourceCrud) SetData() error {
	s.D.Set("byoipv6cidr_blocks", s.Res.Byoipv6CidrBlocks)

	if s.Res.CidrBlock != nil {
		s.D.Set("cidr_block", *s.Res.CidrBlock)
	}

	s.D.Set("cidr_blocks", s.Res.CidrBlocks)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefaultDhcpOptionsId != nil {
		s.D.Set("default_dhcp_options_id", *s.Res.DefaultDhcpOptionsId)
	}

	if s.Res.DefaultRouteTableId != nil {
		s.D.Set("default_route_table_id", *s.Res.DefaultRouteTableId)
	}

	if s.Res.DefaultSecurityListId != nil {
		s.D.Set("default_security_list_id", *s.Res.DefaultSecurityListId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DnsLabel != nil {
		s.D.Set("dns_label", *s.Res.DnsLabel)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("ipv6cidr_blocks", s.Res.Ipv6CidrBlocks)

	s.D.Set("ipv6private_cidr_blocks", s.Res.Ipv6PrivateCidrBlocks)

	s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(s.Res.SecurityAttributes))

	if (s.Res.Ipv6CidrBlocks != nil && len(s.Res.Ipv6CidrBlocks) > 0) || (s.Res.Ipv6PrivateCidrBlocks != nil && len(s.Res.Ipv6PrivateCidrBlocks) > 0) {
		s.D.Set("is_ipv6enabled", true)
	} else {
		s.D.Set("is_ipv6enabled", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnDomainName != nil {
		s.D.Set("vcn_domain_name", *s.Res.VcnDomainName)
	}

	return nil
}

func (s *CoreVcnResourceCrud) mapToByoipv6CidrDetails(fieldKeyFormat string) (oci_core.Byoipv6CidrDetails, error) {
	result := oci_core.Byoipv6CidrDetails{}

	if byoipv6RangeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "byoipv6range_id")); ok {
		tmp := byoipv6RangeId.(string)
		result.Byoipv6RangeId = &tmp
	}

	if ipv6CidrBlock, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipv6cidr_block")); ok {
		tmp := ipv6CidrBlock.(string)
		result.Ipv6CidrBlock = &tmp
	}

	return result, nil
}

func Byoipv6CidrDetailsToMap(obj oci_core.Byoipv6CidrDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Byoipv6RangeId != nil {
		result["byoipv6range_id"] = string(*obj.Byoipv6RangeId)
	}

	if obj.Ipv6CidrBlock != nil {
		result["ipv6cidr_block"] = string(*obj.Ipv6CidrBlock)
	}

	return result
}

func (s *CoreVcnResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeVcnCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VcnId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeVcnCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *CoreVcnResourceCrud) updateCidrBlocks(oldRaw interface{}, newRaw interface{}) error {
	interfaces := oldRaw.([]interface{})
	oldBlocks := make([]string, len(interfaces))
	for i := range interfaces {
		if interfaces[i] != nil {
			oldBlocks[i] = interfaces[i].(string)
		}
	}
	interfaces = newRaw.([]interface{})
	newBlocks := make([]string, len(interfaces))
	for i := range interfaces {
		if interfaces[i] != nil {
			newBlocks[i] = interfaces[i].(string)
		}
	}
	canEdit, operation, oldCidr, newCidr := oneEditAway(oldBlocks, newBlocks)
	if !canEdit {
		return fmt.Errorf("only one add/remove or modification is allowed at once, new cidr_block must be added at the end of list")
	}
	if operation == "add" {
		addVcnCidrRequest := oci_core.AddVcnCidrRequest{}
		idTmp := s.D.Id()
		addVcnCidrRequest.VcnId = &idTmp
		addVcnCidrRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
		addVcnCidrRequest.CidrBlock = &newCidr
		_, err := s.Client.AddVcnCidr(context.Background(), addVcnCidrRequest)
		if err != nil {
			return err
		}
	}
	if operation == "remove" {
		removeVcnCidrRequest := oci_core.RemoveVcnCidrRequest{}
		idTmp := s.D.Id()
		removeVcnCidrRequest.VcnId = &idTmp
		removeVcnCidrRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
		removeVcnCidrRequest.CidrBlock = &oldCidr
		_, err := s.Client.RemoveVcnCidr(context.Background(), removeVcnCidrRequest)
		if err != nil {
			return err
		}
	}
	if operation == "modify" {
		modifyVcnCidrRequest := oci_core.ModifyVcnCidrRequest{}
		idTmp := s.D.Id()
		modifyVcnCidrRequest.VcnId = &idTmp
		modifyVcnCidrRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
		modifyVcnCidrRequest.OriginalCidrBlock = &oldCidr
		modifyVcnCidrRequest.NewCidrBlock = &newCidr
		_, err := s.Client.ModifyVcnCidr(context.Background(), modifyVcnCidrRequest)
		if err != nil {
			return err
		}
	}
	return nil
}

func oneEditAway(oldBlocks []string, newBlocks []string) (bool, string, string, string) {
	if Abs(len(newBlocks)-len(oldBlocks)) > 1 {
		return false, "", "", ""
	}
	if len(newBlocks) == len(oldBlocks) {
		for i := 0; i < len(oldBlocks); i++ {
			if oldBlocks[i] != newBlocks[i] {
				for j := i + 1; j < len(oldBlocks); j++ {
					if oldBlocks[j] != newBlocks[j] {
						return false, "", "", ""
					}
				}
				return true, "modify", oldBlocks[i], newBlocks[i]
			}
		}
	}
	if len(newBlocks) > len(oldBlocks) {
		for i := 0; i < len(oldBlocks); i++ {
			if oldBlocks[i] != newBlocks[i] {
				return false, "", "", ""
			}
		}
		return true, "add", "", newBlocks[len(newBlocks)-1]
	}
	for i := 0; i < len(newBlocks); i++ {
		if oldBlocks[i] != newBlocks[i] {
			for j := i + 1; j < len(newBlocks); j++ {
				if oldBlocks[j] != newBlocks[j-1] {
					return false, "", "", ""
				}
			}
			return true, "remove", oldBlocks[i], ""
		}
	}
	return true, "remove", oldBlocks[len(oldBlocks)-1], ""
}
