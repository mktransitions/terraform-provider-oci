// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

// Options
// To get the body, optional and required are marshalled and merged.
// To get the query string, optional and required are merged.
// To get the header, optional and required are merged.
// Required options get built inline within exported functions, based on
// function parameters.
// A single options struct gets passed to exported functions, representing
// optional params.
// Both required and optional fields need to explicitly tag header, json and
// url, excluding appropriately.

type IfMatchOptions struct {
	IfMatch string `header:"If-Match,omitempty" json:"-" url:"-"`
}

type IfNoneMatchOptions struct {
	IfNoneMatch string `header:"If-None-Match,omitempty" json:"-" url:"-"`
}

type RetryTokenOptions struct {
	RetryToken string `header:"opc-retry-token,omitempty" json:"-" url:"-"`
}

type HeaderOptions struct {
	IfMatchOptions
	RetryTokenOptions
}

type ClientRequestOptions struct {
	OPCClientRequestID string `header:"opc-client-request-id,omitempty" json:"-" url:"-"`
}

type DisplayNameOptions struct {
	DisplayName string `header:"-" json:"displayName,omitempty" url:"-"`
}

type VersionDateOptions struct {
	VersionDate string `header:"-" json:"versionDate,omitempty" url:"-"`
}

// Creation Options

type CreateOptions struct {
	RetryTokenOptions
	DisplayNameOptions
}

type CreateBucketOptions struct {
	Metadata map[string]string `header:"-" json:"metadata,omitempty" url:"-"`
}

type CreateVcnOptions struct {
	CreateOptions
	DnsLabel              string `header:"-" json:"dnsLabel,omitempty" url:"-"`
	DefaultDHCPOptionsID  string `header:"-" json:"defaultDhcpOptionsId,omitempty" url:"-"`
	DefaultRouteTableID   string `header:"-" json:"defaultRouteTableId,omitempty" url:"-"`
	DefaultSecurityListID string `header:"-" json:"defaultSecurityListId,omitempty" url:"-"`
}

type CreateSubnetOptions struct {
	CreateOptions
	DHCPOptionsID   string   `header:"-" json:"dhcpOptionsId,omitempty" url:"-"`
	DNSLabel        string   `header:"-" json:"dnsLabel,omitempty" url:"-"`
	RouteTableID    string   `header:"-" json:"routeTableId,omitempty" url:"-"`
	SecurityListIDs []string `header:"-" json:"securityListIds,omitempty" url:"-"`
}

type LoadBalancerOptions struct {
	ClientRequestOptions
	RetryTokenOptions
}

type CreateLoadBalancerBackendOptions struct {
	LoadBalancerOptions
	Backup  bool `header:"-" json:"backup,omitempty", url:"-"`
	Drain   bool `header:"-" json:"drain,omitempty", url:"-"`
	Offline bool `header:"-" json:"offline,omitempty", url:"-"`
	Weight  int  `header:"-" json:"weight,omitempty", url:"-"`
}

type UpdateLoadBalancerBackendSetOptions struct {
	LoadBalancerOptions
	RetryTokenOptions
	Backends      []Backend        `header:"-" json:"backends,omitempty" url:"-"`
	HealthChecker HealthChecker    `header:"-" json:"healthChecker,omitempty" url:"-"`
	Policy        string           `header:"-" json:"policy,omitempty" url:"-"`
	SSLConfig     SSLConfiguration `header:"-" json:"sslConfiguration,omitempty" url:"-"`
}

type UpdateLoadBalancerListenerOptions struct {
	LoadBalancerOptions
	DefaultBackendSetName string           `header:"-" json:"defaultBackendSetName" url:"-"`
	Port                  int              `header:"-" json:"port" url:"-"`
	Protocol              string           `header:"-" json:"protocol" url:"-"`
	SSLConfig             SSLConfiguration `header:"-" json:"sslConfiguration" url:"-"`
}

type ListLoadBalancerPolicyOptions struct {
	ClientRequestOptions
	ListOptions
}

