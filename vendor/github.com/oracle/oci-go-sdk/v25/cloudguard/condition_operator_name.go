// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ConditionOperatorNameEnum Enum with underlying type: string
type ConditionOperatorNameEnum string

// Set of constants representing the allowable values for ConditionOperatorNameEnum
const (
	ConditionOperatorNameAnd         ConditionOperatorNameEnum = "AND"
	ConditionOperatorNameOr          ConditionOperatorNameEnum = "OR"
	ConditionOperatorNameIn          ConditionOperatorNameEnum = "IN"
	ConditionOperatorNameNotIn       ConditionOperatorNameEnum = "NOT_IN"
	ConditionOperatorNameEquals      ConditionOperatorNameEnum = "EQUALS"
	ConditionOperatorNameNotEquals   ConditionOperatorNameEnum = "NOT_EQUALS"
	ConditionOperatorNameLessThan    ConditionOperatorNameEnum = "LESS_THAN"
	ConditionOperatorNameGreaterThan ConditionOperatorNameEnum = "GREATER_THAN"
	ConditionOperatorNameRange       ConditionOperatorNameEnum = "RANGE"
)

var mappingConditionOperatorName = map[string]ConditionOperatorNameEnum{
	"AND":          ConditionOperatorNameAnd,
	"OR":           ConditionOperatorNameOr,
	"IN":           ConditionOperatorNameIn,
	"NOT_IN":       ConditionOperatorNameNotIn,
	"EQUALS":       ConditionOperatorNameEquals,
	"NOT_EQUALS":   ConditionOperatorNameNotEquals,
	"LESS_THAN":    ConditionOperatorNameLessThan,
	"GREATER_THAN": ConditionOperatorNameGreaterThan,
	"RANGE":        ConditionOperatorNameRange,
}

// GetConditionOperatorNameEnumValues Enumerates the set of values for ConditionOperatorNameEnum
func GetConditionOperatorNameEnumValues() []ConditionOperatorNameEnum {
	values := make([]ConditionOperatorNameEnum, 0)
	for _, v := range mappingConditionOperatorName {
		values = append(values, v)
	}
	return values
}
