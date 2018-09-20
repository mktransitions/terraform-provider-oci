// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func DynamicGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDynamicGroups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dynamic_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DynamicGroupResource()),
			},
		},
	}
}

func readDynamicGroups(d *schema.ResourceData, m interface{}) error {
	sync := &DynamicGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

type DynamicGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListDynamicGroupsResponse
}

func (s *DynamicGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DynamicGroupsDataSourceCrud) Get() error {
	request := oci_identity.ListDynamicGroupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListDynamicGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDynamicGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DynamicGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dynamicGroup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Description != nil {
			dynamicGroup["description"] = *r.Description
		}

		if r.Id != nil {
			dynamicGroup["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			dynamicGroup["inactive_state"] = strconv.FormatInt(*r.InactiveStatus, 10)
		}

		if r.MatchingRule != nil {
			dynamicGroup["matching_rule"] = *r.MatchingRule
		}

		if r.Name != nil {
			dynamicGroup["name"] = *r.Name
		}

		dynamicGroup["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			dynamicGroup["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, dynamicGroup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DynamicGroupsDataSource().Schema["dynamic_groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("dynamic_groups", resources); err != nil {
		return err
	}

	return nil
}
