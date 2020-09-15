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
	RegisterResource("oci_database_autonomous_database_backup", DatabaseAutonomousDatabaseBackupResource())
}

func DatabaseAutonomousDatabaseBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatabaseAutonomousDatabaseBackup,
		Read:     readDatabaseAutonomousDatabaseBackup,
		Delete:   deleteDatabaseAutonomousDatabaseBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"is_automatic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_restorable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseAutonomousDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return CreateResource(d, sync)
}

func readDatabaseAutonomousDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

func deleteDatabaseAutonomousDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseAutonomousDatabaseBackupResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousDatabaseBackup
	DisableNotFoundRetries bool
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseBackupLifecycleStateCreating),
	}
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseBackupLifecycleStateActive),
	}
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseBackupLifecycleStateDeleting),
	}
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseBackupLifecycleStateDeleted),
	}
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) Create() error {
	request := oci_database.CreateAutonomousDatabaseBackupRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousDatabaseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabaseBackup
	return nil
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseBackupRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousDatabaseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabaseBackup
	return nil
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) SetData() error {
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
		s.D.Set("time_ended", s.Res.TimeEnded.Format(time.RFC3339Nano))
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.Format(time.RFC3339Nano))
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
