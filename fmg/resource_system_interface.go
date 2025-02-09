// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Interface configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceCreate,
		Read:   resourceSystemInterfaceRead,
		Update: resourceSystemInterfaceUpdate,
		Delete: resourceSystemInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"autogenerated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_allowaccess": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ip6_autoconf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"rating_service_ip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"serviceaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"speed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_service_ip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterface resource while getting object: %v", err)
	}

	v, _ := d.GetOk("type")
	ag, _ := d.GetOk("autogenerated")
	if v == "physical" || ag == "auto" {
		_, err = c.UpdateSystemInterface(obj, adomv, (*obj)["name"].(string), nil)
	} else {
		_, err = c.CreateSystemInterface(obj, adomv, nil)
	}

	if err != nil {
		return fmt.Errorf("Error creating SystemInterface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemInterfaceRead(d, m)
}

func resourceSystemInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemInterface(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemInterfaceRead(d, m)
}

func resourceSystemInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	v, _ := d.GetOk("type")
	ag, _ := d.GetOk("autogenerated")
	if v == "physical" || ag == "auto" {
		d.SetId("")
		return nil
	}

	err = c.DeleteSystemInterface(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemInterface(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterface resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ip6_address"
	if _, ok := i["ip6-address"]; ok {
		result["ip6_address"] = flattenSystemInterfaceIpv6Ip6Address(i["ip6-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := i["ip6-allowaccess"]; ok {
		result["ip6_allowaccess"] = flattenSystemInterfaceIpv6Ip6Allowaccess(i["ip6-allowaccess"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_autoconf"
	if _, ok := i["ip6-autoconf"]; ok {
		result["ip6_autoconf"] = flattenSystemInterfaceIpv6Ip6Autoconf(i["ip6-autoconf"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemInterfaceIpv6Ip6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6Allowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceIpv6Ip6Autoconf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceRatingServiceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceServiceaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemInterfaceSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceUpdateServiceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("alias", flattenSystemInterfaceAlias(o["alias"], d, "alias")); err != nil {
		if vv, ok := fortiAPIPatch(o["alias"], "SystemInterface-Alias"); ok {
			if err = d.Set("alias", vv); err != nil {
				return fmt.Errorf("Error reading alias: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenSystemInterfaceAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "SystemInterface-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemInterfaceDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemInterface-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemInterfaceIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemInterface-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipv6", flattenSystemInterfaceIpv6(o["ipv6"], d, "ipv6")); err != nil {
			if vv, ok := fortiAPIPatch(o["ipv6"], "SystemInterface-Ipv6"); ok {
				if err = d.Set("ipv6", vv); err != nil {
					return fmt.Errorf("Error reading ipv6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv6"); ok {
			if err = d.Set("ipv6", flattenSystemInterfaceIpv6(o["ipv6"], d, "ipv6")); err != nil {
				if vv, ok := fortiAPIPatch(o["ipv6"], "SystemInterface-Ipv6"); ok {
					if err = d.Set("ipv6", vv); err != nil {
						return fmt.Errorf("Error reading ipv6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ipv6: %v", err)
				}
			}
		}
	}

	if err = d.Set("mtu", flattenSystemInterfaceMtu(o["mtu"], d, "mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["mtu"], "SystemInterface-Mtu"); ok {
			if err = d.Set("mtu", vv); err != nil {
				return fmt.Errorf("Error reading mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mtu: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemInterfaceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("rating_service_ip", flattenSystemInterfaceRatingServiceIp(o["rating-service-ip"], d, "rating_service_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["rating-service-ip"], "SystemInterface-RatingServiceIp"); ok {
			if err = d.Set("rating_service_ip", vv); err != nil {
				return fmt.Errorf("Error reading rating_service_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rating_service_ip: %v", err)
		}
	}

	if err = d.Set("serviceaccess", flattenSystemInterfaceServiceaccess(o["serviceaccess"], d, "serviceaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["serviceaccess"], "SystemInterface-Serviceaccess"); ok {
			if err = d.Set("serviceaccess", vv); err != nil {
				return fmt.Errorf("Error reading serviceaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serviceaccess: %v", err)
		}
	}

	if err = d.Set("speed", flattenSystemInterfaceSpeed(o["speed"], d, "speed")); err != nil {
		if vv, ok := fortiAPIPatch(o["speed"], "SystemInterface-Speed"); ok {
			if err = d.Set("speed", vv); err != nil {
				return fmt.Errorf("Error reading speed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading speed: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemInterfaceStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemInterface-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("update_service_ip", flattenSystemInterfaceUpdateServiceIp(o["update-service-ip"], d, "update_service_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-service-ip"], "SystemInterface-UpdateServiceIp"); ok {
			if err = d.Set("update_service_ip", vv); err != nil {
				return fmt.Errorf("Error reading update_service_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_service_ip: %v", err)
		}
	}

	return nil
}

func flattenSystemInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemInterfaceIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ip6_address"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-address"], _ = expandSystemInterfaceIpv6Ip6Address(d, i["ip6_address"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-allowaccess"], _ = expandSystemInterfaceIpv6Ip6Allowaccess(d, i["ip6_allowaccess"], pre_append)
	} else {
		result["ip6-allowaccess"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip6_autoconf"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-autoconf"], _ = expandSystemInterfaceIpv6Ip6Autoconf(d, i["ip6_autoconf"], pre_append)
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6Allowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceIpv6Ip6Autoconf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceRatingServiceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemInterfaceServiceaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemInterfaceSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceUpdateServiceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectSystemInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("alias"); ok {
		t, err := expandSystemInterfaceAlias(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandSystemInterfaceAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSystemInterfaceDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandSystemInterfaceIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok {
		t, err := expandSystemInterfaceIpv6(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("mtu"); ok {
		t, err := expandSystemInterfaceMtu(d, v, "mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemInterfaceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rating_service_ip"); ok {
		t, err := expandSystemInterfaceRatingServiceIp(d, v, "rating_service_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rating-service-ip"] = t
		}
	}

	if v, ok := d.GetOk("serviceaccess"); ok {
		t, err := expandSystemInterfaceServiceaccess(d, v, "serviceaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serviceaccess"] = t
		}
	}

	if v, ok := d.GetOk("speed"); ok {
		t, err := expandSystemInterfaceSpeed(d, v, "speed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speed"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemInterfaceStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("update_service_ip"); ok {
		t, err := expandSystemInterfaceUpdateServiceIp(d, v, "update_service_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-service-ip"] = t
		}
	}

	return &obj, nil
}
