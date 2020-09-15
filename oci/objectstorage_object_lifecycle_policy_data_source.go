// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v25/objectstorage"
)

func init() {
	RegisterDatasource("oci_objectstorage_object_lifecycle_policy", ObjectStorageObjectLifecyclePolicyDataSource())
}

func ObjectStorageObjectLifecyclePolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["bucket"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["namespace"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(ObjectStorageObjectLifecyclePolicyResource(), fieldMap, readSingularObjectStorageObjectLifecyclePolicy)
}

func readSingularObjectStorageObjectLifecyclePolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectLifecyclePolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient()

	return ReadResource(sync)
}

type ObjectStorageObjectLifecyclePolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.GetObjectLifecyclePolicyResponse
}

func (s *ObjectStorageObjectLifecyclePolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStorageObjectLifecyclePolicyDataSourceCrud) Get() error {
	request := oci_object_storage.GetObjectLifecyclePolicyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "object_storage")

	response, err := s.Client.GetObjectLifecyclePolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ObjectStorageObjectLifecyclePolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	rules := []interface{}{}
	for _, item := range s.Res.Items {
		var objectLifecycleRuleMap = ObjectLifecycleRuleToMap(item)
		fixupObjectNameFilterInclusionPrefixesAsList(objectLifecycleRuleMap)
		rules = append(rules, objectLifecycleRuleMap)
	}
	s.D.Set("rules", rules)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func fixupObjectNameFilterInclusionPrefixesAsList(objectLifecycleRuleMap map[string]interface{}) {
	if objectNameFilterList, exists := objectLifecycleRuleMap["object_name_filter"]; exists {
		if objectNameFilterList, ok := objectNameFilterList.([]interface{}); ok && len(objectNameFilterList) > 0 {
			firstElement := objectNameFilterList[0]

			if objectNameFilterMap, ok := firstElement.(map[string]interface{}); ok {
				if inclusionPrefixesSet, ok := objectNameFilterMap["inclusion_prefixes"].(*schema.Set); ok {
					objectNameFilterMap["inclusion_prefixes"] = inclusionPrefixesSet.List()
				}

				if inclusionPatternsSet, ok := objectNameFilterMap["inclusion_patterns"].(*schema.Set); ok {
					objectNameFilterMap["inclusion_patterns"] = inclusionPatternsSet.List()
				}

				if exclusionPatternsSet, ok := objectNameFilterMap["exclusion_patterns"].(*schema.Set); ok {
					objectNameFilterMap["exclusion_patterns"] = exclusionPatternsSet.List()
				}
			}

		}

	}
}
