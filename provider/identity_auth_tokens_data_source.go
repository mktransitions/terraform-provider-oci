// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func AuthTokensDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAuthTokens,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tokens": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(AuthTokenResource()),
			},
		},
	}
}

func readAuthTokens(d *schema.ResourceData, m interface{}) error {
	sync := &AuthTokensDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

type AuthTokensDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListAuthTokensResponse
}

func (s *AuthTokensDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AuthTokensDataSourceCrud) Get() error {
	request := oci_identity.ListAuthTokensRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListAuthTokens(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AuthTokensDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		authToken := map[string]interface{}{
			"user_id": *r.UserId,
		}

		if r.Description != nil {
			authToken["description"] = *r.Description
		}

		if r.Id != nil {
			authToken["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			authToken["inactive_state"] = strconv.FormatInt(*r.InactiveStatus, 10)
		}

		authToken["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			authToken["time_created"] = r.TimeCreated.String()
		}

		if r.TimeExpires != nil {
			authToken["time_expires"] = r.TimeExpires.String()
		}

		if r.Token != nil {
			authToken["token"] = *r.Token
		}

		resources = append(resources, authToken)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, AuthTokensDataSource().Schema["tokens"].Elem.(*schema.Resource).Schema)
	}

	// @CODEGEN 06/2018: auth_tokens => tokens
	if err := s.D.Set("tokens", resources); err != nil {
		return err
	}

	return nil
}
