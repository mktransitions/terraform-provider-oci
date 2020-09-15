// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_core "github.com/oracle/oci-go-sdk/v25/core"
)

func init() {
	RegisterResource("oci_core_instance_configuration", CoreInstanceConfigurationResource())
}

func CoreInstanceConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreInstanceConfiguration,
		Read:     readCoreInstanceConfiguration,
		Update:   updateCoreInstanceConfiguration,
		Delete:   deleteCoreInstanceConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"instance_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"instance_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"compute",
							}, true),
						},

						// Optional
						"block_volumes": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"attach_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"type": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"iscsi",
														"paravirtualized",
													}, true),
												},

												// Optional
												"device": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_pv_encryption_in_transit_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_read_only": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_shareable": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"use_chap": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"create_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"availability_domain": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
												},
												"backup_policy_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"compartment_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"defined_tags": {
													Type:             schema.TypeMap,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: definedTagsDiffSuppressFunction,
													Elem:             schema.TypeString,
												},
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"freeform_tags": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem:     schema.TypeString,
												},
												"kms_key_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"size_in_gbs": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													ValidateFunc:     validateInt64TypeString,
													DiffSuppressFunc: int64StringDiffSuppressFunction,
												},
												"source_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"type": {
																Type:             schema.TypeString,
																Required:         true,
																ForceNew:         true,
																DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"volume",
																	"volumeBackup",
																}, true),
															},

															// Optional
															"id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"vpus_per_gb": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													ValidateFunc:     validateInt64TypeString,
													DiffSuppressFunc: int64StringDiffSuppressFunction,
												},

												// Computed
											},
										},
									},
									"volume_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"launch_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"agent_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"is_management_disabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_monitoring_disabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"availability_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"recovery_action": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"availability_domain": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"create_vnic_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"assign_public_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"defined_tags": {
													Type:             schema.TypeMap,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: definedTagsDiffSuppressFunction,
													Elem:             schema.TypeString,
												},
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"freeform_tags": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem:     schema.TypeString,
												},
												"hostname_label": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"nsg_ids": {
													Type:     schema.TypeSet,
													Optional: true,
													ForceNew: true,
													Set:      literalTypeHashCodeForSets,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"private_ip": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"skip_source_dest_check": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"subnet_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"dedicated_vm_host_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"defined_tags": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										DiffSuppressFunc: definedTagsDiffSuppressFunction,
										Elem:             schema.TypeString,
									},
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"extended_metadata": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										DiffSuppressFunc: jsonStringDiffSuppressFunction,
										Elem:             schema.TypeString,
									},
									"fault_domain": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem:     schema.TypeString,
									},
									"ipxe_script": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"is_pv_encryption_in_transit_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"launch_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"launch_options": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"boot_volume_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"firmware": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_consistent_volume_naming_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"is_pv_encryption_in_transit_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"network_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"remote_data_volume_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"metadata": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem:     schema.TypeString,
									},
									"preferred_maintenance_action": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"shape": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"shape_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"ocpus": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"source_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"source_type": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"bootVolume",
														"image",
													}, true),
												},

												// Optional
												"boot_volume_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"boot_volume_size_in_gbs": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													ValidateFunc:     validateInt64TypeString,
													DiffSuppressFunc: int64StringDiffSuppressFunction,
												},
												"image_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},
						"secondary_vnics": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"create_vnic_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"assign_public_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"defined_tags": {
													Type:             schema.TypeMap,
													Optional:         true,
													Computed:         true,
													ForceNew:         true,
													DiffSuppressFunc: definedTagsDiffSuppressFunction,
													Elem:             schema.TypeString,
												},
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"freeform_tags": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem:     schema.TypeString,
												},
												"hostname_label": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"nsg_ids": {
													Type:     schema.TypeSet,
													Optional: true,
													ForceNew: true,
													Set:      literalTypeHashCodeForSets,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"private_ip": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"skip_source_dest_check": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"subnet_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"nic_index": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"source": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"INSTANCE",
					"NONE",
				}, true),
			},

			// Computed
			"deferred_fields": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreInstanceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient()

	return CreateResource(d, sync)
}

func readCoreInstanceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient()

	return ReadResource(sync)
}

func updateCoreInstanceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient()

	return UpdateResource(d, sync)
}

func deleteCoreInstanceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreInstanceConfigurationResourceCrud struct {
	BaseCrud
	Client                 *oci_core.ComputeManagementClient
	Res                    *oci_core.InstanceConfiguration
	DisableNotFoundRetries bool
}

func (s *CoreInstanceConfigurationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreInstanceConfigurationResourceCrud) Create() error {
	request := oci_core.CreateInstanceConfigurationRequest{}
	err := s.populateTopLevelPolymorphicCreateInstanceConfigurationRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateInstanceConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceConfiguration
	return nil
}

func (s *CoreInstanceConfigurationResourceCrud) Get() error {
	request := oci_core.GetInstanceConfigurationRequest{}

	tmp := s.D.Id()
	request.InstanceConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetInstanceConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceConfiguration
	return nil
}

func (s *CoreInstanceConfigurationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateInstanceConfigurationRequest{}

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

	tmp := s.D.Id()
	request.InstanceConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateInstanceConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceConfiguration
	return nil
}

func (s *CoreInstanceConfigurationResourceCrud) Delete() error {
	request := oci_core.DeleteInstanceConfigurationRequest{}

	tmp := s.D.Id()
	request.InstanceConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteInstanceConfiguration(context.Background(), request)
	return err
}

func (s *CoreInstanceConfigurationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("deferred_fields", s.Res.DeferredFields)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceDetails != nil {
		instanceDetailsArray := []interface{}{}
		if instanceDetailsMap := InstanceConfigurationInstanceDetailsToMap(&s.Res.InstanceDetails, false); instanceDetailsMap != nil {
			instanceDetailsArray = append(instanceDetailsArray, instanceDetailsMap)
		}
		s.D.Set("instance_details", instanceDetailsArray)
	} else {
		s.D.Set("instance_details", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationAttachVnicDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationAttachVnicDetails, error) {
	result := oci_core.InstanceConfigurationAttachVnicDetails{}

	if createVnicDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "create_vnic_details")); ok {
		if tmpList := createVnicDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "create_vnic_details"), 0)
			tmp, err := s.mapToInstanceConfigurationCreateVnicDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert create_vnic_details, encountered error: %v", err)
			}
			result.CreateVnicDetails = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if nicIndex, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nic_index")); ok {
		tmp := nicIndex.(int)
		result.NicIndex = &tmp
	}

	return result, nil
}

