// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreVolumeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreVolume,
		Read:     readCoreVolume,
		Update:   updateCoreVolume,
		Delete:   deleteCoreVolume,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"autotune_policies": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"autotune_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DETACHED_VOLUME",
								"PERFORMANCE_BASED",
							}, true),
						},

						// Optional
						"max_vpus_per_gb": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"backup_policy_id": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: tfresource.FieldDeprecatedButSupportedThroughAnotherResource("backup_policy_id", "oci_core_volume_backup_policy_assignment"),
			},
			"block_volume_replicas": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"availability_domain": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},

						// Optional
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"xrr_kms_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"block_volume_replica_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kms_key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"cluster_placement_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_volume_replicas_deletion": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
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
			"is_auto_tune_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"size_in_gbs": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"size_in_mbs": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
				Deprecated:       tfresource.FieldDeprecatedForAnother("size_in_mbs", "size_in_gbs"),
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
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"blockVolumeReplica",
								"volume",
								"volumeBackup",
								"volumeBackupDelta",
							}, true),
						},

						// Optional
						"change_block_size_in_bytes": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         false,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"first_backup_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
							ForceNew: true,
						},
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
							ForceNew: true,
						},
						"second_backup_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: false,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"volume_backup_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vpus_per_gb": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"xrc_kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"auto_tuned_vpus_per_gb": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_hydrated": {
				Type:     schema.TypeBool,
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreVolume(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreVolume(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

func updateCoreVolume(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreVolume(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreVolumeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.BlockstorageClient
	Res                    *oci_core.Volume
	DisableNotFoundRetries bool
}

func (s *CoreVolumeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreVolumeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VolumeLifecycleStateProvisioning),
		string(oci_core.VolumeLifecycleStateRestoring),
	}
}

func (s *CoreVolumeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VolumeLifecycleStateAvailable),
	}
}

func (s *CoreVolumeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VolumeLifecycleStateTerminating),
	}
}

func (s *CoreVolumeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VolumeLifecycleStateTerminated),
	}
}

func (s *CoreVolumeResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.VolumeLifecycleStateProvisioning),
	}
}

func (s *CoreVolumeResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.VolumeLifecycleStateAvailable),
	}
}

func (s *CoreVolumeResourceCrud) Create() error {
	request := oci_core.CreateVolumeRequest{}

	if autotunePolicies, ok := s.D.GetOkExists("autotune_policies"); ok {
		interfaces := autotunePolicies.([]interface{})
		tmp := make([]oci_core.AutotunePolicy, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "autotune_policies", stateDataIndex)
			converted, err := s.mapToAutotunePolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("autotune_policies") {
			request.AutotunePolicies = tmp
		}
	}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if backupPolicyId, ok := s.D.GetOkExists("backup_policy_id"); ok {
		tmp := backupPolicyId.(string)
		request.BackupPolicyId = &tmp
	}

	if blockVolumeReplicas, ok := s.D.GetOkExists("block_volume_replicas"); ok {
		interfaces := blockVolumeReplicas.([]interface{})
		tmp := make([]oci_core.BlockVolumeReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "block_volume_replicas", stateDataIndex)
			converted, err := s.mapToBlockVolumeReplicaDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("block_volume_replicas") {
			request.BlockVolumeReplicas = tmp
		}
	}

	if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
		tmp := clusterPlacementGroupId.(string)
		request.ClusterPlacementGroupId = &tmp
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAutoTuneEnabled, ok := s.D.GetOkExists("is_auto_tune_enabled"); ok {
		tmp := isAutoTuneEnabled.(bool)
		request.IsAutoTuneEnabled = &tmp
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if sizeInGBs, ok := s.D.GetOkExists("size_in_gbs"); ok {
		tmp := sizeInGBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SizeInGBs = &tmpInt64
	}

	if sizeInMBs, ok := s.D.GetOkExists("size_in_mbs"); ok {
		tmp := sizeInMBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sizeInMBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SizeInMBs = &tmpInt64
	}

	if request.SizeInMBs != nil && request.SizeInGBs != nil &&
		*request.SizeInMBs > 0 && *request.SizeInGBs > 0 {
		return fmt.Errorf("both size in Megabytes and Gigabytes cannot be set. Specify one or the other, or leave both undefined to use the default size")
	}

	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
			tmp, err := s.mapToVolumeSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SourceDetails = tmp
		}
	}

	if volumeBackupId, ok := s.D.GetOkExists("volume_backup_id"); ok {
		tmp := volumeBackupId.(string)
		request.VolumeBackupId = &tmp
	}

	if vpusPerGB, ok := s.D.GetOkExists("vpus_per_gb"); ok {
		tmp := vpusPerGB.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert vpusPerGB string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.VpusPerGB = &tmpInt64
	}

	if xrcKmsKeyId, ok := s.D.GetOkExists("xrc_kms_key_id"); ok {
		tmp := xrcKmsKeyId.(string)
		request.XrcKmsKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Volume
	return nil
}

