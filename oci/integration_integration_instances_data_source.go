// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_integration "github.com/oracle/oci-go-sdk/v25/integration"
)

func init() {
	RegisterDatasource("oci_integration_integration_instances", IntegrationIntegrationInstancesDataSource())
}

func IntegrationIntegrationInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIntegrationIntegrationInstances,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"integration_instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(IntegrationIntegrationInstanceResource()),
			},
		},
	}
}

func readIntegrationIntegrationInstances(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationIntegrationInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).integrationInstanceClient()

	return ReadResource(sync)
}

type IntegrationIntegrationInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_integration.IntegrationInstanceClient
	Res    *oci_integration.ListIntegrationInstancesResponse
}

func (s *IntegrationIntegrationInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IntegrationIntegrationInstancesDataSourceCrud) Get() error {
	request := oci_integration.ListIntegrationInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_integration.ListIntegrationInstancesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "integration")

	response, err := s.Client.ListIntegrationInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIntegrationInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IntegrationIntegrationInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		integrationInstance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		integrationInstance["consumption_model"] = r.ConsumptionModel

		if r.DisplayName != nil {
			integrationInstance["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			integrationInstance["id"] = *r.Id
		}

		if r.InstanceUrl != nil {
			integrationInstance["instance_url"] = *r.InstanceUrl
		}

		integrationInstance["integration_instance_type"] = r.IntegrationInstanceType

		if r.IsByol != nil {
			integrationInstance["is_byol"] = *r.IsByol
		}

		if r.IsFileServerEnabled != nil {
			integrationInstance["is_file_server_enabled"] = *r.IsFileServerEnabled
		}

		if r.MessagePacks != nil {
			integrationInstance["message_packs"] = *r.MessagePacks
		}

		integrationInstance["state"] = r.LifecycleState

		if r.StateMessage != nil {
			integrationInstance["state_message"] = *r.StateMessage
		}

		if r.TimeCreated != nil {
			integrationInstance["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			integrationInstance["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, integrationInstance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IntegrationIntegrationInstancesDataSource().Schema["integration_instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("integration_instances", resources); err != nil {
		return err
	}

	return nil
}
