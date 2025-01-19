// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bastion

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_bastion "github.com/oracle/oci-go-sdk/v65/bastion"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func BastionBastionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBastionBastion,
		Read:     readBastionBastion,
		Update:   updateBastionBastion,
		Delete:   deleteBastionBastion,
		Schema: map[string]*schema.Schema{
			// Required
			"bastion_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_cidr_block_allow_list": {
				Type:     schema.TypeList,
				Required: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"dns_proxy_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"max_session_ttl_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:             schema.TypeString,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
			},
			"phone_book_entry": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"static_jump_host_ip_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_sessions_allowed": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"private_endpoint_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"target_vcn_id": {
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

func createBastionBastion(d *schema.ResourceData, m interface{}) error {
	sync := &BastionBastionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BastionClient()

	return tfresource.CreateResource(d, sync)
}

func readBastionBastion(d *schema.ResourceData, m interface{}) error {
	sync := &BastionBastionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BastionClient()

	return tfresource.ReadResource(sync)
}

func updateBastionBastion(d *schema.ResourceData, m interface{}) error {
	sync := &BastionBastionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BastionClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteBastionBastion(d *schema.ResourceData, m interface{}) error {
	sync := &BastionBastionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BastionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BastionBastionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bastion.BastionClient
	Res                    *oci_bastion.Bastion
	DisableNotFoundRetries bool
}

func (s *BastionBastionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BastionBastionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bastion.BastionLifecycleStateCreating),
	}
}

func (s *BastionBastionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bastion.BastionLifecycleStateActive),
	}
}

func (s *BastionBastionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bastion.BastionLifecycleStateDeleting),
	}
}

func (s *BastionBastionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bastion.BastionLifecycleStateDeleted),
	}
}

func (s *BastionBastionResourceCrud) Create() error {
	request := oci_bastion.CreateBastionRequest{}

	if bastionType, ok := s.D.GetOkExists("bastion_type"); ok {
		tmp := bastionType.(string)
		request.BastionType = &tmp
	}

	if clientCidrBlockAllowList, ok := s.D.GetOkExists("client_cidr_block_allow_list"); ok {
		interfaces := clientCidrBlockAllowList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("client_cidr_block_allow_list") {
			request.ClientCidrBlockAllowList = tmp
		}
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

	if dnsProxyStatus, ok := s.D.GetOkExists("dns_proxy_status"); ok {
		request.DnsProxyStatus = oci_bastion.BastionDnsProxyStatusEnum(dnsProxyStatus.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if maxSessionTtlInSeconds, ok := s.D.GetOkExists("max_session_ttl_in_seconds"); ok {
		tmp := maxSessionTtlInSeconds.(int)
		request.MaxSessionTtlInSeconds = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if phoneBookEntry, ok := s.D.GetOkExists("phone_book_entry"); ok {
		tmp := phoneBookEntry.(string)
		request.PhoneBookEntry = &tmp
	}

	if staticJumpHostIpAddresses, ok := s.D.GetOkExists("static_jump_host_ip_addresses"); ok {
		interfaces := staticJumpHostIpAddresses.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("static_jump_host_ip_addresses") {
			request.StaticJumpHostIpAddresses = tmp
		}
	}

	if targetSubnetId, ok := s.D.GetOkExists("target_subnet_id"); ok {
		tmp := targetSubnetId.(string)
		request.TargetSubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bastion")

	response, err := s.Client.CreateBastion(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getBastionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bastion"), oci_bastion.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BastionBastionResourceCrud) getBastionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bastion.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	bastionId, err := bastionWaitForWorkRequest(workId, "bastion",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*bastionId)

	return s.Get()
}

func bastionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "bastion", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_bastion.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func bastionWaitForWorkRequest(wId *string, entityType string, action oci_bastion.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bastion.BastionClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bastion")
	retryPolicy.ShouldRetryOperation = bastionWorkRequestShouldRetryFunc(timeout)

	response := oci_bastion.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_bastion.OperationStatusInProgress),
			string(oci_bastion.OperationStatusAccepted),
			string(oci_bastion.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_bastion.OperationStatusSucceeded),
			string(oci_bastion.OperationStatusFailed),
			string(oci_bastion.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_bastion.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_bastion.OperationStatusFailed || response.Status == oci_bastion.OperationStatusCanceled {
		return nil, getErrorFromBastionBastionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBastionBastionWorkRequest(client *oci_bastion.BastionClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bastion.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_bastion.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
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

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *BastionBastionResourceCrud) Get() error {
	request := oci_bastion.GetBastionRequest{}

	tmp := s.D.Id()
	request.BastionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bastion")

	response, err := s.Client.GetBastion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Bastion
	return nil
}

func (s *BastionBastionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_bastion.UpdateBastionRequest{}

	tmp := s.D.Id()
	request.BastionId = &tmp

	if clientCidrBlockAllowList, ok := s.D.GetOkExists("client_cidr_block_allow_list"); ok {
		interfaces := clientCidrBlockAllowList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("client_cidr_block_allow_list") {
			request.ClientCidrBlockAllowList = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if maxSessionTtlInSeconds, ok := s.D.GetOkExists("max_session_ttl_in_seconds"); ok {
		tmp := maxSessionTtlInSeconds.(int)
		request.MaxSessionTtlInSeconds = &tmp
	}

	if staticJumpHostIpAddresses, ok := s.D.GetOkExists("static_jump_host_ip_addresses"); ok {
		interfaces := staticJumpHostIpAddresses.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("static_jump_host_ip_addresses") {
			request.StaticJumpHostIpAddresses = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bastion")

	response, err := s.Client.UpdateBastion(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBastionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bastion"), oci_bastion.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BastionBastionResourceCrud) Delete() error {
	request := oci_bastion.DeleteBastionRequest{}

	tmp := s.D.Id()
	request.BastionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bastion")

	response, err := s.Client.DeleteBastion(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := bastionWaitForWorkRequest(workId, "bastion", oci_bastion.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BastionBastionResourceCrud) SetData() error {
	if s.Res.BastionType != nil {
		s.D.Set("bastion_type", *s.Res.BastionType)
	}

	s.D.Set("client_cidr_block_allow_list", s.Res.ClientCidrBlockAllowList)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("dns_proxy_status", s.Res.DnsProxyStatus)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaxSessionTtlInSeconds != nil {
		s.D.Set("max_session_ttl_in_seconds", *s.Res.MaxSessionTtlInSeconds)
	}

	if s.Res.MaxSessionsAllowed != nil {
		s.D.Set("max_sessions_allowed", *s.Res.MaxSessionsAllowed)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PhoneBookEntry != nil {
		s.D.Set("phone_book_entry", *s.Res.PhoneBookEntry)
	}

	if s.Res.PrivateEndpointIpAddress != nil {
		s.D.Set("private_endpoint_ip_address", *s.Res.PrivateEndpointIpAddress)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("static_jump_host_ip_addresses", s.Res.StaticJumpHostIpAddresses)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetSubnetId != nil {
		s.D.Set("target_subnet_id", *s.Res.TargetSubnetId)
	}

	if s.Res.TargetVcnId != nil {
		s.D.Set("target_vcn_id", *s.Res.TargetVcnId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *BastionBastionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_bastion.ChangeBastionCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BastionId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bastion")

	_, err := s.Client.ChangeBastionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
