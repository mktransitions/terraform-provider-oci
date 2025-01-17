// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_certificates_management "github.com/oracle/oci-go-sdk/v65/certificatesmanagement"
)

func CertificatesManagementCertificateResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCertificatesManagementCertificate,
		Read:     readCertificatesManagementCertificate,
		Update:   updateCertificatesManagementCertificate,
		Delete:   deleteCertificatesManagementCertificate,
		Schema: map[string]*schema.Schema{
			// Required
			"certificate_config": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"config_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"IMPORTED",
								"ISSUED_BY_INTERNAL_CA",
								"MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA",
							}, true),
						},
						"cert_chain_pem": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},
						"certificate_pem": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},
						"private_key_pem": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},
						"private_key_pem_passphrase": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},
						"certificate_profile_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
							ForceNew: true,
						},
						"csr_pem": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},
						"issuer_certificate_authority_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
							ForceNew: true,
						},
						"key_algorithm": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"signature_algorithm": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"subject": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: false,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"common_name": {
										Type:     schema.TypeString,
										Required: true,
										Computed: false,
										ForceNew: true,
									},
									// Optional
									"country": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"distinguished_name_qualifier": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"domain_component": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"generation_qualifier": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"given_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"initials": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"locality_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"organization": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"organizational_unit": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"pseudonym": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"serial_number": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"state_or_province_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"street": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"surname": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"title": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},
									"user_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: false,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"subject_alternative_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: false,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"type": {
										Type:     schema.TypeString,
										Required: true,
										Computed: false,
										ForceNew: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
										Computed: false,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"validity": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: false,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"time_of_validity_not_after": {
										Type:             schema.TypeString,
										Required:         true,
										Computed:         false,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									// Optional
									"time_of_validity_not_before": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         false,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},

									// Computed
								},
							},
						},
						"version_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
						},

						// Computed
					},
				},
			},
			// end certificate config

			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"certificate_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"advance_renewal_period": {
							Type:     schema.TypeString,
							Required: true,
						},
						"renewal_interval": {
							Type:     schema.TypeString,
							Required: true,
						},
						"rule_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CERTIFICATE_RENEWAL_RULE",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"certificate_profile_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_revocation_list_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"custom_formatted_urls": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"object_storage_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"object_storage_bucket_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_storage_namespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_storage_object_name_format": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"config_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_version": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"certificate_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"issuer_ca_version_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"revocation_status": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"revocation_reason": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_of_revocation": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"stages": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"subject_alternative_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_deletion": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"validity": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"time_of_validity_not_after": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_of_validity_not_before": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"version_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"issuer_certificate_authority_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"signature_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subject": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"common_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"country": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"distinguished_name_qualifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain_component": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"generation_qualifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"given_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"initials": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"locality_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organization": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organizational_unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pseudonym": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state_or_province_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"street": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"surname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_deletion": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCertificatesManagementCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readCertificatesManagementCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

func updateCertificatesManagementCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCertificatesManagementCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.DeleteResource(d, sync)
}

type CertificatesManagementCertificateResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_certificates_management.CertificatesManagementClient
	Res                    *oci_certificates_management.Certificate
	DisableNotFoundRetries bool
}

func (s *CertificatesManagementCertificateResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CertificatesManagementCertificateResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_certificates_management.CertificateLifecycleStateCreating),
	}
}

func (s *CertificatesManagementCertificateResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_certificates_management.CertificateLifecycleStateActive),
	}
}

func (s *CertificatesManagementCertificateResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_certificates_management.CertificateLifecycleStateUpdating),
	}
}

func (s *CertificatesManagementCertificateResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_certificates_management.CertificateLifecycleStateActive),
	}
}

func (s *CertificatesManagementCertificateResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_certificates_management.CertificateLifecycleStateDeleting),
		string(oci_certificates_management.CertificateLifecycleStateSchedulingDeletion),
	}
}

func (s *CertificatesManagementCertificateResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_certificates_management.CertificateLifecycleStateDeleted),
		string(oci_certificates_management.CertificateLifecycleStatePendingDeletion),
	}
}

