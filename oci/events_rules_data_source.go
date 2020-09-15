// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_events "github.com/oracle/oci-go-sdk/v25/events"
)

func init() {
	RegisterDatasource("oci_events_rules", EventsRulesDataSource())
}

func EventsRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readEventsRules,
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
			"rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(EventsRuleResource()),
			},
		},
	}
}

func readEventsRules(d *schema.ResourceData, m interface{}) error {
	sync := &EventsRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).eventsClient()

	return ReadResource(sync)
}

type EventsRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_events.EventsClient
	Res    *oci_events.ListRulesResponse
}

func (s *EventsRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EventsRulesDataSourceCrud) Get() error {
	request := oci_events.ListRulesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_events.RuleLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "events")

	response, err := s.Client.ListRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *EventsRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		rule := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Condition != nil {
			rule["condition"] = *r.Condition
		}

		if r.DefinedTags != nil {
			rule["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			rule["description"] = *r.Description
		}

		if r.DisplayName != nil {
			rule["display_name"] = *r.DisplayName
		}

		rule["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			rule["id"] = *r.Id
		}

		if r.IsEnabled != nil {
			rule["is_enabled"] = *r.IsEnabled
		}

		rule["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			rule["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, rule)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, EventsRulesDataSource().Schema["rules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("rules", resources); err != nil {
		return err
	}

	return nil
}
