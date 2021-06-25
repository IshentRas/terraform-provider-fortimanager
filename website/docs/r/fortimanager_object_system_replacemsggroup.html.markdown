---
subcategory: "Object System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_replacemsggroup"
description: |-
  Configure replacement message groups.
---

# fortimanager_object_system_replacemsggroup
Configure replacement message groups.

## Example Usage

```hcl
resource "fortimanager_object_system_replacemsggroup" "trname" {
  comment = "terraform-comment"
  name    = "terraform-tefv"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `admin` - Admin. The structure of `admin` block is documented below.
* `alertmail` - Alertmail. The structure of `alertmail` block is documented below.
* `auth` - Auth. The structure of `auth` block is documented below.
* `automation` - Automation. The structure of `automation` block is documented below. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `comment` - Comment.
* `custom_message` - Custom-Message. The structure of `custom_message` block is documented below.
* `device_detection_portal` - Device-Detection-Portal. The structure of `device_detection_portal` block is documented below. (`ver Controlled FortiOS <= 6.4`)
* `fortiguard_wf` - Fortiguard-Wf. The structure of `fortiguard_wf` block is documented below.
* `ftp` - Ftp. The structure of `ftp` block is documented below.
* `group_type` - Group type. Valid values: `default`, `utm`, `auth`, `ec`, `captive-portal`.

* `http` - Http. The structure of `http` block is documented below.
* `icap` - Icap. The structure of `icap` block is documented below.
* `mail` - Mail. The structure of `mail` block is documented below.
* `mm1` - Mm1. The structure of `mm1` block is documented below. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `mm3` - Mm3. The structure of `mm3` block is documented below. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `mm4` - Mm4. The structure of `mm4` block is documented below. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `mm7` - Mm7. The structure of `mm7` block is documented below. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `mms` - Mms. The structure of `mms` block is documented below. (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `nac_quar` - Nac-Quar. The structure of `nac_quar` block is documented below.
* `name` - Group name.
* `nntp` - Nntp. The structure of `nntp` block is documented below. (`ver Controlled FortiOS <= 6.4`)
* `spam` - Spam. The structure of `spam` block is documented below.
* `sslvpn` - Sslvpn. The structure of `sslvpn` block is documented below.
* `traffic_quota` - Traffic-Quota. The structure of `traffic_quota` block is documented below.
* `utm` - Utm. The structure of `utm` block is documented below.
* `webproxy` - Webproxy. The structure of `webproxy` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `admin` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `alertmail` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `auth` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `automation` block supports (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`):

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `custom_message` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `device_detection_portal` block supports (`ver Controlled FortiOS <= 6.4`):

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `fortiguard_wf` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `ftp` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `http` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `icap` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `mail` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `mm1` block supports (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`):

* `add_smil` - add message encapsulation Valid values: `disable`, `enable`.

* `charset` - character encoding used for replacement message Valid values: `us-ascii`, `utf-8`.

* `class` - message class Valid values: `personal`, `advertisement`, `information`, `automatic`, `not-included`.

* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `from` - from address
* `from_sender` - notification message sent from recipient Valid values: `disable`, `enable`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `image` - Message string.
* `message` - message text
* `msg_type` - Message type.
* `priority` - message priority Valid values: `low`, `normal`, `high`, `not-included`.

* `rsp_status` - response status code Valid values: `ok`, `err-unspecified`, `err-srv-denied`, `err-msg-fmt-corrupt`, `err-snd-addr-unresolv`, `err-msg-not-found`, `err-net-prob`, `err-content-not-accept`, `err-unsupp-msg`.

* `rsp_text` - response text
* `sender_visibility` - sender visibility Valid values: `hide`, `show`, `not-specified`.

* `smil_part` - message encapsulation text
* `subject` - subject text string

The `mm3` block supports (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`):

* `add_html` - add message encapsulation Valid values: `disable`, `enable`.

* `charset` - character encoding used for replacement message Valid values: `us-ascii`, `utf-8`.

* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `from` - from address
* `from_sender` - notification message sent from recipient Valid values: `disable`, `enable`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `html_part` - message encapsulation text
* `image` - Message string.
* `message` - message text
* `msg_type` - Message type.
* `priority` - message priority Valid values: `low`, `normal`, `high`, `not-included`.

* `subject` - subject text string

The `mm4` block supports (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`):

* `add_smil` - add message encapsulation Valid values: `disable`, `enable`.

* `charset` - character encoding used for replacement message Valid values: `us-ascii`, `utf-8`.

* `class` - message class Valid values: `personal`, `advertisement`, `informational`, `auto`, `not-included`.

* `domain` - from address domain
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `from` - from address
* `from_sender` - notification message sent from recipient Valid values: `disable`, `enable`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `image` - Message string.
* `message` - message text
* `msg_type` - Message type.
* `priority` - message priority Valid values: `low`, `normal`, `high`, `not-included`.

* `rsp_status` - response status Valid values: `ok`, `err-unspecified`, `err-srv-denied`, `err-msg-fmt-corrupt`, `err-snd-addr-unresolv`, `err-net-prob`, `err-content-not-accept`, `err-unsupp-msg`.

* `smil_part` - message encapsulation text
* `subject` - subject text string

The `mm7` block supports (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`):

* `add_smil` - add message encapsulation Valid values: `disable`, `enable`.

* `addr_type` - from address type Valid values: `rfc2822-addr`, `number`, `short-code`.

* `allow_content_adaptation` - allow content adaptations Valid values: `disable`, `enable`.

* `charset` - character encoding used for replacement message Valid values: `us-ascii`, `utf-8`.

* `class` - message class Valid values: `personal`, `advertisement`, `informational`, `auto`, `not-included`.

* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `from` - from address
* `from_sender` - notification message sent from recipient Valid values: `disable`, `enable`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `image` - Message string.
* `message` - message text
* `msg_type` - Message type.
* `priority` - message priority Valid values: `low`, `normal`, `high`, `not-included`.

* `rsp_status` - response status Valid values: `success`, `partial-success`, `client-err`, `oper-restrict`, `addr-err`, `addr-not-found`, `content-refused`, `msg-id-not-found`, `link-id-not-found`, `msg-fmt-corrupt`, `app-id-not-found`, `repl-app-id-not-found`, `srv-err`, `not-possible`, `msg-rejected`, `multiple-addr-not-supp`, `app-addr-not-supp`, `gen-service-err`, `improper-ident`, `unsupp-ver`, `unsupp-oper`, `validation-err`, `service-err`, `service-unavail`, `service-denied`, `app-denied`.

* `smil_part` - message encapsulation text
* `subject` - subject text string

The `mms` block supports (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`):

* `buffer` - Message string.
* `charset` - character encoding used for replacement message Valid values: `us-ascii`, `utf-8`.

* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `image` - Message string.
* `msg_type` - Message type.

The `nac_quar` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `nntp` block supports (`ver Controlled FortiOS <= 6.4`):

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `spam` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `sslvpn` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `traffic_quota` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `utm` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.

The `webproxy` block supports:

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem ReplacemsgGroup can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_replacemsggroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
