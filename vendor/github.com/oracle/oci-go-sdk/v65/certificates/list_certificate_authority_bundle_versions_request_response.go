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

// ListCertificateAuthorityBundleVersionsRequest wrapper for the ListCertificateAuthorityBundleVersions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificates/ListCertificateAuthorityBundleVersions.go.html to see an example of how to use ListCertificateAuthorityBundleVersionsRequest.
type ListCertificateAuthorityBundleVersionsRequest struct {

	// The OCID of the certificate authority (CA).
	CertificateAuthorityId *string `mandatory:"true" contributesTo:"path" name:"certificateAuthorityId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can specify only one sort order. The default
	// order for `VERSION_NUMBER` is ascending.
	SortBy ListCertificateAuthorityBundleVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCertificateAuthorityBundleVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCertificateAuthorityBundleVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCertificateAuthorityBundleVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCertificateAuthorityBundleVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCertificateAuthorityBundleVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCertificateAuthorityBundleVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCertificateAuthorityBundleVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCertificateAuthorityBundleVersionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCertificateAuthorityBundleVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCertificateAuthorityBundleVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCertificateAuthorityBundleVersionsResponse wrapper for the ListCertificateAuthorityBundleVersions operation
type ListCertificateAuthorityBundleVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CertificateAuthorityBundleVersionCollection instance
	CertificateAuthorityBundleVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCertificateAuthorityBundleVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCertificateAuthorityBundleVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCertificateAuthorityBundleVersionsSortByEnum Enum with underlying type: string
type ListCertificateAuthorityBundleVersionsSortByEnum string

// Set of constants representing the allowable values for ListCertificateAuthorityBundleVersionsSortByEnum
const (
	ListCertificateAuthorityBundleVersionsSortByVersionNumber ListCertificateAuthorityBundleVersionsSortByEnum = "VERSION_NUMBER"
)

var mappingListCertificateAuthorityBundleVersionsSortByEnum = map[string]ListCertificateAuthorityBundleVersionsSortByEnum{
	"VERSION_NUMBER": ListCertificateAuthorityBundleVersionsSortByVersionNumber,
}

var mappingListCertificateAuthorityBundleVersionsSortByEnumLowerCase = map[string]ListCertificateAuthorityBundleVersionsSortByEnum{
	"version_number": ListCertificateAuthorityBundleVersionsSortByVersionNumber,
}

// GetListCertificateAuthorityBundleVersionsSortByEnumValues Enumerates the set of values for ListCertificateAuthorityBundleVersionsSortByEnum
func GetListCertificateAuthorityBundleVersionsSortByEnumValues() []ListCertificateAuthorityBundleVersionsSortByEnum {
	values := make([]ListCertificateAuthorityBundleVersionsSortByEnum, 0)
	for _, v := range mappingListCertificateAuthorityBundleVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificateAuthorityBundleVersionsSortByEnumStringValues Enumerates the set of values in String for ListCertificateAuthorityBundleVersionsSortByEnum
func GetListCertificateAuthorityBundleVersionsSortByEnumStringValues() []string {
	return []string{
		"VERSION_NUMBER",
	}
}

// GetMappingListCertificateAuthorityBundleVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificateAuthorityBundleVersionsSortByEnum(val string) (ListCertificateAuthorityBundleVersionsSortByEnum, bool) {
	enum, ok := mappingListCertificateAuthorityBundleVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCertificateAuthorityBundleVersionsSortOrderEnum Enum with underlying type: string
type ListCertificateAuthorityBundleVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListCertificateAuthorityBundleVersionsSortOrderEnum
const (
	ListCertificateAuthorityBundleVersionsSortOrderAsc  ListCertificateAuthorityBundleVersionsSortOrderEnum = "ASC"
	ListCertificateAuthorityBundleVersionsSortOrderDesc ListCertificateAuthorityBundleVersionsSortOrderEnum = "DESC"
)

var mappingListCertificateAuthorityBundleVersionsSortOrderEnum = map[string]ListCertificateAuthorityBundleVersionsSortOrderEnum{
	"ASC":  ListCertificateAuthorityBundleVersionsSortOrderAsc,
	"DESC": ListCertificateAuthorityBundleVersionsSortOrderDesc,
}

var mappingListCertificateAuthorityBundleVersionsSortOrderEnumLowerCase = map[string]ListCertificateAuthorityBundleVersionsSortOrderEnum{
	"asc":  ListCertificateAuthorityBundleVersionsSortOrderAsc,
	"desc": ListCertificateAuthorityBundleVersionsSortOrderDesc,
}

// GetListCertificateAuthorityBundleVersionsSortOrderEnumValues Enumerates the set of values for ListCertificateAuthorityBundleVersionsSortOrderEnum
func GetListCertificateAuthorityBundleVersionsSortOrderEnumValues() []ListCertificateAuthorityBundleVersionsSortOrderEnum {
	values := make([]ListCertificateAuthorityBundleVersionsSortOrderEnum, 0)
	for _, v := range mappingListCertificateAuthorityBundleVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificateAuthorityBundleVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListCertificateAuthorityBundleVersionsSortOrderEnum
func GetListCertificateAuthorityBundleVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCertificateAuthorityBundleVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificateAuthorityBundleVersionsSortOrderEnum(val string) (ListCertificateAuthorityBundleVersionsSortOrderEnum, bool) {
	enum, ok := mappingListCertificateAuthorityBundleVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
