---
layout: "oci"
page_title: "Resource Discovery"
sidebar_current: "docs-oci-guide-resource_discovery"
description: |-
  The Oracle Cloud Infrastructure provider. Discovering resources in an existing compartment
---

## Resource Discovery

### Overview

You can use Terraform Resource Discovery to discover deployed resources in your compartment and export them to Terraform configuration and state files. This release supports the most commonly used Oracle Cloud Infrastructure services, such as Compute, Block Volumes, Networking, Load Balancing, Database, and Identity and Access Management (IAM). Please look at the section “Supported Resources” for details.

### Use Cases and Benefits

With this feature, you can perform the following tasks:

* **Move from manually-managed infrastructure to Terraform-managed infrastructure:** You can generate a baseline Terraform state file for your existing infrastructure with a single command, and manage this infrastructure by using Terraform.

* **Detect state drift:** By managing the infrastructure using Terraform, you can detect when the state of your resources changes and differs from the desired configuration.

* **Duplicate or rebuild existing infrastructure:** By creating Terraform configuration files, you can re-create your existing infrastructure architecture in a new tenancy or compartment.

* **Get started with Terraform:** If you’re new to Terraform, you can learn about Terraform’s HCL syntax and how to represent Oracle Cloud Infrastructure resources in HCL.

Please note that this feature is available for version 3.50 and above. The latest version of the terraform-oci-provider can be downloaded using terraform init or by going to https://releases.hashicorp.com/terraform-provider-oci/

### Authentication
To discover resources in your compartment, the terraform-oci-provider will need authentication information about the user, tenancy, and region with which to discover
the resources. It is recommended to specify a user that has access to inspect and read the resources to discover.

Resource discovery supports API Key based authentication and Instance Principal based authentication.

The authentication information can be specified using the following environment variables:

```
export TF_VAR_tenancy_ocid=<value>
export TF_VAR_user_ocid=<value>
export TF_VAR_fingerprint=<value>
export TF_VAR_private_key_path=<path to your private key>
export TF_VAR_region=<region of the resources, e.g. "us-phoenix-1">
```


If your private key is password-encrypted, you may also need to specify a password with this variable:

```
export TF_VAR_private_key_password=<password for private key>
```

The authentication information can also be specified using a configuration file. For details on setting this up, see [SDK and CLI configuration file](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/sdkconfig.htm)
A non-default profile can be set using environment variable:

```
export TF_VAR_config_file_profile=<value>
```


If the parameters have multiple sources, the priority will be in the following order:

    Environment variables
    Non-default profile
    DEFAULT profile


### Usage

Once you have specified the prerequisite authentication settings, the command can be used as follows with a compartment being specified by name or OCID:

```
terraform-provider-oci -command=export -compartment_name=<name of compartment to export> -output_path=<directory under which to generate Terraform files>
```


```
terraform-provider-oci -command=export -compartment_id=<OCID of compartment to export> -output_path=<directory under which to generate Terraform files>
```

This command will discover resources within your compartment and generates Terraform configuration files in the given `output_path`.
The generated `.tf` files contain the Terraform configuration with the resources that the command has discovered.

**Parameter Description**

* `command` - Command to run. Supported commands include:
    * `export` - Discovers Oracle Cloud Infrastructure resources within your compartment and generates Terraform configuration files for them
    * `list_export_resources` - Lists the Terraform Oracle Cloud Infrastructure resources types that can be discovered by the `export` command
    * `list_export_services` - Lists the allowed values for services arguments along with scope in json format
