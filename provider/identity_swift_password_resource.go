// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func SwiftPasswordResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createSwiftPassword,
		Read:     readSwiftPassword,
		Update:   updateSwiftPassword,
		Delete:   deleteSwiftPassword,
		Schema: map[string]*schema.Schema{
			// Required
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"expires_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
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

func createSwiftPassword(d *schema.ResourceData, m interface{}) error {
	sync := &SwiftPasswordResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return CreateResource(d, sync)
}

func readSwiftPassword(d *schema.ResourceData, m interface{}) error {
	sync := &SwiftPasswordResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

func updateSwiftPassword(d *schema.ResourceData, m interface{}) error {
	sync := &SwiftPasswordResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return UpdateResource(d, sync)
}

func deleteSwiftPassword(d *schema.ResourceData, m interface{}) error {
	sync := &SwiftPasswordResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type SwiftPasswordResourceCrud struct {
	BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.SwiftPassword
	DisableNotFoundRetries bool
}

func (s *SwiftPasswordResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *SwiftPasswordResourceCrud) State() oci_identity.SwiftPasswordLifecycleStateEnum {
	return s.Res.LifecycleState
}

func (s *SwiftPasswordResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.SwiftPasswordLifecycleStateCreating),
	}
}

func (s *SwiftPasswordResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.SwiftPasswordLifecycleStateActive),
	}
}

func (s *SwiftPasswordResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.SwiftPasswordLifecycleStateDeleting),
	}
}

func (s *SwiftPasswordResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.SwiftPasswordLifecycleStateDeleted),
	}
}

func (s *SwiftPasswordResourceCrud) Create() error {
	request := oci_identity.CreateSwiftPasswordRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateSwiftPassword(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SwiftPassword
	return nil
}

func (s *SwiftPasswordResourceCrud) Get() error {
	request := oci_identity.ListSwiftPasswordsRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ListSwiftPasswords(context.Background(), request)
	if err != nil {
		return err
	}

	id := s.D.Id()
	for _, item := range response.Items {
		if *item.Id == id {
			s.Res = &item
			return nil
		}
	}

	return nil
}

func (s *SwiftPasswordResourceCrud) Update() error {
	request := oci_identity.UpdateSwiftPasswordRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.SwiftPasswordId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateSwiftPassword(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SwiftPassword
	return nil
}

func (s *SwiftPasswordResourceCrud) Delete() error {
	request := oci_identity.DeleteSwiftPasswordRequest{}

	tmp := s.D.Id()
	request.SwiftPasswordId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteSwiftPassword(context.Background(), request)
	return err
}

func (s *SwiftPasswordResourceCrud) SetData() error {
	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.ExpiresOn != nil {
		s.D.Set("expires_on", *s.Res.ExpiresOn)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.Password != nil {
		s.D.Set("password", *s.Res.Password)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

	return nil
}
