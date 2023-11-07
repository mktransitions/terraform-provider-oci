// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to import custom content
 */

#variable "tenancy_ocid" {}
#variable "user_ocid" {}
#variable "fingerprint" {}
#variable "private_key_path" {}
#variable "compartment_ocid" {}
#variable "region" {}

provider "oci" {
#  tenancy_ocid     = var.tenancy_ocid
#  user_ocid        = var.user_ocid
#  fingerprint      = var.fingerprint
#  private_key_path = var.private_key_path
#  region           = var.region
}
#terraform {
#  required_providers {
#    oci = {
#      source = "oracle/oci"
#      version = "5.19.0"
#    }
#  }
#}
# Fetch namespace name from object store GET /n
/*
data "oci_objectstorage_namespace" "ns" {
  compartment_id             = var.compartment_ocid
}

# Create a custom content
resource "oci_log_analytics_log_analytics_import_custom_content" "importCustomContentNew" {
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  import_custom_content_file = "./files/TFSource1.zip"
}

# Create a custom content with overwrite false
resource "oci_log_analytics_log_analytics_import_custom_content" "importCustomContentOverwriteFalse" {
  depends_on = [oci_log_analytics_log_analytics_import_custom_content.importCustomContentNew]
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  import_custom_content_file = "./files/TFSource1.zip"
  is_overwrite               = "false"
}
*/
# Create a custom content with overwrite true
resource "oci_log_analytics_log_analytics_import_custom_content" "importCustomContentOverwriteTrue" {
#  depends_on = [oci_log_analytics_log_analytics_import_custom_content.importCustomContentOverwriteFalse]
  namespace                  = "dxterraformtest"
  import_custom_content_file = "./files/TFSource1.zip"
  is_overwrite               = "true"
}
