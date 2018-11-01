---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_app_catalog_subscription"
sidebar_current: "docs-oci-resource-core-app_catalog_subscription"
description: |-
  Provides the App Catalog Subscription resource in Oracle Cloud Infrastructure Core service
---

# oci_core_app_catalog_subscription
This resource provides the App Catalog Subscription resource in Oracle Cloud Infrastructure Core service.

Create a subscription for listing resource version for a compartment. It will take some time to propagate to all regions.


## Example Usage

```hcl
resource "oci_core_app_catalog_subscription" "test_app_catalog_subscription" {

	#Optional
	compartment_id = "${var.compartment_id}"
	eula_link = "${var.app_catalog_subscription_eula_link}"
	listing_id = "${oci_core_listing.test_listing.id}"
	listing_resource_version = "${var.app_catalog_subscription_listing_resource_version}"
	oracle_terms_of_use_link = "${var.app_catalog_subscription_oracle_terms_of_use_link}"
	signature = "${var.app_catalog_subscription_signature}"
	time_retrieved = "${var.app_catalog_subscription_time_retrieved}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The compartmentID for the subscription.
* `eula_link` - (Optional) EULA link
* `listing_id` - (Optional) The OCID of the listing.
* `listing_resource_version` - (Optional) Listing resource version.
* `oracle_terms_of_use_link` - (Optional) Oracle TOU link
* `signature` - (Optional) A generated signature for this listing resource version retrieved the agreements API.
* `time_retrieved` - (Optional) Date and time the agreements were retrieved, in RFC3339 format. Example: `2018-03-20T12:32:53.532Z` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartmentID of the subscription.
* `display_name` - The display name of the listing.
* `listing_id` - The ocid of the listing resource.
* `listing_resource_id` - Listing resource id.
* `listing_resource_version` - Listing resource version.
* `publisher_name` - Name of the publisher who published this listing.
* `summary` - The short summary to the listing.
* `time_created` - Date and time at which the subscription was created, in RFC3339 format. Example: `2018-03-20T12:32:53.532Z` 

## Import

AppCatalogSubscriptions can be imported using the `id`, e.g.

```
$ terraform import oci_core_app_catalog_subscription.test_app_catalog_subscription "compartmentId/{compartmentId}/listingId/{listingId}/listingResourceVersion/{listingResourceVersion}"
```

