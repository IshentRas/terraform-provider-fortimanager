---
subcategory: "Security Console"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_securityconsole_install_package"
description: |-
  Copy and install a policy package to devices.
---

# fortimanager_securityconsole_install_package
Copy and install a policy package to devices.

## Example Usage

```hcl
resource "fortimanager_securityconsole_install_package" "trname" {
  fmgadom           = "root"
  adom_rev_comments = "terraform-comment"
  adom_rev_name     = "terraform-tefv"
  dev_rev_comments  = "terraform-comment"
  flags             = ["none", "cp_all_objs", "preview", "generate_rev", "copy_assigned_pkg", "unassign", "ifpolicy_only", "no_ifpolicy", "objs_only", "auto_lock_ws", "check_pkg_st", "copy_only"]
  pkg               = "default"
}
```

## Argument Reference


The following arguments are supported:


* `fmgadom` - Source ADOM name.
* `adom_rev_comments` - If "generate_rev" flag is set, the comment for the new ADOM revision.
* `adom_rev_name` - If "generate_rev" flag is set, the name for the new ADOM revision.
* `dev_rev_comments` - Comments for the device configuration revision that will be generated during install.
* `flags` - cp_all_objs - Assign all objects during global policy assignment. preview - Generate preview cache only. generate_rev - Generate new ADOM revision before install. copy_assigned_pkg - For global policy assignment - copy assigned package from ADOM to device. unassign - Remove global policy from ADOM. ifpolicy_only - Only install interface policies. no_ifpolicy - Install regular policies only - do not install interface policies. objs_only - Install object(s) only - do not install any policies. auto_lock_ws - Automatically lock and unlock workspace when performing security console task. copy_only - Only copy to device db. Valid values: `none`, `cp_all_objs`, `preview`, `generate_rev`, `copy_assigned_pkg`, `unassign`, `ifpolicy_only`, `no_ifpolicy`, `objs_only`, `auto_lock_ws`, `check_pkg_st`, `copy_only`.

* `pkg` - Source package path and name.
* `scope` - Scope. The structure of `scope` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.

The `scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Securityconsole InstallPackage can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_securityconsole_install_package.labelname SecurityconsoleInstallPackage
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
