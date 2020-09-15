// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"sort"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v25/identity"
)

func init() {
	RegisterDatasource("oci_identity_availability_domains", IdentityAvailabilityDomainsDataSource())
}

func IdentityAvailabilityDomainsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityAvailabilityDomains,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"availability_domains": {
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
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readIdentityAvailabilityDomains(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAvailabilityDomainsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return ReadResource(sync)
}

type IdentityAvailabilityDomainsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListAvailabilityDomainsResponse
}

func (s *IdentityAvailabilityDomainsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityAvailabilityDomainsDataSourceCrud) Get() error {
	request := oci_identity.ListAvailabilityDomainsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListAvailabilityDomains(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityAvailabilityDomainsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	items := s.Res.Items

	// sort ADs by name
	sort.Slice(items, func(i, j int) bool {
		return *items[i].Name < *items[j].Name
	})

	for _, r := range items {
		availabilityDomain := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Id != nil {
			availabilityDomain["id"] = *r.Id
		}

		if r.Name != nil {
			availabilityDomain["name"] = *r.Name
		}

		resources = append(resources, availabilityDomain)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityAvailabilityDomainsDataSource().Schema["availability_domains"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("availability_domains", resources); err != nil {
		return err
	}

	return nil
}
