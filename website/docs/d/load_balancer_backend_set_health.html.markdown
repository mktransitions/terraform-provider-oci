---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_backend_set_health"
sidebar_current: "docs-oci-datasource-load_balancer-backend_set_health"
description: |-
  Provides details about a specific Backend Set Health in Oracle Cloud Infrastructure Load Balancer service
---

# Data Source: oci_load_balancer_backend_set_health
This data source provides details about a specific Backend Set Health resource in Oracle Cloud Infrastructure Load Balancer service.

Gets the health status for the specified backend set.

## Example Usage

```hcl
data "oci_load_balancer_backend_set_health" "test_backend_set_health" {
	#Required
	backend_set_name = "${oci_load_balancer_backend_set.test_backend_set.name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
```

## Argument Reference

The following arguments are supported:

* `backend_set_name` - (Required) The name of the backend set to retrieve the health status for.  Example: `example_backend_set` 
* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend set health status to be retrieved.


## Attributes Reference

The following attributes are exported:

* `critical_state_backend_names` - A list of backend servers that are currently in the `CRITICAL` health state. The list identifies each backend server by IP address and port.  Example: `10.0.0.4:8080` 
* `status` - Overall health status of the backend set.
	*  **OK:** All backend servers in the backend set return a status of `OK`.
	*  **WARNING:** Half or more of the backend set's backend servers return a status of `OK` and at least one backend server returns a status of `WARNING`, `CRITICAL`, or `UNKNOWN`.
	*  **CRITICAL:** Fewer than half of the backend set's backend servers return a status of `OK`.
	*  **UNKNOWN:** More than half of the backend set's backend servers return a status of `UNKNOWN`, the system was unable to retrieve metrics, or the backend set does not have a listener attached. 
* `total_backend_count` - The total number of backend servers in this backend set.  Example: `7` 
* `unknown_state_backend_names` - A list of backend servers that are currently in the `UNKNOWN` health state. The list identifies each backend server by IP address and port.  Example: `10.0.0.5:8080` 
* `warning_state_backend_names` - A list of backend servers that are currently in the `WARNING` health state. The list identifies each backend server by IP address and port.  Example: `10.0.0.3:8080` 

