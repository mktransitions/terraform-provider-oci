// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func ListenerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createListener,
		Read:     readListener,
		Update:   updateListener,
		Delete:   deleteListener,
		Schema: map[string]*schema.Schema{
			// Required
			"default_backend_set_name": {
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
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"connection_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"idle_timeout_in_seconds": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     validateInt64TypeString,
							DiffSuppressFunc: int64StringDiffSuppressFunction,
						},

						// Optional

						// Computed
					},
				},
			},
			"hostname_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"path_route_set_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"certificate_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"verify_depth": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  5,
						},
						"verify_peer_certificate": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},

						// Computed
					},
				},
			},

			// Computed
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createListener(d *schema.ResourceData, m interface{}) error {
	sync := &ListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return CreateResource(d, sync)
}

func readListener(d *schema.ResourceData, m interface{}) error {
	sync := &ListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	return ReadResource(sync)
}

func updateListener(d *schema.ResourceData, m interface{}) error {
	sync := &ListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return UpdateResource(d, sync)
}

func deleteListener(d *schema.ResourceData, m interface{}) error {
	sync := &ListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type ListenerResourceCrud struct {
	BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.Listener
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *ListenerResourceCrud) ID() string {
	if s.WorkRequest != nil {
		if s.WorkRequest.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
			return getListenerCompositeId(s.D.Get("name").(string), s.D.Get("load_balancer_id").(string))
		} else {
			return *s.WorkRequest.Id
		}
	}
	return ""
}

func (s *ListenerResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *ListenerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *ListenerResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *ListenerResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *ListenerResourceCrud) Create() error {
	request := oci_load_balancer.CreateListenerRequest{}

	if connectionConfiguration, ok := s.D.GetOkExists("connection_configuration"); ok {
		if tmpList := connectionConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_configuration", 0)
			tmp, err := s.mapToConnectionConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConnectionConfiguration = &tmp
		}
	}

	if defaultBackendSetName, ok := s.D.GetOkExists("default_backend_set_name"); ok {
		tmp := defaultBackendSetName.(string)
		request.DefaultBackendSetName = &tmp
	}

	request.HostnameNames = []string{}
	if hostnameNames, ok := s.D.GetOkExists("hostname_names"); ok {
		interfaces := hostnameNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			tmp[i] = interfaces[i].(string)
		}
		request.HostnameNames = tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if pathRouteSetName, ok := s.D.GetOkExists("path_route_set_name"); ok {
		tmp := pathRouteSetName.(string)
		request.PathRouteSetName = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(string)
		request.Protocol = &tmp
	}

	if sslConfiguration, ok := s.D.GetOkExists("ssl_configuration"); ok {
		if tmpList := sslConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ssl_configuration", 0)
			tmp, err := s.mapToSSLConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SslConfiguration = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateListener(context.Background(), request)
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

func (s *ListenerResourceCrud) Get() (e error) {
	// key: {workRequestID} || {loadBalancerID,name}
	_, stillWorking, err := LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}

	if !strings.HasPrefix(s.D.Id(), "ocid1.loadbalancerworkrequest.") {
		listenerName, loadBalancerId, err := parseListenerCompositeId(s.D.Id())
		if err == nil {
			s.D.Set("name", &listenerName)
			s.D.Set("load_balancer_id", &loadBalancerId)
		} else {
			return err
		}
	}

	res, e := s.GetListener(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
	if e == nil {
		s.Res = res
	}
	return
}

func (s *ListenerResourceCrud) GetListener(loadBalancerID, name string) (*oci_load_balancer.Listener, error) {
	request := oci_load_balancer.GetLoadBalancerRequest{}
	request.LoadBalancerId = &loadBalancerID
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetLoadBalancer(context.Background(), request)
	if err != nil {
		return nil, err
	}
	lb := &response.LoadBalancer
	if lb != nil && lb.Listeners != nil {
		if l, ok := lb.Listeners[name]; ok {
			if l.Name != nil && *l.Name == name {
				return &l, nil
			}
		}
	}
	return nil, fmt.Errorf("Listener %s on load balancer %s does not exist", name, loadBalancerID)
}

func (s *ListenerResourceCrud) Update() error {
	request := oci_load_balancer.UpdateListenerRequest{}

	if connectionConfiguration, ok := s.D.GetOkExists("connection_configuration"); ok {
		if tmpList := connectionConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_configuration", 0)
			tmp, err := s.mapToConnectionConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConnectionConfiguration = &tmp
		}
	}

	if defaultBackendSetName, ok := s.D.GetOkExists("default_backend_set_name"); ok {
		tmp := defaultBackendSetName.(string)
		request.DefaultBackendSetName = &tmp
	}

	request.HostnameNames = []string{}
	if hostnameNames, ok := s.D.GetOkExists("hostname_names"); ok {
		interfaces := hostnameNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			tmp[i] = interfaces[i].(string)
		}
		request.HostnameNames = tmp
	}
	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.ListenerName = &tmp
	}

	if pathRouteSetName, ok := s.D.GetOkExists("path_route_set_name"); ok {
		tmp := pathRouteSetName.(string)
		request.PathRouteSetName = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(string)
		request.Protocol = &tmp
	}

	if sslConfiguration, ok := s.D.GetOkExists("ssl_configuration"); ok {
		if tmpList := sslConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ssl_configuration", 0)
			tmp, err := s.mapToSSLConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SslConfiguration = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateListener(context.Background(), request)
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

func (s *ListenerResourceCrud) Delete() error {
	if strings.Contains(s.D.Id(), "ocid1.loadbalancerworkrequest") {
		return nil
	}
	request := oci_load_balancer.DeleteListenerRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.ListenerName = &tmp
	}
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteListener(context.Background(), request)

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

func (s *ListenerResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}
	listenerName, loadBalancerId, err := parseListenerCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &listenerName)
		s.D.Set("load_balancer_id", &loadBalancerId)
	} else {
		return err
	}

	if s.Res.ConnectionConfiguration != nil {
		s.D.Set("connection_configuration", []interface{}{ConnectionConfigurationToMap(s.Res.ConnectionConfiguration)})
	} else {
		s.D.Set("connection_configuration", []interface{}{})
	}
	if s.Res.DefaultBackendSetName != nil {
		s.D.Set("default_backend_set_name", *s.Res.DefaultBackendSetName)
	}
	if s.Res.HostnameNames != nil {
		s.D.Set("hostname_names", s.Res.HostnameNames)
	}
	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}
	if s.Res.PathRouteSetName != nil {
		s.D.Set("path_route_set_name", *s.Res.PathRouteSetName)
	}
	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}
	if s.Res.Protocol != nil {
		s.D.Set("protocol", *s.Res.Protocol)
	}
	if s.Res.SslConfiguration != nil {
		s.D.Set("ssl_configuration", []interface{}{SSLConfigurationToMap(s.Res.SslConfiguration)})
	} else {
		s.D.Set("ssl_configuration", []interface{}{})
	}

	return nil
}
func getListenerCompositeId(listenerName string, loadBalancerId string) string {
	listenerName = url.PathEscape(listenerName)
	loadBalancerId = url.PathEscape(loadBalancerId)
	compositeId := "loadBalancers/" + loadBalancerId + "/listeners/" + listenerName
	return compositeId
}

