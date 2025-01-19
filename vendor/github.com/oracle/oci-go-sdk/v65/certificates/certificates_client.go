// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Retrieval API
//
// API for retrieving certificates.
//

package certificates

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// CertificatesClient a client for Certificates
type CertificatesClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewCertificatesClientWithConfigurationProvider Creates a new default Certificates client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewCertificatesClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client CertificatesClient, err error) {
	if enabled := common.CheckForEnabledServices("certificates"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newCertificatesClientFromBaseClient(baseClient, provider)
}

// NewCertificatesClientWithOboToken Creates a new default Certificates client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewCertificatesClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client CertificatesClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newCertificatesClientFromBaseClient(baseClient, configProvider)
}

func newCertificatesClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client CertificatesClient, err error) {
	// Certificates service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Certificates"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = CertificatesClient{BaseClient: baseClient}
	client.BasePath = "20210224"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *CertificatesClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("certificates", "https://certificates.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *CertificatesClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *CertificatesClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetCaBundle Gets a ca-bundle bundle.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificates/GetCaBundle.go.html to see an example of how to use GetCaBundle API.
func (client CertificatesClient) GetCaBundle(ctx context.Context, request GetCaBundleRequest) (response GetCaBundleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCaBundle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCaBundleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCaBundleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCaBundleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCaBundleResponse")
	}
	return
}

// getCaBundle implements the OCIOperation interface (enables retrying operations)
func (client CertificatesClient) getCaBundle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/caBundles/{caBundleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCaBundleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/certificates/20210224/CaBundle/GetCaBundle"
		err = common.PostProcessServiceError(err, "Certificates", "GetCaBundle", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCertificateAuthorityBundle Gets a certificate authority bundle that matches either the specified `stage`, `name`, or `versionNumber` parameter.
// If none of these parameters are provided, the bundle for the certificate authority version marked as `CURRENT` will be returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificates/GetCertificateAuthorityBundle.go.html to see an example of how to use GetCertificateAuthorityBundle API.
func (client CertificatesClient) GetCertificateAuthorityBundle(ctx context.Context, request GetCertificateAuthorityBundleRequest) (response GetCertificateAuthorityBundleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCertificateAuthorityBundle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCertificateAuthorityBundleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCertificateAuthorityBundleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCertificateAuthorityBundleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCertificateAuthorityBundleResponse")
	}
	return
}

// getCertificateAuthorityBundle implements the OCIOperation interface (enables retrying operations)
func (client CertificatesClient) getCertificateAuthorityBundle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificateAuthorityBundles/{certificateAuthorityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCertificateAuthorityBundleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/certificates/20210224/CertificateAuthorityBundle/GetCertificateAuthorityBundle"
		err = common.PostProcessServiceError(err, "Certificates", "GetCertificateAuthorityBundle", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCertificateBundle Gets a certificate bundle that matches either the specified `stage`, `versionName`, or `versionNumber` parameter.
// If none of these parameters are provided, the bundle for the certificate version marked as `CURRENT` will be returned.
// By default, the private key is not included in the query result, and a CertificateBundlePublicOnly is returned.
// If the private key is needed, use the CertificateBundleTypeQueryParam parameter to get a CertificateBundleWithPrivateKey response.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificates/GetCertificateBundle.go.html to see an example of how to use GetCertificateBundle API.
func (client CertificatesClient) GetCertificateBundle(ctx context.Context, request GetCertificateBundleRequest) (response GetCertificateBundleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCertificateBundle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCertificateBundleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCertificateBundleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCertificateBundleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCertificateBundleResponse")
	}
	return
}

// getCertificateBundle implements the OCIOperation interface (enables retrying operations)
func (client CertificatesClient) getCertificateBundle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificateBundles/{certificateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCertificateBundleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/certificates/20210224/CertificateBundle/GetCertificateBundle"
		err = common.PostProcessServiceError(err, "Certificates", "GetCertificateBundle", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &certificatebundle{})
	return response, err
}

// ListCertificateAuthorityBundleVersions Lists all certificate authority bundle versions for the specified certificate authority.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificates/ListCertificateAuthorityBundleVersions.go.html to see an example of how to use ListCertificateAuthorityBundleVersions API.
func (client CertificatesClient) ListCertificateAuthorityBundleVersions(ctx context.Context, request ListCertificateAuthorityBundleVersionsRequest) (response ListCertificateAuthorityBundleVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCertificateAuthorityBundleVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCertificateAuthorityBundleVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCertificateAuthorityBundleVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCertificateAuthorityBundleVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCertificateAuthorityBundleVersionsResponse")
	}
	return
}

// listCertificateAuthorityBundleVersions implements the OCIOperation interface (enables retrying operations)
func (client CertificatesClient) listCertificateAuthorityBundleVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificateAuthorityBundles/{certificateAuthorityId}/versions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCertificateAuthorityBundleVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/certificates/20210224/CertificateAuthorityBundleVersionSummary/ListCertificateAuthorityBundleVersions"
		err = common.PostProcessServiceError(err, "Certificates", "ListCertificateAuthorityBundleVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCertificateBundleVersions Lists all certificate bundle versions for the specified certificate.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificates/ListCertificateBundleVersions.go.html to see an example of how to use ListCertificateBundleVersions API.
func (client CertificatesClient) ListCertificateBundleVersions(ctx context.Context, request ListCertificateBundleVersionsRequest) (response ListCertificateBundleVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCertificateBundleVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCertificateBundleVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCertificateBundleVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCertificateBundleVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCertificateBundleVersionsResponse")
	}
	return
}

// listCertificateBundleVersions implements the OCIOperation interface (enables retrying operations)
func (client CertificatesClient) listCertificateBundleVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificateBundles/{certificateId}/versions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCertificateBundleVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/certificates/20210224/CertificateBundleVersionSummary/ListCertificateBundleVersions"
		err = common.PostProcessServiceError(err, "Certificates", "ListCertificateBundleVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
