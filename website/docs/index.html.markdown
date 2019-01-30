---
layout: "oci"
page_title: "Provider: Oracle Cloud Infrastructure"
sidebar_current: "docs-oci-index"
description: |-
  The Oracle Cloud Infrastructure provider is used to interact with the many resources supported by the Oracle Cloud Infrastructure services. The provider needs to be configured with credentials for the Oracle Cloud Account.
---

# Oracle Cloud Infrastructure Provider

The Oracle Cloud Infrastructure provider is used to interact with the many resources supported by the [Oracle Cloud Infrastructure](https://cloud.oracle.com/cloud-infrastructure). The provider needs to be configured with credentials for the Oracle Cloud Infrastructure account.  

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Oracle Cloud Infrastructure provider with an API Key
provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

# Get a list of Availability Domains
data "oci_identity_availability_domains" "ads" {
  compartment_id = "${var.tenancy_ocid}"
}

# Output the result
output "show-ads" {
  value = "${data.oci_identity_availability_domains.ads.availability_domains}"
}

```

## Authentication

The OCI provider supports API Key based authentication and Instance Principal based authentication.

### API Key based authentication  
Calls to OCI using API Key authentication requires that you provide the following credentials:

- `tenancy_ocid` - The global identifier for your account, always shown on the bottom of the web console. 
- `user_ocid` - The identifier of the user account you will be using for Terraform. For information on setting the 
correct policies for your user see [Managing Users](https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingusers.htm).
- `private_key_path` - The path to the private key stored on your computer. The public key portion must be added to the 
user account above in the _API Keys_ section of the web console. For details on how to create and configure keys see 
[Required Keys and OCIDs](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm).
- `fingerprint` - The fingerprint of the public key added in the above user's _API Keys_ section of the web console.
- `region` - The region to target with this provider configuration.

#### Environment variables
It is common to export the above values as environment variables, or source them in different bash profiles when executing 
Terraform commands. Below are OS specific examples for configuring these environment values.

If you primarily work in a single compartment, consider exporting the compartment OCID as well. The tenancy OCID is also 
the OCID of the root compartment, and can be used where any compartment id is required.

##### \*nix
If your Terraform configurations are limited to a single compartment or user, then using this `bash_profile` option be 
sufficient. For more complex environments you may want to maintain multiple sets of environment variables. 
See the [compute single instance example](https://github.com/oracle/terraform-provider-oci/tree/master/docs/examples/compute/instance) for more info.

In your `~/.bash_profile` set these variables:
```
export TF_VAR_tenancy_ocid=<value>
export TF_VAR_compartment_ocid=<value>
export TF_VAR_user_ocid=<value>
export TF_VAR_fingerprint=<value>
export TF_VAR_private_key_path=<value>
``` 

Once you've set these values open a new terminal or source your profile changes:
```
$ source ~/.bash_profile
```

##### Windows

Configuring for Windows usage is largely the same, with one notable exception: you can use PuttyGen to create the public 
and private key as shown above, however, you will need to convert them from PPK format to PEM format.

```
setx TF_VAR_tenancy_ocid <value>
setx TF_VAR_compartment_ocid <value>
setx TF_VAR_user_ocid <value>
setx TF_VAR_fingerprint <value>
setx TF_VAR_private_key_path <value>
```
The variables won't be set for the current session, exit the terminal and reopen.


### Instance Principal Authentication
Instance Principal authentication allows you to run Terraform from an OCI Instance within your Tenancy. To enable Instance 
Principal authentication, set the `auth` attribute to "InstancePrincipal" in the provider definition as below:

```
# Configure the Oracle Cloud Infrastructure provider to use Instance Principal based authentication
provider "oci" {
  auth = "InstancePrincipal"
  region = "${var.region}"
}
```

_Note: this configuration will only work when run from an OCI instance. For more information on using Instance 
Principals, see [this document](https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm)._

## Testing
Credentials must be provided via the environment variables as shown above in order to run acceptance tests.

## Configuring Automatic Retries
While applying, refreshing, or destroying a plan, Terraform may encounter some intermittent OCI errors (such as 429 or 500 errors) that could succeed on retry. 
By default, the Terraform OCI provider will automatically retry such operations for up to 10 minutes. 
The following fields can be specified in the provider block to further configure the retry behavior:

- `disable_auto_retries` - Disable automatic retries for retriable errors.
- `retry_duration_seconds` - The minimum duration (in seconds) to retry a resource operation in response to HTTP 429 and HTTP 500 errors. The actual retry duration may be slightly longer due to jittering of retry operations. This value is ignored if the `disable_auto_retries` field is set to true.

### Concurrency Control using Retry Backoff and Jitter
To alleviate contention between parallel operations against OCI services; the Terraform OCI provider schedules retry attempts using quadratic backoff and full jitter.
Quadratic backoff increases the maximum interval between subsequent retry attempts, while full jitter randomly selects a retry interval within the backoff range.

For example, the wait time between the 1st and 2nd retry attempts is chosen randomly between 1 and 8 seconds. The wait time between the 2nd and 3rd retry attempts is chosen
randomly between 1 and 18 seconds. Regardless of the number of retry attempts, the retry interval time is capped after the 12th attempt at 288 seconds.

Note that the `retry_duration_seconds` field only affects retry duration in response to HTTP 429 and 500 errors; as these errors are more likely to result in success after a long retry duration.
Other HTTP errors (such as 400, 401, 403, 404, and 409) are unlikely to succeed on retry. The `retry_duration_seconds` field does not affect the retry behavior for such errors.
