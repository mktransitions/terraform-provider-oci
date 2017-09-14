# oci\_drg\_attachment

Gets a list of drg attachments.

## Example Usage

```
data "oci_core_drg_attachments" "t" {
    compartment_id = "compartment_id"
    drg_id = "drg_id"
    limit = 1
    page = "page"
    vcn_id = "vcn_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `vcn_id` - (Optional) The OCID of the VCN.
* `drg_id` - (Optional) The OCID of the DRG.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) The page to fetch.

## Attributes Reference

The following attributes are exported:

* `drg_attachments` - The list of images.

## Drg Attachment reference
* `compartment_id` - The OCID of the compartment containing the DRG attachment.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `drg_id` - The OCID of the DRG.
* `id` - The DRG attachment's Oracle ID (OCID).
* `state` - The DRG attachment's current state: [ATTACHING, ATTACHED, DETACHING, DETACHED].
* `time_created` - The date and time the image was created.
* `vcn_id` - The OCID of the VCN.

