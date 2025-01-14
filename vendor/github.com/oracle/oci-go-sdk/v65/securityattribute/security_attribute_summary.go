// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Security Attribute API
//
// Use the Security Attributes API to manage security attributes and security attribute namespaces. For more information, see the documentation for Security Attributes (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attributes.htm) and Security Attribute Nampespaces (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
//

package securityattribute

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityAttributeSummary A security attribute definition that belongs to a specific security attribute namespace.
type SecurityAttributeSummary struct {

	// The OCID of the compartment that contains the security attribute.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of the namespace that contains the security attribute.
	SecurityAttributeNamespaceId *string `mandatory:"false" json:"securityAttributeNamespaceId"`

	// The name of the security attribute namespace that contains the security attribute.
	SecurityAttributeNamespaceName *string `mandatory:"false" json:"securityAttributeNamespaceName"`

	// The OCID of the security attribute.
	Id *string `mandatory:"false" json:"id"`

	// The name assigned to the security attribute during creation. This is the security attribute.
	// The name must be unique within the security attribute namespace and cannot be changed.
	Name *string `mandatory:"false" json:"name"`

	// The description you assign to the security attribute.
	Description *string `mandatory:"false" json:"description"`

	// The data type of the security attribute.
	Type *string `mandatory:"false" json:"type"`

	// Whether the security attribute is retired.
	// See Managing Security Attributes (https://docs.cloud.oracle.com/Content/zero-trust-packet-routing/managing-security-attributes.htm).
	IsRetired *bool `mandatory:"false" json:"isRetired"`

	// The security attribute's current state. After creating a security attribute, make sure its `lifecycleState` is ACTIVE before using it. After retiring a security attribute, make sure its `lifecycleState` is INACTIVE before using it. If you delete a security attribute, you cannot delete another security attribute until the deleted security attribute's `lifecycleState` changes from DELETING to DELETED.
	LifecycleState SecurityAttributeLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Date and time the security attribute was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m SecurityAttributeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityAttributeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSecurityAttributeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSecurityAttributeLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
