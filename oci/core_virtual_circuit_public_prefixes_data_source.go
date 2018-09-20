// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func VirtualCircuitPublicPrefixesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVirtualCircuitPublicPrefixes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"verification_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_circuit_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_circuit_public_prefixes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cidr_block": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"verification_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readVirtualCircuitPublicPrefixes(d *schema.ResourceData, m interface{}) error {
	sync := &VirtualCircuitPublicPrefixesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

type VirtualCircuitPublicPrefixesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListVirtualCircuitPublicPrefixesResponse
}

func (s *VirtualCircuitPublicPrefixesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VirtualCircuitPublicPrefixesDataSourceCrud) Get() error {
	request := oci_core.ListVirtualCircuitPublicPrefixesRequest{}

	if verificationState, ok := s.D.GetOkExists("verification_state"); ok {
		request.VerificationState = oci_core.VirtualCircuitPublicPrefixVerificationStateEnum(verificationState.(string))
	}

	if virtualCircuitId, ok := s.D.GetOkExists("virtual_circuit_id"); ok {
		tmp := virtualCircuitId.(string)
		request.VirtualCircuitId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListVirtualCircuitPublicPrefixes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *VirtualCircuitPublicPrefixesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		virtualCircuitPublicPrefix := map[string]interface{}{}

		if r.CidrBlock != nil {
			virtualCircuitPublicPrefix["cidr_block"] = *r.CidrBlock
		}

		virtualCircuitPublicPrefix["verification_state"] = string(r.VerificationState)

		resources = append(resources, virtualCircuitPublicPrefix)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, VirtualCircuitPublicPrefixesDataSource().Schema["virtual_circuit_public_prefixes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("virtual_circuit_public_prefixes", resources); err != nil {
		return err
	}

	return nil
}
