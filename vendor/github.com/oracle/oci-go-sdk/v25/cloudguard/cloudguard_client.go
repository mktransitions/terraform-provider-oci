// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v25/common"
	"github.com/oracle/oci-go-sdk/v25/common/auth"
	"net/http"
)

//CloudGuardClient a client for CloudGuard
type CloudGuardClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewCloudGuardClientWithConfigurationProvider Creates a new default CloudGuard client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewCloudGuardClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client CloudGuardClient, err error) {
	if provider, err := auth.GetGenericConfigurationProvider(configProvider); err == nil {
		if baseClient, err := common.NewClientWithConfig(provider); err == nil {
			return newCloudGuardClientFromBaseClient(baseClient, configProvider)
		}
	}

	return
}

// NewCloudGuardClientWithOboToken Creates a new default CloudGuard client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewCloudGuardClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client CloudGuardClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newCloudGuardClientFromBaseClient(baseClient, configProvider)
}

func newCloudGuardClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client CloudGuardClient, err error) {
	client = CloudGuardClient{BaseClient: baseClient}
	client.BasePath = "20200131"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *CloudGuardClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("cloudguard", "https://cloudguard-cp-api.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *CloudGuardClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *CloudGuardClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeDetectorRecipeCompartment Moves the DetectorRecipe from current compartment to another.
func (client CloudGuardClient) ChangeDetectorRecipeCompartment(ctx context.Context, request ChangeDetectorRecipeCompartmentRequest) (response ChangeDetectorRecipeCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeDetectorRecipeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDetectorRecipeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDetectorRecipeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDetectorRecipeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDetectorRecipeCompartmentResponse")
	}
	return
}

// changeDetectorRecipeCompartment implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) changeDetectorRecipeCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/detectorRecipes/{detectorRecipeId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeDetectorRecipeCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeManagedListCompartment Moves the ManagedList from current compartment to another.
func (client CloudGuardClient) ChangeManagedListCompartment(ctx context.Context, request ChangeManagedListCompartmentRequest) (response ChangeManagedListCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeManagedListCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeManagedListCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeManagedListCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeManagedListCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeManagedListCompartmentResponse")
	}
	return
}

// changeManagedListCompartment implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) changeManagedListCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedLists/{managedListId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeManagedListCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeResponderRecipeCompartment Moves the ResponderRecipe from current compartment to another.
func (client CloudGuardClient) ChangeResponderRecipeCompartment(ctx context.Context, request ChangeResponderRecipeCompartmentRequest) (response ChangeResponderRecipeCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeResponderRecipeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeResponderRecipeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeResponderRecipeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeResponderRecipeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeResponderRecipeCompartmentResponse")
	}
	return
}

// changeResponderRecipeCompartment implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) changeResponderRecipeCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderRecipes/{responderRecipeId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeResponderRecipeCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDetectorRecipe Creates a DetectorRecipe
func (client CloudGuardClient) CreateDetectorRecipe(ctx context.Context, request CreateDetectorRecipeRequest) (response CreateDetectorRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDetectorRecipeResponse")
	}
	return
}

// createDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createDetectorRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/detectorRecipes")
	if err != nil {
		return nil, err
	}

	var response CreateDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateManagedList Creates a new ManagedList.
func (client CloudGuardClient) CreateManagedList(ctx context.Context, request CreateManagedListRequest) (response CreateManagedListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createManagedList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateManagedListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateManagedListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateManagedListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateManagedListResponse")
	}
	return
}

// createManagedList implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createManagedList(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedLists")
	if err != nil {
		return nil, err
	}

	var response CreateManagedListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateResponderRecipe Create a ResponderRecipe.
func (client CloudGuardClient) CreateResponderRecipe(ctx context.Context, request CreateResponderRecipeRequest) (response CreateResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateResponderRecipeResponse")
	}
	return
}

// createResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createResponderRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderRecipes")
	if err != nil {
		return nil, err
	}

	var response CreateResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTarget Creates a new Target
func (client CloudGuardClient) CreateTarget(ctx context.Context, request CreateTargetRequest) (response CreateTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTargetResponse")
	}
	return
}

// createTarget implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createTarget(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targets")
	if err != nil {
		return nil, err
	}

	var response CreateTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTargetDetectorRecipe Attach a DetectorRecipe with the Target
func (client CloudGuardClient) CreateTargetDetectorRecipe(ctx context.Context, request CreateTargetDetectorRecipeRequest) (response CreateTargetDetectorRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTargetDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTargetDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTargetDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTargetDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTargetDetectorRecipeResponse")
	}
	return
}

// createTargetDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createTargetDetectorRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targets/{targetId}/targetDetectorRecipes")
	if err != nil {
		return nil, err
	}

	var response CreateTargetDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTargetResponderRecipe Attach a ResponderRecipe with the Target
func (client CloudGuardClient) CreateTargetResponderRecipe(ctx context.Context, request CreateTargetResponderRecipeRequest) (response CreateTargetResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTargetResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTargetResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTargetResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTargetResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTargetResponderRecipeResponse")
	}
	return
}

// createTargetResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createTargetResponderRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targets/{targetId}/targetResponderRecipes")
	if err != nil {
		return nil, err
	}

	var response CreateTargetResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDetectorRecipe Deletes a DetectorRecipe identified by detectorRecipeId
func (client CloudGuardClient) DeleteDetectorRecipe(ctx context.Context, request DeleteDetectorRecipeRequest) (response DeleteDetectorRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDetectorRecipeResponse")
	}
	return
}

// deleteDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteDetectorRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/detectorRecipes/{detectorRecipeId}")
	if err != nil {
		return nil, err
	}

	var response DeleteDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteManagedList Deletes a managed list identified by managedListId
func (client CloudGuardClient) DeleteManagedList(ctx context.Context, request DeleteManagedListRequest) (response DeleteManagedListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteManagedList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagedListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagedListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagedListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagedListResponse")
	}
	return
}

// deleteManagedList implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteManagedList(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managedLists/{managedListId}")
	if err != nil {
		return nil, err
	}

	var response DeleteManagedListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteResponderRecipe Delete the ResponderRecipe resource by identifier
func (client CloudGuardClient) DeleteResponderRecipe(ctx context.Context, request DeleteResponderRecipeRequest) (response DeleteResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteResponderRecipeResponse")
	}
	return
}

// deleteResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteResponderRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/responderRecipes/{responderRecipeId}")
	if err != nil {
		return nil, err
	}

	var response DeleteResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTarget Deletes a Target identified by targetId
func (client CloudGuardClient) DeleteTarget(ctx context.Context, request DeleteTargetRequest) (response DeleteTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTargetResponse")
	}
	return
}

// deleteTarget implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteTarget(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/targets/{targetId}")
	if err != nil {
		return nil, err
	}

	var response DeleteTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTargetDetectorRecipe Delete the TargetDetectorRecipe resource by identifier
func (client CloudGuardClient) DeleteTargetDetectorRecipe(ctx context.Context, request DeleteTargetDetectorRecipeRequest) (response DeleteTargetDetectorRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTargetDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTargetDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTargetDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTargetDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTargetDetectorRecipeResponse")
	}
	return
}

// deleteTargetDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteTargetDetectorRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/targets/{targetId}/targetDetectorRecipes/{targetDetectorRecipeId}")
	if err != nil {
		return nil, err
	}

	var response DeleteTargetDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTargetResponderRecipe Delete the TargetResponderRecipe resource by identifier
func (client CloudGuardClient) DeleteTargetResponderRecipe(ctx context.Context, request DeleteTargetResponderRecipeRequest) (response DeleteTargetResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTargetResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTargetResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTargetResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTargetResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTargetResponderRecipeResponse")
	}
	return
}

// deleteTargetResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteTargetResponderRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/targets/{targetId}/targetResponderRecipes/{targetResponderRecipeId}")
	if err != nil {
		return nil, err
	}

	var response DeleteTargetResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExecuteResponderExecution Executes the responder execution. When provided, If-Match is checked against ETag values of the resource.
func (client CloudGuardClient) ExecuteResponderExecution(ctx context.Context, request ExecuteResponderExecutionRequest) (response ExecuteResponderExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.executeResponderExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExecuteResponderExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExecuteResponderExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExecuteResponderExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExecuteResponderExecutionResponse")
	}
	return
}

// executeResponderExecution implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) executeResponderExecution(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderExecutions/{responderExecutionId}/actions/execute")
	if err != nil {
		return nil, err
	}

	var response ExecuteResponderExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConditionMetadataType Returns ConditionType with its details.
func (client CloudGuardClient) GetConditionMetadataType(ctx context.Context, request GetConditionMetadataTypeRequest) (response GetConditionMetadataTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConditionMetadataType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConditionMetadataTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConditionMetadataTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConditionMetadataTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConditionMetadataTypeResponse")
	}
	return
}

// getConditionMetadataType implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getConditionMetadataType(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/conditionMetadataTypes/{conditionMetadataTypeId}")
	if err != nil {
		return nil, err
	}

	var response GetConditionMetadataTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConfiguration GET Cloud Guard Configuration Details for a Tenancy.
func (client CloudGuardClient) GetConfiguration(ctx context.Context, request GetConfigurationRequest) (response GetConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfigurationResponse")
	}
	return
}

// getConfiguration implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configuration")
	if err != nil {
		return nil, err
	}

	var response GetConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDetector Returns a Detector identified by detectorId.
func (client CloudGuardClient) GetDetector(ctx context.Context, request GetDetectorRequest) (response GetDetectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDetector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDetectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDetectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDetectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDetectorResponse")
	}
	return
}

// getDetector implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getDetector(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectors/{detectorId}")
	if err != nil {
		return nil, err
	}

	var response GetDetectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDetectorRecipe Returns a DetectorRecipe identified by detectorRecipeId
func (client CloudGuardClient) GetDetectorRecipe(ctx context.Context, request GetDetectorRecipeRequest) (response GetDetectorRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDetectorRecipeResponse")
	}
	return
}

// getDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getDetectorRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectorRecipes/{detectorRecipeId}")
	if err != nil {
		return nil, err
	}

	var response GetDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDetectorRecipeDetectorRule Get DetectorRule by identifier
func (client CloudGuardClient) GetDetectorRecipeDetectorRule(ctx context.Context, request GetDetectorRecipeDetectorRuleRequest) (response GetDetectorRecipeDetectorRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDetectorRecipeDetectorRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDetectorRecipeDetectorRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDetectorRecipeDetectorRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDetectorRecipeDetectorRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDetectorRecipeDetectorRuleResponse")
	}
	return
}

// getDetectorRecipeDetectorRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getDetectorRecipeDetectorRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectorRecipes/{detectorRecipeId}/detectorRules/{detectorRuleId}")
	if err != nil {
		return nil, err
	}

	var response GetDetectorRecipeDetectorRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDetectorRule Returns a Detector Rule identified by detectorRuleId
func (client CloudGuardClient) GetDetectorRule(ctx context.Context, request GetDetectorRuleRequest) (response GetDetectorRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDetectorRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDetectorRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDetectorRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDetectorRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDetectorRuleResponse")
	}
	return
}

// getDetectorRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getDetectorRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectors/{detectorId}/detectorRules/{detectorRuleId}")
	if err != nil {
		return nil, err
	}

	var response GetDetectorRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedList Returns a managed list identified by managedListId
func (client CloudGuardClient) GetManagedList(ctx context.Context, request GetManagedListRequest) (response GetManagedListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagedList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagedListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagedListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagedListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagedListResponse")
	}
	return
}

// getManagedList implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getManagedList(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedLists/{managedListId}")
	if err != nil {
		return nil, err
	}

	var response GetManagedListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProblem Returns a Problems response
func (client CloudGuardClient) GetProblem(ctx context.Context, request GetProblemRequest) (response GetProblemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProblem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProblemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProblemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProblemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProblemResponse")
	}
	return
}

// getProblem implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getProblem(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/problems/{problemId}")
	if err != nil {
		return nil, err
	}

	var response GetProblemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResponderExecution Returns a Responder Execution identified by responderExecutionId
func (client CloudGuardClient) GetResponderExecution(ctx context.Context, request GetResponderExecutionRequest) (response GetResponderExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResponderExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResponderExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResponderExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResponderExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResponderExecutionResponse")
	}
	return
}

// getResponderExecution implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getResponderExecution(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderExecutions/{responderExecutionId}")
	if err != nil {
		return nil, err
	}

	var response GetResponderExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResponderRecipe Get a ResponderRecipe by identifier
func (client CloudGuardClient) GetResponderRecipe(ctx context.Context, request GetResponderRecipeRequest) (response GetResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResponderRecipeResponse")
	}
	return
}

// getResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getResponderRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderRecipes/{responderRecipeId}")
	if err != nil {
		return nil, err
	}

	var response GetResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResponderRecipeResponderRule Get ResponderRule by identifier
func (client CloudGuardClient) GetResponderRecipeResponderRule(ctx context.Context, request GetResponderRecipeResponderRuleRequest) (response GetResponderRecipeResponderRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResponderRecipeResponderRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResponderRecipeResponderRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResponderRecipeResponderRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResponderRecipeResponderRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResponderRecipeResponderRuleResponse")
	}
	return
}

// getResponderRecipeResponderRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getResponderRecipeResponderRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderRecipes/{responderRecipeId}/responderRules/{responderRuleId}")
	if err != nil {
		return nil, err
	}

	var response GetResponderRecipeResponderRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResponderRule Get a ResponderRule by identifier
func (client CloudGuardClient) GetResponderRule(ctx context.Context, request GetResponderRuleRequest) (response GetResponderRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResponderRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResponderRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResponderRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResponderRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResponderRuleResponse")
	}
	return
}

// getResponderRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getResponderRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderRules/{responderRuleId}")
	if err != nil {
		return nil, err
	}

	var response GetResponderRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTarget Returns a Target identified by targetId
func (client CloudGuardClient) GetTarget(ctx context.Context, request GetTargetRequest) (response GetTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetResponse")
	}
	return
}

// getTarget implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getTarget(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}")
	if err != nil {
		return nil, err
	}

	var response GetTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTargetDetectorRecipe Get a TargetDetectorRecipe by identifier
func (client CloudGuardClient) GetTargetDetectorRecipe(ctx context.Context, request GetTargetDetectorRecipeRequest) (response GetTargetDetectorRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTargetDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetDetectorRecipeResponse")
	}
	return
}

// getTargetDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getTargetDetectorRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetDetectorRecipes/{targetDetectorRecipeId}")
	if err != nil {
		return nil, err
	}

	var response GetTargetDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTargetDetectorRecipeDetectorRule Get DetectorRule by identifier
func (client CloudGuardClient) GetTargetDetectorRecipeDetectorRule(ctx context.Context, request GetTargetDetectorRecipeDetectorRuleRequest) (response GetTargetDetectorRecipeDetectorRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTargetDetectorRecipeDetectorRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetDetectorRecipeDetectorRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetDetectorRecipeDetectorRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetDetectorRecipeDetectorRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetDetectorRecipeDetectorRuleResponse")
	}
	return
}

// getTargetDetectorRecipeDetectorRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getTargetDetectorRecipeDetectorRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetDetectorRecipes/{targetDetectorRecipeId}/detectorRules/{detectorRuleId}")
	if err != nil {
		return nil, err
	}

	var response GetTargetDetectorRecipeDetectorRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTargetResponderRecipe Get a TargetResponderRecipe by identifier
func (client CloudGuardClient) GetTargetResponderRecipe(ctx context.Context, request GetTargetResponderRecipeRequest) (response GetTargetResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTargetResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetResponderRecipeResponse")
	}
	return
}

// getTargetResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getTargetResponderRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetResponderRecipes/{targetResponderRecipeId}")
	if err != nil {
		return nil, err
	}

	var response GetTargetResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTargetResponderRecipeResponderRule Get ResponderRule by identifier
