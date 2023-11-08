// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"bytes"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// ImportCustomContentRequest wrapper for the ImportCustomContent operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ImportCustomContent.go.html to see an example of how to use ImportCustomContentRequest.
type ImportCustomContentRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The file to upload which contains the custom content.
	ImportCustomContentFileBody io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// A flag indicating whether or not to overwrite existing content if a conflict is
	// found during import content operation.
	IsOverwrite *bool `mandatory:"false" contributesTo:"query" name:"isOverwrite"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A value of `100-continue` requests preliminary verification of the request method, path, and headers before the request body is sent.
	// If no error results from such verification, the server will send a 100 (Continue) interim response to indicate readiness for the request body.
	// The only allowed value for this parameter is "100-Continue" (case-insensitive).
	Expect *string `mandatory:"false" contributesTo:"header" name:"expect"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ImportCustomContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ImportCustomContentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	contents := make([]byte, 1335)
	binaryRequestBody.Read(contents)

	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	//if err == nil && binaryRequestBody.Seekable() {
	//	common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	//}
	httpRequest.Body = io.NopCloser(bytes.NewReader(contents))
	//httpRequest.Body = binaryRequestBody
	//bb := make([]byte, 1335)
	//httpRequest.Body.Read(bb)
	//common.Debugf("[SDK]import_custom_content_file body2 ................%v||seekable: %v", string(bb),binaryRequestBody.Seekable())

	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request ImportCustomContentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.ImportCustomContentFileBody)
	if rsc.Seekable() {
		return rsc, true
	}
	return rsc, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ImportCustomContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ImportCustomContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImportCustomContentResponse wrapper for the ImportCustomContent operation
type ImportCustomContentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LogAnalyticsImportCustomContent instance
	LogAnalyticsImportCustomContent `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ImportCustomContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ImportCustomContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
