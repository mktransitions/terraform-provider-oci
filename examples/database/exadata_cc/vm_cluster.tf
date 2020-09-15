// Copyright (c) 2017 2019 Oracle and/or its affiliates. All rights reserved.

resource "oci_database_vm_cluster_network" "test_vm_cluster_network" {
  compartment_id = var.compartment_id
  display_name   = "testVmClusterRecommendedNetwork"
  dns            = ["192.168.10.10"]
  ntp            = ["192.168.10.20"]

  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

  scans {
    hostname = "myprefix1-nsubz-scan"

    ips = [
      "192.168.19.7",
      "192.168.19.6",
      "192.168.19.8",
    ]

    port = 1521
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.169.20.1"
    netmask      = "255.255.0.0"
    network_type = "BACKUP"

    nodes {
      hostname = "myprefix2-cghdm1"
      ip       = "192.169.19.18"
    }

    nodes {
      hostname = "myprefix2-cghdm2"
      ip       = "192.169.19.20"
    }

    vlan_id = "11"
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.168.20.1"
    netmask      = "255.255.0.0"
    network_type = "CLIENT"

    nodes {
      hostname     = "myprefix1-r64zc1"
      ip           = "192.168.19.10"
      vip          = "192.168.19.11"
      vip_hostname = "myprefix1-r64zc1-vip"
    }

    nodes {
      hostname     = "myprefix1-r64zc2"
      ip           = "192.168.19.14"
      vip          = "192.168.19.15"
      vip_hostname = "myprefix1-r64zc2-vip"
    }

    vlan_id = "10"
  }

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }

  validate_vm_cluster_network = true
}

resource "oci_database_vm_cluster" "test_vm_cluster" {
  #Required
  compartment_id            = var.compartment_id
  cpu_core_count            = "4"
  display_name              = "testVmCluster"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  gi_version                = "19.1.0.0"
  ssh_public_keys           = [var.ssh_public_key]
  vm_cluster_network_id     = oci_database_vm_cluster_network.test_vm_cluster_network.id

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }

  is_local_backup_enabled     = "false"
  is_sparse_diskgroup_enabled = "false"
  license_model               = "LICENSE_INCLUDED"
  data_storage_size_in_tbs    = "84"
  db_node_storage_size_in_gbs = "120"
  memory_size_in_gbs          = "60"
}

data "oci_database_vm_cluster_recommended_network" "test_vm_cluster_recommended_network" {
  #Required
  compartment_id            = var.compartment_id
  display_name              = "testVmClusterRecommendedNetwork"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

  networks {
    #Required
    cidr         = "192.168.19.2/16"
    domain       = "oracle.com"
    gateway      = "192.168.20.1"
    netmask      = "255.255.0.0"
    network_type = "CLIENT"
    prefix       = "myprefix1"
    vlan_id      = "10"
  }

  networks {
    #Required
    cidr         = "192.169.19.1/16"
    domain       = "oracle.com"
    gateway      = "192.169.20.1"
    netmask      = "255.255.0.0"
    network_type = "BACKUP"
    prefix       = "myprefix2"
    vlan_id      = "11"
  }

  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }
}

resource "oci_database_db_home" "test_db_home_vm_cluster" {
  vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id

  database {
    admin_password = "BEstrO0ng_#11"
    db_name        = "dbVMClus"
    character_set  = "AL32UTF8"
    ncharacter_set = "AL16UTF16"
    db_workload    = "OLTP"
    pdb_name       = "pdbName"

    freeform_tags = {
      "Department" = "Finance"
    }

    db_backup_config {
      auto_backup_enabled = true
      auto_backup_window  = "SLOT_TWO"

      backup_destination_details {
        id   = oci_database_backup_destination.test_backup_destination_nfs.id
        type = "NFS"
      }
    }
  }

  source       = "VM_CLUSTER_NEW"
  db_version   = "12.1.0.2"
  display_name = "createdDbHome"
}

resource "oci_database_backup_destination" "test_backup_destination_nfs" {
  #Required
  compartment_id = var.compartment_id
  display_name   = "testBackupDestination"
  type           = "NFS"

  #Optional

  freeform_tags = {
    "Department" = "Finance"
  }
  mount_type_details {
    local_mount_point_path = "localMountPointPath"
    mount_type             = "SELF_MOUNT"
  }
}

data "oci_database_vm_cluster_network_download_config_file" "test_vm_cluster_network_download_config_file" {
  #Required
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  vm_cluster_network_id     = oci_database_vm_cluster_network.test_vm_cluster_network.id

  #Optional
  base64_encode_content = true
}

data "oci_database_vm_cluster_networks" "test_vm_cluster_networks" {
  #Required
  compartment_id            = var.compartment_id
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
}

data "oci_database_vm_clusters" "test_vm_clusters" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
}

resource "local_file" "test_vm_cluster_network_downloaded_config_file" {
  content  = data.oci_database_vm_cluster_network_download_config_file.test_vm_cluster_network_download_config_file.content
  filename = "${path.module}/vm_cluster_config.txt"
}