* `compartment_id` - OCID of a compartment to export. If `compartment_id`  or `compartment_name` is not specified, the root compartment will be used.
* `compartment_name` - The name of a compartment to export. Use this instead of `compartment_id` to provide a compartment name.
* `ids` - Comma-separated list of resource IDs to export. The ID could either be an OCID or a Terraform import ID. By default, all resources are exported.
* `list_export_services_path` - Path to output list of supported services in json format, must include json file name
* `output_path` - Path to output generated configurations and state files of the exported compartment
* `services` - Comma-separated list of service resources to export. If not specified, all resources within the given compartment (which excludes identity resources) are exported. The following values can be specified:
    * `analytics` - Discovers analytics resources within the specified compartment
    * `apigateway` - Discovers apigateway resources within the specified compartment
    * `auto_scaling` - Discovers auto_scaling resources within the specified compartment
    * `availability_domain` - Discovers availability domains used by your compartment-level resources. It is recommended to always specify this value
    * `bds` - Discovers big data service resources within the specified compartment
    * `blockchain` - Discovers blockchain resources within the specified compartment
    * `budget` - Discovers budget resources across the entire tenancy
    * `cloud_guard` - Discovers cloud guard resources within the specified compartment
    * `containerengine` - Discovers containerengine resources within the specified compartment
    * `core` - Discovers compute, block storage, and networking resources within the specified compartment
    * `data_safe` - Discovers data_safe resources within the specified compartment
    * `database` - Discovers database resources within the specified compartment
    * `datacatalog` - Discovers datacatalog resources within the specified compartment
    * `dataflow` - Discovers dataflow resources within the specified compartment
    * `dataintegration` - Discovers dataintegration resources within the specified compartment
    * `datascience` - Discovers datascience resources within the specified compartment
    * `dns` - Discovers dns resources (except record) within the specified compartment
    * `email` - Discovers email_sender resources within the specified compartment
    * `email_tenancy` - Discovers email_suppression resources across the entire tenancy
    * `events` - Discovers events resources within the specified compartment
    * `file_storage` - Discovers file_storage resources within the specified compartment
    * `functions` - Discovers functions resources within the specified compartment
    * `health_checks` - Discovers health_checks resources within the specified compartment
    * `identity` - Discovers identity resources across the entire tenancy
    * `integration` - Discovers integration resources within the specified compartment
    * `kms` - Discovers kms resources within the specified compartment
    * `limits` - Discovers limits resources across the entire tenancy
    * `load_balancer` - Discovers load balancer resources within the specified compartment
    * `marketplace` - Discovers marketplace resources within the specified compartment
    * `monitoring` - Discovers monitoring resources within the specified compartment
    * `mysql` - Discovers mysql resources within the specified compartment
    * `nosql` - Discovers nosql resources within the specified compartment
    * `object_storage` - Discovers object storage resources within the specified compartment
    * `oce` - Discovers oce resources within the specified compartment
    * `ocvp` - Discovers ocvp resources within the specified compartment
    * `oda` - Discovers oda resources within the specified compartment
    * `ons` - Discovers ons resources within the specified compartment
    * `osmanagement` - Discovers osmanagement resources within the specified compartment
    * `sch` - Discovers sch resources within the specified compartment
    * `streaming` - Discovers streaming resources within the specified compartment
    * `tagging` - Discovers tag-related resources within the specified compartment
    * `waas` - Discovers waas resources within the specified compartment
* `generate_state` - Provide this flag to import the discovered resources into a state file along with the Terraform configuration
* `tf_version` - The version of terraform syntax to generate for configurations. Default is v0.12. The state file will be written in v0.12 only. The allowed values are:
    * 0.11
    * 0.12

| Arguments | Resources discovered |
| ----------| -------------------- |
| compartment_id = \<empty or tenancy ocid\>  <br> services= \<empty\> or not specified | all tenancy and compartment scope resources <br>  |
| compartment_id = \<empty or tenancy ocid\>  <br> services= \<comma separated list of services\> | tenancy and compartment scope resources for the services specified |
| compartment_id = \<non-root compartment\> <br> services= \<empty\> or not specified | all compartment scope resources only |
| compartment_id = \<non-root compartment\> <br> services=\<comma separated list of services\> | compartment scope resources for the services specified<br>tenancy scope resources will not be discovered even if services with such resources are specified |

> **Notes**:
* The compartment export functionality currently supports discovery of the target compartment. The ability to discover resources in child compartments is not yet supported.
* If using Instance Principals, resources can not be discovered if compartment_id is not specified

### Exit status

While discovering resources if there is any error related to the APIs or service unavailability, the tool will move on to find next resource. All the errors encountered will be displayed after the discovery is complete.

* Exit code 0 - Success
* Exit code 1 - Failure due to errors such as incorrect environment variables, arguments or configuration
* Exit code 2 - Partial Success when resource discovery was not able to find all the resources because of the service failures

### Generated Terraform Configuration Contents

The command will discover resources that are in an active or usable state. Resources that have been terminated or otherwise made inactive are generally excluded from the generated configuration.

By default, the Terraform names of the discovered resources will share the same name as the display name for that resource, if one exists.

The attributes of the resources will be populated with the values that are returned by the Oracle Cloud Infrastructure services.

In some cases, a required or optional attribute may not be discoverable from the Oracle Cloud Infrastructure services and may be omitted from the generated Terraform configuration.
This may be expected behavior from the service, which may prevent discovery of certain sensitive attributes or secrets. In such cases, placeholder value will be set along with a comment like this:

```
admin_password = "<placeholder for missing required attribute>" #Required attribute not found in discovery, placeholder value set to avoid plan failure
```

