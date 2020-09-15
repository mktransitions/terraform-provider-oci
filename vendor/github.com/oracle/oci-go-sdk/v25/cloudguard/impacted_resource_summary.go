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

// ImpactedResourceSummary Impacted Resource summary Definition.
type ImpactedResourceSummary struct {

	// Unique identifier for finding event
	Id *string `mandatory:"true" json:"id"`

	// Unique id of the Impacted Resource
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Problem Id to which the Impacted Resource is associated
	ProblemId *string `mandatory:"true" json:"problemId"`

	// Compartment Id where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the Impacted Resource
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// Type of the Impacted Resource
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Region where the resource is created
	Region *string `mandatory:"true" json:"region"`

	// Time when the problem was identified
	TimeIdentified *common.SDKTime `mandatory:"true" json:"timeIdentified"`
}

func (m ImpactedResourceSummary) String() string {
	return common.PointerString(m)
}
