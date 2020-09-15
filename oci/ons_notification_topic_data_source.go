// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_ons "github.com/oracle/oci-go-sdk/v25/ons"
)

func init() {
	RegisterDatasource("oci_ons_notification_topic", OnsNotificationTopicDataSource())
}

func OnsNotificationTopicDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["topic_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(OnsNotificationTopicResource(), fieldMap, readSingularOnsNotificationTopic)
}

func readSingularOnsNotificationTopic(d *schema.ResourceData, m interface{}) error {
	sync := &OnsNotificationTopicDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).notificationControlPlaneClient()

	return ReadResource(sync)
}

type OnsNotificationTopicDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ons.NotificationControlPlaneClient
	Res    *oci_ons.GetTopicResponse
}

func (s *OnsNotificationTopicDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnsNotificationTopicDataSourceCrud) Get() error {
	request := oci_ons.GetTopicRequest{}

	if topicId, ok := s.D.GetOkExists("topic_id"); ok {
		tmp := topicId.(string)
		request.TopicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "ons")

	response, err := s.Client.GetTopic(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OnsNotificationTopicDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	if s.Res.TopicId != nil {
		s.D.Set("topic_id", *s.Res.TopicId)
		s.D.SetId(*s.Res.TopicId)
	}

	if s.Res.ApiEndpoint != nil {
		s.D.Set("api_endpoint", *s.Res.ApiEndpoint)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Etag != nil {
		s.D.Set("etag", *s.Res.Etag)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
