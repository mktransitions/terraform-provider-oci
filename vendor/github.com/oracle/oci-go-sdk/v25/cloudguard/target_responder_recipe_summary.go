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

// TargetResponderRecipeSummary Summary of ResponderRecipe
type TargetResponderRecipeSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique identifier for Responder Recipe of which this is an extension
	ResponderRecipeId *string `mandatory:"true" json:"responderRecipeId"`

	// ResponderRecipe Identifier Name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// ResponderRecipe Description
	Description *string `mandatory:"true" json:"description"`

	// Owner of ResponderRecipe
	Owner OwnerTypeEnum `mandatory:"true" json:"owner"`

	// The date and time the target responder recipe was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target responder recipe was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Example.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m TargetResponderRecipeSummary) String() string {
	return common.PointerString(m)
}
