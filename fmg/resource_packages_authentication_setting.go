// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure authentication setting.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesAuthenticationSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesAuthenticationSettingUpdate,
		Read:   resourcePackagesAuthenticationSettingRead,
		Update: resourcePackagesAuthenticationSettingUpdate,
		Delete: resourcePackagesAuthenticationSettingDelete,

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
			"active_auth_scheme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"captive_portal_ssl_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"captive_portal_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dev_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rewrite_https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sso_auth_scheme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_cert_ca": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePackagesAuthenticationSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesAuthenticationSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesAuthenticationSetting resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesAuthenticationSetting(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesAuthenticationSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("PackagesAuthenticationSetting")

	return resourcePackagesAuthenticationSettingRead(d, m)
}

func resourcePackagesAuthenticationSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeletePackagesAuthenticationSetting(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesAuthenticationSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesAuthenticationSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadPackagesAuthenticationSetting(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesAuthenticationSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesAuthenticationSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesAuthenticationSetting resource from API: %v", err)
	}
	return nil
}

func flattenPackagesAuthenticationSettingActiveAuthScheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationSettingAuthHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationSettingCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationSettingCaptivePortalIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationSettingCaptivePortalIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationSettingCaptivePortalPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationSettingCaptivePortalSslPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationSettingCaptivePortalType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationSettingCaptivePortal6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationSettingDevRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationSettingRewriteHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationSettingSsoAuthScheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesAuthenticationSettingUserCertCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesAuthenticationSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("active_auth_scheme", flattenPackagesAuthenticationSettingActiveAuthScheme(o["active-auth-scheme"], d, "active_auth_scheme")); err != nil {
		if vv, ok := fortiAPIPatch(o["active-auth-scheme"], "PackagesAuthenticationSetting-ActiveAuthScheme"); ok {
			if err = d.Set("active_auth_scheme", vv); err != nil {
				return fmt.Errorf("Error reading active_auth_scheme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading active_auth_scheme: %v", err)
		}
	}

	if err = d.Set("auth_https", flattenPackagesAuthenticationSettingAuthHttps(o["auth-https"], d, "auth_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-https"], "PackagesAuthenticationSetting-AuthHttps"); ok {
			if err = d.Set("auth_https", vv); err != nil {
				return fmt.Errorf("Error reading auth_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_https: %v", err)
		}
	}

	if err = d.Set("captive_portal", flattenPackagesAuthenticationSettingCaptivePortal(o["captive-portal"], d, "captive_portal")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal"], "PackagesAuthenticationSetting-CaptivePortal"); ok {
			if err = d.Set("captive_portal", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal: %v", err)
		}
	}

	if err = d.Set("captive_portal_ip", flattenPackagesAuthenticationSettingCaptivePortalIp(o["captive-portal-ip"], d, "captive_portal_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-ip"], "PackagesAuthenticationSetting-CaptivePortalIp"); ok {
			if err = d.Set("captive_portal_ip", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_ip: %v", err)
		}
	}

	if err = d.Set("captive_portal_ip6", flattenPackagesAuthenticationSettingCaptivePortalIp6(o["captive-portal-ip6"], d, "captive_portal_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-ip6"], "PackagesAuthenticationSetting-CaptivePortalIp6"); ok {
			if err = d.Set("captive_portal_ip6", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_ip6: %v", err)
		}
	}

	if err = d.Set("captive_portal_port", flattenPackagesAuthenticationSettingCaptivePortalPort(o["captive-portal-port"], d, "captive_portal_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-port"], "PackagesAuthenticationSetting-CaptivePortalPort"); ok {
			if err = d.Set("captive_portal_port", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_port: %v", err)
		}
	}

	if err = d.Set("captive_portal_ssl_port", flattenPackagesAuthenticationSettingCaptivePortalSslPort(o["captive-portal-ssl-port"], d, "captive_portal_ssl_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-ssl-port"], "PackagesAuthenticationSetting-CaptivePortalSslPort"); ok {
			if err = d.Set("captive_portal_ssl_port", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_ssl_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_ssl_port: %v", err)
		}
	}

	if err = d.Set("captive_portal_type", flattenPackagesAuthenticationSettingCaptivePortalType(o["captive-portal-type"], d, "captive_portal_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal-type"], "PackagesAuthenticationSetting-CaptivePortalType"); ok {
			if err = d.Set("captive_portal_type", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal_type: %v", err)
		}
	}

	if err = d.Set("captive_portal6", flattenPackagesAuthenticationSettingCaptivePortal6(o["captive-portal6"], d, "captive_portal6")); err != nil {
		if vv, ok := fortiAPIPatch(o["captive-portal6"], "PackagesAuthenticationSetting-CaptivePortal6"); ok {
			if err = d.Set("captive_portal6", vv); err != nil {
				return fmt.Errorf("Error reading captive_portal6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading captive_portal6: %v", err)
		}
	}

	if err = d.Set("dev_range", flattenPackagesAuthenticationSettingDevRange(o["dev-range"], d, "dev_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["dev-range"], "PackagesAuthenticationSetting-DevRange"); ok {
			if err = d.Set("dev_range", vv); err != nil {
				return fmt.Errorf("Error reading dev_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dev_range: %v", err)
		}
	}

	if err = d.Set("rewrite_https_port", flattenPackagesAuthenticationSettingRewriteHttpsPort(o["rewrite-https-port"], d, "rewrite_https_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["rewrite-https-port"], "PackagesAuthenticationSetting-RewriteHttpsPort"); ok {
			if err = d.Set("rewrite_https_port", vv); err != nil {
				return fmt.Errorf("Error reading rewrite_https_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rewrite_https_port: %v", err)
		}
	}

	if err = d.Set("sso_auth_scheme", flattenPackagesAuthenticationSettingSsoAuthScheme(o["sso-auth-scheme"], d, "sso_auth_scheme")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-auth-scheme"], "PackagesAuthenticationSetting-SsoAuthScheme"); ok {
			if err = d.Set("sso_auth_scheme", vv); err != nil {
				return fmt.Errorf("Error reading sso_auth_scheme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_auth_scheme: %v", err)
		}
	}

	if err = d.Set("user_cert_ca", flattenPackagesAuthenticationSettingUserCertCa(o["user-cert-ca"], d, "user_cert_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-cert-ca"], "PackagesAuthenticationSetting-UserCertCa"); ok {
			if err = d.Set("user_cert_ca", vv); err != nil {
				return fmt.Errorf("Error reading user_cert_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_cert_ca: %v", err)
		}
	}

	return nil
}

func flattenPackagesAuthenticationSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesAuthenticationSettingActiveAuthScheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationSettingAuthHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationSettingCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationSettingCaptivePortalIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationSettingCaptivePortalIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationSettingCaptivePortalPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationSettingCaptivePortalSslPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationSettingCaptivePortalType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationSettingCaptivePortal6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationSettingDevRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationSettingRewriteHttpsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationSettingSsoAuthScheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesAuthenticationSettingUserCertCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesAuthenticationSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("active_auth_scheme"); ok {
		t, err := expandPackagesAuthenticationSettingActiveAuthScheme(d, v, "active_auth_scheme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-auth-scheme"] = t
		}
	}

	if v, ok := d.GetOk("auth_https"); ok {
		t, err := expandPackagesAuthenticationSettingAuthHttps(d, v, "auth_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-https"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal"); ok {
		t, err := expandPackagesAuthenticationSettingCaptivePortal(d, v, "captive_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ip"); ok {
		t, err := expandPackagesAuthenticationSettingCaptivePortalIp(d, v, "captive_portal_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ip"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ip6"); ok {
		t, err := expandPackagesAuthenticationSettingCaptivePortalIp6(d, v, "captive_portal_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ip6"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_port"); ok {
		t, err := expandPackagesAuthenticationSettingCaptivePortalPort(d, v, "captive_portal_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-port"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ssl_port"); ok {
		t, err := expandPackagesAuthenticationSettingCaptivePortalSslPort(d, v, "captive_portal_ssl_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ssl-port"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_type"); ok {
		t, err := expandPackagesAuthenticationSettingCaptivePortalType(d, v, "captive_portal_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-type"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal6"); ok {
		t, err := expandPackagesAuthenticationSettingCaptivePortal6(d, v, "captive_portal6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal6"] = t
		}
	}

	if v, ok := d.GetOk("dev_range"); ok {
		t, err := expandPackagesAuthenticationSettingDevRange(d, v, "dev_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-range"] = t
		}
	}

	if v, ok := d.GetOk("rewrite_https_port"); ok {
		t, err := expandPackagesAuthenticationSettingRewriteHttpsPort(d, v, "rewrite_https_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rewrite-https-port"] = t
		}
	}

	if v, ok := d.GetOk("sso_auth_scheme"); ok {
		t, err := expandPackagesAuthenticationSettingSsoAuthScheme(d, v, "sso_auth_scheme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-auth-scheme"] = t
		}
	}

	if v, ok := d.GetOk("user_cert_ca"); ok {
		t, err := expandPackagesAuthenticationSettingUserCertCa(d, v, "user_cert_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-cert-ca"] = t
		}
	}

	return &obj, nil
}
