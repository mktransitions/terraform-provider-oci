// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v25/mysql"
)

func init() {
	RegisterDatasource("oci_mysql_mysql_versions", MysqlMysqlVersionsDataSource())
}

func MysqlMysqlVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMysqlMysqlVersions,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"version_family": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"versions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readMysqlMysqlVersions(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).mysqlaasClient()

	return ReadResource(sync)
}

type MysqlMysqlVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.MysqlaasClient
	Res    *oci_mysql.ListVersionsResponse
}

func (s *MysqlMysqlVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlMysqlVersionsDataSourceCrud) Get() error {
	request := oci_mysql.ListVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "mysql")

	response, err := s.Client.ListVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MysqlMysqlVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		mysqlVersion := map[string]interface{}{}

		if r.VersionFamily != nil {
			mysqlVersion["version_family"] = *r.VersionFamily
		}

		versions := []interface{}{}
		for _, item := range r.Versions {
			versions = append(versions, VersionToMap(item))
		}
		mysqlVersion["versions"] = versions

		resources = append(resources, mysqlVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, MysqlMysqlVersionsDataSource().Schema["versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("versions", resources); err != nil {
		return err
	}

	return nil
}

func VersionToMap(obj oci_mysql.Version) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
