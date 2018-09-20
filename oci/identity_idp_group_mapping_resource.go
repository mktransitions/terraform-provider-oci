// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func IdpGroupMappingResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createIdpGroupMapping,
		Read:     readIdpGroupMapping,
		Update:   updateIdpGroupMapping,
		Delete:   deleteIdpGroupMapping,
		Schema: map[string]*schema.Schema{
			// Required
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"identity_provider_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"idp_group_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_state": {
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

func createIdpGroupMapping(d *schema.ResourceData, m interface{}) error {
	sync := &IdpGroupMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return CreateResource(d, sync)
}

func readIdpGroupMapping(d *schema.ResourceData, m interface{}) error {
	sync := &IdpGroupMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

func updateIdpGroupMapping(d *schema.ResourceData, m interface{}) error {
	sync := &IdpGroupMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return UpdateResource(d, sync)
}

func deleteIdpGroupMapping(d *schema.ResourceData, m interface{}) error {
	sync := &IdpGroupMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type IdpGroupMappingResourceCrud struct {
	BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.IdpGroupMapping
	DisableNotFoundRetries bool
}

func (s *IdpGroupMappingResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdpGroupMappingResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.IdpGroupMappingLifecycleStateCreating),
	}
}

func (s *IdpGroupMappingResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.IdpGroupMappingLifecycleStateActive),
	}
}

func (s *IdpGroupMappingResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.IdpGroupMappingLifecycleStateDeleting),
	}
}

func (s *IdpGroupMappingResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.IdpGroupMappingLifecycleStateDeleted),
	}
}

func (s *IdpGroupMappingResourceCrud) Create() error {
	request := oci_identity.CreateIdpGroupMappingRequest{}

	if groupId, ok := s.D.GetOkExists("group_id"); ok {
		tmp := groupId.(string)
		request.GroupId = &tmp
	}

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	if idpGroupName, ok := s.D.GetOkExists("idp_group_name"); ok {
		tmp := idpGroupName.(string)
		request.IdpGroupName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateIdpGroupMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdpGroupMapping
	return nil
}

func (s *IdpGroupMappingResourceCrud) Get() error {
	request := oci_identity.GetIdpGroupMappingRequest{}

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	tmp := s.D.Id()
	request.MappingId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetIdpGroupMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdpGroupMapping
	return nil
}

func (s *IdpGroupMappingResourceCrud) Update() error {
	request := oci_identity.UpdateIdpGroupMappingRequest{}

	if groupId, ok := s.D.GetOkExists("group_id"); ok {
		tmp := groupId.(string)
		request.GroupId = &tmp
	}

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	if idpGroupName, ok := s.D.GetOkExists("idp_group_name"); ok {
		tmp := idpGroupName.(string)
		request.IdpGroupName = &tmp
	}

	tmp := s.D.Id()
	request.MappingId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateIdpGroupMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdpGroupMapping
	return nil
}

func (s *IdpGroupMappingResourceCrud) Delete() error {
	request := oci_identity.DeleteIdpGroupMappingRequest{}

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	tmp := s.D.Id()
	request.MappingId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteIdpGroupMapping(context.Background(), request)
	return err
}

func (s *IdpGroupMappingResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.GroupId != nil {
		s.D.Set("group_id", *s.Res.GroupId)
	}

	if s.Res.IdpId != nil {
		s.D.Set("identity_provider_id", *s.Res.IdpId)
	}

	if s.Res.IdpGroupName != nil {
		s.D.Set("idp_group_name", *s.Res.IdpGroupName)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
