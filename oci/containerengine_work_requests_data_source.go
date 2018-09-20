// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
)

func WorkRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWorkRequests,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"work_requests": {
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
						"operation_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resources": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"action_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entity_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"entity_uri": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"identifier": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_accepted": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_finished": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readWorkRequests(d *schema.ResourceData, m interface{}) error {
	sync := &WorkRequestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return ReadResource(sync)
}

type WorkRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListWorkRequestsResponse
}

func (s *WorkRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WorkRequestsDataSourceCrud) Get() error {
	request := oci_containerengine.ListWorkRequestsRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = oci_containerengine.ListWorkRequestsResourceTypeEnum(tmp)
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		interfaces := status.([]interface{})
		tmp := make([]oci_containerengine.ListWorkRequestsStatusEnum, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = oci_containerengine.ListWorkRequestsStatusEnum(toBeConverted.(string))
		}
		request.Status = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "containerengine")

	response, err := s.Client.ListWorkRequests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWorkRequests(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WorkRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		workRequest := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Id != nil {
			workRequest["id"] = *r.Id
		}

		workRequest["operation_type"] = r.OperationType

		_resources := []interface{}{}
		for _, item := range r.Resources {
			_resources = append(_resources, WorkRequestResourceToMap(item))
		}
		workRequest["resources"] = _resources

		workRequest["status"] = r.Status

		if r.TimeAccepted != nil {
			workRequest["time_accepted"] = r.TimeAccepted.String()
		}

		if r.TimeFinished != nil {
			workRequest["time_finished"] = r.TimeFinished.String()
		}

		if r.TimeStarted != nil {
			workRequest["time_started"] = r.TimeStarted.String()
		}

		resources = append(resources, workRequest)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, WorkRequestsDataSource().Schema["work_requests"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("work_requests", resources); err != nil {
		return err
	}

	return nil
}

func WorkRequestResourceToMap(obj oci_containerengine.WorkRequestResource) map[string]interface{} {
	result := map[string]interface{}{}

	result["action_type"] = string(obj.ActionType)

	if obj.EntityType != nil {
		result["entity_type"] = string(*obj.EntityType)
	}

	if obj.EntityUri != nil {
		result["entity_uri"] = string(*obj.EntityUri)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}
