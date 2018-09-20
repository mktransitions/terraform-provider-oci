// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"log"
)

const Version = "3.0.1"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
