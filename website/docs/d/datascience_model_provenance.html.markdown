---
subcategory: "Datascience"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_provenance"
sidebar_current: "docs-oci-datasource-datascience-model_provenance"
description: |-
  Provides details about a specific Model Provenance in Oracle Cloud Infrastructure Datascience service
---

# Data Source: oci_datascience_model_provenance
This data source provides details about a specific Model Provenance resource in Oracle Cloud Infrastructure Datascience service.

Gets provenance information for specified model.

## Example Usage

```hcl
data "oci_datascience_model_provenance" "test_model_provenance" {
	#Required
	model_id = oci_datascience_model.test_model.id
}
```

## Argument Reference

The following arguments are supported:

* `model_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the model.


## Attributes Reference

The following attributes are exported:

* `git_branch` - For model reproducibility purposes. Branch of the git repository associated with model training.
* `git_commit` - For model reproducibility purposes. Commit ID of the git repository associated with model training.
* `repository_url` - For model reproducibility purposes. URL of the git repository associated with model training.
* `script_dir` - For model reproducibility purposes. Path to model artifacts.
* `training_script` - For model reproducibility purposes. Path to the python script or notebook in which the model was trained." 

