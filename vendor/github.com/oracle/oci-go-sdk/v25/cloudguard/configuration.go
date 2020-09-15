// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// Configuration Cloud Guard configuration details of a tenancy.
type Configuration struct {

	// The reporting region value
	ReportingRegion *string `mandatory:"true" json:"reportingRegion"`

	// Status of Cloud Guard Tenant
	Status CloudGuardStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Identifies if Oracle managed resources were created by customers
	SelfManageResources *bool `mandatory:"false" json:"selfManageResources"`
}

func (m Configuration) String() string {
	return common.PointerString(m)
}
