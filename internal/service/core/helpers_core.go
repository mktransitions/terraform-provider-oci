// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"net"
	"strings"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

// This applies the differences between the regular schema and the one
// we supply for default resources, and returns the schema for a default resource
func ConvertToDefaultVcnResourceSchema(resourceSchema *schema.Resource) *schema.Resource {
	if resourceSchema == nil {
		return nil
	}

	resourceSchema.Importer = &schema.ResourceImporter{
		State: ImportDefaultVcnResource,
	}

	resourceSchema.Schema["manage_default_resource_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	}

	resourceSchema.Schema["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	}

	delete(resourceSchema.Schema, "vcn_id")

	return resourceSchema
}

func ImportDefaultVcnResource(d *schema.ResourceData, value interface{}) ([]*schema.ResourceData, error) {
	err := d.Set("manage_default_resource_id", d.Id())
	return []*schema.ResourceData{d}, err
}

func LaunchOptionsToMap(obj *oci_core.LaunchOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["boot_volume_type"] = string(obj.BootVolumeType)

	result["firmware"] = string(obj.Firmware)

	result["network_type"] = string(obj.NetworkType)

	result["remote_data_volume_type"] = string(obj.RemoteDataVolumeType)

	if obj.IsConsistentVolumeNamingEnabled != nil {
		result["is_consistent_volume_naming_enabled"] = bool(*obj.IsConsistentVolumeNamingEnabled)
	}

	if obj.IsPvEncryptionInTransitEnabled != nil {
		result["is_pv_encryption_in_transit_enabled"] = bool(*obj.IsPvEncryptionInTransitEnabled)
	}

	return result
}

func getBackupPolicyId(assetId *string, client *oci_core.BlockstorageClient) (*string, error) {
	request := oci_core.GetVolumeBackupPolicyAssetAssignmentRequest{}
	request.AssetId = assetId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := client.GetVolumeBackupPolicyAssetAssignment(context.Background(), request)
	if err != nil {
		return nil, err
	}

	if len(response.Items) > 0 {
		policyAssignment := response.Items[0]
		return policyAssignment.PolicyId, nil
	} else {
		return nil, nil
	}
}

func (s *CoreVolumeBackupResourceCrud) createBlockStorageSourceRegionClient(region string) error {
	if s.SourceRegionClient == nil {
		sourceBlockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*s.Client.ConfigurationProvider())
		if err != nil {
			return fmt.Errorf("cannot Create client for the source region: %v", err)
		}
		err = tf_client.ConfigureClientVar(&sourceBlockStorageClient.BaseClient)
		if err != nil {
			return fmt.Errorf("cannot configure client for the source region: %v", err)
		}
		s.SourceRegionClient = &sourceBlockStorageClient

	}
	s.SourceRegionClient.SetRegion(region)

	if s.SourceRegionWRClient == nil {
		sourceWorkRequestClient, err := oci_work_requests.NewWorkRequestClientWithConfigurationProvider(*s.workRequestClient.ConfigurationProvider())
		if err != nil {
			return fmt.Errorf("cannot Create WorkRequestclient for the source region: %v", err)
		}
		err = tf_client.ConfigureClientVar(&sourceWorkRequestClient.BaseClient)
		if err != nil {
			return fmt.Errorf("cannot configure client for the source region: %v", err)
		}
		s.SourceRegionWRClient = &sourceWorkRequestClient
	}
	s.SourceRegionWRClient.SetRegion(region)

	return nil
}

func (s *CoreVolumeGroupBackupResourceCrud) createBlockStorageSourceRegionClient(region string) error {
	if s.SourceRegionClient == nil {
		sourceBlockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*s.Client.ConfigurationProvider())
		if err != nil {
			return fmt.Errorf("cannot Create client for the source region: %v", err)
		}
		err = tf_client.ConfigureClientVar(&sourceBlockStorageClient.BaseClient)
		if err != nil {
			return fmt.Errorf("cannot configure client for the source region: %v", err)
		}
		s.SourceRegionClient = &sourceBlockStorageClient
	}
	s.SourceRegionClient.SetRegion(region)

	return nil
}

func (s *CoreBootVolumeBackupResourceCrud) createBlockStorageSourceRegionClient(region string) error {
	if s.SourceRegionClient == nil {
		sourceBlockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*s.Client.ConfigurationProvider())
		if err != nil {
			return fmt.Errorf("cannot Create client for the source region: %v", err)
		}
		err = tf_client.ConfigureClientVar(&sourceBlockStorageClient.BaseClient)
		if err != nil {
			return fmt.Errorf("cannot configure client for the source region: %v", err)
		}
		s.SourceRegionClient = &sourceBlockStorageClient
	}
	s.SourceRegionClient.SetRegion(region)

	if s.SourceRegionWRClient == nil {
		sourceWorkRequestClient, err := oci_work_requests.NewWorkRequestClientWithConfigurationProvider(*s.workRequestClient.ConfigurationProvider())
		if err != nil {
			return fmt.Errorf("cannot Create WorkRequestclient for the source region: %v", err)
		}
		err = tf_client.ConfigureClientVar(&sourceWorkRequestClient.BaseClient)
		if err != nil {
			return fmt.Errorf("cannot configure client for the source region: %v", err)
		}
		s.SourceRegionWRClient = &sourceWorkRequestClient
	}
	s.SourceRegionWRClient.SetRegion(region)

	return nil
}

// This before suppression is required because
// `fd00:aaaa:0123::/48` in request comes back as `fd00:aaaa:123::/48` in response
func ipv6CompressionDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	if old == "" || new == "" {
		return false
	}
	oldIp := strings.Split(old, "/")
	newIp := strings.Split(new, "/")
	if len(oldIp) < 2 || len(newIp) < 2 {
		return false
	}
	oldParsedIp := net.ParseIP(oldIp[0])
	oldSubnetMask := oldIp[1]
	newParsedIp := net.ParseIP(newIp[0])
	newSubnetMask := newIp[1]
	return strings.EqualFold(oldParsedIp.String(), newParsedIp.String()) && strings.EqualFold(oldSubnetMask, newSubnetMask)
}

func ipv6Cidr_blocksSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	if key == "ipv6cidr_blocks.#" {
		if old == "" || new == "" {
			return false
		}

		// old and new should represent size of list
		// if there is a difference in size between old and new values, we have a diff
		if old != new {
			return false
		}

		return true
	}

	return ipv6CompressionDiffSuppressFunction(key, old, new, d)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
