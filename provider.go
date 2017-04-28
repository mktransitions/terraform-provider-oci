// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/terraform-provider-baremetal/core"
	"github.com/oracle/terraform-provider-baremetal/database"
	"github.com/oracle/terraform-provider-baremetal/identity"
	"github.com/oracle/terraform-provider-baremetal/objectstorage"
)

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"tenancy_ocid": "(Required) The tenancy OCID for a user. The tenancy OCID can be found at the bottom of user settings in the Bare Metal console.",
		"user_ocid":    "(Required) The user OCID. This can be found in user settings in the Bare Metal console.",
		"fingerprint":  "(Required) The fingerprint for the user's RSA key. This can be found in user settings in the Bare Metal console.",
		"private_key": "(Optional) A PEM formatted RSA private key for the user.\n" +
			"A private_key or a private_key_path must be provided.",
		"private_key_path": "(Optional) The path to the user's PEM formatted private key.\n" +
			"A private_key or a private_key_path must be provided.",
		"private_key_password": "(Optional) The password used to secure the private key.",
	}
}

// Provider is the adapter for terraform, that gives access to all the resources
func Provider(configfn schema.ConfigureFunc) terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: dataSourcesMap(),
		Schema:         schemaMap(),
		ResourcesMap:   resourcesMap(),
		ConfigureFunc:  configfn,
	}
}

func schemaMap() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tenancy_ocid": {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["tenancy_ocid"],
		},
		"user_ocid": {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["user_ocid"],
		},
		"fingerprint": {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["fingerprint"],
		},
		// Mostly used for testing. Don't put keys in your .tf files
		"private_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Sensitive:   true,
			Description: descriptions["private_key"],
		},
		"private_key_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions["private_key_path"],
		},
		"private_key_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Default:     "",
			Description: descriptions["private_key_password"],
		},
	}
}

func dataSourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"baremetal_core_console_history_data":       core.ConsoleHistoryDataDatasource(),
		"baremetal_core_cpes":                       core.CpeDatasource(),
		"baremetal_core_dhcp_options":               core.DHCPOptionsDatasource(),
		"baremetal_core_drg_attachments":            core.DrgAttachmentDatasource(),
		"baremetal_core_drgs":                       core.DrgDatasource(),
		"baremetal_core_images":                     core.ImageDatasource(),
		"baremetal_core_instances":                  core.InstanceDatasource(),
		"baremetal_core_instance_credentials":       core.InstanceCredentialsDatasource(),
		"baremetal_core_internet_gateways":          core.InternetGatewayDatasource(),
		"baremetal_core_ipsec_config":               core.IPSecConnectionConfigDatasource(),
		"baremetal_core_ipsec_connections":          core.IPSecConnectionsDatasource(),
		"baremetal_core_ipsec_status":               core.IPSecConnectionStatusDatasource(),
		"baremetal_core_route_tables":               core.RouteTableDatasource(),
		"baremetal_core_security_lists":             core.SecurityListDatasource(),
		"baremetal_core_shape":                      core.ShapeDatasource(),
		"baremetal_core_subnets":                    core.SubnetDatasource(),
		"baremetal_core_virtual_networks":           core.VirtualNetworkDatasource(),
		"baremetal_core_vnic":                       core.VnicDatasource(),
		"baremetal_core_vnic_attachments":           core.DatasourceCoreVnicAttachments(),
		"baremetal_core_volume_attachments":         core.VolumeAttachmentDatasource(),
		"baremetal_core_volume_backups":             core.VolumeBackupDatasource(),
		"baremetal_core_volumes":                    core.VolumeDatasource(),
		"baremetal_database_database":               database.DatabaseDatasource(),
		"baremetal_database_databases":              database.DatabasesDatasource(),
		"baremetal_database_db_home":                database.DBHomeDatasource(),
		"baremetal_database_db_homes":               database.DBHomesDatasource(),
		"baremetal_database_db_node":                database.DBNodeDatasource(),
		"baremetal_database_db_nodes":               database.DBNodesDatasource(),
		"baremetal_database_db_system_shapes":       database.DBSystemShapeDatasource(),
		"baremetal_database_db_systems":             database.DBSystemDatasource(),
		"baremetal_database_db_versions":            database.DBVersionDatasource(),
		"baremetal_database_supported_operations":   database.SupportedOperationDatasource(),
		"baremetal_identity_api_keys":               identity.APIKeyDatasource(),
		"baremetal_identity_availability_domains":   identity.AvailabilityDomainDatasource(),
		"baremetal_identity_compartments":           identity.CompartmentDatasource(),
		"baremetal_identity_groups":                 identity.GroupDatasource(),
		"baremetal_identity_policies":               identity.PolicyDatasource(),
		"baremetal_identity_swift_passwords":        identity.SwiftPasswordDatasource(),
		"baremetal_identity_user_group_memberships": identity.UserGroupMembershipDatasource(),
		"baremetal_identity_users":                  identity.UserDatasource(),
		"baremetal_objectstorage_bucket_summaries":  objectstorage.BucketSummaryDatasource(),
		"baremetal_objectstorage_namespace":         objectstorage.NamespaceDatasource(),
		"baremetal_objectstorage_object_head":       objectstorage.ObjectHeadDatasource(),
		"baremetal_objectstorage_objects":           objectstorage.ObjectDatasource(),
	}
}

func resourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"baremetal_core_console_history":           core.ConsoleHistoryResource(),
		"baremetal_core_cpe":                       core.CpeResource(),
		"baremetal_core_dhcp_options":              core.DHCPOptionsResource(),
		"baremetal_core_drg":                       core.DrgResource(),
		"baremetal_core_drg_attachment":            core.DrgAttachmentResource(),
		"baremetal_core_image":                     core.ImageResource(),
		"baremetal_core_instance":                  core.InstanceResource(),
		"baremetal_core_internet_gateway":          core.InternetGatewayResource(),
		"baremetal_core_ipsec":                     core.IPSecConnectionResource(),
		"baremetal_core_route_table":               core.RouteTableResource(),
		"baremetal_core_security_list":             core.SecurityListResource(),
		"baremetal_core_subnet":                    core.SubnetResource(),
		"baremetal_core_virtual_network":           core.VirtualNetworkResource(),
		"baremetal_core_volume":                    core.VolumeResource(),
		"baremetal_core_volume_attachment":         core.VolumeAttachmentResource(),
		"baremetal_core_volume_backup":             core.VolumeBackupResource(),
		"baremetal_database_db_system":             database.DBSystemResource(),
		"baremetal_identity_api_key":               identity.APIKeyResource(),
		"baremetal_identity_compartment":           identity.CompartmentResource(),
		"baremetal_identity_group":                 identity.GroupResource(),
		"baremetal_identity_policy":                identity.PolicyResource(),
		"baremetal_identity_swift_password":        identity.SwiftPasswordResource(),
		"baremetal_identity_ui_password":           identity.UIPasswordResource(),
		"baremetal_identity_user":                  identity.UserResource(),
		"baremetal_identity_user_group_membership": identity.UserGroupMembershipResource(),
		"baremetal_objectstorage_bucket":           objectstorage.BucketResource(),
		"baremetal_objectstorage_object":           objectstorage.ObjectResource(),
	}
}

func providerConfig(d *schema.ResourceData) (client interface{}, err error) {
	tenancyOCID := d.Get("tenancy_ocid").(string)
	userOCID := d.Get("user_ocid").(string)
	fingerprint := d.Get("fingerprint").(string)
	privateKeyBuffer, hasKey := d.Get("private_key").(string)
	privateKeyPath, hasKeyPath := d.Get("private_key_path").(string)
	privateKeyPassword, hasKeyPass := d.Get("private_key_password").(string)

	clientOpts := []baremetal.NewClientOptionsFunc{
		func(o *baremetal.NewClientOptions) {
			o.UserAgent = fmt.Sprintf("baremetal-terraform-v%s", baremetal.SDKVersion)
		},
		func(o *baremetal.NewClientOptions) {
			o.Transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
		},
	}

	if hasKey && privateKeyBuffer != "" {
		clientOpts = append(clientOpts, baremetal.PrivateKeyBytes([]byte(privateKeyBuffer)))
	} else if hasKeyPath && privateKeyPath != "" {
		clientOpts = append(clientOpts, baremetal.PrivateKeyFilePath(privateKeyPath))
	} else {
		err = errors.New("One of private_key or private_key_path is required")
		return
	}

	if hasKeyPass && privateKeyPassword != "" {
		clientOpts = append(clientOpts, baremetal.PrivateKeyPassword(privateKeyPassword))
	}

	client, err = baremetal.NewClient(userOCID, tenancyOCID, fingerprint, clientOpts...)
	return
}
