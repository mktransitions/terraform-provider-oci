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

// CertificateBundleWithPrivateKey A certificate bundle, including the private key.
type CertificateBundleWithPrivateKey struct {

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

	// The private key (in PEM format) for the certificate.
	PrivateKeyPem *string `mandatory:"true" json:"privateKeyPem"`

	// The certificate in PEM format.
	CertificatePem *string `mandatory:"false" json:"certificatePem"`

	// The certificate chain (in PEM format) for the certificate bundle.
	CertChainPem *string `mandatory:"false" json:"certChainPem"`

	// The name of the certificate version.
	VersionName *string `mandatory:"false" json:"versionName"`

	RevocationStatus *RevocationStatus `mandatory:"false" json:"revocationStatus"`

	// An optional passphrase for the private key.
	PrivateKeyPemPassphrase *string `mandatory:"false" json:"privateKeyPemPassphrase"`

	// A list of rotation states for the certificate bundle.
	Stages []VersionStageEnum `mandatory:"true" json:"stages"`
}

// GetCertificateId returns CertificateId
func (m CertificateBundleWithPrivateKey) GetCertificateId() *string {
	return m.CertificateId
}

// GetCertificateName returns CertificateName
func (m CertificateBundleWithPrivateKey) GetCertificateName() *string {
	return m.CertificateName
}

// GetVersionNumber returns VersionNumber
func (m CertificateBundleWithPrivateKey) GetVersionNumber() *int64 {
	return m.VersionNumber
}

// GetSerialNumber returns SerialNumber
func (m CertificateBundleWithPrivateKey) GetSerialNumber() *string {
	return m.SerialNumber
}

// GetCertificatePem returns CertificatePem
func (m CertificateBundleWithPrivateKey) GetCertificatePem() *string {
	return m.CertificatePem
}

// GetCertChainPem returns CertChainPem
func (m CertificateBundleWithPrivateKey) GetCertChainPem() *string {
	return m.CertChainPem
}

// GetTimeCreated returns TimeCreated
func (m CertificateBundleWithPrivateKey) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetValidity returns Validity
func (m CertificateBundleWithPrivateKey) GetValidity() *Validity {
	return m.Validity
}

// GetVersionName returns VersionName
func (m CertificateBundleWithPrivateKey) GetVersionName() *string {
	return m.VersionName
}

// GetStages returns Stages
func (m CertificateBundleWithPrivateKey) GetStages() []VersionStageEnum {
	return m.Stages
}

// GetRevocationStatus returns RevocationStatus
func (m CertificateBundleWithPrivateKey) GetRevocationStatus() *RevocationStatus {
	return m.RevocationStatus
}

func (m CertificateBundleWithPrivateKey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateBundleWithPrivateKey) ValidateEnumValue() (bool, error) {
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
func (m CertificateBundleWithPrivateKey) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCertificateBundleWithPrivateKey CertificateBundleWithPrivateKey
	s := struct {
		DiscriminatorParam string `json:"certificateBundleType"`
		MarshalTypeCertificateBundleWithPrivateKey
	}{
		"CERTIFICATE_CONTENT_WITH_PRIVATE_KEY",
		(MarshalTypeCertificateBundleWithPrivateKey)(m),
	}

	return json.Marshal(&s)
}
