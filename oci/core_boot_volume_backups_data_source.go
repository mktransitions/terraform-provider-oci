// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v25/core"
)

func init() {
	RegisterDatasource("oci_core_boot_volume_backups", CoreBootVolumeBackupsDataSource())
}

func CoreBootVolumeBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreBootVolumeBackups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"boot_volume_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_boot_volume_backup_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"boot_volume_backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CoreBootVolumeBackupResource()),
			},
		},
	}
}

func readCoreBootVolumeBackups(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumeBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return ReadResource(sync)
}

type CoreBootVolumeBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListBootVolumeBackupsResponse
}

func (s *CoreBootVolumeBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreBootVolumeBackupsDataSourceCrud) Get() error {
	request := oci_core.ListBootVolumeBackupsRequest{}

	if bootVolumeId, ok := s.D.GetOkExists("boot_volume_id"); ok {
		tmp := bootVolumeId.(string)
		request.BootVolumeId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if sourceBootVolumeBackupId, ok := s.D.GetOkExists("source_boot_volume_backup_id"); ok {
		tmp := sourceBootVolumeBackupId.(string)
		request.SourceBootVolumeBackupId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.BootVolumeBackupLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListBootVolumeBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBootVolumeBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreBootVolumeBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bootVolumeBackup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.BootVolumeId != nil {
			bootVolumeBackup["boot_volume_id"] = *r.BootVolumeId
		}

		if r.DefinedTags != nil {
			bootVolumeBackup["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			bootVolumeBackup["display_name"] = *r.DisplayName
		}

		if r.ExpirationTime != nil {
			bootVolumeBackup["expiration_time"] = r.ExpirationTime.String()
		}

		bootVolumeBackup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			bootVolumeBackup["id"] = *r.Id
		}

		if r.ImageId != nil {
			bootVolumeBackup["image_id"] = *r.ImageId
		}

		if r.KmsKeyId != nil {
			bootVolumeBackup["kms_key_id"] = *r.KmsKeyId
		}

		if r.SizeInGBs != nil {
			bootVolumeBackup["size_in_gbs"] = strconv.FormatInt(*r.SizeInGBs, 10)
		}

		if r.SourceBootVolumeBackupId != nil {
			bootVolumeBackup["source_boot_volume_backup_id"] = *r.SourceBootVolumeBackupId
		}

		bootVolumeBackup["source_type"] = r.SourceType

		bootVolumeBackup["state"] = r.LifecycleState

		if r.SystemTags != nil {
			bootVolumeBackup["system_tags"] = systemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			bootVolumeBackup["time_created"] = r.TimeCreated.String()
		}

		if r.TimeRequestReceived != nil {
			bootVolumeBackup["time_request_received"] = r.TimeRequestReceived.String()
		}

		bootVolumeBackup["type"] = r.Type

		if r.UniqueSizeInGBs != nil {
			bootVolumeBackup["unique_size_in_gbs"] = strconv.FormatInt(*r.UniqueSizeInGBs, 10)
		}

		resources = append(resources, bootVolumeBackup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreBootVolumeBackupsDataSource().Schema["boot_volume_backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("boot_volume_backups", resources); err != nil {
		return err
	}

	return nil
}
