// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	"regexp"

	"strings"

	"github.com/hashicorp/terraform/helper/validation"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

func KeyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createKey,
		Read:     readKey,
		Update:   updateKey,
		Delete:   deleteKey,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key_shape": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"algorithm": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"length": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// Optional
			"desired_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ENABLED",
					"DISABLED",
				}, false),
			},

			// Computed
			"current_key_version": {
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
			"vault_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createKey(d *schema.ResourceData, m interface{}) error {
	sync := &KeyResourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		return fmt.Errorf("management endpoint missing")
	}
	client, err := m.(*OracleClients).KmsManagementClient(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return CreateResource(d, sync)
}

func readKey(d *schema.ResourceData, m interface{}) error {
	sync := &KeyResourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		//Import use case:
		id := d.Id()
		regex, _ := regexp.Compile("^managementEndpoint/(.*)/keys/(.*)$")
		tokens := regex.FindStringSubmatch(id)
		if len(tokens) == 3 {
			endpoint = tokens[1]
			d.Set("management_endpoint", endpoint)
			d.SetId(tokens[2])
		} else {
			return fmt.Errorf("id %s should be format: managementEndpoint/{managementEndpoint}/keys/{keyId}", id)
		}
	}

	client, err := m.(*OracleClients).KmsManagementClient(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return ReadResource(sync)
}

func updateKey(d *schema.ResourceData, m interface{}) error {
	sync := &KeyResourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		return fmt.Errorf("management endpoint missing")
	}
	client, err := m.(*OracleClients).KmsManagementClient(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return UpdateResource(d, sync)
}

func deleteKey(d *schema.ResourceData, m interface{}) error {
	return nil
}

type KeyResourceCrud struct {
	BaseCrud
	Client                 *oci_kms.KmsManagementClient
	Res                    *oci_kms.Key
	DisableNotFoundRetries bool
}

func (s *KeyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *KeyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_kms.KeyLifecycleStateCreating),
		string(oci_kms.KeyLifecycleStateEnabling),
	}
}

func (s *KeyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_kms.KeyLifecycleStateEnabled),
	}
}

func (s *KeyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_kms.KeyLifecycleStateDisabled),
		string(oci_kms.KeyLifecycleStateDeleting),
		string(oci_kms.KeyLifecycleStateSchedulingDeletion),
	}
}

func (s *KeyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_kms.KeyLifecycleStateDeleted),
		string(oci_kms.KeyLifecycleStatePendingDeletion),
	}
}

func (s *KeyResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_kms.KeyLifecycleStateEnabling),
		string(oci_kms.KeyLifecycleStateDisabling),
	}
}

func (s *KeyResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_kms.KeyLifecycleStateEnabled),
		string(oci_kms.KeyLifecycleStateDisabled),
	}
}

func (s *KeyResourceCrud) Create() error {
	if desiredState, ok := s.D.GetOkExists("desired_state"); ok && !strings.EqualFold(desiredState.(string), "ENABLED") {
		return fmt.Errorf("oci_kms_keys can only be created in ENABLED state")
	}

	request := oci_kms.CreateKeyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if keyShape, ok := s.D.GetOkExists("key_shape"); ok {
		if tmpList := keyShape.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "key_shape", 0)
			tmp, err := s.mapToKeyShape(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.KeyShape = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.CreateKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Key
	return nil
}

func (s *KeyResourceCrud) Get() error {
	request := oci_kms.GetKeyRequest{}

	tmp := s.D.Id()
	request.KeyId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.GetKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Key
	return nil
}

func (s *KeyResourceCrud) Update() error {
	request := oci_kms.UpdateKeyRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.KeyId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.UpdateKey(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response.Key

	// Handle activation/deactivation here
	if desiredState, ok := s.D.GetOkExists("desired_state"); ok && !strings.EqualFold(desiredState.(string), s.D.Get("state").(string)) {
		desiredStateString := desiredState.(string)

		if desiredStateString == "ENABLED" {
			activationRequest := oci_kms.EnableKeyRequest{}
			tmpId := s.D.Id()
			activationRequest.KeyId = &tmpId

			activationRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")
			activationResponse, err := s.Client.EnableKey(context.Background(), activationRequest)
			if err != nil {
				return err
			}
			s.Res = &activationResponse.Key
		} else if desiredStateString == "DISABLED" {
			deactivationRequest := oci_kms.DisableKeyRequest{}
			tmpId := s.D.Id()
			deactivationRequest.KeyId = &tmpId

			deactivationRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")
			deactivationResponse, err := s.Client.DisableKey(context.Background(), deactivationRequest)
			if err != nil {
				return err
			}
			s.Res = &deactivationResponse.Key
		}
	}

	return nil
}

func (s *KeyResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CurrentKeyVersion != nil {
		s.D.Set("current_key_version", *s.Res.CurrentKeyVersion)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("desired_state", s.Res.LifecycleState)

	if s.Res.KeyShape != nil {
		s.D.Set("key_shape", []interface{}{KeyShapeToMap(s.Res.KeyShape)})
	} else {
		s.D.Set("key_shape", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	return nil
}

func (s *KeyResourceCrud) mapToKeyShape(fieldKeyFormat string) (oci_kms.KeyShape, error) {
	result := oci_kms.KeyShape{}

	if algorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "algorithm")); ok {
		tmp := oci_kms.KeyShapeAlgorithmEnum(algorithm.(string))
		result.Algorithm = tmp
	}

	if length, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "length")); ok {
		tmp := length.(int)
		result.Length = &tmp
	}

	return result, nil
}

func KeyShapeToMap(obj *oci_kms.KeyShape) map[string]interface{} {
	result := map[string]interface{}{}

	result["algorithm"] = string(obj.Algorithm)

	if obj.Length != nil {
		result["length"] = int(*obj.Length)
	}

	return result
}
