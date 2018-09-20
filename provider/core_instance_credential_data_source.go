// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func InstanceCredentialDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularInstanceCredential,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularInstanceCredential(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceCredentialDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return ReadResource(sync)
}

type InstanceCredentialDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetWindowsInstanceInitialCredentialsResponse
}

func (s *InstanceCredentialDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *InstanceCredentialDataSourceCrud) Get() error {
	request := oci_core.GetWindowsInstanceInitialCredentialsRequest{}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetWindowsInstanceInitialCredentials(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *InstanceCredentialDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	if s.Res.Password != nil {
		s.D.Set("password", *s.Res.Password)
	}

	if s.Res.Username != nil {
		s.D.Set("username", *s.Res.Username)
	}

	return nil
}
