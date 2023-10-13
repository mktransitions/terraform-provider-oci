// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Redis Service API
//
// Use the Redis Service API to create and manage Redis clusters. A Redis cluster is a memory-based storage solution. For more information, see OCI Caching Service with Redis (https://docs.cloud.oracle.com/iaas/Content/redis/home.htm).
//

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RedisClusterSummary Summary of information about a Redis cluster. A Redis cluster is a memory-based storage solution. For more information, see OCI Caching Service with Redis (https://docs.cloud.oracle.com/iaas/Content/redis/home.htm).
type RedisClusterSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the Redis cluster.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the compartment that contains the Redis cluster.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of nodes in the Redis cluster.
	NodeCount *int `mandatory:"true" json:"nodeCount"`

	// The amount of memory allocated to the Redis cluster's nodes, in gigabytes.
	NodeMemoryInGBs *float32 `mandatory:"true" json:"nodeMemoryInGBs"`

	// The fully qualified domain name (FQDN) of the API endpoint for the Redis cluster's primary node.
	PrimaryFqdn *string `mandatory:"true" json:"primaryFqdn"`

	// The private IP address of the API endpoint for the Redis cluster's primary node.
	PrimaryEndpointIpAddress *string `mandatory:"true" json:"primaryEndpointIpAddress"`

	// The fully qualified domain name (FQDN) of the API endpoint for the Redis cluster's replica nodes.
	ReplicasFqdn *string `mandatory:"true" json:"replicasFqdn"`

	// The private IP address of the API endpoint for the Redis cluster's replica nodes.
	ReplicasEndpointIpAddress *string `mandatory:"true" json:"replicasEndpointIpAddress"`

	// The Redis version that the cluster is running.
	SoftwareVersion RedisClusterSoftwareVersionEnum `mandatory:"true" json:"softwareVersion"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the Redis cluster's subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The current state of the Redis cluster.
	LifecycleState RedisClusterLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, the message might provide actionable information for a resource in `FAILED` state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the Redis cluster was created. An RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the Redis cluster was updated. An RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m RedisClusterSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RedisClusterSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRedisClusterSoftwareVersionEnum(string(m.SoftwareVersion)); !ok && m.SoftwareVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareVersion: %s. Supported values are: %s.", m.SoftwareVersion, strings.Join(GetRedisClusterSoftwareVersionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingRedisClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRedisClusterLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