type ListLoadBalancerOptions struct {
	ClientRequestOptions
	ListOptions
	CompartmentID string `header:"-" json:"-" url:"compartmentId,omitempty"`
	Detail        string `header:"-" json:"-" url:"detail,omitempty"`
}

type UpdateLoadBalancerOptions struct {
	LoadBalancerOptions
	DisplayNameOptions
}

type CreateVolumeOptions struct {
	CreateOptions
	SizeInMBs      int    `header:"-" json:"sizeInMBs,omitempty" url:"-"`
	VolumeBackupID string `header:"-" json:"volumeBackupId,omitempty" url:"-"`
}

type CreatePolicyOptions struct {
	RetryTokenOptions
	VersionDateOptions
}

type LaunchInstanceOptions struct {
	CreateOptions
	HostnameLabel string            `header:"-" json:"hostnameLabel,omitempty" url:"-"`
	Metadata      map[string]string `header:"-" json:"metadata,omitempty" url:"-"`
}

type LaunchDBSystemOptions struct {
	CreateOptions
	DatabaseEdition DatabaseEdition     `header:"-" json:"databaseEdition,omitempty" url:"-"`
	DBHome          createDBHomeDetails `header:"-" json:"dbHome,omitempty" url:"-"`
	DiskRedundancy  DiskRedundancy      `header:"-" json:"diskRedundancy,omitempty" url:"-"`
	Domain          string              `header:"-" json:"domain,omitempty" url:"-"`
	Hostname        string              `header:"-" json:"hostname,omitempty" url:"-"`
}

// Read Options

type GetObjectOptions struct {
	IfMatchOptions
	IfNoneMatchOptions
	ClientRequestOptions
	Range string `header:"Range,omitempty" json:"-" url:"-"`
}

// Update Options

type UpdateOptions struct {
	HeaderOptions
	DisplayNameOptions
}

type IfMatchDisplayNameOptions struct {
	IfMatchOptions
	DisplayNameOptions
}

type UpdateBucketOptions struct {
	IfMatchOptions
	Name      string            `header:"-" json:"name,omitempty" url:"-"`
	Namespace Namespace         `header:"-" json:"namespace,omitempty" url:"-"`
	Metadata  map[string]string `header:"-" json:"metadata,omitempty" url:"-"`
}

type UpdateIdentityOptions struct {
	IfMatchOptions
	Description string `header:"-" json:"description,omitempty" url:"-"`
}

type UpdateUserStateOptions struct {
	IfMatchOptions
	Blocked *bool `header:"-" json:"blocked,omitempty" url:"-"`
}

type UpdatePolicyOptions struct {
	UpdateIdentityOptions
	VersionDateOptions
	Statements []string `header:"-" json:"statements,omitempty" url:"-"`
}

type UpdateDHCPDNSOptions struct {
	CreateOptions
	Options []DHCPDNSOption `header:"-" json:"options,omitempty" url:"-"`
}

type UpdateGatewayOptions struct {
	IfMatchOptions
	DisplayNameOptions
	IsEnabled bool `header:"-" json:"isEnabled,omitempty" url:"-"`
}

type UpdateRouteTableOptions struct {
	CreateOptions
	RouteRules []RouteRule `header:"-" json:"routeRules,omitempty" url:"-"`
}

type UpdateSecurityListOptions struct {
	IfMatchDisplayNameOptions
	EgressRules  []EgressSecurityRule  `header:"-" json:"egressSecurityRules,omitempty" url:"-"`
	IngressRules []IngressSecurityRule `header:"-" json:"ingressSecurityRules,omitempty" url:"-"`
}

type PutObjectOptions struct {
	IfMatchOptions
	IfNoneMatchOptions
	ClientRequestOptions
	Expect          string `header:"Expect,omitempty" json:"-" url:"-"`
	ContentMD5      string `header:"Content-MD5,omitempty" json:"-" url:"-"`
	ContentType     string `header:"Content-Type,omitempty" json:"-" url:"-"`
	ContentLanguage string `header:"Content-Language,omitempty" json:"-" url:"-"`
	ContentEncoding string `header:"Content-Encoding,omitempty" json:"-" url:"-"`

	// TODO: Metadata is handled explicitly during marshal.
	Metadata map[string]string `header:"-" json:"-" url:"-"`
}

