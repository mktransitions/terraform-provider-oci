// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This file demonstrates dns zone management
 */

resource "random_string" "random_prefix" {
  length  = 4
  number  = false
  special = false
}

resource "oci_dns_zone" "zone1" {
  compartment_id = var.compartment_ocid
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example-primary.oci-dns1"
  zone_type      = "PRIMARY"
}

resource "oci_dns_zone" "zone3" {
  compartment_id = var.compartment_ocid
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example3-primary.oci-dns1"
  zone_type      = "PRIMARY"
}

resource "oci_dns_tsig_key" "test_tsig_key" {
  algorithm      = "hmac-sha1"
  compartment_id = var.compartment_ocid
  name           = "test_tsig_key-name"
  secret         = "c2VjcmV0"
}

resource "oci_dns_zone" "zone2" {
  compartment_id = var.compartment_ocid
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example-secondary.oci-dns2"
  zone_type      = "SECONDARY"

  external_masters {
    address     = "77.64.12.1"
    tsig_key_id = oci_dns_tsig_key.test_tsig_key.id
  }

  external_masters {
    address     = "77.64.12.2"
    tsig_key_id = oci_dns_tsig_key.test_tsig_key.id
  }
}

data "oci_dns_zones" "zs" {
  compartment_id = var.compartment_ocid
  name_contains  = "example"
  state          = "ACTIVE"
  zone_type      = "PRIMARY"
  sort_by        = "name" # name|zoneType|timeCreated
  sort_order     = "DESC" # ASC|DESC
}

data "oci_identity_tenancy" "tenancy" {
  tenancy_id = var.tenancy_ocid
}

output "zones" {
  value = data.oci_dns_zones.zs.zones
}

