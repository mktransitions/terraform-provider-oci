// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"crypto/md5"
	"encoding/hex"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ObjectResource() *schema.Resource {
	var objectSchema = map[string]*schema.Schema{
		"namespace": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"bucket": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"object": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"content": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
			StateFunc: func(body interface{}) string {
				v := body.(string)
				if v == "" {
					return ""
				}
				h := md5.Sum([]byte(v))
				return hex.EncodeToString(h[:])
			},
		},
		"content_encoding": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"content_language": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"content_length": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"content_md5": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"content_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"metadata": {
			Type:     schema.TypeMap,
			Optional: true,
			ForceNew: true,
		},
	}

	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createObject,
		Read:     readObject,
		Delete:   deleteObject,
		Schema:   objectSchema,
	}
}

func createObject(d *schema.ResourceData, m interface{}) (e error) {
	sync := &ObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.CreateResource(d, sync)
}

func readObject(d *schema.ResourceData, m interface{}) (e error) {
	sync := &ObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.ReadResource(sync)
}

func deleteObject(d *schema.ResourceData, m interface{}) (e error) {
	sync := &ObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).clientWithoutNotFoundRetries
	return crud.DeleteResource(d, sync)
}

type ObjectResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.Object
}

func (s *ObjectResourceCrud) ID() string {
	return "tfobm-object-" + string(s.Res.Namespace) + "/" + s.Res.Bucket + "/" + s.Res.ID
}

func (s *ObjectResourceCrud) SetData() {
	log.Printf("=======================\n%v\n===================", s.Res)
	s.D.Set("namespace", s.Res.Namespace)
	s.D.Set("bucket", s.Res.Bucket)
	s.D.Set("object", s.Res.ID)
	s.D.Set("content", s.Res.Body)
	s.D.Set("metadata", s.Res.Metadata)
	s.D.Set("content_encoding", s.Res.ContentEncoding)
	s.D.Set("content_language", s.Res.ContentLanguage)
	s.D.Set("content_length", s.Res.ContentLength)
	s.D.Set("content_md5", s.Res.ContentMD5)
	s.D.Set("content_type", s.Res.ContentType)
}

func (s *ObjectResourceCrud) Create() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)
	object := s.D.Get("object").(string)
	content := s.D.Get("content").(string)
	opts := &baremetal.PutObjectOptions{}

	if contentEncoding, ok := s.D.GetOk("content_encoding"); ok {
		opts.ContentEncoding = contentEncoding.(string)
	}

	if contentLanguage, ok := s.D.GetOk("content_language"); ok {
		opts.ContentLanguage = contentLanguage.(string)
	}

	if contentType, ok := s.D.GetOk("content_type"); ok {
		opts.ContentType = contentType.(string)
	}

	if rawMetadata, ok := s.D.GetOk("metadata"); ok {
		metadata := resourceObjectStorageMapToMetadata(rawMetadata.(map[string]interface{}))
		opts.Metadata = metadata
	}
	_, e = s.Client.PutObject(baremetal.Namespace(namespace), bucket, object, []byte(content), opts)
	if e == nil {
		s.Res, e = s.Client.GetObject(baremetal.Namespace(namespace), bucket, object, &baremetal.GetObjectOptions{})
	}
	return
}

func (s *ObjectResourceCrud) Get() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)
	object := s.D.Get("object").(string)
	res, e := s.Client.GetObject(baremetal.Namespace(namespace), bucket, object, &baremetal.GetObjectOptions{})
	if e == nil {
		s.Res = res
	}
	return
}

func (s *ObjectResourceCrud) Delete() (e error) {
	namespace := s.D.Get("namespace").(string)
	bucket := s.D.Get("bucket").(string)
	object := s.D.Get("object").(string)
	opts := &baremetal.DeleteObjectOptions{}

	_, e = s.Client.DeleteObject(baremetal.Namespace(namespace), bucket, object, opts)
	return
}
