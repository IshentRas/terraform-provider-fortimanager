// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectCertificate Template

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectCertificateTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCertificateTemplateCreate,
		Read:   resourceObjectCertificateTemplateRead,
		Update: resourceObjectCertificateTemplateUpdate,
		Delete: resourceObjectCertificateTemplateDelete,

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
			"city": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"curve_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"digest_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_type": &schema.Schema{
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
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"organization_unit": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"scep_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"scep_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subject_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectCertificateTemplateCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectCertificateTemplate(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCertificateTemplate resource while getting object: %v", err)
	}

	_, err = c.CreateObjectCertificateTemplate(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectCertificateTemplate resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCertificateTemplateRead(d, m)
}

func resourceObjectCertificateTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectCertificateTemplate(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCertificateTemplate resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectCertificateTemplate(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCertificateTemplate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCertificateTemplateRead(d, m)
}

func resourceObjectCertificateTemplateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectCertificateTemplate(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCertificateTemplate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCertificateTemplateRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectCertificateTemplate(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCertificateTemplate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCertificateTemplate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCertificateTemplate resource from API: %v", err)
	}
	return nil
}

func flattenObjectCertificateTemplateCity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateCurveName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "secp256r1",
			1: "secp384r1",
			2: "secp521r1",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectCertificateTemplateDigestType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "sha1",
			1: "sha256",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectCertificateTemplateEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "host-ip",
			1: "domain-name",
			2: "email",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectCertificateTemplateKeySize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "512",
			1: "1024",
			2: "1536",
			3: "2048",
			4: "4096",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectCertificateTemplateKeyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "rsa",
			1: "ec",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectCertificateTemplateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateOrganization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateOrganizationUnit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCertificateTemplateScepPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCertificateTemplateScepServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateSubjectName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "external",
			1: "local",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectCertificateTemplate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("city", flattenObjectCertificateTemplateCity(o["city"], d, "city")); err != nil {
		if vv, ok := fortiAPIPatch(o["city"], "ObjectCertificateTemplate-City"); ok {
			if err = d.Set("city", vv); err != nil {
				return fmt.Errorf("Error reading city: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading city: %v", err)
		}
	}

	if err = d.Set("country", flattenObjectCertificateTemplateCountry(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "ObjectCertificateTemplate-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("curve_name", flattenObjectCertificateTemplateCurveName(o["curve-name"], d, "curve_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["curve-name"], "ObjectCertificateTemplate-CurveName"); ok {
			if err = d.Set("curve_name", vv); err != nil {
				return fmt.Errorf("Error reading curve_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading curve_name: %v", err)
		}
	}

	if err = d.Set("digest_type", flattenObjectCertificateTemplateDigestType(o["digest-type"], d, "digest_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["digest-type"], "ObjectCertificateTemplate-DigestType"); ok {
			if err = d.Set("digest_type", vv); err != nil {
				return fmt.Errorf("Error reading digest_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading digest_type: %v", err)
		}
	}

	if err = d.Set("email", flattenObjectCertificateTemplateEmail(o["email"], d, "email")); err != nil {
		if vv, ok := fortiAPIPatch(o["email"], "ObjectCertificateTemplate-Email"); ok {
			if err = d.Set("email", vv); err != nil {
				return fmt.Errorf("Error reading email: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email: %v", err)
		}
	}

	if err = d.Set("id_type", flattenObjectCertificateTemplateIdType(o["id-type"], d, "id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["id-type"], "ObjectCertificateTemplate-IdType"); ok {
			if err = d.Set("id_type", vv); err != nil {
				return fmt.Errorf("Error reading id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading id_type: %v", err)
		}
	}

	if err = d.Set("key_size", flattenObjectCertificateTemplateKeySize(o["key-size"], d, "key_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-size"], "ObjectCertificateTemplate-KeySize"); ok {
			if err = d.Set("key_size", vv); err != nil {
				return fmt.Errorf("Error reading key_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_size: %v", err)
		}
	}

	if err = d.Set("key_type", flattenObjectCertificateTemplateKeyType(o["key-type"], d, "key_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-type"], "ObjectCertificateTemplate-KeyType"); ok {
			if err = d.Set("key_type", vv); err != nil {
				return fmt.Errorf("Error reading key_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_type: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCertificateTemplateName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCertificateTemplate-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("organization", flattenObjectCertificateTemplateOrganization(o["organization"], d, "organization")); err != nil {
		if vv, ok := fortiAPIPatch(o["organization"], "ObjectCertificateTemplate-Organization"); ok {
			if err = d.Set("organization", vv); err != nil {
				return fmt.Errorf("Error reading organization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading organization: %v", err)
		}
	}

	if err = d.Set("organization_unit", flattenObjectCertificateTemplateOrganizationUnit(o["organization-unit"], d, "organization_unit")); err != nil {
		if vv, ok := fortiAPIPatch(o["organization-unit"], "ObjectCertificateTemplate-OrganizationUnit"); ok {
			if err = d.Set("organization_unit", vv); err != nil {
				return fmt.Errorf("Error reading organization_unit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading organization_unit: %v", err)
		}
	}

	if err = d.Set("scep_password", flattenObjectCertificateTemplateScepPassword(o["scep-password"], d, "scep_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["scep-password"], "ObjectCertificateTemplate-ScepPassword"); ok {
			if err = d.Set("scep_password", vv); err != nil {
				return fmt.Errorf("Error reading scep_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scep_password: %v", err)
		}
	}

	if err = d.Set("scep_server", flattenObjectCertificateTemplateScepServer(o["scep-server"], d, "scep_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["scep-server"], "ObjectCertificateTemplate-ScepServer"); ok {
			if err = d.Set("scep_server", vv); err != nil {
				return fmt.Errorf("Error reading scep_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scep_server: %v", err)
		}
	}

	if err = d.Set("state", flattenObjectCertificateTemplateState(o["state"], d, "state")); err != nil {
		if vv, ok := fortiAPIPatch(o["state"], "ObjectCertificateTemplate-State"); ok {
			if err = d.Set("state", vv); err != nil {
				return fmt.Errorf("Error reading state: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading state: %v", err)
		}
	}

	if err = d.Set("subject_name", flattenObjectCertificateTemplateSubjectName(o["subject-name"], d, "subject_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["subject-name"], "ObjectCertificateTemplate-SubjectName"); ok {
			if err = d.Set("subject_name", vv); err != nil {
				return fmt.Errorf("Error reading subject_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subject_name: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectCertificateTemplateType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectCertificateTemplate-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectCertificateTemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCertificateTemplateCity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateCurveName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateDigestType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateKeySize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateKeyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateOrganization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateOrganizationUnit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCertificateTemplateScepPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCertificateTemplateScepServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateSubjectName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCertificateTemplate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("city"); ok {
		t, err := expandObjectCertificateTemplateCity(d, v, "city")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["city"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok {
		t, err := expandObjectCertificateTemplateCountry(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("curve_name"); ok {
		t, err := expandObjectCertificateTemplateCurveName(d, v, "curve_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["curve-name"] = t
		}
	}

	if v, ok := d.GetOk("digest_type"); ok {
		t, err := expandObjectCertificateTemplateDigestType(d, v, "digest_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digest-type"] = t
		}
	}

	if v, ok := d.GetOk("email"); ok {
		t, err := expandObjectCertificateTemplateEmail(d, v, "email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email"] = t
		}
	}

	if v, ok := d.GetOk("id_type"); ok {
		t, err := expandObjectCertificateTemplateIdType(d, v, "id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id-type"] = t
		}
	}

	if v, ok := d.GetOk("key_size"); ok {
		t, err := expandObjectCertificateTemplateKeySize(d, v, "key_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-size"] = t
		}
	}

	if v, ok := d.GetOk("key_type"); ok {
		t, err := expandObjectCertificateTemplateKeyType(d, v, "key_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectCertificateTemplateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("organization"); ok {
		t, err := expandObjectCertificateTemplateOrganization(d, v, "organization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["organization"] = t
		}
	}

	if v, ok := d.GetOk("organization_unit"); ok {
		t, err := expandObjectCertificateTemplateOrganizationUnit(d, v, "organization_unit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["organization-unit"] = t
		}
	}

	if v, ok := d.GetOk("scep_password"); ok {
		t, err := expandObjectCertificateTemplateScepPassword(d, v, "scep_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-password"] = t
		}
	}

	if v, ok := d.GetOk("scep_server"); ok {
		t, err := expandObjectCertificateTemplateScepServer(d, v, "scep_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-server"] = t
		}
	}

	if v, ok := d.GetOk("state"); ok {
		t, err := expandObjectCertificateTemplateState(d, v, "state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["state"] = t
		}
	}

	if v, ok := d.GetOk("subject_name"); ok {
		t, err := expandObjectCertificateTemplateSubjectName(d, v, "subject_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subject-name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandObjectCertificateTemplateType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
