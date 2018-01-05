// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

var transportSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"source_port_range": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"min": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
			"max": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	},
}

var icmpSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"code": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	},
}

func DefaultSecurityListResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: crud.ImportDefaultResource,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createSecurityList,
		Read:     readSecurityList,
		Update:   updateSecurityList,
		Delete:   deleteSecurityList,
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"egress_security_rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:     schema.TypeString,
							Required: true,
						},
						"icmp_options": icmpSchema,
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tcp_options": transportSchema,
						"udp_options": transportSchema,
						"stateless": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"manage_default_resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ingress_security_rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_options": icmpSchema,
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"source": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tcp_options": transportSchema,
						"udp_options": transportSchema,
						"stateless": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
			"state": {
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

func SecurityListResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createSecurityList,
		Read:     readSecurityList,
		Update:   updateSecurityList,
		Delete:   deleteSecurityList,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"egress_security_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:     schema.TypeString,
							Required: true,
						},
						"icmp_options": icmpSchema,
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tcp_options": transportSchema,
						"udp_options": transportSchema,
						"stateless": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ingress_security_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_options": icmpSchema,
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"source": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tcp_options": transportSchema,
						"udp_options": transportSchema,
						"stateless": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSecurityList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	crd := &SecurityListResourceCrud{}
	crd.D = d
	crd.Client = client.client
	return crud.CreateResource(d, crd)
}

func readSecurityList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	crd := &SecurityListResourceCrud{}
	crd.D = d
	crd.Client = client.client
	return crud.ReadResource(crd)
}

func updateSecurityList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	crd := &SecurityListResourceCrud{}
	crd.D = d
	crd.Client = client.client
	return crud.UpdateResource(d, crd)
}

func deleteSecurityList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	crd := &SecurityListResourceCrud{}
	crd.D = d
	crd.Client = client.clientWithoutNotFoundRetries
	return crud.DeleteResource(d, crd)
}

type SecurityListResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.SecurityList
}

func (s *SecurityListResourceCrud) ID() string {
	return s.Res.ID
}

func (s *SecurityListResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *SecurityListResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *SecurityListResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *SecurityListResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *SecurityListResourceCrud) State() string {
	return s.Res.State
}

func (s *SecurityListResourceCrud) Create() (e error) {
	// If we are creating a default resource, then don't have to
	// actually create it. Just set the ID and update it.
	if defaultId, ok := s.D.GetOk("manage_default_resource_id"); ok {
		s.D.SetId(defaultId.(string))
		e = s.Update()
		return
	}

	compartmentID := s.D.Get("compartment_id").(string)
	egress := s.buildEgressRules()
	ingress := s.buildIngressRules()
	vcnID := s.D.Get("vcn_id").(string)

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = s.D.Get("display_name").(string)

	s.Res, e = s.Client.CreateSecurityList(compartmentID, vcnID, egress, ingress, opts)

	return
}

func (s *SecurityListResourceCrud) Get() (e error) {
	res, e := s.Client.GetSecurityList(s.D.Id())
	if e == nil {
		s.Res = res

		// If this is a default resource that we removed earlier, then
		// we need to assume that the parent resource will remove it
		// and notify terraform of it. Otherwise, terraform will
		// see that the resource is still available and error out
		deleteTargetState := s.DeletedTarget()[0]
		if _, ok := s.D.GetOk("manage_default_resource_id"); ok &&
			s.D.Get("state") == deleteTargetState {
			s.Res.State = deleteTargetState
		}
	}
	return
}

func (s *SecurityListResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateSecurityListOptions{}

	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	if egress := s.buildEgressRules(); egress != nil {
		opts.EgressRules = egress
	}
	if ingress := s.buildIngressRules(); ingress != nil {
		opts.IngressRules = ingress
	}

	s.Res, e = s.Client.UpdateSecurityList(s.D.Id(), opts)

	return
}

