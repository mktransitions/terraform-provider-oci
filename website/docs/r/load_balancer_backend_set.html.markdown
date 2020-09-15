---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_backend_set"
sidebar_current: "docs-oci-resource-load_balancer-backend_set"
description: |-
  Provides the Backend Set resource in Oracle Cloud Infrastructure Load Balancer service
---

# oci_load_balancer_backend_set
This resource provides the Backend Set resource in Oracle Cloud Infrastructure Load Balancer service.

Adds a backend set to a load balancer.

## Supported Aliases

* `oci_load_balancer_backendset`

## Example Usage

```hcl
resource "oci_load_balancer_backend_set" "test_backend_set" {
	#Required
	health_checker {
		#Required
		protocol = var.backend_set_health_checker_protocol

		#Optional
		interval_ms = var.backend_set_health_checker_interval_ms
		port = var.backend_set_health_checker_port
		response_body_regex = var.backend_set_health_checker_response_body_regex
		retries = var.backend_set_health_checker_retries
		return_code = var.backend_set_health_checker_return_code
		timeout_in_millis = var.backend_set_health_checker_timeout_in_millis
		url_path = var.backend_set_health_checker_url_path
	}
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	name = var.backend_set_name
	policy = var.backend_set_policy

	#Optional
	lb_cookie_session_persistence_configuration {

		#Optional
		cookie_name = var.backend_set_lb_cookie_session_persistence_configuration_cookie_name
		disable_fallback = var.backend_set_lb_cookie_session_persistence_configuration_disable_fallback
		domain = var.backend_set_lb_cookie_session_persistence_configuration_domain
		is_http_only = var.backend_set_lb_cookie_session_persistence_configuration_is_http_only
		is_secure = var.backend_set_lb_cookie_session_persistence_configuration_is_secure
		max_age_in_seconds = var.backend_set_lb_cookie_session_persistence_configuration_max_age_in_seconds
		path = var.backend_set_lb_cookie_session_persistence_configuration_path
	}
	session_persistence_configuration {
		#Required
		cookie_name = var.backend_set_session_persistence_configuration_cookie_name

		#Optional
		disable_fallback = var.backend_set_session_persistence_configuration_disable_fallback
	}
	ssl_configuration {
		#Required
		certificate_name = oci_load_balancer_certificate.test_certificate.certificate_name

		#Optional
		verify_depth = var.backend_set_ssl_configuration_verify_depth
		verify_peer_certificate = var.backend_set_ssl_configuration_verify_peer_certificate
		protocols = ["TLSv1.1", "TLSv1.2"]
		cipher_suite_name = oci_load_balancer_ssl_cipher_suite.example_ssl_cipher_suite.name
		server_order_preference = ENABLED
	}
}
```
**Note:** The `sessionPersistenceConfiguration` (application cookie stickiness) and `lbCookieSessionPersistenceConfiguration`
      (LB cookie stickiness) attributes are mutually exclusive. To avoid returning an error, configure only one of these two
      attributes per backend set.

## Argument Reference

The following arguments are supported:

* `health_checker` - (Required) (Updatable) 
	* `interval_ms` - (Optional) (Updatable) The interval between health checks, in milliseconds.  Example: `10000` 
	* `port` - (Optional) (Updatable) The backend server port against which to run the health check. If the port is not specified, the load balancer uses the port information from the `Backend` object.  Example: `8080` 
	* `protocol` - (Required) (Updatable) The protocol the health check must use; either HTTP or TCP.  Example: `HTTP` 
	* `response_body_regex` - (Optional) (Updatable) A regular expression for parsing the response body from the backend server.  Example: `^((?!false).|\s)*$` 
	* `retries` - (Optional) (Updatable) The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state.  Example: `3` 
	* `return_code` - (Optional) (Updatable) The status code a healthy backend server should return.  Example: `200` 
	* `timeout_in_millis` - (Optional) (Updatable) The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period.  Example: `3000` 
	* `url_path` - (Optional) (Updatable) The path against which to run the health check.  Example: `/healthcheck` 
