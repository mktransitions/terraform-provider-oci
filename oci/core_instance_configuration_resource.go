// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func InstanceConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createInstanceConfiguration,
		Read:     readInstanceConfiguration,
		Update:   updateInstanceConfiguration,
		Delete:   deleteInstanceConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_details": {
				Type:     schema.TypeList,
				Required: true,
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
												"display_name": {
													Type:     schema.TypeString,
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
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
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
									"availability_domain": {
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
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"hostname_label": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
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
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem:     schema.TypeString,
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
									"metadata": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem:     schema.TypeString,
									},
									"shape": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
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
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"hostname_label": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
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

func createInstanceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient

	return CreateResource(d, sync)
}

func readInstanceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient

	return ReadResource(sync)
}

func updateInstanceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient

	return UpdateResource(d, sync)
}

func deleteInstanceConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeManagementClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type InstanceConfigurationResourceCrud struct {
	BaseCrud
	Client                 *oci_core.ComputeManagementClient
	Res                    *oci_core.InstanceConfiguration
	DisableNotFoundRetries bool
}

func (s *InstanceConfigurationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *InstanceConfigurationResourceCrud) Create() error {
	request := oci_core.CreateInstanceConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CreateInstanceConfiguration.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.CreateInstanceConfiguration.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.CreateInstanceConfiguration.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.CreateInstanceConfiguration.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if instanceDetails, ok := s.D.GetOkExists("instance_details"); ok {
		if tmpList := instanceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "instance_details", 0)
			tmp, err := s.mapToInstanceConfigurationInstanceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CreateInstanceConfiguration.InstanceDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateInstanceConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceConfiguration
	return nil
}

func (s *InstanceConfigurationResourceCrud) Get() error {
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

func (s *InstanceConfigurationResourceCrud) Update() error {
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

func (s *InstanceConfigurationResourceCrud) Delete() error {
	request := oci_core.DeleteInstanceConfigurationRequest{}

	tmp := s.D.Id()
	request.InstanceConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteInstanceConfiguration(context.Background(), request)
	return err
}

func (s *InstanceConfigurationResourceCrud) SetData() error {
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
		if instanceDetailsMap := InstanceConfigurationInstanceDetailsToMap(&s.Res.InstanceDetails); instanceDetailsMap != nil {
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

func (s *InstanceConfigurationResourceCrud) mapToInstanceConfigurationAttachVnicDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationAttachVnicDetails, error) {
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

func InstanceConfigurationAttachVnicDetailsToMap(obj oci_core.InstanceConfigurationAttachVnicDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CreateVnicDetails != nil {
		result["create_vnic_details"] = []interface{}{InstanceConfigurationCreateVnicDetailsToMap(obj.CreateVnicDetails)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.NicIndex != nil {
		result["nic_index"] = int(*obj.NicIndex)
	}

	return result
}

func (s *InstanceConfigurationResourceCrud) mapToInstanceConfigurationAttachVolumeDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationAttachVolumeDetails, error) {
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
		if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if isReadOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_read_only")); ok {
			tmp := isReadOnly.(bool)
			details.IsReadOnly = &tmp
		}
		baseObject = details
	case strings.ToLower("paravirtualized"):
		details := oci_core.InstanceConfigurationParavirtualizedAttachVolumeDetails{}
		if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if isReadOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_read_only")); ok {
			tmp := isReadOnly.(bool)
			details.IsReadOnly = &tmp
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

		result["display_name"] = *v.DisplayName

		if v.IsReadOnly != nil {
			result["is_read_only"] = bool(*v.IsReadOnly)
		}

		if v.UseChap != nil {
			result["use_chap"] = bool(*v.UseChap)
		}
	case oci_core.InstanceConfigurationParavirtualizedAttachVolumeDetails:
		result["type"] = "paravirtualized"

		result["display_name"] = *v.DisplayName

		if v.IsReadOnly != nil {
			result["is_read_only"] = bool(*v.IsReadOnly)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *InstanceConfigurationResourceCrud) mapToInstanceConfigurationBlockVolumeDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationBlockVolumeDetails, error) {
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

func (s *InstanceConfigurationResourceCrud) mapToInstanceConfigurationCreateVnicDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationCreateVnicDetails, error) {
	result := oci_core.InstanceConfigurationCreateVnicDetails{}

	if assignPublicIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assign_public_ip")); ok {
		tmp := assignPublicIp.(bool)
		result.AssignPublicIp = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if hostnameLabel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname_label")); ok {
		tmp := hostnameLabel.(string)
		result.HostnameLabel = &tmp
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

func InstanceConfigurationCreateVnicDetailsToMap(obj *oci_core.InstanceConfigurationCreateVnicDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssignPublicIp != nil {
		result["assign_public_ip"] = bool(*obj.AssignPublicIp)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.HostnameLabel != nil {
		result["hostname_label"] = string(*obj.HostnameLabel)
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

func (s *InstanceConfigurationResourceCrud) mapToInstanceConfigurationCreateVolumeDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationCreateVolumeDetails, error) {
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

	return result
}

func (s *InstanceConfigurationResourceCrud) mapToInstanceConfigurationInstanceDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationInstanceDetails, error) {
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
			details.BlockVolumes = []oci_core.InstanceConfigurationBlockVolumeDetails{}
			interfaces := blockVolumes.([]interface{})
			tmp := make([]oci_core.InstanceConfigurationBlockVolumeDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf(fieldKeyFormat, fmt.Sprintf("%s.%d.%%s", "block_volumes", stateDataIndex))
				converted, err := s.mapToInstanceConfigurationBlockVolumeDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return nil, err
				}
				tmp[i] = converted
			}
			details.BlockVolumes = tmp
		}
		if launchDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "launch_details")); ok {
			if tmpList := launchDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf(fieldKeyFormat, fmt.Sprintf("%s.%d.%%s", "launch_details", 0))
				tmp, err := s.mapToInstanceConfigurationLaunchInstanceDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return nil, err
				}
				details.LaunchDetails = &tmp
			}
		}
		if secondaryVnics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secondary_vnics")); ok {
			details.SecondaryVnics = []oci_core.InstanceConfigurationAttachVnicDetails{}
			interfaces := secondaryVnics.([]interface{})
			tmp := make([]oci_core.InstanceConfigurationAttachVnicDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf(fieldKeyFormat, fmt.Sprintf("%s.%d.%%s", "secondary_vnics", stateDataIndex))
				converted, err := s.mapToInstanceConfigurationAttachVnicDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return nil, err
				}
				tmp[i] = converted
			}
			details.SecondaryVnics = tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown instance_type '%v' was specified", instanceType)
	}
	return baseObject, nil
}