func (s *SecurityListResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("vcn_id", s.Res.VcnID)

	confEgressRules, confIngressRules := buildConfRuleLists(s.Res)
	s.D.Set("egress_security_rules", confEgressRules)
	s.D.Set("ingress_security_rules", confIngressRules)
}

func (s *SecurityListResourceCrud) reset() (e error) {
	opts := &baremetal.UpdateSecurityListOptions{
		IngressRules: []baremetal.IngressSecurityRule{},
		EgressRules:  []baremetal.EgressSecurityRule{},
	}

	_, e = s.Client.UpdateSecurityList(s.D.Id(), opts)
	return
}

func (s *SecurityListResourceCrud) Delete() (e error) {
	if _, ok := s.D.GetOk("manage_default_resource_id"); ok {
		// We can't actually delete a default resource.
		// Clear out its settings and mark it as deleted.
		e = s.reset()
		s.D.Set("state", s.DeletedTarget()[0])
		return
	}

	return s.Client.DeleteSecurityList(s.D.Id(), nil)
}

func (s *SecurityListResourceCrud) buildEgressRules() (sdkRules []baremetal.EgressSecurityRule) {
	sdkRules = []baremetal.EgressSecurityRule{}
	for _, val := range s.D.Get("egress_security_rules").([]interface{}) {
		confRule := val.(map[string]interface{})

		sdkRule := baremetal.EgressSecurityRule{
			Destination: confRule["destination"].(string),
			ICMPOptions: s.buildICMPOptions(confRule),
			Protocol:    confRule["protocol"].(string),
			TCPOptions:  s.buildTCPOptions(confRule),
			UDPOptions:  s.buildUDPOptions(confRule),
			IsStateless: confRule["stateless"].(bool),
		}

		sdkRules = append(sdkRules, sdkRule)
	}
	return
}

func (s *SecurityListResourceCrud) buildIngressRules() (sdkRules []baremetal.IngressSecurityRule) {
	sdkRules = []baremetal.IngressSecurityRule{}
	for _, val := range s.D.Get("ingress_security_rules").([]interface{}) {
		confRule := val.(map[string]interface{})

		sdkRule := baremetal.IngressSecurityRule{
			ICMPOptions: s.buildICMPOptions(confRule),
			Protocol:    confRule["protocol"].(string),
			Source:      confRule["source"].(string),
			TCPOptions:  s.buildTCPOptions(confRule),
			UDPOptions:  s.buildUDPOptions(confRule),
			IsStateless: confRule["stateless"].(bool),
		}

		sdkRules = append(sdkRules, sdkRule)
	}
	return
}

func (s *SecurityListResourceCrud) buildICMPOptions(conf map[string]interface{}) (opts *baremetal.ICMPOptions) {
	l := conf["icmp_options"].([]interface{})
	if len(l) > 0 {
		confOpts := l[0].(map[string]interface{})
		opts = &baremetal.ICMPOptions{
			Code: uint64(confOpts["code"].(int)),
			Type: uint64(confOpts["type"].(int)),
		}
	}
	return
}

func (s *SecurityListResourceCrud) buildTCPOptions(conf map[string]interface{}) (opts *baremetal.TCPOptions) {
	options := conf["tcp_options"].([]interface{})
	if len(options) > 0 {
		sourcePortRange, destinationPortRange := s.buildSourceAndDestinationPortRanges(options)
		opts = &baremetal.TCPOptions{
			DestinationPortRange: destinationPortRange,
			SourcePortRange:      sourcePortRange,
		}
	}
	return
}

func (s *SecurityListResourceCrud) buildUDPOptions(conf map[string]interface{}) (opts *baremetal.UDPOptions) {
	options := conf["udp_options"].([]interface{})
	if len(options) > 0 {
		sourcePortRange, destinationPortRange := s.buildSourceAndDestinationPortRanges(options)
		opts = &baremetal.UDPOptions{
			DestinationPortRange: destinationPortRange,
			SourcePortRange:      sourcePortRange,
		}
	}
	return
}

