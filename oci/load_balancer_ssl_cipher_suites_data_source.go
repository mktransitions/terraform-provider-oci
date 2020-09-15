// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v25/loadbalancer"
)

func init() {
	RegisterDatasource("oci_load_balancer_ssl_cipher_suites", LoadBalancerSslCipherSuitesDataSource())
}

func LoadBalancerSslCipherSuitesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerSslCipherSuites,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"load_balancer_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_cipher_suites": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     LoadBalancerSslCipherSuiteResource(),
			},
		},
	}
}

func readLoadBalancerSslCipherSuites(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerSslCipherSuitesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()

	return ReadResource(sync)
}

type LoadBalancerSslCipherSuitesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListSSLCipherSuitesResponse
}

func (s *LoadBalancerSslCipherSuitesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerSslCipherSuitesDataSourceCrud) Get() error {
	request := oci_load_balancer.ListSSLCipherSuitesRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListSSLCipherSuites(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerSslCipherSuitesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		sslCipherSuite := map[string]interface{}{}

		if r.Name != nil {
			sslCipherSuite["name"] = *r.Name
		}

		resources = append(resources, sslCipherSuite)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, LoadBalancerSslCipherSuitesDataSource().Schema["ssl_cipher_suites"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("ssl_cipher_suites", resources); err != nil {
		return err
	}

	return nil
}
