// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/crud"
	"github.com/oracle/terraform-provider-baremetal/options"
)

func VolumeDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readVolumes,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volumes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     VolumeResource(),
			},
		},
	}
}

func readVolumes(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &VolumeDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type VolumeDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListVolumes
}

func (s *VolumeDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListVolumesOptions{}
	options.SetListOptions(s.D, &opts.ListOptions)
	if val, ok := s.D.GetOk("availability_domain"); ok {
		opts.AvailabilityDomain = val.(string)
	}

	s.Res = &baremetal.ListVolumes{Volumes: []baremetal.Volume{}}

	for {
		var list *baremetal.ListVolumes
		if list, e = s.Client.ListVolumes(compartmentID, opts); e != nil {
			break
		}

		s.Res.Volumes = append(s.Res.Volumes, list.Volumes...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *VolumeDatasourceCrud) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		volumes := []map[string]interface{}{}
		for _, v := range s.Res.Volumes {
			vol := map[string]interface{}{
				"availability_domain": v.AvailabilityDomain,
				"compartment_id":      v.CompartmentID,
				"display_name":        v.DisplayName,
				"id":                  v.ID,
				"size_in_mbs":         v.SizeInMBs,
				"state":               v.State,
				"time_created":        v.TimeCreated.String(),
			}
			volumes = append(volumes, vol)
		}
		s.D.Set("volumes", volumes)
	}
	return
}
