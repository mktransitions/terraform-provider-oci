// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func BucketResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createBucket,
		Read:     readBucket,
		Update:   updateBucket,
		Delete:   deleteBucket,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"access_type": {
				Type:     schema.TypeString,
				Computed: false,
				Default:  baremetal.NoPublicAccess,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(baremetal.NoPublicAccess),
					string(baremetal.ObjectRead)}, true),
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func createBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.CreateResource(d, sync)
}

func readBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.ReadResource(sync)
}

func updateBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.UpdateResource(d, sync)
}

func deleteBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).clientWithoutNotFoundRetries
	return crud.DeleteResource(d, sync)
}

type BucketResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.Bucket
}

func (s *BucketResourceCrud) ID() string {
	return string(s.Res.Namespace) + "/" + s.Res.Name
}

// todo: this delay doesnt seem necessary but should be well tested in real world scenarios before removal
func (s *BucketResourceCrud) ExtraWaitPostCreateDelete() time.Duration {
	return time.Duration(10 * time.Second)
}

func (s *BucketResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("name", s.Res.Name)
	s.D.Set("namespace", s.Res.Namespace)
	s.D.Set("metadata", s.Res.Metadata)
	s.D.Set("created_by", s.Res.CreatedBy)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("accessType", s.Res.AccessType)
}

func (s *BucketResourceCrud) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	name := s.D.Get("name").(string)
	namespace := s.D.Get("namespace").(string)
	opts := &baremetal.CreateBucketOptions{}

	if rawMetadata, ok := s.D.GetOk("metadata"); ok {
		metadata := resourceObjectStorageMapToMetadata(rawMetadata.(map[string]interface{}))
		opts.Metadata = metadata
	}

	accessType, _ := s.D.GetOk("access_type") //guaranteed to be there with Default value
	opts.AccessType = baremetal.BucketAccessType(accessType.(string))
	s.Res, e = s.Client.CreateBucket(compartmentID, name, baremetal.Namespace(namespace), opts)
	return
}

func (s *BucketResourceCrud) Get() (e error) {
	name := s.D.Get("name").(string)
	namespace := s.D.Get("namespace").(string)
	res, e := s.Client.GetBucket(name, baremetal.Namespace(namespace))
	if e == nil {
		s.Res = res
	}
	return
}

func (s *BucketResourceCrud) Update() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	name := s.D.Get("name").(string)
	namespace := s.D.Get("namespace").(string)
	opts := &baremetal.UpdateBucketOptions{}
	if rawMetadata, ok := s.D.GetOk("metadata"); ok {
		metadata := resourceObjectStorageMapToMetadata(rawMetadata.(map[string]interface{}))
		opts.Metadata = metadata
	}

	accessType, _ := s.D.GetOk("access_type") //guaranteed to be there with Default value
	opts.AccessType = baremetal.BucketAccessType(accessType.(string))
	s.Res, e = s.Client.UpdateBucket(compartmentID, name, baremetal.Namespace(namespace), opts)
	return
}

func (s *BucketResourceCrud) Delete() (e error) {
	name := s.D.Get("name").(string)
	namespace := s.D.Get("namespace").(string)
	return s.Client.DeleteBucket(name, baremetal.Namespace(namespace), nil)
}
