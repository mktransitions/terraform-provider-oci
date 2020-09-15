// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_logging "github.com/oracle/oci-go-sdk/v25/logging"
)

func init() {
	RegisterDatasource("oci_logging_logs", LoggingLogsDataSource())
}

func LoggingLogsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoggingLogs,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_resource": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_service": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"logs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(LoggingLogResource()),
			},
		},
	}
}

func readLoggingLogs(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loggingManagementClient()

	return ReadResource(sync)
}

type LoggingLogsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_logging.LoggingManagementClient
	Res    *oci_logging.ListLogsResponse
}

func (s *LoggingLogsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoggingLogsDataSourceCrud) Get() error {
	request := oci_logging.ListLogsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if logGroupId, ok := s.D.GetOkExists("log_group_id"); ok {
		tmp := logGroupId.(string)
		request.LogGroupId = &tmp
	}

	if logType, ok := s.D.GetOkExists("log_type"); ok {
		request.LogType = oci_logging.ListLogsLogTypeEnum(logType.(string))
	}

	if sourceResource, ok := s.D.GetOkExists("source_resource"); ok {
		tmp := sourceResource.(string)
		request.SourceResource = &tmp
	}

	if sourceService, ok := s.D.GetOkExists("source_service"); ok {
		tmp := sourceService.(string)
		request.SourceService = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_logging.ListLogsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "logging")

	response, err := s.Client.ListLogs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLogs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LoggingLogsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		log := map[string]interface{}{
			"log_group_id": *r.LogGroupId,
		}

		if r.CompartmentId != nil {
			log["compartment_id"] = *r.CompartmentId
		}

		if r.Configuration != nil {
			log["configuration"] = []interface{}{LogConfigurationToMap(r.Configuration)}
		} else {
			log["configuration"] = nil
		}

		if r.DefinedTags != nil {
			log["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			log["display_name"] = *r.DisplayName
		}

		log["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			log["id"] = *r.Id
		}

		if r.IsEnabled != nil {
			log["is_enabled"] = *r.IsEnabled
		}

		log["log_type"] = r.LogType

		if r.RetentionDuration != nil {
			log["retention_duration"] = *r.RetentionDuration
		}

		log["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			log["time_created"] = r.TimeCreated.String()
		}

		if r.TimeLastModified != nil {
			log["time_last_modified"] = r.TimeLastModified.String()
		}

		resources = append(resources, log)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, LoggingLogsDataSource().Schema["logs"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("logs", resources); err != nil {
		return err
	}

	return nil
}
