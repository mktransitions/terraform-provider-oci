// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v25/common"
	oci_nosql "github.com/oracle/oci-go-sdk/v25/nosql"
)

func init() {
	RegisterResource("oci_nosql_table", NosqlTableResource())
}

func NosqlTableResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createNosqlTable,
		Read:     readNosqlTable,
		Update:   updateNosqlTable,
		Delete:   deleteNosqlTable,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ddl_statement": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"table_limits": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"max_read_units": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"max_storage_in_gbs": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"max_write_units": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schema": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"columns": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"default_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_nullable": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"primary_key": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"shard_key": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ttl": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createNosqlTable(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).nosqlClient()

	return CreateResource(d, sync)
}

func readNosqlTable(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).nosqlClient()

	return ReadResource(sync)
}

func updateNosqlTable(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).nosqlClient()

	return UpdateResource(d, sync)
}

func deleteNosqlTable(d *schema.ResourceData, m interface{}) error {
	sync := &NosqlTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).nosqlClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type NosqlTableResourceCrud struct {
	BaseCrud
	Client                 *oci_nosql.NosqlClient
	Res                    *oci_nosql.Table
	DisableNotFoundRetries bool
}

func (s *NosqlTableResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *NosqlTableResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_nosql.TableLifecycleStateCreating),
	}
}

func (s *NosqlTableResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_nosql.TableLifecycleStateActive),
	}
}

func (s *NosqlTableResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_nosql.TableLifecycleStateDeleting),
	}
}

func (s *NosqlTableResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_nosql.TableLifecycleStateDeleted),
	}
}

func (s *NosqlTableResourceCrud) Create() error {
	request := oci_nosql.CreateTableRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if ddlStatement, ok := s.D.GetOkExists("ddl_statement"); ok {
		tmp := ddlStatement.(string)
		request.DdlStatement = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.Name = &tmp
	}

	if tableLimits, ok := s.D.GetOkExists("table_limits"); ok {
		if tmpList := tableLimits.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "table_limits", 0)
			tmp, err := s.mapToTableLimits(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TableLimits = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.CreateTable(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getTableFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "nosql"), oci_nosql.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *NosqlTableResourceCrud) getTableFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_nosql.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	tableId, err := tableWaitForWorkRequest(workId, "TABLE",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, tableId)
		_, cancelErr := s.Client.DeleteWorkRequest(context.Background(),
			oci_nosql.DeleteWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}

	// For update, we send multiple requests and we don't want to override the state file for each request
	if actionTypeEnum == oci_nosql.WorkRequestResourceActionTypeUpdated {
		return nil
	}
	s.D.SetId(*tableId)
	s.D.Set("table_name_or_id", *tableId)

	return s.Get()
}

func tableWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "nosql", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_nosql.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func tableWaitForWorkRequest(wId *string, entityType string, action oci_nosql.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_nosql.NosqlClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "nosql")
	retryPolicy.ShouldRetryOperation = tableWorkRequestShouldRetryFunc(timeout)

	response := oci_nosql.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_nosql.WorkRequestStatusInProgress),
			string(oci_nosql.WorkRequestStatusAccepted),
			string(oci_nosql.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_nosql.WorkRequestStatusSucceeded),
			string(oci_nosql.WorkRequestStatusFailed),
			string(oci_nosql.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_nosql.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}
	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(*res.EntityType, entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest didn't do all its intended tasks, if the errors is set; so we should check for it
	if identifier == nil || response.WorkRequest.Status == oci_nosql.WorkRequestStatusFailed || response.WorkRequest.Status == oci_nosql.WorkRequestStatusCanceled {
		return nil, getErrorFromTableWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromTableWorkRequest(client *oci_nosql.NosqlClient, wId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_nosql.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_nosql.ListWorkRequestErrorsRequest{
			WorkRequestId: wId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *NosqlTableResourceCrud) Get() error {
	request := oci_nosql.GetTableRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.TableNameOrId = &tmp
	} else if s.D.Id() != "" {
		tmp := s.D.Id()
		request.TableNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.GetTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Table
	return nil
}

func (s *NosqlTableResourceCrud) Update() error {
	if _, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			fromCompartmentId := oldRaw.(string)
			toCompartmentId := newRaw.(string)
			err := s.updateCompartment(fromCompartmentId, toCompartmentId)
			if err != nil {
				return err
			}
		}
	}

	defer func() {
		// get latest state of the instance
		err := s.Get()
		if err != nil {
			log.Printf("[ERROR] unable to invoke GET() after UPDATE '%v'", err)
		}
		// write latest state
		if err := s.SetData(); err != nil {
			log.Printf("[ERROR] unable to invoke setData() '%v'", err)
		}
	}()

	request := oci_nosql.UpdateTableRequest{}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "nosql")

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.TableNameOrId = &tmp
	} else if s.D.Id() != "" {
		tmp := s.D.Id()
		request.TableNameOrId = &tmp
	}

	if ddlStatement, ok := s.D.GetOkExists("ddl_statement"); ok && s.D.HasChange("ddl_statement") {
		tmp := ddlStatement.(string)
		request.DdlStatement = &tmp
		err := sendUpdateRequest(s, request)
		if err != nil {
			return err
		}
		request.DdlStatement = nil
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok && s.D.HasChange("defined_tags") {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
		err = sendUpdateRequest(s, request)
		if err != nil {
			return err
		}
		request.DefinedTags = nil
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
		err := sendUpdateRequest(s, request)
		if err != nil {
			return err
		}
		request.FreeformTags = nil
	}

	if tableLimits, ok := s.D.GetOkExists("table_limits"); ok && s.D.HasChange("table_limits") {
		if tmpList := tableLimits.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "table_limits", 0)
			tmp, err := s.mapToTableLimits(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TableLimits = &tmp
		}
		err := sendUpdateRequest(s, request)
		if err != nil {
			return err
		}
		request.TableLimits = nil
	}

	return nil
}

