// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_budget "github.com/oracle/oci-go-sdk/v25/budget"

	oci_common "github.com/oracle/oci-go-sdk/v25/common"
)

func init() {
	RegisterOracleClient("oci_budget.BudgetClient", &OracleClient{initClientFn: initBudgetBudgetClient})
}

func initBudgetBudgetClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_budget.NewBudgetClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) budgetClient() *oci_budget.BudgetClient {
	return m.GetClient("oci_budget.BudgetClient").(*oci_budget.BudgetClient)
}