func (s *CertificatesManagementCertificateResourceCrud) Create() error {
	request := oci_certificates_management.CreateCertificateRequest{}

	if certificateConfig, ok := s.D.GetOkExists("certificate_config"); ok {
		if tmpList := certificateConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_config", 0)
			tmp, err := s.mapToCreateCertificateConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CertificateConfig = tmp
		}
	}

	if certificateRules, ok := s.D.GetOkExists("certificate_rules"); ok {
		interfaces := certificateRules.([]interface{})
		tmp := make([]oci_certificates_management.CertificateRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_rules", stateDataIndex)
			converted, err := s.mapToCertificateRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("certificate_rules") {
			request.CertificateRules = tmp
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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	response, err := s.Client.CreateCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Certificate
	return nil
}

func (s *CertificatesManagementCertificateResourceCrud) Get() error {
	request := oci_certificates_management.GetCertificateRequest{}

	tmp := s.D.Id()
	request.CertificateId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	response, err := s.Client.GetCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Certificate
	return nil
}

func (s *CertificatesManagementCertificateResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_certificates_management.UpdateCertificateRequest{}

	if certificateConfig, ok := s.D.GetOkExists("certificate_config"); ok && s.D.HasChange("certificate_config") {
		if tmpList := certificateConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_config", 0)
			tmp, err := s.mapToUpdateCertificateConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CertificateConfig = tmp
		}
	}

	//if currentVersionNumber, ok := s.D.GetOkExists("current_version_number"); ok {
	//	tmp := currentVersionNumber.(string)
	//	tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
	//	if err != nil {
	//		return fmt.Errorf("unable to convert currentVersionNumber string: %s to an int64 and encountered error: %v", tmp, err)
	//	}
	//	request.CurrentVersionNumber = &tmpInt64
	//}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	// only update if request has updates
	if request.CertificateConfig != nil || request.FreeformTags != nil || request.Description != nil || request.DefinedTags != nil {
		tmp := s.D.Id()
		request.CertificateId = &tmp

		response, err := s.Client.UpdateCertificate(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res = &response.Certificate
	}

	// update rules
	if _, ok := s.D.GetOkExists("certificate_rules"); ok && s.D.HasChange("certificate_rules") {
		err := s.UpdateRules()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *CertificatesManagementCertificateResourceCrud) UpdateRules() error {
	request := oci_certificates_management.UpdateCertificateRequest{}

	tmp := s.D.Id()
	request.CertificateId = &tmp

	if certificateRules, ok := s.D.GetOkExists("certificate_rules"); ok && s.D.HasChange("certificate_rules") {
		interfaces := certificateRules.([]interface{})
		tmp := make([]oci_certificates_management.CertificateRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_rules", stateDataIndex)
			converted, err := s.mapToCertificateRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("certificate_rules") {
			request.CertificateRules = tmp
		} else {
			return nil
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	response, err := s.Client.UpdateCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Certificate
	return s.Get()
}

func (s *CertificatesManagementCertificateResourceCrud) Delete() error {
	request := oci_certificates_management.ScheduleCertificateDeletionRequest{}

	if timeOfDeletion, ok := s.D.GetOkExists("time_of_deletion"); ok {
		tmpTime, err := time.Parse(time.RFC3339Nano, timeOfDeletion.(string))
		if err != nil {
			return err
		}
		request.TimeOfDeletion = &oci_common.SDKTime{Time: tmpTime}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")
	tmp := s.D.Id()
	request.CertificateId = &tmp

	_, err := s.Client.ScheduleCertificateDeletion(context.Background(), request)
	return err
}

func (s *CertificatesManagementCertificateResourceCrud) SetData() error {
	s.D.Set("certificate_profile_type", s.Res.CertificateProfileType)

	if s.Res.CertificateRevocationListDetails != nil {
		s.D.Set("certificate_revocation_list_details", []interface{}{CertificateRevocationListDetailsToMap(s.Res.CertificateRevocationListDetails)})
	} else {
		s.D.Set("certificate_revocation_list_details", nil)
	}

	certificateRules := []interface{}{}
	for _, item := range s.Res.CertificateRules {
		certificateRules = append(certificateRules, CertificateRuleToMap(item))
	}
	s.D.Set("certificate_rules", certificateRules)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("config_type", s.Res.ConfigType)

	if s.Res.CurrentVersion != nil {
		s.D.Set("current_version", []interface{}{CertificateVersionSummaryToMap(s.Res.CurrentVersion)})
	} else {
		s.D.Set("current_version", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IssuerCertificateAuthorityId != nil {
		s.D.Set("issuer_certificate_authority_id", *s.Res.IssuerCertificateAuthorityId)
	}

	s.D.Set("key_algorithm", s.Res.KeyAlgorithm)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("signature_algorithm", s.Res.SignatureAlgorithm)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Subject != nil {
		s.D.Set("subject", []interface{}{CertificateSubjectToMap(s.Res.Subject)})
	} else {
		s.D.Set("subject", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	return nil
}

func CertificateRuleToMap(obj oci_certificates_management.CertificateRule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_certificates_management.CertificateRenewalRule:
		result["rule_type"] = "CERTIFICATE_RENEWAL_RULE"

		if v.AdvanceRenewalPeriod != nil {
			result["advance_renewal_period"] = string(*v.AdvanceRenewalPeriod)
		}

		if v.RenewalInterval != nil {
			result["renewal_interval"] = string(*v.RenewalInterval)
		}
	default:
		log.Printf("[WARN] Received 'rule_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func CertificateSubjectAlternativeNameToMap(obj oci_certificates_management.CertificateSubjectAlternativeName) map[string]interface{} {
	result := map[string]interface{}{}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func CertificatesManagementCertificateSummaryToMap(obj oci_certificates_management.CertificateSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["certificate_profile_type"] = obj.CertificateProfileType

	if obj.CertificateRules != nil {
		certificateRules := []interface{}{}
		for _, item := range obj.CertificateRules {
			certificateRules = append(certificateRules, CertificateRuleToMap(item))
		}
		result["certificate_rules"] = certificateRules
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = *obj.CompartmentId
	}

	result["config_type"] = obj.ConfigType

	if obj.CurrentVersionSummary != nil {
		result["current_version"] = []interface{}{CertificateVersionSummaryToMap(obj.CurrentVersionSummary)}
	} else {
		result["current_version"] = nil
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = *obj.Description
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = *obj.Id
	}

	if obj.IssuerCertificateAuthorityId != nil {
		result["issuer_certificate_authority_id"] = *obj.IssuerCertificateAuthorityId
	}

	result["key_algorithm"] = obj.KeyAlgorithm

	if obj.Name != nil {
		result["name"] = *obj.Name
	}

	result["signature_algorithm"] = obj.SignatureAlgorithm

	result["state"] = obj.LifecycleState

	if obj.Subject != nil {
		result["subject"] = []interface{}{CertificateSubjectToMap(obj.Subject)}
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeOfDeletion != nil {
		result["time_of_deletion"] = obj.TimeOfDeletion.String()
	}

	return result
}

func CertificateVersionSummaryToMap(obj *oci_certificates_management.CertificateVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CertificateId != nil {
		result["certificate_id"] = string(*obj.CertificateId)
	}

	if obj.IssuerCaVersionNumber != nil {
		result["issuer_ca_version_number"] = strconv.FormatInt(*obj.IssuerCaVersionNumber, 10)
	}

	if obj.RevocationStatus != nil {
		result["revocation_status"] = []interface{}{RevocationStatusToMap(obj.RevocationStatus)}
	}

	if obj.SerialNumber != nil {
		result["serial_number"] = string(*obj.SerialNumber)
	}

	result["stages"] = obj.Stages

	if obj.SubjectAlternativeNames != nil {
		subjectAlternativeNames := []interface{}{}
		for _, item := range obj.SubjectAlternativeNames {
			subjectAlternativeNames = append(subjectAlternativeNames, CertificateSubjectAlternativeNameToMap(item))
		}
		result["subject_alternative_names"] = subjectAlternativeNames
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeOfDeletion != nil {
		result["time_of_deletion"] = obj.TimeOfDeletion.String()
	}

	if obj.Validity != nil {
		result["validity"] = []interface{}{ValidityToMap(obj.Validity)}
	}

	if obj.VersionName != nil {
		result["version_name"] = string(*obj.VersionName)
	}

	if obj.VersionNumber != nil {
		result["version_number"] = strconv.FormatInt(*obj.VersionNumber, 10)
	}

	return result
}

func (s *CertificatesManagementCertificateResourceCrud) mapToCreateCertificateConfigDetails(fieldKeyFormat string) (oci_certificates_management.CreateCertificateConfigDetails, error) {
	var baseObject oci_certificates_management.CreateCertificateConfigDetails
	//discriminator
	configTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_type"))
	var configType string
	if ok {
		configType = configTypeRaw.(string)
	} else {
		configType = "" // default value
	}
	switch strings.ToLower(configType) {
	case strings.ToLower("IMPORTED"):
		details := oci_certificates_management.CreateCertificateByImportingConfigDetails{}
		if certChainPem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cert_chain_pem")); ok {
			tmp := certChainPem.(string)
			details.CertChainPem = &tmp
		}
		if certificatePem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_pem")); ok {
			tmp := certificatePem.(string)
			details.CertificatePem = &tmp
		}
		if privateKeyPem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_key_pem")); ok {
			tmp := privateKeyPem.(string)
			details.PrivateKeyPem = &tmp
		}
		if privateKeyPemPassphrase, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_key_pem_passphrase")); ok {
			tmp := privateKeyPemPassphrase.(string)
			details.PrivateKeyPemPassphrase = &tmp
		}
		if versionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_name")); ok {
			tmp := versionName.(string)
			details.VersionName = &tmp
		}

		baseObject = details
	case strings.ToLower("ISSUED_BY_INTERNAL_CA"):
		details := oci_certificates_management.CreateCertificateIssuedByInternalCaConfigDetails{}
		if validity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validity")); ok {
			if tmpList := validity.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "validity"), 0)
				tmp, err := s.mapToValidity(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert validity, encountered error: %v", err)
				}
				details.Validity = &tmp
			}
		}
		if versionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_name")); ok {
			tmp := versionName.(string)
			details.VersionName = &tmp
		}

		if profileType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_profile_type")); ok {
			details.CertificateProfileType = oci_certificates_management.CertificateProfileTypeEnum(profileType.(string))
		}

		if issuerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "issuer_certificate_authority_id")); ok {
			tmp := issuerId.(string)
			details.IssuerCertificateAuthorityId = &tmp
		}

		if subject, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subject")); ok {
			if tmpList := subject.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "subject"), 0)
				tmp, err := s.mapToCertificateSubject(fieldKeyFormatNextLevel)
				if err != nil {
					return baseObject, err
				}
				details.Subject = &tmp
			}
		}

		if keyAlgorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_algorithm")); ok {
			details.KeyAlgorithm = oci_certificates_management.KeyAlgorithmEnum(keyAlgorithm.(string))
		}
		if signatureAlgorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "signature_algorithm")); ok {
			details.SignatureAlgorithm = oci_certificates_management.SignatureAlgorithmEnum(signatureAlgorithm.(string))
		}

		if subjectAlternativeNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subject_alternative_names")); ok {
			interfaces := subjectAlternativeNames.([]interface{})
			tmp := make([]oci_certificates_management.CertificateSubjectAlternativeName, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "subject_alternative_names"), stateDataIndex)
				converted, err := s.mapToCertificateSubjectAlternativeNames(fieldKeyFormatNextLevel)
				if err != nil {
					return baseObject, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("subject_alternative_names") {
				details.SubjectAlternativeNames = tmp
			}
		}

		baseObject = details
	case strings.ToLower("MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA"):
		details := oci_certificates_management.CreateCertificateManagedExternallyIssuedByInternalCaConfigDetails{}
		if csrPem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "csr_pem")); ok {
			tmp := csrPem.(string)
			details.CsrPem = &tmp
		}
		if validity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validity")); ok {
			if tmpList := validity.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "validity"), 0)
				tmp, err := s.mapToValidity(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert validity, encountered error: %v", err)
				}
				details.Validity = &tmp
			}
		}
		if versionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_name")); ok {
			tmp := versionName.(string)
			details.VersionName = &tmp
		}

		if issuerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "issuer_certificate_authority_id")); ok {
			tmp := issuerId.(string)
			details.IssuerCertificateAuthorityId = &tmp
		}

		baseObject = details
	default:
		return nil, fmt.Errorf("unknown config_type '%v' was specified", configType)
	}
	return baseObject, nil
}

