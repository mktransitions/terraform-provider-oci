// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v25/database"
)

func init() {
	RegisterDatasource("oci_database_backups", DatabaseBackupsDataSource())
}

func DatabaseBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseBackups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DatabaseBackupResource()),
			},
		},
	}
}

func readDatabaseBackups(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListBackupsResponse
}

func (s *DatabaseBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseBackupsDataSourceCrud) Get() error {
	request := oci_database.ListBackupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		backup := map[string]interface{}{}

		if r.AvailabilityDomain != nil {
			backup["availability_domain"] = *r.AvailabilityDomain
		}

		if r.CompartmentId != nil {
			backup["compartment_id"] = *r.CompartmentId
		}

		backup["database_edition"] = r.DatabaseEdition

		if r.DatabaseId != nil {
			backup["database_id"] = *r.DatabaseId
		}

		if r.DatabaseSizeInGBs != nil {
			backup["database_size_in_gbs"] = *r.DatabaseSizeInGBs
		}

		if r.DisplayName != nil {
			backup["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			backup["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			backup["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.Shape != nil {
			backup["shape"] = *r.Shape
		}

		backup["state"] = r.LifecycleState

		if r.TimeEnded != nil {
			backup["time_ended"] = r.TimeEnded.Format(time.RFC3339Nano)
		}

		if r.TimeStarted != nil {
			backup["time_started"] = r.TimeStarted.Format(time.RFC3339Nano)
		}

		backup["type"] = r.Type

		if r.Version != nil {
			backup["version"] = *r.Version
		}

		resources = append(resources, backup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseBackupsDataSource().Schema["backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("backups", resources); err != nil {
		return err
	}

	return nil
}
