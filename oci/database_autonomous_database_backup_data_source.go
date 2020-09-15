// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v25/database"
)

func init() {
	RegisterDatasource("oci_database_autonomous_database_backup", DatabaseAutonomousDatabaseBackupDataSource())
}

func DatabaseAutonomousDatabaseBackupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["autonomous_database_backup_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DatabaseAutonomousDatabaseBackupResource(), fieldMap, readSingularDatabaseAutonomousDatabaseBackup)
}

func readSingularDatabaseAutonomousDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseAutonomousDatabaseBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousDatabaseBackupResponse
}

func (s *DatabaseAutonomousDatabaseBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseBackupDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseBackupRequest{}

	if autonomousDatabaseBackupId, ok := s.D.GetOkExists("autonomous_database_backup_id"); ok {
		tmp := autonomousDatabaseBackupId.(string)
		request.AutonomousDatabaseBackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousDatabaseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousDatabaseBackupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AutonomousDatabaseId != nil {
		s.D.Set("autonomous_database_id", *s.Res.AutonomousDatabaseId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseSizeInTBs != nil {
		s.D.Set("database_size_in_tbs", *s.Res.DatabaseSizeInTBs)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IsAutomatic != nil {
		s.D.Set("is_automatic", *s.Res.IsAutomatic)
	}

	if s.Res.IsRestorable != nil {
		s.D.Set("is_restorable", *s.Res.IsRestorable)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