func parseListenerCompositeId(compositeId string) (listenerName string, loadBalancerId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("loadBalancers/.*/listeners/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	loadBalancerId, _ = url.PathUnescape(parts[1])
	listenerName, _ = url.PathUnescape(parts[3])

	return
}

func (s *ListenerResourceCrud) mapToConnectionConfiguration(fieldKeyFormat string) (oci_load_balancer.ConnectionConfiguration, error) {
	result := oci_load_balancer.ConnectionConfiguration{}

	if idleTimeoutInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idle_timeout_in_seconds")); ok {
		tmp := idleTimeoutInSeconds.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert idleTimeoutInSeconds string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.IdleTimeout = &tmpInt64
	}

	return result, nil
}

func ConnectionConfigurationToMap(obj *oci_load_balancer.ConnectionConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IdleTimeout != nil {
		result["idle_timeout_in_seconds"] = strconv.FormatInt(*obj.IdleTimeout, 10)
	}

	return result
}

func (s *ListenerResourceCrud) mapToSSLConfigurationDetails(fieldKeyFormat string) (oci_load_balancer.SslConfigurationDetails, error) {
	result := oci_load_balancer.SslConfigurationDetails{}

	if certificateName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_name")); ok {
		tmp := certificateName.(string)
		result.CertificateName = &tmp
	}

	if verifyDepth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_depth")); ok {
		tmp := verifyDepth.(int)
		result.VerifyDepth = &tmp
	}

	if verifyPeerCertificate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_peer_certificate")); ok {
		tmp := verifyPeerCertificate.(bool)
		result.VerifyPeerCertificate = &tmp
	}

	return result, nil
}

// @CODEGEN 08/2018 - Method SSLConfigurationDetailsToMap is available in load_balancer_backend_set_resource.go
