// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v25/database"
)

func init() {
	RegisterDatasource("oci_database_autonomous_container_databases", DatabaseAutonomousContainerDatabasesDataSource())
}

func DatabaseAutonomousContainerDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousContainerDatabases,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"autonomous_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"infrastructure_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_container_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DatabaseAutonomousContainerDatabaseResource()),
			},
		},
	}
}

func readDatabaseAutonomousContainerDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseAutonomousContainerDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousContainerDatabasesResponse
}

func (s *DatabaseAutonomousContainerDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousContainerDatabasesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousContainerDatabasesRequest{}

	if autonomousExadataInfrastructureId, ok := s.D.GetOkExists("autonomous_exadata_infrastructure_id"); ok {
		tmp := autonomousExadataInfrastructureId.(string)
		request.AutonomousExadataInfrastructureId = &tmp
	}

	if autonomousVmClusterId, ok := s.D.GetOkExists("autonomous_vm_cluster_id"); ok {
		tmp := autonomousVmClusterId.(string)
		request.AutonomousVmClusterId = &tmp
	}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if infrastructureType, ok := s.D.GetOkExists("infrastructure_type"); ok {
		request.InfrastructureType = oci_database.AutonomousContainerDatabaseSummaryInfrastructureTypeEnum(infrastructureType.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousContainerDatabaseSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousContainerDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousContainerDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousContainerDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousContainerDatabase := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AutonomousExadataInfrastructureId != nil {
			autonomousContainerDatabase["autonomous_exadata_infrastructure_id"] = *r.AutonomousExadataInfrastructureId
		}

		if r.AutonomousVmClusterId != nil {
			autonomousContainerDatabase["autonomous_vm_cluster_id"] = *r.AutonomousVmClusterId
		}

		if r.AvailabilityDomain != nil {
			autonomousContainerDatabase["availability_domain"] = *r.AvailabilityDomain
		}

		if r.BackupConfig != nil {
			autonomousContainerDatabase["backup_config"] = []interface{}{AutonomousContainerDatabaseBackupConfigToMap(r.BackupConfig, nil, true)}
		} else {
			autonomousContainerDatabase["backup_config"] = nil
		}

		if r.DbUniqueName != nil {
			autonomousContainerDatabase["db_unique_name"] = *r.DbUniqueName
		}

		if r.DbVersion != nil {
			autonomousContainerDatabase["db_version"] = *r.DbVersion
		}

		if r.DefinedTags != nil {
			autonomousContainerDatabase["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			autonomousContainerDatabase["display_name"] = *r.DisplayName
		}

		autonomousContainerDatabase["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			autonomousContainerDatabase["id"] = *r.Id
		}

		autonomousContainerDatabase["infrastructure_type"] = r.InfrastructureType

		if r.LastMaintenanceRunId != nil {
			autonomousContainerDatabase["last_maintenance_run_id"] = *r.LastMaintenanceRunId
		}

		if r.LifecycleDetails != nil {
			autonomousContainerDatabase["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.MaintenanceWindow != nil {
			autonomousContainerDatabase["maintenance_window"] = []interface{}{MaintenanceWindowToMap(r.MaintenanceWindow)}
		} else {
			autonomousContainerDatabase["maintenance_window"] = nil
		}

		if r.NextMaintenanceRunId != nil {
			autonomousContainerDatabase["next_maintenance_run_id"] = *r.NextMaintenanceRunId
		}

		if r.PatchId != nil {
			autonomousContainerDatabase["patch_id"] = *r.PatchId
		}

		autonomousContainerDatabase["patch_model"] = r.PatchModel

		autonomousContainerDatabase["service_level_agreement_type"] = r.ServiceLevelAgreementType

		autonomousContainerDatabase["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			autonomousContainerDatabase["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, autonomousContainerDatabase)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousContainerDatabasesDataSource().Schema["autonomous_container_databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_container_databases", resources); err != nil {
		return err
	}

	return nil
}
