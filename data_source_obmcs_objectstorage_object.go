// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"strconv"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ObjectDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readObjects,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"objects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     resourceObjectSummary(),
			},
		},
	}
}

func resourceObjectSummary() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"md5": {
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

func readObjects(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	reader := &ObjectDatasourceCrud{}
	reader.D = d
	reader.Client = client

	return crud.ReadResource(reader)
}

type ObjectDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListObjects
}

func (s *ObjectDatasourceCrud) Get() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)

	opts := &baremetal.ListObjectsOptions{
		Fields: "name,size,md5,timeCreated",
	}

	if prefix, ok := s.D.GetOk("prefix"); ok {
		opts.Prefix = prefix.(string)
	}
	if start, ok := s.D.GetOk("start"); ok {
		opts.Start = start.(string)
	}
	if end, ok := s.D.GetOk("end"); ok {
		opts.End = end.(string)
	}
	if limit, ok := s.D.GetOk("limit"); ok {
		opts.Limit = uint64(limit.(int))
	}

	s.Res = &baremetal.ListObjects{Objects: []baremetal.ObjectSummary{}}

	for {
		var list *baremetal.ListObjects
		if list, e = s.Client.ListObjects(baremetal.Namespace(namespace), bucket, opts); e != nil {
			break
		}

		s.Res.Objects = append(s.Res.Objects, list.Objects...)

		if list.NextStartWith == "" {
			break
		}

		opts.Start = list.NextStartWith
	}

	return
}

func (s *ObjectDatasourceCrud) SetData() {

	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Objects {
			res := map[string]interface{}{
				"name":         v.Name,
				"size":         strconv.FormatUint(v.Size, 10),
				"md5":          v.MD5,
				"time_created": v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("objects", resources)
	}
	return
}