func (s *CertificatesManagementCertificateResourceCrud) mapToCertificateSubjectAlternativeNames(fieldKeyFormat string) (oci_certificates_management.CertificateSubjectAlternativeName, error) {
	var subjectAlternativeName oci_certificates_management.CertificateSubjectAlternativeName
	if sanType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		subjectAlternativeName.Type = oci_certificates_management.CertificateSubjectAlternativeNameTypeEnum(sanType.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		subjectAlternativeName.Value = &tmp
	}

	return subjectAlternativeName, nil
}

func (s *CertificatesManagementCertificateResourceCrud) mapToCertificateSubject(fieldKeyFormat string) (oci_certificates_management.CertificateSubject, error) {
	result := oci_certificates_management.CertificateSubject{}

	if commonName, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "common_name")); ok {
		tmp := commonName.(string)
		result.CommonName = &tmp
	}

	if country, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "country")); ok {
		tmp := country.(string)
		result.Country = &tmp
	}

	if distinguishedNameQualifier, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "distinguished_name_qualifier")); ok {
		tmp := distinguishedNameQualifier.(string)
		result.DistinguishedNameQualifier = &tmp
	}

	if domainComponent, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "domain_component")); ok {
		tmp := domainComponent.(string)
		result.DomainComponent = &tmp
	}

	if generationQualifier, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "generation_qualifier")); ok {
		tmp := generationQualifier.(string)
		result.GenerationQualifier = &tmp
	}

	if givenName, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "given_name")); ok {
		tmp := givenName.(string)
		result.GivenName = &tmp
	}

	if initials, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "initials")); ok {
		tmp := initials.(string)
		result.Initials = &tmp
	}

	if localityName, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "locality_name")); ok {
		tmp := localityName.(string)
		result.LocalityName = &tmp
	}

	if organization, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "organization")); ok {
		tmp := organization.(string)
		result.Organization = &tmp
	}

	if organizationalUnit, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "organizational_unit")); ok {
		tmp := organizationalUnit.(string)
		result.OrganizationalUnit = &tmp
	}

	if pseudonym, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "pseudonym")); ok {
		tmp := pseudonym.(string)
		result.Pseudonym = &tmp
	}

	if serialNumber, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "serial_number")); ok {
		tmp := serialNumber.(string)
		result.SerialNumber = &tmp
	}

	if stateOrProvinceName, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "state_or_province_name")); ok {
		tmp := stateOrProvinceName.(string)
		result.StateOrProvinceName = &tmp
	}

	if street, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "street")); ok {
		tmp := street.(string)
		result.Street = &tmp
	}

	if surname, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "surname")); ok {
		tmp := surname.(string)
		result.Surname = &tmp
	}

	if title, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "title")); ok {
		tmp := title.(string)
		result.Title = &tmp
	}

	if userId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "user_id")); ok {
		tmp := userId.(string)
		result.UserId = &tmp
	}

	return result, nil
}

