// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"github.com/oracle/oci-go-sdk/v65/common"
	tf_adm "github.com/oracle/terraform-provider-oci/internal/service/adm"
	tf_ai_anomaly_detection "github.com/oracle/terraform-provider-oci/internal/service/ai_anomaly_detection"
	tf_ai_document "github.com/oracle/terraform-provider-oci/internal/service/ai_document"
	tf_ai_language "github.com/oracle/terraform-provider-oci/internal/service/ai_language"
	tf_ai_vision "github.com/oracle/terraform-provider-oci/internal/service/ai_vision"
	tf_analytics "github.com/oracle/terraform-provider-oci/internal/service/analytics"
	tf_announcements_service "github.com/oracle/terraform-provider-oci/internal/service/announcements_service"
	tf_apigateway "github.com/oracle/terraform-provider-oci/internal/service/apigateway"
	tf_apm "github.com/oracle/terraform-provider-oci/internal/service/apm"
	tf_apm_config "github.com/oracle/terraform-provider-oci/internal/service/apm_config"
	tf_apm_synthetics "github.com/oracle/terraform-provider-oci/internal/service/apm_synthetics"
	tf_apm_traces "github.com/oracle/terraform-provider-oci/internal/service/apm_traces"
	tf_appmgmt_control "github.com/oracle/terraform-provider-oci/internal/service/appmgmt_control"
	tf_artifacts "github.com/oracle/terraform-provider-oci/internal/service/artifacts"
	tf_audit "github.com/oracle/terraform-provider-oci/internal/service/audit"
	tf_autoscaling "github.com/oracle/terraform-provider-oci/internal/service/autoscaling"
	tf_bastion "github.com/oracle/terraform-provider-oci/internal/service/bastion"
	tf_bds "github.com/oracle/terraform-provider-oci/internal/service/bds"
	tf_blockchain "github.com/oracle/terraform-provider-oci/internal/service/blockchain"
	tf_budget "github.com/oracle/terraform-provider-oci/internal/service/budget"
	tf_capacity_management "github.com/oracle/terraform-provider-oci/internal/service/capacity_management"
	tf_certificates "github.com/oracle/terraform-provider-oci/internal/service/certificates"
	tf_certificates_management "github.com/oracle/terraform-provider-oci/internal/service/certificates_management"
	tf_cloud_bridge "github.com/oracle/terraform-provider-oci/internal/service/cloud_bridge"
	tf_cloud_guard "github.com/oracle/terraform-provider-oci/internal/service/cloud_guard"
	tf_cloud_migrations "github.com/oracle/terraform-provider-oci/internal/service/cloud_migrations"
	tf_cluster_placement_groups "github.com/oracle/terraform-provider-oci/internal/service/cluster_placement_groups"
	tf_compute_cloud_at_customer "github.com/oracle/terraform-provider-oci/internal/service/compute_cloud_at_customer"
	tf_computeinstanceagent "github.com/oracle/terraform-provider-oci/internal/service/computeinstanceagent"
	tf_container_instances "github.com/oracle/terraform-provider-oci/internal/service/container_instances"
	tf_containerengine "github.com/oracle/terraform-provider-oci/internal/service/containerengine"
	tf_core "github.com/oracle/terraform-provider-oci/internal/service/core"
	tf_data_labeling_service "github.com/oracle/terraform-provider-oci/internal/service/data_labeling_service"
	tf_data_safe "github.com/oracle/terraform-provider-oci/internal/service/data_safe"
	tf_database "github.com/oracle/terraform-provider-oci/internal/service/database"
	tf_database_management "github.com/oracle/terraform-provider-oci/internal/service/database_management"
	tf_database_migration "github.com/oracle/terraform-provider-oci/internal/service/database_migration"
	tf_database_tools "github.com/oracle/terraform-provider-oci/internal/service/database_tools"
	tf_datacatalog "github.com/oracle/terraform-provider-oci/internal/service/datacatalog"
	tf_dataflow "github.com/oracle/terraform-provider-oci/internal/service/dataflow"
	tf_dataintegration "github.com/oracle/terraform-provider-oci/internal/service/dataintegration"
	tf_datascience "github.com/oracle/terraform-provider-oci/internal/service/datascience"
	tf_delegate_access_control "github.com/oracle/terraform-provider-oci/internal/service/delegate_access_control"
	tf_demand_signal "github.com/oracle/terraform-provider-oci/internal/service/demand_signal"
	tf_desktops "github.com/oracle/terraform-provider-oci/internal/service/desktops"
	tf_devops "github.com/oracle/terraform-provider-oci/internal/service/devops"
	tf_disaster_recovery "github.com/oracle/terraform-provider-oci/internal/service/disaster_recovery"
	tf_dns "github.com/oracle/terraform-provider-oci/internal/service/dns"
	tf_em_warehouse "github.com/oracle/terraform-provider-oci/internal/service/em_warehouse"
	tf_email "github.com/oracle/terraform-provider-oci/internal/service/email"
	tf_events "github.com/oracle/terraform-provider-oci/internal/service/events"
	tf_file_storage "github.com/oracle/terraform-provider-oci/internal/service/file_storage"
	tf_fleet_apps_management "github.com/oracle/terraform-provider-oci/internal/service/fleet_apps_management"
	tf_fleet_software_update "github.com/oracle/terraform-provider-oci/internal/service/fleet_software_update"
	tf_functions "github.com/oracle/terraform-provider-oci/internal/service/functions"
	tf_fusion_apps "github.com/oracle/terraform-provider-oci/internal/service/fusion_apps"
	tf_generative_ai "github.com/oracle/terraform-provider-oci/internal/service/generative_ai"
	tf_generic_artifacts_content "github.com/oracle/terraform-provider-oci/internal/service/generic_artifacts_content"
	tf_golden_gate "github.com/oracle/terraform-provider-oci/internal/service/golden_gate"
	tf_health_checks "github.com/oracle/terraform-provider-oci/internal/service/health_checks"
	tf_identity "github.com/oracle/terraform-provider-oci/internal/service/identity"
	tf_identity_data_plane "github.com/oracle/terraform-provider-oci/internal/service/identity_data_plane"
	tf_identity_domains "github.com/oracle/terraform-provider-oci/internal/service/identity_domains"
	tf_integration "github.com/oracle/terraform-provider-oci/internal/service/integration"
	tf_jms "github.com/oracle/terraform-provider-oci/internal/service/jms"
	tf_jms_java_downloads "github.com/oracle/terraform-provider-oci/internal/service/jms_java_downloads"
	tf_kms "github.com/oracle/terraform-provider-oci/internal/service/kms"
	tf_license_manager "github.com/oracle/terraform-provider-oci/internal/service/license_manager"
	tf_limits "github.com/oracle/terraform-provider-oci/internal/service/limits"
	tf_load_balancer "github.com/oracle/terraform-provider-oci/internal/service/load_balancer"
	tf_log_analytics "github.com/oracle/terraform-provider-oci/internal/service/log_analytics"
	tf_logging "github.com/oracle/terraform-provider-oci/internal/service/logging"
	tf_management_agent "github.com/oracle/terraform-provider-oci/internal/service/management_agent"
	tf_management_dashboard "github.com/oracle/terraform-provider-oci/internal/service/management_dashboard"
	tf_marketplace "github.com/oracle/terraform-provider-oci/internal/service/marketplace"
	tf_media_services "github.com/oracle/terraform-provider-oci/internal/service/media_services"
	tf_metering_computation "github.com/oracle/terraform-provider-oci/internal/service/metering_computation"
	tf_monitoring "github.com/oracle/terraform-provider-oci/internal/service/monitoring"
	tf_mysql "github.com/oracle/terraform-provider-oci/internal/service/mysql"
	tf_network_firewall "github.com/oracle/terraform-provider-oci/internal/service/network_firewall"
	tf_network_load_balancer "github.com/oracle/terraform-provider-oci/internal/service/network_load_balancer"
	tf_nosql "github.com/oracle/terraform-provider-oci/internal/service/nosql"
	tf_objectstorage "github.com/oracle/terraform-provider-oci/internal/service/objectstorage"
	tf_oce "github.com/oracle/terraform-provider-oci/internal/service/oce"
	tf_ocvp "github.com/oracle/terraform-provider-oci/internal/service/ocvp"
	tf_oda "github.com/oracle/terraform-provider-oci/internal/service/oda"
	tf_onesubscription "github.com/oracle/terraform-provider-oci/internal/service/onesubscription"
	tf_ons "github.com/oracle/terraform-provider-oci/internal/service/ons"
	tf_opa "github.com/oracle/terraform-provider-oci/internal/service/opa"
	tf_opensearch "github.com/oracle/terraform-provider-oci/internal/service/opensearch"
	tf_operator_access_control "github.com/oracle/terraform-provider-oci/internal/service/operator_access_control"
	tf_opsi "github.com/oracle/terraform-provider-oci/internal/service/opsi"
	tf_optimizer "github.com/oracle/terraform-provider-oci/internal/service/optimizer"
	tf_os_management_hub "github.com/oracle/terraform-provider-oci/internal/service/os_management_hub"
	tf_osmanagement "github.com/oracle/terraform-provider-oci/internal/service/osmanagement"
	tf_osp_gateway "github.com/oracle/terraform-provider-oci/internal/service/osp_gateway"
	tf_osub_billing_schedule "github.com/oracle/terraform-provider-oci/internal/service/osub_billing_schedule"
	tf_osub_organization_subscription "github.com/oracle/terraform-provider-oci/internal/service/osub_organization_subscription"
	tf_osub_subscription "github.com/oracle/terraform-provider-oci/internal/service/osub_subscription"
	tf_osub_usage "github.com/oracle/terraform-provider-oci/internal/service/osub_usage"
	tf_psql "github.com/oracle/terraform-provider-oci/internal/service/psql"
	tf_queue "github.com/oracle/terraform-provider-oci/internal/service/queue"
	tf_recovery "github.com/oracle/terraform-provider-oci/internal/service/recovery"
	tf_redis "github.com/oracle/terraform-provider-oci/internal/service/redis"
	tf_resource_scheduler "github.com/oracle/terraform-provider-oci/internal/service/resource_scheduler"
	tf_resourcemanager "github.com/oracle/terraform-provider-oci/internal/service/resourcemanager"
	tf_sch "github.com/oracle/terraform-provider-oci/internal/service/sch"
	tf_secrets "github.com/oracle/terraform-provider-oci/internal/service/secrets"
	tf_security_attribute "github.com/oracle/terraform-provider-oci/internal/service/security_attribute"
	tf_service_catalog "github.com/oracle/terraform-provider-oci/internal/service/service_catalog"
	tf_service_manager_proxy "github.com/oracle/terraform-provider-oci/internal/service/service_manager_proxy"
	tf_service_mesh "github.com/oracle/terraform-provider-oci/internal/service/service_mesh"
	tf_stack_monitoring "github.com/oracle/terraform-provider-oci/internal/service/stack_monitoring"
	tf_streaming "github.com/oracle/terraform-provider-oci/internal/service/streaming"
	tf_usage_proxy "github.com/oracle/terraform-provider-oci/internal/service/usage_proxy"
	tf_vault "github.com/oracle/terraform-provider-oci/internal/service/vault"
	tf_vbs_inst "github.com/oracle/terraform-provider-oci/internal/service/vbs_inst"
	tf_visual_builder "github.com/oracle/terraform-provider-oci/internal/service/visual_builder"
	tf_vn_monitoring "github.com/oracle/terraform-provider-oci/internal/service/vn_monitoring"
	tf_vulnerability_scanning "github.com/oracle/terraform-provider-oci/internal/service/vulnerability_scanning"
	tf_waa "github.com/oracle/terraform-provider-oci/internal/service/waa"
	tf_waas "github.com/oracle/terraform-provider-oci/internal/service/waas"
	tf_waf "github.com/oracle/terraform-provider-oci/internal/service/waf"
	tf_zpr "github.com/oracle/terraform-provider-oci/internal/service/zpr"
)