func buildPortRange(conf []interface{}) (portRange *baremetal.PortRange) {
	if len(conf) > 0 && conf[0] != nil {
		mapConf := conf[0].(map[string]interface{})

		max := mapConf["max"].(int)
		min := mapConf["min"].(int)

		// Max and Min default to 0, and that is not a valid port number, so we can assume that if
		// the value is 0 then the user has not set the port number.
		// Also, note that if either max or min is set, then the service will return an error if both are not
		// set. However, we want to create the PortRange if either is set and let the service return the error.
		if max != 0 || min != 0 {
			portRange = &baremetal.PortRange{
				Max: uint64(max),
				Min: uint64(min),
			}
		}
	}
	return
}

func (s *SecurityListResourceCrud) buildSourceAndDestinationPortRanges(conf []interface{}) (sourcePortRange, destinationPortRange *baremetal.PortRange) {
	if len(conf) > 0 && conf[0] != nil {
		mapConf := conf[0].(map[string]interface{})
		sourcePortRange = buildPortRange(mapConf["source_port_range"].([]interface{}))
		destinationPortRange = buildPortRange(conf)
	}

	return
}

// Used to build rule lists for SetData.
func buildConfRuleLists(res *baremetal.SecurityList) (confEgressRules, confIngressRules []map[string]interface{}) {
	for _, egressRule := range res.EgressSecurityRules {
		confEgressRule := buildConfRule(
			egressRule.Protocol,
			egressRule.ICMPOptions,
			egressRule.TCPOptions,
			egressRule.UDPOptions,
			&egressRule.IsStateless,
		)
		confEgressRule["destination"] = egressRule.Destination
		confEgressRules = append(confEgressRules, confEgressRule)
	}

	for _, ingressRule := range res.IngressSecurityRules {
		confIngressRule := buildConfRule(
			ingressRule.Protocol,
			ingressRule.ICMPOptions,
			ingressRule.TCPOptions,
			ingressRule.UDPOptions,
			&ingressRule.IsStateless,
		)
		confIngressRule["source"] = ingressRule.Source
		confIngressRules = append(confIngressRules, confIngressRule)
	}

	return
}

// Used to build rules for SetData.
func buildConfRule(
	protocol string,
	icmpOpts *baremetal.ICMPOptions,
	tcpOpts *baremetal.TCPOptions,
	udpOpts *baremetal.UDPOptions,
	stateless *bool,
) (confRule map[string]interface{}) {
	confRule = map[string]interface{}{}
	confRule["protocol"] = protocol
	if icmpOpts != nil {
		confRule["icmp_options"] = buildConfICMPOptions(icmpOpts)
	}
	if tcpOpts != nil {
		confRule["tcp_options"] = buildConfTransportOptions(tcpOpts.DestinationPortRange, tcpOpts.SourcePortRange)
	}
	if udpOpts != nil {
		confRule["udp_options"] = buildConfTransportOptions(udpOpts.DestinationPortRange, udpOpts.SourcePortRange)
	}
	if stateless != nil {
		confRule["stateless"] = *stateless
	}
	return confRule
}

// Used to build ICMP options for SetData.
func buildConfICMPOptions(opts *baremetal.ICMPOptions) (list []interface{}) {
	confOpts := map[string]interface{}{
		"code": int(opts.Code),
		"type": int(opts.Type),
	}
	return []interface{}{confOpts}
}

// Used to build TCP/UDP port ranges for SetData.
func buildConfTransportOptions(destinationPortRange *baremetal.PortRange, sourcePortRange *baremetal.PortRange) (list []interface{}) {
	confOpts := map[string]interface{}{}
	if destinationPortRange != nil {
		confOpts["max"] = int(destinationPortRange.Max)
		confOpts["min"] = int(destinationPortRange.Min)
	}

	if sourcePortRange != nil {
		confOpts["source_port_range"] = []interface{}{map[string]interface{}{
			"max": int(sourcePortRange.Max),
			"min": int(sourcePortRange.Min),
		}}
	}

	return []interface{}{confOpts}
}
