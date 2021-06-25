---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_pxgrid"
description: |-
  ObjectUser Pxgrid
---

# fortimanager_object_user_pxgrid
ObjectUser Pxgrid

## Example Usage

```hcl
resource "fortimanager_object_user_pxgrid" "trname" {
  name   = "terr-user-pxgrid"
  server = "192.168.1.1"
  status = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `ca_cert` - Ca-Cert.
* `client_cert` - Client-Cert.
* `name` - Name.
* `server` - Server.
* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Pxgrid can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_pxgrid.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
