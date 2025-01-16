// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "compartment_id" {}

variable "unified_agent_configuration_defined_tags_value" {
  default = "value"
}

variable "unified_agent_configuration_description" {
  default = "description2"
}

variable "unified_agent_configuration_display_name" {
  default = "tf-agentConfigName1"
}

variable "unified_agent_configuration_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "unified_agent_configuration_group_association_group_list" {
  default = [""]
}

variable "unified_agent_configuration_is_compartment_id_in_subtree" {
  default = false
}

variable "unified_agent_configuration_is_enabled" {
  default = true
}

variable "unified_agent_configuration_service_configuration_configuration_type" {
  default = "LOGGING"
}

variable "unified_agent_configuration_service_configuration_sources_channels" {
  default = ["Security"]
}

variable "unified_agent_configuration_service_configuration_sources_name" {
  default = "name"
}

variable "unified_agent_configuration_service_configuration_sources_parser_delimiter" {
  default = "delimiter"
}

variable "unified_agent_configuration_service_configuration_sources_parser_expression" {
  default = "expression"
}

variable "unified_agent_configuration_service_configuration_sources_parser_field_time_key" {
  default = "fieldTimeKey"
}

variable "unified_agent_configuration_service_configuration_sources_parser_format" {
  default = []
}

variable "unified_agent_configuration_service_configuration_sources_parser_format_firstline" {
  default = "formatFirstline"
}

variable "unified_agent_configuration_service_configuration_sources_parser_grok_failure_key" {
  default = "grokFailureKey"
}

variable "unified_agent_configuration_service_configuration_sources_parser_grok_name_key" {
  default = "grokNameKey"
}

variable "unified_agent_configuration_service_configuration_sources_parser_is_estimate_current_event" {
  default = false
}

variable "unified_agent_configuration_service_configuration_sources_parser_is_keep_time_key" {
  default = false
}

variable "unified_agent_configuration_service_configuration_sources_parser_is_null_empty_string" {
  default = false
}

variable "unified_agent_configuration_service_configuration_sources_parser_is_support_colonless_ident" {
  default = false
}

variable "unified_agent_configuration_service_configuration_sources_parser_is_with_priority" {
  default = false
}

variable "unified_agent_configuration_service_configuration_sources_parser_keys" {
  default = []
}

variable "unified_agent_configuration_service_configuration_sources_parser_message_format" {
  default = "RFC3164"
}

variable "unified_agent_configuration_service_configuration_sources_parser_message_key" {
  default = "messageKey"
}

variable "unified_agent_configuration_service_configuration_sources_parser_multi_line_start_regexp" {
  default = "multiLineStartRegexp"
}

variable "unified_agent_configuration_service_configuration_sources_parser_null_value_pattern" {
  default = "nullValuePattern"
}

variable "unified_agent_configuration_service_configuration_sources_parser_parser_type" {
  default = "AUDITD"
}

variable "unified_agent_configuration_service_configuration_sources_parser_patterns_field_time_format" {
  default = "fieldTimeFormat"
}

variable "unified_agent_configuration_service_configuration_sources_parser_patterns_field_time_key" {
  default = "fieldTimeKey"
}

variable "unified_agent_configuration_service_configuration_sources_parser_patterns_field_time_zone" {
  default = "fieldTimeZone"
}

variable "unified_agent_configuration_service_configuration_sources_parser_patterns_name" {
  default = "name"
}

variable "unified_agent_configuration_service_configuration_sources_parser_patterns_pattern" {
  default = "pattern"
}

variable "unified_agent_configuration_service_configuration_sources_parser_rfc5424time_format" {
  default = "rfc5424TimeFormat"
}

variable "unified_agent_configuration_service_configuration_sources_parser_syslog_parser_type" {
  default = "STRING"
}

variable "unified_agent_configuration_service_configuration_sources_parser_time_format" {
  default = "timeFormat"
}

variable "unified_agent_configuration_service_configuration_sources_parser_time_type" {
  default = "FLOAT"
}

variable "unified_agent_configuration_service_configuration_sources_parser_timeout_in_milliseconds" {
  default = 10
}

variable "unified_agent_configuration_service_configuration_sources_parser_types" {
  default = "types"
}

variable "unified_agent_configuration_service_configuration_sources_paths" {
  default = []
}

variable "unified_agent_configuration_service_configuration_sources_source_type" {
  default = "LOG_TAIL"
}

variable "unified_agent_configuration_state" {
  default = "ACTIVE"
}

variable "log_group_defined_tags_value" {
  default = "value2"
}

variable "test_log_group_id" {}
variable "test_log_id" {}

