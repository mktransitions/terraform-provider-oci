// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v25/common"
	"net/http"
)

// ListResponderActivitiesRequest wrapper for the ListResponderActivities operation
type ListResponderActivitiesRequest struct {

	// OCId of the problem.
	ProblemId *string `mandatory:"true" contributesTo:"path" name:"problemId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListResponderActivitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for responderRuleName is ascending. If no value is specified timeCreated is default.
	SortBy ListResponderActivitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResponderActivitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResponderActivitiesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResponderActivitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListResponderActivitiesResponse wrapper for the ListResponderActivities operation
type ListResponderActivitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResponderActivityCollection instances
	ResponderActivityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResponderActivitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResponderActivitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResponderActivitiesSortOrderEnum Enum with underlying type: string
type ListResponderActivitiesSortOrderEnum string

// Set of constants representing the allowable values for ListResponderActivitiesSortOrderEnum
const (
	ListResponderActivitiesSortOrderAsc  ListResponderActivitiesSortOrderEnum = "ASC"
	ListResponderActivitiesSortOrderDesc ListResponderActivitiesSortOrderEnum = "DESC"
)

var mappingListResponderActivitiesSortOrder = map[string]ListResponderActivitiesSortOrderEnum{
	"ASC":  ListResponderActivitiesSortOrderAsc,
	"DESC": ListResponderActivitiesSortOrderDesc,
}

// GetListResponderActivitiesSortOrderEnumValues Enumerates the set of values for ListResponderActivitiesSortOrderEnum
func GetListResponderActivitiesSortOrderEnumValues() []ListResponderActivitiesSortOrderEnum {
	values := make([]ListResponderActivitiesSortOrderEnum, 0)
	for _, v := range mappingListResponderActivitiesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListResponderActivitiesSortByEnum Enum with underlying type: string
type ListResponderActivitiesSortByEnum string

// Set of constants representing the allowable values for ListResponderActivitiesSortByEnum
const (
	ListResponderActivitiesSortByTimecreated       ListResponderActivitiesSortByEnum = "timeCreated"
	ListResponderActivitiesSortByResponderrulename ListResponderActivitiesSortByEnum = "responderRuleName"
)

var mappingListResponderActivitiesSortBy = map[string]ListResponderActivitiesSortByEnum{
	"timeCreated":       ListResponderActivitiesSortByTimecreated,
	"responderRuleName": ListResponderActivitiesSortByResponderrulename,
}

// GetListResponderActivitiesSortByEnumValues Enumerates the set of values for ListResponderActivitiesSortByEnum
func GetListResponderActivitiesSortByEnumValues() []ListResponderActivitiesSortByEnum {
	values := make([]ListResponderActivitiesSortByEnum, 0)
	for _, v := range mappingListResponderActivitiesSortBy {
		values = append(values, v)
	}
	return values
}
