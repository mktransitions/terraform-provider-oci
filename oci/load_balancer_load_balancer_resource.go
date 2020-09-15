// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v25/loadbalancer"
)

func init() {
	RegisterResource("oci_load_balancer_load_balancer", LoadBalancerLoadBalancerResource())
}

func LoadBalancerLoadBalancerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createLoadBalancerLoadBalancer,
		Read:     readLoadBalancerLoadBalancer,
		Update:   updateLoadBalancerLoadBalancer,
		Delete:   deleteLoadBalancerLoadBalancer,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_ids": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"ip_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_private": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"network_security_group_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      literalTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"ip_address_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_public": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"ip_addresses": {
				Type:       schema.TypeList,
				Computed:   true,
				Deprecated: FieldDeprecatedForAnother("ip_addresses", "ip_address_details"),
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoadBalancerLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()

	return CreateResource(d, sync)
}

func readLoadBalancerLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()

	return ReadResource(sync)
}

func updateLoadBalancerLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()

	return UpdateResource(d, sync)
}

func deleteLoadBalancerLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type LoadBalancerLoadBalancerResourceCrud struct {
	BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.LoadBalancer
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *LoadBalancerLoadBalancerResourceCrud) ID() string {
	id, workSuccess := LoadBalancerResourceID(s.Res, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		return *s.WorkRequest.LoadBalancerId
	}
	return ""
}

func (s *LoadBalancerLoadBalancerResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateCreating),
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerLoadBalancerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateActive),
		string(oci_load_balancer.LoadBalancerLifecycleStateFailed),
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerLoadBalancerResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateDeleting),
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerLoadBalancerResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateDeleted),
	}
}

func (s *LoadBalancerLoadBalancerResourceCrud) Create() error {
	request := oci_load_balancer.CreateLoadBalancerRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if ipMode, ok := s.D.GetOkExists("ip_mode"); ok {
		request.IpMode = oci_load_balancer.CreateLoadBalancerDetailsIpModeEnum(ipMode.(string))
	}

	if isPrivate, ok := s.D.GetOkExists("is_private"); ok {
		tmp := isPrivate.(bool)
		request.IsPrivate = &tmp
	}

	if networkSecurityGroupIds, ok := s.D.GetOkExists("network_security_group_ids"); ok {
		set := networkSecurityGroupIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("network_security_group_ids") {
			request.NetworkSecurityGroupIds = tmp
		}
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.ShapeName = &tmp
	}

	if subnetIds, ok := s.D.GetOkExists("subnet_ids"); ok {
		interfaces := subnetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("subnet_ids") {
			request.SubnetIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerLoadBalancerResourceCrud) Get() error {
	id, stillWorking, err := LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	if id == "" && s.WorkRequest != nil {
		id = *s.WorkRequest.LoadBalancerId
		s.D.SetId(id)
	}

	request := oci_load_balancer.GetLoadBalancerRequest{}

	tmp := s.D.Id()
	request.LoadBalancerId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LoadBalancer
	return nil
}

func (s *LoadBalancerLoadBalancerResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	if shape, ok := s.D.GetOkExists("shape"); ok && s.D.HasChange("shape") {
		oldRaw, newRaw := s.D.GetChange("shape")
		if newRaw != "" && oldRaw != "" {
			err := s.updateShape(shape)
			if err != nil {
				return err
			}
		}
	}

	if s.D.HasChange("network_security_group_ids") {
		err := s.updateNetworkSecurityGroups()
		if err != nil {
			return fmt.Errorf("unable to update 'network_security_group_ids', error: %v", err)
		}
	}
	request := oci_load_balancer.UpdateLoadBalancerRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.LoadBalancerId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *LoadBalancerLoadBalancerResourceCrud) Delete() error {
	request := oci_load_balancer.DeleteLoadBalancerRequest{}

	tmp := s.D.Id()
	request.LoadBalancerId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerLoadBalancerResourceCrud) SetData() error {
	if s.Res == nil || s.Res.Id == nil {
		return nil
	}
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	ipAddresses := []string{}
	ipMode := "IPV4"
	for _, ad := range s.Res.IpAddresses {
		if ad.IpAddress != nil {
			ipAddresses = append(ipAddresses, *ad.IpAddress)
		}
		tmp := *ad.IpAddress
		if !isIPV4(tmp) {
			ipMode = "IPV6"
		}
	}
	s.D.Set("ip_mode", ipMode)
	s.D.Set("ip_addresses", ipAddresses)

	ipAddressDetails := []interface{}{}

	for _, item := range s.Res.IpAddresses {
		ipAddressDetails = append(ipAddressDetails, IpAddressToMap(item))
	}

	s.D.Set("ip_address_details", ipAddressDetails)

	if s.Res.IsPrivate != nil {
		s.D.Set("is_private", *s.Res.IsPrivate)
	}

	networkSecurityGroupIds := []interface{}{}
	for _, item := range s.Res.NetworkSecurityGroupIds {
		networkSecurityGroupIds = append(networkSecurityGroupIds, item)
	}
	s.D.Set("network_security_group_ids", schema.NewSet(literalTypeHashCodeForSets, networkSecurityGroupIds))

	if s.Res.ShapeName != nil {
		s.D.Set("shape", *s.Res.ShapeName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subnet_ids", s.Res.SubnetIds)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", systemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func IpAddressToMap(obj oci_load_balancer.IpAddress) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.IsPublic != nil {
		result["is_public"] = bool(*obj.IsPublic)
	}

	return result
}

func isIPV4(ipAddress string) bool {
	return strings.Contains(ipAddress, ".")
}

func (s *LoadBalancerLoadBalancerResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_load_balancer.ChangeLoadBalancerCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.LoadBalancerId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	_, err := s.Client.ChangeLoadBalancerCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerLoadBalancerResourceCrud) updateNetworkSecurityGroups() error {
	updateNsgIdsRequest := oci_load_balancer.UpdateNetworkSecurityGroupsRequest{}

	//@Codegen: Unless explicitly specified by the user, network_security_group_ids will not be supplied as the feature may or may not be supported
	if networkSecurityGroupIds, ok := s.D.GetOkExists("network_security_group_ids"); ok {
		set := networkSecurityGroupIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		updateNsgIdsRequest.NetworkSecurityGroupIds = tmp
	}

	tmp := s.D.Id()
	updateNsgIdsRequest.LoadBalancerId = &tmp

	updateNsgIdsRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateNetworkSecurityGroups(context.Background(), updateNsgIdsRequest)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerLoadBalancerResourceCrud) updateShape(shape interface{}) error {
	changeShapeRequest := oci_load_balancer.UpdateLoadBalancerShapeRequest{}

	shapeTmp := shape.(string)
	changeShapeRequest.ShapeName = &shapeTmp

	idTmp := s.D.Id()
	changeShapeRequest.LoadBalancerId = &idTmp

	changeShapeRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateLoadBalancerShape(context.Background(), changeShapeRequest)
	if err != nil {
		return err
	}
	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}
