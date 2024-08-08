// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_data_safe_alert", DataSafeAlertDataSource())
	tfresource.RegisterDatasource("oci_data_safe_alert_analytic", DataSafeAlertAnalyticDataSource())
	tfresource.RegisterDatasource("oci_data_safe_alert_policies", DataSafeAlertPoliciesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_alert_policy", DataSafeAlertPolicyDataSource())
	tfresource.RegisterDatasource("oci_data_safe_alert_policy_rule", DataSafeAlertPolicyRuleDataSource())
	tfresource.RegisterDatasource("oci_data_safe_alert_policy_rules", DataSafeAlertPolicyRulesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_alerts", DataSafeAlertsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_archive_retrieval", DataSafeAuditArchiveRetrievalDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_archive_retrievals", DataSafeAuditArchiveRetrievalsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_event", DataSafeAuditEventDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_event_analytic", DataSafeAuditEventAnalyticDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_events", DataSafeAuditEventsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_policies", DataSafeAuditPoliciesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_policy", DataSafeAuditPolicyDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_profile", DataSafeAuditProfileDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_profile_analytic", DataSafeAuditProfileAnalyticDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_profile_available_audit_volume", DataSafeAuditProfileAvailableAuditVolumeDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_profile_available_audit_volumes", DataSafeAuditProfileAvailableAuditVolumesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_profile_collected_audit_volume", DataSafeAuditProfileCollectedAuditVolumeDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_profile_collected_audit_volumes", DataSafeAuditProfileCollectedAuditVolumesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_profiles", DataSafeAuditProfilesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_trail", DataSafeAuditTrailDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_trail_analytic", DataSafeAuditTrailAnalyticDataSource())
	tfresource.RegisterDatasource("oci_data_safe_audit_trails", DataSafeAuditTrailsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_compatible_formats_for_data_type", DataSafeCompatibleFormatsForDataTypeDataSource())
	tfresource.RegisterDatasource("oci_data_safe_compatible_formats_for_sensitive_type", DataSafeCompatibleFormatsForSensitiveTypeDataSource())
	tfresource.RegisterDatasource("oci_data_safe_data_safe_configuration", DataSafeDataSafeConfigurationDataSource())
	tfresource.RegisterDatasource("oci_data_safe_data_safe_private_endpoint", DataSafeDataSafePrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_data_safe_data_safe_private_endpoints", DataSafeDataSafePrivateEndpointsDataSource())

	tfresource.RegisterDatasource("oci_data_safe_discovery_analytic", DataSafeDiscoveryAnalyticDataSource())

	tfresource.RegisterDatasource("oci_data_safe_discovery_analytics", DataSafeDiscoveryAnalyticsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_discovery_job", DataSafeDiscoveryJobDataSource())
	tfresource.RegisterDatasource("oci_data_safe_discovery_jobs", DataSafeDiscoveryJobsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_discovery_jobs_result", DataSafeDiscoveryJobsResultDataSource())
	tfresource.RegisterDatasource("oci_data_safe_discovery_jobs_results", DataSafeDiscoveryJobsResultsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_library_masking_format", DataSafeLibraryMaskingFormatDataSource())
	tfresource.RegisterDatasource("oci_data_safe_library_masking_formats", DataSafeLibraryMaskingFormatsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_list_user_grants", DataSafeListUserGrantsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_analytic", DataSafeMaskingAnalyticDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_analytics", DataSafeMaskingAnalyticsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_policies", DataSafeMaskingPoliciesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_policies_masking_column", DataSafeMaskingPoliciesMaskingColumnDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_policies_masking_columns", DataSafeMaskingPoliciesMaskingColumnsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_policy", DataSafeMaskingPolicyDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_policy_health_report", DataSafeMaskingPolicyHealthReportDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_policy_health_report_logs", DataSafeMaskingPolicyHealthReportLogsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_policy_health_reports", DataSafeMaskingPolicyHealthReportsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_policy_masking_objects", DataSafeMaskingPolicyMaskingObjectsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_policy_masking_schemas", DataSafeMaskingPolicyMaskingSchemasDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_report", DataSafeMaskingReportDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_reports", DataSafeMaskingReportsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_reports_masked_column", DataSafeMaskingReportsMaskedColumnDataSource())
	tfresource.RegisterDatasource("oci_data_safe_masking_reports_masked_columns", DataSafeMaskingReportsMaskedColumnsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_on_prem_connector", DataSafeOnPremConnectorDataSource())
	tfresource.RegisterDatasource("oci_data_safe_on_prem_connectors", DataSafeOnPremConnectorsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_report", DataSafeReportDataSource())
	tfresource.RegisterDatasource("oci_data_safe_report_content", DataSafeReportContentDataSource())
	tfresource.RegisterDatasource("oci_data_safe_report_definition", DataSafeReportDefinitionDataSource())
	tfresource.RegisterDatasource("oci_data_safe_report_definitions", DataSafeReportDefinitionsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_reports", DataSafeReportsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sdm_masking_policy_difference", DataSafeSdmMaskingPolicyDifferenceDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sdm_masking_policy_difference_difference_column", DataSafeSdmMaskingPolicyDifferenceDifferenceColumnDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sdm_masking_policy_difference_difference_columns", DataSafeSdmMaskingPolicyDifferenceDifferenceColumnsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sdm_masking_policy_differences", DataSafeSdmMaskingPolicyDifferencesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_assessment", DataSafeSecurityAssessmentDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_assessment_comparison", DataSafeSecurityAssessmentComparisonDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_assessment_finding_analytics", DataSafeSecurityAssessmentFindingAnalyticsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_assessment_finding", DataSafeSecurityAssessmentFindingsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_assessment_findings", DataSafeSecurityAssessmentFindingsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_assessment_security_feature_analytics", DataSafeSecurityAssessmentSecurityFeatureAnalyticsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_assessment_security_features", DataSafeSecurityAssessmentSecurityFeaturesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_assessment_findings_change_audit_logs", DataSafeSecurityAssessmentFindingsChangeAuditLogsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_assessments", DataSafeSecurityAssessmentsDataSource())

	tfresource.RegisterDatasource("oci_data_safe_security_policies", DataSafeSecurityPoliciesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_policy", DataSafeSecurityPolicyDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_policy_deployment", DataSafeSecurityPolicyDeploymentDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_policy_deployment_security_policy_entry_state", DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_policy_deployment_security_policy_entry_states", DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStatesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_policy_deployments", DataSafeSecurityPolicyDeploymentsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_policy_report", DataSafeSecurityPolicyReportDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_policy_report_database_table_access_entries", DataSafeSecurityPolicyReportDatabaseTableAccessEntriesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_policy_report_database_table_access_entry", DataSafeSecurityPolicyReportDatabaseTableAccessEntryDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_policy_report_database_view_access_entries", DataSafeSecurityPolicyReportDatabaseViewAccessEntriesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_policy_report_database_view_access_entry", DataSafeSecurityPolicyReportDatabaseViewAccessEntryDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_policy_report_role_grant_paths", DataSafeSecurityPolicyReportRoleGrantPathsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_security_policy_reports", DataSafeSecurityPolicyReportsDataSource())

	tfresource.RegisterDatasource("oci_data_safe_sensitive_data_model", DataSafeSensitiveDataModelDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sensitive_data_model_sensitive_objects", DataSafeSensitiveDataModelSensitiveObjectsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sensitive_data_model_sensitive_schemas", DataSafeSensitiveDataModelSensitiveSchemasDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sensitive_data_model_sensitive_types", DataSafeSensitiveDataModelSensitiveTypesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sensitive_data_models", DataSafeSensitiveDataModelsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sensitive_data_models_sensitive_column", DataSafeSensitiveDataModelsSensitiveColumnDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sensitive_data_models_sensitive_columns", DataSafeSensitiveDataModelsSensitiveColumnsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sensitive_type", DataSafeSensitiveTypeDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sensitive_types", DataSafeSensitiveTypesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sql_collection", DataSafeSqlCollectionDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sql_collection_analytics", DataSafeSqlCollectionAnalyticsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sql_collection_log_insights", DataSafeSqlCollectionLogInsightsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sql_collections", DataSafeSqlCollectionsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sql_firewall_allowed_sql", DataSafeSqlFirewallAllowedSqlDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sql_firewall_allowed_sql_analytics", DataSafeSqlFirewallAllowedSqlAnalyticsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sql_firewall_allowed_sqls", DataSafeSqlFirewallAllowedSqlsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sql_firewall_policies", DataSafeSqlFirewallPoliciesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sql_firewall_policy", DataSafeSqlFirewallPolicyDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sql_firewall_policy_analytics", DataSafeSqlFirewallPolicyAnalyticsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sql_firewall_violation_analytics", DataSafeSqlFirewallViolationAnalyticsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_sql_firewall_violations", DataSafeSqlFirewallViolationsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_target_alert_policy_association", DataSafeTargetAlertPolicyAssociationDataSource())
	tfresource.RegisterDatasource("oci_data_safe_target_alert_policy_associations", DataSafeTargetAlertPolicyAssociationsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_target_database", DataSafeTargetDatabaseDataSource())
	tfresource.RegisterDatasource("oci_data_safe_target_database_peer_target_database", DataSafeTargetDatabasePeerTargetDatabaseDataSource())
	tfresource.RegisterDatasource("oci_data_safe_target_database_peer_target_databases", DataSafeTargetDatabasePeerTargetDatabasesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_target_database_roles", DataSafeTargetDatabaseRolesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_target_database_role", DataSafeTargetDatabaseRolesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_target_databases", DataSafeTargetDatabasesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_target_databases_columns", DataSafeTargetDatabasesColumnsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_target_databases_schemas", DataSafeTargetDatabasesSchemasDataSource())
	tfresource.RegisterDatasource("oci_data_safe_target_databases_tables", DataSafeTargetDatabasesTablesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_user_assessment", DataSafeUserAssessmentDataSource())
	tfresource.RegisterDatasource("oci_data_safe_user_assessment_comparison", DataSafeUserAssessmentComparisonDataSource())
	tfresource.RegisterDatasource("oci_data_safe_user_assessment_profile_analytics", DataSafeUserAssessmentProfileAnalyticsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_user_assessment_profiles", DataSafeUserAssessmentProfilesDataSource())
	tfresource.RegisterDatasource("oci_data_safe_user_assessment_user_access_analytics", DataSafeUserAssessmentUserAccessAnalyticsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_user_assessment_user_analytics", DataSafeUserAssessmentUserAnalyticsDataSource())
	tfresource.RegisterDatasource("oci_data_safe_user_assessment_users", DataSafeUserAssessmentUsersDataSource())
	tfresource.RegisterDatasource("oci_data_safe_user_assessments", DataSafeUserAssessmentsDataSource())
}
