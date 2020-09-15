// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ProblemLifecycleStateEnum Enum with underlying type: string
type ProblemLifecycleStateEnum string

// Set of constants representing the allowable values for ProblemLifecycleStateEnum
const (
	ProblemLifecycleStateActive   ProblemLifecycleStateEnum = "ACTIVE"
	ProblemLifecycleStateInactive ProblemLifecycleStateEnum = "INACTIVE"
)

var mappingProblemLifecycleState = map[string]ProblemLifecycleStateEnum{
	"ACTIVE":   ProblemLifecycleStateActive,
	"INACTIVE": ProblemLifecycleStateInactive,
}

// GetProblemLifecycleStateEnumValues Enumerates the set of values for ProblemLifecycleStateEnum
func GetProblemLifecycleStateEnumValues() []ProblemLifecycleStateEnum {
	values := make([]ProblemLifecycleStateEnum, 0)
	for _, v := range mappingProblemLifecycleState {
		values = append(values, v)
	}
	return values
}
