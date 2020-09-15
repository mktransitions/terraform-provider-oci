// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_nosql "github.com/oracle/oci-go-sdk/v25/nosql"
)

func init() {
	RegisterDatasource("oci_nosql_indexes", NosqlIndexesDataSource())
}

func NosqlIndexesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNosqlIndexes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"table_name_or_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(NosqlIndexResource()),
			},
		},
	}
}

func readNosqlIndexes(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlIndexesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).nosqlClient()

	return ReadResource(sync)
}

type NosqlIndexesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_nosql.NosqlClient
	Res    *oci_nosql.ListIndexesResponse
}

func (s *NosqlIndexesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NosqlIndexesDataSourceCrud) Get() error {
	request := oci_nosql.ListIndexesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_nosql.ListIndexesLifecycleStateEnum(state.(string))
	}

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.TableNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "nosql")

	response, err := s.Client.ListIndexes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NosqlIndexesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	resources := []map[string]interface{}{}
	for _, item := range s.Res.Items {
		resources = append(resources, IndexSummaryToMap(item))
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, NosqlIndexesDataSource().Schema["index_collection"].Elem.(*schema.Resource).Schema)
	}
	s.D.Set("index_collection", resources)

	return nil
}
