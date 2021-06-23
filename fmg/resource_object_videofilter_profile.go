// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure VideoFilter profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectVideofilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVideofilterProfileCreate,
		Read:   resourceObjectVideofilterProfileRead,
		Update: resourceObjectVideofilterProfileUpdate,
		Delete: resourceObjectVideofilterProfileDelete,

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
			"dailymotion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiguard_category": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"category_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
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
			"vimeo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vimeo_restrict": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"youtube": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"youtube_channel_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"youtube_restrict": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectVideofilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVideofilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVideofilterProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVideofilterProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVideofilterProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVideofilterProfileRead(d, m)
}

func resourceObjectVideofilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVideofilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVideofilterProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVideofilterProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVideofilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVideofilterProfileRead(d, m)
}

func resourceObjectVideofilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectVideofilterProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVideofilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVideofilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectVideofilterProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVideofilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVideofilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVideofilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectVideofilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileDailymotion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileFortiguardCategory(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenObjectVideofilterProfileFortiguardCategoryFilters(i["filters"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectVideofilterProfileFortiguardCategoryFilters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenObjectVideofilterProfileFortiguardCategoryFiltersAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectVideofilterProfileFortiguardCategory-Filters-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category_id"
		if _, ok := i["category-id"]; ok {
			v := flattenObjectVideofilterProfileFortiguardCategoryFiltersCategoryId(i["category-id"], d, pre_append)
			tmp["category_id"] = fortiAPISubPartPatch(v, "ObjectVideofilterProfileFortiguardCategory-Filters-CategoryId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectVideofilterProfileFortiguardCategoryFiltersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectVideofilterProfileFortiguardCategory-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectVideofilterProfileFortiguardCategoryFiltersLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectVideofilterProfileFortiguardCategory-Filters-Log")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersCategoryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileVimeo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileVimeoRestrict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileYoutube(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileYoutubeChannelFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileYoutubeRestrict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVideofilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenObjectVideofilterProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectVideofilterProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("dailymotion", flattenObjectVideofilterProfileDailymotion(o["dailymotion"], d, "dailymotion")); err != nil {
		if vv, ok := fortiAPIPatch(o["dailymotion"], "ObjectVideofilterProfile-Dailymotion"); ok {
			if err = d.Set("dailymotion", vv); err != nil {
				return fmt.Errorf("Error reading dailymotion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dailymotion: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fortiguard_category", flattenObjectVideofilterProfileFortiguardCategory(o["fortiguard-category"], d, "fortiguard_category")); err != nil {
			if vv, ok := fortiAPIPatch(o["fortiguard-category"], "ObjectVideofilterProfile-FortiguardCategory"); ok {
				if err = d.Set("fortiguard_category", vv); err != nil {
					return fmt.Errorf("Error reading fortiguard_category: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fortiguard_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fortiguard_category"); ok {
			if err = d.Set("fortiguard_category", flattenObjectVideofilterProfileFortiguardCategory(o["fortiguard-category"], d, "fortiguard_category")); err != nil {
				if vv, ok := fortiAPIPatch(o["fortiguard-category"], "ObjectVideofilterProfile-FortiguardCategory"); ok {
					if err = d.Set("fortiguard_category", vv); err != nil {
						return fmt.Errorf("Error reading fortiguard_category: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fortiguard_category: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectVideofilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVideofilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vimeo", flattenObjectVideofilterProfileVimeo(o["vimeo"], d, "vimeo")); err != nil {
		if vv, ok := fortiAPIPatch(o["vimeo"], "ObjectVideofilterProfile-Vimeo"); ok {
			if err = d.Set("vimeo", vv); err != nil {
				return fmt.Errorf("Error reading vimeo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vimeo: %v", err)
		}
	}

	if err = d.Set("vimeo_restrict", flattenObjectVideofilterProfileVimeoRestrict(o["vimeo-restrict"], d, "vimeo_restrict")); err != nil {
		if vv, ok := fortiAPIPatch(o["vimeo-restrict"], "ObjectVideofilterProfile-VimeoRestrict"); ok {
			if err = d.Set("vimeo_restrict", vv); err != nil {
				return fmt.Errorf("Error reading vimeo_restrict: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vimeo_restrict: %v", err)
		}
	}

	if err = d.Set("youtube", flattenObjectVideofilterProfileYoutube(o["youtube"], d, "youtube")); err != nil {
		if vv, ok := fortiAPIPatch(o["youtube"], "ObjectVideofilterProfile-Youtube"); ok {
			if err = d.Set("youtube", vv); err != nil {
				return fmt.Errorf("Error reading youtube: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading youtube: %v", err)
		}
	}

	if err = d.Set("youtube_channel_filter", flattenObjectVideofilterProfileYoutubeChannelFilter(o["youtube-channel-filter"], d, "youtube_channel_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["youtube-channel-filter"], "ObjectVideofilterProfile-YoutubeChannelFilter"); ok {
			if err = d.Set("youtube_channel_filter", vv); err != nil {
				return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
		}
	}

	if err = d.Set("youtube_restrict", flattenObjectVideofilterProfileYoutubeRestrict(o["youtube-restrict"], d, "youtube_restrict")); err != nil {
		if vv, ok := fortiAPIPatch(o["youtube-restrict"], "ObjectVideofilterProfile-YoutubeRestrict"); ok {
			if err = d.Set("youtube_restrict", vv); err != nil {
				return fmt.Errorf("Error reading youtube_restrict: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading youtube_restrict: %v", err)
		}
	}

	return nil
}

func flattenObjectVideofilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVideofilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileDailymotion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileFortiguardCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok {
		result["filters"], _ = expandObjectVideofilterProfileFortiguardCategoryFilters(d, i["filters"], pre_append)
	} else {
		result["filters"] = make([]string, 0)
	}

	return result, nil
}

func expandObjectVideofilterProfileFortiguardCategoryFilters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandObjectVideofilterProfileFortiguardCategoryFiltersAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category-id"], _ = expandObjectVideofilterProfileFortiguardCategoryFiltersCategoryId(d, i["category_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectVideofilterProfileFortiguardCategoryFiltersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandObjectVideofilterProfileFortiguardCategoryFiltersLog(d, i["log"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectVideofilterProfileFortiguardCategoryFiltersAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileFortiguardCategoryFiltersCategoryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileFortiguardCategoryFiltersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileFortiguardCategoryFiltersLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileVimeo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileVimeoRestrict(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileYoutube(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileYoutubeChannelFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileYoutubeRestrict(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVideofilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectVideofilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dailymotion"); ok {
		t, err := expandObjectVideofilterProfileDailymotion(d, v, "dailymotion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dailymotion"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_category"); ok {
		t, err := expandObjectVideofilterProfileFortiguardCategory(d, v, "fortiguard_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-category"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectVideofilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vimeo"); ok {
		t, err := expandObjectVideofilterProfileVimeo(d, v, "vimeo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vimeo"] = t
		}
	}

	if v, ok := d.GetOk("vimeo_restrict"); ok {
		t, err := expandObjectVideofilterProfileVimeoRestrict(d, v, "vimeo_restrict")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vimeo-restrict"] = t
		}
	}

	if v, ok := d.GetOk("youtube"); ok {
		t, err := expandObjectVideofilterProfileYoutube(d, v, "youtube")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube"] = t
		}
	}

	if v, ok := d.GetOk("youtube_channel_filter"); ok {
		t, err := expandObjectVideofilterProfileYoutubeChannelFilter(d, v, "youtube_channel_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-channel-filter"] = t
		}
	}

	if v, ok := d.GetOk("youtube_restrict"); ok {
		t, err := expandObjectVideofilterProfileYoutubeRestrict(d, v, "youtube_restrict")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-restrict"] = t
		}
	}

	return &obj, nil
}
