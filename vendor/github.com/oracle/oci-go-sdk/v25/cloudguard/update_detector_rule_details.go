// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v25/common"
)

// UpdateDetectorRuleDetails Details of a Detector Rule to be overriden in Detector Recipe
type UpdateDetectorRuleDetails struct {

	// Enables the control
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The Risk Level
	RiskLevel RiskLevelEnum `mandatory:"true" json:"riskLevel"`

	// Configuration details
	Configurations []DetectorConfiguration `mandatory:"false" json:"configurations"`

	Condition Condition `mandatory:"false" json:"condition"`

	// user defined labels for a detector rule
	Labels []string `mandatory:"false" json:"labels"`
}

func (m UpdateDetectorRuleDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDetectorRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Configurations []DetectorConfiguration `json:"configurations"`
		Condition      condition               `json:"condition"`
		Labels         []string                `json:"labels"`
		IsEnabled      *bool                   `json:"isEnabled"`
		RiskLevel      RiskLevelEnum           `json:"riskLevel"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Configurations = make([]DetectorConfiguration, len(model.Configurations))
	for i, n := range model.Configurations {
		m.Configurations[i] = n
	}

	nn, e = model.Condition.UnmarshalPolymorphicJSON(model.Condition.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Condition = nn.(Condition)
	} else {
		m.Condition = nil
	}

	m.Labels = make([]string, len(model.Labels))
	for i, n := range model.Labels {
		m.Labels[i] = n
	}

	m.IsEnabled = model.IsEnabled

	m.RiskLevel = model.RiskLevel

	return
}