func (client CloudGuardClient) GetTargetResponderRecipeResponderRule(ctx context.Context, request GetTargetResponderRecipeResponderRuleRequest) (response GetTargetResponderRecipeResponderRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTargetResponderRecipeResponderRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetResponderRecipeResponderRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetResponderRecipeResponderRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetResponderRecipeResponderRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetResponderRecipeResponderRuleResponse")
	}
	return
}

// getTargetResponderRecipeResponderRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getTargetResponderRecipeResponderRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetResponderRecipes/{targetResponderRecipeId}/responderRules/{responderRuleId}")
	if err != nil {
		return nil, err
	}

	var response GetTargetResponderRecipeResponderRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConditionMetadataTypes Returns a list of condition types.
func (client CloudGuardClient) ListConditionMetadataTypes(ctx context.Context, request ListConditionMetadataTypesRequest) (response ListConditionMetadataTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConditionMetadataTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConditionMetadataTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConditionMetadataTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConditionMetadataTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConditionMetadataTypesResponse")
	}
	return
}

// listConditionMetadataTypes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listConditionMetadataTypes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/conditionMetadataTypes")
	if err != nil {
		return nil, err
	}

	var response ListConditionMetadataTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDetectorRecipeDetectorRules Returns a list of DetectorRule associated with DetectorRecipe.
func (client CloudGuardClient) ListDetectorRecipeDetectorRules(ctx context.Context, request ListDetectorRecipeDetectorRulesRequest) (response ListDetectorRecipeDetectorRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDetectorRecipeDetectorRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDetectorRecipeDetectorRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDetectorRecipeDetectorRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDetectorRecipeDetectorRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDetectorRecipeDetectorRulesResponse")
	}
	return
}

// listDetectorRecipeDetectorRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listDetectorRecipeDetectorRules(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectorRecipes/{detectorRecipeId}/detectorRules")
	if err != nil {
		return nil, err
	}

	var response ListDetectorRecipeDetectorRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDetectorRecipes Returns a list of all Detector Recipes in a compartment
// The ListDetectorRecipes operation returns only the detector recipes in `compartmentId` passed.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListDetectorRecipes on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
func (client CloudGuardClient) ListDetectorRecipes(ctx context.Context, request ListDetectorRecipesRequest) (response ListDetectorRecipesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDetectorRecipes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDetectorRecipesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDetectorRecipesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDetectorRecipesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDetectorRecipesResponse")
	}
	return
}

// listDetectorRecipes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listDetectorRecipes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectorRecipes")
	if err != nil {
		return nil, err
	}

	var response ListDetectorRecipesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDetectorRules Returns a list of detector rules for the detectorId passed.
func (client CloudGuardClient) ListDetectorRules(ctx context.Context, request ListDetectorRulesRequest) (response ListDetectorRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDetectorRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDetectorRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDetectorRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDetectorRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDetectorRulesResponse")
	}
	return
}

// listDetectorRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listDetectorRules(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectors/{detectorId}/detectorRules")
	if err != nil {
		return nil, err
	}

	var response ListDetectorRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDetectors Returns detector catalog - list of detectors supported by Cloud Guard
func (client CloudGuardClient) ListDetectors(ctx context.Context, request ListDetectorsRequest) (response ListDetectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDetectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDetectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDetectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDetectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDetectorsResponse")
	}
	return
}

// listDetectors implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listDetectors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectors")
	if err != nil {
		return nil, err
	}

	var response ListDetectorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListImpactedResources Returns a list of Impacted Resources for a CloudGuard Problem
func (client CloudGuardClient) ListImpactedResources(ctx context.Context, request ListImpactedResourcesRequest) (response ListImpactedResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listImpactedResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListImpactedResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListImpactedResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListImpactedResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListImpactedResourcesResponse")
	}
	return
}

// listImpactedResources implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listImpactedResources(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/problems/{problemId}/impactedResources")
	if err != nil {
		return nil, err
	}

	var response ListImpactedResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedListTypes Returns all ManagedList types supported by Cloud Guard
func (client CloudGuardClient) ListManagedListTypes(ctx context.Context, request ListManagedListTypesRequest) (response ListManagedListTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedListTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedListTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedListTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedListTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedListTypesResponse")
	}
	return
}

// listManagedListTypes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listManagedListTypes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedListTypes")
	if err != nil {
		return nil, err
	}

	var response ListManagedListTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedLists Returns a list of ListManagedLists.
// The ListManagedLists operation returns only the managed lists in `compartmentId` passed.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return ManagedLists in only
// those compartments for which the requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListManagedLists on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
func (client CloudGuardClient) ListManagedLists(ctx context.Context, request ListManagedListsRequest) (response ListManagedListsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedLists, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedListsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedListsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedListsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedListsResponse")
	}
	return
}

// listManagedLists implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listManagedLists(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedLists")
	if err != nil {
		return nil, err
	}

	var response ListManagedListsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProblemHistories Returns a list of Actions done on CloudGuard Problem
func (client CloudGuardClient) ListProblemHistories(ctx context.Context, request ListProblemHistoriesRequest) (response ListProblemHistoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProblemHistories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProblemHistoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProblemHistoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProblemHistoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProblemHistoriesResponse")
	}
	return
}

// listProblemHistories implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listProblemHistories(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/problems/{problemId}/histories")
	if err != nil {
		return nil, err
	}

	var response ListProblemHistoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProblems Returns a list of all Problems identified by the Cloud Guard
