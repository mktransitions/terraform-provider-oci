// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomain,
		Read:     readIdentityDomain,
		Update:   updateIdentityDomain,
		Delete:   deleteIdentityDomain,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"home_region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"admin_email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"admin_first_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"admin_last_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"admin_user_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_hidden_on_login": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_notification_bypassed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_primary_email_required": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"home_region_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replica_regions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{ // Add a validator to restrict invalid user input
					string(oci_identity.DomainLifecycleStateActive),
					string(oci_identity.DomainLifecycleStateInactive),
				}, true),
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityDomain(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_identity.DomainLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_identity.DomainLifecycleStateInactive {
			powerOff = true
		}
	}

	if err := tfresource.CreateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.deActivate(); err != nil {
			return err
		}
		sync.D.Set("state", oci_identity.DomainLifecycleStateInactive)
	}
	return nil
}

func readIdentityDomain(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityDomain(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_identity.DomainLifecycleStateActive == oci_identity.DomainLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_identity.DomainLifecycleStateInactive == oci_identity.DomainLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.activate(); err != nil {
			return err
		}
		sync.D.Set("state", oci_identity.DomainLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.deActivate(); err != nil {
			return err
		}
		sync.D.Set("state", oci_identity.DomainLifecycleStateInactive)
	}

	return nil
}

func deleteIdentityDomain(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityDomainResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.Domain
	DisableNotFoundRetries bool
}

func (s *IdentityDomainResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.DomainLifecycleStateCreating),
	}
}

func (s *IdentityDomainResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.DomainLifecycleStateActive),
	}
}

func (s *IdentityDomainResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.DomainLifecycleStateDeleting),
	}
}

