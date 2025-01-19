// Copyright (c) 2024, Néfix Estrada. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CorePublicIpAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCorePublicIpAttachment,
		Read:     readCorePublicIpAttachment,
		Update:   updateCorePublicIpAttachment,
		Delete:   deleteCorePublicIpAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"public_ip_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"private_ip_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCorePublicIpAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCorePublicIpAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCorePublicIpAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCorePublicIpAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.DeleteResource(d, sync)
}

type CorePublicIpAttachmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.PublicIp
	DisableNotFoundRetries bool
}

func (s *CorePublicIpAttachmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CorePublicIpAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateAssigning),
		string(oci_core.PublicIpLifecycleStateUnassigning),
	}
}

func (s *CorePublicIpAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateAvailable),
		string(oci_core.PublicIpLifecycleStateAssigned),
	}
}

func (s *CorePublicIpAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateUnassigning),
	}
}

func (s *CorePublicIpAttachmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateUnassigned),
		string(oci_core.PublicIpLifecycleStateAvailable),
	}
}

func (s *CorePublicIpAttachmentResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateAssigning),
		string(oci_core.PublicIpLifecycleStateUnassigning),
	}
}

func (s *CorePublicIpAttachmentResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateAvailable),
		string(oci_core.PublicIpLifecycleStateAssigned),
	}
}

func (s *CorePublicIpAttachmentResourceCrud) Create() error {
	request := oci_core.UpdatePublicIpRequest{}

	if publicIpId, ok := s.D.GetOkExists("public_ip_id"); ok {
		tmp := publicIpId.(string)
		request.PublicIpId = &tmp
	}

	if privateIpId, ok := s.D.GetOkExists("private_ip_id"); ok {
		tmp := privateIpId.(string)
		request.PrivateIpId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdatePublicIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PublicIp
	return nil
}

func (s *CorePublicIpAttachmentResourceCrud) Get() error {
	request := oci_core.GetPublicIpRequest{}

	tmp := s.D.Id()
	request.PublicIpId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetPublicIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PublicIp
	return nil
}

func (s *CorePublicIpAttachmentResourceCrud) Update() error {
	request := oci_core.UpdatePublicIpRequest{}

	if s.D.HasChange("private_ip_id") {
		if privateIpId, ok := s.D.GetOkExists("private_ip_id"); ok {
			tmp := privateIpId.(string)
			request.PrivateIpId = &tmp
		}
	}

	tmp := s.D.Id()
	request.PublicIpId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdatePublicIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PublicIp
	return nil
}

func (s *CorePublicIpAttachmentResourceCrud) Delete() error {
	request := oci_core.UpdatePublicIpRequest{}

	tmp := s.D.Id()
	request.PublicIpId = &tmp

	tmp2 := ""
	request.PrivateIpId = &tmp2

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.UpdatePublicIp(context.Background(), request)
	return err
}

func (s *CorePublicIpAttachmentResourceCrud) SetData() error {
	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.PrivateIpId != nil {
		s.D.Set("private_ip_id", *s.Res.PrivateIpId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
