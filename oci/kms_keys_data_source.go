// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

func KeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readKeys,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(KeyResource()),
			},
		},
	}
}

func readKeys(d *schema.ResourceData, m interface{}) error {
	sync := &KeysDataSourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		return fmt.Errorf("management endpoint missing")
	}
	client, err := m.(*OracleClients).KmsManagementClient(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return ReadResource(sync)
}

type KeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsManagementClient
	Res    *oci_kms.ListKeysResponse
}

func (s *KeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KeysDataSourceCrud) Get() error {
	request := oci_kms.ListKeysRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "kms")

	response, err := s.Client.ListKeys(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListKeys(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *KeysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		key := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DisplayName != nil {
			key["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			key["id"] = *r.Id
		}

		key["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			key["time_created"] = r.TimeCreated.String()
		}

		if r.VaultId != nil {
			key["vault_id"] = *r.VaultId
		}

		resources = append(resources, key)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, KeysDataSource().Schema["keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("keys", resources); err != nil {
		return err
	}

	return nil
}
