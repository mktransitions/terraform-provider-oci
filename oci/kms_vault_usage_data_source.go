// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v25/keymanagement"
)

func init() {
	RegisterDatasource("oci_kms_vault_usage", KmsVaultUsageDataSource())
}

func KmsVaultUsageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularKmsVaultUsage,
		Schema: map[string]*schema.Schema{
			"vault_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"key_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"key_version_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularKmsVaultUsage(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultUsageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient()

	return ReadResource(sync)
}

type KmsVaultUsageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsVaultClient
	Res    *oci_kms.GetVaultUsageResponse
}

func (s *KmsVaultUsageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsVaultUsageDataSourceCrud) Get() error {
	request := oci_kms.GetVaultUsageRequest{}

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "kms")

	response, err := s.Client.GetVaultUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *KmsVaultUsageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	if s.Res.KeyCount != nil {
		s.D.Set("key_count", *s.Res.KeyCount)
	}

	if s.Res.KeyVersionCount != nil {
		s.D.Set("key_version_count", *s.Res.KeyVersionCount)
	}

	return nil
}
