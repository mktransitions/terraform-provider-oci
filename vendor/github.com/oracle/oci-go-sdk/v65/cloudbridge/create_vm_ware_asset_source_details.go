// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateVmWareAssetSourceDetails Description of an asset source.
type CreateVmWareAssetSourceDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment for the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the environment.
	EnvironmentId *string `mandatory:"true" json:"environmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the inventory that will contain created assets.
	InventoryId *string `mandatory:"true" json:"inventoryId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that is going to be used to create assets.
	AssetsCompartmentId *string `mandatory:"true" json:"assetsCompartmentId"`

	// Endpoint for VMware asset discovery and replication in the form of ```https://<host>:<port>/sdk```
	VcenterEndpoint *string `mandatory:"true" json:"vcenterEndpoint"`

	DiscoveryCredentials *AssetSourceCredentials `mandatory:"true" json:"discoveryCredentials"`

	// A user-friendly name for the asset source. Does not have to be unique, and it's mutable.
	// Avoid entering confidential information. The name is generated by the service if it is not
	// explicitly provided.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the discovery schedule that is going to be attached to the created asset.
	DiscoveryScheduleId *string `mandatory:"false" json:"discoveryScheduleId"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	ReplicationCredentials *AssetSourceCredentials `mandatory:"false" json:"replicationCredentials"`

	// Flag indicating whether historical metrics are collected for assets, originating from this asset source.
	AreHistoricalMetricsCollected *bool `mandatory:"false" json:"areHistoricalMetricsCollected"`

	// Flag indicating whether real-time metrics are collected for assets, originating from this asset source.
	AreRealtimeMetricsCollected *bool `mandatory:"false" json:"areRealtimeMetricsCollected"`
}

//GetDisplayName returns DisplayName
func (m CreateVmWareAssetSourceDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m CreateVmWareAssetSourceDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetEnvironmentId returns EnvironmentId
func (m CreateVmWareAssetSourceDetails) GetEnvironmentId() *string {
	return m.EnvironmentId
}

//GetInventoryId returns InventoryId
func (m CreateVmWareAssetSourceDetails) GetInventoryId() *string {
	return m.InventoryId
}

//GetAssetsCompartmentId returns AssetsCompartmentId
func (m CreateVmWareAssetSourceDetails) GetAssetsCompartmentId() *string {
	return m.AssetsCompartmentId
}

//GetDiscoveryScheduleId returns DiscoveryScheduleId
func (m CreateVmWareAssetSourceDetails) GetDiscoveryScheduleId() *string {
	return m.DiscoveryScheduleId
}

//GetFreeformTags returns FreeformTags
func (m CreateVmWareAssetSourceDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateVmWareAssetSourceDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m CreateVmWareAssetSourceDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m CreateVmWareAssetSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVmWareAssetSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateVmWareAssetSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateVmWareAssetSourceDetails CreateVmWareAssetSourceDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateVmWareAssetSourceDetails
	}{
		"VMWARE",
		(MarshalTypeCreateVmWareAssetSourceDetails)(m),
	}

	return json.Marshal(&s)
}