func sendUpdateRequest(s *NosqlTableResourceCrud, request oci_nosql.UpdateTableRequest) error {
	response, err := s.Client.UpdateTable(context.Background(), request)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	err = s.getTableFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "nosql"), oci_nosql.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	return nil
}

func (s *NosqlTableResourceCrud) Delete() error {
	request := oci_nosql.DeleteTableRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isIfExists, ok := s.D.GetOkExists("is_if_exists"); ok {
		tmp := isIfExists.(bool)
		request.IsIfExists = &tmp
	}

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		request.TableNameOrId = &tmp
	} else if s.D.Id() != "" {
		tmp := s.D.Id()
		request.TableNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.DeleteTable(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := tableWaitForWorkRequest(workId, "TABLE",
		oci_nosql.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *NosqlTableResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DdlStatement != nil {
		s.D.Set("ddl_statement", *s.Res.DdlStatement)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Id != nil {
		s.D.SetId(*s.Res.Id)
	}

	if s.Res.Schema != nil {
		s.D.Set("schema", []interface{}{SchemaToMap(s.Res.Schema)})
	} else {
		s.D.Set("schema", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TableLimits != nil {
		s.D.Set("table_limits", []interface{}{TableLimitsToMap(s.Res.TableLimits)})
	} else {
		s.D.Set("table_limits", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ColumnToMap(obj oci_nosql.Column) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultValue != nil {
		result["default_value"] = string(*obj.DefaultValue)
	}

	if obj.IsNullable != nil {
		result["is_nullable"] = bool(*obj.IsNullable)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func SchemaToMap(obj *oci_nosql.Schema) map[string]interface{} {
	result := map[string]interface{}{}

	columns := []interface{}{}
	for _, item := range obj.Columns {
		columns = append(columns, ColumnToMap(item))
	}
	result["columns"] = columns

	result["primary_key"] = obj.PrimaryKey

	result["shard_key"] = obj.ShardKey

	if obj.Ttl != nil {
		result["ttl"] = int(*obj.Ttl)
	}

	return result
}

func (s *NosqlTableResourceCrud) mapToTableLimits(fieldKeyFormat string) (oci_nosql.TableLimits, error) {
	result := oci_nosql.TableLimits{}

	if maxReadUnits, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_read_units")); ok {
		tmp := maxReadUnits.(int)
		result.MaxReadUnits = &tmp
	}

	if maxStorageInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_storage_in_gbs")); ok {
		tmp := maxStorageInGBs.(int)
		result.MaxStorageInGBs = &tmp
	}

	if maxWriteUnits, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_write_units")); ok {
		tmp := maxWriteUnits.(int)
		result.MaxWriteUnits = &tmp
	}

	return result, nil
}

func TableLimitsToMap(obj *oci_nosql.TableLimits) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaxReadUnits != nil {
		result["max_read_units"] = int(*obj.MaxReadUnits)
	}

	if obj.MaxStorageInGBs != nil {
		result["max_storage_in_gbs"] = int(*obj.MaxStorageInGBs)
	}

	if obj.MaxWriteUnits != nil {
		result["max_write_units"] = int(*obj.MaxWriteUnits)
	}

	return result
}

func TableSummaryToMap(obj oci_nosql.TableSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TableLimits != nil {
		result["table_limits"] = []interface{}{TableLimitsToMap(obj.TableLimits)}
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *NosqlTableResourceCrud) updateCompartment(fromCompartmentId, toCompartmentId string) error {
	changeCompartmentRequest := oci_nosql.ChangeTableCompartmentRequest{}

	changeCompartmentRequest.FromCompartmentId = &fromCompartmentId

	if tableNameOrId, ok := s.D.GetOkExists("table_name_or_id"); ok {
		tmp := tableNameOrId.(string)
		changeCompartmentRequest.TableNameOrId = &tmp
	} else if s.D.Id() != "" {
		tmp := s.D.Id()
		changeCompartmentRequest.TableNameOrId = &tmp
	}

	changeCompartmentRequest.ToCompartmentId = &toCompartmentId

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "nosql")

	response, err := s.Client.ChangeTableCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getTableFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "nosql"), oci_nosql.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	return nil
}
