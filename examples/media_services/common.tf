// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "tenancy_ocid" {}

variable "user_ocid" {}

variable "fingerprint" {}

variable "private_key_path" {}

variable "region" {}

variable "compartment_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "media_asset_state" {
  default = "ACTIVE"
}

variable "media_asset_type" {
  default = "AUDIO"
}

variable "defined_tags_value" {
  default = "value"
}

variable "display_name" {
  default = "displayName"
}

variable "freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "id" {
  default = "id"
}

variable "active_state" {
  default = "ACTIVE"
}

variable "accepted_state" {
  default = "ACCEPTED"
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "examples-tag-namespace-all"
  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
  is_cost_tracking = false
  is_retired       = false
}

variable "kms_vault_id" {}

data "oci_kms_vault" "test_vault" {
  #Required
  vault_id = var.kms_vault_id
}

data "oci_kms_keys" "test_keys_dependency_RSA" {
  #Required
  compartment_id = var.tenancy_ocid
  management_endpoint = data.oci_kms_vault.test_vault.management_endpoint
  algorithm = "RSA"

  filter {
    name = "state"
    values = ["ENABLED", "UPDATING"]
  }
}