func CreateCertificateConfigDetailsToMap(obj *oci_certificates_management.CreateCertificateConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_certificates_management.UpdateCertificateByImportingConfigDetails:
		result["config_type"] = "IMPORTED"

		if v.CertChainPem != nil {
			result["cert_chain_pem"] = string(*v.CertChainPem)
		}

		if v.CertificatePem != nil {
			result["certificate_pem"] = string(*v.CertificatePem)
		}

		if v.PrivateKeyPem != nil {
			result["private_key_pem"] = string(*v.PrivateKeyPem)
		}

		if v.PrivateKeyPemPassphrase != nil {
			result["private_key_pem_passphrase"] = string(*v.PrivateKeyPemPassphrase)
		}
	case oci_certificates_management.UpdateCertificateIssuedByInternalCaConfigDetails:
		result["config_type"] = "ISSUED_BY_INTERNAL_CA"

		if v.Validity != nil {
			result["validity"] = []interface{}{ValidityToMap(v.Validity)}
		}
	case oci_certificates_management.UpdateCertificateManagedExternallyIssuedByInternalCaConfigDetails:
		result["config_type"] = "MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA"

		if v.CsrPem != nil {
			result["csr_pem"] = string(*v.CsrPem)
		}

		if v.Validity != nil {
			result["validity"] = []interface{}{ValidityToMap(v.Validity)}
		}
	default:
		log.Printf("[WARN] Received 'config_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CertificatesManagementCertificateResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_certificates_management.ChangeCertificateCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CertificateId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	_, err := s.Client.ChangeCertificateCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *CertificatesManagementCertificateResourceCrud) mapToCertificateRule(fieldKeyFormat string) (oci_certificates_management.CertificateRule, error) {
	var baseObject oci_certificates_management.CertificateRule

	//discriminator
	ruleTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_type"))
	var ruleType string
	if ok {
		ruleType = ruleTypeRaw.(string)
	} else {
		ruleType = "" // default value
	}
	switch strings.ToLower(ruleType) {
	case strings.ToLower("CERTIFICATE_RENEWAL_RULE"):
		details := oci_certificates_management.CertificateRenewalRule{}
		if advanceRenewalPeriod, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "advance_renewal_period")); ok {
			tmp := advanceRenewalPeriod.(string)
			details.AdvanceRenewalPeriod = &tmp
		}
		if renewalInterval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "renewal_interval")); ok {
			tmp := renewalInterval.(string)
			details.RenewalInterval = &tmp
		}
		baseObject = &details
	default:
		log.Printf("[WARN] Received 'rule_type' of unknown type %v", ruleType)
		return baseObject, nil
	}

	return baseObject, nil
}

