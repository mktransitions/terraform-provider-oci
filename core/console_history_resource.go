// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func ConsoleHistoryResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: crud.DefaultTimeout,
		Create:   createConsoleHistory,
		Read:     readConsoleHistory,
		Delete:   deleteConsoleHistory,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createConsoleHistory(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	ichCrud := &ConsoleHistoryResourceCrud{}
	ichCrud.D = d
	ichCrud.Client = client
	return crud.CreateResource(d, ichCrud)
}

func readConsoleHistory(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	ichCrud := &ConsoleHistoryResourceCrud{}
	ichCrud.D = d
	ichCrud.Client = client
	return crud.ReadResource(ichCrud)
}

func deleteConsoleHistory(d *schema.ResourceData, m interface{}) (e error) {
	return fmt.Errorf("console history resource: console history %v cannot be deleted", d.Id())
}
