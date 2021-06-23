// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure TACACS+ server entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserTacacs() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserTacacsCreate,
		Read:   resourceObjectUserTacacsRead,
		Update: resourceObjectUserTacacsUpdate,
		Delete: resourceObjectUserTacacsDelete,

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
			"authen_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_scope": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"authen_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authorization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface_select_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"secondary_key": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"secondary_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tertiary_key": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"tertiary_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"secondary_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"secondary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tertiary_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"tertiary_server": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceObjectUserTacacsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserTacacs(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserTacacs resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserTacacs(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserTacacs resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserTacacsRead(d, m)
}

func resourceObjectUserTacacsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserTacacs(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserTacacs resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserTacacs(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserTacacs resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserTacacsRead(d, m)
}

func resourceObjectUserTacacsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserTacacs(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserTacacs resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserTacacsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserTacacs(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserTacacs resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserTacacs(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserTacacs resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserTacacsAuthenType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenObjectUserTacacsDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authen_type"
		if _, ok := i["authen-type"]; ok {
			v := flattenObjectUserTacacsDynamicMappingAuthenType(i["authen-type"], d, pre_append)
			tmp["authen_type"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-AuthenType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization"
		if _, ok := i["authorization"]; ok {
			v := flattenObjectUserTacacsDynamicMappingAuthorization(i["authorization"], d, pre_append)
			tmp["authorization"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-Authorization")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenObjectUserTacacsDynamicMappingInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenObjectUserTacacsDynamicMappingInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			v := flattenObjectUserTacacsDynamicMappingKey(i["key"], d, pre_append)
			tmp["key"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-Key")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectUserTacacsDynamicMappingPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_key"
		if _, ok := i["secondary-key"]; ok {
			v := flattenObjectUserTacacsDynamicMappingSecondaryKey(i["secondary-key"], d, pre_append)
			tmp["secondary_key"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-SecondaryKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_server"
		if _, ok := i["secondary-server"]; ok {
			v := flattenObjectUserTacacsDynamicMappingSecondaryServer(i["secondary-server"], d, pre_append)
			tmp["secondary_server"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-SecondaryServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenObjectUserTacacsDynamicMappingServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenObjectUserTacacsDynamicMappingSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_key"
		if _, ok := i["tertiary-key"]; ok {
			v := flattenObjectUserTacacsDynamicMappingTertiaryKey(i["tertiary-key"], d, pre_append)
			tmp["tertiary_key"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-TertiaryKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_server"
		if _, ok := i["tertiary-server"]; ok {
			v := flattenObjectUserTacacsDynamicMappingTertiaryServer(i["tertiary-server"], d, pre_append)
			tmp["tertiary_server"] = fortiAPISubPartPatch(v, "ObjectUserTacacs-DynamicMapping-TertiaryServer")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserTacacsDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectUserTacacsDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserTacacsDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserTacacsDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserTacacsDynamicMapping-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserTacacsDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsDynamicMappingAuthenType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsDynamicMappingAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsDynamicMappingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsDynamicMappingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsDynamicMappingKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserTacacsDynamicMappingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsDynamicMappingSecondaryKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserTacacsDynamicMappingSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsDynamicMappingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsDynamicMappingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsDynamicMappingTertiaryKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserTacacsDynamicMappingTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserTacacsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsSecondaryKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserTacacsSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserTacacsTertiaryKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserTacacsTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserTacacs(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("authen_type", flattenObjectUserTacacsAuthenType(o["authen-type"], d, "authen_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["authen-type"], "ObjectUserTacacs-AuthenType"); ok {
			if err = d.Set("authen_type", vv); err != nil {
				return fmt.Errorf("Error reading authen_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authen_type: %v", err)
		}
	}

	if err = d.Set("authorization", flattenObjectUserTacacsAuthorization(o["authorization"], d, "authorization")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorization"], "ObjectUserTacacs-Authorization"); ok {
			if err = d.Set("authorization", vv); err != nil {
				return fmt.Errorf("Error reading authorization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorization: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectUserTacacsDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserTacacs-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectUserTacacsDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserTacacs-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface", flattenObjectUserTacacsInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectUserTacacs-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectUserTacacsInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectUserTacacs-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("key", flattenObjectUserTacacsKey(o["key"], d, "key")); err != nil {
		if vv, ok := fortiAPIPatch(o["key"], "ObjectUserTacacs-Key"); ok {
			if err = d.Set("key", vv); err != nil {
				return fmt.Errorf("Error reading key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserTacacsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserTacacs-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectUserTacacsPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectUserTacacs-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("secondary_key", flattenObjectUserTacacsSecondaryKey(o["secondary-key"], d, "secondary_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-key"], "ObjectUserTacacs-SecondaryKey"); ok {
			if err = d.Set("secondary_key", vv); err != nil {
				return fmt.Errorf("Error reading secondary_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_key: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenObjectUserTacacsSecondaryServer(o["secondary-server"], d, "secondary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-server"], "ObjectUserTacacs-SecondaryServer"); ok {
			if err = d.Set("secondary_server", vv); err != nil {
				return fmt.Errorf("Error reading secondary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectUserTacacsServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectUserTacacs-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectUserTacacsSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectUserTacacs-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("tertiary_key", flattenObjectUserTacacsTertiaryKey(o["tertiary-key"], d, "tertiary_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["tertiary-key"], "ObjectUserTacacs-TertiaryKey"); ok {
			if err = d.Set("tertiary_key", vv); err != nil {
				return fmt.Errorf("Error reading tertiary_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tertiary_key: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenObjectUserTacacsTertiaryServer(o["tertiary-server"], d, "tertiary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tertiary-server"], "ObjectUserTacacs-TertiaryServer"); ok {
			if err = d.Set("tertiary_server", vv); err != nil {
				return fmt.Errorf("Error reading tertiary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	return nil
}

func flattenObjectUserTacacsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserTacacsAuthenType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["_scope"], _ = expandObjectUserTacacsDynamicMappingScope(d, i["_scope"], pre_append)
		} else {
			tmp["_scope"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authen_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authen-type"], _ = expandObjectUserTacacsDynamicMappingAuthenType(d, i["authen_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authorization"], _ = expandObjectUserTacacsDynamicMappingAuthorization(d, i["authorization"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandObjectUserTacacsDynamicMappingInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface-select-method"], _ = expandObjectUserTacacsDynamicMappingInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key"], _ = expandObjectUserTacacsDynamicMappingKey(d, i["key"], pre_append)
		} else {
			tmp["key"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandObjectUserTacacsDynamicMappingPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["secondary-key"], _ = expandObjectUserTacacsDynamicMappingSecondaryKey(d, i["secondary_key"], pre_append)
		} else {
			tmp["secondary-key"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["secondary-server"], _ = expandObjectUserTacacsDynamicMappingSecondaryServer(d, i["secondary_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server"], _ = expandObjectUserTacacsDynamicMappingServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source-ip"], _ = expandObjectUserTacacsDynamicMappingSourceIp(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tertiary-key"], _ = expandObjectUserTacacsDynamicMappingTertiaryKey(d, i["tertiary_key"], pre_append)
		} else {
			tmp["tertiary-key"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tertiary-server"], _ = expandObjectUserTacacsDynamicMappingTertiaryServer(d, i["tertiary_server"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserTacacsDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectUserTacacsDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vdom"], _ = expandObjectUserTacacsDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserTacacsDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsDynamicMappingAuthenType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsDynamicMappingAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsDynamicMappingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsDynamicMappingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsDynamicMappingKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserTacacsDynamicMappingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsDynamicMappingSecondaryKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserTacacsDynamicMappingSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsDynamicMappingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsDynamicMappingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsDynamicMappingTertiaryKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserTacacsDynamicMappingTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserTacacsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsSecondaryKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserTacacsSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserTacacsTertiaryKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserTacacsTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserTacacs(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authen_type"); ok {
		t, err := expandObjectUserTacacsAuthenType(d, v, "authen_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authen-type"] = t
		}
	}

	if v, ok := d.GetOk("authorization"); ok {
		t, err := expandObjectUserTacacsAuthorization(d, v, "authorization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorization"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok {
		t, err := expandObjectUserTacacsDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandObjectUserTacacsInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		t, err := expandObjectUserTacacsInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok {
		t, err := expandObjectUserTacacsKey(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectUserTacacsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandObjectUserTacacsPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("secondary_key"); ok {
		t, err := expandObjectUserTacacsSecondaryKey(d, v, "secondary_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-key"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok {
		t, err := expandObjectUserTacacsSecondaryServer(d, v, "secondary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandObjectUserTacacsServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandObjectUserTacacsSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_key"); ok {
		t, err := expandObjectUserTacacsTertiaryKey(d, v, "tertiary_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-key"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok {
		t, err := expandObjectUserTacacsTertiaryServer(d, v, "tertiary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	return &obj, nil
}
