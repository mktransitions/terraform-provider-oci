// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v25/datasafe"
)

func init() {
	RegisterDatasource("oci_data_safe_data_safe_private_endpoint", DataSafeDataSafePrivateEndpointDataSource())
}

func DataSafeDataSafePrivateEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["data_safe_private_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DataSafeDataSafePrivateEndpointResource(), fieldMap, readSingularDataSafeDataSafePrivateEndpoint)
}

func readSingularDataSafeDataSafePrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDataSafePrivateEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataSafeClient()

	return ReadResource(sync)
}

type DataSafeDataSafePrivateEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetDataSafePrivateEndpointResponse
}

func (s *DataSafeDataSafePrivateEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeDataSafePrivateEndpointDataSourceCrud) Get() error {
	request := oci_data_safe.GetDataSafePrivateEndpointRequest{}

	if dataSafePrivateEndpointId, ok := s.D.GetOkExists("data_safe_private_endpoint_id"); ok {
		tmp := dataSafePrivateEndpointId.(string)
		request.DataSafePrivateEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "data_safe")

	response, err := s.Client.GetDataSafePrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeDataSafePrivateEndpointDataSourceCrud) SetData() error {
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

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EndpointFqdn != nil {
		s.D.Set("endpoint_fqdn", *s.Res.EndpointFqdn)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.PrivateEndpointId != nil {
		s.D.Set("private_endpoint_id", *s.Res.PrivateEndpointId)
	}

	if s.Res.PrivateEndpointIp != nil {
		s.D.Set("private_endpoint_ip", *s.Res.PrivateEndpointIp)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
