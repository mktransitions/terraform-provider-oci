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

// CertificateAuthorityBundle The contents of the certificate, properties of the certificate (and certificate version), and user-provided contextual metadata for the certificate.
type CertificateAuthorityBundle struct {

	// The OCID of the certificate authority (CA).
	CertificateAuthorityId *string `mandatory:"true" json:"certificateAuthorityId"`

	// The name of the CA.
	CertificateAuthorityName *string `mandatory:"true" json:"certificateAuthorityName"`

	// A unique certificate identifier used in certificate revocation tracking, formatted as octets.
	// Example: `03 AC FC FA CC B3 CB 02 B8 F8 DE F5 85 E7 7B FF`
	SerialNumber *string `mandatory:"true" json:"serialNumber"`

	// The certificate (in PEM format) for this CA version.
	CertificatePem *string `mandatory:"true" json:"certificatePem"`

	// A property indicating when the CA was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The version number of the CA.
	VersionNumber *int64 `mandatory:"true" json:"versionNumber"`

	Validity *Validity `mandatory:"true" json:"validity"`

	// A list of rotation states for this CA.
	Stages []VersionStageEnum `mandatory:"true" json:"stages"`

	// The certificate chain (in PEM format) for this CA version.
	CertChainPem *string `mandatory:"false" json:"certChainPem"`

	// The name of the CA.
	VersionName *string `mandatory:"false" json:"versionName"`

	RevocationStatus *RevocationStatus `mandatory:"false" json:"revocationStatus"`
}

func (m CertificateAuthorityBundle) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateAuthorityBundle) ValidateEnumValue() (bool, error) {
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