// Delete Options

type DeleteObjectOptions struct {
	IfMatchOptions
	ClientRequestOptions
}

// List Options

type PageListOptions struct {
	Page string `header:"-" json:"-" url:"page,omitempty"`
}

type LimitListOptions struct {
	Limit uint64 `header:"-" json:"-" url:"limit,omitempty"`
}

type ListOptions struct {
	LimitListOptions
	PageListOptions
}

type DisplayNameListOptions struct {
	DisplayName string `header:"-" json:"-" url:"displayName,omitempty"`
}

type AvailabilityDomainListOptions struct {
	AvailabilityDomain string `header:"-" json:"-" url:"availabilityDomain,omitempty"`
}

type DrgIDListOptions struct {
	DrgID string `header:"-" json:"-" url:"drgId,omitempty"`
}

type InstanceIDListOptions struct {
	InstanceID string `header:"-" json:"-" url:"instanceId,omitempty"`
}

type ListInstancesOptions struct {
	AvailabilityDomainListOptions
	DisplayNameListOptions
	ListOptions
}

type ListConsoleHistoriesOptions struct {
	AvailabilityDomainListOptions
	InstanceIDListOptions
	ListOptions
}

type ListDrgAttachmentsOptions struct {
	DrgIDListOptions
	ListOptions
	VcnID string `header:"-" json:"-" url:"vcnId,omitempty"`
}

type ListImagesOptions struct {
	DisplayNameListOptions
	ListOptions
	OperatingSystem        string `header:"-" json:"-" url:"operatingSystem,omitempty"`
	OperatingSystemVersion string `header:"-" json:"-" url:"operatingSystemVersion,omitempty"`
}

type ListIPSecConnsOptions struct {
	DrgIDListOptions
	ListOptions
	CpeID string `header:"-" json:"-" url:"cpeId,omitempty"`
}

type ListShapesOptions struct {
	AvailabilityDomainListOptions
	ListOptions
	ImageID string `header:"-" json:"-" url:"imageId,omitempty"`
}

type ListVnicAttachmentsOptions struct {
	AvailabilityDomainListOptions
	InstanceIDListOptions
	ListOptions
	VnicID string `header:"-" json:"-" url:"vnicId,omitempty"`
}

type ListVolumesOptions struct {
	AvailabilityDomainListOptions
	ListOptions
}

type ListVolumeAttachmentsOptions struct {
	AvailabilityDomainListOptions
	InstanceIDListOptions
	ListOptions
	VolumeID string `header:"-" json:"-" url:"volumeId,omitempty"`
}

type ListBackupsOptions struct {
	ListOptions
	VolumeID string `header:"-" json:"-" url:"volumeId,omitempty"`
}

type ListMembershipsOptions struct {
	ListOptions
	GroupID string `header:"-" json:"-" url:"groupId,omitempty"`
	UserID  string `header:"-" json:"-" url:"userId,omitempty"`
}

type ListBucketsOptions struct {
	ListOptions
	ClientRequestOptions
}

type ListObjectsOptions struct {
	ClientRequestOptions
	LimitListOptions
	Prefix    string `header:"-" json:"-" url:"prefix"`
	Start     string `header:"-" json:"-" url:"start"`
	End       string `header:"-" json:"-" url:"end"`
	Delimiter string `header:"-" json:"-" url:"delimiter"`
	Fields    string `header:"-" json:"-" url:"fields"`
}

// Misc Options

type HeadObjectOptions struct {
	IfMatchOptions
	IfNoneMatchOptions
	ClientRequestOptions
}

type ConsoleHistoryDataOptions struct {
	Length uint64 `header:"-" json:"-" url:"length,omitempty"`
	Offset uint64 `header:"-" json:"-" url:"offset,omitempty"`
}