// The ListProblems operation returns only the problems in `compartmentId` passed.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListProblems on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
func (client CloudGuardClient) ListProblems(ctx context.Context, request ListProblemsRequest) (response ListProblemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProblems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProblemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProblemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProblemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProblemsResponse")
	}
	return
}

// listProblems implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listProblems(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/problems")
	if err != nil {
		return nil, err
	}

	var response ListProblemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRecommendations Returns a list of all Recommendations.
func (client CloudGuardClient) ListRecommendations(ctx context.Context, request ListRecommendationsRequest) (response ListRecommendationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRecommendations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRecommendationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRecommendationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRecommendationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRecommendationsResponse")
	}
	return
}

// listRecommendations implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listRecommendations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/recommendations")
	if err != nil {
		return nil, err
	}

	var response ListRecommendationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResourceTypes Returns a list of resource types.
func (client CloudGuardClient) ListResourceTypes(ctx context.Context, request ListResourceTypesRequest) (response ListResourceTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceTypesResponse")
	}
	return
}

// listResourceTypes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResourceTypes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceTypes")
	if err != nil {
		return nil, err
	}

	var response ListResourceTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResponderActivities Returns a list of Responder activities done on CloudGuard Problem
func (client CloudGuardClient) ListResponderActivities(ctx context.Context, request ListResponderActivitiesRequest) (response ListResponderActivitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResponderActivities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResponderActivitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResponderActivitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResponderActivitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResponderActivitiesResponse")
	}
	return
}

// listResponderActivities implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResponderActivities(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/problems/{problemId}/responderActivities")
	if err != nil {
		return nil, err
	}

	var response ListResponderActivitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResponderExecutions Returns a list of Responder Executions. A Responder Execution is an entity that tracks the collective execution of multiple Responder Rule Executions for a given Problem.
func (client CloudGuardClient) ListResponderExecutions(ctx context.Context, request ListResponderExecutionsRequest) (response ListResponderExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResponderExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResponderExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResponderExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResponderExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResponderExecutionsResponse")
	}
	return
}

// listResponderExecutions implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResponderExecutions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderExecutions")
	if err != nil {
		return nil, err
	}

	var response ListResponderExecutionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResponderRecipeResponderRules Returns a list of ResponderRule associated with ResponderRecipe.
func (client CloudGuardClient) ListResponderRecipeResponderRules(ctx context.Context, request ListResponderRecipeResponderRulesRequest) (response ListResponderRecipeResponderRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResponderRecipeResponderRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResponderRecipeResponderRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResponderRecipeResponderRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResponderRecipeResponderRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResponderRecipeResponderRulesResponse")
	}
	return
}

// listResponderRecipeResponderRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResponderRecipeResponderRules(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderRecipes/{responderRecipeId}/responderRules")
	if err != nil {
		return nil, err
	}

	var response ListResponderRecipeResponderRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResponderRecipes Returns a list of all ResponderRecipes in a compartment
// The ListResponderRecipe operation returns only the targets in `compartmentId` passed.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListResponderRecipe on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
func (client CloudGuardClient) ListResponderRecipes(ctx context.Context, request ListResponderRecipesRequest) (response ListResponderRecipesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResponderRecipes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResponderRecipesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResponderRecipesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResponderRecipesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResponderRecipesResponse")
	}
	return
}

// listResponderRecipes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResponderRecipes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderRecipes")
	if err != nil {
		return nil, err
	}

	var response ListResponderRecipesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResponderRules Returns a list of ResponderRule.
func (client CloudGuardClient) ListResponderRules(ctx context.Context, request ListResponderRulesRequest) (response ListResponderRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResponderRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResponderRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResponderRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResponderRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResponderRulesResponse")
	}
	return
}

// listResponderRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResponderRules(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderRules")
	if err != nil {
		return nil, err
	}

	var response ListResponderRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargetDetectorRecipeDetectorRules Returns a list of DetectorRule associated with DetectorRecipe within a Target.
func (client CloudGuardClient) ListTargetDetectorRecipeDetectorRules(ctx context.Context, request ListTargetDetectorRecipeDetectorRulesRequest) (response ListTargetDetectorRecipeDetectorRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargetDetectorRecipeDetectorRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetDetectorRecipeDetectorRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetDetectorRecipeDetectorRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetDetectorRecipeDetectorRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetDetectorRecipeDetectorRulesResponse")
	}
	return
}

// listTargetDetectorRecipeDetectorRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listTargetDetectorRecipeDetectorRules(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetDetectorRecipes/{targetDetectorRecipeId}/detectorRules")
	if err != nil {
		return nil, err
	}

	var response ListTargetDetectorRecipeDetectorRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargetDetectorRecipes Returns a list of all detector recipes associated with the target identified by targetId
func (client CloudGuardClient) ListTargetDetectorRecipes(ctx context.Context, request ListTargetDetectorRecipesRequest) (response ListTargetDetectorRecipesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargetDetectorRecipes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetDetectorRecipesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetDetectorRecipesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetDetectorRecipesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetDetectorRecipesResponse")
	}
	return
}

// listTargetDetectorRecipes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listTargetDetectorRecipes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetDetectorRecipes")
	if err != nil {
		return nil, err
	}

	var response ListTargetDetectorRecipesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargetResponderRecipeResponderRules Returns a list of ResponderRule associated with ResponderRecipe within a Target.
