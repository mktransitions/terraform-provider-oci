// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func IdentityProviderGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityProviderGroups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"identity_provider_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"identity_provider_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_provider_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readIdentityProviderGroups(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityProviderGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

type IdentityProviderGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListIdentityProviderGroupsResponse
}

func (s *IdentityProviderGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityProviderGroupsDataSourceCrud) Get() error {
	request := oci_identity.ListIdentityProviderGroupsRequest{}

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListIdentityProviderGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIdentityProviderGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityProviderGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		identityProviderGroup := map[string]interface{}{
			"identity_provider_id": *r.IdentityProviderId,
		}

		if r.DisplayName != nil {
			identityProviderGroup["display_name"] = *r.DisplayName
		}

		if r.ExternalIdentifier != nil {
			identityProviderGroup["external_identifier"] = *r.ExternalIdentifier
		}

		if r.Id != nil {
			identityProviderGroup["id"] = *r.Id
		}

		if r.TimeCreated != nil {
			identityProviderGroup["time_created"] = r.TimeCreated.String()
		}

		if r.TimeModified != nil {
			identityProviderGroup["time_modified"] = *r.TimeModified
		}

		resources = append(resources, identityProviderGroup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityProviderGroupsDataSource().Schema["identity_provider_groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("identity_provider_groups", resources); err != nil {
		return err
	}

	return nil
}
