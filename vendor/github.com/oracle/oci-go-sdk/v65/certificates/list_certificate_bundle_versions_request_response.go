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

// ListCertificateBundleVersionsRequest wrapper for the ListCertificateBundleVersions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificates/ListCertificateBundleVersions.go.html to see an example of how to use ListCertificateBundleVersionsRequest.
type ListCertificateBundleVersionsRequest struct {

	// The OCID of the certificate.
	CertificateId *string `mandatory:"true" contributesTo:"path" name:"certificateId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can specify only one sort order. The default
	// order for `VERSION_NUMBER` is ascending.
	SortBy ListCertificateBundleVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCertificateBundleVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCertificateBundleVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCertificateBundleVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCertificateBundleVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCertificateBundleVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCertificateBundleVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCertificateBundleVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCertificateBundleVersionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCertificateBundleVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCertificateBundleVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCertificateBundleVersionsResponse wrapper for the ListCertificateBundleVersions operation
type ListCertificateBundleVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CertificateBundleVersionCollection instance
	CertificateBundleVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCertificateBundleVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCertificateBundleVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCertificateBundleVersionsSortByEnum Enum with underlying type: string
type ListCertificateBundleVersionsSortByEnum string

// Set of constants representing the allowable values for ListCertificateBundleVersionsSortByEnum
const (
	ListCertificateBundleVersionsSortByVersionNumber ListCertificateBundleVersionsSortByEnum = "VERSION_NUMBER"
)

var mappingListCertificateBundleVersionsSortByEnum = map[string]ListCertificateBundleVersionsSortByEnum{
	"VERSION_NUMBER": ListCertificateBundleVersionsSortByVersionNumber,
}

var mappingListCertificateBundleVersionsSortByEnumLowerCase = map[string]ListCertificateBundleVersionsSortByEnum{
	"version_number": ListCertificateBundleVersionsSortByVersionNumber,
}

// GetListCertificateBundleVersionsSortByEnumValues Enumerates the set of values for ListCertificateBundleVersionsSortByEnum
func GetListCertificateBundleVersionsSortByEnumValues() []ListCertificateBundleVersionsSortByEnum {
	values := make([]ListCertificateBundleVersionsSortByEnum, 0)
	for _, v := range mappingListCertificateBundleVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificateBundleVersionsSortByEnumStringValues Enumerates the set of values in String for ListCertificateBundleVersionsSortByEnum
func GetListCertificateBundleVersionsSortByEnumStringValues() []string {
	return []string{
		"VERSION_NUMBER",
	}
}

// GetMappingListCertificateBundleVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificateBundleVersionsSortByEnum(val string) (ListCertificateBundleVersionsSortByEnum, bool) {
	enum, ok := mappingListCertificateBundleVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCertificateBundleVersionsSortOrderEnum Enum with underlying type: string
type ListCertificateBundleVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListCertificateBundleVersionsSortOrderEnum
const (
	ListCertificateBundleVersionsSortOrderAsc  ListCertificateBundleVersionsSortOrderEnum = "ASC"
	ListCertificateBundleVersionsSortOrderDesc ListCertificateBundleVersionsSortOrderEnum = "DESC"
)

var mappingListCertificateBundleVersionsSortOrderEnum = map[string]ListCertificateBundleVersionsSortOrderEnum{
	"ASC":  ListCertificateBundleVersionsSortOrderAsc,
	"DESC": ListCertificateBundleVersionsSortOrderDesc,
}

var mappingListCertificateBundleVersionsSortOrderEnumLowerCase = map[string]ListCertificateBundleVersionsSortOrderEnum{
	"asc":  ListCertificateBundleVersionsSortOrderAsc,
	"desc": ListCertificateBundleVersionsSortOrderDesc,
}

// GetListCertificateBundleVersionsSortOrderEnumValues Enumerates the set of values for ListCertificateBundleVersionsSortOrderEnum
func GetListCertificateBundleVersionsSortOrderEnumValues() []ListCertificateBundleVersionsSortOrderEnum {
	values := make([]ListCertificateBundleVersionsSortOrderEnum, 0)
	for _, v := range mappingListCertificateBundleVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificateBundleVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListCertificateBundleVersionsSortOrderEnum
func GetListCertificateBundleVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCertificateBundleVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificateBundleVersionsSortOrderEnum(val string) (ListCertificateBundleVersionsSortOrderEnum, bool) {
	enum, ok := mappingListCertificateBundleVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
