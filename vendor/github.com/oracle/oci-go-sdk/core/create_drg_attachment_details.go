// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDrgAttachmentDetails The representation of CreateDrgAttachmentDetails
type CreateDrgAttachmentDetails struct {

	// The OCID of the DRG.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The OCID of the VCN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// A user-friendly name. Does not have to be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the route table the DRG attachment will use.
	// If you don't specify a route table here, the DRG attachment is created without an associated route
	// table. The Networking service does NOT automatically associate the attached VCN's default route table
	// with the DRG attachment.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`
}

func (m CreateDrgAttachmentDetails) String() string {
	return common.PointerString(m)
}
