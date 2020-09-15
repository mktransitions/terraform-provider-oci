// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_datascience "github.com/oracle/oci-go-sdk/v25/datascience"
)

func init() {
	RegisterResource("oci_datascience_notebook_session", DatascienceNotebookSessionResource())
}

func DatascienceNotebookSessionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatascienceNotebookSession,
		Read:     readDatascienceNotebookSession,
		Update:   updateDatascienceNotebookSession,
		Delete:   deleteDatascienceNotebookSession,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notebook_session_configuration_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"block_storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_datascience.NotebookSessionLifecycleStateActive),
					string(oci_datascience.NotebookSessionLifecycleStateInactive),
				}, true),
			},

			// Computed
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"notebook_session_url": {
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

func createDatascienceNotebookSession(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceNotebookSessionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	var deactivateNotebookSession = false
	if state, ok := sync.D.GetOkExists("state"); ok {
		desiredState := oci_datascience.NotebookSessionLifecycleStateEnum(strings.ToUpper(state.(string)))
		if desiredState == oci_datascience.NotebookSessionLifecycleStateInactive {
			deactivateNotebookSession = true
		}
	}

	if e := CreateResource(d, sync); e != nil {
		return e
	}
	if deactivateNotebookSession {
		if e := sync.DeactivateNotebookSession(); e != nil {
			return e
		}
		sync.D.Set("state", oci_datascience.NotebookSessionLifecycleStateInactive)
	}
	return nil
}

func readDatascienceNotebookSession(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceNotebookSessionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	return ReadResource(sync)
}

func updateDatascienceNotebookSession(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceNotebookSessionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	// Activate/Deactivate NotebookSession
	activate, deactivate := false, false

	if sync.D.HasChange("state") {
		desiredState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_datascience.NotebookSessionLifecycleStateActive == oci_datascience.NotebookSessionLifecycleStateEnum(desiredState) {
			activate = true
		} else if oci_datascience.NotebookSessionLifecycleStateInactive == oci_datascience.NotebookSessionLifecycleStateEnum(desiredState) {
			deactivate = true
		}
	} else {
		currentState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_datascience.NotebookSessionLifecycleStateActive == oci_datascience.NotebookSessionLifecycleStateEnum(currentState) {
			activate = true
			deactivate = true
		}
	}

	if deactivate {
		if err := sync.DeactivateNotebookSession(); err != nil {
			return err
		}
		sync.D.Set("state", oci_datascience.NotebookSessionLifecycleStateInactive)
	}
	if err := UpdateResource(d, sync); err != nil {
		return err
	}

	if activate {
		if err := sync.ActivateNotebookSession(); err != nil {
			return err
		}
		sync.D.Set("state", oci_datascience.NotebookSessionLifecycleStateActive)
	}
	return nil
}

func deleteDatascienceNotebookSession(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceNotebookSessionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatascienceNotebookSessionResourceCrud struct {
	BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.NotebookSession
	DisableNotFoundRetries bool
}

func (s *DatascienceNotebookSessionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatascienceNotebookSessionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datascience.NotebookSessionLifecycleStateCreating),
	}
}

func (s *DatascienceNotebookSessionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.NotebookSessionLifecycleStateActive),
	}
}

func (s *DatascienceNotebookSessionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datascience.NotebookSessionLifecycleStateDeleting),
	}
}

func (s *DatascienceNotebookSessionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.NotebookSessionLifecycleStateDeleted),
	}
}

func (s *DatascienceNotebookSessionResourceCrud) Create() error {
	request := oci_datascience.CreateNotebookSessionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if notebookSessionConfigurationDetails, ok := s.D.GetOkExists("notebook_session_configuration_details"); ok {
		if tmpList := notebookSessionConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notebook_session_configuration_details", 0)
			tmp, err := s.mapToNotebookSessionConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NotebookSessionConfigurationDetails = &tmp
		}
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateNotebookSession(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NotebookSession
	return nil
}

func (s *DatascienceNotebookSessionResourceCrud) Get() error {
	request := oci_datascience.GetNotebookSessionRequest{}

	tmp := s.D.Id()
	request.NotebookSessionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetNotebookSession(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NotebookSession
	return nil
}

func (s *DatascienceNotebookSessionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdateNotebookSessionRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if notebookSessionConfigurationDetails, ok := s.D.GetOkExists("notebook_session_configuration_details"); ok {
		if tmpList := notebookSessionConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notebook_session_configuration_details", 0)
			tmp, err := s.mapToNotebookSessionConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NotebookSessionConfigurationDetails = &tmp
		}
	}

	tmp := s.D.Id()
	request.NotebookSessionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateNotebookSession(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NotebookSession
	return nil
}

func (s *DatascienceNotebookSessionResourceCrud) Delete() error {
	request := oci_datascience.DeleteNotebookSessionRequest{}

	tmp := s.D.Id()
	request.NotebookSessionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeleteNotebookSession(context.Background(), request)
	return err
}

func (s *DatascienceNotebookSessionResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NotebookSessionConfigurationDetails != nil {
		s.D.Set("notebook_session_configuration_details", []interface{}{NotebookSessionConfigurationDetailsToMap(s.Res.NotebookSessionConfigurationDetails)})
	} else {
		s.D.Set("notebook_session_configuration_details", nil)
	}

	if s.Res.NotebookSessionUrl != nil {
		s.D.Set("notebook_session_url", *s.Res.NotebookSessionUrl)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *DatascienceNotebookSessionResourceCrud) mapToNotebookSessionConfigurationDetails(fieldKeyFormat string) (oci_datascience.NotebookSessionConfigurationDetails, error) {
	result := oci_datascience.NotebookSessionConfigurationDetails{}

	if blockStorageSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_storage_size_in_gbs")); ok {
		tmp := blockStorageSizeInGBs.(int)
		result.BlockStorageSizeInGBs = &tmp
	}

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		result.Shape = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func NotebookSessionConfigurationDetailsToMap(obj *oci_datascience.NotebookSessionConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BlockStorageSizeInGBs != nil {
		result["block_storage_size_in_gbs"] = int(*obj.BlockStorageSizeInGBs)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *DatascienceNotebookSessionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangeNotebookSessionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.NotebookSessionId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ChangeNotebookSessionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *DatascienceNotebookSessionResourceCrud) ActivateNotebookSession() error {
	request := oci_datascience.ActivateNotebookSessionRequest{}

	tmp := s.D.Id()
	request.NotebookSessionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ActivateNotebookSession(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_datascience.NotebookSessionLifecycleStateActive }
	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceNotebookSessionResourceCrud) DeactivateNotebookSession() error {
	request := oci_datascience.DeactivateNotebookSessionRequest{}

	tmp := s.D.Id()
	request.NotebookSessionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeactivateNotebookSession(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_datascience.NotebookSessionLifecycleStateInactive }
	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}
