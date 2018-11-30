// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func InstanceConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readInstanceConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(InstanceConfigurationResource()),
			},
		},
	}
}

func readInstanceConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient

	return ReadResource(sync)
}

type InstanceConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.ListInstanceConfigurationsResponse
}

func (s *InstanceConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *InstanceConfigurationsDataSourceCrud) Get() error {
	request := oci_core.ListInstanceConfigurationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListInstanceConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInstanceConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *InstanceConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instanceConfiguration := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DisplayName != nil {
			instanceConfiguration["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			instanceConfiguration["id"] = *r.Id
		}

		if r.TimeCreated != nil {
			instanceConfiguration["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, instanceConfiguration)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, InstanceConfigurationsDataSource().Schema["instance_configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instance_configurations", resources); err != nil {
		return err
	}

	return nil
}
