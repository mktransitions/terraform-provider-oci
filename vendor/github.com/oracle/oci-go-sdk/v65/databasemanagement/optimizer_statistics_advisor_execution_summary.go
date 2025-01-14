// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OptimizerStatisticsAdvisorExecutionSummary The summary of the Optimizer Statistics Advisor execution.
type OptimizerStatisticsAdvisorExecutionSummary struct {

	// The name of the Optimizer Statistics Advisor task.
	TaskName *string `mandatory:"true" json:"taskName"`

	// The name of the Optimizer Statistics Advisor execution.
	ExecutionName *string `mandatory:"true" json:"executionName"`

	// The start time of the time range to retrieve the Optimizer Statistics Advisor execution of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	TimeStart *common.SDKTime `mandatory:"true" json:"timeStart"`

	// The end time of the time range to retrieve the Optimizer Statistics Advisor execution of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	TimeEnd *common.SDKTime `mandatory:"true" json:"timeEnd"`

	// The status of the Optimizer Statistics Advisor execution.
	Status OptimizerStatisticsAdvisorExecutionSummaryStatusEnum `mandatory:"true" json:"status"`

	// The Optimizer Statistics Advisor execution status message, if any.
	StatusMessage *string `mandatory:"false" json:"statusMessage"`

	// The errors in the Optimizer Statistics Advisor execution, if any.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// The number of findings generated by the Optimizer Statistics Advisor execution.
	Findings *int `mandatory:"false" json:"findings"`
}

func (m OptimizerStatisticsAdvisorExecutionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OptimizerStatisticsAdvisorExecutionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOptimizerStatisticsAdvisorExecutionSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOptimizerStatisticsAdvisorExecutionSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OptimizerStatisticsAdvisorExecutionSummaryStatusEnum Enum with underlying type: string
type OptimizerStatisticsAdvisorExecutionSummaryStatusEnum string

// Set of constants representing the allowable values for OptimizerStatisticsAdvisorExecutionSummaryStatusEnum
const (
	OptimizerStatisticsAdvisorExecutionSummaryStatusExecuting   OptimizerStatisticsAdvisorExecutionSummaryStatusEnum = "EXECUTING"
	OptimizerStatisticsAdvisorExecutionSummaryStatusCompleted   OptimizerStatisticsAdvisorExecutionSummaryStatusEnum = "COMPLETED"
	OptimizerStatisticsAdvisorExecutionSummaryStatusInterrupted OptimizerStatisticsAdvisorExecutionSummaryStatusEnum = "INTERRUPTED"
	OptimizerStatisticsAdvisorExecutionSummaryStatusCancelled   OptimizerStatisticsAdvisorExecutionSummaryStatusEnum = "CANCELLED"
	OptimizerStatisticsAdvisorExecutionSummaryStatusFatalError  OptimizerStatisticsAdvisorExecutionSummaryStatusEnum = "FATAL_ERROR"
)

var mappingOptimizerStatisticsAdvisorExecutionSummaryStatusEnum = map[string]OptimizerStatisticsAdvisorExecutionSummaryStatusEnum{
	"EXECUTING":   OptimizerStatisticsAdvisorExecutionSummaryStatusExecuting,
	"COMPLETED":   OptimizerStatisticsAdvisorExecutionSummaryStatusCompleted,
	"INTERRUPTED": OptimizerStatisticsAdvisorExecutionSummaryStatusInterrupted,
	"CANCELLED":   OptimizerStatisticsAdvisorExecutionSummaryStatusCancelled,
	"FATAL_ERROR": OptimizerStatisticsAdvisorExecutionSummaryStatusFatalError,
}

var mappingOptimizerStatisticsAdvisorExecutionSummaryStatusEnumLowerCase = map[string]OptimizerStatisticsAdvisorExecutionSummaryStatusEnum{
	"executing":   OptimizerStatisticsAdvisorExecutionSummaryStatusExecuting,
	"completed":   OptimizerStatisticsAdvisorExecutionSummaryStatusCompleted,
	"interrupted": OptimizerStatisticsAdvisorExecutionSummaryStatusInterrupted,
	"cancelled":   OptimizerStatisticsAdvisorExecutionSummaryStatusCancelled,
	"fatal_error": OptimizerStatisticsAdvisorExecutionSummaryStatusFatalError,
}

// GetOptimizerStatisticsAdvisorExecutionSummaryStatusEnumValues Enumerates the set of values for OptimizerStatisticsAdvisorExecutionSummaryStatusEnum
func GetOptimizerStatisticsAdvisorExecutionSummaryStatusEnumValues() []OptimizerStatisticsAdvisorExecutionSummaryStatusEnum {
	values := make([]OptimizerStatisticsAdvisorExecutionSummaryStatusEnum, 0)
	for _, v := range mappingOptimizerStatisticsAdvisorExecutionSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOptimizerStatisticsAdvisorExecutionSummaryStatusEnumStringValues Enumerates the set of values in String for OptimizerStatisticsAdvisorExecutionSummaryStatusEnum
func GetOptimizerStatisticsAdvisorExecutionSummaryStatusEnumStringValues() []string {
	return []string{
		"EXECUTING",
		"COMPLETED",
		"INTERRUPTED",
		"CANCELLED",
		"FATAL_ERROR",
	}
}

// GetMappingOptimizerStatisticsAdvisorExecutionSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOptimizerStatisticsAdvisorExecutionSummaryStatusEnum(val string) (OptimizerStatisticsAdvisorExecutionSummaryStatusEnum, bool) {
	enum, ok := mappingOptimizerStatisticsAdvisorExecutionSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
