---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_internetservicegroup"
description: |-
  Configure group of Internet Service.
---

# fortimanager_object_firewall_internetservicegroup
Configure group of Internet Service.

## Example Usage

```hcl
resource "fortimanager_object_firewall_internetservicegroup" "trname" {
  comment   = "fdsafdsafds"
  direction = "destination"
  member = [
    "8X8-8X8.Cloud",
    "ADP-DNS",
  ]
  name = "dd"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `direction` - How this service may be used (source, destination or both). Valid values: `both`, `source`, `destination`.

* `member` - Internet Service group member.
* `name` - Internet Service group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall InternetServiceGroup can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_internetservicegroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
