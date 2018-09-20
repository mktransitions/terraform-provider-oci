// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func AutonomousDataWarehouseDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularAutonomousDataWarehouse,
		Schema: map[string]*schema.Schema{
			"autonomous_data_warehouse_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_strings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"high": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"low": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"medium": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"license_model": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_console_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularAutonomousDataWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDataWarehouseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type AutonomousDataWarehouseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousDataWarehouseResponse
}

func (s *AutonomousDataWarehouseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AutonomousDataWarehouseDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousDataWarehouseRequest{}

	if autonomousDataWarehouseId, ok := s.D.GetOkExists("autonomous_data_warehouse_id"); ok {
		tmp := autonomousDataWarehouseId.(string)
		request.AutonomousDataWarehouseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousDataWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AutonomousDataWarehouseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{AutonomousDataWarehouseConnectionStringsToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ServiceConsoleUrl != nil {
		s.D.Set("service_console_url", *s.Res.ServiceConsoleUrl)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
