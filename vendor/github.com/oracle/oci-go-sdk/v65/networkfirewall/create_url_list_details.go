// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateUrlListDetails The Request for creating the URL List
type CreateUrlListDetails struct {

	// Unique name to identify the group of urls to be used in the policy rules.
	Name *string `mandatory:"true" json:"name"`

	// List of urls.
	Urls []UrlPattern `mandatory:"true" json:"urls"`
}

func (m CreateUrlListDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateUrlListDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateUrlListDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name *string      `json:"name"`
		Urls []urlpattern `json:"urls"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	m.Urls = make([]UrlPattern, len(model.Urls))
	for i, n := range model.Urls {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Urls[i] = nn.(UrlPattern)
		} else {
			m.Urls[i] = nil
		}
	}

	return
}
