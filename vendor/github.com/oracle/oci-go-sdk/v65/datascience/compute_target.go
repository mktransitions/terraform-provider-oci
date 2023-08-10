// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeTarget A compute target.
type ComputeTarget struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute target.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the resource was created in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: 2020-08-06T21:10:29.41Z
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the compute target.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment associated with the compute target.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"true" json:"displayName"`

	ComputeConfigurationDetails ComputeConfigurationDetails `mandatory:"true" json:"computeConfigurationDetails"`

	// The state of the compute target.
	LifecycleState ComputeTargetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A short description of the compute target.
	Description *string `mandatory:"false" json:"description"`

	// Details about the state of the compute target.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ComputeTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingComputeTargetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetComputeTargetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ComputeTarget) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                 *string                           `json:"description"`
		LifecycleDetails            *string                           `json:"lifecycleDetails"`
		FreeformTags                map[string]string                 `json:"freeformTags"`
		DefinedTags                 map[string]map[string]interface{} `json:"definedTags"`
		Id                          *string                           `json:"id"`
		TimeCreated                 *common.SDKTime                   `json:"timeCreated"`
		ProjectId                   *string                           `json:"projectId"`
		CompartmentId               *string                           `json:"compartmentId"`
		DisplayName                 *string                           `json:"displayName"`
		ComputeConfigurationDetails computeconfigurationdetails       `json:"computeConfigurationDetails"`
		LifecycleState              ComputeTargetLifecycleStateEnum   `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.TimeCreated = model.TimeCreated

	m.ProjectId = model.ProjectId

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	nn, e = model.ComputeConfigurationDetails.UnmarshalPolymorphicJSON(model.ComputeConfigurationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ComputeConfigurationDetails = nn.(ComputeConfigurationDetails)
	} else {
		m.ComputeConfigurationDetails = nil
	}

	m.LifecycleState = model.LifecycleState

	return
}
