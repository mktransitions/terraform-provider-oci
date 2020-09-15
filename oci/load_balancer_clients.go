// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_load_balancer "github.com/oracle/oci-go-sdk/v25/loadbalancer"

	oci_common "github.com/oracle/oci-go-sdk/v25/common"
)

func init() {
	RegisterOracleClient("oci_load_balancer.LoadBalancerClient", &OracleClient{initClientFn: initLoadbalancerLoadBalancerClient})
}

func initLoadbalancerLoadBalancerClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_load_balancer.NewLoadBalancerClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) loadBalancerClient() *oci_load_balancer.LoadBalancerClient {
	return m.GetClient("oci_load_balancer.LoadBalancerClient").(*oci_load_balancer.LoadBalancerClient)
}
