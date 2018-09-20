// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func VirtualCircuitsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVirtualCircuits,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_circuits": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(VirtualCircuitDataSource()),
			},
		},
	}
}

func readVirtualCircuits(d *schema.ResourceData, m interface{}) error {
	sync := &VirtualCircuitsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

type VirtualCircuitsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListVirtualCircuitsResponse
}

func (s *VirtualCircuitsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VirtualCircuitsDataSourceCrud) Get() error {
	request := oci_core.ListVirtualCircuitsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.VirtualCircuitLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListVirtualCircuits(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVirtualCircuits(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *VirtualCircuitsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		virtualCircuit := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.BandwidthShapeName != nil {
			virtualCircuit["bandwidth_shape_name"] = *r.BandwidthShapeName
		}

		virtualCircuit["bgp_management"] = r.BgpManagement

		virtualCircuit["bgp_session_state"] = r.BgpSessionState

		crossConnectMappings := []interface{}{}
		for _, item := range r.CrossConnectMappings {
			crossConnectMappings = append(crossConnectMappings, CrossConnectMappingToMap(item))
		}
		virtualCircuit["cross_connect_mappings"] = crossConnectMappings

		if r.CustomerBgpAsn != nil {
			virtualCircuit["customer_bgp_asn"] = *r.CustomerBgpAsn
		}

		if r.DisplayName != nil {
			virtualCircuit["display_name"] = *r.DisplayName
		}

		if r.GatewayId != nil {
			virtualCircuit["gateway_id"] = *r.GatewayId
		}

		if r.Id != nil {
			virtualCircuit["id"] = *r.Id
		}

		if r.OracleBgpAsn != nil {
			virtualCircuit["oracle_bgp_asn"] = *r.OracleBgpAsn
		}

		if r.ProviderServiceId != nil {
			virtualCircuit["provider_service_id"] = *r.ProviderServiceId
		}

		virtualCircuit["provider_state"] = r.ProviderState

		publicPrefixes := []interface{}{}
		for _, item := range r.PublicPrefixes {
			publicPrefixes = append(publicPrefixes, CreateVirtualCircuitPublicPrefixDetailsToMap(item))
		}
		virtualCircuit["public_prefixes"] = publicPrefixes

		if r.ReferenceComment != nil {
			virtualCircuit["reference_comment"] = *r.ReferenceComment
		}

		if r.Region != nil {
			virtualCircuit["region"] = *r.Region
		}

		virtualCircuit["service_type"] = r.ServiceType

		virtualCircuit["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			virtualCircuit["time_created"] = r.TimeCreated.String()
		}

		virtualCircuit["type"] = r.Type

		resources = append(resources, virtualCircuit)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, VirtualCircuitsDataSource().Schema["virtual_circuits"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("virtual_circuits", resources); err != nil {
		return err
	}

	return nil
}
