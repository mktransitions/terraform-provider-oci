// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v25/core"
)

func init() {
	RegisterResource("oci_core_public_ip", CorePublicIpResource())
}

func CorePublicIpResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCorePublicIp,
		Read:     readCorePublicIp,
		Update:   updateCorePublicIp,
		Delete:   deleteCorePublicIp,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"lifetime": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
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
			"private_ip_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
			"assigned_entity_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assigned_entity_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
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

func createCorePublicIp(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return CreateResource(d, sync)
}

func readCorePublicIp(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

func updateCorePublicIp(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return UpdateResource(d, sync)
}

func deleteCorePublicIp(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CorePublicIpResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.PublicIp
	DisableNotFoundRetries bool
}

func (s *CorePublicIpResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CorePublicIpResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateProvisioning),
		string(oci_core.PublicIpLifecycleStateAssigning),
	}
}

func (s *CorePublicIpResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateAvailable),
		string(oci_core.PublicIpLifecycleStateAssigned),
	}
}

func (s *CorePublicIpResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateUnassigning),
		string(oci_core.PublicIpLifecycleStateTerminating),
	}
}

func (s *CorePublicIpResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateUnassigned),
		string(oci_core.PublicIpLifecycleStateTerminated),
	}
}

func (s *CorePublicIpResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateProvisioning),
		string(oci_core.PublicIpLifecycleStateAssigning),
		string(oci_core.PublicIpLifecycleStateUnassigning),
	}
}

func (s *CorePublicIpResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateAvailable),
		string(oci_core.PublicIpLifecycleStateAssigned),
	}
}

func (s *CorePublicIpResourceCrud) Create() error {
	request := oci_core.CreatePublicIpRequest{}

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

	if lifetime, ok := s.D.GetOkExists("lifetime"); ok {
		request.Lifetime = oci_core.CreatePublicIpDetailsLifetimeEnum(lifetime.(string))
	}

	if privateIpId, ok := s.D.GetOkExists("private_ip_id"); ok {
		tmp := privateIpId.(string)
		request.PrivateIpId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreatePublicIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PublicIp
	return nil
}

func (s *CorePublicIpResourceCrud) Get() error {
	request := oci_core.GetPublicIpRequest{}

	tmp := s.D.Id()
	request.PublicIpId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetPublicIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PublicIp
	return nil
}

func (s *CorePublicIpResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdatePublicIpRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	// Wrapping in "HasChange" conditionals because the service will treat the PUT as a PATCH.
	if s.D.HasChange("display_name") {
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			request.DisplayName = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if s.D.HasChange("private_ip_id") {
		if privateIpId, ok := s.D.GetOkExists("private_ip_id"); ok {
			tmp := privateIpId.(string)
			request.PrivateIpId = &tmp
		}
	}

	tmp := s.D.Id()
	request.PublicIpId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdatePublicIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PublicIp
	return nil
}

func (s *CorePublicIpResourceCrud) Delete() error {
	request := oci_core.DeletePublicIpRequest{}

	tmp := s.D.Id()
	request.PublicIpId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeletePublicIp(context.Background(), request)
	return err
}

func (s *CorePublicIpResourceCrud) SetData() error {
	if s.Res.AssignedEntityId != nil {
		s.D.Set("assigned_entity_id", *s.Res.AssignedEntityId)
	}

	s.D.Set("assigned_entity_type", s.Res.AssignedEntityType)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
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

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	s.D.Set("lifetime", s.Res.Lifetime)

	if s.Res.PrivateIpId != nil {
		s.D.Set("private_ip_id", *s.Res.PrivateIpId)
	}

	s.D.Set("scope", s.Res.Scope)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CorePublicIpResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangePublicIpCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.PublicIpId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangePublicIpCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