The missing required attributes will also be added to lifecycle ignore_changes. This is done to avoid terraform plan failure when moving manually-managed infrastructure to Terraform-managed infrastructure.
Any changes made to such fields will not reflect in terraform plan. If you want to update these fields, remove them from `ignore_changes`.

Resources that are dependent on availability domains will be generated under `availability_domain.tf` file. These include:
* oci\_core\_boot\_volume
* oci\_file\_storage\_file\_system
* oci\_file\_storage\_mount\_target
* oci\_file\_storage\_snapshot

### Exporting Identity Resources

Some resources, such as identity resources, may exist only at the tenancy level and cannot be discovered within a specific compartment. To discover such resources, specify
the following command.

```
terraform-provider-oci -command=export -output_path=<directory under which to generate Terraform files> -services=identity
```

> **Note**: When exporting identity resources, a `compartment_id` is not required. If a `compartment_id` is specified, the value will be ignored for discovering identity resources.


### Exporting Resources to Another Compartment
Once the user has reviewed the generated configuration and made the necessary changes to reflect the desired settings, the configuration can be used with Terraform.
One such use case is the re-deploying of those resources in a new compartment or tenancy, using Terraform.

To do so, specify the following environment variables:

```
export TF_VAR_tenancy_ocid=<new tenancy OCID>
export TF_VAR_compartment_ocid=<new compartment OCID>
```

And run

```
terraform apply
```

### Generating a Terraform State File

Using this command it is also possible to generate a Terraform state file to manage the discovered resources. To do so, run the following command:

```
terraform-provider-oci -command=export -compartment_id=<compartment to export> -output_path=<directory under which to generate Terraform files> -generate_state
```

The results of this command are both the `.tf` files representing the Terraform configuration and a `terraform.tfstate` file representing the state.

> **Note** The Terraform state file generated by this command is currently compatible with Terraform v0.12.4 and above


### Supported Resources
As of this writing, the list of Terraform services and resources that can be discovered by the command is as follows.
The list of supported resources can also be retrieved by running this command:

```
terraform-provider-oci -command=list_export_resources
```

analytics
    
* oci\_analytics\_analytics\_instance

apigateway
    
* oci\_apigateway\_gateway
* oci\_apigateway\_deployment

auto_scaling
    
* oci\_autoscaling\_auto\_scaling\_configuration

bds
    
* oci\_bds\_bds\_instance

blockchain
    
* oci\_blockchain\_blockchain\_platform
* oci\_blockchain\_peer
* oci\_blockchain\_osn

budget
    
* oci\_budget\_budget
* oci\_budget\_alert\_rule

cloud_guard
    
* oci\_cloud\_guard\_target
* oci\_cloud\_guard\_managed\_list
* oci\_cloud\_guard\_responder\_recipe
* oci\_cloud\_guard\_detector\_recipe

containerengine
    
* oci\_containerengine\_cluster
* oci\_containerengine\_node\_pool

core
    
* oci\_core\_boot\_volume\_backup
* oci\_core\_boot\_volume
* oci\_core\_console\_history
* oci\_core\_cluster\_network
* oci\_core\_compute\_image\_capability\_schema
* oci\_core\_cpe
* oci\_core\_cross\_connect\_group
* oci\_core\_cross\_connect
* oci\_core\_dhcp\_options
* oci\_core\_drg\_attachment
* oci\_core\_drg
* oci\_core\_dedicated\_vm\_host
* oci\_core\_image
* oci\_core\_instance\_configuration
* oci\_core\_instance\_console\_connection
* oci\_core\_instance\_pool
* oci\_core\_instance
* oci\_core\_internet\_gateway
* oci\_core\_ipsec
* oci\_core\_local\_peering\_gateway
* oci\_core\_nat\_gateway
* oci\_core\_network\_security\_group
* oci\_core\_network\_security\_group\_security\_rule
* oci\_core\_private\_ip
* oci\_core\_public\_ip
* oci\_core\_remote\_peering\_connection
* oci\_core\_route\_table
* oci\_core\_security\_list
* oci\_core\_service\_gateway
* oci\_core\_subnet
* oci\_core\_vcn
* oci\_core\_vlan
* oci\_core\_virtual\_circuit
* oci\_core\_vnic\_attachment
* oci\_core\_volume\_attachment
* oci\_core\_volume\_backup
* oci\_core\_volume\_backup\_policy
* oci\_core\_volume\_backup\_policy\_assignment
* oci\_core\_volume\_group
* oci\_core\_volume\_group\_backup
* oci\_core\_volume

