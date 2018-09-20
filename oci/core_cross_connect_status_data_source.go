// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CrossConnectStatusDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCrossConnectStatus,
		Schema: map[string]*schema.Schema{
			"cross_connect_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			// @CODEGEN 07/2018: Remove duplicated fields in computed that are also required
			"interface_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"light_level_ind_bm": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"light_level_indicator": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCrossConnectStatus(d *schema.ResourceData, m interface{}) error {
	sync := &CrossConnectStatusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

type CrossConnectStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetCrossConnectStatusResponse
}

func (s *CrossConnectStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CrossConnectStatusDataSourceCrud) Get() error {
	request := oci_core.GetCrossConnectStatusRequest{}

	if crossConnectId, ok := s.D.GetOkExists("cross_connect_id"); ok {
		tmp := crossConnectId.(string)
		request.CrossConnectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetCrossConnectStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CrossConnectStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	s.D.Set("interface_state", s.Res.InterfaceState)

	if s.Res.LightLevelIndBm != nil {
		s.D.Set("light_level_ind_bm", *s.Res.LightLevelIndBm)
	}

	s.D.Set("light_level_indicator", s.Res.LightLevelIndicator)

	return nil
}
