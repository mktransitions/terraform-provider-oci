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

// UpdateTargetResponderRecipe The information to be updated in attached Target ResponderRecipe
type UpdateTargetResponderRecipe struct {

	// Identifier for ResponderRecipe.
	TargetResponderRecipeId *string `mandatory:"true" json:"targetResponderRecipeId"`

	// Update responder rules associated with reponder recipe in a target.
	ResponderRules []UpdateTargetRecipeResponderRuleDetails `mandatory:"true" json:"responderRules"`
}

func (m UpdateTargetResponderRecipe) String() string {
	return common.PointerString(m)
}
