---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_operations_insights_warehouse_rotate_warehouse_wallet"
sidebar_current: "docs-oci-resource-opsi-operations_insights_warehouse_rotate_warehouse_wallet"
description: |-
  Provides the Operations Insights Warehouse Rotate Warehouse Wallet resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_operations_insights_warehouse_rotate_warehouse_wallet
This resource provides the Operations Insights Warehouse Rotate Warehouse Wallet resource in Oracle Cloud Infrastructure Opsi service.

Rotate the ADW wallet for Operations Insights Warehouse using which the Hub data is exposed.

## Example Usage

```hcl
resource "oci_opsi_operations_insights_warehouse_rotate_warehouse_wallet" "test_operations_insights_warehouse_rotate_warehouse_wallet" {
	#Required
	operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id
}
```

## Argument Reference

The following arguments are supported:

* `operations_insights_warehouse_id` - (Required) Unique Ops Insights Warehouse identifier


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Operations Insights Warehouse Rotate Warehouse Wallet
	* `update` - (Defaults to 20 minutes), when updating the Operations Insights Warehouse Rotate Warehouse Wallet
	* `delete` - (Defaults to 20 minutes), when destroying the Operations Insights Warehouse Rotate Warehouse Wallet


## Import

OperationsInsightsWarehouseRotateWarehouseWallet can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_operations_insights_warehouse_rotate_warehouse_wallet.test_operations_insights_warehouse_rotate_warehouse_wallet "id"
```

