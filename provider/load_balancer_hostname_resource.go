// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func HostnameResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createHostname,
		Read:     readHostname,
		Update:   updateHostname,
		Delete:   deleteHostname,
		Schema: map[string]*schema.Schema{
			// Required
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
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

			// Optional

			// Computed
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createHostname(d *schema.ResourceData, m interface{}) error {
	sync := &HostnameResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return CreateResource(d, sync)
}

func readHostname(d *schema.ResourceData, m interface{}) error {
	sync := &HostnameResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return ReadResource(sync)
}

func updateHostname(d *schema.ResourceData, m interface{}) error {
	sync := &HostnameResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return UpdateResource(d, sync)
}

func deleteHostname(d *schema.ResourceData, m interface{}) error {
	sync := &HostnameResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type HostnameResourceCrud struct {
	BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.Hostname
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *HostnameResourceCrud) ID() string {
	if s.WorkRequest != nil {
		if s.WorkRequest.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
			return getHostnameCompositeId(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
		} else {
			return *s.WorkRequest.Id
		}
	}
	return ""
}

func (s *HostnameResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *HostnameResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *HostnameResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *HostnameResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *HostnameResourceCrud) Create() error {
	request := oci_load_balancer.CreateHostnameRequest{}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		request.Hostname = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateHostname(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *HostnameResourceCrud) Get() error {
	_, stillWorking, err := LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	request := oci_load_balancer.GetHostnameRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if !strings.HasPrefix(s.D.Id(), "ocid1.loadbalancerworkrequest.") {
		loadBalancerId, name, err := parseHostnameCompositeId(s.D.Id())
		if err == nil {
			request.LoadBalancerId = &loadBalancerId
			request.Name = &name
		} else {
			return err
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetHostname(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Hostname
	return nil
}

func (s *HostnameResourceCrud) Update() error {
	request := oci_load_balancer.UpdateHostnameRequest{}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		request.Hostname = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateHostname(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *HostnameResourceCrud) Delete() error {
	if strings.Contains(s.D.Id(), "ocid1.loadbalancerworkrequest") {
		return nil
	}
	request := oci_load_balancer.DeleteHostnameRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteHostname(context.Background(), request)

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *HostnameResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	loadBalancerId, name, err := parseHostnameCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("load_balancer_id", &loadBalancerId)
		s.D.Set("name", &name)
	} else {
		return err
	}

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	return nil
}

func getHostnameCompositeId(loadBalancerId string, name string) string {
	loadBalancerId = url.PathEscape(loadBalancerId)
	name = url.PathEscape(name)
	compositeId := "loadBalancers/" + loadBalancerId + "/hostnames/" + name
	return compositeId
}

func parseHostnameCompositeId(compositeId string) (loadBalancerId string, name string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("loadBalancers/.*/hostnames/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	loadBalancerId, _ = url.PathUnescape(parts[1])
	name, _ = url.PathUnescape(parts[3])

	return
}
