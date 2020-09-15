// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v25/datacatalog"
)

func init() {
	RegisterDatasource("oci_datacatalog_catalog", DatacatalogCatalogDataSource())
}

func DatacatalogCatalogDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["catalog_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DatacatalogCatalogResource(), fieldMap, readSingularDatacatalogCatalog)
}

func readSingularDatacatalogCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataCatalogClient()

	return ReadResource(sync)
}

type DatacatalogCatalogDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacatalog.DataCatalogClient
	Res    *oci_datacatalog.GetCatalogResponse
}

func (s *DatacatalogCatalogDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatacatalogCatalogDataSourceCrud) Get() error {
	request := oci_datacatalog.GetCatalogRequest{}

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "datacatalog")

	response, err := s.Client.GetCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatacatalogCatalogDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("attached_catalog_private_endpoints", s.Res.AttachedCatalogPrivateEndpoints)

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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NumberOfObjects != nil {
		s.D.Set("number_of_objects", *s.Res.NumberOfObjects)
	}

	if s.Res.ServiceApiUrl != nil {
		s.D.Set("service_api_url", *s.Res.ServiceApiUrl)
	}

	if s.Res.ServiceConsoleUrl != nil {
		s.D.Set("service_console_url", *s.Res.ServiceConsoleUrl)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
