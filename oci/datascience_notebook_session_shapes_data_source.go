// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v25/datascience"
)

func init() {
	RegisterDatasource("oci_datascience_notebook_session_shapes", DatascienceNotebookSessionShapesDataSource())
}

func DatascienceNotebookSessionShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceNotebookSessionShapes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notebook_session_shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"core_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"memory_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatascienceNotebookSessionShapes(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceNotebookSessionShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	return ReadResource(sync)
}

type DatascienceNotebookSessionShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListNotebookSessionShapesResponse
}

func (s *DatascienceNotebookSessionShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceNotebookSessionShapesDataSourceCrud) Get() error {
	request := oci_datascience.ListNotebookSessionShapesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "datascience")

	response, err := s.Client.ListNotebookSessionShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNotebookSessionShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceNotebookSessionShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		notebookSessionShape := map[string]interface{}{}

		if r.CoreCount != nil {
			notebookSessionShape["core_count"] = *r.CoreCount
		}

		if r.MemoryInGBs != nil {
			notebookSessionShape["memory_in_gbs"] = *r.MemoryInGBs
		}

		if r.Name != nil {
			notebookSessionShape["name"] = *r.Name
		}

		resources = append(resources, notebookSessionShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatascienceNotebookSessionShapesDataSource().Schema["notebook_session_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("notebook_session_shapes", resources); err != nil {
		return err
	}

	return nil
}