resource "oci_logging_unified_agent_configuration" "test_unified_agent_configuration" {
  #Required
  compartment_id = var.compartment_id
  is_enabled     = var.unified_agent_configuration_is_enabled
  service_configuration {
    #Required
    configuration_type = var.unified_agent_configuration_service_configuration_configuration_type

    #Required field destination for service_configuration
    destination {
      #Required field for destination
      log_object_id = var.test_log_id
      operational_metrics_configuration {
        destination {
          compartment_id = var.compartment_id
        }
        source {
          type = "UMA_METRICS"
          record_input {
            namespace = "tf_test_namespace"
            resource_group = "tf-test-resource-group"
          }
        }
      }
    }
    sources {
      #Required
      source_type = var.unified_agent_configuration_service_configuration_sources_source_type

      #Optional
      # channels for windows only
      # channels = var.unified_agent_configuration_service_configuration_sources_channels
      name     = var.unified_agent_configuration_service_configuration_sources_name
      parser {
        parser_type = "CRI"
        is_merge_cri_fields = false
        nested_parser {
          time_format = "%Y-%m-%dT%H:%M:%S.%L%z"
          field_time_key = "time"
          is_keep_time_key = true
        }
      }
      paths = ["/var/log/*"]
    }
    sources {
      #Required
      source_type = var.unified_agent_configuration_service_configuration_sources_source_type

      #Optional
      # channels for windows only
      # channels = var.unified_agent_configuration_service_configuration_sources_channels
      name     = var.unified_agent_configuration_service_configuration_sources_name
      parser {
        parser_type = "NONE"
      }
      paths = ["/var/log/*"]
    }
  }

  description   = var.unified_agent_configuration_description
  display_name  = var.unified_agent_configuration_display_name
  freeform_tags = var.unified_agent_configuration_freeform_tags
  group_association {

    #Optional
    group_list = ["ocid1.dynamicgroup.oc1..aaaaaaaatqbpurg4jtr57dthka4lbykvsqajjmynecixwgsfgu2z36wf4kgq"]
  }

  lifecycle {
    ignore_changes = [ defined_tags ]
  }
}

data "oci_logging_unified_agent_configurations" "test_unified_agent_configurations" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name                 = var.unified_agent_configuration_display_name
  group_id                     = var.test_log_group_id
  is_compartment_id_in_subtree = var.unified_agent_configuration_is_compartment_id_in_subtree
  log_id                       = var.test_log_id
  state                        = var.unified_agent_configuration_state
}

resource "oci_logging_unified_agent_configuration" "test_unified_agent_configuration_1" {
  #Required
  compartment_id = var.compartment_id
  is_enabled     = var.unified_agent_configuration_is_enabled
  service_configuration {
    #Required
    configuration_type = "LOGGING"

    #Required field destination for service_configuration
    destination {
      #Required field for destination
      log_object_id = var.test_log_id
      operational_metrics_configuration {
        destination {
          compartment_id = var.compartment_id
        }
        source {
          type = "UMA_METRICS"
          record_input {
            namespace = "tf_test_namespace"
            resource_group = "tf-test-resource-group"
          }
        }
      }
    }
    sources {
      #Required
      source_type = "LOG_TAIL"
      name     = var.unified_agent_configuration_service_configuration_sources_name
      advanced_options {
          is_read_from_head = true
      }
      parser {
        parser_type = "JSON"
        field_time_key = "time"
        is_keep_time_key = true
        timeout_in_milliseconds = 1000
      }
      paths = ["/var/log/*"]
    }
    sources {
      #Required
      source_type = "WINDOWS_EVENT_LOG"
      name     = "windows_event_test"
      channels = ["system"]
    }
    # could test custom_plugin sources here too
    unified_agent_configuration_filter {
      filter_type = "GREP_FILTER"
      allow_list {
        key = "key"
        pattern = "value"
      }
      deny_list {
        key = "key"
        pattern = "value"
      }
      name = "test"
    }
  }

  #Optional
  description   = var.unified_agent_configuration_description
  display_name  = "test_unified_agent_configuration_1"
  freeform_tags = var.unified_agent_configuration_freeform_tags
  group_association {
    #Optional
    group_list = ["ocid1.dynamicgroup.oc1..aaaaaaaatqbpurg4jtr57dthka4lbykvsqajjmynecixwgsfgu2z36wf4kgq"]
  }

  lifecycle {
    ignore_changes = [ defined_tags ]
  }
}
data "oci_logging_unified_agent_configurations" "test_unified_agent_configurations_KUBERNETES" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name  = "test_unified_agent_configuration_1"
  group_id                     = var.test_log_group_id
  is_compartment_id_in_subtree = var.unified_agent_configuration_is_compartment_id_in_subtree
  log_id                       = var.test_log_id
  state                        = var.unified_agent_configuration_state
}

