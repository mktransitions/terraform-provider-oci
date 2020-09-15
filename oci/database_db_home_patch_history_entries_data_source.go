// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v25/database"
)

func init() {
	RegisterDatasource("oci_database_db_home_patch_history_entries", DatabaseDbHomePatchHistoryEntriesDataSource())
}

func DatabaseDbHomePatchHistoryEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbHomePatchHistoryEntries,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"db_home_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"patch_history_entries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"patch_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseDbHomePatchHistoryEntries(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbHomePatchHistoryEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseDbHomePatchHistoryEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbHomePatchHistoryEntriesResponse
}

func (s *DatabaseDbHomePatchHistoryEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbHomePatchHistoryEntriesDataSourceCrud) Get() error {
	request := oci_database.ListDbHomePatchHistoryEntriesRequest{}

	if dbHomeId, ok := s.D.GetOkExists("db_home_id"); ok {
		tmp := dbHomeId.(string)
		request.DbHomeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListDbHomePatchHistoryEntries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbHomePatchHistoryEntries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbHomePatchHistoryEntriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbHomePatchHistoryEntry := map[string]interface{}{}

		dbHomePatchHistoryEntry["action"] = r.Action

		if r.Id != nil {
			dbHomePatchHistoryEntry["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			dbHomePatchHistoryEntry["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.PatchId != nil {
			dbHomePatchHistoryEntry["patch_id"] = *r.PatchId
		}

		dbHomePatchHistoryEntry["state"] = r.LifecycleState

		if r.TimeEnded != nil {
			dbHomePatchHistoryEntry["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeStarted != nil {
			dbHomePatchHistoryEntry["time_started"] = r.TimeStarted.String()
		}

		resources = append(resources, dbHomePatchHistoryEntry)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseDbHomePatchHistoryEntriesDataSource().Schema["patch_history_entries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("patch_history_entries", resources); err != nil {
		return err
	}

	return nil
}