func InstanceConfigurationAttachVnicDetailsToMap(obj oci_core.InstanceConfigurationAttachVnicDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CreateVnicDetails != nil {
		result["create_vnic_details"] = []interface{}{InstanceConfigurationCreateVnicDetailsToMap(obj.CreateVnicDetails, datasource)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.NicIndex != nil {
		result["nic_index"] = int(*obj.NicIndex)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationAttachVolumeDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationAttachVolumeDetails, error) {
	var baseObject oci_core.InstanceConfigurationAttachVolumeDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("iscsi"):
		details := oci_core.InstanceConfigurationIscsiAttachVolumeDetails{}
		if useChap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "use_chap")); ok {
			tmp := useChap.(bool)
			details.UseChap = &tmp
		}
		if device, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "device")); ok {
			tmp := device.(string)
			details.Device = &tmp
		}
		if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if isReadOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_read_only")); ok {
			tmp := isReadOnly.(bool)
			details.IsReadOnly = &tmp
		}
		if isShareable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_shareable")); ok {
			tmp := isShareable.(bool)
			details.IsShareable = &tmp
		}
		baseObject = details
	case strings.ToLower("paravirtualized"):
		details := oci_core.InstanceConfigurationParavirtualizedAttachVolumeDetails{}
		if isPvEncryptionInTransitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pv_encryption_in_transit_enabled")); ok {
			tmp := isPvEncryptionInTransitEnabled.(bool)
			details.IsPvEncryptionInTransitEnabled = &tmp
		}
		if device, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "device")); ok {
			tmp := device.(string)
			details.Device = &tmp
		}
		if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if isReadOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_read_only")); ok {
			tmp := isReadOnly.(bool)
			details.IsReadOnly = &tmp
		}
		if isShareable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_shareable")); ok {
			tmp := isShareable.(bool)
			details.IsShareable = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func InstanceConfigurationAttachVolumeDetailsToMap(obj *oci_core.InstanceConfigurationAttachVolumeDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.InstanceConfigurationIscsiAttachVolumeDetails:
		result["type"] = "iscsi"

		if v.DisplayName != nil {
			result["display_name"] = *v.DisplayName
		}

		if v.IsReadOnly != nil {
			result["is_read_only"] = bool(*v.IsReadOnly)
		}

		if v.UseChap != nil {
			result["use_chap"] = bool(*v.UseChap)
		}

		if v.Device != nil {
			result["device"] = *v.Device
		}

		if v.IsShareable != nil {
			result["is_shareable"] = bool(*v.IsShareable)
		}
	case oci_core.InstanceConfigurationParavirtualizedAttachVolumeDetails:
		result["type"] = "paravirtualized"

		if v.IsPvEncryptionInTransitEnabled != nil {
			result["is_pv_encryption_in_transit_enabled"] = bool(*v.IsPvEncryptionInTransitEnabled)
		}

		if v.DisplayName != nil {
			result["display_name"] = *v.DisplayName
		}

		if v.Device != nil {
			result["device"] = *v.Device
		}

		if v.IsReadOnly != nil {
			result["is_read_only"] = bool(*v.IsReadOnly)
		}

		if v.IsShareable != nil {
			result["is_shareable"] = bool(*v.IsShareable)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationAvailabilityConfig(fieldKeyFormat string) (oci_core.InstanceConfigurationAvailabilityConfig, error) {
	result := oci_core.InstanceConfigurationAvailabilityConfig{}

	if recoveryAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_action")); ok {
		result.RecoveryAction = oci_core.InstanceConfigurationAvailabilityConfigRecoveryActionEnum(recoveryAction.(string))
	}

	return result, nil
}

func InstanceConfigurationAvailabilityConfigToMap(obj *oci_core.InstanceConfigurationAvailabilityConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["recovery_action"] = string(obj.RecoveryAction)

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationBlockVolumeDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationBlockVolumeDetails, error) {
	result := oci_core.InstanceConfigurationBlockVolumeDetails{}

	if attachDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attach_details")); ok {
		if tmpList := attachDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "attach_details"), 0)
			tmp, err := s.mapToInstanceConfigurationAttachVolumeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert attach_details, encountered error: %v", err)
			}
			result.AttachDetails = tmp
		}
	}

	if createDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "create_details")); ok {
		if tmpList := createDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "create_details"), 0)
			tmp, err := s.mapToInstanceConfigurationCreateVolumeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert create_details, encountered error: %v", err)
			}
			result.CreateDetails = &tmp
		}
	}

	if volumeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "volume_id")); ok {
		tmp := volumeId.(string)
		result.VolumeId = &tmp
	}

	return result, nil
}

