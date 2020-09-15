// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_dns "github.com/oracle/oci-go-sdk/v25/dns"
)

func init() {
	RegisterResource("oci_dns_tsig_key", DnsTsigKeyResource())
}

func DnsTsigKeyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDnsTsigKey,
		Read:     readDnsTsigKey,
		Update:   updateDnsTsigKey,
		Delete:   deleteDnsTsigKey,
		Schema: map[string]*schema.Schema{
			// Required
			"algorithm": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"secret": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
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

			// Computed
			"self": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDnsTsigKey(d *schema.ResourceData, m interface{}) error {
	sync := &DnsTsigKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return CreateResource(d, sync)
}

func readDnsTsigKey(d *schema.ResourceData, m interface{}) error {
	sync := &DnsTsigKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return ReadResource(sync)
}

func updateDnsTsigKey(d *schema.ResourceData, m interface{}) error {
	sync := &DnsTsigKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return UpdateResource(d, sync)
}

func deleteDnsTsigKey(d *schema.ResourceData, m interface{}) error {
	sync := &DnsTsigKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DnsTsigKeyResourceCrud struct {
	BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.TsigKey
	DisableNotFoundRetries bool
}

func (s *DnsTsigKeyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DnsTsigKeyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dns.TsigKeyLifecycleStateCreating),
	}
}

func (s *DnsTsigKeyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dns.TsigKeyLifecycleStateActive),
	}
}

func (s *DnsTsigKeyResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DnsTsigKeyResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DnsTsigKeyResourceCrud) Create() error {
	request := oci_dns.CreateTsigKeyRequest{}

	if algorithm, ok := s.D.GetOkExists("algorithm"); ok {
		tmp := algorithm.(string)
		request.Algorithm = &tmp
	}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if secret, ok := s.D.GetOkExists("secret"); ok {
		tmp := secret.(string)
		request.Secret = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.CreateTsigKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TsigKey
	return nil
}

func (s *DnsTsigKeyResourceCrud) Get() error {
	request := oci_dns.GetTsigKeyRequest{}

	tmp := s.D.Id()
	request.TsigKeyId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.GetTsigKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TsigKey
	return nil
}

func (s *DnsTsigKeyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dns.UpdateTsigKeyRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.TsigKeyId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateTsigKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TsigKey
	return nil
}

func (s *DnsTsigKeyResourceCrud) Delete() error {
	request := oci_dns.DeleteTsigKeyRequest{}

	tmp := s.D.Id()
	request.TsigKeyId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.DeleteTsigKey(context.Background(), request)
	return err
}

func (s *DnsTsigKeyResourceCrud) SetData() error {
	if s.Res.Algorithm != nil {
		s.D.Set("algorithm", *s.Res.Algorithm)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Secret != nil {
		s.D.Set("secret", *s.Res.Secret)
	}

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DnsTsigKeyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dns.ChangeTsigKeyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.TsigKeyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.ChangeTsigKeyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