func (s *IdentityDomainResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *IdentityDomainResourceCrud) Create() error {
	request := oci_identity.CreateDomainRequest{}

	if adminEmail, ok := s.D.GetOkExists("admin_email"); ok {
		tmp := adminEmail.(string)
		request.AdminEmail = &tmp
	}

	if adminFirstName, ok := s.D.GetOkExists("admin_first_name"); ok {
		tmp := adminFirstName.(string)
		request.AdminFirstName = &tmp
	}

	if adminLastName, ok := s.D.GetOkExists("admin_last_name"); ok {
		tmp := adminLastName.(string)
		request.AdminLastName = &tmp
	}

	if adminUserName, ok := s.D.GetOkExists("admin_user_name"); ok {
		tmp := adminUserName.(string)
		request.AdminUserName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if homeRegion, ok := s.D.GetOkExists("home_region"); ok {
		tmp := homeRegion.(string)
		request.HomeRegion = &tmp
	}

	if isHiddenOnLogin, ok := s.D.GetOkExists("is_hidden_on_login"); ok {
		tmp := isHiddenOnLogin.(bool)
		request.IsHiddenOnLogin = &tmp
	}

	if isNotificationBypassed, ok := s.D.GetOkExists("is_notification_bypassed"); ok {
		tmp := isNotificationBypassed.(bool)
		request.IsNotificationBypassed = &tmp
	}

	if isPrimaryEmailRequired, ok := s.D.GetOkExists("is_primary_email_required"); ok {
		tmp := isPrimaryEmailRequired.(bool)
		request.IsPrimaryEmailRequired = &tmp
	}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		tmp := licenseType.(string)
		request.LicenseType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateDomain(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_identity.GetIamWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetIamWorkRequest(context.Background(),
		oci_identity.GetIamWorkRequestRequest{
			IamWorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "domain") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getDomainFromWorkRequest(workId, oci_identity.IamWorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IdentityDomainResourceCrud) getDomainFromWorkRequest(workId *string,
	actionTypeEnum oci_identity.IamWorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	domainId, err := domainWaitForWorkRequest(workId, "domain",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*domainId)

	return s.Get()
}

func domainWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "identity", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_identity.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func domainWaitForWorkRequest(wId *string, entityType string, action oci_identity.IamWorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_identity.IdentityClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "identity")
	retryPolicy.ShouldRetryOperation = domainWorkRequestShouldRetryFunc(timeout)

	response := oci_identity.GetIamWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_identity.IamWorkRequestStatusAccepted),
			string(oci_identity.IamWorkRequestStatusInProgress),
			string(oci_identity.IamWorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_identity.IamWorkRequestStatusSucceeded),
			string(oci_identity.IamWorkRequestStatusFailed),
			string(oci_identity.IamWorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetIamWorkRequest(context.Background(),
				oci_identity.GetIamWorkRequestRequest{
					IamWorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			return response, string(response.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_identity.IamWorkRequestStatusFailed || response.Status == oci_identity.IamWorkRequestStatusCanceled {
		return nil, getErrorFromIdentityDomainWorkRequest(client, wId, retryPolicy, response.OperationType)
	}

	return identifier, nil
}

func getErrorFromIdentityDomainWorkRequest(client *oci_identity.IdentityClient, workId *string, retryPolicy *oci_common.RetryPolicy, operationType oci_identity.IamWorkRequestOperationTypeEnum) error {

	errorMessage, err := getErrorMessageFromIdentityDomainWorkRequest(client, workId, retryPolicy)
	if err != nil {
		return err
	}
	workRequestErr := fmt.Errorf("oci_identity_domain: iam work request did not succeed, workId: %s, action: %s. ErrorMessage: %s", *workId, operationType, errorMessage)
	return workRequestErr
}

func getErrorMessageFromIdentityDomainWorkRequest(client *oci_identity.IdentityClient, workId *string, retryPolicy *oci_common.RetryPolicy) (string, error) {
	errorMessage := ""
	response, err := client.ListIamWorkRequestErrors(context.Background(),
		oci_identity.ListIamWorkRequestErrorsRequest{
			IamWorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})

	if err != nil {
		return errorMessage, err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage = strings.Join(allErrs, "\n")

	return errorMessage, nil
}

func (s *IdentityDomainResourceCrud) Get() error {
	request := oci_identity.GetDomainRequest{}

	tmp := s.D.Id()
	request.DomainId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetDomain(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Domain
	return nil
}

func (s *IdentityDomainResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok && s.D.HasChange("license_type") {
		oldRaw, newRaw := s.D.GetChange("license_type")
		if newRaw != "" && oldRaw != "" {
			err := s.updateLicenseType(licenseType)
			if err != nil {
				return err
			}
		}
	}

	request := oci_identity.UpdateDomainRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.DomainId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isHiddenOnLogin, ok := s.D.GetOkExists("is_hidden_on_login"); ok {
		tmp := isHiddenOnLogin.(bool)
		request.IsHiddenOnLogin = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateDomain(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDomainFromWorkRequest(workId, oci_identity.IamWorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *IdentityDomainResourceCrud) Delete() error {
	tmp := s.D.Id()

	deactivateRequest := oci_identity.DeactivateDomainRequest{}

	deactivateRequest.DomainId = &tmp

	deactivateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	deactivateResponse, err := s.Client.DeactivateDomain(context.Background(), deactivateRequest)
	if err != nil {
		return err
	}

	deactivateWorkId := deactivateResponse.OpcWorkRequestId
	// Wait until it finishes
	_, deactivateWorkRequestErr := domainWaitForWorkRequest(deactivateWorkId, "domain",
		oci_identity.IamWorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)

	if deactivateWorkRequestErr != nil {
		return deactivateWorkRequestErr
	}

	deleteRequest := oci_identity.DeleteDomainRequest{}

	deleteRequest.DomainId = &tmp

	deleteRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	deleteResponse, err := s.Client.DeleteDomain(context.Background(), deleteRequest)
	if err != nil {
		return err
	}

	deleteWorkId := deleteResponse.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := domainWaitForWorkRequest(deleteWorkId, "domain",
		oci_identity.IamWorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *IdentityDomainResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HomeRegion != nil {
		s.D.Set("home_region", *s.Res.HomeRegion)
	}

	if s.Res.HomeRegionUrl != nil {
		s.D.Set("home_region_url", *s.Res.HomeRegionUrl)
	}

	if s.Res.IsHiddenOnLogin != nil {
		s.D.Set("is_hidden_on_login", *s.Res.IsHiddenOnLogin)
	}

	if s.Res.LicenseType != nil {
		s.D.Set("license_type", *s.Res.LicenseType)
	}

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	replicaRegions := []interface{}{}
	for _, item := range s.Res.ReplicaRegions {
		replicaRegions = append(replicaRegions, ReplicatedRegionDetailsToMap(item))
	}
	s.D.Set("replica_regions", replicaRegions)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.Url != nil {
		s.D.Set("url", *s.Res.Url)
	}

	return nil
}

func ReplicatedRegionDetailsToMap(obj oci_identity.ReplicatedRegionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["state"] = string(obj.State)

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func (s *IdentityDomainResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_identity.ChangeDomainCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DomainId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ChangeDomainCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDomainFromWorkRequest(workId, oci_identity.IamWorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *IdentityDomainResourceCrud) updateLicenseType(licenseType interface{}) error {
	changeDomainLicenseTypeRequest := oci_identity.ChangeDomainLicenseTypeRequest{}

	licenseTypeTmp := licenseType.(string)
	changeDomainLicenseTypeRequest.LicenseType = &licenseTypeTmp

	idTmp := s.D.Id()
	changeDomainLicenseTypeRequest.DomainId = &idTmp

	changeDomainLicenseTypeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ChangeDomainLicenseType(context.Background(), changeDomainLicenseTypeRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDomainFromWorkRequest(workId, oci_identity.IamWorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *IdentityDomainResourceCrud) activate() error {
	activateDomainRequest := oci_identity.ActivateDomainRequest{}

	idTmp := s.D.Id()
	activateDomainRequest.DomainId = &idTmp

	activateDomainRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ActivateDomain(context.Background(), activateDomainRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDomainFromWorkRequest(workId, oci_identity.IamWorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *IdentityDomainResourceCrud) deActivate() error {
	deactivateDomainRequest := oci_identity.DeactivateDomainRequest{}

	idTmp := s.D.Id()
	deactivateDomainRequest.DomainId = &idTmp

	deactivateDomainRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.DeactivateDomain(context.Background(), deactivateDomainRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDomainFromWorkRequest(workId, oci_identity.IamWorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