func (s *CertificatesManagementCertificateResourceCrud) mapToUpdateCertificateConfigDetails(fieldKeyFormat string) (oci_certificates_management.UpdateCertificateConfigDetails, error) {
	var baseObject oci_certificates_management.UpdateCertificateConfigDetails
	//discriminator
	configTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_type"))
	var configType string
	if ok {
		configType = configTypeRaw.(string)
	} else {
		configType = "" // default value
	}
	switch strings.ToLower(configType) {
	case strings.ToLower("IMPORTED"):
		details := oci_certificates_management.UpdateCertificateByImportingConfigDetails{}
		if certChainPem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cert_chain_pem")); ok {
			tmp := certChainPem.(string)
			details.CertChainPem = &tmp
		}
		if certificatePem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_pem")); ok {
			tmp := certificatePem.(string)
			details.CertificatePem = &tmp
		}
		if privateKeyPem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_key_pem")); ok {
			tmp := privateKeyPem.(string)
			details.PrivateKeyPem = &tmp
		}
		if privateKeyPemPassphrase, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_key_pem_passphrase")); ok {
			tmp := privateKeyPemPassphrase.(string)
			details.PrivateKeyPemPassphrase = &tmp
		}

		baseObject = details
	case strings.ToLower("ISSUED_BY_INTERNAL_CA"):
		details := oci_certificates_management.UpdateCertificateIssuedByInternalCaConfigDetails{}
		if validity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validity")); ok {
			if tmpList := validity.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "validity"), 0)
				tmp, err := s.mapToValidity(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert validity, encountered error: %v", err)
				}
				details.Validity = &tmp
			}
		}
		if versionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_name")); ok {
			tmp := versionName.(string)
			details.VersionName = &tmp
		}

		baseObject = details
	case strings.ToLower("MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA"):
		details := oci_certificates_management.UpdateCertificateManagedExternallyIssuedByInternalCaConfigDetails{}
		if csrPem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "csr_pem")); ok {
			tmp := csrPem.(string)
			details.CsrPem = &tmp
		}
		if validity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validity")); ok {
			if tmpList := validity.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "validity"), 0)
				tmp, err := s.mapToValidity(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert validity, encountered error: %v", err)
				}
				details.Validity = &tmp
			}
		}
		if versionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_name")); ok {
			tmp := versionName.(string)
			details.VersionName = &tmp
		}

		baseObject = details
	default:
		return nil, fmt.Errorf("unknown config_type '%v' was specified", configType)
	}

	return baseObject, nil
}

func (s *CertificatesManagementCertificateResourceCrud) mapToValidity(fieldKeyFormat string) (oci_certificates_management.Validity, error) {
	result := oci_certificates_management.Validity{}

	if timeOfValidityNotAfter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_of_validity_not_after")); ok {
		tmp, err := time.Parse(time.RFC3339Nano, timeOfValidityNotAfter.(string))
		if err != nil {
			return result, err
		}
		result.TimeOfValidityNotAfter = &oci_common.SDKTime{Time: tmp}
	}

	if timeOfValidityNotBefore, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_of_validity_not_before")); ok {
		tmp, err := time.Parse(time.RFC3339Nano, timeOfValidityNotBefore.(string))
		if err != nil {
			return result, err
		}
		result.TimeOfValidityNotBefore = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}