func InstanceConfigurationBlockVolumeDetailsToMap(obj oci_core.InstanceConfigurationBlockVolumeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AttachDetails != nil {
		attachDetailsArray := []interface{}{}
		if attachDetailsMap := InstanceConfigurationAttachVolumeDetailsToMap(&obj.AttachDetails); attachDetailsMap != nil {
			attachDetailsArray = append(attachDetailsArray, attachDetailsMap)
		}
		result["attach_details"] = attachDetailsArray
	}

	if obj.CreateDetails != nil {
		result["create_details"] = []interface{}{InstanceConfigurationCreateVolumeDetailsToMap(obj.CreateDetails)}
	}

	if obj.VolumeId != nil {
		result["volume_id"] = string(*obj.VolumeId)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationCreateVnicDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationCreateVnicDetails, error) {
	result := oci_core.InstanceConfigurationCreateVnicDetails{}

	if assignPublicIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_public_ip")); ok {
		tmp := assignPublicIp.(bool)
		result.AssignPublicIp = &tmp
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname_label")); ok {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
	}

	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
			result.NsgIds = tmp
		}
	}

	if privateIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_ip")); ok {
		tmp := privateIp.(string)
		result.PrivateIp = &tmp
	}

	if skipSourceDestCheck, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "skip_source_dest_check")); ok {
		tmp := skipSourceDestCheck.(bool)
		result.SkipSourceDestCheck = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func InstanceConfigurationCreateVnicDetailsToMap(obj *oci_core.InstanceConfigurationCreateVnicDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssignPublicIp != nil {
		result["assign_public_ip"] = bool(*obj.AssignPublicIp)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HostnameLabel != nil {
		result["hostname_label"] = string(*obj.HostnameLabel)
	}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(literalTypeHashCodeForSets, nsgIds)
	}

	if obj.PrivateIp != nil {
		result["private_ip"] = string(*obj.PrivateIp)
	}

	if obj.SkipSourceDestCheck != nil {
		result["skip_source_dest_check"] = bool(*obj.SkipSourceDestCheck)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationCreateVolumeDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationCreateVolumeDetails, error) {
	result := oci_core.InstanceConfigurationCreateVolumeDetails{}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if backupPolicyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_policy_id")); ok {
		tmp := backupPolicyId.(string)
		result.BackupPolicyId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
		tmp := kmsKeyId.(string)
		result.KmsKeyId = &tmp
	}

	if sizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "size_in_gbs")); ok {
		tmp := sizeInGBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert sizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.SizeInGBs = &tmpInt64
	}

	if sourceDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_details")); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_details"), 0)
			tmp, err := s.mapToInstanceConfigurationVolumeSourceDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_details, encountered error: %v", err)
			}
			result.SourceDetails = &tmp
		}
	}

	if vpusPerGB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpus_per_gb")); ok {
		tmp := vpusPerGB.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert vpusPerGB string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.VpusPerGB = &tmpInt64
	}

	return result, nil
}

func InstanceConfigurationCreateVolumeDetailsToMap(obj *oci_core.InstanceConfigurationCreateVolumeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.BackupPolicyId != nil {
		result["backup_policy_id"] = string(*obj.BackupPolicyId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = string(*obj.KmsKeyId)
	}

	if obj.SizeInGBs != nil {
		result["size_in_gbs"] = strconv.FormatInt(*obj.SizeInGBs, 10)
	}

	if obj.SourceDetails != nil {
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := InstanceConfigurationVolumeSourceDetailsToMap(&obj.SourceDetails); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		result["source_details"] = sourceDetailsArray
	}

	if obj.VpusPerGB != nil {
		result["vpus_per_gb"] = strconv.FormatInt(*obj.VpusPerGB, 10)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationInstanceDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationInstanceDetails, error) {
	var baseObject oci_core.InstanceConfigurationInstanceDetails
	//discriminator
	instanceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_type"))
	var instanceType string
	if ok {
		instanceType = instanceTypeRaw.(string)
	} else {
		instanceType = "" // default value
	}
	switch strings.ToLower(instanceType) {
	case strings.ToLower("compute"):
		details := oci_core.ComputeInstanceDetails{}
		if blockVolumes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_volumes")); ok {
			interfaces := blockVolumes.([]interface{})
			tmp := make([]oci_core.InstanceConfigurationBlockVolumeDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "block_volumes"), stateDataIndex)
				converted, err := s.mapToInstanceConfigurationBlockVolumeDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "block_volumes")) {
				details.BlockVolumes = tmp
			}
		}
		if launchDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "launch_details")); ok {
			if tmpList := launchDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "launch_details"), 0)
				tmp, err := s.mapToInstanceConfigurationLaunchInstanceDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert launch_details, encountered error: %v", err)
				}
				details.LaunchDetails = &tmp
			}
		}
		if secondaryVnics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secondary_vnics")); ok {
			interfaces := secondaryVnics.([]interface{})
			tmp := make([]oci_core.InstanceConfigurationAttachVnicDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "secondary_vnics"), stateDataIndex)
				converted, err := s.mapToInstanceConfigurationAttachVnicDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "secondary_vnics")) {
				details.SecondaryVnics = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown instance_type '%v' was specified", instanceType)
	}
	return baseObject, nil
}