data_safe
    
* oci\_data\_safe\_data\_safe\_private\_endpoint

database
    
* oci\_database\_autonomous\_container\_database
* oci\_database\_autonomous\_database
* oci\_database\_autonomous\_exadata\_infrastructure
* oci\_database\_autonomous\_vm\_cluster
* oci\_database\_backup\_destination
* oci\_database\_backup
* oci\_database\_database
* oci\_database\_db\_home
* oci\_database\_db\_system
* oci\_database\_exadata\_infrastructure
* oci\_database\_vm\_cluster\_network
* oci\_database\_vm\_cluster
* oci\_database\_database\_software\_image

datacatalog
    
* oci\_datacatalog\_catalog
* oci\_datacatalog\_data\_asset
* oci\_datacatalog\_connection
* oci\_datacatalog\_catalog\_private\_endpoint

dataflow
    
* oci\_dataflow\_application
* oci\_dataflow\_private\_endpoint

dataintegration
    
* oci\_dataintegration\_workspace

datascience
    
* oci\_datascience\_project
* oci\_datascience\_notebook\_session
* oci\_datascience\_model
* oci\_datascience\_model\_provenance

dns
    
* oci\_dns\_zone
* oci\_dns\_steering\_policy
* oci\_dns\_steering\_policy\_attachment
* oci\_dns\_tsig\_key
* oci\_dns\_rrset

email
    
* oci\_email\_sender

email_tenancy

* oci\_email\_suppression

events
    
* oci\_events\_rule

file_storage
    
* oci\_file\_storage\_file\_system
* oci\_file\_storage\_mount\_target
* oci\_file\_storage\_export
* oci\_file\_storage\_snapshot

functions
    
* oci\_functions\_application
* oci\_functions\_function

health_checks
    
* oci\_health\_checks\_http\_monitor
* oci\_health\_checks\_ping\_monitor

identity
    
* oci\_identity\_api\_key
* oci\_identity\_authentication\_policy
* oci\_identity\_auth\_token
* oci\_identity\_compartment
* oci\_identity\_customer\_secret\_key
* oci\_identity\_dynamic\_group
* oci\_identity\_group
* oci\_identity\_identity\_provider
* oci\_identity\_idp\_group\_mapping
* oci\_identity\_policy
* oci\_identity\_smtp\_credential
* oci\_identity\_swift\_password
* oci\_identity\_ui\_password
* oci\_identity\_user\_group\_membership
* oci\_identity\_user
* oci\_identity\_network\_source

integration
    
* oci\_integration\_integration\_instance

kms
    
* oci\_kms\_key
* oci\_kms\_key\_version
* oci\_kms\_vault

limits
    
* oci\_limits\_quota

load_balancer
    
* oci\_load\_balancer\_backend
* oci\_load\_balancer\_backend\_set
* oci\_load\_balancer\_certificate
* oci\_load\_balancer\_hostname
* oci\_load\_balancer\_listener
* oci\_load\_balancer\_load\_balancer
* oci\_load\_balancer\_path\_route\_set
* oci\_load\_balancer\_rule\_set

marketplace
    
* oci\_marketplace\_accepted\_agreement

monitoring
    
* oci\_monitoring\_alarm

mysql
    
* oci\_mysql\_mysql\_backup
* oci\_mysql\_mysql\_db\_system

nosql
    
* oci\_nosql\_table
* oci\_nosql\_index

object_storage
    
* oci\_objectstorage\_bucket
* oci\_objectstorage\_object\_lifecycle\_policy
* oci\_objectstorage\_object
* oci\_objectstorage\_preauthrequest
* oci\_objectstorage\_replication\_policy

oce
    
* oci\_oce\_oce\_instance

ocvp
    
* oci\_ocvp\_sddc
* oci\_ocvp\_esxi\_host

oda
    
* oci\_oda\_oda\_instance

ons
    
* oci\_ons\_notification\_topic
* oci\_ons\_subscription

osmanagement
    
* oci\_osmanagement\_managed\_instance\_group
* oci\_osmanagement\_software\_source

sch
    
* oci\_sch\_service\_connector

streaming
    
* oci\_streaming\_connect\_harness
* oci\_streaming\_stream\_pool
* oci\_streaming\_stream

tagging
    
* oci\_identity\_tag\_default
* oci\_identity\_tag\_namespace
* oci\_identity\_tag

waas
    
* oci\_waas\_address\_list
* oci\_waas\_custom\_protection\_rule
* oci\_waas\_http\_redirect
* oci\_waas\_waas\_policy
