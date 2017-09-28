// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DBHomesDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDBHomes,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": {
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
			"db_homes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DBHomeDatasource(),
			},
		},
	}
}

func readDBHomes(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &DBHomesDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

type DBHomesDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDBHomes
}

func (s *DBHomesDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	dbSystemID := s.D.Get("db_system_id").(string)

	opts := &baremetal.ListOptions{}
	options.SetPageOptions(s.D, &opts.PageListOptions)
	options.SetLimitOptions(s.D, &opts.LimitListOptions)

	s.Res = &baremetal.ListDBHomes{}

	for {
		var list *baremetal.ListDBHomes
		if list, e = s.Client.ListDBHomes(
			compartmentID, dbSystemID, opts,
		); e != nil {
			break
		}

		s.Res.DBHomes = append(s.Res.DBHomes, list.DBHomes...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *DBHomesDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.DBHomes {
			res := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"db_system_id":   v.DBSystemID,
				"display_name":   v.DisplayName,
				"id":             v.ID,
				"state":          v.State,
				"time_created":   v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("db_homes", resources)
	}
	return
}