func InstanceConfigurationInstanceDetailsToMap(obj *oci_core.InstanceConfigurationInstanceDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.ComputeInstanceDetails:
		result["instance_type"] = "compute"

		blockVolumes := []interface{}{}
		for _, item := range v.BlockVolumes {
			blockVolumes = append(blockVolumes, InstanceConfigurationBlockVolumeDetailsToMap(item))
		}
		result["block_volumes"] = blockVolumes

		if v.LaunchDetails != nil {
			result["launch_details"] = []interface{}{InstanceConfigurationLaunchInstanceDetailsToMap(v.LaunchDetails, datasource)}
		}

		secondaryVnics := []interface{}{}
		for _, item := range v.SecondaryVnics {
			secondaryVnics = append(secondaryVnics, InstanceConfigurationAttachVnicDetailsToMap(item, datasource))
		}
		result["secondary_vnics"] = secondaryVnics
	default:
		log.Printf("[WARN] Received 'instance_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationInstanceSourceDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationInstanceSourceDetails, error) {
	var baseObject oci_core.InstanceConfigurationInstanceSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("bootVolume"):
		details := oci_core.InstanceConfigurationInstanceSourceViaBootVolumeDetails{}
		if bootVolumeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_id")); ok {
			tmp := bootVolumeId.(string)
			details.BootVolumeId = &tmp
		}
		baseObject = details
	case strings.ToLower("image"):
		details := oci_core.InstanceConfigurationInstanceSourceViaImageDetails{}
		if bootVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_size_in_gbs")); ok {
			tmp := bootVolumeSizeInGBs.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert bootVolumeSizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.BootVolumeSizeInGBs = &tmpInt64
		}
		if imageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_id")); ok {
			tmp := imageId.(string)
			details.ImageId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func InstanceConfigurationInstanceSourceDetailsToMap(obj *oci_core.InstanceConfigurationInstanceSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.InstanceConfigurationInstanceSourceViaBootVolumeDetails:
		result["source_type"] = "bootVolume"

		if v.BootVolumeId != nil {
			result["boot_volume_id"] = string(*v.BootVolumeId)
		}
	case oci_core.InstanceConfigurationInstanceSourceViaImageDetails:
		result["source_type"] = "image"

		if v.BootVolumeSizeInGBs != nil {
			result["boot_volume_size_in_gbs"] = strconv.FormatInt(*v.BootVolumeSizeInGBs, 10)
		}

		if v.ImageId != nil {
			result["image_id"] = string(*v.ImageId)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationLaunchInstanceAgentConfigDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationLaunchInstanceAgentConfigDetails, error) {
	result := oci_core.InstanceConfigurationLaunchInstanceAgentConfigDetails{}

	if isManagementDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_management_disabled")); ok {
		tmp := isManagementDisabled.(bool)
		result.IsManagementDisabled = &tmp
	}

	if isMonitoringDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_monitoring_disabled")); ok {
		tmp := isMonitoringDisabled.(bool)
		result.IsMonitoringDisabled = &tmp
	}

	return result, nil
}

