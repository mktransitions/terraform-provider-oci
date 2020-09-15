// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example file shows how to maintain authentication policy for the current tenancy.
 */

variable "authentication_policy_password_policy_is_lowercase_characters_required" {
  default = true
}

variable "authentication_policy_password_policy_is_numeric_characters_required" {
  default = true
}

variable "authentication_policy_password_policy_is_special_characters_required" {
  default = true
}

variable "authentication_policy_password_policy_is_uppercase_characters_required" {
  default = true
}

variable "authentication_policy_password_policy_is_username_containment_allowed" {
  default = false
}

variable "authentication_policy_password_policy_minimum_password_length" {
  default = 11
}

resource "oci_identity_authentication_policy" "test_authentication_policy" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  password_policy {
    #Optional
    is_lowercase_characters_required = var.authentication_policy_password_policy_is_lowercase_characters_required
    is_numeric_characters_required   = var.authentication_policy_password_policy_is_numeric_characters_required
    is_special_characters_required   = var.authentication_policy_password_policy_is_special_characters_required
    is_uppercase_characters_required = var.authentication_policy_password_policy_is_uppercase_characters_required
    is_username_containment_allowed  = var.authentication_policy_password_policy_is_username_containment_allowed
    minimum_password_length          = var.authentication_policy_password_policy_minimum_password_length
  }

  network_policy {
    #Optional
    network_source_ids = [oci_identity_network_source.test_network_source.id] // remove this before destroy oci_identity_network_source
  }
}

data "oci_identity_authentication_policy" "test_authentication_policy" {
  #Required
  compartment_id = var.tenancy_ocid
}

