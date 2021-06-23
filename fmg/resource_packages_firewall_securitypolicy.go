// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NGFW IPv4/IPv6 application policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesFirewallSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallSecurityPolicyCreate,
		Read:   resourcePackagesFirewallSecurityPolicyRead,
		Update: resourcePackagesFirewallSecurityPolicyUpdate,
		Delete: resourcePackagesFirewallSecurityPolicyDelete,

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
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_group": &schema.Schema{
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
			"application_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cifs_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dlp_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enforce_default_app_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_filter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsso_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icap_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_custom_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_custom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_custom_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"learning_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mms_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"profile_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_deny_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_filter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"videofilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voip_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePackagesFirewallSecurityPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesFirewallSecurityPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallSecurityPolicy resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesFirewallSecurityPolicy(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallSecurityPolicy resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallSecurityPolicyRead(d, m)
}

func resourcePackagesFirewallSecurityPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesFirewallSecurityPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallSecurityPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesFirewallSecurityPolicy(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallSecurityPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallSecurityPolicyRead(d, m)
}

func resourcePackagesFirewallSecurityPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	err = c.DeletePackagesFirewallSecurityPolicy(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesFirewallSecurityPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallSecurityPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	if pkg == "" {
		pkg = importOptionChecking(m.(*FortiClient).Cfg, "pkg")
		if err = d.Set("pkg", pkg); err != nil {
			return fmt.Errorf("Error set params pkg: %v", err)
		}
	}
	var paralist []string
	paralist = append(paralist, pkg)

	o, err := c.ReadPackagesFirewallSecurityPolicy(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallSecurityPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallSecurityPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallSecurityPolicy resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallSecurityPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyAppCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyAppGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyApplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenPackagesFirewallSecurityPolicyApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyAvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyCifsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyDlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyDnsfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyDstaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyDstaddr4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyEnforceDefaultAppPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyFileFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyFssoGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyGlobalLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyIcapProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceSrcCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceSrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceSrcGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceSrcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceSrcId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyInternetServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyLearningMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyLogtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyLogtrafficStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyMmsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyProfileGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicySendDenyPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicySrcaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicySrcaddr4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicySrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicySshFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicySslSshProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyUrlCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyUtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyVideofilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyVoipProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallSecurityPolicyWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesFirewallSecurityPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenPackagesFirewallSecurityPolicyAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "PackagesFirewallSecurityPolicy-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("app_category", flattenPackagesFirewallSecurityPolicyAppCategory(o["app-category"], d, "app_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-category"], "PackagesFirewallSecurityPolicy-AppCategory"); ok {
			if err = d.Set("app_category", vv); err != nil {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_category: %v", err)
		}
	}

	if err = d.Set("app_group", flattenPackagesFirewallSecurityPolicyAppGroup(o["app-group"], d, "app_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-group"], "PackagesFirewallSecurityPolicy-AppGroup"); ok {
			if err = d.Set("app_group", vv); err != nil {
				return fmt.Errorf("Error reading app_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_group: %v", err)
		}
	}

	if err = d.Set("application", flattenPackagesFirewallSecurityPolicyApplication(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "PackagesFirewallSecurityPolicy-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("application_list", flattenPackagesFirewallSecurityPolicyApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "PackagesFirewallSecurityPolicy-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenPackagesFirewallSecurityPolicyAvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile"], "PackagesFirewallSecurityPolicy-AvProfile"); ok {
			if err = d.Set("av_profile", vv); err != nil {
				return fmt.Errorf("Error reading av_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("cifs_profile", flattenPackagesFirewallSecurityPolicyCifsProfile(o["cifs-profile"], d, "cifs_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["cifs-profile"], "PackagesFirewallSecurityPolicy-CifsProfile"); ok {
			if err = d.Set("cifs_profile", vv); err != nil {
				return fmt.Errorf("Error reading cifs_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("comments", flattenPackagesFirewallSecurityPolicyComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesFirewallSecurityPolicy-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenPackagesFirewallSecurityPolicyDlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor"], "PackagesFirewallSecurityPolicy-DlpSensor"); ok {
			if err = d.Set("dlp_sensor", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenPackagesFirewallSecurityPolicyDnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnsfilter-profile"], "PackagesFirewallSecurityPolicy-DnsfilterProfile"); ok {
			if err = d.Set("dnsfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenPackagesFirewallSecurityPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "PackagesFirewallSecurityPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenPackagesFirewallSecurityPolicyDstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr-negate"], "PackagesFirewallSecurityPolicy-DstaddrNegate"); ok {
			if err = d.Set("dstaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr4", flattenPackagesFirewallSecurityPolicyDstaddr4(o["dstaddr4"], d, "dstaddr4")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr4"], "PackagesFirewallSecurityPolicy-Dstaddr4"); ok {
			if err = d.Set("dstaddr4", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr4: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenPackagesFirewallSecurityPolicyDstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "PackagesFirewallSecurityPolicy-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenPackagesFirewallSecurityPolicyDstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "PackagesFirewallSecurityPolicy-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenPackagesFirewallSecurityPolicyEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile"], "PackagesFirewallSecurityPolicy-EmailfilterProfile"); ok {
			if err = d.Set("emailfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("enforce_default_app_port", flattenPackagesFirewallSecurityPolicyEnforceDefaultAppPort(o["enforce-default-app-port"], d, "enforce_default_app_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-default-app-port"], "PackagesFirewallSecurityPolicy-EnforceDefaultAppPort"); ok {
			if err = d.Set("enforce_default_app_port", vv); err != nil {
				return fmt.Errorf("Error reading enforce_default_app_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_default_app_port: %v", err)
		}
	}

	if err = d.Set("file_filter_profile", flattenPackagesFirewallSecurityPolicyFileFilterProfile(o["file-filter-profile"], d, "file_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-filter-profile"], "PackagesFirewallSecurityPolicy-FileFilterProfile"); ok {
			if err = d.Set("file_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading file_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_filter_profile: %v", err)
		}
	}

	if err = d.Set("fsso_groups", flattenPackagesFirewallSecurityPolicyFssoGroups(o["fsso-groups"], d, "fsso_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-groups"], "PackagesFirewallSecurityPolicy-FssoGroups"); ok {
			if err = d.Set("fsso_groups", vv); err != nil {
				return fmt.Errorf("Error reading fsso_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_groups: %v", err)
		}
	}

	if err = d.Set("global_label", flattenPackagesFirewallSecurityPolicyGlobalLabel(o["global-label"], d, "global_label")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-label"], "PackagesFirewallSecurityPolicy-GlobalLabel"); ok {
			if err = d.Set("global_label", vv); err != nil {
				return fmt.Errorf("Error reading global_label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_label: %v", err)
		}
	}

	if err = d.Set("groups", flattenPackagesFirewallSecurityPolicyGroups(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "PackagesFirewallSecurityPolicy-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenPackagesFirewallSecurityPolicyIcapProfile(o["icap-profile"], d, "icap_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["icap-profile"], "PackagesFirewallSecurityPolicy-IcapProfile"); ok {
			if err = d.Set("icap_profile", vv); err != nil {
				return fmt.Errorf("Error reading icap_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenPackagesFirewallSecurityPolicyInternetService(o["internet-service"], d, "internet_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service"], "PackagesFirewallSecurityPolicy-InternetService"); ok {
			if err = d.Set("internet_service", vv); err != nil {
				return fmt.Errorf("Error reading internet_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", flattenPackagesFirewallSecurityPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom"], "PackagesFirewallSecurityPolicy-InternetServiceCustom"); ok {
			if err = d.Set("internet_service_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_custom_group", flattenPackagesFirewallSecurityPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-custom-group"], "PackagesFirewallSecurityPolicy-InternetServiceCustomGroup"); ok {
			if err = d.Set("internet_service_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_group", flattenPackagesFirewallSecurityPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-group"], "PackagesFirewallSecurityPolicy-InternetServiceGroup"); ok {
			if err = d.Set("internet_service_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_group: %v", err)
		}
	}

	if err = d.Set("internet_service_name", flattenPackagesFirewallSecurityPolicyInternetServiceName(o["internet-service-name"], d, "internet_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-name"], "PackagesFirewallSecurityPolicy-InternetServiceName"); ok {
			if err = d.Set("internet_service_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_name: %v", err)
		}
	}

	if err = d.Set("internet_service_id", flattenPackagesFirewallSecurityPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-id"], "PackagesFirewallSecurityPolicy-InternetServiceId"); ok {
			if err = d.Set("internet_service_id", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_id: %v", err)
		}
	}

	if err = d.Set("internet_service_negate", flattenPackagesFirewallSecurityPolicyInternetServiceNegate(o["internet-service-negate"], d, "internet_service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-negate"], "PackagesFirewallSecurityPolicy-InternetServiceNegate"); ok {
			if err = d.Set("internet_service_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service_src", flattenPackagesFirewallSecurityPolicyInternetServiceSrc(o["internet-service-src"], d, "internet_service_src")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src"], "PackagesFirewallSecurityPolicy-InternetServiceSrc"); ok {
			if err = d.Set("internet_service_src", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom", flattenPackagesFirewallSecurityPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom"], "PackagesFirewallSecurityPolicy-InternetServiceSrcCustom"); ok {
			if err = d.Set("internet_service_src_custom", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom_group", flattenPackagesFirewallSecurityPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-custom-group"], "PackagesFirewallSecurityPolicy-InternetServiceSrcCustomGroup"); ok {
			if err = d.Set("internet_service_src_custom_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_group", flattenPackagesFirewallSecurityPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-group"], "PackagesFirewallSecurityPolicy-InternetServiceSrcGroup"); ok {
			if err = d.Set("internet_service_src_group", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_name", flattenPackagesFirewallSecurityPolicyInternetServiceSrcName(o["internet-service-src-name"], d, "internet_service_src_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-name"], "PackagesFirewallSecurityPolicy-InternetServiceSrcName"); ok {
			if err = d.Set("internet_service_src_name", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_name: %v", err)
		}
	}

	if err = d.Set("internet_service_src_id", flattenPackagesFirewallSecurityPolicyInternetServiceSrcId(o["internet-service-src-id"], d, "internet_service_src_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-id"], "PackagesFirewallSecurityPolicy-InternetServiceSrcId"); ok {
			if err = d.Set("internet_service_src_id", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_id: %v", err)
		}
	}

	if err = d.Set("internet_service_src_negate", flattenPackagesFirewallSecurityPolicyInternetServiceSrcNegate(o["internet-service-src-negate"], d, "internet_service_src_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-src-negate"], "PackagesFirewallSecurityPolicy-InternetServiceSrcNegate"); ok {
			if err = d.Set("internet_service_src_negate", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenPackagesFirewallSecurityPolicyIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "PackagesFirewallSecurityPolicy-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("learning_mode", flattenPackagesFirewallSecurityPolicyLearningMode(o["learning-mode"], d, "learning_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["learning-mode"], "PackagesFirewallSecurityPolicy-LearningMode"); ok {
			if err = d.Set("learning_mode", vv); err != nil {
				return fmt.Errorf("Error reading learning_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learning_mode: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenPackagesFirewallSecurityPolicyLogtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "PackagesFirewallSecurityPolicy-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("logtraffic_start", flattenPackagesFirewallSecurityPolicyLogtrafficStart(o["logtraffic-start"], d, "logtraffic_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic-start"], "PackagesFirewallSecurityPolicy-LogtrafficStart"); ok {
			if err = d.Set("logtraffic_start", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic_start: %v", err)
		}
	}

	if err = d.Set("mms_profile", flattenPackagesFirewallSecurityPolicyMmsProfile(o["mms-profile"], d, "mms_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["mms-profile"], "PackagesFirewallSecurityPolicy-MmsProfile"); ok {
			if err = d.Set("mms_profile", vv); err != nil {
				return fmt.Errorf("Error reading mms_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mms_profile: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesFirewallSecurityPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesFirewallSecurityPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("policyid", flattenPackagesFirewallSecurityPolicyPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "PackagesFirewallSecurityPolicy-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("profile_group", flattenPackagesFirewallSecurityPolicyProfileGroup(o["profile-group"], d, "profile_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-group"], "PackagesFirewallSecurityPolicy-ProfileGroup"); ok {
			if err = d.Set("profile_group", vv); err != nil {
				return fmt.Errorf("Error reading profile_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenPackagesFirewallSecurityPolicyProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-protocol-options"], "PackagesFirewallSecurityPolicy-ProfileProtocolOptions"); ok {
			if err = d.Set("profile_protocol_options", vv); err != nil {
				return fmt.Errorf("Error reading profile_protocol_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenPackagesFirewallSecurityPolicyProfileType(o["profile-type"], d, "profile_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-type"], "PackagesFirewallSecurityPolicy-ProfileType"); ok {
			if err = d.Set("profile_type", vv); err != nil {
				return fmt.Errorf("Error reading profile_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	if err = d.Set("schedule", flattenPackagesFirewallSecurityPolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "PackagesFirewallSecurityPolicy-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("send_deny_packet", flattenPackagesFirewallSecurityPolicySendDenyPacket(o["send-deny-packet"], d, "send_deny_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-deny-packet"], "PackagesFirewallSecurityPolicy-SendDenyPacket"); ok {
			if err = d.Set("send_deny_packet", vv); err != nil {
				return fmt.Errorf("Error reading send_deny_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_deny_packet: %v", err)
		}
	}

	if err = d.Set("service", flattenPackagesFirewallSecurityPolicyService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "PackagesFirewallSecurityPolicy-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenPackagesFirewallSecurityPolicyServiceNegate(o["service-negate"], d, "service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-negate"], "PackagesFirewallSecurityPolicy-ServiceNegate"); ok {
			if err = d.Set("service_negate", vv); err != nil {
				return fmt.Errorf("Error reading service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenPackagesFirewallSecurityPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "PackagesFirewallSecurityPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenPackagesFirewallSecurityPolicySrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr-negate"], "PackagesFirewallSecurityPolicy-SrcaddrNegate"); ok {
			if err = d.Set("srcaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr4", flattenPackagesFirewallSecurityPolicySrcaddr4(o["srcaddr4"], d, "srcaddr4")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr4"], "PackagesFirewallSecurityPolicy-Srcaddr4"); ok {
			if err = d.Set("srcaddr4", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr4: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenPackagesFirewallSecurityPolicySrcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "PackagesFirewallSecurityPolicy-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenPackagesFirewallSecurityPolicySrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "PackagesFirewallSecurityPolicy-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenPackagesFirewallSecurityPolicySshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-filter-profile"], "PackagesFirewallSecurityPolicy-SshFilterProfile"); ok {
			if err = d.Set("ssh_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenPackagesFirewallSecurityPolicySslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-ssh-profile"], "PackagesFirewallSecurityPolicy-SslSshProfile"); ok {
			if err = d.Set("ssl_ssh_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesFirewallSecurityPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesFirewallSecurityPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("url_category", flattenPackagesFirewallSecurityPolicyUrlCategory(o["url-category"], d, "url_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-category"], "PackagesFirewallSecurityPolicy-UrlCategory"); ok {
			if err = d.Set("url_category", vv); err != nil {
				return fmt.Errorf("Error reading url_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_category: %v", err)
		}
	}

	if err = d.Set("users", flattenPackagesFirewallSecurityPolicyUsers(o["users"], d, "users")); err != nil {
		if vv, ok := fortiAPIPatch(o["users"], "PackagesFirewallSecurityPolicy-Users"); ok {
			if err = d.Set("users", vv); err != nil {
				return fmt.Errorf("Error reading users: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenPackagesFirewallSecurityPolicyUtmStatus(o["utm-status"], d, "utm_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["utm-status"], "PackagesFirewallSecurityPolicy-UtmStatus"); ok {
			if err = d.Set("utm_status", vv); err != nil {
				return fmt.Errorf("Error reading utm_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenPackagesFirewallSecurityPolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "PackagesFirewallSecurityPolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("videofilter_profile", flattenPackagesFirewallSecurityPolicyVideofilterProfile(o["videofilter-profile"], d, "videofilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["videofilter-profile"], "PackagesFirewallSecurityPolicy-VideofilterProfile"); ok {
			if err = d.Set("videofilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading videofilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading videofilter_profile: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenPackagesFirewallSecurityPolicyVoipProfile(o["voip-profile"], d, "voip_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip-profile"], "PackagesFirewallSecurityPolicy-VoipProfile"); ok {
			if err = d.Set("voip_profile", vv); err != nil {
				return fmt.Errorf("Error reading voip_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenPackagesFirewallSecurityPolicyWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "PackagesFirewallSecurityPolicy-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallSecurityPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallSecurityPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyAppCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyAppGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandPackagesFirewallSecurityPolicyApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyAvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyCifsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyDlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyDnsfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyDstaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyDstaddr4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyEmailfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyEnforceDefaultAppPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyFileFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyFssoGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyGlobalLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyIcapProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceSrcId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyInternetServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyLearningMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyLogtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyLogtrafficStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyMmsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyProfileGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicySchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicySendDenyPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicySrcaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicySrcaddr4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicySrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicySshFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicySslSshProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyUrlCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyUtmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyVideofilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyVoipProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallSecurityPolicyWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallSecurityPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok {
		t, err := expandPackagesFirewallSecurityPolicyAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("app_category"); ok {
		t, err := expandPackagesFirewallSecurityPolicyAppCategory(d, v, "app_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("app_group"); ok {
		t, err := expandPackagesFirewallSecurityPolicyAppGroup(d, v, "app_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-group"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok {
		t, err := expandPackagesFirewallSecurityPolicyApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok {
		t, err := expandPackagesFirewallSecurityPolicyApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok {
		t, err := expandPackagesFirewallSecurityPolicyAvProfile(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("cifs_profile"); ok {
		t, err := expandPackagesFirewallSecurityPolicyCifsProfile(d, v, "cifs_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs-profile"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandPackagesFirewallSecurityPolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok {
		t, err := expandPackagesFirewallSecurityPolicyDlpSensor(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok {
		t, err := expandPackagesFirewallSecurityPolicyDnsfilterProfile(d, v, "dnsfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {
		t, err := expandPackagesFirewallSecurityPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok {
		t, err := expandPackagesFirewallSecurityPolicyDstaddrNegate(d, v, "dstaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr4"); ok {
		t, err := expandPackagesFirewallSecurityPolicyDstaddr4(d, v, "dstaddr4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr4"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok {
		t, err := expandPackagesFirewallSecurityPolicyDstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok {
		t, err := expandPackagesFirewallSecurityPolicyDstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok {
		t, err := expandPackagesFirewallSecurityPolicyEmailfilterProfile(d, v, "emailfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("enforce_default_app_port"); ok {
		t, err := expandPackagesFirewallSecurityPolicyEnforceDefaultAppPort(d, v, "enforce_default_app_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-default-app-port"] = t
		}
	}

	if v, ok := d.GetOk("file_filter_profile"); ok {
		t, err := expandPackagesFirewallSecurityPolicyFileFilterProfile(d, v, "file_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("fsso_groups"); ok {
		t, err := expandPackagesFirewallSecurityPolicyFssoGroups(d, v, "fsso_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-groups"] = t
		}
	}

	if v, ok := d.GetOk("global_label"); ok {
		t, err := expandPackagesFirewallSecurityPolicyGlobalLabel(d, v, "global_label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-label"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok {
		t, err := expandPackagesFirewallSecurityPolicyGroups(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("icap_profile"); ok {
		t, err := expandPackagesFirewallSecurityPolicyIcapProfile(d, v, "icap_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetService(d, v, "internet_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceCustom(d, v, "internet_service_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceGroup(d, v, "internet_service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceName(d, v, "internet_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_id"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceId(d, v, "internet_service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_negate"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceNegate(d, v, "internet_service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceSrc(d, v, "internet_service_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom_group"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_group"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_name"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceSrcName(d, v, "internet_service_src_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_id"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceSrcId(d, v, "internet_service_src_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_negate"); ok {
		t, err := expandPackagesFirewallSecurityPolicyInternetServiceSrcNegate(d, v, "internet_service_src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok {
		t, err := expandPackagesFirewallSecurityPolicyIpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("learning_mode"); ok {
		t, err := expandPackagesFirewallSecurityPolicyLearningMode(d, v, "learning_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learning-mode"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok {
		t, err := expandPackagesFirewallSecurityPolicyLogtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic_start"); ok {
		t, err := expandPackagesFirewallSecurityPolicyLogtrafficStart(d, v, "logtraffic_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic-start"] = t
		}
	}

	if v, ok := d.GetOk("mms_profile"); ok {
		t, err := expandPackagesFirewallSecurityPolicyMmsProfile(d, v, "mms_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mms-profile"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandPackagesFirewallSecurityPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok {
		t, err := expandPackagesFirewallSecurityPolicyPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("profile_group"); ok {
		t, err := expandPackagesFirewallSecurityPolicyProfileGroup(d, v, "profile_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-group"] = t
		}
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok {
		t, err := expandPackagesFirewallSecurityPolicyProfileProtocolOptions(d, v, "profile_protocol_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("profile_type"); ok {
		t, err := expandPackagesFirewallSecurityPolicyProfileType(d, v, "profile_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandPackagesFirewallSecurityPolicySchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("send_deny_packet"); ok {
		t, err := expandPackagesFirewallSecurityPolicySendDenyPacket(d, v, "send_deny_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-deny-packet"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {
		t, err := expandPackagesFirewallSecurityPolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok {
		t, err := expandPackagesFirewallSecurityPolicyServiceNegate(d, v, "service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {
		t, err := expandPackagesFirewallSecurityPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok {
		t, err := expandPackagesFirewallSecurityPolicySrcaddrNegate(d, v, "srcaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr4"); ok {
		t, err := expandPackagesFirewallSecurityPolicySrcaddr4(d, v, "srcaddr4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr4"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok {
		t, err := expandPackagesFirewallSecurityPolicySrcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok {
		t, err := expandPackagesFirewallSecurityPolicySrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok {
		t, err := expandPackagesFirewallSecurityPolicySshFilterProfile(d, v, "ssh_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok {
		t, err := expandPackagesFirewallSecurityPolicySslSshProfile(d, v, "ssl_ssh_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandPackagesFirewallSecurityPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("url_category"); ok {
		t, err := expandPackagesFirewallSecurityPolicyUrlCategory(d, v, "url_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-category"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok {
		t, err := expandPackagesFirewallSecurityPolicyUsers(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok {
		t, err := expandPackagesFirewallSecurityPolicyUtmStatus(d, v, "utm_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandPackagesFirewallSecurityPolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("videofilter_profile"); ok {
		t, err := expandPackagesFirewallSecurityPolicyVideofilterProfile(d, v, "videofilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["videofilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("voip_profile"); ok {
		t, err := expandPackagesFirewallSecurityPolicyVoipProfile(d, v, "voip_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok {
		t, err := expandPackagesFirewallSecurityPolicyWebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	return &obj, nil
}
