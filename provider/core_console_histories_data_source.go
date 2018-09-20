// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func ConsoleHistoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readConsoleHistories,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"console_histories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(ConsoleHistoryResource()),
			},
		},
	}
}

func readConsoleHistories(d *schema.ResourceData, m interface{}) error {
	sync := &ConsoleHistoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return ReadResource(sync)
}

type ConsoleHistoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListConsoleHistoriesResponse
}

func (s *ConsoleHistoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ConsoleHistoriesDataSourceCrud) Get() error {
	request := oci_core.ListConsoleHistoriesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.ConsoleHistoryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListConsoleHistories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListConsoleHistories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ConsoleHistoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		consoleHistory := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			consoleHistory["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DefinedTags != nil {
			consoleHistory["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			consoleHistory["display_name"] = *r.DisplayName
		}

		consoleHistory["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			consoleHistory["id"] = *r.Id
		}

		if r.InstanceId != nil {
			consoleHistory["instance_id"] = *r.InstanceId
		}

		consoleHistory["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			consoleHistory["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, consoleHistory)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ConsoleHistoriesDataSource().Schema["console_histories"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("console_histories", resources); err != nil {
		return err
	}

	return nil
}
