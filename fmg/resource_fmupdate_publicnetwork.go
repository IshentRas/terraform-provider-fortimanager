// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Enable/disable access to the public FortiGuard.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdatePublicnetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdatePublicnetworkUpdate,
		Read:   resourceFmupdatePublicnetworkRead,
		Update: resourceFmupdatePublicnetworkUpdate,
		Delete: resourceFmupdatePublicnetworkDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdatePublicnetworkUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdatePublicnetwork(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdatePublicnetwork resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdatePublicnetwork(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdatePublicnetwork resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdatePublicnetwork")

	return resourceFmupdatePublicnetworkRead(d, m)
}

func resourceFmupdatePublicnetworkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdatePublicnetwork(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdatePublicnetwork resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdatePublicnetworkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdatePublicnetwork(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdatePublicnetwork resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdatePublicnetwork(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdatePublicnetwork resource from API: %v", err)
	}
	return nil
}

func flattenFmupdatePublicnetworkStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdatePublicnetwork(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenFmupdatePublicnetworkStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FmupdatePublicnetwork-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFmupdatePublicnetworkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdatePublicnetworkStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdatePublicnetwork(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFmupdatePublicnetworkStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
