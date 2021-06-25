---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_bleprofile"
description: |-
  Configure Bluetooth Low Energy profile.
---

# fortimanager_object_wirelesscontroller_bleprofile
Configure Bluetooth Low Energy profile.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_bleprofile" "trname" {
  beacon_interval          = 100
  ble_scanning             = "disable"
  comment                  = "This is a Terraform example"
  eddystone_instance       = "abcdef"
  eddystone_namespace      = "0102030405"
  eddystone_url            = "http://www.fortinet.com"
  eddystone_url_encode_hex = "00666f7274696e65-7407"
  ibeacon_uuid             = "005ea414-cbd1-11e5-9956-625662870761"
  major_id                 = 1000
  minor_id                 = 2000
  name                     = "terr-wils-ctler-ble-profile"
  txpower                  = "0"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `advertising` - Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.

* `beacon_interval` - Beacon interval (default = 100 msec).
* `ble_scanning` - Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `disable`, `enable`.

* `comment` - Comment.
* `eddystone_instance` - Eddystone instance ID.
* `eddystone_namespace` - Eddystone namespace ID.
* `eddystone_url` - Eddystone URL.
* `eddystone_url_encode_hex` - Eddystone encoded URL hexadecimal string
* `ibeacon_uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `major_id` - Major ID.
* `minor_id` - Minor ID.
* `name` - Bluetooth Low Energy profile name.
* `txpower` - Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController BleProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_bleprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
