---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_securityexemptlist"
description: |-
  Configure security exemption list.
---

# fortimanager_object_user_securityexemptlist
Configure security exemption list.

## Example Usage

```hcl
resource "fortimanager_object_user_securityexemptlist" "trname" {
  description = "terraform-description"
  name        = "terraform-tefv-security"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `description` - Description.
* `name` - Name of the exempt list.
* `rule` - Rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `dstaddr` - Destination addresses or address groups.
* `id` - ID.
* `service` - Destination services.
* `srcaddr` - Source addresses or address groups.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser SecurityExemptList can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_securityexemptlist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
