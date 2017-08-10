// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"log"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

func LoadBalancerBackendSetResource() *schema.Resource {
	return &schema.Resource{
		Create: createLoadBalancerBackendSet,
		Read:   readLoadBalancerBackendSet,
		Update: updateLoadBalancerBackendSet,
		Delete: deleteLoadBalancerBackendSet,
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_checker":                    HealthCheckerSchema,
			"ssl_configuration":                 SSLConfigSchema,
			"session_persistence_configuration": SessionPersistenceConfigSchema,
			"backend": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     LoadBalancerBackendResource(),
			},
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerBackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.CreateResource(d, sync)
}

func readLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerBackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.ReadResource(sync)
}

func updateLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerBackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.UpdateResource(d, sync)
}

func deleteLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerBackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.DeleteResource(d, sync)
}

type LoadBalancerBackendSetResourceCrud struct {
	crud.BaseCrud
	WorkRequest  *baremetal.WorkRequest
	Resource     *baremetal.BackendSet
	ResourceName string
}

func (s *LoadBalancerBackendSetResourceCrud) ID() string {
	id, workSuccess := crud.LoadBalancerResourceID(s.Resource, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		return s.D.Get("name").(string)
	}
	return ""
}

func (s *LoadBalancerBackendSetResourceCrud) CreatedPending() []string {
	return []string{
		baremetal.ResourceWaitingForWorkRequest,
		baremetal.WorkRequestInProgress,
		baremetal.WorkRequestAccepted,
	}
}

func (s *LoadBalancerBackendSetResourceCrud) CreatedTarget() []string {
	return []string{
		baremetal.ResourceSucceededWorkRequest,
		baremetal.WorkRequestSucceeded,
		baremetal.ResourceFailed,
	}
}

func (s *LoadBalancerBackendSetResourceCrud) DeletedPending() []string {
	return []string{
		baremetal.ResourceWaitingForWorkRequest,
		baremetal.WorkRequestInProgress,
		baremetal.WorkRequestAccepted,
	}
}

func (s *LoadBalancerBackendSetResourceCrud) DeletedTarget() []string {
	return []string{
		baremetal.ResourceSucceededWorkRequest,
		baremetal.WorkRequestSucceeded,
	}
}

func (s *LoadBalancerBackendSetResourceCrud) Create() (e error) {
	workReqID, e := s.Client.CreateBackendSet(
		s.D.Get("load_balancer_id").(string),
		s.D.Get("name").(string),
		s.D.Get("policy").(string),
		s.backends(),
		s.healthChecker(),
		s.sslConfig(),
		s.sessionPersistenceConfig(),
		nil,
	)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}

func (s *LoadBalancerBackendSetResourceCrud) Get() (e error) {
	_, stillWorking, err := crud.LoadBalancerResourceGet(s.BaseCrud, s.WorkRequest)
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	s.Resource, e = s.Client.GetBackendSet(
		s.D.Get("load_balancer_id").(string),
		s.D.Get("name").(string),
		nil,
	)
	return
}

func (s *LoadBalancerBackendSetResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateLoadBalancerBackendSetOptions{}

	opts.HealthChecker = s.healthChecker()
	opts.SSLConfig = s.sslConfig()
	opts.Policy = s.D.Get("policy").(string)

	// This is hacky and a race condition, but works for now. Ideally backends are not a required parameter to a backendset update
	bes, err := s.Client.GetBackendSet(s.D.Get("load_balancer_id").(string), s.D.Id(), nil)
	if err != nil {
		return err
	}
	opts.Backends = bes.Backends

	log.Printf("BACKENDS: %v\n", opts.Backends)
	var workReqID string
	workReqID, e = s.Client.UpdateBackendSet(s.D.Get("load_balancer_id").(string), s.D.Id(), opts)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	if e != nil {
		return
	}
	e = crud.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest)
	if e != nil {
		return
	}
	return s.Get()
}

