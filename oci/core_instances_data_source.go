// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v25/core"
)

func init() {
	RegisterDatasource("oci_core_instances", CoreInstancesDataSource())
}

func CoreInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreInstances,
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CoreInstanceResource()),
			},
		},
	}
}

func readCoreInstances(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()

	return ReadResource(sync)
}

type CoreInstancesDataSourceCrud struct {
	BaseCrud
	Client *oci_core.ComputeClient
	Res    *oci_core.ListInstancesResponse
}

func (s *CoreInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstancesDataSourceCrud) Get() error {
	request := oci_core.ListInstancesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.InstanceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AgentConfig != nil {
			instance["agent_config"] = []interface{}{InstanceAgentConfigToMap(r.AgentConfig)}
		} else {
			instance["agent_config"] = nil
		}

		if r.AvailabilityConfig != nil {
			instance["availability_config"] = []interface{}{InstanceAvailabilityConfigToMap(r.AvailabilityConfig)}
		} else {
			instance["availability_config"] = nil
		}

		if r.AvailabilityDomain != nil {
			instance["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DedicatedVmHostId != nil {
			instance["dedicated_vm_host_id"] = *r.DedicatedVmHostId
		}

		if r.DefinedTags != nil {
			instance["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			instance["display_name"] = *r.DisplayName
		}

		if r.ExtendedMetadata != nil {
			instance["extended_metadata"] = convertNestedMapToFlatMap(r.ExtendedMetadata)
		}

		if r.FaultDomain != nil {
			instance["fault_domain"] = *r.FaultDomain
		}

		instance["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			instance["id"] = *r.Id
		}

		if r.ImageId != nil {
			instance["image"] = *r.ImageId
		}

		if r.IpxeScript != nil {
			instance["ipxe_script"] = *r.IpxeScript
		}

		instance["launch_mode"] = r.LaunchMode

		if r.LaunchOptions != nil {
			instance["launch_options"] = []interface{}{LaunchOptionsToMap(r.LaunchOptions)}
		} else {
			instance["launch_options"] = nil
		}

		if r.Metadata != nil {
			instance["metadata"] = r.Metadata
		}

		if r.Region != nil {
			instance["region"] = *r.Region
		}

		if r.Shape != nil {
			instance["shape"] = *r.Shape
		}

		if r.ShapeConfig != nil {
			instance["shape_config"] = []interface{}{InstanceShapeConfigToMap(r.ShapeConfig)}
		} else {
			instance["shape_config"] = nil
		}

		if r.SourceDetails != nil {
			sourceDetailsArray := []interface{}{}
			if sourceDetailsMap := InstanceSourceDetailsToMap(&r.SourceDetails, nil, nil); sourceDetailsMap != nil {
				sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
			}
			instance["source_details"] = sourceDetailsArray
		} else {
			instance["source_details"] = nil
		}

		instance["state"] = r.LifecycleState

		if r.SystemTags != nil {
			instance["system_tags"] = systemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			instance["time_created"] = r.TimeCreated.String()
		}

		if r.TimeMaintenanceRebootDue != nil {
			instance["time_maintenance_reboot_due"] = r.TimeMaintenanceRebootDue.String()
		}

		resources = append(resources, instance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreInstancesDataSource().Schema["instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instances", resources); err != nil {
		return err
	}

	return nil
}

func convertNestedMapToFlatMap(m map[string]interface{}) map[string]string {
	flatMap := make(map[string]string)
	var ok bool
	for key, val := range m {
		if flatMap[key], ok = val.(string); !ok {
			mapValStr, err := json.Marshal(val)
			if err != nil {
				mapValStr = []byte{}
			}
			flatMap[key] = string(mapValStr)
		}
	}
	return flatMap
}
