// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v25/common"
	"net/http"
)

// ListResponderRulesRequest wrapper for the ListResponderRules operation
type ListResponderRulesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListResponderRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListResponderRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListResponderRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResponderRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResponderRulesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResponderRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListResponderRulesResponse wrapper for the ListResponderRules operation
type ListResponderRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResponderRuleCollection instances
	ResponderRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResponderRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResponderRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResponderRulesLifecycleStateEnum Enum with underlying type: string
type ListResponderRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListResponderRulesLifecycleStateEnum
const (
	ListResponderRulesLifecycleStateCreating ListResponderRulesLifecycleStateEnum = "CREATING"
	ListResponderRulesLifecycleStateUpdating ListResponderRulesLifecycleStateEnum = "UPDATING"
	ListResponderRulesLifecycleStateActive   ListResponderRulesLifecycleStateEnum = "ACTIVE"
	ListResponderRulesLifecycleStateInactive ListResponderRulesLifecycleStateEnum = "INACTIVE"
	ListResponderRulesLifecycleStateDeleting ListResponderRulesLifecycleStateEnum = "DELETING"
	ListResponderRulesLifecycleStateDeleted  ListResponderRulesLifecycleStateEnum = "DELETED"
	ListResponderRulesLifecycleStateFailed   ListResponderRulesLifecycleStateEnum = "FAILED"
)

var mappingListResponderRulesLifecycleState = map[string]ListResponderRulesLifecycleStateEnum{
	"CREATING": ListResponderRulesLifecycleStateCreating,
	"UPDATING": ListResponderRulesLifecycleStateUpdating,
	"ACTIVE":   ListResponderRulesLifecycleStateActive,
	"INACTIVE": ListResponderRulesLifecycleStateInactive,
	"DELETING": ListResponderRulesLifecycleStateDeleting,
	"DELETED":  ListResponderRulesLifecycleStateDeleted,
	"FAILED":   ListResponderRulesLifecycleStateFailed,
}

// GetListResponderRulesLifecycleStateEnumValues Enumerates the set of values for ListResponderRulesLifecycleStateEnum
func GetListResponderRulesLifecycleStateEnumValues() []ListResponderRulesLifecycleStateEnum {
	values := make([]ListResponderRulesLifecycleStateEnum, 0)
	for _, v := range mappingListResponderRulesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListResponderRulesSortOrderEnum Enum with underlying type: string
type ListResponderRulesSortOrderEnum string

// Set of constants representing the allowable values for ListResponderRulesSortOrderEnum
const (
	ListResponderRulesSortOrderAsc  ListResponderRulesSortOrderEnum = "ASC"
	ListResponderRulesSortOrderDesc ListResponderRulesSortOrderEnum = "DESC"
)

var mappingListResponderRulesSortOrder = map[string]ListResponderRulesSortOrderEnum{
	"ASC":  ListResponderRulesSortOrderAsc,
	"DESC": ListResponderRulesSortOrderDesc,
}

// GetListResponderRulesSortOrderEnumValues Enumerates the set of values for ListResponderRulesSortOrderEnum
func GetListResponderRulesSortOrderEnumValues() []ListResponderRulesSortOrderEnum {
	values := make([]ListResponderRulesSortOrderEnum, 0)
	for _, v := range mappingListResponderRulesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListResponderRulesSortByEnum Enum with underlying type: string
type ListResponderRulesSortByEnum string

// Set of constants representing the allowable values for ListResponderRulesSortByEnum
const (
	ListResponderRulesSortByTimecreated ListResponderRulesSortByEnum = "timeCreated"
	ListResponderRulesSortByDisplayname ListResponderRulesSortByEnum = "displayName"
)

var mappingListResponderRulesSortBy = map[string]ListResponderRulesSortByEnum{
	"timeCreated": ListResponderRulesSortByTimecreated,
	"displayName": ListResponderRulesSortByDisplayname,
}

// GetListResponderRulesSortByEnumValues Enumerates the set of values for ListResponderRulesSortByEnum
func GetListResponderRulesSortByEnumValues() []ListResponderRulesSortByEnum {
	values := make([]ListResponderRulesSortByEnum, 0)
	for _, v := range mappingListResponderRulesSortBy {
		values = append(values, v)
	}
	return values
}
