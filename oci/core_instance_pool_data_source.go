// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v25/core"
)

func init() {
	RegisterDatasource("oci_core_instance_pool", CoreInstancePoolDataSource())
}

func CoreInstancePoolDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["instance_pool_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(CoreInstancePoolResource(), fieldMap, readSingularCoreInstancePool)
}

func readSingularCoreInstancePool(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancePoolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient()

	return ReadResource(sync)
}

type CoreInstancePoolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.GetInstancePoolResponse
}

func (s *CoreInstancePoolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstancePoolDataSourceCrud) Get() error {
	request := oci_core.GetInstancePoolRequest{}

	if instancePoolId, ok := s.D.GetOkExists("instance_pool_id"); ok {
		tmp := instancePoolId.(string)
		request.InstancePoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetInstancePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreInstancePoolDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceConfigurationId != nil {
		s.D.Set("instance_configuration_id", *s.Res.InstanceConfigurationId)
	}

	loadBalancers := []interface{}{}
	for _, item := range s.Res.LoadBalancers {
		loadBalancers = append(loadBalancers, InstancePoolLoadBalancerAttachmentToMap(item))
	}
	s.D.Set("load_balancers", loadBalancers)

	placementConfigurations := []interface{}{}
	for _, item := range s.Res.PlacementConfigurations {
		placementConfigurations = append(placementConfigurations, InstancePoolPlacementConfigurationToMap(item))
	}
	s.D.Set("placement_configurations", placementConfigurations)

	if s.Res.Size != nil {
		s.D.Set("size", *s.Res.Size)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