func InstanceConfigurationLaunchInstanceAgentConfigDetailsToMap(obj *oci_core.InstanceConfigurationLaunchInstanceAgentConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsManagementDisabled != nil {
		result["is_management_disabled"] = bool(*obj.IsManagementDisabled)
	}

	if obj.IsMonitoringDisabled != nil {
		result["is_monitoring_disabled"] = bool(*obj.IsMonitoringDisabled)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationLaunchInstanceDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationLaunchInstanceDetails, error) {
	result := oci_core.InstanceConfigurationLaunchInstanceDetails{}

	if agentConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "agent_config")); ok {
		if tmpList := agentConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "agent_config"), 0)
			tmp, err := s.mapToInstanceConfigurationLaunchInstanceAgentConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert agent_config, encountered error: %v", err)
			}
			result.AgentConfig = &tmp
		}
	}

	if availabilityConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_config")); ok {
		if tmpList := availabilityConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "availability_config"), 0)
			tmp, err := s.mapToInstanceConfigurationAvailabilityConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert availability_config, encountered error: %v", err)
			}
			result.AvailabilityConfig = &tmp
		}
	}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if createVnicDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "create_vnic_details")); ok {
		if tmpList := createVnicDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "create_vnic_details"), 0)
			tmp, err := s.mapToInstanceConfigurationCreateVnicDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert create_vnic_details, encountered error: %v", err)
			}
			result.CreateVnicDetails = &tmp
		}
	}

	if dedicatedVmHostId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dedicated_vm_host_id")); ok {
		tmp := dedicatedVmHostId.(string)
		result.DedicatedVmHostId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if extendedMetadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "extended_metadata")); ok {
		extendedMetadata, err := mapToExtendedMetadata(extendedMetadata.(map[string]interface{}))
		if err != nil {
			return result, err
		}
		result.ExtendedMetadata = extendedMetadata
	}

	if faultDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fault_domain")); ok {
		tmp := faultDomain.(string)
		result.FaultDomain = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if ipxeScript, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipxe_script")); ok {
		tmp := ipxeScript.(string)
		result.IpxeScript = &tmp
	}

	if isPvEncryptionInTransitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pv_encryption_in_transit_enabled")); ok {
		tmp := isPvEncryptionInTransitEnabled.(bool)
		result.IsPvEncryptionInTransitEnabled = &tmp
	}

	if launchMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "launch_mode")); ok {
		result.LaunchMode = oci_core.InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum(launchMode.(string))
	}

	if launchOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "launch_options")); ok {
		if tmpList := launchOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "launch_options"), 0)
			tmp, err := s.mapToInstanceConfigurationLaunchOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert launch_options, encountered error: %v", err)
			}
			result.LaunchOptions = &tmp
		}
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		result.Metadata = objectMapToStringMap(metadata.(map[string]interface{}))
	}

	if preferredMaintenanceAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preferred_maintenance_action")); ok {
		result.PreferredMaintenanceAction = oci_core.InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum(preferredMaintenanceAction.(string))
	}

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		result.Shape = &tmp
	}

	if shapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_config")); ok {
		if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "shape_config"), 0)
			tmp, err := s.mapToInstanceConfigurationLaunchInstanceShapeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert shape_config, encountered error: %v", err)
			}
			result.ShapeConfig = &tmp
		}
	}

	if sourceDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_details")); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_details"), 0)
			tmp, err := s.mapToInstanceConfigurationInstanceSourceDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_details, encountered error: %v", err)
			}
			result.SourceDetails = &tmp
		}
	}

	return result, nil
}

func InstanceConfigurationLaunchInstanceDetailsToMap(obj *oci_core.InstanceConfigurationLaunchInstanceDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentConfig != nil {
		result["agent_config"] = []interface{}{InstanceConfigurationLaunchInstanceAgentConfigDetailsToMap(obj.AgentConfig)}
	}

	if obj.AvailabilityConfig != nil {
		result["availability_config"] = []interface{}{InstanceConfigurationAvailabilityConfigToMap(obj.AvailabilityConfig)}
	}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CreateVnicDetails != nil {
		result["create_vnic_details"] = []interface{}{InstanceConfigurationCreateVnicDetailsToMap(obj.CreateVnicDetails, datasource)}
	}

	if obj.DedicatedVmHostId != nil {
		result["dedicated_vm_host_id"] = string(*obj.DedicatedVmHostId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["extended_metadata"] = genericMapToJsonMap(obj.ExtendedMetadata)

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.IpxeScript != nil {
		result["ipxe_script"] = string(*obj.IpxeScript)
	}

	if obj.IsPvEncryptionInTransitEnabled != nil {
		result["is_pv_encryption_in_transit_enabled"] = bool(*obj.IsPvEncryptionInTransitEnabled)
	}

	result["launch_mode"] = string(obj.LaunchMode)

	if obj.LaunchOptions != nil {
		result["launch_options"] = []interface{}{InstanceConfigurationLaunchOptionsToMap(obj.LaunchOptions)}
	}

	result["metadata"] = obj.Metadata

	result["preferred_maintenance_action"] = string(obj.PreferredMaintenanceAction)

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.ShapeConfig != nil {
		result["shape_config"] = []interface{}{InstanceConfigurationLaunchInstanceShapeConfigDetailsToMap(obj.ShapeConfig)}
	}

	if obj.SourceDetails != nil {
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := InstanceConfigurationInstanceSourceDetailsToMap(&obj.SourceDetails); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		result["source_details"] = sourceDetailsArray
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationLaunchInstanceShapeConfigDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationLaunchInstanceShapeConfigDetails, error) {
	result := oci_core.InstanceConfigurationLaunchInstanceShapeConfigDetails{}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	return result, nil
}

func InstanceConfigurationLaunchInstanceShapeConfigDetailsToMap(obj *oci_core.InstanceConfigurationLaunchInstanceShapeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationLaunchOptions(fieldKeyFormat string) (oci_core.InstanceConfigurationLaunchOptions, error) {
	result := oci_core.InstanceConfigurationLaunchOptions{}

	if bootVolumeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_type")); ok {
		result.BootVolumeType = oci_core.InstanceConfigurationLaunchOptionsBootVolumeTypeEnum(bootVolumeType.(string))
	}

	if firmware, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "firmware")); ok {
		result.Firmware = oci_core.InstanceConfigurationLaunchOptionsFirmwareEnum(firmware.(string))
	}

	if isConsistentVolumeNamingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_consistent_volume_naming_enabled")); ok {
		tmp := isConsistentVolumeNamingEnabled.(bool)
		result.IsConsistentVolumeNamingEnabled = &tmp
	}

	if isPvEncryptionInTransitEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_pv_encryption_in_transit_enabled")); ok {
		tmp := isPvEncryptionInTransitEnabled.(bool)
		result.IsPvEncryptionInTransitEnabled = &tmp
	}

	if networkType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_type")); ok {
		result.NetworkType = oci_core.InstanceConfigurationLaunchOptionsNetworkTypeEnum(networkType.(string))
	}

	if remoteDataVolumeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remote_data_volume_type")); ok {
		result.RemoteDataVolumeType = oci_core.InstanceConfigurationLaunchOptionsRemoteDataVolumeTypeEnum(remoteDataVolumeType.(string))
	}

	return result, nil
}

