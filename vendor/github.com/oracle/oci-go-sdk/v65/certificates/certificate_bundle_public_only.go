// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Retrieval API
//
// API for retrieving certificates.
//

package certificates

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CertificateBundlePublicOnly A certificate bundle, not including the private key.
type CertificateBundlePublicOnly struct {

	// The OCID of the certificate.
	CertificateId *string `mandatory:"true" json:"certificateId"`

	// The name of the certificate.
	CertificateName *string `mandatory:"true" json:"certificateName"`

	// The version number of the certificate.
	VersionNumber *int64 `mandatory:"true" json:"versionNumber"`

	// A unique certificate identifier used in certificate revocation tracking, formatted as octets.
	// Example: `03 AC FC FA CC B3 CB 02 B8 F8 DE F5 85 E7 7B FF`
	SerialNumber *string `mandatory:"true" json:"serialNumber"`

	// An optional property indicating when the certificate version was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	Validity *Validity `mandatory:"true" json:"validity"`

	// The certificate in PEM format.
	CertificatePem *string `mandatory:"false" json:"certificatePem"`

	// The certificate chain (in PEM format) for the certificate bundle.
	CertChainPem *string `mandatory:"false" json:"certChainPem"`

	// The name of the certificate version.
	VersionName *string `mandatory:"false" json:"versionName"`

	RevocationStatus *RevocationStatus `mandatory:"false" json:"revocationStatus"`

	// A list of rotation states for the certificate bundle.
	Stages []VersionStageEnum `mandatory:"true" json:"stages"`
}

// GetCertificateId returns CertificateId
func (m CertificateBundlePublicOnly) GetCertificateId() *string {
	return m.CertificateId
}

// GetCertificateName returns CertificateName
func (m CertificateBundlePublicOnly) GetCertificateName() *string {
	return m.CertificateName
}

// GetVersionNumber returns VersionNumber
func (m CertificateBundlePublicOnly) GetVersionNumber() *int64 {
	return m.VersionNumber
}

// GetSerialNumber returns SerialNumber
func (m CertificateBundlePublicOnly) GetSerialNumber() *string {
	return m.SerialNumber
}

// GetCertificatePem returns CertificatePem
func (m CertificateBundlePublicOnly) GetCertificatePem() *string {
	return m.CertificatePem
}

// GetCertChainPem returns CertChainPem
func (m CertificateBundlePublicOnly) GetCertChainPem() *string {
	return m.CertChainPem
}

// GetTimeCreated returns TimeCreated
func (m CertificateBundlePublicOnly) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetValidity returns Validity
func (m CertificateBundlePublicOnly) GetValidity() *Validity {
	return m.Validity
}

// GetVersionName returns VersionName
func (m CertificateBundlePublicOnly) GetVersionName() *string {
	return m.VersionName
}

// GetStages returns Stages
func (m CertificateBundlePublicOnly) GetStages() []VersionStageEnum {
	return m.Stages
}

// GetRevocationStatus returns RevocationStatus
func (m CertificateBundlePublicOnly) GetRevocationStatus() *RevocationStatus {
	return m.RevocationStatus
}

func (m CertificateBundlePublicOnly) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateBundlePublicOnly) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m CertificateBundlePublicOnly) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCertificateBundlePublicOnly CertificateBundlePublicOnly
	s := struct {
		DiscriminatorParam string `json:"certificateBundleType"`
		MarshalTypeCertificateBundlePublicOnly
	}{
		"CERTIFICATE_CONTENT_PUBLIC_ONLY",
		(MarshalTypeCertificateBundlePublicOnly)(m),
	}

	return json.Marshal(&s)
}
