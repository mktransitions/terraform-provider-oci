// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

var tenancyResourceGraphs = map[string]TerraformResourceGraph{
	"budget":        budgetResourceGraph,
	"email_tenancy": emailTenancyResourceGraph,
	"identity":      identityResourceGraph,
	"limits":        limitsResourceGraph,
}

var compartmentResourceGraphs = map[string]TerraformResourceGraph{
	"availability_domain": availabilityDomainResourceGraph,
	"analytics":           analyticsResourceGraph,
	"apigateway":          apigatewayResourceGraph,
	"auto_scaling":        autoScalingResourceGraph,
	"bds":                 bdsResourceGraph,
	"blockchain":          blockchainResourceGraph,
	"cloud_guard":         cloudGuardResourceGraph,
	"containerengine":     containerengineResourceGraph,
	"core":                coreResourceGraph,
	"data_safe":           dataSafeResourceGraph,
	"database":            databaseResourceGraph,
	"datacatalog":         datacatalogResourceGraph,
	"dataflow":            dataflowResourceGraph,
	"dataintegration":     dataintegrationResourceGraph,
	"datascience":         datascienceResourceGraph,
	"dns":                 dnsResourceGraph,
	"email":               emailResourceGraph,
	"events":              eventsResourceGraph,
	"file_storage":        fileStorageResourceGraph,
	"functions":           functionsResourceGraph,
	"health_checks":       healthChecksResourceGraph,
	"integration":         integrationResourceGraph,
	"kms":                 kmsResourceGraph,
	"load_balancer":       loadBalancerResourceGraph,
	"marketplace":         marketplaceResourceGraph,
	"monitoring":          monitoringResourceGraph,
	"mysql":               mysqlResourceGraph,
	"nosql":               nosqlResourceGraph,
	"object_storage":      objectStorageResourceGraph,
	"oce":                 oceResourceGraph,
	"ocvp":                ocvpResourceGraph,
	"oda":                 odaResourceGraph,
	"ons":                 onsResourceGraph,
	"osmanagement":        osmanagementResourceGraph,
	"sch":                 schResourceGraph,
	"streaming":           streamingResourceGraph,
	"tagging":             taggingResourceGraph,
	"waas":                waasResourceGraph,
}

var availabilityDomainResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportIdentityAvailabilityDomainHints,
		},
	},
	"oci_identity_availability_domain": {
		{
			TerraformResourceHints: exportCoreBootVolumeHints,
			datasourceQueryParams: map[string]string{
				"availability_domain": "name",
			},
		},
		{
			TerraformResourceHints: exportFileStorageFileSystemHints,
			datasourceQueryParams: map[string]string{
				"availability_domain": "name",
			},
		},
		{
			TerraformResourceHints: exportFileStorageMountTargetHints,
			datasourceQueryParams: map[string]string{
				"availability_domain": "name",
			},
		},
	},
	"oci_file_storage_file_system": {
		{
			TerraformResourceHints: exportFileStorageSnapshotHints,
			datasourceQueryParams: map[string]string{
				"file_system_id": "id",
			},
		},
	},
}

var analyticsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAnalyticsAnalyticsInstanceHints},
	},
}

var apigatewayResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportApigatewayGatewayHints},
		{TerraformResourceHints: exportApigatewayDeploymentHints},
	},
}

var autoScalingResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAutoScalingAutoScalingConfigurationHints},
	},
}

var bdsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportBdsBdsInstanceHints},
	},
}

var blockchainResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportBlockchainBlockchainPlatformHints},
	},
	"oci_blockchain_blockchain_platform": {
		{
			TerraformResourceHints: exportBlockchainOsnHints,
			datasourceQueryParams: map[string]string{
				"blockchain_platform_id": "id",
			},
		},
		{
			TerraformResourceHints: exportBlockchainPeerHints,
			datasourceQueryParams: map[string]string{
				"blockchain_platform_id": "id",
			},
		},
	},
}

var budgetResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{
			TerraformResourceHints: exportBudgetBudgetHints,
			datasourceQueryParams: map[string]string{
				"target_type": "'ALL'",
			},
		},
	},
	"oci_budget_budget": {
		{
			TerraformResourceHints: exportBudgetAlertRuleHints,
			datasourceQueryParams: map[string]string{
				"budget_id": "id",
			},
		},
	},
}

var cloudGuardResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCloudGuardTargetHints},
		{TerraformResourceHints: exportCloudGuardManagedListHints},
		{TerraformResourceHints: exportCloudGuardResponderRecipeHints},
		{TerraformResourceHints: exportCloudGuardDetectorRecipeHints},
	},
}

var containerengineResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportContainerengineClusterHints},
		{TerraformResourceHints: exportContainerengineNodePoolHints},
	},
}

var coreResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCoreBootVolumeBackupHints},
		{TerraformResourceHints: exportCoreConsoleHistoryHints},
		{TerraformResourceHints: exportCoreClusterNetworkHints},
		{TerraformResourceHints: exportCoreComputeImageCapabilitySchemaHints},
		{TerraformResourceHints: exportCoreCpeHints},
		{TerraformResourceHints: exportCoreCrossConnectGroupHints},
		{TerraformResourceHints: exportCoreCrossConnectHints},
		{TerraformResourceHints: exportCoreDrgAttachmentHints},
		{TerraformResourceHints: exportCoreDrgHints},
		{TerraformResourceHints: exportCoreDedicatedVmHostHints},
		{TerraformResourceHints: exportCoreImageHints},
		{TerraformResourceHints: exportCoreInstanceConfigurationHints},
		{TerraformResourceHints: exportCoreInstanceConsoleConnectionHints},
		{TerraformResourceHints: exportCoreInstancePoolHints},
		{TerraformResourceHints: exportCoreInstanceHints},
		{TerraformResourceHints: exportCoreInternetGatewayHints},
		{TerraformResourceHints: exportCoreIpSecConnectionHints},
		{TerraformResourceHints: exportCoreLocalPeeringGatewayHints},
		{TerraformResourceHints: exportCoreNetworkSecurityGroupHints},
		{
			TerraformResourceHints: exportCorePublicIpHints,
			datasourceQueryParams: map[string]string{
				"scope": "'REGION'",
			},
		},
		{TerraformResourceHints: exportCoreRemotePeeringConnectionHints},
		{TerraformResourceHints: exportCoreServiceGatewayHints},
		{TerraformResourceHints: exportCoreSubnetHints},
		{TerraformResourceHints: exportCoreVcnHints},
		{TerraformResourceHints: exportCoreVirtualCircuitHints},
		{TerraformResourceHints: exportCoreVolumeAttachmentHints},
		{TerraformResourceHints: exportCoreVolumeBackupHints},
		{TerraformResourceHints: exportCoreVolumeBackupPolicyHints},
		{TerraformResourceHints: exportCoreVolumeGroupHints},
		{TerraformResourceHints: exportCoreVolumeGroupBackupHints},
		{TerraformResourceHints: exportCoreVolumeHints},
	},
	"oci_core_boot_volume": {
		{
			TerraformResourceHints: exportCoreVolumeBackupPolicyAssignmentHints,
			datasourceQueryParams: map[string]string{
				"asset_id": "id",
			},
		},
	},
	"oci_core_instance": {
		{
			TerraformResourceHints: exportCoreVnicAttachmentHints,
			datasourceQueryParams: map[string]string{
				"instance_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreVolumeBackupPolicyAssignmentHints,
			datasourceQueryParams: map[string]string{
				"asset_id": "boot_volume_id",
			},
		},
	},
	"oci_core_network_security_group": {
		{
			TerraformResourceHints: exportCoreNetworkSecurityGroupSecurityRuleHints,
			datasourceQueryParams: map[string]string{
				"network_security_group_id": "id",
			},
		},
	},
	"oci_core_subnet": {
		{
			TerraformResourceHints: exportCorePrivateIpHints,
			datasourceQueryParams: map[string]string{
				"subnet_id": "id",
			},
		},
	},
	"oci_core_vcn": {
		{
			TerraformResourceHints: exportCoreDhcpOptionsHints,
			datasourceQueryParams: map[string]string{
				"vcn_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreNatGatewayHints,
			datasourceQueryParams: map[string]string{
				"vcn_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreRouteTableHints,
			datasourceQueryParams: map[string]string{
				"vcn_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreSecurityListHints,
			datasourceQueryParams: map[string]string{
				"vcn_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreVlanHints,
			datasourceQueryParams: map[string]string{
				"vcn_id": "id",
			},
		},
	},
	"oci_core_volume": {
		{
			TerraformResourceHints: exportCoreVolumeBackupPolicyAssignmentHints,
			datasourceQueryParams: map[string]string{
				"asset_id": "id",
			},
		},
	},
}

var dataSafeResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataSafeDataSafePrivateEndpointHints},
	},
}

var databaseResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatabaseAutonomousContainerDatabaseHints},
		{TerraformResourceHints: exportDatabaseAutonomousDatabaseHints},
		{TerraformResourceHints: exportDatabaseAutonomousExadataInfrastructureHints},
		{TerraformResourceHints: exportDatabaseAutonomousVmClusterHints},
		{TerraformResourceHints: exportDatabaseBackupDestinationHints},
		{TerraformResourceHints: exportDatabaseBackupHints},
		{TerraformResourceHints: exportDatabaseDbSystemHints},
		{TerraformResourceHints: exportDatabaseExadataInfrastructureHints},
		{TerraformResourceHints: exportDatabaseVmClusterHints},
		{TerraformResourceHints: exportDatabaseDatabaseSoftwareImageHints},
	},
	"oci_database_db_home": {
		{
			TerraformResourceHints: exportDatabaseDatabaseHints,
			datasourceQueryParams: map[string]string{
				"db_home_id": "id",
			},
		},
	},
	"oci_database_db_system": {
		{
			TerraformResourceHints: exportDatabaseDbHomeHints,
			datasourceQueryParams: map[string]string{
				"db_system_id": "id",
			},
		},
	},
	"oci_database_exadata_infrastructure": {
		{
			TerraformResourceHints: exportDatabaseVmClusterNetworkHints,
			datasourceQueryParams: map[string]string{
				"exadata_infrastructure_id": "id",
			},
		},
	},
	"oci_database_vm_cluster": {
		{
			TerraformResourceHints: exportDatabaseDbHomeHints,
			datasourceQueryParams: map[string]string{
				"vm_cluster_id": "id",
			},
		},
	},
}

var datacatalogResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatacatalogCatalogHints},
		{TerraformResourceHints: exportDatacatalogCatalogPrivateEndpointHints},
	},
	"oci_datacatalog_catalog": {
		{
			TerraformResourceHints: exportDatacatalogConnectionHints,
			datasourceQueryParams: map[string]string{
				"catalog_id":     "id",
				"data_asset_key": "data_asset_key",
			},
		},
		{
			TerraformResourceHints: exportDatacatalogDataAssetHints,
			datasourceQueryParams: map[string]string{
				"catalog_id": "id",
			},
		},
	},
	"oci_datacatalog_data_asset": {
		{
			TerraformResourceHints: exportDatacatalogConnectionHints,
			datasourceQueryParams: map[string]string{
				"data_asset_key": "key",
				"catalog_id":     "catalog_id",
			},
		},
	},
}

var dataflowResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataflowApplicationHints},
		{TerraformResourceHints: exportDataflowPrivateEndpointHints},
	},
}

var dataintegrationResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataintegrationWorkspaceHints},
	},
}

var datascienceResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatascienceProjectHints},
		{TerraformResourceHints: exportDatascienceNotebookSessionHints},
		{TerraformResourceHints: exportDatascienceModelHints},
	},
	"oci_datascience_model": {
		{
			TerraformResourceHints: exportDatascienceModelProvenanceHints,
			datasourceQueryParams: map[string]string{
				"model_id": "id",
			},
		},
	},
}

var dnsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDnsZoneHints},
		{TerraformResourceHints: exportDnsSteeringPolicyHints},
		{TerraformResourceHints: exportDnsSteeringPolicyAttachmentHints},
		{TerraformResourceHints: exportDnsTsigKeyHints},
	},
	"oci_dns_zone": {
		{
			TerraformResourceHints: exportDnsRrsetHints,
			datasourceQueryParams: map[string]string{
				"zone_name_or_id": "id",
			},
		},
	},
}

var emailResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportEmailSenderHints},
	},
}

var emailTenancyResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportEmailSuppressionHints},
	},
}

var eventsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportEventsRuleHints},
	},
}

var fileStorageResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportFileStorageExportHints},
	},
}

var functionsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportFunctionsApplicationHints},
	},
	"oci_functions_application": {
		{
			TerraformResourceHints: exportFunctionsFunctionHints,
			datasourceQueryParams: map[string]string{
				"application_id": "id",
			},
		},
	},
}

var healthChecksResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportHealthChecksHttpMonitorHints},
		{TerraformResourceHints: exportHealthChecksPingMonitorHints},
	},
}

var identityResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportIdentityAuthenticationPolicyHints},
		{TerraformResourceHints: exportIdentityCompartmentHints},
		{TerraformResourceHints: exportIdentityDynamicGroupHints},
		{TerraformResourceHints: exportIdentityGroupHints},
		{
			TerraformResourceHints: exportIdentityIdentityProviderHints,
			datasourceQueryParams:  map[string]string{"protocol": "'SAML2'"},
		},
		{TerraformResourceHints: exportIdentityPolicyHints},
		{TerraformResourceHints: exportIdentityUserHints},
		{TerraformResourceHints: exportIdentityNetworkSourceHints},
	},
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportIdentityCompartmentHints,
			datasourceQueryParams:  map[string]string{"compartment_id": "id"},
		},
		{
			TerraformResourceHints: exportIdentityPolicyHints,
			datasourceQueryParams:  map[string]string{"compartment_id": "id"},
		},
	},
	"oci_identity_identity_provider": {
		{
			TerraformResourceHints: exportIdentityIdpGroupMappingHints,
			datasourceQueryParams: map[string]string{
				"identity_provider_id": "id",
			},
		},
	},
	"oci_identity_user": {
		{
			TerraformResourceHints: exportIdentityApiKeyHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityAuthTokenHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityCustomerSecretKeyHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentitySmtpCredentialHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentitySwiftPasswordHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityUiPasswordHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIdentityUserGroupMembershipHints,
			datasourceQueryParams: map[string]string{
				"user_id": "id",
			},
		},
	},
}

var integrationResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportIntegrationIntegrationInstanceHints},
	},
}

var kmsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportKmsVaultHints},
	},
	"oci_kms_key": {
		{
			TerraformResourceHints: exportKmsKeyVersionHints,
			datasourceQueryParams: map[string]string{
				"key_id":              "id",
				"management_endpoint": "management_endpoint",
			},
		},
	},
	"oci_kms_vault": {
		{
			TerraformResourceHints: exportKmsKeyHints,
			datasourceQueryParams: map[string]string{
				"management_endpoint": "management_endpoint",
			},
		},
	},
}

var limitsResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportLimitsQuotaHints},
	},
}

var loadBalancerResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportLoadBalancerLoadBalancerHints},
	},
	"oci_load_balancer_backend_set": {
		{
			TerraformResourceHints: exportLoadBalancerBackendHints,
			datasourceQueryParams: map[string]string{
				"backendset_name":  "name",
				"load_balancer_id": "load_balancer_id",
			},
		},
		{TerraformResourceHints: exportLoadBalancerListenerHints},
	},
	"oci_load_balancer_load_balancer": {
		{
			TerraformResourceHints: exportLoadBalancerBackendSetHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerCertificateHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerHostnameHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerPathRouteSetHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerRuleSetHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
	},
}

var marketplaceResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportMarketplaceAcceptedAgreementHints},
	},
}

var monitoringResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportMonitoringAlarmHints},
	},
}

var mysqlResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportMysqlMysqlBackupHints},
		{TerraformResourceHints: exportMysqlMysqlDbSystemHints},
	},
}

var nosqlResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportNosqlTableHints},
	},
	"oci_nosql_table": {
		{
			TerraformResourceHints: exportNosqlIndexHints,
			datasourceQueryParams: map[string]string{
				"table_name_or_id": "id",
			},
		},
	},
}

var objectStorageResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportObjectStorageNamespaceHints},
	},
	"oci_objectstorage_bucket": {
		{
			TerraformResourceHints: exportObjectStorageObjectHints,
			datasourceQueryParams: map[string]string{
				"bucket":    "name",
				"namespace": "namespace",
			},
		},
		{
			TerraformResourceHints: exportObjectStorageObjectLifecyclePolicyHints,
			datasourceQueryParams: map[string]string{
				"namespace": "namespace",
				"bucket":    "name",
			},
		},
		{
			TerraformResourceHints: exportObjectStoragePreauthenticatedRequestHints,
			datasourceQueryParams: map[string]string{
				"namespace": "namespace",
				"bucket":    "name",
			},
		},
		{
			TerraformResourceHints: exportObjectStorageReplicationPolicyHints,
			datasourceQueryParams: map[string]string{
				"namespace": "namespace",
				"bucket":    "name",
			},
		},
	},
	"oci_objectstorage_namespace": {
		{
			TerraformResourceHints: exportObjectStorageBucketHints,
			datasourceQueryParams: map[string]string{
				"namespace": "namespace",
			},
		},
	},
}

var oceResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOceOceInstanceHints},
	},
}

var ocvpResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOcvpSddcHints},
	},
	"oci_ocvp_sddc": {
		{
			TerraformResourceHints: exportOcvpEsxiHostHints,
			datasourceQueryParams: map[string]string{
				"sddc_id": "id",
			},
		},
	},
}

var odaResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOdaOdaInstanceHints},
	},
}

var onsResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOnsNotificationTopicHints},
		{TerraformResourceHints: exportOnsSubscriptionHints},
	},
}

var osmanagementResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOsmanagementManagedInstanceGroupHints},
		{TerraformResourceHints: exportOsmanagementSoftwareSourceHints},
	},
}

var schResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportSchServiceConnectorHints},
	},
}

var streamingResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportStreamingConnectHarnessHints},
		{TerraformResourceHints: exportStreamingStreamPoolHints},
		{TerraformResourceHints: exportStreamingStreamHints},
	},
}

var waasResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportWaasAddressListHints},
		{TerraformResourceHints: exportWaasCustomProtectionRuleHints},
		{TerraformResourceHints: exportWaasHttpRedirectHints},
		{TerraformResourceHints: exportWaasWaasPolicyHints},
	},
}

var taggingResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportIdentityTagDefaultHints},
		{TerraformResourceHints: exportIdentityTagNamespaceHints},
	},
	"oci_identity_tag_namespace": {
		{
			TerraformResourceHints: exportIdentityTagHints,
			datasourceQueryParams: map[string]string{
				"tag_namespace_id": "id",
			},
		},
	},
}
