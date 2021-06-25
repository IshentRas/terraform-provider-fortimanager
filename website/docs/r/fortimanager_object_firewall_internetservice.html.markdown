---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_internetservice"
description: |-
  Show Internet Service application.
---

# fortimanager_object_firewall_internetservice
Show Internet Service application.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `city` - City sequence number list. (`ver Controlled FortiOS >= 6.4`)
* `country` - Country sequence number list. (`ver Controlled FortiOS >= 6.4`)
* `database` - Database. Valid values: `isdb`, `irdb`.

* `direction` - Direction. Valid values: `src`, `dst`, `both`.

* `extra_ip_range_number` - Extra-Ip-Range-Number.
* `icon_id` - Icon-Id.
* `fosid` - Id.
* `ip_number` - Ip-Number.
* `ip_range_number` - Ip-Range-Number.
* `jitter_threshold` - Jitter-Threshold. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `latency_threshold` - Latency-Threshold. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `name` - Name.
* `obsolete` - Obsolete.
* `region` - Region sequence number list. (`ver Controlled FortiOS >= 6.4`)
* `packetloss_threshold` - Packetloss-Threshold. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `reputation` - Reputation. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `singularity` - Singularity.
* `sld_id` - Sld-Id. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall InternetService can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_internetservice.labelname ObjectFirewallInternetService
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
