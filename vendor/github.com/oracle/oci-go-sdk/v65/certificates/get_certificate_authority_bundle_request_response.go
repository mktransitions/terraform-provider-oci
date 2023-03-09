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

// GetCertificateAuthorityBundleRequest wrapper for the GetCertificateAuthorityBundle operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificates/GetCertificateAuthorityBundle.go.html to see an example of how to use GetCertificateAuthorityBundleRequest.
type GetCertificateAuthorityBundleRequest struct {

	// The OCID of the certificate authority (CA).
	CertificateAuthorityId *string `mandatory:"true" contributesTo:"path" name:"certificateAuthorityId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The version number of the certificate authority (CA).
	VersionNumber *int64 `mandatory:"false" contributesTo:"query" name:"versionNumber"`

	// The name of the certificate authority (CA). (This might be referred to as the name of the CA version, as every CA consists of at least one version.) Names are unique across versions of a given CA.
	CertificateAuthorityVersionName *string `mandatory:"false" contributesTo:"query" name:"certificateAuthorityVersionName"`

	// The rotation state of the certificate version.
	Stage GetCertificateAuthorityBundleStageEnum `mandatory:"false" contributesTo:"query" name:"stage" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetCertificateAuthorityBundleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetCertificateAuthorityBundleRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetCertificateAuthorityBundleRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetCertificateAuthorityBundleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetCertificateAuthorityBundleRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetCertificateAuthorityBundleStageEnum(string(request.Stage)); !ok && request.Stage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stage: %s. Supported values are: %s.", request.Stage, strings.Join(GetGetCertificateAuthorityBundleStageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetCertificateAuthorityBundleResponse wrapper for the GetCertificateAuthorityBundle operation
type GetCertificateAuthorityBundleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CertificateAuthorityBundle instance
	CertificateAuthorityBundle `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCertificateAuthorityBundleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetCertificateAuthorityBundleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetCertificateAuthorityBundleStageEnum Enum with underlying type: string
type GetCertificateAuthorityBundleStageEnum string

// Set of constants representing the allowable values for GetCertificateAuthorityBundleStageEnum
const (
	GetCertificateAuthorityBundleStageCurrent    GetCertificateAuthorityBundleStageEnum = "CURRENT"
	GetCertificateAuthorityBundleStagePending    GetCertificateAuthorityBundleStageEnum = "PENDING"
	GetCertificateAuthorityBundleStageLatest     GetCertificateAuthorityBundleStageEnum = "LATEST"
	GetCertificateAuthorityBundleStagePrevious   GetCertificateAuthorityBundleStageEnum = "PREVIOUS"
	GetCertificateAuthorityBundleStageDeprecated GetCertificateAuthorityBundleStageEnum = "DEPRECATED"
)

var mappingGetCertificateAuthorityBundleStageEnum = map[string]GetCertificateAuthorityBundleStageEnum{
	"CURRENT":    GetCertificateAuthorityBundleStageCurrent,
	"PENDING":    GetCertificateAuthorityBundleStagePending,
	"LATEST":     GetCertificateAuthorityBundleStageLatest,
	"PREVIOUS":   GetCertificateAuthorityBundleStagePrevious,
	"DEPRECATED": GetCertificateAuthorityBundleStageDeprecated,
}

var mappingGetCertificateAuthorityBundleStageEnumLowerCase = map[string]GetCertificateAuthorityBundleStageEnum{
	"current":    GetCertificateAuthorityBundleStageCurrent,
	"pending":    GetCertificateAuthorityBundleStagePending,
	"latest":     GetCertificateAuthorityBundleStageLatest,
	"previous":   GetCertificateAuthorityBundleStagePrevious,
	"deprecated": GetCertificateAuthorityBundleStageDeprecated,
}

// GetGetCertificateAuthorityBundleStageEnumValues Enumerates the set of values for GetCertificateAuthorityBundleStageEnum
func GetGetCertificateAuthorityBundleStageEnumValues() []GetCertificateAuthorityBundleStageEnum {
	values := make([]GetCertificateAuthorityBundleStageEnum, 0)
	for _, v := range mappingGetCertificateAuthorityBundleStageEnum {
		values = append(values, v)
	}
	return values
}

// GetGetCertificateAuthorityBundleStageEnumStringValues Enumerates the set of values in String for GetCertificateAuthorityBundleStageEnum
func GetGetCertificateAuthorityBundleStageEnumStringValues() []string {
	return []string{
		"CURRENT",
		"PENDING",
		"LATEST",
		"PREVIOUS",
		"DEPRECATED",
	}
}

// GetMappingGetCertificateAuthorityBundleStageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetCertificateAuthorityBundleStageEnum(val string) (GetCertificateAuthorityBundleStageEnum, bool) {
	enum, ok := mappingGetCertificateAuthorityBundleStageEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
