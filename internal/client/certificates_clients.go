// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_certificates "github.com/oracle/oci-go-sdk/v65/certificates"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_certificates.CertificatesClient", &OracleClient{InitClientFn: initCertificatesCertificatesClient})
}

func initCertificatesCertificatesClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_certificates.NewCertificatesClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) CertificatesClient() *oci_certificates.CertificatesClient {
	return m.GetClient("oci_certificates.CertificatesClient").(*oci_certificates.CertificatesClient)
}