func (client CloudGuardClient) ListTargetResponderRecipeResponderRules(ctx context.Context, request ListTargetResponderRecipeResponderRulesRequest) (response ListTargetResponderRecipeResponderRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargetResponderRecipeResponderRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetResponderRecipeResponderRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetResponderRecipeResponderRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetResponderRecipeResponderRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetResponderRecipeResponderRulesResponse")
	}
	return
}

// listTargetResponderRecipeResponderRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listTargetResponderRecipeResponderRules(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetResponderRecipes/{targetResponderRecipeId}/responderRules")
	if err != nil {
		return nil, err
	}

	var response ListTargetResponderRecipeResponderRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargetResponderRecipes Returns a list of all responder recipes associated with the target identified by targetId
func (client CloudGuardClient) ListTargetResponderRecipes(ctx context.Context, request ListTargetResponderRecipesRequest) (response ListTargetResponderRecipesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargetResponderRecipes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetResponderRecipesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetResponderRecipesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetResponderRecipesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetResponderRecipesResponse")
	}
	return
}

// listTargetResponderRecipes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listTargetResponderRecipes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetResponderRecipes")
	if err != nil {
		return nil, err
	}

	var response ListTargetResponderRecipesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargets Returns a list of all Targets in a compartment
// The ListTargets operation returns only the targets in `compartmentId` passed.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListTargets on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
func (client CloudGuardClient) ListTargets(ctx context.Context, request ListTargetsRequest) (response ListTargetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetsResponse")
	}
	return
}

// listTargets implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listTargets(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets")
	if err != nil {
		return nil, err
	}

	var response ListTargetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestRiskScores Examines the number of problems related to the resource and the relative severity of those problems.
func (client CloudGuardClient) RequestRiskScores(ctx context.Context, request RequestRiskScoresRequest) (response RequestRiskScoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestRiskScores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestRiskScoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestRiskScoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestRiskScoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestRiskScoresResponse")
	}
	return
}

// requestRiskScores implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestRiskScores(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/riskScores")
	if err != nil {
		return nil, err
	}

	var response RequestRiskScoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSecurityScoreSummarizedTrend Measures the number of resources examined across all regions and compares it with the
// number of problems detected, for a given time period.
func (client CloudGuardClient) RequestSecurityScoreSummarizedTrend(ctx context.Context, request RequestSecurityScoreSummarizedTrendRequest) (response RequestSecurityScoreSummarizedTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSecurityScoreSummarizedTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSecurityScoreSummarizedTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSecurityScoreSummarizedTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSecurityScoreSummarizedTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSecurityScoreSummarizedTrendResponse")
	}
	return
}

// requestSecurityScoreSummarizedTrend implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSecurityScoreSummarizedTrend(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityScores/actions/summarizeTrend")
	if err != nil {
		return nil, err
	}

	var response RequestSecurityScoreSummarizedTrendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSecurityScores Measures the number of resources examined across all regions and compares it with the number of problems detected.
func (client CloudGuardClient) RequestSecurityScores(ctx context.Context, request RequestSecurityScoresRequest) (response RequestSecurityScoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSecurityScores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSecurityScoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSecurityScoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSecurityScoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSecurityScoresResponse")
	}
	return
}

// requestSecurityScores implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSecurityScores(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityScores")
	if err != nil {
		return nil, err
	}

	var response RequestSecurityScoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedActivityProblems Returns the summary of Activity type problems identified by cloud guard, for a given set of dimensions.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform summarize API on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
// The compartmentId to be passed with `accessLevel` and `compartmentIdInSubtree` params has to be the root
// compartment id (tenant-id) only.
func (client CloudGuardClient) RequestSummarizedActivityProblems(ctx context.Context, request RequestSummarizedActivityProblemsRequest) (response RequestSummarizedActivityProblemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedActivityProblems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedActivityProblemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedActivityProblemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedActivityProblemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedActivityProblemsResponse")
	}
	return
}

// requestSummarizedActivityProblems implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedActivityProblems(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/summarizeActivityProblems")
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedActivityProblemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedProblems Returns the number of problems identified by cloud guard, for a given set of dimensions.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform summarize API on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
func (client CloudGuardClient) RequestSummarizedProblems(ctx context.Context, request RequestSummarizedProblemsRequest) (response RequestSummarizedProblemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedProblems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedProblemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedProblemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedProblemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedProblemsResponse")
	}
	return
}

// requestSummarizedProblems implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedProblems(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/summarize")
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedProblemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedResponderExecutions Returns the number of Responder Executions, for a given set of dimensions.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform summarize API on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
func (client CloudGuardClient) RequestSummarizedResponderExecutions(ctx context.Context, request RequestSummarizedResponderExecutionsRequest) (response RequestSummarizedResponderExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedResponderExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedResponderExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedResponderExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedResponderExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedResponderExecutionsResponse")
	}
	return
}

// requestSummarizedResponderExecutions implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedResponderExecutions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderExecutions/actions/summarize")
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedResponderExecutionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedRiskScores DEPRECATED
func (client CloudGuardClient) RequestSummarizedRiskScores(ctx context.Context, request RequestSummarizedRiskScoresRequest) (response RequestSummarizedRiskScoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedRiskScores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedRiskScoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedRiskScoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedRiskScoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedRiskScoresResponse")
	}
	return
}

