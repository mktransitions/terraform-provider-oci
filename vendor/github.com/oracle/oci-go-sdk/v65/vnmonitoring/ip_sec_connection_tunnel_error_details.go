// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IpSecConnectionTunnelErrorDetails Details for an error on an IPSec tunnel.
type IpSecConnectionTunnelErrorDetails struct {

	// Unique ID generated for each error report.
	Id *string `mandatory:"true" json:"id"`

	// Unique code describes the error type.
	ErrorCode *string `mandatory:"true" json:"errorCode"`

	// A detailed description of the error.
	ErrorDescription *string `mandatory:"true" json:"errorDescription"`

	// Resolution for the error.
	Solution *string `mandatory:"true" json:"solution"`

	// Link to more Oracle resources or relevant documentation.
	OciResourcesLink *string `mandatory:"true" json:"ociResourcesLink"`

	// Timestamp when the error occurred.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m IpSecConnectionTunnelErrorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IpSecConnectionTunnelErrorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
