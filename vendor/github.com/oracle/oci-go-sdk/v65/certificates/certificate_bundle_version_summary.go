// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Retrieval API
//
// API for retrieving certificates.
//

package certificates

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CertificateBundleVersionSummary The properties of the certificate bundle. Certificate bundle version summary objects do not include the actual contents of the certificate.
type CertificateBundleVersionSummary struct {

	// The OCID of the certificate.
	CertificateId *string `mandatory:"true" json:"certificateId"`

	// The name of the certificate.
	CertificateName *string `mandatory:"true" json:"certificateName"`

	// The version number of the certificate.
	VersionNumber *int64 `mandatory:"true" json:"versionNumber"`

	// An optional property indicating when the certificate version was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A list of rotation states for this certificate bundle version.
	Stages []VersionStageEnum `mandatory:"true" json:"stages"`

	// A unique certificate identifier used in certificate revocation tracking, formatted as octets.
	// Example: `03 AC FC FA CC B3 CB 02 B8 F8 DE F5 85 E7 7B FF`
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	// The name of the certificate version.
	VersionName *string `mandatory:"false" json:"versionName"`

	Validity *Validity `mandatory:"false" json:"validity"`

	// An optional property indicating when to delete the certificate version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	RevocationStatus *RevocationStatus `mandatory:"false" json:"revocationStatus"`
}

func (m CertificateBundleVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateBundleVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Stages {
		if _, ok := GetMappingVersionStageEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stages: %s. Supported values are: %s.", val, strings.Join(GetVersionStageEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
