// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func VolumeGroupBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createVolumeGroupBackup,
		Read:     readVolumeGroupBackup,
		Update:   updateVolumeGroupBackup,
		Delete:   deleteVolumeGroupBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"volume_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
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
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"size_in_gbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_mbs": {
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
			"time_request_received": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_size_in_gbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_size_in_mbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_backup_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createVolumeGroupBackup(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeGroupBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return CreateResource(d, sync)
}

func readVolumeGroupBackup(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeGroupBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return ReadResource(sync)
}

func updateVolumeGroupBackup(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeGroupBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return UpdateResource(d, sync)
}

func deleteVolumeGroupBackup(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeGroupBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type VolumeGroupBackupResourceCrud struct {
	BaseCrud
	Client                 *oci_core.BlockstorageClient
	Res                    *oci_core.VolumeGroupBackup
	DisableNotFoundRetries bool
}

func (s *VolumeGroupBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VolumeGroupBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VolumeGroupBackupLifecycleStateCreating),
		string(oci_core.VolumeGroupBackupLifecycleStateRequestReceived),
	}
}

func (s *VolumeGroupBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VolumeGroupBackupLifecycleStateCommitted),
		string(oci_core.VolumeGroupBackupLifecycleStateAvailable),
	}
}

func (s *VolumeGroupBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VolumeGroupBackupLifecycleStateTerminating),
	}
}

func (s *VolumeGroupBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VolumeGroupBackupLifecycleStateTerminated),
	}
}

func (s *VolumeGroupBackupResourceCrud) Create() error {
	request := oci_core.CreateVolumeGroupBackupRequest{}

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

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_core.CreateVolumeGroupBackupDetailsTypeEnum(type_.(string))
	}

	if volumeGroupId, ok := s.D.GetOkExists("volume_group_id"); ok {
		tmp := volumeGroupId.(string)
		request.VolumeGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVolumeGroupBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeGroupBackup
	return nil
}

func (s *VolumeGroupBackupResourceCrud) Get() error {
	request := oci_core.GetVolumeGroupBackupRequest{}

	tmp := s.D.Id()
	request.VolumeGroupBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVolumeGroupBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeGroupBackup
	return nil
}

func (s *VolumeGroupBackupResourceCrud) Update() error {
	request := oci_core.UpdateVolumeGroupBackupRequest{}

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
	request.VolumeGroupBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateVolumeGroupBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeGroupBackup
	return nil
}

func (s *VolumeGroupBackupResourceCrud) Delete() error {
	request := oci_core.DeleteVolumeGroupBackupRequest{}

	tmp := s.D.Id()
	request.VolumeGroupBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVolumeGroupBackup(context.Background(), request)
	return err
}

func (s *VolumeGroupBackupResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", strconv.FormatInt(*s.Res.SizeInGBs, 10))
	}

	if s.Res.SizeInMBs != nil {
		s.D.Set("size_in_mbs", strconv.FormatInt(*s.Res.SizeInMBs, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeRequestReceived != nil {
		s.D.Set("time_request_received", s.Res.TimeRequestReceived.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.UniqueSizeInGbs != nil {
		s.D.Set("unique_size_in_gbs", strconv.FormatInt(*s.Res.UniqueSizeInGbs, 10))
	}

	if s.Res.UniqueSizeInMbs != nil {
		s.D.Set("unique_size_in_mbs", strconv.FormatInt(*s.Res.UniqueSizeInMbs, 10))
	}

	s.D.Set("volume_backup_ids", s.Res.VolumeBackupIds)

	if s.Res.VolumeGroupId != nil {
		s.D.Set("volume_group_id", *s.Res.VolumeGroupId)
	}

	return nil
}
