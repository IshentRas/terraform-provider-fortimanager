// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Email Filter profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectEmailfilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectEmailfilterProfileCreate,
		Read:   resourceObjectEmailfilterProfileRead,
		Update: resourceObjectEmailfilterProfileUpdate,
		Delete: resourceObjectEmailfilterProfileDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gmail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"imap": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_msg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mapi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"msn_hotmail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"other_webmails": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"pop3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_msg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smtp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hdrip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_msg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"spam_bal_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_bwl_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_bword_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_bword_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"spam_filtering": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_iptrust_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_log_fortiguard_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_mheader_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_rbl_table": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectEmailfilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectEmailfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectEmailfilterProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectEmailfilterProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectEmailfilterProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectEmailfilterProfileRead(d, m)
}

func resourceObjectEmailfilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectEmailfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectEmailfilterProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEmailfilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectEmailfilterProfileRead(d, m)
}

func resourceObjectEmailfilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectEmailfilterProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectEmailfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectEmailfilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectEmailfilterProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectEmailfilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEmailfilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectEmailfilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileGmail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectEmailfilterProfileGmailLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenObjectEmailfilterProfileGmailLogAll(i["log-all"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectEmailfilterProfileGmailLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileGmailLogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileImap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectEmailfilterProfileImapAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectEmailfilterProfileImapLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenObjectEmailfilterProfileImapLogAll(i["log-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenObjectEmailfilterProfileImapTagMsg(i["tag-msg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenObjectEmailfilterProfileImapTagType(i["tag-type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectEmailfilterProfileImapAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileImapLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileImapLogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileImapTagMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileImapTagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectEmailfilterProfileMapi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectEmailfilterProfileMapiAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectEmailfilterProfileMapiLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenObjectEmailfilterProfileMapiLogAll(i["log-all"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectEmailfilterProfileMapiAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileMapiLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileMapiLogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileMsnHotmail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectEmailfilterProfileMsnHotmailLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenObjectEmailfilterProfileMsnHotmailLogAll(i["log-all"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectEmailfilterProfileMsnHotmailLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileMsnHotmailLogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectEmailfilterProfileOtherWebmails(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenObjectEmailfilterProfileOtherWebmailsLogAll(i["log-all"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectEmailfilterProfileOtherWebmailsLogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfilePop3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectEmailfilterProfilePop3Action(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectEmailfilterProfilePop3Log(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenObjectEmailfilterProfilePop3LogAll(i["log-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenObjectEmailfilterProfilePop3TagMsg(i["tag-msg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenObjectEmailfilterProfilePop3TagType(i["tag-type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectEmailfilterProfilePop3Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfilePop3Log(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfilePop3LogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfilePop3TagMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfilePop3TagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectEmailfilterProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSmtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectEmailfilterProfileSmtpAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "hdrip"
	if _, ok := i["hdrip"]; ok {
		result["hdrip"] = flattenObjectEmailfilterProfileSmtpHdrip(i["hdrip"], d, pre_append)
	}

	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenObjectEmailfilterProfileSmtpLocalOverride(i["local-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectEmailfilterProfileSmtpLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenObjectEmailfilterProfileSmtpLogAll(i["log-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenObjectEmailfilterProfileSmtpTagMsg(i["tag-msg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenObjectEmailfilterProfileSmtpTagType(i["tag-type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectEmailfilterProfileSmtpAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSmtpHdrip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSmtpLocalOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSmtpLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSmtpLogAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSmtpTagMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSmtpTagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectEmailfilterProfileSpamBalTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSpamBwlTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSpamBwordTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSpamBwordThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSpamFiltering(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSpamIptrustTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSpamLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSpamLogFortiguardResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSpamMheaderTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEmailfilterProfileSpamRblTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectEmailfilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenObjectEmailfilterProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectEmailfilterProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("external", flattenObjectEmailfilterProfileExternal(o["external"], d, "external")); err != nil {
		if vv, ok := fortiAPIPatch(o["external"], "ObjectEmailfilterProfile-External"); ok {
			if err = d.Set("external", vv); err != nil {
				return fmt.Errorf("Error reading external: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if err = d.Set("feature_set", flattenObjectEmailfilterProfileFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["feature-set"], "ObjectEmailfilterProfile-FeatureSet"); ok {
			if err = d.Set("feature_set", vv); err != nil {
				return fmt.Errorf("Error reading feature_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("gmail", flattenObjectEmailfilterProfileGmail(o["gmail"], d, "gmail")); err != nil {
			if vv, ok := fortiAPIPatch(o["gmail"], "ObjectEmailfilterProfile-Gmail"); ok {
				if err = d.Set("gmail", vv); err != nil {
					return fmt.Errorf("Error reading gmail: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading gmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gmail"); ok {
			if err = d.Set("gmail", flattenObjectEmailfilterProfileGmail(o["gmail"], d, "gmail")); err != nil {
				if vv, ok := fortiAPIPatch(o["gmail"], "ObjectEmailfilterProfile-Gmail"); ok {
					if err = d.Set("gmail", vv); err != nil {
						return fmt.Errorf("Error reading gmail: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading gmail: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("imap", flattenObjectEmailfilterProfileImap(o["imap"], d, "imap")); err != nil {
			if vv, ok := fortiAPIPatch(o["imap"], "ObjectEmailfilterProfile-Imap"); ok {
				if err = d.Set("imap", vv); err != nil {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading imap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imap"); ok {
			if err = d.Set("imap", flattenObjectEmailfilterProfileImap(o["imap"], d, "imap")); err != nil {
				if vv, ok := fortiAPIPatch(o["imap"], "ObjectEmailfilterProfile-Imap"); ok {
					if err = d.Set("imap", vv); err != nil {
						return fmt.Errorf("Error reading imap: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mapi", flattenObjectEmailfilterProfileMapi(o["mapi"], d, "mapi")); err != nil {
			if vv, ok := fortiAPIPatch(o["mapi"], "ObjectEmailfilterProfile-Mapi"); ok {
				if err = d.Set("mapi", vv); err != nil {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mapi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mapi"); ok {
			if err = d.Set("mapi", flattenObjectEmailfilterProfileMapi(o["mapi"], d, "mapi")); err != nil {
				if vv, ok := fortiAPIPatch(o["mapi"], "ObjectEmailfilterProfile-Mapi"); ok {
					if err = d.Set("mapi", vv); err != nil {
						return fmt.Errorf("Error reading mapi: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("msn_hotmail", flattenObjectEmailfilterProfileMsnHotmail(o["msn-hotmail"], d, "msn_hotmail")); err != nil {
			if vv, ok := fortiAPIPatch(o["msn-hotmail"], "ObjectEmailfilterProfile-MsnHotmail"); ok {
				if err = d.Set("msn_hotmail", vv); err != nil {
					return fmt.Errorf("Error reading msn_hotmail: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading msn_hotmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("msn_hotmail"); ok {
			if err = d.Set("msn_hotmail", flattenObjectEmailfilterProfileMsnHotmail(o["msn-hotmail"], d, "msn_hotmail")); err != nil {
				if vv, ok := fortiAPIPatch(o["msn-hotmail"], "ObjectEmailfilterProfile-MsnHotmail"); ok {
					if err = d.Set("msn_hotmail", vv); err != nil {
						return fmt.Errorf("Error reading msn_hotmail: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading msn_hotmail: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectEmailfilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectEmailfilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("options", flattenObjectEmailfilterProfileOptions(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "ObjectEmailfilterProfile-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("other_webmails", flattenObjectEmailfilterProfileOtherWebmails(o["other-webmails"], d, "other_webmails")); err != nil {
			if vv, ok := fortiAPIPatch(o["other-webmails"], "ObjectEmailfilterProfile-OtherWebmails"); ok {
				if err = d.Set("other_webmails", vv); err != nil {
					return fmt.Errorf("Error reading other_webmails: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading other_webmails: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("other_webmails"); ok {
			if err = d.Set("other_webmails", flattenObjectEmailfilterProfileOtherWebmails(o["other-webmails"], d, "other_webmails")); err != nil {
				if vv, ok := fortiAPIPatch(o["other-webmails"], "ObjectEmailfilterProfile-OtherWebmails"); ok {
					if err = d.Set("other_webmails", vv); err != nil {
						return fmt.Errorf("Error reading other_webmails: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading other_webmails: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("pop3", flattenObjectEmailfilterProfilePop3(o["pop3"], d, "pop3")); err != nil {
			if vv, ok := fortiAPIPatch(o["pop3"], "ObjectEmailfilterProfile-Pop3"); ok {
				if err = d.Set("pop3", vv); err != nil {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pop3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3"); ok {
			if err = d.Set("pop3", flattenObjectEmailfilterProfilePop3(o["pop3"], d, "pop3")); err != nil {
				if vv, ok := fortiAPIPatch(o["pop3"], "ObjectEmailfilterProfile-Pop3"); ok {
					if err = d.Set("pop3", vv); err != nil {
						return fmt.Errorf("Error reading pop3: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			}
		}
	}

	if err = d.Set("replacemsg_group", flattenObjectEmailfilterProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "ObjectEmailfilterProfile-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("smtp", flattenObjectEmailfilterProfileSmtp(o["smtp"], d, "smtp")); err != nil {
			if vv, ok := fortiAPIPatch(o["smtp"], "ObjectEmailfilterProfile-Smtp"); ok {
				if err = d.Set("smtp", vv); err != nil {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading smtp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtp"); ok {
			if err = d.Set("smtp", flattenObjectEmailfilterProfileSmtp(o["smtp"], d, "smtp")); err != nil {
				if vv, ok := fortiAPIPatch(o["smtp"], "ObjectEmailfilterProfile-Smtp"); ok {
					if err = d.Set("smtp", vv); err != nil {
						return fmt.Errorf("Error reading smtp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			}
		}
	}

	if err = d.Set("spam_bal_table", flattenObjectEmailfilterProfileSpamBalTable(o["spam-bal-table"], d, "spam_bal_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-bal-table"], "ObjectEmailfilterProfile-SpamBalTable"); ok {
			if err = d.Set("spam_bal_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_bal_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_bal_table: %v", err)
		}
	}

	if err = d.Set("spam_bwl_table", flattenObjectEmailfilterProfileSpamBwlTable(o["spam-bwl-table"], d, "spam_bwl_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-bwl-table"], "ObjectEmailfilterProfile-SpamBwlTable"); ok {
			if err = d.Set("spam_bwl_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_bwl_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_bwl_table: %v", err)
		}
	}

	if err = d.Set("spam_bword_table", flattenObjectEmailfilterProfileSpamBwordTable(o["spam-bword-table"], d, "spam_bword_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-bword-table"], "ObjectEmailfilterProfile-SpamBwordTable"); ok {
			if err = d.Set("spam_bword_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_bword_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_bword_table: %v", err)
		}
	}

	if err = d.Set("spam_bword_threshold", flattenObjectEmailfilterProfileSpamBwordThreshold(o["spam-bword-threshold"], d, "spam_bword_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-bword-threshold"], "ObjectEmailfilterProfile-SpamBwordThreshold"); ok {
			if err = d.Set("spam_bword_threshold", vv); err != nil {
				return fmt.Errorf("Error reading spam_bword_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_bword_threshold: %v", err)
		}
	}

	if err = d.Set("spam_filtering", flattenObjectEmailfilterProfileSpamFiltering(o["spam-filtering"], d, "spam_filtering")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-filtering"], "ObjectEmailfilterProfile-SpamFiltering"); ok {
			if err = d.Set("spam_filtering", vv); err != nil {
				return fmt.Errorf("Error reading spam_filtering: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_filtering: %v", err)
		}
	}

	if err = d.Set("spam_iptrust_table", flattenObjectEmailfilterProfileSpamIptrustTable(o["spam-iptrust-table"], d, "spam_iptrust_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-iptrust-table"], "ObjectEmailfilterProfile-SpamIptrustTable"); ok {
			if err = d.Set("spam_iptrust_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_iptrust_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_iptrust_table: %v", err)
		}
	}

	if err = d.Set("spam_log", flattenObjectEmailfilterProfileSpamLog(o["spam-log"], d, "spam_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-log"], "ObjectEmailfilterProfile-SpamLog"); ok {
			if err = d.Set("spam_log", vv); err != nil {
				return fmt.Errorf("Error reading spam_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_log: %v", err)
		}
	}

	if err = d.Set("spam_log_fortiguard_response", flattenObjectEmailfilterProfileSpamLogFortiguardResponse(o["spam-log-fortiguard-response"], d, "spam_log_fortiguard_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-log-fortiguard-response"], "ObjectEmailfilterProfile-SpamLogFortiguardResponse"); ok {
			if err = d.Set("spam_log_fortiguard_response", vv); err != nil {
				return fmt.Errorf("Error reading spam_log_fortiguard_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_log_fortiguard_response: %v", err)
		}
	}

	if err = d.Set("spam_mheader_table", flattenObjectEmailfilterProfileSpamMheaderTable(o["spam-mheader-table"], d, "spam_mheader_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-mheader-table"], "ObjectEmailfilterProfile-SpamMheaderTable"); ok {
			if err = d.Set("spam_mheader_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_mheader_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_mheader_table: %v", err)
		}
	}

	if err = d.Set("spam_rbl_table", flattenObjectEmailfilterProfileSpamRblTable(o["spam-rbl-table"], d, "spam_rbl_table")); err != nil {
		if vv, ok := fortiAPIPatch(o["spam-rbl-table"], "ObjectEmailfilterProfile-SpamRblTable"); ok {
			if err = d.Set("spam_rbl_table", vv); err != nil {
				return fmt.Errorf("Error reading spam_rbl_table: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spam_rbl_table: %v", err)
		}
	}

	return nil
}

func flattenObjectEmailfilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectEmailfilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileFeatureSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileGmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandObjectEmailfilterProfileGmailLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandObjectEmailfilterProfileGmailLogAll(d, i["log_all"], pre_append)
	}

	return result, nil
}

func expandObjectEmailfilterProfileGmailLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileGmailLogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileImap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandObjectEmailfilterProfileImapAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandObjectEmailfilterProfileImapLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandObjectEmailfilterProfileImapLogAll(d, i["log_all"], pre_append)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-msg"], _ = expandObjectEmailfilterProfileImapTagMsg(d, i["tag_msg"], pre_append)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-type"], _ = expandObjectEmailfilterProfileImapTagType(d, i["tag_type"], pre_append)
	} else {
		result["tag-type"] = make([]string, 0)
	}

	return result, nil
}

func expandObjectEmailfilterProfileImapAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileImapLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileImapLogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileImapTagMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileImapTagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectEmailfilterProfileMapi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandObjectEmailfilterProfileMapiAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandObjectEmailfilterProfileMapiLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandObjectEmailfilterProfileMapiLogAll(d, i["log_all"], pre_append)
	}

	return result, nil
}

func expandObjectEmailfilterProfileMapiAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileMapiLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileMapiLogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileMsnHotmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandObjectEmailfilterProfileMsnHotmailLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandObjectEmailfilterProfileMsnHotmailLogAll(d, i["log_all"], pre_append)
	}

	return result, nil
}

func expandObjectEmailfilterProfileMsnHotmailLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileMsnHotmailLogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectEmailfilterProfileOtherWebmails(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandObjectEmailfilterProfileOtherWebmailsLogAll(d, i["log_all"], pre_append)
	}

	return result, nil
}

func expandObjectEmailfilterProfileOtherWebmailsLogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfilePop3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandObjectEmailfilterProfilePop3Action(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandObjectEmailfilterProfilePop3Log(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandObjectEmailfilterProfilePop3LogAll(d, i["log_all"], pre_append)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-msg"], _ = expandObjectEmailfilterProfilePop3TagMsg(d, i["tag_msg"], pre_append)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-type"], _ = expandObjectEmailfilterProfilePop3TagType(d, i["tag_type"], pre_append)
	} else {
		result["tag-type"] = make([]string, 0)
	}

	return result, nil
}

func expandObjectEmailfilterProfilePop3Action(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfilePop3Log(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfilePop3LogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfilePop3TagMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfilePop3TagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectEmailfilterProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSmtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandObjectEmailfilterProfileSmtpAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "hdrip"
	if _, ok := d.GetOk(pre_append); ok {
		result["hdrip"], _ = expandObjectEmailfilterProfileSmtpHdrip(d, i["hdrip"], pre_append)
	}
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok {
		result["local-override"], _ = expandObjectEmailfilterProfileSmtpLocalOverride(d, i["local_override"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandObjectEmailfilterProfileSmtpLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandObjectEmailfilterProfileSmtpLogAll(d, i["log_all"], pre_append)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-msg"], _ = expandObjectEmailfilterProfileSmtpTagMsg(d, i["tag_msg"], pre_append)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-type"], _ = expandObjectEmailfilterProfileSmtpTagType(d, i["tag_type"], pre_append)
	} else {
		result["tag-type"] = make([]string, 0)
	}

	return result, nil
}

func expandObjectEmailfilterProfileSmtpAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSmtpHdrip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSmtpLocalOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSmtpLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSmtpLogAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSmtpTagMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSmtpTagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectEmailfilterProfileSpamBalTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSpamBwlTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSpamBwordTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSpamBwordThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSpamFiltering(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSpamIptrustTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSpamLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSpamLogFortiguardResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSpamMheaderTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEmailfilterProfileSpamRblTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectEmailfilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectEmailfilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("external"); ok {
		t, err := expandObjectEmailfilterProfileExternal(d, v, "external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok {
		t, err := expandObjectEmailfilterProfileFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("gmail"); ok {
		t, err := expandObjectEmailfilterProfileGmail(d, v, "gmail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gmail"] = t
		}
	}

	if v, ok := d.GetOk("imap"); ok {
		t, err := expandObjectEmailfilterProfileImap(d, v, "imap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok {
		t, err := expandObjectEmailfilterProfileMapi(d, v, "mapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("msn_hotmail"); ok {
		t, err := expandObjectEmailfilterProfileMsnHotmail(d, v, "msn_hotmail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msn-hotmail"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectEmailfilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandObjectEmailfilterProfileOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("other_webmails"); ok {
		t, err := expandObjectEmailfilterProfileOtherWebmails(d, v, "other_webmails")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["other-webmails"] = t
		}
	}

	if v, ok := d.GetOk("pop3"); ok {
		t, err := expandObjectEmailfilterProfilePop3(d, v, "pop3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandObjectEmailfilterProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("smtp"); ok {
		t, err := expandObjectEmailfilterProfileSmtp(d, v, "smtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp"] = t
		}
	}

	if v, ok := d.GetOk("spam_bal_table"); ok {
		t, err := expandObjectEmailfilterProfileSpamBalTable(d, v, "spam_bal_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bal-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_bwl_table"); ok {
		t, err := expandObjectEmailfilterProfileSpamBwlTable(d, v, "spam_bwl_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bwl-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_bword_table"); ok {
		t, err := expandObjectEmailfilterProfileSpamBwordTable(d, v, "spam_bword_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bword-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_bword_threshold"); ok {
		t, err := expandObjectEmailfilterProfileSpamBwordThreshold(d, v, "spam_bword_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bword-threshold"] = t
		}
	}

	if v, ok := d.GetOk("spam_filtering"); ok {
		t, err := expandObjectEmailfilterProfileSpamFiltering(d, v, "spam_filtering")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-filtering"] = t
		}
	}

	if v, ok := d.GetOk("spam_iptrust_table"); ok {
		t, err := expandObjectEmailfilterProfileSpamIptrustTable(d, v, "spam_iptrust_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-iptrust-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_log"); ok {
		t, err := expandObjectEmailfilterProfileSpamLog(d, v, "spam_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-log"] = t
		}
	}

	if v, ok := d.GetOk("spam_log_fortiguard_response"); ok {
		t, err := expandObjectEmailfilterProfileSpamLogFortiguardResponse(d, v, "spam_log_fortiguard_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-log-fortiguard-response"] = t
		}
	}

	if v, ok := d.GetOk("spam_mheader_table"); ok {
		t, err := expandObjectEmailfilterProfileSpamMheaderTable(d, v, "spam_mheader_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-mheader-table"] = t
		}
	}

	if v, ok := d.GetOk("spam_rbl_table"); ok {
		t, err := expandObjectEmailfilterProfileSpamRblTable(d, v, "spam_rbl_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-rbl-table"] = t
		}
	}

	return &obj, nil
}
