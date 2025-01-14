// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Associations Associations for the runbook.
type Associations struct {

	// A set of tasks to execute in the runbook.
	Tasks []Task `mandatory:"true" json:"tasks"`

	// The groups of the runbook.
	Groups []Group `mandatory:"true" json:"groups"`

	ExecutionWorkflowDetails *ExecutionWorkflowDetails `mandatory:"true" json:"executionWorkflowDetails"`

	RollbackWorkflowDetails *RollbackWorkflowDetails `mandatory:"false" json:"rollbackWorkflowDetails"`

	// The version of the runbook.
	Version *string `mandatory:"false" json:"version"`
}

func (m Associations) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Associations) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