func (s *LoadBalancerBackendSetResourceCrud) SetData() {
	if s.Resource == nil {
		return
	}
	s.D.Set("policy", s.Resource.Policy)
	s.D.Set("name", s.Resource.Name)
	if s.Resource.HealthChecker != nil {
		s.D.Set("health_checker", map[string]interface{}{
			"interval_ms":         s.Resource.HealthChecker.IntervalInMS,
			"port":                s.Resource.HealthChecker.Port,
			"protocol":            s.Resource.HealthChecker.Protocol,
			"response_body_regex": s.Resource.HealthChecker.ResponseBodyRegex,
			"url_path":            s.Resource.HealthChecker.URLPath,
		})
	}

	if s.Resource.SSLConfig != nil {
		s.D.Set("ssl_configuration", map[string]interface{}{
			"certificate_name":        s.Resource.SSLConfig.CertificateName,
			"verify_depth":            s.Resource.SSLConfig.VerifyDepth,
			"verify_peer_certificate": s.Resource.SSLConfig.VerifyPeerCertificate,
		})
	}

	if s.Resource.SessionPersistenceConfig != nil {
		s.D.Set("session_persistence_configuration", map[string]interface{}{
			"cookie_name":      s.Resource.SessionPersistenceConfig.CookieName,
			"disable_fallback": s.Resource.SessionPersistenceConfig.DisableFallback,
		})
	}

	backends := make([]map[string]interface{}, len(s.Resource.Backends))
	for i, v := range s.Resource.Backends {
		backends[i] = map[string]interface{}{
			"backup":     v.Backup,
			"drain":      v.Drain,
			"ip_address": v.IPAddress,
			"offline":    v.Offline,
			"port":       v.Port,
			"weight":     v.Weight,
		}
	}
	s.D.Set("backend", backends)
}

func (s *LoadBalancerBackendSetResourceCrud) Delete() (e error) {
	var workReqID string
	workReqID, e = s.Client.DeleteBackendSet(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string), nil)
	if e != nil {
		return
	}
	s.D.SetId(workReqID)
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}

func (s *LoadBalancerBackendSetResourceCrud) sslConfig() (sslConfig *baremetal.SSLConfiguration) {
	vs := s.D.Get("ssl_configuration").([]interface{})
	if len(vs) == 1 {
		sslConfig = new(baremetal.SSLConfiguration)
		v := vs[0].(map[string]interface{})
		sslConfig.CertificateName = v["certificate_name"].(string)
		sslConfig.VerifyDepth = v["verify_depth"].(int)
		sslConfig.VerifyPeerCertificate = v["verify_peer_certificate"].(bool)
		return sslConfig
	}

	return nil
}

func (s *LoadBalancerBackendSetResourceCrud) sessionPersistenceConfig() (sessionPersistenceConfig *baremetal.SessionPersistenceConfiguration) {
	vs := s.D.Get("session_persistence_configuration").([]interface{})
	if len(vs) == 1 {
		sessionPersistenceConfig = new(baremetal.SessionPersistenceConfiguration)
		v := vs[0].(map[string]interface{})
		sessionPersistenceConfig.CookieName = v["cookie_name"].(string)
		sessionPersistenceConfig.DisableFallback = v["disable_fallback"].(bool)
		return sessionPersistenceConfig
	}

	return nil
}

func (s *LoadBalancerBackendSetResourceCrud) healthChecker() *baremetal.HealthChecker {

	vs := s.D.Get("health_checker").([]interface{})
	if len(vs) == 1 {
		healthChecker := new(baremetal.HealthChecker)
		v := vs[0].(map[string]interface{})
		healthChecker.IntervalInMS = v["interval_ms"].(int)
		healthChecker.Port = v["port"].(int)
		healthChecker.Protocol = v["protocol"].(string)
		healthChecker.ResponseBodyRegex = v["response_body_regex"].(string)
		healthChecker.URLPath = v["url_path"].(string)
		return healthChecker
	}
	return nil
}
func (s *LoadBalancerBackendSetResourceCrud) backends() []baremetal.Backend {
	vs := s.D.Get("backend").([]interface{})
	backends := make([]baremetal.Backend, len(vs))
	for i := range vs {
		v := vs[i].(map[string]interface{})
		backends[i] = baremetal.Backend{
			Backup:    v["backup"].(bool),
			Drain:     v["drain"].(bool),
			IPAddress: v["ip_address"].(string),
			Offline:   v["offline"].(bool),
			Port:      v["port"].(int),
			Weight:    v["weight"].(int),
		}
	}
	return backends
}
