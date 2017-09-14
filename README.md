## NOTICE
**The terraform provider has been renamed, see [this wiki](https://github.com/oracle/terraform-provider-oci/wiki/Oracle-Terraform-Provider-Name-Change) for information on migration steps.**

*Legacy provider documentation (for v1.0.18 and earlier) can be found [here](https://github.com/oracle/terraform-provider-oci/tree/v1.0.18/docs)* 
 

    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
# Terraform provider for Oracle Cloud Infrastructure

[![wercker status](https://app.wercker.com/status/666d2ee10f45dde41189bb03248aadf9/s/master "wercker status")](https://app.wercker.com/project/byKey/666d2ee10f45dde41189bb03248aadf9)

Oracle customers now have access to an enterprise class, developer friendly orchestration tool they can use to manage [Oracle Cloud Infrastructure](https://cloud.oracle.com/en_US/bare-metal) resources as well as the [Oracle Compute Cloud](https://github.com/oracle/terraform-provider-compute).

This Terraform provider is OSS, available to all OCI customers at no charge.

## Compatibility
The provider is compatible with Terraform .9.\*.

### Coverage
The Terraform provider provides coverage for the entire OCI API, with some minor exceptions.

## Getting started
Be sure to read the FAQ and Writing Terraform configurations for OCI in [/docs](https://github.com/oracle/terraform-provider-oci/tree/master/docs).

### Download Terraform
Download the appropriate **.9.x binary** for your platform.  
https://www.terraform.io/downloads.html

**NOTE** Terraform v.10.x introduces a change to plugin management where 
previous v.9.x configuration no longer applies. See note below.


### Install Terraform
https://www.terraform.io/intro/getting-started/install.html

### Get the Oracle Cloud Infrastructure Terraform provider
https://github.com/oracle/terraform-provider-oci/releases

Unpack the provider. Terraform v.10.x introduces a change to plugin 
management where v.9.x configuration no longer applies. To be compatible 
with both terraform v.9.x and v.10.x, put the provider in the following 
location:

#### On \*nix
```
~/.terraform.d/plugins/
```

Then create the `~/.terraformrc` file that specifies the path to the 
`oci` provider **(only required for v.9.x)**.
```
providers {
  oci = "~/.terraform.d/plugins/terraform-provider-oci"
}
```

#### On Windows
```
%APPDATA%/terraform.d/plugins/
```

Then create `%APPDATA%/terraform.rc` that specifies the path to the 
`oci` provider **(only required for v.9.x)**.
```
providers {
  oci = "%appdata%/terraform.d/plugins/terraform-provider-oci"
}
```

### Export credentials
Required Keys and OCIDs - https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm

If you primarily work in a single compartment consider exporting that compartment's OCID as well. Remember that the tenancy OCID is also the OCID of the root compartment.

#### \*nix
If your TF configurations are limited to a single compartment/user then 
using this `bash_profile` option will work well. For more complex 
environments you may want to maintain multiple sets of environment 
variables. 
See the [compute single instance example](https://github.com/oracle/terraform-provider-oci/tree/master/docs/examples/compute/instance) for more info.

In your ~/.bash_profile set these variables
```
export TF_VAR_tenancy_ocid=
export TF_VAR_user_ocid=
export TF_VAR_fingerprint=
export TF_VAR_private_key_path=<fully qualified path>`
```

Once you've set these values open a new terminal or source your profile changes
```
$ source ~/.bash_profile
```

#### Windows
```
setx TF_VAR_tenancy_ocid <value>
setx TF_VAR_user_ocid <value>
setx TF_VAR_fingerprint <value>
setx TF_VAR_private_key_path <value>
```
The variables won't be set for the current session, exit the terminal and reopen.

## Deploy an example configuration
Download the [compute single instance example](https://github.com/oracle/terraform-provider-oci/tree/master/docs/examples/compute/instance).

Edit it to include the OCID of the compartment you want to create the VCN. Remember that the tenancy OCID is the compartment OCID of your root compartment.

You should always plan, then apply a configuration -
```
# from the compute/instance directory
$ terraform plan
  
# Make sure the plan looks right.
$ terraform apply
```
## OCI resource and datasource details
https://github.com/oracle/terraform-provider-oci/tree/master/docs

## Getting help
You can file an issue against the project
https://github.com/oracle/terraform-provider-oci/issues

or meet us in the OCI forums
https://community.oracle.com/community/cloud_computing/bare-metal

## Known issues

[Github issues](https://github.com/oracle/terraform-provider-oci/issues)

## About the provider
This provider was written on behalf of Oracle by [MustWin.](http://mustwin.com/)
