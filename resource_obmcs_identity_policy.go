// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"crypto/md5"
	"encoding/hex"
	"strings"

	"github.com/oracle/terraform-provider-oci/crud"
)

func PolicyResource() *schema.Resource {
	policySchema := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"description": {
			Type:     schema.TypeString,
			Required: true,
		},
		"compartment_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"state": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"time_created": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"time_modified": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"statements": {
			Type:             schema.TypeList,
			Required:         true,
			DiffSuppressFunc: ignorePolicyFormatDiff,
			Elem:             &schema.Schema{Type: schema.TypeString},
		},
		"ETag": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"policyHash": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"lastUpdateETag": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"inactive_state": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"version_date": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createPolicy,
		Read:     readPolicy,
		Update:   updatePolicy,
		Delete:   deletePolicy,
		Schema:   policySchema,
	}
}

func ignorePolicyFormatDiff(k string, old string, new string, d *schema.ResourceData) bool {
	oldHash := getOrDefault(d, "policyHash", "")
	newHash := getMD5Hash(toStringArray(d.Get("statements")))
	oldETag := getOrDefault(d, "lastUpdateETag", "")
	currentETag := getOrDefault(d, "ETag", "")
	suppressDiff := strings.EqualFold(oldHash, newHash) && strings.EqualFold(oldETag, currentETag)

	return suppressDiff
}

func getOrDefault(d *schema.ResourceData, key string, defaultValue string) string {
	valueString := defaultValue
	if value, ok := d.GetOk(key); ok {
		valueString = value.(string)
	}

	return valueString
}

func createPolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &PolicyResourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.CreateResource(d, sync)
}

func readPolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &PolicyResourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

func updatePolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &PolicyResourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.UpdateResource(d, sync)
}

func deletePolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &PolicyResourceCrud{}
	sync.D = d
	sync.Client = client.clientWithoutNotFoundRetries
	return sync.Delete()
}

type PolicyResourceCrud struct {
	*crud.IdentitySync
	crud.BaseCrud
	Res *baremetal.Policy
}

func (s *PolicyResourceCrud) ID() string {
	return s.Res.ID
}

func (s *PolicyResourceCrud) State() string {
	return s.Res.State
}

func (s *PolicyResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *PolicyResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceActive}
}

func (s *PolicyResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDeleting}
}

func (s *PolicyResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDeleted}
}

func toStringArray(vals interface{}) []string {
	arr := vals.([]interface{})
	result := []string{}
	for _, val := range arr {
		result = append(result, val.(string))
	}
	return result
}

func (s *PolicyResourceCrud) Create() (e error) {
	name := s.D.Get("name").(string)
	description := s.D.Get("description").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	statements := toStringArray(s.D.Get("statements"))

	s.Res, e = s.Client.CreatePolicy(name, description, compartmentID, statements, nil)

	if e == nil {
		s.D.Set("policyHash", getMD5Hash(statements))
		s.D.Set("lastUpdateETag", s.Res.ETag)
	}

	return
}

func (s *PolicyResourceCrud) Get() (e error) {
	res, e := s.Client.GetPolicy(s.D.Id())
	if e == nil {
		s.Res = res
	}
	return
}

func (s *PolicyResourceCrud) Update() (e error) {
	opts := &baremetal.UpdatePolicyOptions{}
	if description, ok := s.D.GetOk("description"); ok {
		opts.Description = description.(string)
	}

	policyHash := ""
	if rawStatements, ok := s.D.GetOk("statements"); ok {
		statements := toStringArray(rawStatements)
		opts.Statements = statements
		policyHash = getMD5Hash(statements)
	}

	s.Res, e = s.Client.UpdatePolicy(s.D.Id(), opts)
	if e == nil {
		s.D.Set("policyHash", policyHash)
		s.D.Set("lastUpdateETag", s.Res.ETag)
	}
	return
}

func (s *PolicyResourceCrud) SetData() {
	s.D.Set("statements", s.Res.Statements)
	s.D.Set("ETag", s.Res.ETag)
	s.D.Set("name", s.Res.Name)
	s.D.Set("description", s.Res.Description)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func getMD5Hash(values []string) string {
	hash := md5.Sum([]byte(strings.Join(values, "#")))
	return hex.EncodeToString(hash[:])
}

func (s *PolicyResourceCrud) Delete() (e error) {
	return s.Client.DeletePolicy(s.D.Id(), nil)
}