func InstanceConfigurationInstanceDetailsToMap(obj *oci_core.InstanceConfigurationInstanceDetails) map[string]interface{} {
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
			result["launch_details"] = []interface{}{InstanceConfigurationLaunchInstanceDetailsToMap(v.LaunchDetails)}
		}

		secondaryVnics := []interface{}{}
		for _, item := range v.SecondaryVnics {
			secondaryVnics = append(secondaryVnics, InstanceConfigurationAttachVnicDetailsToMap(item))
		}
		result["secondary_vnics"] = secondaryVnics
	default:
		log.Printf("[WARN] Received 'instance_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *InstanceConfigurationResourceCrud) mapToInstanceConfigurationInstanceSourceDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationInstanceSourceDetails, error) {
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

		if v.ImageId != nil {
			result["image_id"] = string(*v.ImageId)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *InstanceConfigurationResourceCrud) mapToInstanceConfigurationLaunchInstanceDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationLaunchInstanceDetails, error) {
	result := oci_core.InstanceConfigurationLaunchInstanceDetails{}

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

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if ipxeScript, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipxe_script")); ok {
		tmp := ipxeScript.(string)
		result.IpxeScript = &tmp
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		result.Metadata = objectMapToStringMap(metadata.(map[string]interface{}))
	}

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		result.Shape = &tmp
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

func InstanceConfigurationLaunchInstanceDetailsToMap(obj *oci_core.InstanceConfigurationLaunchInstanceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CreateVnicDetails != nil {
		result["create_vnic_details"] = []interface{}{InstanceConfigurationCreateVnicDetailsToMap(obj.CreateVnicDetails)}
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["extended_metadata"] = obj.ExtendedMetadata

	result["freeform_tags"] = obj.FreeformTags

	if obj.IpxeScript != nil {
		result["ipxe_script"] = string(*obj.IpxeScript)
	}

	result["metadata"] = obj.Metadata

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
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

func (s *InstanceConfigurationResourceCrud) mapToInstanceConfigurationVolumeSourceDetails(fieldKeyFormat string) (oci_core.InstanceConfigurationVolumeSourceDetails, error) {
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