func InstanceConfigurationLaunchOptionsToMap(obj *oci_core.InstanceConfigurationLaunchOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["boot_volume_type"] = string(obj.BootVolumeType)

	result["firmware"] = string(obj.Firmware)

	if obj.IsConsistentVolumeNamingEnabled != nil {
		result["is_consistent_volume_naming_enabled"] = bool(*obj.IsConsistentVolumeNamingEnabled)
	}

	if obj.IsPvEncryptionInTransitEnabled != nil {
		result["is_pv_encryption_in_transit_enabled"] = bool(*obj.IsPvEncryptionInTransitEnabled)
	}

	result["network_type"] = string(obj.NetworkType)

	result["remote_data_volume_type"] = string(obj.RemoteDataVolumeType)

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) mapToInstanceConfigurationVolumeSourceDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationVolumeSourceDetails, error) {
	var baseObject oci_core.InstanceConfigurationVolumeSourceDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("volume"):
		details := oci_core.InstanceConfigurationVolumeSourceFromVolumeDetails{}
		if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	case strings.ToLower("volumeBackup"):
		details := oci_core.InstanceConfigurationVolumeSourceFromVolumeBackupDetails{}
		if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func InstanceConfigurationVolumeSourceDetailsToMap(obj *oci_core.InstanceConfigurationVolumeSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.InstanceConfigurationVolumeSourceFromVolumeDetails:
		result["type"] = "volume"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	case oci_core.InstanceConfigurationVolumeSourceFromVolumeBackupDetails:
		result["type"] = "volumeBackup"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreInstanceConfigurationResourceCrud) populateTopLevelPolymorphicCreateInstanceConfigurationRequest(request *oci_core.CreateInstanceConfigurationRequest) error {
	//discriminator
	sourceRaw, ok := s.D.GetOkExists("source")
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "NONE" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("INSTANCE"):
		details := oci_core.CreateInstanceConfigurationFromInstanceDetails{}
		if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
			tmp := instanceId.(string)
			details.InstanceId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateInstanceConfiguration = details
	case strings.ToLower("NONE"):
		details := oci_core.CreateInstanceConfigurationDetails{}
		if instanceDetails, ok := s.D.GetOkExists("instance_details"); ok {
			if tmpList := instanceDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "instance_details", 0)
				tmp, err := s.mapToInstanceConfigurationInstanceDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.InstanceDetails = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateInstanceConfiguration = details
	default:
		return fmt.Errorf("unknown source '%v' was specified", source)
	}
	return nil
}

func (s *CoreInstanceConfigurationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeInstanceConfigurationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.InstanceConfigurationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeInstanceConfigurationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
