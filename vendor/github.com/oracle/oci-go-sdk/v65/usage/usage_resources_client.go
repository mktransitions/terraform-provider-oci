// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage Proxy API
//
// Use the Usage Proxy API to list Oracle Support Rewards, view related detailed usage information, and manage users who redeem rewards. For more information, see Oracle Support Rewards Overview (https://docs.cloud.oracle.com/iaas/Content/Billing/Concepts/supportrewardsoverview.htm).
//

package usage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ResourcesClient a client for Resources
type ResourcesClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewResourcesClientWithConfigurationProvider Creates a new default Resources client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewResourcesClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ResourcesClient, err error) {
	if enabled := common.CheckForEnabledServices("usage"); !enabled {
		return client, fmt.Errorf("the Alloy configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local alloy_config file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newResourcesClientFromBaseClient(baseClient, provider)
}

// NewResourcesClientWithOboToken Creates a new default Resources client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewResourcesClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ResourcesClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newResourcesClientFromBaseClient(baseClient, configProvider)
}

func newResourcesClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ResourcesClient, err error) {
	// Resources service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Resources"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ResourcesClient{BaseClient: baseClient}
	client.BasePath = "20190111"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ResourcesClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("identity", "https://identity.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ResourcesClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *ResourcesClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListResourceQuota Returns the resource quota details under a tenancy
// > **Important**: Calls to this API will only succeed against the endpoint in the home region.
// A default retry strategy applies to this operation ListResourceQuota()
func (client ResourcesClient) ListResourceQuota(ctx context.Context, request ListResourceQuotaRequest) (response ListResourceQuotaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceQuota, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceQuotaResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceQuotaResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceQuotaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceQuotaResponse")
	}
	return
}

// listResourceQuota implements the OCIOperation interface (enables retrying operations)
func (client ResourcesClient) listResourceQuota(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resources/quota", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourceQuotaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage-proxy/20190111/ResourceQuotumSummary/ListResourceQuota"
		err = common.PostProcessServiceError(err, "Resources", "ListResourceQuota", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResources Returns the resource details for a service
// > **Important**: Calls to this API will only succeed against the endpoint in the home region.
// A default retry strategy applies to this operation ListResources()
func (client ResourcesClient) ListResources(ctx context.Context, request ListResourcesRequest) (response ListResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourcesResponse")
	}
	return
}

// listResources implements the OCIOperation interface (enables retrying operations)
func (client ResourcesClient) listResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage-proxy/20190111/ResourceSummary/ListResources"
		err = common.PostProcessServiceError(err, "Resources", "ListResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
