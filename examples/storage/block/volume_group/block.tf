// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example demonstrates how to spin up a block volume
 *
 * See examples/compute/instance/ for a real world scenario
 */
variable "tenancy_ocid" {
}

variable "compartment_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "DBSize" {
  default = "50" // size in GBs, min: 50, max 16384
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_volume" "t" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "-tf-volume"
  size_in_gbs         = var.DBSize
}

resource "oci_core_volume" "t2" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "-tf-volume-with-backup-policy"
  size_in_gbs         = var.DBSize
}

resource "oci_core_volume_backup_policy_assignment" "policy" {
  asset_id  = oci_core_volume.t2.id
  policy_id = data.oci_core_volume_backup_policies.test_boot_volume_backup_policies.volume_backup_policies[0].id
}

data "oci_core_volume_backup_policies" "test_boot_volume_backup_policies" {
  filter {
    name   = "display_name"
    values = ["bronze"]
  }
}

data "oci_core_volumes" "test_volumes" {
  compartment_id = var.compartment_ocid

  filter {
    name   = "id"
    values = [oci_core_volume.t.id]
  }
}

output "volumes" {
  value = data.oci_core_volumes.test_volumes.volumes
}

/*
 * Examples for volume groups
 */

// Example 1: Case of volume group sourced from source volumes

// Create additional volumes to have multiple volumes in the volume group
resource "oci_core_volume" "test_volume" {
  count               = 2
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = format("-tf-volume-%d", count.index + 1)
  size_in_gbs         = var.DBSize
}

resource "oci_core_volume_group" "test_volume_group_from_vol_ids" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  source_details {
    #Required
    type = "volumeIds"

    // Mix of named volume and splatted multiple volumes
     volume_ids = concat([oci_core_volume.t.id], oci_core_volume.test_volume.*.id)
  }

  #Optional
  display_name = "test-volume-group-from-vol-ids"
}

data "oci_core_volume_groups" "test_volume_groups_from_vol_ids" {
  #Required
  compartment_id = var.compartment_ocid

  filter {
    name   = "id"
    values = [oci_core_volume_group.test_volume_group_from_vol_ids.id]
  }
}

output "volumeGroupsSourcedFromVolIds" {
  value = data.oci_core_volume_groups.test_volume_groups_from_vol_ids.volume_groups
}

output "volumeGroupVolumeIdsSourcedFromVolIds" {
  value = oci_core_volume_group.test_volume_group_from_vol_ids.volume_ids
}

// Example 2: Case of volume group cloned from another volume group

resource "oci_core_volume_group" "test_volume_group_from_vol_group" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  source_details {
    #Required
    type = "volumeGroupId"

    # Use the volume group created in Example 1
    volume_group_id = oci_core_volume_group.test_volume_group_from_vol_ids.id
  }

  #Optional
  display_name = "test-volume-group-from-vol-group"
}

data "oci_core_volume_groups" "test_volume_groups_from_vol_group" {
  #Required
  compartment_id = var.compartment_ocid

  filter {
    name   = "id"
    values = [oci_core_volume_group.test_volume_group_from_vol_group.id]
  }
}

output "volumeGroupsSourcedFromVolGroup" {
  value = data.oci_core_volume_groups.test_volume_groups_from_vol_group.volume_groups
}

output "volumeGroupVolumeIdsSourcedFromVolGroup" {
  value = oci_core_volume_group.test_volume_group_from_vol_group.volume_ids
}

// Example 3: Case of volume group restored from volume group backup

resource "oci_core_volume_group" "test_volume_group_from_vol_group_backup" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  source_details {
    #Required
    type                   = "volumeGroupBackupId"
    volume_group_backup_id = oci_core_volume_group_backup.test_volume_group_backup.id
  }

  #Optional
  display_name = "test-volume-group-from-vol-group-backup"
}

data "oci_core_volume_groups" "test_volume_groups_from_vol_group_backup" {
  #Required
  compartment_id = var.compartment_ocid

  filter {
    name   = "id"
    values = [oci_core_volume_group.test_volume_group_from_vol_group_backup.id]
  }
}

output "volumeGroupsSourcedFromVolGroupBackup" {
  value = data.oci_core_volume_groups.test_volume_groups_from_vol_group_backup.volume_groups
}

output "volumeGroupVolumeIdsSourcedFromVolGroupBackup" {
  value = oci_core_volume_group.test_volume_group_from_vol_group_backup.volume_ids
}

/*
 * Examples for volume group backup
 */
resource "oci_core_volume_group_backup" "test_volume_group_backup" {
  #Required
  volume_group_id = oci_core_volume_group.test_volume_group_from_vol_ids.id

  #Optional
  display_name = "tf-volume-group-backup"
  type         = "INCREMENTAL"
}

data "oci_core_volume_group_backups" "test_volume_group_backups" {
  #Required
  compartment_id = var.compartment_ocid

  filter {
    name   = "id"
    values = [oci_core_volume_group_backup.test_volume_group_backup.id]
  }
}

output "volumeGroupBackups" {
  value = data.oci_core_volume_group_backups.test_volume_group_backups.volume_group_backups
}

