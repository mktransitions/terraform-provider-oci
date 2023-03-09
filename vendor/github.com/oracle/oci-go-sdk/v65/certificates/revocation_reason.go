// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Retrieval API
//
// API for retrieving certificates.
//

package certificates

import (
	"strings"
)

// RevocationReasonEnum Enum with underlying type: string
type RevocationReasonEnum string

// Set of constants representing the allowable values for RevocationReasonEnum
const (
	RevocationReasonUnspecified          RevocationReasonEnum = "UNSPECIFIED"
	RevocationReasonKeyCompromise        RevocationReasonEnum = "KEY_COMPROMISE"
	RevocationReasonCaCompromise         RevocationReasonEnum = "CA_COMPROMISE"
	RevocationReasonAffiliationChanged   RevocationReasonEnum = "AFFILIATION_CHANGED"
	RevocationReasonSuperseded           RevocationReasonEnum = "SUPERSEDED"
	RevocationReasonCessationOfOperation RevocationReasonEnum = "CESSATION_OF_OPERATION"
	RevocationReasonPrivilegeWithdrawn   RevocationReasonEnum = "PRIVILEGE_WITHDRAWN"
	RevocationReasonAaCompromise         RevocationReasonEnum = "AA_COMPROMISE"
)

var mappingRevocationReasonEnum = map[string]RevocationReasonEnum{
	"UNSPECIFIED":            RevocationReasonUnspecified,
	"KEY_COMPROMISE":         RevocationReasonKeyCompromise,
	"CA_COMPROMISE":          RevocationReasonCaCompromise,
	"AFFILIATION_CHANGED":    RevocationReasonAffiliationChanged,
	"SUPERSEDED":             RevocationReasonSuperseded,
	"CESSATION_OF_OPERATION": RevocationReasonCessationOfOperation,
	"PRIVILEGE_WITHDRAWN":    RevocationReasonPrivilegeWithdrawn,
	"AA_COMPROMISE":          RevocationReasonAaCompromise,
}

var mappingRevocationReasonEnumLowerCase = map[string]RevocationReasonEnum{
	"unspecified":            RevocationReasonUnspecified,
	"key_compromise":         RevocationReasonKeyCompromise,
	"ca_compromise":          RevocationReasonCaCompromise,
	"affiliation_changed":    RevocationReasonAffiliationChanged,
	"superseded":             RevocationReasonSuperseded,
	"cessation_of_operation": RevocationReasonCessationOfOperation,
	"privilege_withdrawn":    RevocationReasonPrivilegeWithdrawn,
	"aa_compromise":          RevocationReasonAaCompromise,
}

// GetRevocationReasonEnumValues Enumerates the set of values for RevocationReasonEnum
func GetRevocationReasonEnumValues() []RevocationReasonEnum {
	values := make([]RevocationReasonEnum, 0)
	for _, v := range mappingRevocationReasonEnum {
		values = append(values, v)
	}
	return values
}

// GetRevocationReasonEnumStringValues Enumerates the set of values in String for RevocationReasonEnum
func GetRevocationReasonEnumStringValues() []string {
	return []string{
		"UNSPECIFIED",
		"KEY_COMPROMISE",
		"CA_COMPROMISE",
		"AFFILIATION_CHANGED",
		"SUPERSEDED",
		"CESSATION_OF_OPERATION",
		"PRIVILEGE_WITHDRAWN",
		"AA_COMPROMISE",
	}
}

// GetMappingRevocationReasonEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRevocationReasonEnum(val string) (RevocationReasonEnum, bool) {
	enum, ok := mappingRevocationReasonEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
