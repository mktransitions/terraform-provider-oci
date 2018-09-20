// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

func PreauthenticatedRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readPreauthenticatedRequests,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"object_name_prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"preauthenticated_requests": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(PreauthenticatedRequestResource()),
			},
		},
	}
}

func readPreauthenticatedRequests(d *schema.ResourceData, m interface{}) error {
	sync := &PreauthenticatedRequestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return ReadResource(sync)
}

type PreauthenticatedRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.ListPreauthenticatedRequestsResponse
}

func (s *PreauthenticatedRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PreauthenticatedRequestsDataSourceCrud) Get() error {
	request := oci_object_storage.ListPreauthenticatedRequestsRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if objectNamePrefix, ok := s.D.GetOkExists("object_name_prefix"); ok {
		tmp := objectNamePrefix.(string)
		request.ObjectNamePrefix = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "object_storage")

	response, err := s.Client.ListPreauthenticatedRequests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPreauthenticatedRequests(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *PreauthenticatedRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		preauthenticatedRequest := map[string]interface{}{}

		preauthenticatedRequest["access_type"] = r.AccessType

		if r.Id != nil {
			preauthenticatedRequest["id"] = *r.Id
		}

		if r.Name != nil {
			preauthenticatedRequest["name"] = *r.Name
		}

		if r.ObjectName != nil {
			preauthenticatedRequest["object"] = *r.ObjectName
		}

		if r.TimeCreated != nil {
			preauthenticatedRequest["time_created"] = r.TimeCreated.String()
		}

		if r.TimeExpires != nil {
			preauthenticatedRequest["time_expires"] = r.TimeExpires.String()
		}

		resources = append(resources, preauthenticatedRequest)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, PreauthenticatedRequestsDataSource().Schema["preauthenticated_requests"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("preauthenticated_requests", resources); err != nil {
		return err
	}

	return nil
}
