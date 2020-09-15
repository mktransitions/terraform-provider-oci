// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v25/usageapi"
)

func init() {
	RegisterDatasource("oci_metering_computation_configuration", MeteringComputationConfigurationDataSource())
}

func MeteringComputationConfigurationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMeteringComputationConfiguration,
		Schema: map[string]*schema.Schema{
			"tenant_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"values": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func readSingularMeteringComputationConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).usageapiClient()

	return ReadResource(sync)
}

type MeteringComputationConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.RequestSummarizedConfigurationsResponse
}

func (s *MeteringComputationConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationConfigurationDataSourceCrud) Get() error {
	request := oci_metering_computation.RequestSummarizedConfigurationsRequest{}

	if tenantId, ok := s.D.GetOkExists("tenant_id"); ok {
		tmp := tenantId.(string)
		request.TenantId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "metering_computation")

	response, err := s.Client.RequestSummarizedConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MeteringComputationConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ConfigurationToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func ConfigurationToMap(obj oci_metering_computation.Configuration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["values"] = obj.Values

	return result
}
