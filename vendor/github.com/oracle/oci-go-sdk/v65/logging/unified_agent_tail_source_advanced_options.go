// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, agent configurations, log data models,
// continuous queries, and managed continuous queries.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UnifiedAgentTailSourceAdvancedOptions Advanced options for logging configuration
type UnifiedAgentTailSourceAdvancedOptions struct {

	// Enable the stat watcher based on inotify
	IsEnableStatWatcher *bool `mandatory:"false" json:"isEnableStatWatcher"`

	// Follow inodes instead of following file names.
	IsFollowInodes *bool `mandatory:"false" json:"isFollowInodes"`

	// Starts to read the logs from the head of the file or the last read position recorded in pos_file, not tail.
	IsReadFromHead *bool `mandatory:"false" json:"isReadFromHead"`

	// The interval of doing compaction of pos file.
	PosFileCompactionInterval *string `mandatory:"false" json:"posFileCompactionInterval"`
}

func (m UnifiedAgentTailSourceAdvancedOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentTailSourceAdvancedOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
