---
subcategory: "Dataflow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_invoke_runs"
sidebar_current: "docs-oci-datasource-dataflow-invoke_runs"
description: |-
  Provides the list of Invoke Runs in Oracle Cloud Infrastructure Dataflow service
---

# Data Source: oci_dataflow_invoke_runs
This data source provides the list of Invoke Runs in Oracle Cloud Infrastructure Dataflow service.

Lists all runs of an application in the specified compartment.


## Example Usage

```hcl
data "oci_dataflow_invoke_runs" "test_invoke_runs" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	application_id = oci_dataflow_application.test_application.id
	display_name = var.invoke_run_display_name
	display_name_starts_with = var.invoke_run_display_name_starts_with
	owner_principal_id = oci_dataflow_owner_principal.test_owner_principal.id
	state = var.invoke_run_state
	time_created_greater_than = var.invoke_run_time_created_greater_than
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The ID of the application. 
* `compartment_id` - (Required) The OCID of the compartment. 
* `display_name` - (Optional) The query parameter for the Spark application name. 
* `display_name_starts_with` - (Optional) The displayName prefix. 
* `owner_principal_id` - (Optional) The OCID of the user who created the resource. 
* `state` - (Optional) The LifecycleState of the run. 
* `time_created_greater_than` - (Optional) The epoch time that the resource was created. 


## Attributes Reference

The following attributes are exported:

* `runs` - The list of runs.

### InvokeRun Reference

The following attributes are exported:

* `application_id` - The application ID. 
* `archive_uri` - An Oracle Cloud Infrastructure URI of an archive.zip file containing custom dependencies that may be used to support the execution a Python, Java, or Scala application. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `arguments` - The arguments passed to the running application as command line arguments.  An argument is either a plain text or a placeholder. Placeholders are replaced using values from the parameters map.  Each placeholder specified must be represented in the parameters map else the request (POST or PUT) will fail with a HTTP 400 status code.  Placeholders are specified as `Service Api Spec`, where `name` is the name of the parameter. Example:  `[ "--input", "${input_file}", "--name", "John Doe" ]` If "input_file" has a value of "mydata.xml", then the value above will be translated to `--input mydata.xml --name "John Doe"` 
* `class_name` - The class for the application. 
* `compartment_id` - The OCID of a compartment. 
* `configuration` - The Spark configuration passed to the running process. See https://spark.apache.org/docs/latest/configuration.html#available-properties. Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" } Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is not allowed to be overwritten will cause a 400 status to be returned. 
* `data_read_in_bytes` - The data read by the run in bytes. 
* `data_written_in_bytes` - The data written by the run in bytes. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. This name is not necessarily unique. 
* `driver_shape` - The VM shape for the driver. Sets the driver cores and memory. 
* `executor_shape` - The VM shape for the executors. Sets the executor cores and memory. 
* `file_uri` - An Oracle Cloud Infrastructure URI of the file containing the application to execute. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The ID of a run. 
* `language` - The Spark language. 
* `lifecycle_details` - The detailed messages about the lifecycle state. 
* `logs_bucket_uri` - An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 
* `num_executors` - The number of executor VMs requested. 
* `opc_request_id` - Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID. 
* `owner_principal_id` - The OCID of the user who created the resource. 
* `owner_user_name` - The username of the user who created the resource.  If the username of the owner does not exist, `null` will be returned and the caller should refer to the ownerPrincipalId value instead. 
* `parameters` - An array of name/value pairs used to fill placeholders found in properties like `Application.arguments`.  The name must be a string of one or more word characters (a-z, A-Z, 0-9, _).  The value can be a string of 0 or more characters of any kind. Example:  [ { name: "iterations", value: "10"}, { name: "input_file", value: "mydata.xml" }, { name: "variable_x", value: "${x}"} ] 
	* `name` - The name of the parameter.  It must be a string of one or more word characters (a-z, A-Z, 0-9, _). Examples: "iterations", "input_file" 
	* `value` - The value of the parameter. It must be a string of 0 or more characters of any kind. Examples: "" (empty string), "10", "mydata.xml", "${x}" 
* `private_endpoint_dns_zones` - An array of DNS zone names. Example: `[ "app.examplecorp.com", "app.examplecorp2.com" ]` 
* `private_endpoint_id` - The OCID of a private endpoint. 
* `private_endpoint_max_host_count` - The maximum number of hosts to be accessed through the private endpoint. This value is used to calculate the relevant CIDR block and should be a multiple of 256.  If the value is not a multiple of 256, it is rounded up to the next multiple of 256. For example, 300 is rounded up to 512. 
* `private_endpoint_nsg_ids` - An array of network security group OCIDs. 
* `private_endpoint_subnet_id` - The OCID of a subnet. 
* `run_duration_in_milliseconds` - The duration of the run in milliseconds. 
* `spark_version` - The Spark version utilized to run the application. 
* `state` - The current state of this run. 
* `time_created` - The date and time a application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `time_updated` - The date and time a application was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `total_ocpu` - The total number of oCPU requested by the run. 
* `warehouse_bucket_uri` - An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory for BATCH SQL runs. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat. 

