// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_file_storage "github.com/oracle/oci-go-sdk/v25/filestorage"

	oci_common "github.com/oracle/oci-go-sdk/v25/common"
)

func init() {
	RegisterOracleClient("oci_file_storage.FileStorageClient", &OracleClient{initClientFn: initFilestorageFileStorageClient})
}

func initFilestorageFileStorageClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_file_storage.NewFileStorageClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) fileStorageClient() *oci_file_storage.FileStorageClient {
	return m.GetClient("oci_file_storage.FileStorageClient").(*oci_file_storage.FileStorageClient)
}