// requestSummarizedRiskScores implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedRiskScores(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/summarizeRiskScore")
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedRiskScoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedSecurityScores DEPRECATED
func (client CloudGuardClient) RequestSummarizedSecurityScores(ctx context.Context, request RequestSummarizedSecurityScoresRequest) (response RequestSummarizedSecurityScoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedSecurityScores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedSecurityScoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedSecurityScoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedSecurityScoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedSecurityScoresResponse")
	}
	return
}

// requestSummarizedSecurityScores implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedSecurityScores(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/summarizeSecurityScore")
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedSecurityScoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedTrendProblems Returns the number of problems identified by cloud guard, for a given time period.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform summarize API on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
func (client CloudGuardClient) RequestSummarizedTrendProblems(ctx context.Context, request RequestSummarizedTrendProblemsRequest) (response RequestSummarizedTrendProblemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedTrendProblems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedTrendProblemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedTrendProblemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedTrendProblemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedTrendProblemsResponse")
	}
	return
}

// requestSummarizedTrendProblems implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedTrendProblems(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/summarizeTrend")
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedTrendProblemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedTrendResponderExecutions Returns the number of remediations performed by Responders, for a given time period.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform summarize API on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
func (client CloudGuardClient) RequestSummarizedTrendResponderExecutions(ctx context.Context, request RequestSummarizedTrendResponderExecutionsRequest) (response RequestSummarizedTrendResponderExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedTrendResponderExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedTrendResponderExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedTrendResponderExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedTrendResponderExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedTrendResponderExecutionsResponse")
	}
	return
}

// requestSummarizedTrendResponderExecutions implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedTrendResponderExecutions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderExecutions/actions/summarizeTrend")
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedTrendResponderExecutionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedTrendSecurityScores DEPRECATED
func (client CloudGuardClient) RequestSummarizedTrendSecurityScores(ctx context.Context, request RequestSummarizedTrendSecurityScoresRequest) (response RequestSummarizedTrendSecurityScoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedTrendSecurityScores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedTrendSecurityScoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedTrendSecurityScoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedTrendSecurityScoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedTrendSecurityScoresResponse")
	}
	return
}

// requestSummarizedTrendSecurityScores implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedTrendSecurityScores(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/summarizeSecurityScoreTrend")
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedTrendSecurityScoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SkipBulkResponderExecution Skips the execution for a bulk of responder executions
// The operation is atomic in nature
func (client CloudGuardClient) SkipBulkResponderExecution(ctx context.Context, request SkipBulkResponderExecutionRequest) (response SkipBulkResponderExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.skipBulkResponderExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SkipBulkResponderExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SkipBulkResponderExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SkipBulkResponderExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SkipBulkResponderExecutionResponse")
	}
	return
}

// skipBulkResponderExecution implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) skipBulkResponderExecution(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderExecutions/actions/bulkSkip")
	if err != nil {
		return nil, err
	}

	var response SkipBulkResponderExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SkipResponderExecution Skips the execution of the responder execution. When provided, If-Match is checked against ETag values of the resource.
func (client CloudGuardClient) SkipResponderExecution(ctx context.Context, request SkipResponderExecutionRequest) (response SkipResponderExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.skipResponderExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SkipResponderExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SkipResponderExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SkipResponderExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SkipResponderExecutionResponse")
	}
	return
}

// skipResponderExecution implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) skipResponderExecution(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderExecutions/{responderExecutionId}/actions/skip")
	if err != nil {
		return nil, err
	}

	var response SkipResponderExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TriggerResponder push the problem to responder
func (client CloudGuardClient) TriggerResponder(ctx context.Context, request TriggerResponderRequest) (response TriggerResponderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.triggerResponder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TriggerResponderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TriggerResponderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TriggerResponderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TriggerResponderResponse")
	}
	return
}

// triggerResponder implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) triggerResponder(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/{problemId}/actions/triggerResponder")
	if err != nil {
		return nil, err
	}

	var response TriggerResponderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBulkProblemStatus Updates the statuses in bulk for a list of problems
// The operation is atomic in nature
func (client CloudGuardClient) UpdateBulkProblemStatus(ctx context.Context, request UpdateBulkProblemStatusRequest) (response UpdateBulkProblemStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBulkProblemStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBulkProblemStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBulkProblemStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBulkProblemStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBulkProblemStatusResponse")
	}
	return
}

// updateBulkProblemStatus implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateBulkProblemStatus(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/bulkUpdateStatus")
	if err != nil {
		return nil, err
	}

	var response UpdateBulkProblemStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConfiguration Enable/Disable Cloud Guard. The reporting region cannot be updated once created.
func (client CloudGuardClient) UpdateConfiguration(ctx context.Context, request UpdateConfigurationRequest) (response UpdateConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConfigurationResponse")
	}
	return
}

// updateConfiguration implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateConfiguration(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/configuration")
	if err != nil {
		return nil, err
	}

	var response UpdateConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDetectorRecipe Updates a detector recipe identified by detectorRecipeId
func (client CloudGuardClient) UpdateDetectorRecipe(ctx context.Context, request UpdateDetectorRecipeRequest) (response UpdateDetectorRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDetectorRecipeResponse")
	}
	return
}

// updateDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateDetectorRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/detectorRecipes/{detectorRecipeId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDetectorRecipeDetectorRule Update the DetectorRule by identifier
func (client CloudGuardClient) UpdateDetectorRecipeDetectorRule(ctx context.Context, request UpdateDetectorRecipeDetectorRuleRequest) (response UpdateDetectorRecipeDetectorRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDetectorRecipeDetectorRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDetectorRecipeDetectorRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDetectorRecipeDetectorRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDetectorRecipeDetectorRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDetectorRecipeDetectorRuleResponse")
	}
	return
}

// updateDetectorRecipeDetectorRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateDetectorRecipeDetectorRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/detectorRecipes/{detectorRecipeId}/detectorRules/{detectorRuleId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDetectorRecipeDetectorRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateManagedList Updates a managed list identified by managedListId
func (client CloudGuardClient) UpdateManagedList(ctx context.Context, request UpdateManagedListRequest) (response UpdateManagedListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateManagedList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagedListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagedListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagedListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagedListResponse")
	}
	return
}

// updateManagedList implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateManagedList(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managedLists/{managedListId}")
	if err != nil {
		return nil, err
	}

	var response UpdateManagedListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateProblemStatus updates the problem details
func (client CloudGuardClient) UpdateProblemStatus(ctx context.Context, request UpdateProblemStatusRequest) (response UpdateProblemStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateProblemStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateProblemStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateProblemStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateProblemStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateProblemStatusResponse")
	}
	return
}

// updateProblemStatus implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateProblemStatus(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/{problemId}/actions/updateStatus")
	if err != nil {
		return nil, err
	}

	var response UpdateProblemStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateResponderRecipe Update the ResponderRecipe resource by identifier
func (client CloudGuardClient) UpdateResponderRecipe(ctx context.Context, request UpdateResponderRecipeRequest) (response UpdateResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateResponderRecipeResponse")
	}
	return
}

// updateResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateResponderRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/responderRecipes/{responderRecipeId}")
	if err != nil {
		return nil, err
	}

	var response UpdateResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateResponderRecipeResponderRule Update the ResponderRule by identifier
func (client CloudGuardClient) UpdateResponderRecipeResponderRule(ctx context.Context, request UpdateResponderRecipeResponderRuleRequest) (response UpdateResponderRecipeResponderRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateResponderRecipeResponderRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateResponderRecipeResponderRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateResponderRecipeResponderRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateResponderRecipeResponderRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateResponderRecipeResponderRuleResponse")
	}
	return
}

// updateResponderRecipeResponderRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateResponderRecipeResponderRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/responderRecipes/{responderRecipeId}/responderRules/{responderRuleId}")
	if err != nil {
		return nil, err
	}

	var response UpdateResponderRecipeResponderRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTarget Updates a Target identified by targetId
func (client CloudGuardClient) UpdateTarget(ctx context.Context, request UpdateTargetRequest) (response UpdateTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetResponse")
	}
	return
}

// updateTarget implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateTarget(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targets/{targetId}")
	if err != nil {
		return nil, err
	}

	var response UpdateTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTargetDetectorRecipe Update the TargetDetectorRecipe resource by identifier
func (client CloudGuardClient) UpdateTargetDetectorRecipe(ctx context.Context, request UpdateTargetDetectorRecipeRequest) (response UpdateTargetDetectorRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTargetDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetDetectorRecipeResponse")
	}
	return
}

// updateTargetDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateTargetDetectorRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targets/{targetId}/targetDetectorRecipes/{targetDetectorRecipeId}")
	if err != nil {
		return nil, err
	}

	var response UpdateTargetDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTargetDetectorRecipeDetectorRule Update the DetectorRule by identifier
func (client CloudGuardClient) UpdateTargetDetectorRecipeDetectorRule(ctx context.Context, request UpdateTargetDetectorRecipeDetectorRuleRequest) (response UpdateTargetDetectorRecipeDetectorRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTargetDetectorRecipeDetectorRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetDetectorRecipeDetectorRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetDetectorRecipeDetectorRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetDetectorRecipeDetectorRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetDetectorRecipeDetectorRuleResponse")
	}
	return
}

// updateTargetDetectorRecipeDetectorRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateTargetDetectorRecipeDetectorRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targets/{targetId}/targetDetectorRecipes/{targetDetectorRecipeId}/detectorRules/{detectorRuleId}")
	if err != nil {
		return nil, err
	}

	var response UpdateTargetDetectorRecipeDetectorRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTargetResponderRecipe Update the TargetResponderRecipe resource by identifier
func (client CloudGuardClient) UpdateTargetResponderRecipe(ctx context.Context, request UpdateTargetResponderRecipeRequest) (response UpdateTargetResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTargetResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetResponderRecipeResponse")
	}
	return
}

// updateTargetResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateTargetResponderRecipe(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targets/{targetId}/targetResponderRecipes/{targetResponderRecipeId}")
	if err != nil {
		return nil, err
	}

	var response UpdateTargetResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTargetResponderRecipeResponderRule Update the ResponderRule by identifier
func (client CloudGuardClient) UpdateTargetResponderRecipeResponderRule(ctx context.Context, request UpdateTargetResponderRecipeResponderRuleRequest) (response UpdateTargetResponderRecipeResponderRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTargetResponderRecipeResponderRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetResponderRecipeResponderRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetResponderRecipeResponderRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetResponderRecipeResponderRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetResponderRecipeResponderRuleResponse")
	}
	return
}

// updateTargetResponderRecipeResponderRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateTargetResponderRecipeResponderRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targets/{targetId}/targetResponderRecipes/{targetResponderRecipeId}/responderRules/{responderRuleId}")
	if err != nil {
		return nil, err
	}

	var response UpdateTargetResponderRecipeResponderRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
