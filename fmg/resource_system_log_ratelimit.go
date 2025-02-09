// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Logging rate limit.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLogRatelimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogRatelimitUpdate,
		Read:   resourceSystemLogRatelimitRead,
		Update: resourceSystemLogRatelimitUpdate,
		Delete: resourceSystemLogRatelimitDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"device": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"filter_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ratelimit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"device_ratelimit_default": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_ratelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemLogRatelimitUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogRatelimit(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogRatelimit resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogRatelimit(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogRatelimit resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogRatelimit")

	return resourceSystemLogRatelimitRead(d, m)
}

func resourceSystemLogRatelimitDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogRatelimit(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogRatelimit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogRatelimitRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogRatelimit(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogRatelimit resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogRatelimit(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogRatelimit resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogRatelimitDeviceSLR(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := i["device"]; ok {
			v := flattenSystemLogRatelimitDeviceDeviceSLR(i["device"], d, pre_append)
			tmp["device"] = fortiAPISubPartPatch(v, "SystemLogRatelimit-Device-Device")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := i["filter-type"]; ok {
			v := flattenSystemLogRatelimitDeviceFilterTypeSLR(i["filter-type"], d, pre_append)
			tmp["filter_type"] = fortiAPISubPartPatch(v, "SystemLogRatelimit-Device-FilterType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemLogRatelimitDeviceIdSLR(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemLogRatelimit-Device-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ratelimit"
		if _, ok := i["ratelimit"]; ok {
			v := flattenSystemLogRatelimitDeviceRatelimitSLR(i["ratelimit"], d, pre_append)
			tmp["ratelimit"] = fortiAPISubPartPatch(v, "SystemLogRatelimit-Device-Ratelimit")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLogRatelimitDeviceDeviceSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitDeviceFilterTypeSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitDeviceIdSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitDeviceRatelimitSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitDeviceRatelimitDefaultSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitModeSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogRatelimitSystemRatelimitSLR(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogRatelimit(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("device", flattenSystemLogRatelimitDeviceSLR(o["device"], d, "device")); err != nil {
			if vv, ok := fortiAPIPatch(o["device"], "SystemLogRatelimit-Device"); ok {
				if err = d.Set("device", vv); err != nil {
					return fmt.Errorf("Error reading device: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading device: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("device"); ok {
			if err = d.Set("device", flattenSystemLogRatelimitDeviceSLR(o["device"], d, "device")); err != nil {
				if vv, ok := fortiAPIPatch(o["device"], "SystemLogRatelimit-Device"); ok {
					if err = d.Set("device", vv); err != nil {
						return fmt.Errorf("Error reading device: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading device: %v", err)
				}
			}
		}
	}

	if err = d.Set("device_ratelimit_default", flattenSystemLogRatelimitDeviceRatelimitDefaultSLR(o["device-ratelimit-default"], d, "device_ratelimit_default")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-ratelimit-default"], "SystemLogRatelimit-DeviceRatelimitDefault"); ok {
			if err = d.Set("device_ratelimit_default", vv); err != nil {
				return fmt.Errorf("Error reading device_ratelimit_default: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_ratelimit_default: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemLogRatelimitModeSLR(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemLogRatelimit-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("system_ratelimit", flattenSystemLogRatelimitSystemRatelimitSLR(o["system-ratelimit"], d, "system_ratelimit")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-ratelimit"], "SystemLogRatelimit-SystemRatelimit"); ok {
			if err = d.Set("system_ratelimit", vv); err != nil {
				return fmt.Errorf("Error reading system_ratelimit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_ratelimit: %v", err)
		}
	}

	return nil
}

func flattenSystemLogRatelimitFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogRatelimitDeviceSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device"], _ = expandSystemLogRatelimitDeviceDeviceSLR(d, i["device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-type"], _ = expandSystemLogRatelimitDeviceFilterTypeSLR(d, i["filter_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemLogRatelimitDeviceIdSLR(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ratelimit"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ratelimit"], _ = expandSystemLogRatelimitDeviceRatelimitSLR(d, i["ratelimit"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLogRatelimitDeviceDeviceSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitDeviceFilterTypeSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitDeviceIdSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitDeviceRatelimitSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitDeviceRatelimitDefaultSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitModeSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogRatelimitSystemRatelimitSLR(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogRatelimit(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok {
		t, err := expandSystemLogRatelimitDeviceSLR(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("device_ratelimit_default"); ok {
		t, err := expandSystemLogRatelimitDeviceRatelimitDefaultSLR(d, v, "device_ratelimit_default")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-ratelimit-default"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemLogRatelimitModeSLR(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("system_ratelimit"); ok {
		t, err := expandSystemLogRatelimitSystemRatelimitSLR(d, v, "system_ratelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-ratelimit"] = t
		}
	}

	return &obj, nil
}
