// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v25/containerengine"
)

func init() {
	RegisterDatasource("oci_containerengine_node_pool_option", ContainerengineNodePoolOptionDataSource())
}

func ContainerengineNodePoolOptionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularContainerengineNodePoolOption,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"node_pool_option_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"images": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"kubernetes_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"image_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularContainerengineNodePoolOption(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolOptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient()

	return ReadResource(sync)
}

type ContainerengineNodePoolOptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetNodePoolOptionsResponse
}

func (s *ContainerengineNodePoolOptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineNodePoolOptionDataSourceCrud) Get() error {
	request := oci_containerengine.GetNodePoolOptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if nodePoolOptionId, ok := s.D.GetOkExists("node_pool_option_id"); ok {
		tmp := nodePoolOptionId.(string)
		request.NodePoolOptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "containerengine")

	response, err := s.Client.GetNodePoolOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineNodePoolOptionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	s.D.Set("images", s.Res.Images)

	s.D.Set("kubernetes_versions", s.Res.KubernetesVersions)

	s.D.Set("shapes", s.Res.Shapes)

	sources := []interface{}{}
	for _, item := range s.Res.Sources {
		sources = append(sources, NodeSourceOptionToMap(&item))
	}
	s.D.Set("sources", sources)

	return nil
}
