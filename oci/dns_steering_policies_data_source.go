// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v25/common"
	oci_dns "github.com/oracle/oci-go-sdk/v25/dns"
)

func init() {
	RegisterDatasource("oci_dns_steering_policies", DnsSteeringPoliciesDataSource())
}

func DnsSteeringPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDnsSteeringPolicies,
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
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_check_monitor_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"steering_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DnsSteeringPolicyResource()),
			},
		},
	}
}

func readDnsSteeringPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &DnsSteeringPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient()

	return ReadResource(sync)
}

type DnsSteeringPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.ListSteeringPoliciesResponse
}

func (s *DnsSteeringPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsSteeringPoliciesDataSourceCrud) Get() error {
	request := oci_dns.ListSteeringPoliciesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if healthCheckMonitorId, ok := s.D.GetOkExists("health_check_monitor_id"); ok {
		tmp := healthCheckMonitorId.(string)
		request.HealthCheckMonitorId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dns.SteeringPolicySummaryLifecycleStateEnum(state.(string))
	}

	if template, ok := s.D.GetOkExists("template"); ok {
		tmp := template.(string)
		request.Template = &tmp
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "dns")

	response, err := s.Client.ListSteeringPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSteeringPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DnsSteeringPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		steeringPolicy := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			steeringPolicy["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			steeringPolicy["display_name"] = *r.DisplayName
		}

		steeringPolicy["freeform_tags"] = r.FreeformTags

		if r.HealthCheckMonitorId != nil {
			steeringPolicy["health_check_monitor_id"] = *r.HealthCheckMonitorId
		}

		if r.Id != nil {
			steeringPolicy["id"] = *r.Id
		}

		if r.Self != nil {
			steeringPolicy["self"] = *r.Self
		}

		steeringPolicy["state"] = r.LifecycleState

		steeringPolicy["template"] = r.Template

		if r.TimeCreated != nil {
			steeringPolicy["time_created"] = r.TimeCreated.String()
		}

		if r.Ttl != nil {
			steeringPolicy["ttl"] = *r.Ttl
		}

		resources = append(resources, steeringPolicy)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DnsSteeringPoliciesDataSource().Schema["steering_policies"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("steering_policies", resources); err != nil {
		return err
	}

	return nil
}
