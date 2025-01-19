// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"bytes"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreRouteTableResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreRouteTable,
		Read:     readCoreRouteTable,
		Update:   updateCoreRouteTable,
		Delete:   deleteCoreRouteTable,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"route_rules": {
				Type: schema.TypeSet,
				// Code-gen and specs say this should be required
				// Keep it optional to continue to allow empty route_rules and avoid a breaking change.
				Optional: true,
				MinItems: 0,
				Set:      routeRulesHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"network_entity_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"cidr_block": {
							Type:       schema.TypeString,
							Optional:   true,
							Computed:   true,
							Deprecated: tfresource.FieldDeprecatedForAnother("cidr_block", "destination"),
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destination": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destination_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &CoreRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &CoreRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &CoreRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &CoreRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreRouteTableResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.RouteTable
	DisableNotFoundRetries bool
}

func (s *CoreRouteTableResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreRouteTableResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.RouteTableLifecycleStateProvisioning),
	}
}

func (s *CoreRouteTableResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.RouteTableLifecycleStateAvailable),
	}
}

func (s *CoreRouteTableResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.RouteTableLifecycleStateTerminating),
	}
}

func (s *CoreRouteTableResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.RouteTableLifecycleStateTerminated),
	}
}

func (s *CoreRouteTableResourceCrud) Create() error {
	request := oci_core.CreateRouteTableRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RouteRules = []oci_core.RouteRule{}
	if routeRules, ok := s.D.GetOkExists("route_rules"); ok {
		set := routeRules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.RouteRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := routeRulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "route_rules", stateDataIndex)
			converted, err := s.mapToRouteRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.RouteRules = tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RouteTable
	return nil
}

func (s *CoreRouteTableResourceCrud) Get() error {
	request := oci_core.GetRouteTableRequest{}

	tmp := s.D.Id()
	request.RtId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RouteTable
	return nil
}

func (s *CoreRouteTableResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateRouteTableRequest{}

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

	request.RouteRules = []oci_core.RouteRule{}
	if routeRules, ok := s.D.GetOkExists("route_rules"); ok {
		set := routeRules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.RouteRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := routeRulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "route_rules", stateDataIndex)
			converted, err := s.mapToRouteRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.RouteRules = tmp
	}

	tmp := s.D.Id()
	request.RtId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RouteTable
	return nil
}

func (s *CoreRouteTableResourceCrud) Delete() error {
	request := oci_core.DeleteRouteTableRequest{}

	tmp := s.D.Id()
	request.RtId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteRouteTable(context.Background(), request)
	return err
}

func (s *CoreRouteTableResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	routeRules := []interface{}{}
	for _, item := range s.Res.RouteRules {
		routeRules = append(routeRules, RouteRuleToMap(item))
	}
	s.D.Set("route_rules", schema.NewSet(routeRulesHashCodeForSets, routeRules))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if _, ok := s.D.GetOk("manage_default_resource_id"); !ok {
		if s.Res.VcnId != nil {
			s.D.Set("vcn_id", *s.Res.VcnId)
		}
	}

	return nil
}

func (s *CoreRouteTableResourceCrud) mapToRouteRule(fieldKeyFormat string) (oci_core.RouteRule, error) {
	result := oci_core.RouteRule{}

	// @CODEGEN We need this change because the service will return both cidr_block and destination.
	// Without this change on Update operations terraform will send both paremeters since they are both in the statefile.
	// The service will complain if both parameters are not the same on the Update operation so we need to make sure only the relevant one in sent to the service.
	destinationType, destinationTypePresent := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_type"))
	if destinationTypePresent {
		tmp := oci_core.RouteRuleDestinationTypeEnum(destinationType.(string))
		result.DestinationType = tmp
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok && description != "" {
		tmp := description.(string)
		result.Description = &tmp
	}

	cidrBlockChanged := false
	cidrBlock, cidrBlockPresent := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cidr_block"))
	if cidrBlockPresent && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cidr_block")) {
		cidrBlockChanged = true
	}

	destinationChanged := false
	destination, destinationPresent := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination"))
	if destinationPresent && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "destination")) {
		tmp := destination.(string)
		result.Destination = &tmp
		destinationChanged = true
	}

	if destinationType == string(oci_core.RouteRuleDestinationTypeServiceCidrBlock) || (!destinationChanged && !cidrBlockChanged) {
		tmp := destination.(string)
		result.Destination = &tmp
	}

	if !destinationChanged && cidrBlockPresent && cidrBlock != "" && destinationType != string(oci_core.RouteRuleDestinationTypeServiceCidrBlock) {
		tmp := cidrBlock.(string)
		result.CidrBlock = &tmp
	}

	if networkEntityId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_entity_id")); ok {
		tmp := networkEntityId.(string)
		result.NetworkEntityId = &tmp
	}

	if routeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "route_type")); ok {
		result.RouteType = oci_core.RouteRuleRouteTypeEnum(routeType.(string))
	}

	return result, nil
}

func RouteRuleToMap(obj oci_core.RouteRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CidrBlock != nil {
		result["cidr_block"] = string(*obj.CidrBlock)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Destination != nil {
		result["destination"] = string(*obj.Destination)
	}

	result["destination_type"] = string(obj.DestinationType)
	result["route_type"] = string(obj.RouteType)

	if obj.NetworkEntityId != nil {
		result["network_entity_id"] = string(*obj.NetworkEntityId)
	}

	return result
}

func routeRulesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	/* @CODEGEN The user needs to provide either cidr_block or destination.
	 * We need to make them both the same in the hashing function otherwise there will be a diff on every apply.
	 * This is because the service will return both fields
	 */
	cidrBlock, cidrBlockPresent := m["cidr_block"]
	destination, destinationPresent := m["destination"]
	if cidrBlockPresent && cidrBlock != "" {
		buf.WriteString(fmt.Sprintf("%v-", cidrBlock))
	} else if destinationPresent && destination != "" {
		buf.WriteString(fmt.Sprintf("%v-", destination))
	}
	if description, ok := m["description"]; ok && description != "" {
		buf.WriteString(fmt.Sprintf("%v-", description))
	}
	if destinationPresent && destination != "" {
		buf.WriteString(fmt.Sprintf("%v-", destination))
	} else if cidrBlockPresent && cidrBlock != "" {
		buf.WriteString(fmt.Sprintf("%v-", cidrBlock))
	}
	if destinationType, ok := m["destination_type"]; ok && destinationType != "" {
		buf.WriteString(fmt.Sprintf("%v-", destinationType))
	} else {
		buf.WriteString(fmt.Sprintf("%v-", oci_core.RouteRuleDestinationTypeCidrBlock))
	}

	if networkEntityId, ok := m["network_entity_id"]; ok && networkEntityId != "" {
		buf.WriteString(fmt.Sprintf("%v-", networkEntityId))
	}

	if routeType, ok := m["route_type"]; ok && routeType != "" {
		buf.WriteString(fmt.Sprintf("%v-", routeType))
	}

	return utils.GetStringHashcode(buf.String())

}
func (s *CoreRouteTableResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeRouteTableCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.RtId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeRouteTableCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
