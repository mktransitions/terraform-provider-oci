// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"bytes"
	"context"
	"fmt"
	"github.com/oracle/terraform-provider-oci/internal/utils"
	"io"
	"io/ioutil"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"
)

func LogAnalyticsLogAnalyticsImportCustomContentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsLogAnalyticsImportCustomContent,
		Read:     readLogAnalyticsLogAnalyticsImportCustomContent,
		Delete:   deleteLogAnalyticsLogAnalyticsImportCustomContent,
		Schema: map[string]*schema.Schema{
			// Required
			"import_custom_content_file": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"expect": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_overwrite": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"change_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"conflict_field_display_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"conflict_parser_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"conflict_source_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"created_field_display_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"created_parser_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"created_source_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"updated_field_display_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"updated_parser_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"updated_source_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"content_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"field_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"parser_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createLogAnalyticsLogAnalyticsImportCustomContent(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsImportCustomContentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsLogAnalyticsImportCustomContent(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteLogAnalyticsLogAnalyticsImportCustomContent(d *schema.ResourceData, m interface{}) error {
	return nil
}

type LogAnalyticsLogAnalyticsImportCustomContentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.LogAnalyticsImportCustomContent
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsLogAnalyticsImportCustomContentResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("LogAnalyticsLogAnalyticsImportCustomContentResource-", LogAnalyticsLogAnalyticsImportCustomContentResource(), s.D)
}

func (s *LogAnalyticsLogAnalyticsImportCustomContentResourceCrud) Create() error {
	request := oci_log_analytics.ImportCustomContentRequest{}

	if importFile, ok := s.D.GetOkExists("import_custom_content_file"); ok {
		tmp := importFile.(string)
		contents, err := ioutil.ReadFile(tmp)
		if err != nil {
			return fmt.Errorf("the specified content file is not available: %q", err)
		}
		utils.Debugf("Length...........%d", len(contents))
		body := io.NopCloser(bytes.NewReader(contents))
		request.ImportCustomContentFileBody = body
	}

	if expect, ok := s.D.GetOkExists("expect"); ok {
		tmp := expect.(string)
		request.Expect = &tmp
	}

	if isOverwrite, ok := s.D.GetOkExists("is_overwrite"); ok {
		tmp := isOverwrite.(bool)
		request.IsOverwrite = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.ImportCustomContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LogAnalyticsImportCustomContent
	return nil
}

func (s *LogAnalyticsLogAnalyticsImportCustomContentResourceCrud) SetData() error {
	if s.Res.ChangeList != nil {
		s.D.Set("change_list", []interface{}{LogAnalyticsImportCustomChangeListToMap(s.Res.ChangeList)})
	} else {
		s.D.Set("change_list", nil)
	}

	if s.Res.ContentName != nil {
		s.D.Set("content_name", *s.Res.ContentName)
	}

	s.D.Set("field_names", s.Res.FieldNames)

	s.D.Set("parser_names", s.Res.ParserNames)

	s.D.Set("source_names", s.Res.SourceNames)

	return nil
}

func LogAnalyticsImportCustomChangeListToMap(obj *oci_log_analytics.LogAnalyticsImportCustomChangeList) map[string]interface{} {
	result := map[string]interface{}{}

	result["conflict_field_display_names"] = obj.ConflictFieldDisplayNames

	result["conflict_parser_names"] = obj.ConflictParserNames

	result["conflict_source_names"] = obj.ConflictSourceNames

	result["created_field_display_names"] = obj.CreatedFieldDisplayNames

	result["created_parser_names"] = obj.CreatedParserNames

	result["created_source_names"] = obj.CreatedSourceNames

	result["updated_field_display_names"] = obj.UpdatedFieldDisplayNames

	result["updated_parser_names"] = obj.UpdatedParserNames

	result["updated_source_names"] = obj.UpdatedSourceNames

	return result
}