resource "oci_logging_unified_agent_configuration" "test_unified_agent_configuration_monitoring_KUBERNETES" {
  #Required
  compartment_id = var.compartment_id
  is_enabled     = var.unified_agent_configuration_is_enabled
  service_configuration {
    #Required
    configuration_type = "MONITORING"

    application_configurations {
      destination {
        compartment_id = var.compartment_id
        metrics_namespace = "tf_test_namespace"
      }
      source_type = "KUBERNETES"
      source {
        name = "kubernetes_source"
        scrape_targets {
          k8s_namespace = "kube-system"
          resource_group = "tf-test-resource-group"
          resource_type = "PODS"
          service_name = "kubernetes"
        }
      }
      unified_agent_configuration_filter {
        filter_type = "KUBERNETES_FILTER"
        name = "kubernetes_filter"
        allow_list = ["allow_list"]
        deny_list = ["deny_list"]
      }
    }
  }

  #Optional
  description   = var.unified_agent_configuration_description
  display_name  = "test_unified_agent_configuration_monitoring_KUBERNETES"
  freeform_tags = var.unified_agent_configuration_freeform_tags
  group_association {
    #Optional
    group_list = ["ocid1.dynamicgroup.oc1..aaaaaaaatqbpurg4jtr57dthka4lbykvsqajjmynecixwgsfgu2z36wf4kgq"]
  }

  lifecycle {
    ignore_changes = [ defined_tags ]
  }
}

data "oci_logging_unified_agent_configurations" "test_unified_agent_configuration_monitoring_KUBERNETES" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name  = "test_unified_agent_configuration_monitoring_KUBERNETES"
#  group_id                     = var.test_log_group_id
#  is_compartment_id_in_subtree = var.unified_agent_configuration_is_compartment_id_in_subtree
#  log_id                       = var.test_log_id
  filter {
    name   = "id"
    values = [oci_logging_unified_agent_configuration.test_unified_agent_configuration_monitoring_KUBERNETES.id]
  }
#  state                        = var.unified_agent_configuration_state
}


resource "oci_logging_unified_agent_configuration" "test_unified_agent_configuration_monitoring_TAIL" {
  #Required
  compartment_id = var.compartment_id
  is_enabled     = var.unified_agent_configuration_is_enabled
  service_configuration {
    #Required
    configuration_type = "MONITORING"

    application_configurations {
      destination {
        compartment_id = var.compartment_id
        metrics_namespace = "tf_test_namespace"
      }
      source_type = "TAIL"
      sources {
        #Required
        source_type = "LOG_TAIL"
        name     = "test_unified_agent_configuration_monitoring_1_sources_name_0"
        parser {
          parser_type = "JSON"
          field_time_key = "time"
          is_keep_time_key = true
          timeout_in_milliseconds = 1000
        }
        paths = ["/var/log/*"]
      }
      sources {
        #Required
        source_type = "LOG_TAIL"
        name     = "test_unified_agent_configuration_monitoring_1_sources_name_1"
        parser {
          parser_type = "REGEXP"
          expression = "regexp"
        }
        paths = ["/var/log1/*"]
      }
    }
  }

  #Optional
  description   = var.unified_agent_configuration_description
  display_name  = "test_unified_agent_configuration_monitoring_TAIL"
  freeform_tags = var.unified_agent_configuration_freeform_tags
  group_association {
    #Optional
    group_list = ["ocid1.dynamicgroup.oc1..aaaaaaaatqbpurg4jtr57dthka4lbykvsqajjmynecixwgsfgu2z36wf4kgq"]
  }

  lifecycle {
    ignore_changes = [ defined_tags ]
  }
}

resource "oci_logging_unified_agent_configuration" "test_unified_agent_configuration_monitoring_URL" {
  #Required
  compartment_id = var.compartment_id
  is_enabled     = var.unified_agent_configuration_is_enabled
  service_configuration {
    #Required
    configuration_type = "MONITORING"

    application_configurations {
      destination {
        compartment_id = var.compartment_id
        metrics_namespace = "tf_test_namespace"
      }
      source_type = "URL"
      source {
        name = "url_source"
        scrape_targets {
          name = "url_scrape"
          url = "http://example.com"
        }
      }
      unified_agent_configuration_filter {
        filter_type = "URL_FILTER"
        name = "url_filter"
        allow_list = ["allow_list"]
        deny_list = ["deny_list"]
      }
    }
  }

  #Optional
  description   = var.unified_agent_configuration_description
  display_name  = "test_unified_agent_configuration_monitoring_URL"
  freeform_tags = var.unified_agent_configuration_freeform_tags
  group_association {
    #Optional
    group_list = ["ocid1.dynamicgroup.oc1..aaaaaaaatqbpurg4jtr57dthka4lbykvsqajjmynecixwgsfgu2z36wf4kgq"]
  }

  lifecycle {
    ignore_changes = [ defined_tags ]
  }
}