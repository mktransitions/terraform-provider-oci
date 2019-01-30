// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func VolumeAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createVolumeAttachment,
		Read:     readVolumeAttachment,
		Delete:   deleteVolumeAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"attachment_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"iscsi",
					"paravirtualized",
				}, true),
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"volume_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"use_chap": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"chap_secret": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"chap_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
				// The legacy provider required this, but the API no longer accepts it. Keep as optional
				// to avoid a breaking change. The value will be ignored if defined in the config.
				Optional: true,
			},
			"ipv4": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iqn": { // iSCSI Qualified Name per RFC 3720
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
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
		},
	}
}

func createVolumeAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return CreateResource(d, sync)
}

func readVolumeAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return ReadResource(sync)
}

func deleteVolumeAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type VolumeAttachmentResourceCrud struct {
	BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.VolumeAttachment
	DisableNotFoundRetries bool
}

func (s *VolumeAttachmentResourceCrud) ID() string {
	volumeAttachment := *s.Res
	return *volumeAttachment.GetId()
}

func (s *VolumeAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VolumeAttachmentLifecycleStateAttaching),
	}
}

func (s *VolumeAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VolumeAttachmentLifecycleStateAttached),
	}
}

func (s *VolumeAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VolumeAttachmentLifecycleStateDetaching),
	}
}

func (s *VolumeAttachmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VolumeAttachmentLifecycleStateDetached),
	}
}

func (s *VolumeAttachmentResourceCrud) Create() error {
	request := oci_core.AttachVolumeRequest{}
	err := s.populateTopLevelPolymorphicAttachVolumeRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.AttachVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeAttachment
	return nil
}

func (s *VolumeAttachmentResourceCrud) Get() error {
	request := oci_core.GetVolumeAttachmentRequest{}

	tmp := s.D.Id()
	request.VolumeAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVolumeAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeAttachment
	return nil
}

func (s *VolumeAttachmentResourceCrud) Delete() error {
	request := oci_core.DetachVolumeRequest{}

	tmp := s.D.Id()
	request.VolumeAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DetachVolume(context.Background(), request)
	return err
}

func (s *VolumeAttachmentResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_core.IScsiVolumeAttachment:
		s.D.Set("attachment_type", "iscsi")

		if v.ChapSecret != nil {
			s.D.Set("chap_secret", *v.ChapSecret)
		}

		if v.ChapUsername != nil {
			s.D.Set("chap_username", *v.ChapUsername)
		}

		if v.Ipv4 != nil {
			s.D.Set("ipv4", *v.Ipv4)
		}

		if v.Iqn != nil {
			s.D.Set("iqn", *v.Iqn)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		if v.AvailabilityDomain != nil {
			s.D.Set("availability_domain", *v.AvailabilityDomain)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.Device != nil {
			s.D.Set("device", *v.Device)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.InstanceId != nil {
			s.D.Set("instance_id", *v.InstanceId)
		}

		if v.IsPvEncryptionInTransitEnabled != nil {
			s.D.Set("is_pv_encryption_in_transit_enabled", *v.IsPvEncryptionInTransitEnabled)
		}

		if v.IsReadOnly != nil {
			s.D.Set("is_read_only", *v.IsReadOnly)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.VolumeId != nil {
			s.D.Set("volume_id", *v.VolumeId)
		}
	case oci_core.ParavirtualizedVolumeAttachment:
		s.D.Set("attachment_type", "paravirtualized")

		if v.AvailabilityDomain != nil {
			s.D.Set("availability_domain", *v.AvailabilityDomain)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.Device != nil {
			s.D.Set("device", *v.Device)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.InstanceId != nil {
			s.D.Set("instance_id", *v.InstanceId)
		}

		if v.IsPvEncryptionInTransitEnabled != nil {
			s.D.Set("is_pv_encryption_in_transit_enabled", *v.IsPvEncryptionInTransitEnabled)
		}

		if v.IsReadOnly != nil {
			s.D.Set("is_read_only", *v.IsReadOnly)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.VolumeId != nil {
			s.D.Set("volume_id", *v.VolumeId)
		}
	default:
		log.Printf("[WARN] Received 'attachment_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *VolumeAttachmentResourceCrud) populateTopLevelPolymorphicAttachVolumeRequest(request *oci_core.AttachVolumeRequest) error {
	//discriminator
	attachmentTypeRaw, ok := s.D.GetOkExists("attachment_type")
	var attachmentType string
	if ok {
		attachmentType = attachmentTypeRaw.(string)
	} else {
		attachmentType = "" // default value
	}
	switch strings.ToLower(attachmentType) {
	case strings.ToLower("iscsi"):
		details := oci_core.AttachIScsiVolumeDetails{}
		if useChap, ok := s.D.GetOkExists("use_chap"); ok {
			tmp := useChap.(bool)
			details.UseChap = &tmp
		}
		if device, ok := s.D.GetOkExists("device"); ok {
			tmp := device.(string)
			details.Device = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
			tmp := instanceId.(string)
			details.InstanceId = &tmp
		}
		if isReadOnly, ok := s.D.GetOkExists("is_read_only"); ok {
			tmp := isReadOnly.(bool)
			details.IsReadOnly = &tmp
		}
		if volumeId, ok := s.D.GetOkExists("volume_id"); ok {
			tmp := volumeId.(string)
			details.VolumeId = &tmp
		}
		request.AttachVolumeDetails = details
	case strings.ToLower("paravirtualized"):
		details := oci_core.AttachParavirtualizedVolumeDetails{}
		if device, ok := s.D.GetOkExists("device"); ok {
			tmp := device.(string)
			details.Device = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
			tmp := instanceId.(string)
			details.InstanceId = &tmp
		}
		if isPvEncryptionInTransitEnabled, ok := s.D.GetOkExists("is_pv_encryption_in_transit_enabled"); ok {
			tmp := isPvEncryptionInTransitEnabled.(bool)
			details.IsPvEncryptionInTransitEnabled = &tmp
		}
		if isReadOnly, ok := s.D.GetOkExists("is_read_only"); ok {
			tmp := isReadOnly.(bool)
			details.IsReadOnly = &tmp
		}
		if volumeId, ok := s.D.GetOkExists("volume_id"); ok {
			tmp := volumeId.(string)
			details.VolumeId = &tmp
		}
		request.AttachVolumeDetails = details
	default:
		return fmt.Errorf("unknown attachment_type '%v' was specified", attachmentType)
	}
	return nil
}
