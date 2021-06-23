// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure connection to SDN Connector.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemSdnConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemSdnConnectorCreate,
		Read:   resourceObjectSystemSdnConnectorRead,
		Update: resourceObjectSystemSdnConnectorUpdate,
		Delete: resourceObjectSystemSdnConnectorDelete,

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
			"_local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"api_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"azure_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_secret": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compute_generation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"gcp_project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ibm_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ibm_region_gen1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ibm_region_gen2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_passwd": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"last_update": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"login_endpoint": &schema.Schema{
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
			"nic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"public_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"resource_group": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"nsx_cert_fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oci_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oci_fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oci_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oci_region_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rest_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rest_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rest_sport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rest_ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"route_table": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"resource_group": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"next_hop": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"subscription_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"secret_key": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"secret_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"service_account": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"updating": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"use_metadata_iam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vcenter_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vmx_image_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vmx_service_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpc_id": &schema.Schema{
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

func resourceObjectSystemSdnConnectorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemSdnConnector(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnector resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemSdnConnector(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnector resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemSdnConnectorRead(d, m)
}

func resourceObjectSystemSdnConnectorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemSdnConnector(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnector resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemSdnConnector(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemSdnConnectorRead(d, m)
}

func resourceObjectSystemSdnConnectorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemSdnConnector(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemSdnConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemSdnConnectorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemSdnConnector(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemSdnConnector(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnector resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemSdnConnectorLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorAccessKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorApiKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemSdnConnectorAzureRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorClientId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorClientSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemSdnConnectorCompartmentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorComputeGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorExternalIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemSdnConnectorExternalIpName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnector-ExternalIp-Name")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemSdnConnectorExternalIpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorGcpProject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorHaStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorIbmRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorIbmRegionGen1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorIbmRegionGen2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorKeyPasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemSdnConnectorLastUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorLoginEndpoint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorNic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenObjectSystemSdnConnectorNicIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnector-Nic-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSystemSdnConnectorNicName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnector-Nic-Name")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemSdnConnectorNicIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemSdnConnectorNicIpName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnectorNic-Ip-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "public_ip"
		if _, ok := i["public-ip"]; ok {
			v := flattenObjectSystemSdnConnectorNicIpPublicIp(i["public-ip"], d, pre_append)
			tmp["public_ip"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnectorNic-Ip-PublicIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := i["resource-group"]; ok {
			v := flattenObjectSystemSdnConnectorNicIpResourceGroup(i["resource-group"], d, pre_append)
			tmp["resource_group"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnectorNic-Ip-ResourceGroup")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemSdnConnectorNicIpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorNicIpPublicIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorNicIpResourceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorNicName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorNsxCertFingerprint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorOciCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorOciFingerprint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorOciRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorOciRegionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemSdnConnectorPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorResourceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorResourceUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorRestInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorRestPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemSdnConnectorRestSport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorRestSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorRoute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemSdnConnectorRouteName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnector-Route-Name")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemSdnConnectorRouteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorRouteTable(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemSdnConnectorRouteTableName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnector-RouteTable-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := i["resource-group"]; ok {
			v := flattenObjectSystemSdnConnectorRouteTableResourceGroup(i["resource-group"], d, pre_append)
			tmp["resource_group"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnector-RouteTable-ResourceGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route"
		if _, ok := i["route"]; ok {
			v := flattenObjectSystemSdnConnectorRouteTableRoute(i["route"], d, pre_append)
			tmp["route"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnector-RouteTable-Route")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subscription_id"
		if _, ok := i["subscription-id"]; ok {
			v := flattenObjectSystemSdnConnectorRouteTableSubscriptionId(i["subscription-id"], d, pre_append)
			tmp["subscription_id"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnector-RouteTable-SubscriptionId")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemSdnConnectorRouteTableName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorRouteTableResourceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorRouteTableRoute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSystemSdnConnectorRouteTableRouteName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnectorRouteTable-Route-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop"
		if _, ok := i["next-hop"]; ok {
			v := flattenObjectSystemSdnConnectorRouteTableRouteNextHop(i["next-hop"], d, pre_append)
			tmp["next_hop"] = fortiAPISubPartPatch(v, "ObjectSystemSdnConnectorRouteTable-Route-NextHop")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSystemSdnConnectorRouteTableRouteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorRouteTableRouteNextHop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorRouteTableSubscriptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorSecretKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemSdnConnectorSecretToken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorServerList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemSdnConnectorServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorServiceAccount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorSubscriptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorTenantId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorUpdating(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorUseMetadataIam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorVcenterPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSystemSdnConnectorVcenterServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorVcenterUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorVmxImageUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorVmxServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorVpcId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemSdnConnector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("_local_cert", flattenObjectSystemSdnConnectorLocalCert(o["_local_cert"], d, "_local_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["_local_cert"], "ObjectSystemSdnConnector-LocalCert"); ok {
			if err = d.Set("_local_cert", vv); err != nil {
				return fmt.Errorf("Error reading _local_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _local_cert: %v", err)
		}
	}

	if err = d.Set("access_key", flattenObjectSystemSdnConnectorAccessKey(o["access-key"], d, "access_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-key"], "ObjectSystemSdnConnector-AccessKey"); ok {
			if err = d.Set("access_key", vv); err != nil {
				return fmt.Errorf("Error reading access_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_key: %v", err)
		}
	}

	if err = d.Set("api_key", flattenObjectSystemSdnConnectorApiKey(o["api-key"], d, "api_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["api-key"], "ObjectSystemSdnConnector-ApiKey"); ok {
			if err = d.Set("api_key", vv); err != nil {
				return fmt.Errorf("Error reading api_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading api_key: %v", err)
		}
	}

	if err = d.Set("azure_region", flattenObjectSystemSdnConnectorAzureRegion(o["azure-region"], d, "azure_region")); err != nil {
		if vv, ok := fortiAPIPatch(o["azure-region"], "ObjectSystemSdnConnector-AzureRegion"); ok {
			if err = d.Set("azure_region", vv); err != nil {
				return fmt.Errorf("Error reading azure_region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading azure_region: %v", err)
		}
	}

	if err = d.Set("client_id", flattenObjectSystemSdnConnectorClientId(o["client-id"], d, "client_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-id"], "ObjectSystemSdnConnector-ClientId"); ok {
			if err = d.Set("client_id", vv); err != nil {
				return fmt.Errorf("Error reading client_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_id: %v", err)
		}
	}

	if err = d.Set("client_secret", flattenObjectSystemSdnConnectorClientSecret(o["client-secret"], d, "client_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-secret"], "ObjectSystemSdnConnector-ClientSecret"); ok {
			if err = d.Set("client_secret", vv); err != nil {
				return fmt.Errorf("Error reading client_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_secret: %v", err)
		}
	}

	if err = d.Set("compartment_id", flattenObjectSystemSdnConnectorCompartmentId(o["compartment-id"], d, "compartment_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["compartment-id"], "ObjectSystemSdnConnector-CompartmentId"); ok {
			if err = d.Set("compartment_id", vv); err != nil {
				return fmt.Errorf("Error reading compartment_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading compartment_id: %v", err)
		}
	}

	if err = d.Set("compute_generation", flattenObjectSystemSdnConnectorComputeGeneration(o["compute-generation"], d, "compute_generation")); err != nil {
		if vv, ok := fortiAPIPatch(o["compute-generation"], "ObjectSystemSdnConnector-ComputeGeneration"); ok {
			if err = d.Set("compute_generation", vv); err != nil {
				return fmt.Errorf("Error reading compute_generation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading compute_generation: %v", err)
		}
	}

	if err = d.Set("domain", flattenObjectSystemSdnConnectorDomain(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ObjectSystemSdnConnector-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("external_ip", flattenObjectSystemSdnConnectorExternalIp(o["external-ip"], d, "external_ip")); err != nil {
			if vv, ok := fortiAPIPatch(o["external-ip"], "ObjectSystemSdnConnector-ExternalIp"); ok {
				if err = d.Set("external_ip", vv); err != nil {
					return fmt.Errorf("Error reading external_ip: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading external_ip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("external_ip"); ok {
			if err = d.Set("external_ip", flattenObjectSystemSdnConnectorExternalIp(o["external-ip"], d, "external_ip")); err != nil {
				if vv, ok := fortiAPIPatch(o["external-ip"], "ObjectSystemSdnConnector-ExternalIp"); ok {
					if err = d.Set("external_ip", vv); err != nil {
						return fmt.Errorf("Error reading external_ip: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading external_ip: %v", err)
				}
			}
		}
	}

	if err = d.Set("gcp_project", flattenObjectSystemSdnConnectorGcpProject(o["gcp-project"], d, "gcp_project")); err != nil {
		if vv, ok := fortiAPIPatch(o["gcp-project"], "ObjectSystemSdnConnector-GcpProject"); ok {
			if err = d.Set("gcp_project", vv); err != nil {
				return fmt.Errorf("Error reading gcp_project: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gcp_project: %v", err)
		}
	}

	if err = d.Set("group_name", flattenObjectSystemSdnConnectorGroupName(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "ObjectSystemSdnConnector-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("ha_status", flattenObjectSystemSdnConnectorHaStatus(o["ha-status"], d, "ha_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-status"], "ObjectSystemSdnConnector-HaStatus"); ok {
			if err = d.Set("ha_status", vv); err != nil {
				return fmt.Errorf("Error reading ha_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_status: %v", err)
		}
	}

	if err = d.Set("ibm_region", flattenObjectSystemSdnConnectorIbmRegion(o["ibm-region"], d, "ibm_region")); err != nil {
		if vv, ok := fortiAPIPatch(o["ibm-region"], "ObjectSystemSdnConnector-IbmRegion"); ok {
			if err = d.Set("ibm_region", vv); err != nil {
				return fmt.Errorf("Error reading ibm_region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ibm_region: %v", err)
		}
	}

	if err = d.Set("ibm_region_gen1", flattenObjectSystemSdnConnectorIbmRegionGen1(o["ibm-region-gen1"], d, "ibm_region_gen1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ibm-region-gen1"], "ObjectSystemSdnConnector-IbmRegionGen1"); ok {
			if err = d.Set("ibm_region_gen1", vv); err != nil {
				return fmt.Errorf("Error reading ibm_region_gen1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ibm_region_gen1: %v", err)
		}
	}

	if err = d.Set("ibm_region_gen2", flattenObjectSystemSdnConnectorIbmRegionGen2(o["ibm-region-gen2"], d, "ibm_region_gen2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ibm-region-gen2"], "ObjectSystemSdnConnector-IbmRegionGen2"); ok {
			if err = d.Set("ibm_region_gen2", vv); err != nil {
				return fmt.Errorf("Error reading ibm_region_gen2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ibm_region_gen2: %v", err)
		}
	}

	if err = d.Set("key_passwd", flattenObjectSystemSdnConnectorKeyPasswd(o["key-passwd"], d, "key_passwd")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-passwd"], "ObjectSystemSdnConnector-KeyPasswd"); ok {
			if err = d.Set("key_passwd", vv); err != nil {
				return fmt.Errorf("Error reading key_passwd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_passwd: %v", err)
		}
	}

	if err = d.Set("last_update", flattenObjectSystemSdnConnectorLastUpdate(o["last-update"], d, "last_update")); err != nil {
		if vv, ok := fortiAPIPatch(o["last-update"], "ObjectSystemSdnConnector-LastUpdate"); ok {
			if err = d.Set("last_update", vv); err != nil {
				return fmt.Errorf("Error reading last_update: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_update: %v", err)
		}
	}

	if err = d.Set("login_endpoint", flattenObjectSystemSdnConnectorLoginEndpoint(o["login-endpoint"], d, "login_endpoint")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-endpoint"], "ObjectSystemSdnConnector-LoginEndpoint"); ok {
			if err = d.Set("login_endpoint", vv); err != nil {
				return fmt.Errorf("Error reading login_endpoint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_endpoint: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSystemSdnConnectorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemSdnConnector-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nic", flattenObjectSystemSdnConnectorNic(o["nic"], d, "nic")); err != nil {
			if vv, ok := fortiAPIPatch(o["nic"], "ObjectSystemSdnConnector-Nic"); ok {
				if err = d.Set("nic", vv); err != nil {
					return fmt.Errorf("Error reading nic: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nic"); ok {
			if err = d.Set("nic", flattenObjectSystemSdnConnectorNic(o["nic"], d, "nic")); err != nil {
				if vv, ok := fortiAPIPatch(o["nic"], "ObjectSystemSdnConnector-Nic"); ok {
					if err = d.Set("nic", vv); err != nil {
						return fmt.Errorf("Error reading nic: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nic: %v", err)
				}
			}
		}
	}

	if err = d.Set("nsx_cert_fingerprint", flattenObjectSystemSdnConnectorNsxCertFingerprint(o["nsx-cert-fingerprint"], d, "nsx_cert_fingerprint")); err != nil {
		if vv, ok := fortiAPIPatch(o["nsx-cert-fingerprint"], "ObjectSystemSdnConnector-NsxCertFingerprint"); ok {
			if err = d.Set("nsx_cert_fingerprint", vv); err != nil {
				return fmt.Errorf("Error reading nsx_cert_fingerprint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nsx_cert_fingerprint: %v", err)
		}
	}

	if err = d.Set("oci_cert", flattenObjectSystemSdnConnectorOciCert(o["oci-cert"], d, "oci_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["oci-cert"], "ObjectSystemSdnConnector-OciCert"); ok {
			if err = d.Set("oci_cert", vv); err != nil {
				return fmt.Errorf("Error reading oci_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oci_cert: %v", err)
		}
	}

	if err = d.Set("oci_fingerprint", flattenObjectSystemSdnConnectorOciFingerprint(o["oci-fingerprint"], d, "oci_fingerprint")); err != nil {
		if vv, ok := fortiAPIPatch(o["oci-fingerprint"], "ObjectSystemSdnConnector-OciFingerprint"); ok {
			if err = d.Set("oci_fingerprint", vv); err != nil {
				return fmt.Errorf("Error reading oci_fingerprint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oci_fingerprint: %v", err)
		}
	}

	if err = d.Set("oci_region", flattenObjectSystemSdnConnectorOciRegion(o["oci-region"], d, "oci_region")); err != nil {
		if vv, ok := fortiAPIPatch(o["oci-region"], "ObjectSystemSdnConnector-OciRegion"); ok {
			if err = d.Set("oci_region", vv); err != nil {
				return fmt.Errorf("Error reading oci_region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oci_region: %v", err)
		}
	}

	if err = d.Set("oci_region_type", flattenObjectSystemSdnConnectorOciRegionType(o["oci-region-type"], d, "oci_region_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["oci-region-type"], "ObjectSystemSdnConnector-OciRegionType"); ok {
			if err = d.Set("oci_region_type", vv); err != nil {
				return fmt.Errorf("Error reading oci_region_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oci_region_type: %v", err)
		}
	}

	if err = d.Set("password", flattenObjectSystemSdnConnectorPassword(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "ObjectSystemSdnConnector-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("private_key", flattenObjectSystemSdnConnectorPrivateKey(o["private-key"], d, "private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-key"], "ObjectSystemSdnConnector-PrivateKey"); ok {
			if err = d.Set("private_key", vv); err != nil {
				return fmt.Errorf("Error reading private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_key: %v", err)
		}
	}

	if err = d.Set("region", flattenObjectSystemSdnConnectorRegion(o["region"], d, "region")); err != nil {
		if vv, ok := fortiAPIPatch(o["region"], "ObjectSystemSdnConnector-Region"); ok {
			if err = d.Set("region", vv); err != nil {
				return fmt.Errorf("Error reading region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region: %v", err)
		}
	}

	if err = d.Set("resource_group", flattenObjectSystemSdnConnectorResourceGroup(o["resource-group"], d, "resource_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["resource-group"], "ObjectSystemSdnConnector-ResourceGroup"); ok {
			if err = d.Set("resource_group", vv); err != nil {
				return fmt.Errorf("Error reading resource_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resource_group: %v", err)
		}
	}

	if err = d.Set("resource_url", flattenObjectSystemSdnConnectorResourceUrl(o["resource-url"], d, "resource_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["resource-url"], "ObjectSystemSdnConnector-ResourceUrl"); ok {
			if err = d.Set("resource_url", vv); err != nil {
				return fmt.Errorf("Error reading resource_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resource_url: %v", err)
		}
	}

	if err = d.Set("rest_interface", flattenObjectSystemSdnConnectorRestInterface(o["rest-interface"], d, "rest_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["rest-interface"], "ObjectSystemSdnConnector-RestInterface"); ok {
			if err = d.Set("rest_interface", vv); err != nil {
				return fmt.Errorf("Error reading rest_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rest_interface: %v", err)
		}
	}

	if err = d.Set("rest_password", flattenObjectSystemSdnConnectorRestPassword(o["rest-password"], d, "rest_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["rest-password"], "ObjectSystemSdnConnector-RestPassword"); ok {
			if err = d.Set("rest_password", vv); err != nil {
				return fmt.Errorf("Error reading rest_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rest_password: %v", err)
		}
	}

	if err = d.Set("rest_sport", flattenObjectSystemSdnConnectorRestSport(o["rest-sport"], d, "rest_sport")); err != nil {
		if vv, ok := fortiAPIPatch(o["rest-sport"], "ObjectSystemSdnConnector-RestSport"); ok {
			if err = d.Set("rest_sport", vv); err != nil {
				return fmt.Errorf("Error reading rest_sport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rest_sport: %v", err)
		}
	}

	if err = d.Set("rest_ssl", flattenObjectSystemSdnConnectorRestSsl(o["rest-ssl"], d, "rest_ssl")); err != nil {
		if vv, ok := fortiAPIPatch(o["rest-ssl"], "ObjectSystemSdnConnector-RestSsl"); ok {
			if err = d.Set("rest_ssl", vv); err != nil {
				return fmt.Errorf("Error reading rest_ssl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rest_ssl: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("route", flattenObjectSystemSdnConnectorRoute(o["route"], d, "route")); err != nil {
			if vv, ok := fortiAPIPatch(o["route"], "ObjectSystemSdnConnector-Route"); ok {
				if err = d.Set("route", vv); err != nil {
					return fmt.Errorf("Error reading route: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading route: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("route"); ok {
			if err = d.Set("route", flattenObjectSystemSdnConnectorRoute(o["route"], d, "route")); err != nil {
				if vv, ok := fortiAPIPatch(o["route"], "ObjectSystemSdnConnector-Route"); ok {
					if err = d.Set("route", vv); err != nil {
						return fmt.Errorf("Error reading route: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading route: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("route_table", flattenObjectSystemSdnConnectorRouteTable(o["route-table"], d, "route_table")); err != nil {
			if vv, ok := fortiAPIPatch(o["route-table"], "ObjectSystemSdnConnector-RouteTable"); ok {
				if err = d.Set("route_table", vv); err != nil {
					return fmt.Errorf("Error reading route_table: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading route_table: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("route_table"); ok {
			if err = d.Set("route_table", flattenObjectSystemSdnConnectorRouteTable(o["route-table"], d, "route_table")); err != nil {
				if vv, ok := fortiAPIPatch(o["route-table"], "ObjectSystemSdnConnector-RouteTable"); ok {
					if err = d.Set("route_table", vv); err != nil {
						return fmt.Errorf("Error reading route_table: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading route_table: %v", err)
				}
			}
		}
	}

	if err = d.Set("secret_key", flattenObjectSystemSdnConnectorSecretKey(o["secret-key"], d, "secret_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["secret-key"], "ObjectSystemSdnConnector-SecretKey"); ok {
			if err = d.Set("secret_key", vv); err != nil {
				return fmt.Errorf("Error reading secret_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secret_key: %v", err)
		}
	}

	if err = d.Set("secret_token", flattenObjectSystemSdnConnectorSecretToken(o["secret-token"], d, "secret_token")); err != nil {
		if vv, ok := fortiAPIPatch(o["secret-token"], "ObjectSystemSdnConnector-SecretToken"); ok {
			if err = d.Set("secret_token", vv); err != nil {
				return fmt.Errorf("Error reading secret_token: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secret_token: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectSystemSdnConnectorServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectSystemSdnConnector-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_list", flattenObjectSystemSdnConnectorServerList(o["server-list"], d, "server_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-list"], "ObjectSystemSdnConnector-ServerList"); ok {
			if err = d.Set("server_list", vv); err != nil {
				return fmt.Errorf("Error reading server_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_list: %v", err)
		}
	}

	if err = d.Set("server_port", flattenObjectSystemSdnConnectorServerPort(o["server-port"], d, "server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-port"], "ObjectSystemSdnConnector-ServerPort"); ok {
			if err = d.Set("server_port", vv); err != nil {
				return fmt.Errorf("Error reading server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("service_account", flattenObjectSystemSdnConnectorServiceAccount(o["service-account"], d, "service_account")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-account"], "ObjectSystemSdnConnector-ServiceAccount"); ok {
			if err = d.Set("service_account", vv); err != nil {
				return fmt.Errorf("Error reading service_account: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_account: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectSystemSdnConnectorStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectSystemSdnConnector-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("subscription_id", flattenObjectSystemSdnConnectorSubscriptionId(o["subscription-id"], d, "subscription_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["subscription-id"], "ObjectSystemSdnConnector-SubscriptionId"); ok {
			if err = d.Set("subscription_id", vv); err != nil {
				return fmt.Errorf("Error reading subscription_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subscription_id: %v", err)
		}
	}

	if err = d.Set("tenant_id", flattenObjectSystemSdnConnectorTenantId(o["tenant-id"], d, "tenant_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant-id"], "ObjectSystemSdnConnector-TenantId"); ok {
			if err = d.Set("tenant_id", vv); err != nil {
				return fmt.Errorf("Error reading tenant_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant_id: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectSystemSdnConnectorType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSystemSdnConnector-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenObjectSystemSdnConnectorUpdateInterval(o["update-interval"], d, "update_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-interval"], "ObjectSystemSdnConnector-UpdateInterval"); ok {
			if err = d.Set("update_interval", vv); err != nil {
				return fmt.Errorf("Error reading update_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("updating", flattenObjectSystemSdnConnectorUpdating(o["updating"], d, "updating")); err != nil {
		if vv, ok := fortiAPIPatch(o["updating"], "ObjectSystemSdnConnector-Updating"); ok {
			if err = d.Set("updating", vv); err != nil {
				return fmt.Errorf("Error reading updating: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading updating: %v", err)
		}
	}

	if err = d.Set("use_metadata_iam", flattenObjectSystemSdnConnectorUseMetadataIam(o["use-metadata-iam"], d, "use_metadata_iam")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-metadata-iam"], "ObjectSystemSdnConnector-UseMetadataIam"); ok {
			if err = d.Set("use_metadata_iam", vv); err != nil {
				return fmt.Errorf("Error reading use_metadata_iam: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_metadata_iam: %v", err)
		}
	}

	if err = d.Set("user_id", flattenObjectSystemSdnConnectorUserId(o["user-id"], d, "user_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-id"], "ObjectSystemSdnConnector-UserId"); ok {
			if err = d.Set("user_id", vv); err != nil {
				return fmt.Errorf("Error reading user_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_id: %v", err)
		}
	}

	if err = d.Set("username", flattenObjectSystemSdnConnectorUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "ObjectSystemSdnConnector-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("vcenter_password", flattenObjectSystemSdnConnectorVcenterPassword(o["vcenter-password"], d, "vcenter_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["vcenter-password"], "ObjectSystemSdnConnector-VcenterPassword"); ok {
			if err = d.Set("vcenter_password", vv); err != nil {
				return fmt.Errorf("Error reading vcenter_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vcenter_password: %v", err)
		}
	}

	if err = d.Set("vcenter_server", flattenObjectSystemSdnConnectorVcenterServer(o["vcenter-server"], d, "vcenter_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["vcenter-server"], "ObjectSystemSdnConnector-VcenterServer"); ok {
			if err = d.Set("vcenter_server", vv); err != nil {
				return fmt.Errorf("Error reading vcenter_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vcenter_server: %v", err)
		}
	}

	if err = d.Set("vcenter_username", flattenObjectSystemSdnConnectorVcenterUsername(o["vcenter-username"], d, "vcenter_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["vcenter-username"], "ObjectSystemSdnConnector-VcenterUsername"); ok {
			if err = d.Set("vcenter_username", vv); err != nil {
				return fmt.Errorf("Error reading vcenter_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vcenter_username: %v", err)
		}
	}

	if err = d.Set("vmx_image_url", flattenObjectSystemSdnConnectorVmxImageUrl(o["vmx-image-url"], d, "vmx_image_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["vmx-image-url"], "ObjectSystemSdnConnector-VmxImageUrl"); ok {
			if err = d.Set("vmx_image_url", vv); err != nil {
				return fmt.Errorf("Error reading vmx_image_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vmx_image_url: %v", err)
		}
	}

	if err = d.Set("vmx_service_name", flattenObjectSystemSdnConnectorVmxServiceName(o["vmx-service-name"], d, "vmx_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["vmx-service-name"], "ObjectSystemSdnConnector-VmxServiceName"); ok {
			if err = d.Set("vmx_service_name", vv); err != nil {
				return fmt.Errorf("Error reading vmx_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vmx_service_name: %v", err)
		}
	}

	if err = d.Set("vpc_id", flattenObjectSystemSdnConnectorVpcId(o["vpc-id"], d, "vpc_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpc-id"], "ObjectSystemSdnConnector-VpcId"); ok {
			if err = d.Set("vpc_id", vv); err != nil {
				return fmt.Errorf("Error reading vpc_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpc_id: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemSdnConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemSdnConnectorLocalCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorAccessKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorApiKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemSdnConnectorAzureRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorClientId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorClientSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemSdnConnectorCompartmentId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorComputeGeneration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorExternalIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectSystemSdnConnectorExternalIpName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemSdnConnectorExternalIpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorGcpProject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorHaStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorIbmRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorIbmRegionGen1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorIbmRegionGen2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorKeyPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemSdnConnectorLastUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorLoginEndpoint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorNic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandObjectSystemSdnConnectorNicIp(d, i["ip"], pre_append)
		} else {
			tmp["ip"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectSystemSdnConnectorNicName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemSdnConnectorNicIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectSystemSdnConnectorNicIpName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "public_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["public-ip"], _ = expandObjectSystemSdnConnectorNicIpPublicIp(d, i["public_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["resource-group"], _ = expandObjectSystemSdnConnectorNicIpResourceGroup(d, i["resource_group"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemSdnConnectorNicIpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorNicIpPublicIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorNicIpResourceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorNicName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorNsxCertFingerprint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorOciCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorOciFingerprint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorOciRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorOciRegionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemSdnConnectorPrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorResourceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorResourceUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorRestInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorRestPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemSdnConnectorRestSport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorRestSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectSystemSdnConnectorRouteName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemSdnConnectorRouteName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorRouteTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectSystemSdnConnectorRouteTableName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["resource-group"], _ = expandObjectSystemSdnConnectorRouteTableResourceGroup(d, i["resource_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route"], _ = expandObjectSystemSdnConnectorRouteTableRoute(d, i["route"], pre_append)
		} else {
			tmp["route"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subscription_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["subscription-id"], _ = expandObjectSystemSdnConnectorRouteTableSubscriptionId(d, i["subscription_id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemSdnConnectorRouteTableName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorRouteTableResourceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorRouteTableRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectSystemSdnConnectorRouteTableRouteName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop"], _ = expandObjectSystemSdnConnectorRouteTableRouteNextHop(d, i["next_hop"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSystemSdnConnectorRouteTableRouteName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorRouteTableRouteNextHop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorRouteTableSubscriptionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorSecretKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemSdnConnectorSecretToken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorServerList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemSdnConnectorServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorServiceAccount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorSubscriptionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorTenantId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorUpdateInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorUpdating(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorUseMetadataIam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorUserId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorVcenterPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSystemSdnConnectorVcenterServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorVcenterUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorVmxImageUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorVmxServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorVpcId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemSdnConnector(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_local_cert"); ok {
		t, err := expandObjectSystemSdnConnectorLocalCert(d, v, "_local_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_local_cert"] = t
		}
	}

	if v, ok := d.GetOk("access_key"); ok {
		t, err := expandObjectSystemSdnConnectorAccessKey(d, v, "access_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-key"] = t
		}
	}

	if v, ok := d.GetOk("api_key"); ok {
		t, err := expandObjectSystemSdnConnectorApiKey(d, v, "api_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-key"] = t
		}
	}

	if v, ok := d.GetOk("azure_region"); ok {
		t, err := expandObjectSystemSdnConnectorAzureRegion(d, v, "azure_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-region"] = t
		}
	}

	if v, ok := d.GetOk("client_id"); ok {
		t, err := expandObjectSystemSdnConnectorClientId(d, v, "client_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-id"] = t
		}
	}

	if v, ok := d.GetOk("client_secret"); ok {
		t, err := expandObjectSystemSdnConnectorClientSecret(d, v, "client_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-secret"] = t
		}
	}

	if v, ok := d.GetOk("compartment_id"); ok {
		t, err := expandObjectSystemSdnConnectorCompartmentId(d, v, "compartment_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compartment-id"] = t
		}
	}

	if v, ok := d.GetOk("compute_generation"); ok {
		t, err := expandObjectSystemSdnConnectorComputeGeneration(d, v, "compute_generation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compute-generation"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandObjectSystemSdnConnectorDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("external_ip"); ok {
		t, err := expandObjectSystemSdnConnectorExternalIp(d, v, "external_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-ip"] = t
		}
	}

	if v, ok := d.GetOk("gcp_project"); ok {
		t, err := expandObjectSystemSdnConnectorGcpProject(d, v, "gcp_project")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-project"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok {
		t, err := expandObjectSystemSdnConnectorGroupName(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("ha_status"); ok {
		t, err := expandObjectSystemSdnConnectorHaStatus(d, v, "ha_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-status"] = t
		}
	}

	if v, ok := d.GetOk("ibm_region"); ok {
		t, err := expandObjectSystemSdnConnectorIbmRegion(d, v, "ibm_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibm-region"] = t
		}
	}

	if v, ok := d.GetOk("ibm_region_gen1"); ok {
		t, err := expandObjectSystemSdnConnectorIbmRegionGen1(d, v, "ibm_region_gen1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibm-region-gen1"] = t
		}
	}

	if v, ok := d.GetOk("ibm_region_gen2"); ok {
		t, err := expandObjectSystemSdnConnectorIbmRegionGen2(d, v, "ibm_region_gen2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibm-region-gen2"] = t
		}
	}

	if v, ok := d.GetOk("key_passwd"); ok {
		t, err := expandObjectSystemSdnConnectorKeyPasswd(d, v, "key_passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-passwd"] = t
		}
	}

	if v, ok := d.GetOk("last_update"); ok {
		t, err := expandObjectSystemSdnConnectorLastUpdate(d, v, "last_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-update"] = t
		}
	}

	if v, ok := d.GetOk("login_endpoint"); ok {
		t, err := expandObjectSystemSdnConnectorLoginEndpoint(d, v, "login_endpoint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-endpoint"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectSystemSdnConnectorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nic"); ok {
		t, err := expandObjectSystemSdnConnectorNic(d, v, "nic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nic"] = t
		}
	}

	if v, ok := d.GetOk("nsx_cert_fingerprint"); ok {
		t, err := expandObjectSystemSdnConnectorNsxCertFingerprint(d, v, "nsx_cert_fingerprint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nsx-cert-fingerprint"] = t
		}
	}

	if v, ok := d.GetOk("oci_cert"); ok {
		t, err := expandObjectSystemSdnConnectorOciCert(d, v, "oci_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-cert"] = t
		}
	}

	if v, ok := d.GetOk("oci_fingerprint"); ok {
		t, err := expandObjectSystemSdnConnectorOciFingerprint(d, v, "oci_fingerprint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-fingerprint"] = t
		}
	}

	if v, ok := d.GetOk("oci_region"); ok {
		t, err := expandObjectSystemSdnConnectorOciRegion(d, v, "oci_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-region"] = t
		}
	}

	if v, ok := d.GetOk("oci_region_type"); ok {
		t, err := expandObjectSystemSdnConnectorOciRegionType(d, v, "oci_region_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-region-type"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandObjectSystemSdnConnectorPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok {
		t, err := expandObjectSystemSdnConnectorPrivateKey(d, v, "private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	}

	if v, ok := d.GetOk("region"); ok {
		t, err := expandObjectSystemSdnConnectorRegion(d, v, "region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region"] = t
		}
	}

	if v, ok := d.GetOk("resource_group"); ok {
		t, err := expandObjectSystemSdnConnectorResourceGroup(d, v, "resource_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-group"] = t
		}
	}

	if v, ok := d.GetOk("resource_url"); ok {
		t, err := expandObjectSystemSdnConnectorResourceUrl(d, v, "resource_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-url"] = t
		}
	}

	if v, ok := d.GetOk("rest_interface"); ok {
		t, err := expandObjectSystemSdnConnectorRestInterface(d, v, "rest_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-interface"] = t
		}
	}

	if v, ok := d.GetOk("rest_password"); ok {
		t, err := expandObjectSystemSdnConnectorRestPassword(d, v, "rest_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-password"] = t
		}
	}

	if v, ok := d.GetOk("rest_sport"); ok {
		t, err := expandObjectSystemSdnConnectorRestSport(d, v, "rest_sport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-sport"] = t
		}
	}

	if v, ok := d.GetOk("rest_ssl"); ok {
		t, err := expandObjectSystemSdnConnectorRestSsl(d, v, "rest_ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-ssl"] = t
		}
	}

	if v, ok := d.GetOk("route"); ok {
		t, err := expandObjectSystemSdnConnectorRoute(d, v, "route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route"] = t
		}
	}

	if v, ok := d.GetOk("route_table"); ok {
		t, err := expandObjectSystemSdnConnectorRouteTable(d, v, "route_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-table"] = t
		}
	}

	if v, ok := d.GetOk("secret_key"); ok {
		t, err := expandObjectSystemSdnConnectorSecretKey(d, v, "secret_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret-key"] = t
		}
	}

	if v, ok := d.GetOk("secret_token"); ok {
		t, err := expandObjectSystemSdnConnectorSecretToken(d, v, "secret_token")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret-token"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandObjectSystemSdnConnectorServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_list"); ok {
		t, err := expandObjectSystemSdnConnectorServerList(d, v, "server_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-list"] = t
		}
	}

	if v, ok := d.GetOk("server_port"); ok {
		t, err := expandObjectSystemSdnConnectorServerPort(d, v, "server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-port"] = t
		}
	}

	if v, ok := d.GetOk("service_account"); ok {
		t, err := expandObjectSystemSdnConnectorServiceAccount(d, v, "service_account")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-account"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandObjectSystemSdnConnectorStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("subscription_id"); ok {
		t, err := expandObjectSystemSdnConnectorSubscriptionId(d, v, "subscription_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subscription-id"] = t
		}
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		t, err := expandObjectSystemSdnConnectorTenantId(d, v, "tenant_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandObjectSystemSdnConnectorType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok {
		t, err := expandObjectSystemSdnConnectorUpdateInterval(d, v, "update_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	if v, ok := d.GetOk("updating"); ok {
		t, err := expandObjectSystemSdnConnectorUpdating(d, v, "updating")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["updating"] = t
		}
	}

	if v, ok := d.GetOk("use_metadata_iam"); ok {
		t, err := expandObjectSystemSdnConnectorUseMetadataIam(d, v, "use_metadata_iam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-metadata-iam"] = t
		}
	}

	if v, ok := d.GetOk("user_id"); ok {
		t, err := expandObjectSystemSdnConnectorUserId(d, v, "user_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-id"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandObjectSystemSdnConnectorUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("vcenter_password"); ok {
		t, err := expandObjectSystemSdnConnectorVcenterPassword(d, v, "vcenter_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcenter-password"] = t
		}
	}

	if v, ok := d.GetOk("vcenter_server"); ok {
		t, err := expandObjectSystemSdnConnectorVcenterServer(d, v, "vcenter_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcenter-server"] = t
		}
	}

	if v, ok := d.GetOk("vcenter_username"); ok {
		t, err := expandObjectSystemSdnConnectorVcenterUsername(d, v, "vcenter_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcenter-username"] = t
		}
	}

	if v, ok := d.GetOk("vmx_image_url"); ok {
		t, err := expandObjectSystemSdnConnectorVmxImageUrl(d, v, "vmx_image_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vmx-image-url"] = t
		}
	}

	if v, ok := d.GetOk("vmx_service_name"); ok {
		t, err := expandObjectSystemSdnConnectorVmxServiceName(d, v, "vmx_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vmx-service-name"] = t
		}
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		t, err := expandObjectSystemSdnConnectorVpcId(d, v, "vpc_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpc-id"] = t
		}
	}

	return &obj, nil
}
