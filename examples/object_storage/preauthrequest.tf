// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to create preauthenticated requests for objects and buckets.
 */

resource "oci_objectstorage_preauthrequest" "bucket_par" {
  namespace    = data.oci_objectstorage_namespace.ns.namespace
  bucket       = oci_objectstorage_bucket.bucket1.name
  name         = "parOnBucket"
  access_type  = "AnyObjectWrite" //Other configurations accepted are ObjectWrite, ObjectReadWrite
  time_expires = "2020-12-10T23:00:00Z"
}

resource "oci_objectstorage_preauthrequest" "object_par" {
  namespace    = data.oci_objectstorage_namespace.ns.namespace
  bucket       = oci_objectstorage_bucket.bucket1.name
  object       = oci_objectstorage_object.object1.object
  name         = "objectPar"
  access_type  = "ObjectRead" // ObjectRead, ObjectWrite, ObjectReadWrite, AnyObjectWrite
  time_expires = "2020-12-29T23:00:00Z"
}

output "par_request_url" {
  value = "https://objectstorage.${var.region}.oraclecloud.com${oci_objectstorage_preauthrequest.object_par.access_uri}"
}

