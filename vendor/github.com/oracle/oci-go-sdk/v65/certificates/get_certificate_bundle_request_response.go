// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package certificates

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetCertificateBundleRequest wrapper for the GetCertificateBundle operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificates/GetCertificateBundle.go.html to see an example of how to use GetCertificateBundleRequest.
type GetCertificateBundleRequest struct {

	// The OCID of the certificate.
	CertificateId *string `mandatory:"true" contributesTo:"path" name:"certificateId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The version number of the certificate. The default value is 0, which means that this query parameter is ignored.
	VersionNumber *int64 `mandatory:"false" contributesTo:"query" name:"versionNumber"`

	// The name of the certificate. (This might be referred to as the name of the certificate version, as every certificate consists of at least one version.) Names are unique across versions of a given certificate.
	CertificateVersionName *string `mandatory:"false" contributesTo:"query" name:"certificateVersionName"`

	// The rotation state of the certificate version.
	Stage GetCertificateBundleStageEnum `mandatory:"false" contributesTo:"query" name:"stage" omitEmpty:"true"`

	// The type of certificate bundle. By default, the private key fields are not returned. When querying for certificate bundles, to return results with certificate contents, the private key in PEM format, and the private key passphrase, specify the value of this parameter as `CERTIFICATE_CONTENT_WITH_PRIVATE_KEY`.
	CertificateBundleType GetCertificateBundleCertificateBundleTypeEnum `mandatory:"false" contributesTo:"query" name:"certificateBundleType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetCertificateBundleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetCertificateBundleRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetCertificateBundleRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetCertificateBundleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetCertificateBundleRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetCertificateBundleStageEnum(string(request.Stage)); !ok && request.Stage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stage: %s. Supported values are: %s.", request.Stage, strings.Join(GetGetCertificateBundleStageEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetCertificateBundleCertificateBundleTypeEnum(string(request.CertificateBundleType)); !ok && request.CertificateBundleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateBundleType: %s. Supported values are: %s.", request.CertificateBundleType, strings.Join(GetGetCertificateBundleCertificateBundleTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetCertificateBundleResponse wrapper for the GetCertificateBundle operation
type GetCertificateBundleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CertificateBundle instance
	CertificateBundle `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCertificateBundleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetCertificateBundleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetCertificateBundleStageEnum Enum with underlying type: string
type GetCertificateBundleStageEnum string

// Set of constants representing the allowable values for GetCertificateBundleStageEnum
const (
	GetCertificateBundleStageCurrent    GetCertificateBundleStageEnum = "CURRENT"
	GetCertificateBundleStagePending    GetCertificateBundleStageEnum = "PENDING"
	GetCertificateBundleStageLatest     GetCertificateBundleStageEnum = "LATEST"
	GetCertificateBundleStagePrevious   GetCertificateBundleStageEnum = "PREVIOUS"
	GetCertificateBundleStageDeprecated GetCertificateBundleStageEnum = "DEPRECATED"
)

var mappingGetCertificateBundleStageEnum = map[string]GetCertificateBundleStageEnum{
	"CURRENT":    GetCertificateBundleStageCurrent,
	"PENDING":    GetCertificateBundleStagePending,
	"LATEST":     GetCertificateBundleStageLatest,
	"PREVIOUS":   GetCertificateBundleStagePrevious,
	"DEPRECATED": GetCertificateBundleStageDeprecated,
}

var mappingGetCertificateBundleStageEnumLowerCase = map[string]GetCertificateBundleStageEnum{
	"current":    GetCertificateBundleStageCurrent,
	"pending":    GetCertificateBundleStagePending,
	"latest":     GetCertificateBundleStageLatest,
	"previous":   GetCertificateBundleStagePrevious,
	"deprecated": GetCertificateBundleStageDeprecated,
}

// GetGetCertificateBundleStageEnumValues Enumerates the set of values for GetCertificateBundleStageEnum
func GetGetCertificateBundleStageEnumValues() []GetCertificateBundleStageEnum {
	values := make([]GetCertificateBundleStageEnum, 0)
	for _, v := range mappingGetCertificateBundleStageEnum {
		values = append(values, v)
	}
	return values
}

// GetGetCertificateBundleStageEnumStringValues Enumerates the set of values in String for GetCertificateBundleStageEnum
func GetGetCertificateBundleStageEnumStringValues() []string {
	return []string{
		"CURRENT",
		"PENDING",
		"LATEST",
		"PREVIOUS",
		"DEPRECATED",
	}
}

// GetMappingGetCertificateBundleStageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetCertificateBundleStageEnum(val string) (GetCertificateBundleStageEnum, bool) {
	enum, ok := mappingGetCertificateBundleStageEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetCertificateBundleCertificateBundleTypeEnum Enum with underlying type: string
type GetCertificateBundleCertificateBundleTypeEnum string

// Set of constants representing the allowable values for GetCertificateBundleCertificateBundleTypeEnum
const (
	GetCertificateBundleCertificateBundleTypePublicOnly     GetCertificateBundleCertificateBundleTypeEnum = "CERTIFICATE_CONTENT_PUBLIC_ONLY"
	GetCertificateBundleCertificateBundleTypeWithPrivateKey GetCertificateBundleCertificateBundleTypeEnum = "CERTIFICATE_CONTENT_WITH_PRIVATE_KEY"
)

var mappingGetCertificateBundleCertificateBundleTypeEnum = map[string]GetCertificateBundleCertificateBundleTypeEnum{
	"CERTIFICATE_CONTENT_PUBLIC_ONLY":      GetCertificateBundleCertificateBundleTypePublicOnly,
	"CERTIFICATE_CONTENT_WITH_PRIVATE_KEY": GetCertificateBundleCertificateBundleTypeWithPrivateKey,
}

var mappingGetCertificateBundleCertificateBundleTypeEnumLowerCase = map[string]GetCertificateBundleCertificateBundleTypeEnum{
	"certificate_content_public_only":      GetCertificateBundleCertificateBundleTypePublicOnly,
	"certificate_content_with_private_key": GetCertificateBundleCertificateBundleTypeWithPrivateKey,
}

// GetGetCertificateBundleCertificateBundleTypeEnumValues Enumerates the set of values for GetCertificateBundleCertificateBundleTypeEnum
func GetGetCertificateBundleCertificateBundleTypeEnumValues() []GetCertificateBundleCertificateBundleTypeEnum {
	values := make([]GetCertificateBundleCertificateBundleTypeEnum, 0)
	for _, v := range mappingGetCertificateBundleCertificateBundleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetCertificateBundleCertificateBundleTypeEnumStringValues Enumerates the set of values in String for GetCertificateBundleCertificateBundleTypeEnum
func GetGetCertificateBundleCertificateBundleTypeEnumStringValues() []string {
	return []string{
		"CERTIFICATE_CONTENT_PUBLIC_ONLY",
		"CERTIFICATE_CONTENT_WITH_PRIVATE_KEY",
	}
}

// GetMappingGetCertificateBundleCertificateBundleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetCertificateBundleCertificateBundleTypeEnum(val string) (GetCertificateBundleCertificateBundleTypeEnum, bool) {
	enum, ok := mappingGetCertificateBundleCertificateBundleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
