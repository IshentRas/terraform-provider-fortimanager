// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure custom services.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallServiceCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallServiceCustomCreate,
		Read:   resourceObjectFirewallServiceCustomRead,
		Update: resourceObjectFirewallServiceCustomUpdate,
		Delete: resourceObjectFirewallServiceCustomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
					"inherit",
				}, false),
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"app_service_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_reset_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_object": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"helper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmpcode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icmptype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iprange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sctp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_halfclose_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_halfopen_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_rst_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_timewait_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"udp_idle_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"udp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallServiceCustomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallServiceCustom(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallServiceCustom resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallServiceCustom(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallServiceCustom resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallServiceCustomRead(d, m)
}

func resourceObjectFirewallServiceCustomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallServiceCustom(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallServiceCustom resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallServiceCustom(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallServiceCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallServiceCustomRead(d, m)
}

func resourceObjectFirewallServiceCustomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallServiceCustom(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallServiceCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallServiceCustomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallServiceCustom(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallServiceCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallServiceCustom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallServiceCustom resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallServiceCustomAppCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallServiceCustomAppServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallServiceCustomCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenObjectFirewallServiceCustomCheckResetRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomHelper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomIcmpcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomIcmptype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomIprange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomProtocolNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomSctpPortrange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomSessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomTcpHalfcloseTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomTcpHalfopenTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomTcpPortrange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomTcpRstTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomTcpTimewaitTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomUdpIdleTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallServiceCustomUdpPortrange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenObjectFirewallServiceCustomVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallServiceCustom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("app_category", flattenObjectFirewallServiceCustomAppCategory(o["app-category"], d, "app_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-category"], "ObjectFirewallServiceCustom-AppCategory"); ok {
			if err = d.Set("app_category", vv); err != nil {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_category: %v", err)
		}
	}

	if err = d.Set("app_service_type", flattenObjectFirewallServiceCustomAppServiceType(o["app-service-type"], d, "app_service_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-service-type"], "ObjectFirewallServiceCustom-AppServiceType"); ok {
			if err = d.Set("app_service_type", vv); err != nil {
				return fmt.Errorf("Error reading app_service_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_service_type: %v", err)
		}
	}

	if err = d.Set("application", flattenObjectFirewallServiceCustomApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "ObjectFirewallServiceCustom-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("category", flattenObjectFirewallServiceCustomCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "ObjectFirewallServiceCustom-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("check_reset_range", flattenObjectFirewallServiceCustomCheckResetRange(o["check-reset-range"], d, "check_reset_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-reset-range"], "ObjectFirewallServiceCustom-CheckResetRange"); ok {
			if err = d.Set("check_reset_range", vv); err != nil {
				return fmt.Errorf("Error reading check_reset_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_reset_range: %v", err)
		}
	}

	if err = d.Set("color", flattenObjectFirewallServiceCustomColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFirewallServiceCustom-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectFirewallServiceCustomComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallServiceCustom-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenObjectFirewallServiceCustomFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "ObjectFirewallServiceCustom-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenObjectFirewallServiceCustomFqdn(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "ObjectFirewallServiceCustom-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("global_object", flattenObjectFirewallServiceCustomGlobalObject(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "ObjectFirewallServiceCustom-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("helper", flattenObjectFirewallServiceCustomHelper(o["helper"], d, "helper")); err != nil {
		if vv, ok := fortiAPIPatch(o["helper"], "ObjectFirewallServiceCustom-Helper"); ok {
			if err = d.Set("helper", vv); err != nil {
				return fmt.Errorf("Error reading helper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading helper: %v", err)
		}
	}

	if err = d.Set("icmpcode", flattenObjectFirewallServiceCustomIcmpcode(o["icmpcode"], d, "icmpcode")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmpcode"], "ObjectFirewallServiceCustom-Icmpcode"); ok {
			if err = d.Set("icmpcode", vv); err != nil {
				return fmt.Errorf("Error reading icmpcode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmpcode: %v", err)
		}
	}

	if err = d.Set("icmptype", flattenObjectFirewallServiceCustomIcmptype(o["icmptype"], d, "icmptype")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmptype"], "ObjectFirewallServiceCustom-Icmptype"); ok {
			if err = d.Set("icmptype", vv); err != nil {
				return fmt.Errorf("Error reading icmptype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmptype: %v", err)
		}
	}

	if err = d.Set("iprange", flattenObjectFirewallServiceCustomIprange(o["iprange"], d, "iprange")); err != nil {
		if vv, ok := fortiAPIPatch(o["iprange"], "ObjectFirewallServiceCustom-Iprange"); ok {
			if err = d.Set("iprange", vv); err != nil {
				return fmt.Errorf("Error reading iprange: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iprange: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallServiceCustomName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallServiceCustom-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectFirewallServiceCustomProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectFirewallServiceCustom-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("protocol_number", flattenObjectFirewallServiceCustomProtocolNumber(o["protocol-number"], d, "protocol_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol-number"], "ObjectFirewallServiceCustom-ProtocolNumber"); ok {
			if err = d.Set("protocol_number", vv); err != nil {
				return fmt.Errorf("Error reading protocol_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol_number: %v", err)
		}
	}

	if err = d.Set("proxy", flattenObjectFirewallServiceCustomProxy(o["proxy"], d, "proxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy"], "ObjectFirewallServiceCustom-Proxy"); ok {
			if err = d.Set("proxy", vv); err != nil {
				return fmt.Errorf("Error reading proxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("sctp_portrange", flattenObjectFirewallServiceCustomSctpPortrange(o["sctp-portrange"], d, "sctp_portrange")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-portrange"], "ObjectFirewallServiceCustom-SctpPortrange"); ok {
			if err = d.Set("sctp_portrange", vv); err != nil {
				return fmt.Errorf("Error reading sctp_portrange: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_portrange: %v", err)
		}
	}

	if err = d.Set("session_ttl", flattenObjectFirewallServiceCustomSessionTtl(o["session-ttl"], d, "session_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-ttl"], "ObjectFirewallServiceCustom-SessionTtl"); ok {
			if err = d.Set("session_ttl", vv); err != nil {
				return fmt.Errorf("Error reading session_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("tcp_halfclose_timer", flattenObjectFirewallServiceCustomTcpHalfcloseTimer(o["tcp-halfclose-timer"], d, "tcp_halfclose_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-halfclose-timer"], "ObjectFirewallServiceCustom-TcpHalfcloseTimer"); ok {
			if err = d.Set("tcp_halfclose_timer", vv); err != nil {
				return fmt.Errorf("Error reading tcp_halfclose_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_halfclose_timer: %v", err)
		}
	}

	if err = d.Set("tcp_halfopen_timer", flattenObjectFirewallServiceCustomTcpHalfopenTimer(o["tcp-halfopen-timer"], d, "tcp_halfopen_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-halfopen-timer"], "ObjectFirewallServiceCustom-TcpHalfopenTimer"); ok {
			if err = d.Set("tcp_halfopen_timer", vv); err != nil {
				return fmt.Errorf("Error reading tcp_halfopen_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_halfopen_timer: %v", err)
		}
	}

	if err = d.Set("tcp_portrange", flattenObjectFirewallServiceCustomTcpPortrange(o["tcp-portrange"], d, "tcp_portrange")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-portrange"], "ObjectFirewallServiceCustom-TcpPortrange"); ok {
			if err = d.Set("tcp_portrange", vv); err != nil {
				return fmt.Errorf("Error reading tcp_portrange: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_portrange: %v", err)
		}
	}

	if err = d.Set("tcp_rst_timer", flattenObjectFirewallServiceCustomTcpRstTimer(o["tcp-rst-timer"], d, "tcp_rst_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-rst-timer"], "ObjectFirewallServiceCustom-TcpRstTimer"); ok {
			if err = d.Set("tcp_rst_timer", vv); err != nil {
				return fmt.Errorf("Error reading tcp_rst_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_rst_timer: %v", err)
		}
	}

	if err = d.Set("tcp_timewait_timer", flattenObjectFirewallServiceCustomTcpTimewaitTimer(o["tcp-timewait-timer"], d, "tcp_timewait_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-timewait-timer"], "ObjectFirewallServiceCustom-TcpTimewaitTimer"); ok {
			if err = d.Set("tcp_timewait_timer", vv); err != nil {
				return fmt.Errorf("Error reading tcp_timewait_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_timewait_timer: %v", err)
		}
	}

	if err = d.Set("udp_idle_timer", flattenObjectFirewallServiceCustomUdpIdleTimer(o["udp-idle-timer"], d, "udp_idle_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-idle-timer"], "ObjectFirewallServiceCustom-UdpIdleTimer"); ok {
			if err = d.Set("udp_idle_timer", vv); err != nil {
				return fmt.Errorf("Error reading udp_idle_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_idle_timer: %v", err)
		}
	}

	if err = d.Set("udp_portrange", flattenObjectFirewallServiceCustomUdpPortrange(o["udp-portrange"], d, "udp_portrange")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-portrange"], "ObjectFirewallServiceCustom-UdpPortrange"); ok {
			if err = d.Set("udp_portrange", vv); err != nil {
				return fmt.Errorf("Error reading udp_portrange: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_portrange: %v", err)
		}
	}

	if err = d.Set("visibility", flattenObjectFirewallServiceCustomVisibility(o["visibility"], d, "visibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["visibility"], "ObjectFirewallServiceCustom-Visibility"); ok {
			if err = d.Set("visibility", vv); err != nil {
				return fmt.Errorf("Error reading visibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallServiceCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallServiceCustomAppCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallServiceCustomAppServiceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallServiceCustomCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomCheckResetRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomHelper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomIcmpcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomIcmptype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomIprange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomProtocolNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomSctpPortrange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomSessionTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomTcpHalfcloseTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomTcpHalfopenTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomTcpPortrange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomTcpRstTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomTcpTimewaitTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomUdpIdleTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomUdpPortrange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallServiceCustomVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallServiceCustom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_category"); ok {
		t, err := expandObjectFirewallServiceCustomAppCategory(d, v, "app_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("app_service_type"); ok {
		t, err := expandObjectFirewallServiceCustomAppServiceType(d, v, "app_service_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-service-type"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok {
		t, err := expandObjectFirewallServiceCustomApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok {
		t, err := expandObjectFirewallServiceCustomCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("check_reset_range"); ok {
		t, err := expandObjectFirewallServiceCustomCheckResetRange(d, v, "check_reset_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-reset-range"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok {
		t, err := expandObjectFirewallServiceCustomColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectFirewallServiceCustomComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok {
		t, err := expandObjectFirewallServiceCustomFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok {
		t, err := expandObjectFirewallServiceCustomFqdn(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok {
		t, err := expandObjectFirewallServiceCustomGlobalObject(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("helper"); ok {
		t, err := expandObjectFirewallServiceCustomHelper(d, v, "helper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["helper"] = t
		}
	}

	if v, ok := d.GetOk("icmpcode"); ok {
		t, err := expandObjectFirewallServiceCustomIcmpcode(d, v, "icmpcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmpcode"] = t
		}
	}

	if v, ok := d.GetOk("icmptype"); ok {
		t, err := expandObjectFirewallServiceCustomIcmptype(d, v, "icmptype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmptype"] = t
		}
	}

	if v, ok := d.GetOk("iprange"); ok {
		t, err := expandObjectFirewallServiceCustomIprange(d, v, "iprange")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iprange"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectFirewallServiceCustomName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandObjectFirewallServiceCustomProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("protocol_number"); ok {
		t, err := expandObjectFirewallServiceCustomProtocolNumber(d, v, "protocol_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol-number"] = t
		}
	}

	if v, ok := d.GetOk("proxy"); ok {
		t, err := expandObjectFirewallServiceCustomProxy(d, v, "proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	if v, ok := d.GetOk("sctp_portrange"); ok {
		t, err := expandObjectFirewallServiceCustomSctpPortrange(d, v, "sctp_portrange")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-portrange"] = t
		}
	}

	if v, ok := d.GetOk("session_ttl"); ok {
		t, err := expandObjectFirewallServiceCustomSessionTtl(d, v, "session_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("tcp_halfclose_timer"); ok {
		t, err := expandObjectFirewallServiceCustomTcpHalfcloseTimer(d, v, "tcp_halfclose_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-halfclose-timer"] = t
		}
	}

	if v, ok := d.GetOk("tcp_halfopen_timer"); ok {
		t, err := expandObjectFirewallServiceCustomTcpHalfopenTimer(d, v, "tcp_halfopen_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-halfopen-timer"] = t
		}
	}

	if v, ok := d.GetOk("tcp_portrange"); ok {
		t, err := expandObjectFirewallServiceCustomTcpPortrange(d, v, "tcp_portrange")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-portrange"] = t
		}
	}

	if v, ok := d.GetOk("tcp_rst_timer"); ok {
		t, err := expandObjectFirewallServiceCustomTcpRstTimer(d, v, "tcp_rst_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-rst-timer"] = t
		}
	}

	if v, ok := d.GetOk("tcp_timewait_timer"); ok {
		t, err := expandObjectFirewallServiceCustomTcpTimewaitTimer(d, v, "tcp_timewait_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-timewait-timer"] = t
		}
	}

	if v, ok := d.GetOk("udp_idle_timer"); ok {
		t, err := expandObjectFirewallServiceCustomUdpIdleTimer(d, v, "udp_idle_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-idle-timer"] = t
		}
	}

	if v, ok := d.GetOk("udp_portrange"); ok {
		t, err := expandObjectFirewallServiceCustomUdpPortrange(d, v, "udp_portrange")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-portrange"] = t
		}
	}

	if v, ok := d.GetOk("visibility"); ok {
		t, err := expandObjectFirewallServiceCustomVisibility(d, v, "visibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	}

	return &obj, nil
}
