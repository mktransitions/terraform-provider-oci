// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func LoadBalancerShapeDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerShapes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readLoadBalancerShapes(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &LoadBalancerShapeDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

type LoadBalancerShapeDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListLoadBalancerShapes
}

func (s *LoadBalancerShapeDatasourceCrud) Get() (e error) {
	cID := s.D.Get("compartment_id").(string)
	s.Res, e = s.Client.ListLoadBalancerShapes(cID, nil)
	return
}

func (s *LoadBalancerShapeDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}

		for _, v := range s.Res.LoadBalancerShapes {
			res := map[string]interface{}{
				"name": v.Name,
			}
			resources = append(resources, res)

		}

		if f, fOk := s.D.GetOk("filter"); fOk {
			resources = ApplyFilters(f.(*schema.Set), resources)
		}

		s.D.Set("shapes", resources)
	}
	return
}
