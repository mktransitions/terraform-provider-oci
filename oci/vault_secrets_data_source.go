// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_vault "github.com/oracle/oci-go-sdk/v25/vault"
)

func init() {
	RegisterDatasource("oci_vault_secrets", VaultSecretsDataSource())
}

func VaultSecretsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVaultSecrets,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vault_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secrets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"secret_name": {
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
						"time_of_current_version_expiry": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_deletion": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vault_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readVaultSecrets(d *schema.ResourceData, m interface{}) error {
	sync := &VaultSecretsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).vaultsClient()

	return ReadResource(sync)
}

type VaultSecretsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_vault.VaultsClient
	Res    *oci_vault.ListSecretsResponse
}

func (s *VaultSecretsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VaultSecretsDataSourceCrud) Get() error {
	request := oci_vault.ListSecretsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_vault.SecretSummaryLifecycleStateEnum(state.(string))
	}

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "vault")

	response, err := s.Client.ListSecrets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSecrets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *VaultSecretsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		secret := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			secret["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			secret["description"] = *r.Description
		}

		secret["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			secret["id"] = *r.Id
		}

		if r.KeyId != nil {
			secret["key_id"] = *r.KeyId
		}

		if r.LifecycleDetails != nil {
			secret["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.SecretName != nil {
			secret["secret_name"] = *r.SecretName
		}

		secret["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			secret["time_created"] = r.TimeCreated.String()
		}

		if r.TimeOfCurrentVersionExpiry != nil {
			secret["time_of_current_version_expiry"] = r.TimeOfCurrentVersionExpiry.String()
		}

		if r.TimeOfDeletion != nil {
			secret["time_of_deletion"] = r.TimeOfDeletion.String()
		}

		if r.VaultId != nil {
			secret["vault_id"] = *r.VaultId
		}

		resources = append(resources, secret)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, VaultSecretsDataSource().Schema["secrets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("secrets", resources); err != nil {
		return err
	}

	return nil
}