func init() {
	if common.CheckForEnabledServices("adm") {
		tf_adm.RegisterDatasource()
	}
	if common.CheckForEnabledServices("aianomalydetection") {
		tf_ai_anomaly_detection.RegisterDatasource()
	}
	if common.CheckForEnabledServices("aidocument") {
		tf_ai_document.RegisterDatasource()
	}
	if common.CheckForEnabledServices("ailanguage") {
		tf_ai_language.RegisterDatasource()
	}
	if common.CheckForEnabledServices("aivision") {
		tf_ai_vision.RegisterDatasource()
	}
	if common.CheckForEnabledServices("analytics") {
		tf_analytics.RegisterDatasource()
	}
	if common.CheckForEnabledServices("announcementsservice") {
		tf_announcements_service.RegisterDatasource()
	}
	if common.CheckForEnabledServices("apigateway") {
		tf_apigateway.RegisterDatasource()
	}
	if common.CheckForEnabledServices("apm") {
		tf_apm.RegisterDatasource()
	}
	if common.CheckForEnabledServices("apmconfig") {
		tf_apm_config.RegisterDatasource()
	}
	if common.CheckForEnabledServices("apmsynthetics") {
		tf_apm_synthetics.RegisterDatasource()
	}
	if common.CheckForEnabledServices("apmtraces") {
		tf_apm_traces.RegisterDatasource()
	}
	if common.CheckForEnabledServices("appmgmtcontrol") {
		tf_appmgmt_control.RegisterDatasource()
	}
	if common.CheckForEnabledServices("artifacts") {
		tf_artifacts.RegisterDatasource()
	}
	if common.CheckForEnabledServices("audit") {
		tf_audit.RegisterDatasource()
	}
	if common.CheckForEnabledServices("autoscaling") {
		tf_autoscaling.RegisterDatasource()
	}
	if common.CheckForEnabledServices("bastion") {
		tf_bastion.RegisterDatasource()
	}
	if common.CheckForEnabledServices("bds") {
		tf_bds.RegisterDatasource()
	}
	if common.CheckForEnabledServices("blockchain") {
		tf_blockchain.RegisterDatasource()
	}
	if common.CheckForEnabledServices("budget") {
		tf_budget.RegisterDatasource()
	}
	if common.CheckForEnabledServices("capacitymanagement") {
		tf_capacity_management.RegisterDatasource()
	}
	if common.CheckForEnabledServices("certificates") {
		tf_certificates.RegisterDatasource()
	}
	if common.CheckForEnabledServices("certificatesmanagement") {
		tf_certificates_management.RegisterDatasource()
	}
	if common.CheckForEnabledServices("cloudbridge") {
		tf_cloud_bridge.RegisterDatasource()
	}
	if common.CheckForEnabledServices("cloudguard") {
		tf_cloud_guard.RegisterDatasource()
	}
	if common.CheckForEnabledServices("cloudmigrations") {
		tf_cloud_migrations.RegisterDatasource()
	}
	if common.CheckForEnabledServices("clusterplacementgroups") {
		tf_cluster_placement_groups.RegisterDatasource()
	}
	if common.CheckForEnabledServices("computecloudatcustomer") {
		tf_compute_cloud_at_customer.RegisterDatasource()
	}
	if common.CheckForEnabledServices("computeinstanceagent") {
		tf_computeinstanceagent.RegisterDatasource()
	}
	if common.CheckForEnabledServices("containerinstances") {
		tf_container_instances.RegisterDatasource()
	}
	if common.CheckForEnabledServices("containerengine") {
		tf_containerengine.RegisterDatasource()
	}
	if common.CheckForEnabledServices("core") {
		tf_core.RegisterDatasource()
	}
	if common.CheckForEnabledServices("datalabelingservice") {
		tf_data_labeling_service.RegisterDatasource()
	}
	if common.CheckForEnabledServices("datasafe") {
		tf_data_safe.RegisterDatasource()
	}
	if common.CheckForEnabledServices("database") {
		tf_database.RegisterDatasource()
	}
	if common.CheckForEnabledServices("databasemanagement") {
		tf_database_management.RegisterDatasource()
	}
	if common.CheckForEnabledServices("databasemigration") {
		tf_database_migration.RegisterDatasource()
	}
	if common.CheckForEnabledServices("databasetools") {
		tf_database_tools.RegisterDatasource()
	}
	if common.CheckForEnabledServices("datacatalog") {
		tf_datacatalog.RegisterDatasource()
	}
	if common.CheckForEnabledServices("dataflow") {
		tf_dataflow.RegisterDatasource()
	}
	if common.CheckForEnabledServices("dataintegration") {
		tf_dataintegration.RegisterDatasource()
	}
	if common.CheckForEnabledServices("datascience") {
		tf_datascience.RegisterDatasource()
	}
	if common.CheckForEnabledServices("delegateaccesscontrol") {
		tf_delegate_access_control.RegisterDatasource()
	}
	if common.CheckForEnabledServices("demandsignal") {
		tf_demand_signal.RegisterDatasource()
	}
	if common.CheckForEnabledServices("desktops") {
		tf_desktops.RegisterDatasource()
	}
	if common.CheckForEnabledServices("devops") {
		tf_devops.RegisterDatasource()
	}
	if common.CheckForEnabledServices("disasterrecovery") {
		tf_disaster_recovery.RegisterDatasource()
	}
	if common.CheckForEnabledServices("dns") {
		tf_dns.RegisterDatasource()
	}
	if common.CheckForEnabledServices("emwarehouse") {
		tf_em_warehouse.RegisterDatasource()
	}
	if common.CheckForEnabledServices("email") {
		tf_email.RegisterDatasource()
	}
	if common.CheckForEnabledServices("events") {
		tf_events.RegisterDatasource()
	}
	if common.CheckForEnabledServices("filestorage") {
		tf_file_storage.RegisterDatasource()
	}
	if common.CheckForEnabledServices("fleetsoftwareupdate") {
		tf_fleet_software_update.RegisterDatasource()
	}
	if common.CheckForEnabledServices("fleetappsmanagement") {
		tf_fleet_apps_management.RegisterDatasource()
	}
	if common.CheckForEnabledServices("functions") {
		tf_functions.RegisterDatasource()
	}
	if common.CheckForEnabledServices("fusionapps") {
		tf_fusion_apps.RegisterDatasource()
	}
	if common.CheckForEnabledServices("generativeai") {
		tf_generative_ai.RegisterDatasource()
	}
	if common.CheckForEnabledServices("genericartifactscontent") {
		tf_generic_artifacts_content.RegisterDatasource()
	}
	if common.CheckForEnabledServices("goldengate") {
		tf_golden_gate.RegisterDatasource()
	}
	if common.CheckForEnabledServices("healthchecks") {
		tf_health_checks.RegisterDatasource()
	}
	if common.CheckForEnabledServices("identity") {
		tf_identity.RegisterDatasource()
	}
	if common.CheckForEnabledServices("identitydataplane") {
		tf_identity_data_plane.RegisterDatasource()
	}
	if common.CheckForEnabledServices("identitydomains") {
		tf_identity_domains.RegisterDatasource()
	}
	if common.CheckForEnabledServices("integration") {
		tf_integration.RegisterDatasource()
	}
	if common.CheckForEnabledServices("jms") {
		tf_jms.RegisterDatasource()
	}
	if common.CheckForEnabledServices("jmsjavadownloads") {
		tf_jms_java_downloads.RegisterDatasource()
	}
	if common.CheckForEnabledServices("kms") {
		tf_kms.RegisterDatasource()
	}
	if common.CheckForEnabledServices("licensemanager") {
		tf_license_manager.RegisterDatasource()
	}
	if common.CheckForEnabledServices("limits") {
		tf_limits.RegisterDatasource()
	}
	if common.CheckForEnabledServices("loadbalancer") {
		tf_load_balancer.RegisterDatasource()
	}
	if common.CheckForEnabledServices("loganalytics") {
		tf_log_analytics.RegisterDatasource()
	}
	if common.CheckForEnabledServices("logging") {
		tf_logging.RegisterDatasource()
	}
	if common.CheckForEnabledServices("managementagent") {
		tf_management_agent.RegisterDatasource()
	}
	if common.CheckForEnabledServices("managementdashboard") {
		tf_management_dashboard.RegisterDatasource()
	}
	if common.CheckForEnabledServices("marketplace") {
		tf_marketplace.RegisterDatasource()
	}
	if common.CheckForEnabledServices("mediaservices") {
		tf_media_services.RegisterDatasource()
	}
	if common.CheckForEnabledServices("meteringcomputation") {
		tf_metering_computation.RegisterDatasource()
	}
	if common.CheckForEnabledServices("monitoring") {
		tf_monitoring.RegisterDatasource()
	}
	if common.CheckForEnabledServices("mysql") {
		tf_mysql.RegisterDatasource()
	}
	if common.CheckForEnabledServices("networkfirewall") {
		tf_network_firewall.RegisterDatasource()
	}
	if common.CheckForEnabledServices("networkloadbalancer") {
		tf_network_load_balancer.RegisterDatasource()
	}
	if common.CheckForEnabledServices("nosql") {
		tf_nosql.RegisterDatasource()
	}
	if common.CheckForEnabledServices("objectstorage") {
		tf_objectstorage.RegisterDatasource()
	}
	if common.CheckForEnabledServices("oce") {
		tf_oce.RegisterDatasource()
	}
	if common.CheckForEnabledServices("ocvp") {
		tf_ocvp.RegisterDatasource()
	}
	if common.CheckForEnabledServices("oda") {
		tf_oda.RegisterDatasource()
	}
	if common.CheckForEnabledServices("onesubscription") {
		tf_onesubscription.RegisterDatasource()
	}
	if common.CheckForEnabledServices("ons") {
		tf_ons.RegisterDatasource()
	}
	if common.CheckForEnabledServices("opa") {
		tf_opa.RegisterDatasource()
	}
	if common.CheckForEnabledServices("opensearch") {
		tf_opensearch.RegisterDatasource()
	}
	if common.CheckForEnabledServices("operatoraccesscontrol") {
		tf_operator_access_control.RegisterDatasource()
	}
	if common.CheckForEnabledServices("opsi") {
		tf_opsi.RegisterDatasource()
	}
	if common.CheckForEnabledServices("optimizer") {
		tf_optimizer.RegisterDatasource()
	}
	if common.CheckForEnabledServices("os_management_hub") {
		tf_os_management_hub.RegisterDatasource()
	}
	if common.CheckForEnabledServices("osmanagement") {
		tf_osmanagement.RegisterDatasource()
	}
	if common.CheckForEnabledServices("ospgateway") {
		tf_osp_gateway.RegisterDatasource()
	}
	if common.CheckForEnabledServices("osubbillingschedule") {
		tf_osub_billing_schedule.RegisterDatasource()
	}
	if common.CheckForEnabledServices("osuborganizationsubscription") {
		tf_osub_organization_subscription.RegisterDatasource()
	}
	if common.CheckForEnabledServices("osubsubscription") {
		tf_osub_subscription.RegisterDatasource()
	}
	if common.CheckForEnabledServices("osubusage") {
		tf_osub_usage.RegisterDatasource()
	}
	if common.CheckForEnabledServices("psql") {
		tf_psql.RegisterDatasource()
	}
	if common.CheckForEnabledServices("queue") {
		tf_queue.RegisterDatasource()
	}
	if common.CheckForEnabledServices("recovery") {
		tf_recovery.RegisterDatasource()
	}
	if common.CheckForEnabledServices("redis") {
		tf_redis.RegisterDatasource()
	}
	if common.CheckForEnabledServices("resourcescheduler") {
		tf_resource_scheduler.RegisterDatasource()
	}
	if common.CheckForEnabledServices("resourcemanager") {
		tf_resourcemanager.RegisterDatasource()
	}
	if common.CheckForEnabledServices("sch") {
		tf_sch.RegisterDatasource()
	}
	if common.CheckForEnabledServices("secrets") {
		tf_secrets.RegisterDatasource()
	}
	if common.CheckForEnabledServices("securityattribute") {
		tf_security_attribute.RegisterDatasource()
	}
	if common.CheckForEnabledServices("servicecatalog") {
		tf_service_catalog.RegisterDatasource()
	}
	if common.CheckForEnabledServices("servicemanagerproxy") {
		tf_service_manager_proxy.RegisterDatasource()
	}
	if common.CheckForEnabledServices("servicemesh") {
		tf_service_mesh.RegisterDatasource()
	}
	if common.CheckForEnabledServices("stackmonitoring") {
		tf_stack_monitoring.RegisterDatasource()
	}
	if common.CheckForEnabledServices("streaming") {
		tf_streaming.RegisterDatasource()
	}
	if common.CheckForEnabledServices("usageproxy") {
		tf_usage_proxy.RegisterDatasource()
	}
	if common.CheckForEnabledServices("vault") {
		tf_vault.RegisterDatasource()
	}
	if common.CheckForEnabledServices("vbsinst") {
		tf_vbs_inst.RegisterDatasource()
	}
	if common.CheckForEnabledServices("visualbuilder") {
		tf_visual_builder.RegisterDatasource()
	}
	if common.CheckForEnabledServices("vnmonitoring") {
		tf_vn_monitoring.RegisterDatasource()
	}
	if common.CheckForEnabledServices("vulnerabilityscanning") {
		tf_vulnerability_scanning.RegisterDatasource()
	}
	if common.CheckForEnabledServices("waa") {
		tf_waa.RegisterDatasource()
	}
	if common.CheckForEnabledServices("waas") {
		tf_waas.RegisterDatasource()
	}
	if common.CheckForEnabledServices("waf") {
		tf_waf.RegisterDatasource()
	}
	if common.CheckForEnabledServices("zpr") {
		tf_zpr.RegisterDatasource()
	}
}
