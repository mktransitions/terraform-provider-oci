// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/email"
)

func SuppressionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createSuppression,
		Read:     readSuppression,
		Delete:   deleteSuppression,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"email_address": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},

			// Optional

			// Computed
			"reason": {
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

func createSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &SuppressionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient

	return CreateResource(d, sync)
}

func readSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &SuppressionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient

	return ReadResource(sync)
}

func deleteSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &SuppressionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type SuppressionResourceCrud struct {
	BaseCrud
	Client                 *oci_email.EmailClient
	Res                    *oci_email.Suppression
	DisableNotFoundRetries bool
}

func (s *SuppressionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *SuppressionResourceCrud) Create() error {
	request := oci_email.CreateSuppressionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if emailAddress, ok := s.D.GetOkExists("email_address"); ok {
		tmp := emailAddress.(string)
		request.EmailAddress = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.CreateSuppression(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Suppression
	return nil
}

func (s *SuppressionResourceCrud) Get() error {
	request := oci_email.GetSuppressionRequest{}

	tmp := s.D.Id()
	request.SuppressionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.GetSuppression(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Suppression
	return nil
}

func (s *SuppressionResourceCrud) Delete() error {
	request := oci_email.DeleteSuppressionRequest{}

	tmp := s.D.Id()
	request.SuppressionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "email")

	_, err := s.Client.DeleteSuppression(context.Background(), request)
	return err
}

func (s *SuppressionResourceCrud) SetData() error {
	if s.Res.EmailAddress != nil {
		s.D.Set("email_address", *s.Res.EmailAddress)
	}

	s.D.Set("reason", s.Res.Reason)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
