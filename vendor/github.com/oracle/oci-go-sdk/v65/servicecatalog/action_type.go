// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"strings"
)

// ActionTypeEnum Enum with underlying type: string
type ActionTypeEnum string

// Set of constants representing the allowable values for ActionTypeEnum
const (
	ActionTypeCreated ActionTypeEnum = "CREATED"
	ActionTypeUpdated ActionTypeEnum = "UPDATED"
	ActionTypeDeleted ActionTypeEnum = "DELETED"
	ActionTypeFailed  ActionTypeEnum = "FAILED"
)

var mappingActionTypeEnum = map[string]ActionTypeEnum{
	"CREATED": ActionTypeCreated,
	"UPDATED": ActionTypeUpdated,
	"DELETED": ActionTypeDeleted,
	"FAILED":  ActionTypeFailed,
}

var mappingActionTypeEnumLowerCase = map[string]ActionTypeEnum{
	"created": ActionTypeCreated,
	"updated": ActionTypeUpdated,
	"deleted": ActionTypeDeleted,
	"failed":  ActionTypeFailed,
}

// GetActionTypeEnumValues Enumerates the set of values for ActionTypeEnum
func GetActionTypeEnumValues() []ActionTypeEnum {
	values := make([]ActionTypeEnum, 0)
	for _, v := range mappingActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetActionTypeEnumStringValues Enumerates the set of values in String for ActionTypeEnum
func GetActionTypeEnumStringValues() []string {
	return []string{
		"CREATED",
		"UPDATED",
		"DELETED",
		"FAILED",
	}
}

// GetMappingActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionTypeEnum(val string) (ActionTypeEnum, bool) {
	enum, ok := mappingActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