* `lb_cookie_session_persistence_configuration` - (Optional) (Updatable) 
	* `cookie_name` - (Optional) (Updatable) The name of the cookie inserted by the load balancer. If this field is not configured, the cookie name defaults to "X-Oracle-BMC-LBS-Route".  Example: `example_cookie`

		**Notes:**
		*  Ensure that the cookie name used at the backend application servers is different from the cookie name used at the load balancer. To minimize the chance of name collision, Oracle recommends that you use a prefix such as "X-Oracle-OCI-" for this field.
		*  If a backend server and the load balancer both insert cookies with the same name, the client or browser behavior can vary depending on the domain and path values associated with the cookie. If the name, domain, and path values of the `Set-cookie` generated by a backend server and the `Set-cookie` generated by the load balancer are all the same, the client or browser treats them as one cookie and returns only one of the cookie values in subsequent requests. If both `Set-cookie` names are the same, but the domain and path names are different, the client or browser treats them as two different cookies. 
	* `disable_fallback` - (Optional) (Updatable) Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: `false` 
	* `domain` - (Optional) (Updatable) The domain in which the cookie is valid. The `Set-cookie` header inserted by the load balancer contains a domain attribute with the specified value.

		This attribute has no default value. If you do not specify a value, the load balancer does not insert the domain attribute into the `Set-cookie` header.

		**Notes:**
		*  [RFC 6265 - HTTP State Management Mechanism](https://www.ietf.org/rfc/rfc6265.txt) describes client and browser behavior when the domain attribute is present or not present in the `Set-cookie` header.

		If the value of the `Domain` attribute is `example.com` in the `Set-cookie` header, the client includes the same cookie in the `Cookie` header when making HTTP requests to `example.com`, `www.example.com`, and `www.abc.example.com`. If the `Domain` attribute is not present, the client returns the cookie only for the domain to which the original request was made.
		*  Ensure that this attribute specifies the correct domain value. If the `Domain` attribute in the `Set-cookie` header does not include the domain to which the original request was made, the client or browser might reject the cookie. As specified in RFC 6265, the client accepts a cookie with the `Domain` attribute value `example.com` or `www.example.com` sent from `www.example.com`. It does not accept a cookie with the `Domain` attribute `abc.example.com` or `www.abc.example.com` sent from `www.example.com`.

		Example: `example.com` 
	* `is_http_only` - (Optional) (Updatable) Whether the `Set-cookie` header should contain the `HttpOnly` attribute. If `true`, the `Set-cookie` header inserted by the load balancer contains the `HttpOnly` attribute, which limits the scope of the cookie to HTTP requests. This attribute directs the client or browser to omit the cookie when providing access to cookies through non-HTTP APIs. For example, it restricts the cookie from JavaScript channels.  Example: `true` 
	* `is_secure` - (Optional) (Updatable) Whether the `Set-cookie` header should contain the `Secure` attribute. If `true`, the `Set-cookie` header inserted by the load balancer contains the `Secure` attribute, which directs the client or browser to send the cookie only using a secure protocol.

		**Note:** If you set this field to `true`, you cannot associate the corresponding backend set with an HTTP listener.

		Example: `true` 
	* `max_age_in_seconds` - (Optional) (Updatable) The amount of time the cookie remains valid. The `Set-cookie` header inserted by the load balancer contains a `Max-Age` attribute with the specified value.

		The specified value must be at least one second. There is no default value for this attribute. If you do not specify a value, the load balancer does not include the `Max-Age` attribute in the `Set-cookie` header. In most cases, the client or browser retains the cookie until the current session ends, as defined by the client.

		Example: `3600` 
	* `path` - (Optional) (Updatable) The path in which the cookie is valid. The `Set-cookie header` inserted by the load balancer contains a `Path` attribute with the specified value.

		Clients include the cookie in an HTTP request only if the path portion of the request-uri matches, or is a subdirectory of, the cookie's `Path` attribute.

		The default value is `/`.

		Example: `/example` 
* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer on which to add a backend set.
* `name` - (Required) A friendly name for the backend set. It must be unique and it cannot be changed.

	Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot contain spaces. Avoid entering confidential information.

	Example: `example_backend_set` 
* `policy` - (Required) (Updatable) The load balancer policy for the backend set. To get a list of available policies, use the [ListPolicies](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancerPolicy/ListPolicies) operation.  Example: `LEAST_CONNECTIONS` 
* `session_persistence_configuration` - (Optional) (Updatable) 
	* `cookie_name` - (Required) (Updatable) The name of the cookie used to detect a session initiated by the backend server. Use '*' to specify that any cookie set by the backend causes the session to persist.  Example: `example_cookie` 
	* `disable_fallback` - (Optional) (Updatable) Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: `false` 
* `ssl_configuration` - (Optional) (Updatable) 
	* `certificate_name` - (Required) (Updatable) A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
	* `verify_depth` - (Optional) (Updatable) The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - (Optional) (Updatable) Whether the load balancer listener should verify peer certificates.  Example: `true` 
	* `protocols` - (Optional) (Updatable) A list of SSL protocols the load balancer must support for HTTPS or SSL connections. The load balancer uses SSL protocols to establish a secure connection between a client and a server. A secure connection ensures that all data passed between the client and the server is private. The Load Balancing service supports the following protocols:  TLSv1  TLSv1.1  TLSv1.2  If this field is not specified, TLSv1.2 is the default.  Example: `["TLSv1.1", "TLSv1.2"]`
    * `cipher_suite_name` - (Optional) (Updatable) The name of the cipher suite to use for HTTPS or SSL connections. If this field is not specified, the default is `oci-default-ssl-cipher-suite-v1`. Example: `example_cipher_suite`
    * `server_order_preference` - (Optional) (Updatable) When this attribute is set to ENABLED, the system gives preference to the server ciphers over the client ciphers.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `backend` - 
	* `backup` - Whether the load balancer should treat this server as a backup unit. If `true`, the load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.

		**Note:** You cannot add a backend server marked as `backup` to a backend set that uses the IP Hash policy.

		Example: `false` 
	* `drain` - Whether the load balancer should drain this server. Servers marked "drain" receive no new incoming traffic.  Example: `false` 
	* `ip_address` - The IP address of the backend server.  Example: `10.0.0.3` 
	* `name` - A read-only field showing the IP address and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080` 
	* `offline` - Whether the load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
	* `port` - The communication port for the backend server.  Example: `8080` 
	* `weight` - The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections as a server weighted '1'. For more information on load balancing policies, see [How Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 
* `health_checker` - 
	* `interval_ms` - The interval between health checks, in milliseconds. The default is 30000 (30 seconds).  Example: `30000` 
	* `port` - The backend server port against which to run the health check. If the port is not specified, the load balancer uses the port information from the `Backend` object.  Example: `8080` 
	* `protocol` - The protocol the health check must use; either HTTP or TCP.  Example: `HTTP` 
	* `response_body_regex` - A regular expression for parsing the response body from the backend server.  Example: `^((?!false).|\s)*$` 
	* `retries` - The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state. Defaults to 3.  Example: `3` 
	* `return_code` - The status code a healthy backend server should return. If you configure the health check policy to use the HTTP protocol, you can use common HTTP status codes such as "200".  Example: `200` 
	* `timeout_in_millis` - The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period. Defaults to 3000 (3 seconds).  Example: `3000` 
	* `url_path` - The path against which to run the health check.  Example: `/healthcheck` 
* `lb_cookie_session_persistence_configuration` - 
	* `cookie_name` - The name of the cookie inserted by the load balancer. If this field is not configured, the cookie name defaults to "X-Oracle-BMC-LBS-Route".  Example: `example_cookie`

		**Notes:**
		*  Ensure that the cookie name used at the backend application servers is different from the cookie name used at the load balancer. To minimize the chance of name collision, Oracle recommends that you use a prefix such as "X-Oracle-OCI-" for this field.
		*  If a backend server and the load balancer both insert cookies with the same name, the client or browser behavior can vary depending on the domain and path values associated with the cookie. If the name, domain, and path values of the `Set-cookie` generated by a backend server and the `Set-cookie` generated by the load balancer are all the same, the client or browser treats them as one cookie and returns only one of the cookie values in subsequent requests. If both `Set-cookie` names are the same, but the domain and path names are different, the client or browser treats them as two different cookies. 
	* `disable_fallback` - Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: `false` 
	* `domain` - The domain in which the cookie is valid. The `Set-cookie` header inserted by the load balancer contains a domain attribute with the specified value.

		This attribute has no default value. If you do not specify a value, the load balancer does not insert the domain attribute into the `Set-cookie` header.

		**Notes:**
		*  [RFC 6265 - HTTP State Management Mechanism](https://www.ietf.org/rfc/rfc6265.txt) describes client and browser behavior when the domain attribute is present or not present in the `Set-cookie` header.

		If the value of the `Domain` attribute is `example.com` in the `Set-cookie` header, the client includes the same cookie in the `Cookie` header when making HTTP requests to `example.com`, `www.example.com`, and `www.abc.example.com`. If the `Domain` attribute is not present, the client returns the cookie only for the domain to which the original request was made.
		*  Ensure that this attribute specifies the correct domain value. If the `Domain` attribute in the `Set-cookie` header does not include the domain to which the original request was made, the client or browser might reject the cookie. As specified in RFC 6265, the client accepts a cookie with the `Domain` attribute value `example.com` or `www.example.com` sent from `www.example.com`. It does not accept a cookie with the `Domain` attribute `abc.example.com` or `www.abc.example.com` sent from `www.example.com`.

		Example: `example.com` 
	* `is_http_only` - Whether the `Set-cookie` header should contain the `HttpOnly` attribute. If `true`, the `Set-cookie` header inserted by the load balancer contains the `HttpOnly` attribute, which limits the scope of the cookie to HTTP requests. This attribute directs the client or browser to omit the cookie when providing access to cookies through non-HTTP APIs. For example, it restricts the cookie from JavaScript channels.  Example: `true` 
	* `is_secure` - Whether the `Set-cookie` header should contain the `Secure` attribute. If `true`, the `Set-cookie` header inserted by the load balancer contains the `Secure` attribute, which directs the client or browser to send the cookie only using a secure protocol.

		**Note:** If you set this field to `true`, you cannot associate the corresponding backend set with an HTTP listener.

		Example: `true` 
	* `max_age_in_seconds` - The amount of time the cookie remains valid. The `Set-cookie` header inserted by the load balancer contains a `Max-Age` attribute with the specified value.

		The specified value must be at least one second. There is no default value for this attribute. If you do not specify a value, the load balancer does not include the `Max-Age` attribute in the `Set-cookie` header. In most cases, the client or browser retains the cookie until the current session ends, as defined by the client.

		Example: `3600` 
	* `path` - The path in which the cookie is valid. The `Set-cookie header` inserted by the load balancer contains a `Path` attribute with the specified value.

		Clients include the cookie in an HTTP request only if the path portion of the request-uri matches, or is a subdirectory of, the cookie's `Path` attribute.

		The default value is `/`.

		Example: `/example` 
* `name` - A friendly name for the backend set. It must be unique and it cannot be changed.

	Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot contain spaces. Avoid entering confidential information.

	Example: `example_backend_set` 
* `policy` - The load balancer policy for the backend set. To get a list of available policies, use the [ListPolicies](https://docs.cloud.oracle.com/iaas/api/#/en/loadbalancer/20170115/LoadBalancerPolicy/ListPolicies) operation.  Example: `LEAST_CONNECTIONS` 
* `session_persistence_configuration` - 
	* `cookie_name` - The name of the cookie used to detect a session initiated by the backend server. Use '*' to specify that any cookie set by the backend causes the session to persist.  Example: `example_cookie` 
	* `disable_fallback` - Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: `false` 
* `ssl_configuration` - 
	* `certificate_name` - A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `example_certificate_bundle` 
	* `verify_depth` - The maximum depth for peer certificate chain verification.  Example: `3` 
	* `verify_peer_certificate` - Whether the load balancer listener should verify peer certificates. Defaults to true.   Example: `true` 

## Import

BackendSets can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_backend_set.test_backend_set "loadBalancers/{loadBalancerId}/backendSets/{backendSetName}" 
```

