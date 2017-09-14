#     ___  ____     _    ____ _     _____
#    / _ \|  _ \   / \  / ___| |   | ____|
#   | | | | |_) | / _ \| |   | |   |  _|
#   | |_| |  _ < / ___ | |___| |___| |___
#    \___/|_| \_/_/   \_\____|_____|_____|

## PAR(preauthenticated requests) management.
## This example creates two preauthenticated requests, one on a bucket and the second on an object
## Export the environmental variables used for configuration.
## Provide the script with an object and bucket present in your namespace

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "namespace_name" {}
variable "region" {}

variable "bucket_name" {}
variable "object_name" {}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}


resource "oci_objectstorage_preauthrequest" "parOnBucket" {
  namespace = "${var.namespace_name}"
  bucket = "${var.bucket_name}"
  name = "parOnBucket"
  access_type = "AnyObjectWrite" //Other configurations accepted are ObjectWrite, ObjectReadWrite
  time_expires = "2019-11-10T23:00:00Z"
}

resource "oci_objectstorage_preauthrequest" "parOnObject" {
  namespace = "${var.namespace_name}"
  bucket = "${var.bucket_name}"
  object = "${var.object_name}"
  name = "parOnObject"
  access_type = "ObjectRead" //Other configurations accepted are ObjectWrite, ObjectReadWrite
  time_expires = "2019-11-10T23:00:00Z"
}

output "par_request_url" {
  value = "${oci_objectstorage_preauthrequest.parOnObject.access_uri}"
}

output "par_request_name" {
  value = "${oci_objectstorage_preauthrequest.parOnObject.name}"
}
