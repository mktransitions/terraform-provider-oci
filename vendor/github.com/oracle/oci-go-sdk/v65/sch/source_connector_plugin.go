// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.cloud.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SourceConnectorPlugin A connector plugin for fetching data from a source service.
// For configuration instructions, see
// Creating a Connector (https://docs.cloud.oracle.com/iaas/Content/connector-hub/create-service-connector.htm).
type SourceConnectorPlugin struct {

	// The service to be called by the connector plugin.
	// Example: `QueueSource`
	Name *string `mandatory:"true" json:"name"`

	// The date and time when this plugin became available.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2023-09-09T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Gets the specified connector plugin configuration information in OpenAPI specification format.
	Schema *string `mandatory:"false" json:"schema"`

	// The estimated maximum period of time the data will be kept at the source.
	// The duration is specified as a string in ISO 8601 format (P1D for one day or P30D for thrity days).
	MaxRetention *string `mandatory:"false" json:"maxRetention"`

	// The estimated throughput range (LOW, MEDIUM, HIGH).
	EstimatedThroughput EstimatedThroughputEnum `mandatory:"false" json:"estimatedThroughput,omitempty"`

	// The current state of the service connector.
	LifecycleState ConnectorPluginLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetName returns Name
func (m SourceConnectorPlugin) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m SourceConnectorPlugin) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetEstimatedThroughput returns EstimatedThroughput
func (m SourceConnectorPlugin) GetEstimatedThroughput() EstimatedThroughputEnum {
	return m.EstimatedThroughput
}

// GetLifecycleState returns LifecycleState
func (m SourceConnectorPlugin) GetLifecycleState() ConnectorPluginLifecycleStateEnum {
	return m.LifecycleState
}

// GetDisplayName returns DisplayName
func (m SourceConnectorPlugin) GetDisplayName() *string {
	return m.DisplayName
}

// GetSchema returns Schema
func (m SourceConnectorPlugin) GetSchema() *string {
	return m.Schema
}

func (m SourceConnectorPlugin) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SourceConnectorPlugin) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEstimatedThroughputEnum(string(m.EstimatedThroughput)); !ok && m.EstimatedThroughput != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EstimatedThroughput: %s. Supported values are: %s.", m.EstimatedThroughput, strings.Join(GetEstimatedThroughputEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConnectorPluginLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectorPluginLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SourceConnectorPlugin) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSourceConnectorPlugin SourceConnectorPlugin
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeSourceConnectorPlugin
	}{
		"SOURCE",
		(MarshalTypeSourceConnectorPlugin)(m),
	}

	return json.Marshal(&s)
}
