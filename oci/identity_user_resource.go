// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func UserResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createUser,
		Read:     readUser,
		Update:   updateUser,
		Delete:   deleteUser,
		Schema: map[string]*schema.Schema{
			// The legacy provider exposed this as read-only/computed. The API requires this param. For legacy users who are
			// not supplying a value, make it optional, behind the scenes it will use the tenancy ocid if not supplied.
			// If a user supplies the value, then changes it, it requires forcing new.
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			// Required
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"capabilities": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"can_use_api_keys": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"can_use_auth_tokens": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"can_use_console_password": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"can_use_customer_secret_keys": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"can_use_smtp_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"external_identifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"identity_provider_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_state": {
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
			// @Deprecated: time_modified (removed)
			"time_modified": {
				Type:       schema.TypeString,
				Deprecated: FieldDeprecated("time_modified"),
				Computed:   true,
			},
		},
	}
}

func createUser(d *schema.ResourceData, m interface{}) error {
	sync := &UserResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.Configuration = m.(*OracleClients).configuration

	return CreateResource(d, sync)
}

func readUser(d *schema.ResourceData, m interface{}) error {
	sync := &UserResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

func updateUser(d *schema.ResourceData, m interface{}) error {
	sync := &UserResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return UpdateResource(d, sync)
}

func deleteUser(d *schema.ResourceData, m interface{}) error {
	sync := &UserResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type UserResourceCrud struct {
	BaseCrud
	Client                 *oci_identity.IdentityClient
	Configuration          map[string]string
	Res                    *oci_identity.User
	DisableNotFoundRetries bool
}

func (s *UserResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *UserResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.UserLifecycleStateCreating),
	}
}

func (s *UserResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.UserLifecycleStateActive),
	}
}

func (s *UserResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.UserLifecycleStateDeleting),
	}
}

func (s *UserResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.UserLifecycleStateDeleted),
	}
}

func (s *UserResourceCrud) Create() error {
	request := oci_identity.CreateUserRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	} else { // @next-break: remove
		// Prevent potentially inferring wrong TenancyOCID from InstancePrincipal
		if auth := s.Configuration["auth"]; strings.ToLower(auth) == strings.ToLower(authInstancePrincipalSetting) {
			return fmt.Errorf("compartment_id must be specified for this resource")
		}
		// Maintain legacy contract of compartment_id defaulting to tenancy ocid if not specified
		c := *s.Client.ConfigurationProvider()
		if c == nil {
			return fmt.Errorf("cannot access tenancyOCID")
		}
		tenancy, err := c.TenancyOCID()
		if err != nil {
			return err
		}
		request.CompartmentId = &tenancy
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *UserResourceCrud) Get() error {
	request := oci_identity.GetUserRequest{}

	tmp := s.D.Id()
	request.UserId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *UserResourceCrud) Update() error {
	request := oci_identity.UpdateUserRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.UserId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *UserResourceCrud) Delete() error {
	request := oci_identity.DeleteUserRequest{}

	tmp := s.D.Id()
	request.UserId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteUser(context.Background(), request)
	return err
}

func (s *UserResourceCrud) SetData() error {
	if s.Res.Capabilities != nil {
		s.D.Set("capabilities", []interface{}{UserCapabilitiesToMap(s.Res.Capabilities)})
	} else {
		s.D.Set("capabilities", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.ExternalIdentifier != nil {
		s.D.Set("external_identifier", *s.Res.ExternalIdentifier)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IdentityProviderId != nil {
		s.D.Set("identity_provider_id", *s.Res.IdentityProviderId)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func UserCapabilitiesToMap(obj *oci_identity.UserCapabilities) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CanUseApiKeys != nil {
		result["can_use_api_keys"] = bool(*obj.CanUseApiKeys)
	}

	if obj.CanUseAuthTokens != nil {
		result["can_use_auth_tokens"] = bool(*obj.CanUseAuthTokens)
	}

	if obj.CanUseConsolePassword != nil {
		result["can_use_console_password"] = bool(*obj.CanUseConsolePassword)
	}

	if obj.CanUseCustomerSecretKeys != nil {
		result["can_use_customer_secret_keys"] = bool(*obj.CanUseCustomerSecretKeys)
	}

	if obj.CanUseSmtpCredentials != nil {
		result["can_use_smtp_credentials"] = bool(*obj.CanUseSmtpCredentials)
	}

	return result
}
