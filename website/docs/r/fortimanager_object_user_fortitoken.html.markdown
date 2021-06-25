---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_fortitoken"
description: |-
  Configure FortiToken.
---

# fortimanager_object_user_fortitoken
Configure FortiToken.

## Example Usage

```hcl
resource "fortimanager_object_user_fortitoken" "labelname" {
  license       = "FTMTRIAL0BDD59BA"
  serial_number = "FTKMOB099A321EA2"
  status        = "active"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comments` - Comment.
* `license` - Mobile token license.
* `serial_number` - Serial number.
* `status` - Status Valid values: `lock`, `active`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{serial_number}}.

## Import

ObjectUser Fortitoken can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_fortitoken.labelname {{serial_number}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
