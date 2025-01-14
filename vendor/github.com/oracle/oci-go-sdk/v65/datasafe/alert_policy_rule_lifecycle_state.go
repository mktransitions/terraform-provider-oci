// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// AlertPolicyRuleLifecycleStateEnum Enum with underlying type: string
type AlertPolicyRuleLifecycleStateEnum string

// Set of constants representing the allowable values for AlertPolicyRuleLifecycleStateEnum
const (
	AlertPolicyRuleLifecycleStateCreating AlertPolicyRuleLifecycleStateEnum = "CREATING"
	AlertPolicyRuleLifecycleStateUpdating AlertPolicyRuleLifecycleStateEnum = "UPDATING"
	AlertPolicyRuleLifecycleStateActive   AlertPolicyRuleLifecycleStateEnum = "ACTIVE"
	AlertPolicyRuleLifecycleStateDeleting AlertPolicyRuleLifecycleStateEnum = "DELETING"
	AlertPolicyRuleLifecycleStateFailed   AlertPolicyRuleLifecycleStateEnum = "FAILED"
)

var mappingAlertPolicyRuleLifecycleStateEnum = map[string]AlertPolicyRuleLifecycleStateEnum{
	"CREATING": AlertPolicyRuleLifecycleStateCreating,
	"UPDATING": AlertPolicyRuleLifecycleStateUpdating,
	"ACTIVE":   AlertPolicyRuleLifecycleStateActive,
	"DELETING": AlertPolicyRuleLifecycleStateDeleting,
	"FAILED":   AlertPolicyRuleLifecycleStateFailed,
}

var mappingAlertPolicyRuleLifecycleStateEnumLowerCase = map[string]AlertPolicyRuleLifecycleStateEnum{
	"creating": AlertPolicyRuleLifecycleStateCreating,
	"updating": AlertPolicyRuleLifecycleStateUpdating,
	"active":   AlertPolicyRuleLifecycleStateActive,
	"deleting": AlertPolicyRuleLifecycleStateDeleting,
	"failed":   AlertPolicyRuleLifecycleStateFailed,
}

// GetAlertPolicyRuleLifecycleStateEnumValues Enumerates the set of values for AlertPolicyRuleLifecycleStateEnum
func GetAlertPolicyRuleLifecycleStateEnumValues() []AlertPolicyRuleLifecycleStateEnum {
	values := make([]AlertPolicyRuleLifecycleStateEnum, 0)
	for _, v := range mappingAlertPolicyRuleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertPolicyRuleLifecycleStateEnumStringValues Enumerates the set of values in String for AlertPolicyRuleLifecycleStateEnum
func GetAlertPolicyRuleLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
	}
}

// GetMappingAlertPolicyRuleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertPolicyRuleLifecycleStateEnum(val string) (AlertPolicyRuleLifecycleStateEnum, bool) {
	enum, ok := mappingAlertPolicyRuleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
