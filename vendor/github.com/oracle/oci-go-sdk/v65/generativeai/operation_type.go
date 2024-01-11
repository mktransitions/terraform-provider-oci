// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service API
//
// **Generative AI Service**
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable LLMs that cover a wide range of use cases for text generation. Use the playground to try out the models out-of-the-box or create and host your own fine-tuned custom models based on your own data on dedicated AI clusters.
//

package generativeai

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateModel              OperationTypeEnum = "CREATE_MODEL"
	OperationTypeDeleteModel              OperationTypeEnum = "DELETE_MODEL"
	OperationTypeMoveModel                OperationTypeEnum = "MOVE_MODEL"
	OperationTypeCreateDedicatedAiCluster OperationTypeEnum = "CREATE_DEDICATED_AI_CLUSTER"
	OperationTypeDeleteDedicatedAiCluster OperationTypeEnum = "DELETE_DEDICATED_AI_CLUSTER"
	OperationTypeUpdateDedicatedAiCluster OperationTypeEnum = "UPDATE_DEDICATED_AI_CLUSTER"
	OperationTypeMoveDedicatedAiCluster   OperationTypeEnum = "MOVE_DEDICATED_AI_CLUSTER"
	OperationTypeCreateEndpoint           OperationTypeEnum = "CREATE_ENDPOINT"
	OperationTypeDeleteEndpoint           OperationTypeEnum = "DELETE_ENDPOINT"
	OperationTypeUpdateEndpoint           OperationTypeEnum = "UPDATE_ENDPOINT"
	OperationTypeMoveEndpoint             OperationTypeEnum = "MOVE_ENDPOINT"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_MODEL":                OperationTypeCreateModel,
	"DELETE_MODEL":                OperationTypeDeleteModel,
	"MOVE_MODEL":                  OperationTypeMoveModel,
	"CREATE_DEDICATED_AI_CLUSTER": OperationTypeCreateDedicatedAiCluster,
	"DELETE_DEDICATED_AI_CLUSTER": OperationTypeDeleteDedicatedAiCluster,
	"UPDATE_DEDICATED_AI_CLUSTER": OperationTypeUpdateDedicatedAiCluster,
	"MOVE_DEDICATED_AI_CLUSTER":   OperationTypeMoveDedicatedAiCluster,
	"CREATE_ENDPOINT":             OperationTypeCreateEndpoint,
	"DELETE_ENDPOINT":             OperationTypeDeleteEndpoint,
	"UPDATE_ENDPOINT":             OperationTypeUpdateEndpoint,
	"MOVE_ENDPOINT":               OperationTypeMoveEndpoint,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_model":                OperationTypeCreateModel,
	"delete_model":                OperationTypeDeleteModel,
	"move_model":                  OperationTypeMoveModel,
	"create_dedicated_ai_cluster": OperationTypeCreateDedicatedAiCluster,
	"delete_dedicated_ai_cluster": OperationTypeDeleteDedicatedAiCluster,
	"update_dedicated_ai_cluster": OperationTypeUpdateDedicatedAiCluster,
	"move_dedicated_ai_cluster":   OperationTypeMoveDedicatedAiCluster,
	"create_endpoint":             OperationTypeCreateEndpoint,
	"delete_endpoint":             OperationTypeDeleteEndpoint,
	"update_endpoint":             OperationTypeUpdateEndpoint,
	"move_endpoint":               OperationTypeMoveEndpoint,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_MODEL",
		"DELETE_MODEL",
		"MOVE_MODEL",
		"CREATE_DEDICATED_AI_CLUSTER",
		"DELETE_DEDICATED_AI_CLUSTER",
		"UPDATE_DEDICATED_AI_CLUSTER",
		"MOVE_DEDICATED_AI_CLUSTER",
		"CREATE_ENDPOINT",
		"DELETE_ENDPOINT",
		"UPDATE_ENDPOINT",
		"MOVE_ENDPOINT",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