func (s *CoreVolumeResourceCrud) Get() error {
	request := oci_core.GetVolumeRequest{}

	tmp := s.D.Id()
	request.VolumeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Volume
	return nil
}

func (s *CoreVolumeResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateVolumeRequest{}

	if autotunePolicies, ok := s.D.GetOkExists("autotune_policies"); ok {
		interfaces := autotunePolicies.([]interface{})
		tmp := make([]oci_core.AutotunePolicy, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "autotune_policies", stateDataIndex)
			converted, err := s.mapToAutotunePolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("autotune_policies") {
			request.AutotunePolicies = tmp
		}
	}

	if blockVolumeReplicas, ok := s.D.GetOkExists("block_volume_replicas"); ok {
		interfaces := blockVolumeReplicas.([]interface{})
		tmp := make([]oci_core.BlockVolumeReplicaDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "block_volume_replicas", stateDataIndex)
			converted, err := s.mapToBlockVolumeReplicaDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("block_volume_replicas") {
			request.BlockVolumeReplicas = tmp
		}
	}

	if blockVolumeReplicasDeletion, ok := s.D.GetOkExists("block_volume_replicas_deletion"); ok {
		tmp := blockVolumeReplicasDeletion.(bool)
		if tmp == true {
			request.BlockVolumeReplicas = []oci_core.BlockVolumeReplicaDetails{}
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAutoTuneEnabled, ok := s.D.GetOkExists("is_auto_tune_enabled"); ok {
		tmp := isAutoTuneEnabled.(bool)
		request.IsAutoTuneEnabled = &tmp
	}

	if s.D.HasChange("kms_key_id") {
		keyUpdateRequest := oci_core.UpdateVolumeKmsKeyRequest{}

		volumeId := s.D.Id()
		keyUpdateRequest.VolumeId = &volumeId

		tmp := s.D.Get("kms_key_id").(string)
		keyUpdateRequest.KmsKeyId = &tmp

		keyUpdateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

		_, err := s.Client.UpdateVolumeKmsKey(context.Background(), keyUpdateRequest)
		if err != nil {
			return err
		}
	}

	if sizeInGBs, ok := s.D.GetOkExists("size_in_gbs"); ok {
		tmp := sizeInGBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SizeInGBs = &tmpInt64
	}

	tmp := s.D.Id()
	request.VolumeId = &tmp

	if vpusPerGB, ok := s.D.GetOkExists("vpus_per_gb"); ok {
		tmp := vpusPerGB.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert vpusPerGB string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.VpusPerGB = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Volume
	return nil
}

func (s *CoreVolumeResourceCrud) Delete() error {
	request := oci_core.DeleteVolumeRequest{}

	tmp := s.D.Id()
	request.VolumeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVolume(context.Background(), request)
	return err
}

func (s *CoreVolumeResourceCrud) SetData() error {
	if s.Res.AutoTunedVpusPerGB != nil {
		s.D.Set("auto_tuned_vpus_per_gb", strconv.FormatInt(*s.Res.AutoTunedVpusPerGB, 10))
	}

	autotunePolicies := []interface{}{}
	for _, item := range s.Res.AutotunePolicies {
		autotunePolicies = append(autotunePolicies, BlockVolumeAutotunePolicyToMap(item))
	}
	s.D.Set("autotune_policies", autotunePolicies)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	blockVolumeReplicas := []interface{}{}
	for _, item := range s.Res.BlockVolumeReplicas {
		blockVolumeReplicas = append(blockVolumeReplicas, BlockVolumeReplicaInfoToMap(item))
	}
	s.D.Set("block_volume_replicas", blockVolumeReplicas)

	if s.Res.ClusterPlacementGroupId != nil {
		s.D.Set("cluster_placement_group_id", *s.Res.ClusterPlacementGroupId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutoTuneEnabled != nil {
		s.D.Set("is_auto_tune_enabled", *s.Res.IsAutoTuneEnabled)
	}

	if s.Res.IsHydrated != nil {
		s.D.Set("is_hydrated", *s.Res.IsHydrated)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", strconv.FormatInt(*s.Res.SizeInGBs, 10))
	}

	if s.Res.SizeInMBs != nil {
		s.D.Set("size_in_mbs", strconv.FormatInt(*s.Res.SizeInMBs, 10))
	}

	if s.Res.SourceDetails != nil {
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := VolumeSourceDetailsToMap(&s.Res.SourceDetails); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		s.D.Set("source_details", sourceDetailsArray)
	} else {
		s.D.Set("source_details", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	} else {
		system_tags := map[string]interface{}{}
		s.D.Set("system_tags", system_tags)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VolumeGroupId != nil {
		s.D.Set("volume_group_id", *s.Res.VolumeGroupId)
	}

	if s.Res.VpusPerGB != nil {
		s.D.Set("vpus_per_gb", strconv.FormatInt(*s.Res.VpusPerGB, 10))
	}

	// Add backup policy id from the other API
	backupPolicyId, err := getBackupPolicyId(s.Res.Id, s.Client)
	if err != nil {
		log.Printf("[ERROR] Received an error when fetching backup policy id %v", err)
	} else if backupPolicyId != nil {
		s.D.Set("backup_policy_id", backupPolicyId)
	}
	return nil
}

func (s *CoreVolumeResourceCrud) mapToAutotunePolicy(fieldKeyFormat string) (oci_core.AutotunePolicy, error) {
	var baseObject oci_core.AutotunePolicy
	//discriminator
	autotuneTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "autotune_type"))
	var autotuneType string
	if ok {
		autotuneType = autotuneTypeRaw.(string)
	} else {
		autotuneType = "" // default value
	}
	switch strings.ToLower(autotuneType) {
	case strings.ToLower("DETACHED_VOLUME"):
		details := oci_core.DetachedVolumeAutotunePolicy{}
		baseObject = details
	case strings.ToLower("PERFORMANCE_BASED"):
		details := oci_core.PerformanceBasedAutotunePolicy{}
		if maxVpusPerGB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_vpus_per_gb")); ok {
			tmp := maxVpusPerGB.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert maxVpusPerGB string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.MaxVpusPerGB = &tmpInt64
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown autotune_type '%v' was specified", autotuneType)
	}
	return baseObject, nil
}

func BlockVolumeAutotunePolicyToMap(obj oci_core.AutotunePolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_core.DetachedVolumeAutotunePolicy:
		result["autotune_type"] = "DETACHED_VOLUME"
	case oci_core.PerformanceBasedAutotunePolicy:
		result["autotune_type"] = "PERFORMANCE_BASED"

		if v.MaxVpusPerGB != nil {
			result["max_vpus_per_gb"] = strconv.FormatInt(*v.MaxVpusPerGB, 10)
		}
	default:
		log.Printf("[WARN] Received 'autotune_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *CoreVolumeResourceCrud) mapToBlockVolumeReplicaDetails(fieldKeyFormat string) (oci_core.BlockVolumeReplicaDetails, error) {
	result := oci_core.BlockVolumeReplicaDetails{}

	if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
		tmp := availabilityDomain.(string)
		result.AvailabilityDomain = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if xrrKmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "xrr_kms_key_id")); ok {
		tmp := xrrKmsKeyId.(string)
		result.XrrKmsKeyId = &tmp
	}

	return result, nil
}

func BlockVolumeReplicaInfoToMap(obj oci_core.BlockVolumeReplicaInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.BlockVolumeReplicaId != nil {
		result["block_volume_replica_id"] = string(*obj.BlockVolumeReplicaId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = string(*obj.KmsKeyId)
	}

	return result
}

func (s *CoreVolumeResourceCrud) mapToVolumeSourceDetails(fieldKeyFormat string) (oci_core.VolumeSourceDetails, error) {
	var baseObject oci_core.VolumeSourceDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("blockVolumeReplica"):
		details := oci_core.VolumeSourceFromBlockVolumeReplicaDetails{}
		if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	case strings.ToLower("volume"):
		details := oci_core.VolumeSourceFromVolumeDetails{}
		if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	case strings.ToLower("volumeBackup"):
		details := oci_core.VolumeSourceFromVolumeBackupDetails{}
		if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	case strings.ToLower("volumeBackupDelta"):
		details := oci_core.VolumeSourceFromVolumeBackupDeltaDetails{}
		if changeBlockSizeInBytes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "change_block_size_in_bytes")); ok {
			tmp := changeBlockSizeInBytes.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert changeBlockSizeInBytes string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.ChangeBlockSizeInBytes = &tmpInt64
		}
		if firstBackupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "first_backup_id")); ok {
			tmp := firstBackupId.(string)
			details.FirstBackupId = &tmp
		}
		if secondBackupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "second_backup_id")); ok {
			tmp := secondBackupId.(string)
			details.SecondBackupId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func VolumeSourceDetailsToMap(obj *oci_core.VolumeSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.VolumeSourceFromBlockVolumeReplicaDetails:
		result["type"] = "blockVolumeReplica"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	case oci_core.VolumeSourceFromVolumeDetails:
		result["type"] = "volume"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	case oci_core.VolumeSourceFromVolumeBackupDetails:
		result["type"] = "volumeBackup"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	case oci_core.VolumeSourceFromVolumeBackupDeltaDetails:
		result["type"] = "volumeBackupDelta"

		if v.ChangeBlockSizeInBytes != nil {
			result["change_block_size_in_bytes"] = strconv.FormatInt(*v.ChangeBlockSizeInBytes, 10)
		}

		if v.FirstBackupId != nil {
			result["first_backup_id"] = string(*v.FirstBackupId)
		}

		if v.SecondBackupId != nil {
			result["second_backup_id"] = string(*v.SecondBackupId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreVolumeResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeVolumeCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VolumeId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeVolumeCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
