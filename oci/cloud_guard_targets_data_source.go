// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v25/cloudguard"
)

func init() {
	RegisterDatasource("oci_cloud_guard_targets", CloudGuardTargetsDataSource())
}

func CloudGuardTargetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudGuardTargets,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(CloudGuardTargetResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudGuardTargets(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardTargetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).cloudGuardClient()

	return ReadResource(sync)
}

type CloudGuardTargetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListTargetsResponse
}

func (s *CloudGuardTargetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardTargetsDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListTargetsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_cloud_guard.ListTargetsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_guard.ListTargetsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListTargets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTargets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudGuardTargetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}
	target := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TargetSummaryToMap(item))
	}
	target["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, CloudGuardTargetsDataSource().Schema["target_collection"].Elem.(*schema.Resource).Schema)
		target["items"] = items
	}

	resources = append(resources, target)
	if err := s.D.Set("target_collection", resources); err != nil {
		return err
	}

	return nil
}
