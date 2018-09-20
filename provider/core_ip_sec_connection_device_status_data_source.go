// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func IpSecConnectionDeviceStatusDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularIpSecConnectionDeviceStatus,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"ipsec_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						// @Deprecated 01/2018: time_state_modifed => time_state_modified
						"time_state_modifed": {
							Deprecated: FieldDeprecatedForAnother("time_state_modifed", "time_state_modified"),
							Type:       schema.TypeString,
							Computed:   true,
						},
						"time_state_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularIpSecConnectionDeviceStatus(d *schema.ResourceData, m interface{}) error {
	sync := &IpSecConnectionDeviceStatusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

type IpSecConnectionDeviceStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetIPSecConnectionDeviceStatusResponse
}

func (s *IpSecConnectionDeviceStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IpSecConnectionDeviceStatusDataSourceCrud) Get() error {
	request := oci_core.GetIPSecConnectionDeviceStatusRequest{}

	if ipsecId, ok := s.D.GetOkExists("ipsec_id"); ok {
		tmp := ipsecId.(string)
		request.IpscId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetIPSecConnectionDeviceStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IpSecConnectionDeviceStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	tunnels := []map[string]interface{}{}
	for _, item := range s.Res.Tunnels {
		tunnels = append(tunnels, TunnelStatusToMap(item))
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		tunnels = ApplyFilters(f.(*schema.Set), tunnels, IpSecConnectionDeviceStatusDataSource().Schema["tunnels"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tunnels", tunnels); err != nil {
		return err
	}

	return nil
}

func TunnelStatusToMap(obj oci_core.TunnelStatus) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeStateModified != nil {
		// @Deprecated 01/2018: time_state_modifed => time_state_modified
		result["time_state_modifed"] = obj.TimeStateModified.String()

		result["time_state_modified"] = obj.TimeStateModified.String()
	}

	return result
}
