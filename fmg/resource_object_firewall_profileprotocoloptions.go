// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure protocol options.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallProfileProtocolOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallProfileProtocolOptionsCreate,
		Read:   resourceObjectFirewallProfileProtocolOptionsRead,
		Update: resourceObjectFirewallProfileProtocolOptionsUpdate,
		Delete: resourceObjectFirewallProfileProtocolOptionsDelete,

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
			"cifs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_controller": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"file_filter": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entries": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"comment": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"direction": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"file_type": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
													Computed: true,
												},
												"filter": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"protocol": &schema.Schema{
													Type:     schema.TypeSet,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_keytab": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"keytab": &schema.Schema{
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
									"principal": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_window_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comfort_amount": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"comfort_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stream_based_uncompressed_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"block_page_status_code": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"comfort_amount": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"comfort_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fortinet_bar": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortinet_bar_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"post_lang": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"range_block": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"retry_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stream_based_uncompressed_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"streaming_content_bypass": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"strip_x_forwarded_for": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"switching_protocols": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_window_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_non_http": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"unknown_http_version": &schema.Schema{
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
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mail_signature": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"signature": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
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
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
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
			"nntp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"oversize_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pop3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
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
			"rpc_over_http": &schema.Schema{
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
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_busy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssh": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comfort_amount": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"comfort_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stream_based_uncompressed_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"tcp_window_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"switching_protocols_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallProfileProtocolOptionsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallProfileProtocolOptions(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallProfileProtocolOptions resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallProfileProtocolOptions(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallProfileProtocolOptions resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallProfileProtocolOptionsRead(d, m)
}

func resourceObjectFirewallProfileProtocolOptionsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallProfileProtocolOptions(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptions resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallProfileProtocolOptions(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallProfileProtocolOptionsRead(d, m)
}

func resourceObjectFirewallProfileProtocolOptionsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallProfileProtocolOptions(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallProfileProtocolOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallProfileProtocolOptionsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallProfileProtocolOptions(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallProfileProtocolOptions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptions resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallProfileProtocolOptionsCifs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "domain_controller"
	if _, ok := i["domain-controller"]; ok {
		result["domain_controller"] = flattenObjectFirewallProfileProtocolOptionsCifsDomainController(i["domain-controller"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_filter"
	if _, ok := i["file-filter"]; ok {
		result["file_filter"] = flattenObjectFirewallProfileProtocolOptionsCifsFileFilter(i["file-filter"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectFirewallProfileProtocolOptionsCifsOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsCifsOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallProfileProtocolOptionsCifsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenObjectFirewallProfileProtocolOptionsCifsScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_credential_type"
	if _, ok := i["server-credential-type"]; ok {
		result["server_credential_type"] = flattenObjectFirewallProfileProtocolOptionsCifsServerCredentialType(i["server-credential-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_keytab"
	if _, ok := i["server-keytab"]; ok {
		result["server_keytab"] = flattenObjectFirewallProfileProtocolOptionsCifsServerKeytab(i["server-keytab"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallProfileProtocolOptionsCifsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := i["tcp-window-maximum"]; ok {
		result["tcp_window_maximum"] = flattenObjectFirewallProfileProtocolOptionsCifsTcpWindowMaximum(i["tcp-window-maximum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := i["tcp-window-minimum"]; ok {
		result["tcp_window_minimum"] = flattenObjectFirewallProfileProtocolOptionsCifsTcpWindowMinimum(i["tcp-window-minimum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := i["tcp-window-size"]; ok {
		result["tcp_window_size"] = flattenObjectFirewallProfileProtocolOptionsCifsTcpWindowSize(i["tcp-window-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := i["tcp-window-type"]; ok {
		result["tcp_window_type"] = flattenObjectFirewallProfileProtocolOptionsCifsTcpWindowType(i["tcp-window-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenObjectFirewallProfileProtocolOptionsCifsUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallProfileProtocolOptionsCifsDomainController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsFileFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "entries"
	if _, ok := i["entries"]; ok {
		result["entries"] = flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntries(i["entries"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenObjectFirewallProfileProtocolOptionsCifsFileFilterLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallProfileProtocolOptionsCifsFileFilterStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectFirewallProfileProtocolOptionsCifsFileFilter-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "ObjectFirewallProfileProtocolOptionsCifsFileFilter-Entries-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesDirection(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "ObjectFirewallProfileProtocolOptionsCifsFileFilter-Entries-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := i["file-type"]; ok {
			v := flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesFileType(i["file-type"], d, pre_append)
			tmp["file_type"] = fortiAPISubPartPatch(v, "ObjectFirewallProfileProtocolOptionsCifsFileFilter-Entries-FileType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesFilter(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "ObjectFirewallProfileProtocolOptionsCifsFileFilter-Entries-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectFirewallProfileProtocolOptionsCifsFileFilter-Entries-Protocol")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesFileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsCifsFileFilterLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsFileFilterStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsCifsOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallProfileProtocolOptionsCifsScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsServerCredentialType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsServerKeytab(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keytab"
		if _, ok := i["keytab"]; ok {
			v := flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabKeytab(i["keytab"], d, pre_append)
			tmp["keytab"] = fortiAPISubPartPatch(v, "ObjectFirewallProfileProtocolOptionsCifs-ServerKeytab-Keytab")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			v := flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabPassword(i["password"], d, pre_append)
			tmp["password"] = fortiAPISubPartPatch(v, "ObjectFirewallProfileProtocolOptionsCifs-ServerKeytab-Password")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "principal"
		if _, ok := i["principal"]; ok {
			v := flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabPrincipal(i["principal"], d, pre_append)
			tmp["principal"] = fortiAPISubPartPatch(v, "ObjectFirewallProfileProtocolOptionsCifs-ServerKeytab-Principal")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabKeytab(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabPrincipal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsTcpWindowMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsTcpWindowMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsTcpWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsTcpWindowType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsDns(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallProfileProtocolOptionsDnsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallProfileProtocolOptionsDnsStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallProfileProtocolOptionsDnsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallProfileProtocolOptionsDnsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := i["comfort-amount"]; ok {
		result["comfort_amount"] = flattenObjectFirewallProfileProtocolOptionsFtpComfortAmount(i["comfort-amount"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := i["comfort-interval"]; ok {
		result["comfort_interval"] = flattenObjectFirewallProfileProtocolOptionsFtpComfortInterval(i["comfort-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenObjectFirewallProfileProtocolOptionsFtpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectFirewallProfileProtocolOptionsFtpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsFtpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallProfileProtocolOptionsFtpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenObjectFirewallProfileProtocolOptionsFtpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = flattenObjectFirewallProfileProtocolOptionsFtpSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallProfileProtocolOptionsFtpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := i["stream-based-uncompressed-limit"]; ok {
		result["stream_based_uncompressed_limit"] = flattenObjectFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit(i["stream-based-uncompressed-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := i["tcp-window-maximum"]; ok {
		result["tcp_window_maximum"] = flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowMaximum(i["tcp-window-maximum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := i["tcp-window-minimum"]; ok {
		result["tcp_window_minimum"] = flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowMinimum(i["tcp-window-minimum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := i["tcp-window-size"]; ok {
		result["tcp_window_size"] = flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowSize(i["tcp-window-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := i["tcp-window-type"]; ok {
		result["tcp_window_type"] = flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowType(i["tcp-window-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenObjectFirewallProfileProtocolOptionsFtpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallProfileProtocolOptionsFtpComfortAmount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpComfortInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsFtpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallProfileProtocolOptionsFtpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "block_page_status_code"
	if _, ok := i["block-page-status-code"]; ok {
		result["block_page_status_code"] = flattenObjectFirewallProfileProtocolOptionsHttpBlockPageStatusCode(i["block-page-status-code"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := i["comfort-amount"]; ok {
		result["comfort_amount"] = flattenObjectFirewallProfileProtocolOptionsHttpComfortAmount(i["comfort-amount"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := i["comfort-interval"]; ok {
		result["comfort_interval"] = flattenObjectFirewallProfileProtocolOptionsHttpComfortInterval(i["comfort-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortinet_bar"
	if _, ok := i["fortinet-bar"]; ok {
		result["fortinet_bar"] = flattenObjectFirewallProfileProtocolOptionsHttpFortinetBar(i["fortinet-bar"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortinet_bar_port"
	if _, ok := i["fortinet-bar-port"]; ok {
		result["fortinet_bar_port"] = flattenObjectFirewallProfileProtocolOptionsHttpFortinetBarPort(i["fortinet-bar-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenObjectFirewallProfileProtocolOptionsHttpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectFirewallProfileProtocolOptionsHttpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsHttpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallProfileProtocolOptionsHttpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "post_lang"
	if _, ok := i["post-lang"]; ok {
		result["post_lang"] = flattenObjectFirewallProfileProtocolOptionsHttpPostLang(i["post-lang"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenObjectFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "range_block"
	if _, ok := i["range-block"]; ok {
		result["range_block"] = flattenObjectFirewallProfileProtocolOptionsHttpRangeBlock(i["range-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "retry_count"
	if _, ok := i["retry-count"]; ok {
		result["retry_count"] = flattenObjectFirewallProfileProtocolOptionsHttpRetryCount(i["retry-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenObjectFirewallProfileProtocolOptionsHttpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = flattenObjectFirewallProfileProtocolOptionsHttpSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallProfileProtocolOptionsHttpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := i["stream-based-uncompressed-limit"]; ok {
		result["stream_based_uncompressed_limit"] = flattenObjectFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit(i["stream-based-uncompressed-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "streaming_content_bypass"
	if _, ok := i["streaming-content-bypass"]; ok {
		result["streaming_content_bypass"] = flattenObjectFirewallProfileProtocolOptionsHttpStreamingContentBypass(i["streaming-content-bypass"], d, pre_append)
	}

	pre_append = pre + ".0." + "strip_x_forwarded_for"
	if _, ok := i["strip-x-forwarded-for"]; ok {
		result["strip_x_forwarded_for"] = flattenObjectFirewallProfileProtocolOptionsHttpStripXForwardedFor(i["strip-x-forwarded-for"], d, pre_append)
	}

	pre_append = pre + ".0." + "switching_protocols"
	if _, ok := i["switching-protocols"]; ok {
		result["switching_protocols"] = flattenObjectFirewallProfileProtocolOptionsHttpSwitchingProtocols(i["switching-protocols"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := i["tcp-window-maximum"]; ok {
		result["tcp_window_maximum"] = flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowMaximum(i["tcp-window-maximum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := i["tcp-window-minimum"]; ok {
		result["tcp_window_minimum"] = flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowMinimum(i["tcp-window-minimum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := i["tcp-window-size"]; ok {
		result["tcp_window_size"] = flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowSize(i["tcp-window-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := i["tcp-window-type"]; ok {
		result["tcp_window_type"] = flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowType(i["tcp-window-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_non_http"
	if _, ok := i["tunnel-non-http"]; ok {
		result["tunnel_non_http"] = flattenObjectFirewallProfileProtocolOptionsHttpTunnelNonHttp(i["tunnel-non-http"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenObjectFirewallProfileProtocolOptionsHttpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "unknown_http_version"
	if _, ok := i["unknown-http-version"]; ok {
		result["unknown_http_version"] = flattenObjectFirewallProfileProtocolOptionsHttpUnknownHttpVersion(i["unknown-http-version"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallProfileProtocolOptionsHttpBlockPageStatusCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpComfortAmount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpComfortInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpFortinetBar(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpFortinetBarPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsHttpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallProfileProtocolOptionsHttpPostLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpRangeBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpRetryCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpStreamingContentBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpStripXForwardedFor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpSwitchingProtocols(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpTcpWindowType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpTunnelNonHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsHttpUnknownHttpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsImap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenObjectFirewallProfileProtocolOptionsImapInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectFirewallProfileProtocolOptionsImapOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsImapOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallProfileProtocolOptionsImapPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenObjectFirewallProfileProtocolOptionsImapProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenObjectFirewallProfileProtocolOptionsImapScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = flattenObjectFirewallProfileProtocolOptionsImapSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallProfileProtocolOptionsImapStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenObjectFirewallProfileProtocolOptionsImapUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallProfileProtocolOptionsImapInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsImapOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsImapOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsImapPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallProfileProtocolOptionsImapProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsImapScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsImapSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsImapStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsImapUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsMailSignature(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "signature"
	if _, ok := i["signature"]; ok {
		result["signature"] = flattenObjectFirewallProfileProtocolOptionsMailSignatureSignature(i["signature"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallProfileProtocolOptionsMailSignatureStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallProfileProtocolOptionsMailSignatureSignature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsMailSignatureStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsMapi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectFirewallProfileProtocolOptionsMapiOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsMapiOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallProfileProtocolOptionsMapiPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenObjectFirewallProfileProtocolOptionsMapiScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallProfileProtocolOptionsMapiStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenObjectFirewallProfileProtocolOptionsMapiUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallProfileProtocolOptionsMapiOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsMapiOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsMapiPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallProfileProtocolOptionsMapiScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsMapiStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsMapiUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsNntp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenObjectFirewallProfileProtocolOptionsNntpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectFirewallProfileProtocolOptionsNntpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsNntpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallProfileProtocolOptionsNntpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenObjectFirewallProfileProtocolOptionsNntpProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenObjectFirewallProfileProtocolOptionsNntpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallProfileProtocolOptionsNntpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenObjectFirewallProfileProtocolOptionsNntpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallProfileProtocolOptionsNntpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsNntpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsNntpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsNntpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallProfileProtocolOptionsNntpProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsNntpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsNntpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsNntpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsOversizeLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsPop3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenObjectFirewallProfileProtocolOptionsPop3InspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectFirewallProfileProtocolOptionsPop3Options(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsPop3OversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallProfileProtocolOptionsPop3Ports(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenObjectFirewallProfileProtocolOptionsPop3ProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenObjectFirewallProfileProtocolOptionsPop3ScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = flattenObjectFirewallProfileProtocolOptionsPop3SslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallProfileProtocolOptionsPop3Status(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenObjectFirewallProfileProtocolOptionsPop3UncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallProfileProtocolOptionsPop3InspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsPop3Options(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsPop3OversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsPop3Ports(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallProfileProtocolOptionsPop3ProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsPop3ScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsPop3SslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsPop3Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsPop3UncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsRpcOverHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSmtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenObjectFirewallProfileProtocolOptionsSmtpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectFirewallProfileProtocolOptionsSmtpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsSmtpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenObjectFirewallProfileProtocolOptionsSmtpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenObjectFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenObjectFirewallProfileProtocolOptionsSmtpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_busy"
	if _, ok := i["server-busy"]; ok {
		result["server_busy"] = flattenObjectFirewallProfileProtocolOptionsSmtpServerBusy(i["server-busy"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = flattenObjectFirewallProfileProtocolOptionsSmtpSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectFirewallProfileProtocolOptionsSmtpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenObjectFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallProfileProtocolOptionsSmtpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSmtpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsSmtpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSmtpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSmtpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSmtpServerBusy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSmtpSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSmtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSsh(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := i["comfort-amount"]; ok {
		result["comfort_amount"] = flattenObjectFirewallProfileProtocolOptionsSshComfortAmount(i["comfort-amount"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := i["comfort-interval"]; ok {
		result["comfort_interval"] = flattenObjectFirewallProfileProtocolOptionsSshComfortInterval(i["comfort-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenObjectFirewallProfileProtocolOptionsSshOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsSshOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenObjectFirewallProfileProtocolOptionsSshScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = flattenObjectFirewallProfileProtocolOptionsSshSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := i["stream-based-uncompressed-limit"]; ok {
		result["stream_based_uncompressed_limit"] = flattenObjectFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit(i["stream-based-uncompressed-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := i["tcp-window-maximum"]; ok {
		result["tcp_window_maximum"] = flattenObjectFirewallProfileProtocolOptionsSshTcpWindowMaximum(i["tcp-window-maximum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := i["tcp-window-minimum"]; ok {
		result["tcp_window_minimum"] = flattenObjectFirewallProfileProtocolOptionsSshTcpWindowMinimum(i["tcp-window-minimum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := i["tcp-window-size"]; ok {
		result["tcp_window_size"] = flattenObjectFirewallProfileProtocolOptionsSshTcpWindowSize(i["tcp-window-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := i["tcp-window-type"]; ok {
		result["tcp_window_type"] = flattenObjectFirewallProfileProtocolOptionsSshTcpWindowType(i["tcp-window-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenObjectFirewallProfileProtocolOptionsSshUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenObjectFirewallProfileProtocolOptionsSshUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectFirewallProfileProtocolOptionsSshComfortAmount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSshComfortInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSshOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsSshOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSshScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSshSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSshTcpWindowMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSshTcpWindowMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSshTcpWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSshTcpWindowType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSshUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSshUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsSwitchingProtocolsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallProfileProtocolOptions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("cifs", flattenObjectFirewallProfileProtocolOptionsCifs(o["cifs"], d, "cifs")); err != nil {
			if vv, ok := fortiAPIPatch(o["cifs"], "ObjectFirewallProfileProtocolOptions-Cifs"); ok {
				if err = d.Set("cifs", vv); err != nil {
					return fmt.Errorf("Error reading cifs: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cifs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cifs"); ok {
			if err = d.Set("cifs", flattenObjectFirewallProfileProtocolOptionsCifs(o["cifs"], d, "cifs")); err != nil {
				if vv, ok := fortiAPIPatch(o["cifs"], "ObjectFirewallProfileProtocolOptions-Cifs"); ok {
					if err = d.Set("cifs", vv); err != nil {
						return fmt.Errorf("Error reading cifs: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cifs: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenObjectFirewallProfileProtocolOptionsComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallProfileProtocolOptions-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dns", flattenObjectFirewallProfileProtocolOptionsDns(o["dns"], d, "dns")); err != nil {
			if vv, ok := fortiAPIPatch(o["dns"], "ObjectFirewallProfileProtocolOptions-Dns"); ok {
				if err = d.Set("dns", vv); err != nil {
					return fmt.Errorf("Error reading dns: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dns"); ok {
			if err = d.Set("dns", flattenObjectFirewallProfileProtocolOptionsDns(o["dns"], d, "dns")); err != nil {
				if vv, ok := fortiAPIPatch(o["dns"], "ObjectFirewallProfileProtocolOptions-Dns"); ok {
					if err = d.Set("dns", vv); err != nil {
						return fmt.Errorf("Error reading dns: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dns: %v", err)
				}
			}
		}
	}

	if err = d.Set("feature_set", flattenObjectFirewallProfileProtocolOptionsFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["feature-set"], "ObjectFirewallProfileProtocolOptions-FeatureSet"); ok {
			if err = d.Set("feature_set", vv); err != nil {
				return fmt.Errorf("Error reading feature_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ftp", flattenObjectFirewallProfileProtocolOptionsFtp(o["ftp"], d, "ftp")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftp"], "ObjectFirewallProfileProtocolOptions-Ftp"); ok {
				if err = d.Set("ftp", vv); err != nil {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftp"); ok {
			if err = d.Set("ftp", flattenObjectFirewallProfileProtocolOptionsFtp(o["ftp"], d, "ftp")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftp"], "ObjectFirewallProfileProtocolOptions-Ftp"); ok {
					if err = d.Set("ftp", vv); err != nil {
						return fmt.Errorf("Error reading ftp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("http", flattenObjectFirewallProfileProtocolOptionsHttp(o["http"], d, "http")); err != nil {
			if vv, ok := fortiAPIPatch(o["http"], "ObjectFirewallProfileProtocolOptions-Http"); ok {
				if err = d.Set("http", vv); err != nil {
					return fmt.Errorf("Error reading http: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading http: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("http"); ok {
			if err = d.Set("http", flattenObjectFirewallProfileProtocolOptionsHttp(o["http"], d, "http")); err != nil {
				if vv, ok := fortiAPIPatch(o["http"], "ObjectFirewallProfileProtocolOptions-Http"); ok {
					if err = d.Set("http", vv); err != nil {
						return fmt.Errorf("Error reading http: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading http: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("imap", flattenObjectFirewallProfileProtocolOptionsImap(o["imap"], d, "imap")); err != nil {
			if vv, ok := fortiAPIPatch(o["imap"], "ObjectFirewallProfileProtocolOptions-Imap"); ok {
				if err = d.Set("imap", vv); err != nil {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading imap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imap"); ok {
			if err = d.Set("imap", flattenObjectFirewallProfileProtocolOptionsImap(o["imap"], d, "imap")); err != nil {
				if vv, ok := fortiAPIPatch(o["imap"], "ObjectFirewallProfileProtocolOptions-Imap"); ok {
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
		if err = d.Set("mail_signature", flattenObjectFirewallProfileProtocolOptionsMailSignature(o["mail-signature"], d, "mail_signature")); err != nil {
			if vv, ok := fortiAPIPatch(o["mail-signature"], "ObjectFirewallProfileProtocolOptions-MailSignature"); ok {
				if err = d.Set("mail_signature", vv); err != nil {
					return fmt.Errorf("Error reading mail_signature: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mail_signature: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mail_signature"); ok {
			if err = d.Set("mail_signature", flattenObjectFirewallProfileProtocolOptionsMailSignature(o["mail-signature"], d, "mail_signature")); err != nil {
				if vv, ok := fortiAPIPatch(o["mail-signature"], "ObjectFirewallProfileProtocolOptions-MailSignature"); ok {
					if err = d.Set("mail_signature", vv); err != nil {
						return fmt.Errorf("Error reading mail_signature: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mail_signature: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mapi", flattenObjectFirewallProfileProtocolOptionsMapi(o["mapi"], d, "mapi")); err != nil {
			if vv, ok := fortiAPIPatch(o["mapi"], "ObjectFirewallProfileProtocolOptions-Mapi"); ok {
				if err = d.Set("mapi", vv); err != nil {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mapi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mapi"); ok {
			if err = d.Set("mapi", flattenObjectFirewallProfileProtocolOptionsMapi(o["mapi"], d, "mapi")); err != nil {
				if vv, ok := fortiAPIPatch(o["mapi"], "ObjectFirewallProfileProtocolOptions-Mapi"); ok {
					if err = d.Set("mapi", vv); err != nil {
						return fmt.Errorf("Error reading mapi: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectFirewallProfileProtocolOptionsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallProfileProtocolOptions-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nntp", flattenObjectFirewallProfileProtocolOptionsNntp(o["nntp"], d, "nntp")); err != nil {
			if vv, ok := fortiAPIPatch(o["nntp"], "ObjectFirewallProfileProtocolOptions-Nntp"); ok {
				if err = d.Set("nntp", vv); err != nil {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nntp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nntp"); ok {
			if err = d.Set("nntp", flattenObjectFirewallProfileProtocolOptionsNntp(o["nntp"], d, "nntp")); err != nil {
				if vv, ok := fortiAPIPatch(o["nntp"], "ObjectFirewallProfileProtocolOptions-Nntp"); ok {
					if err = d.Set("nntp", vv); err != nil {
						return fmt.Errorf("Error reading nntp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			}
		}
	}

	if err = d.Set("oversize_log", flattenObjectFirewallProfileProtocolOptionsOversizeLog(o["oversize-log"], d, "oversize_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["oversize-log"], "ObjectFirewallProfileProtocolOptions-OversizeLog"); ok {
			if err = d.Set("oversize_log", vv); err != nil {
				return fmt.Errorf("Error reading oversize_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oversize_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("pop3", flattenObjectFirewallProfileProtocolOptionsPop3(o["pop3"], d, "pop3")); err != nil {
			if vv, ok := fortiAPIPatch(o["pop3"], "ObjectFirewallProfileProtocolOptions-Pop3"); ok {
				if err = d.Set("pop3", vv); err != nil {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pop3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3"); ok {
			if err = d.Set("pop3", flattenObjectFirewallProfileProtocolOptionsPop3(o["pop3"], d, "pop3")); err != nil {
				if vv, ok := fortiAPIPatch(o["pop3"], "ObjectFirewallProfileProtocolOptions-Pop3"); ok {
					if err = d.Set("pop3", vv); err != nil {
						return fmt.Errorf("Error reading pop3: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			}
		}
	}

	if err = d.Set("replacemsg_group", flattenObjectFirewallProfileProtocolOptionsReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "ObjectFirewallProfileProtocolOptions-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("rpc_over_http", flattenObjectFirewallProfileProtocolOptionsRpcOverHttp(o["rpc-over-http"], d, "rpc_over_http")); err != nil {
		if vv, ok := fortiAPIPatch(o["rpc-over-http"], "ObjectFirewallProfileProtocolOptions-RpcOverHttp"); ok {
			if err = d.Set("rpc_over_http", vv); err != nil {
				return fmt.Errorf("Error reading rpc_over_http: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rpc_over_http: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("smtp", flattenObjectFirewallProfileProtocolOptionsSmtp(o["smtp"], d, "smtp")); err != nil {
			if vv, ok := fortiAPIPatch(o["smtp"], "ObjectFirewallProfileProtocolOptions-Smtp"); ok {
				if err = d.Set("smtp", vv); err != nil {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading smtp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtp"); ok {
			if err = d.Set("smtp", flattenObjectFirewallProfileProtocolOptionsSmtp(o["smtp"], d, "smtp")); err != nil {
				if vv, ok := fortiAPIPatch(o["smtp"], "ObjectFirewallProfileProtocolOptions-Smtp"); ok {
					if err = d.Set("smtp", vv); err != nil {
						return fmt.Errorf("Error reading smtp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ssh", flattenObjectFirewallProfileProtocolOptionsSsh(o["ssh"], d, "ssh")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssh"], "ObjectFirewallProfileProtocolOptions-Ssh"); ok {
				if err = d.Set("ssh", vv); err != nil {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssh: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssh"); ok {
			if err = d.Set("ssh", flattenObjectFirewallProfileProtocolOptionsSsh(o["ssh"], d, "ssh")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssh"], "ObjectFirewallProfileProtocolOptions-Ssh"); ok {
					if err = d.Set("ssh", vv); err != nil {
						return fmt.Errorf("Error reading ssh: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			}
		}
	}

	if err = d.Set("switching_protocols_log", flattenObjectFirewallProfileProtocolOptionsSwitchingProtocolsLog(o["switching-protocols-log"], d, "switching_protocols_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["switching-protocols-log"], "ObjectFirewallProfileProtocolOptions-SwitchingProtocolsLog"); ok {
			if err = d.Set("switching_protocols_log", vv); err != nil {
				return fmt.Errorf("Error reading switching_protocols_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switching_protocols_log: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallProfileProtocolOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallProfileProtocolOptionsCifs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "domain_controller"
	if _, ok := d.GetOk(pre_append); ok {
		result["domain-controller"], _ = expandObjectFirewallProfileProtocolOptionsCifsDomainController(d, i["domain_controller"], pre_append)
	}
	pre_append = pre + ".0." + "file_filter"
	if _, ok := d.GetOk(pre_append); ok {
		result["file-filter"], _ = expandObjectFirewallProfileProtocolOptionsCifsFileFilter(d, i["file_filter"], pre_append)
	} else {
		result["file-filter"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectFirewallProfileProtocolOptionsCifsOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsCifsOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallProfileProtocolOptionsCifsPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandObjectFirewallProfileProtocolOptionsCifsScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "server_credential_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["server-credential-type"], _ = expandObjectFirewallProfileProtocolOptionsCifsServerCredentialType(d, i["server_credential_type"], pre_append)
	}
	pre_append = pre + ".0." + "server_keytab"
	if _, ok := d.GetOk(pre_append); ok {
		result["server-keytab"], _ = expandObjectFirewallProfileProtocolOptionsCifsServerKeytab(d, i["server_keytab"], pre_append)
	} else {
		result["server-keytab"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallProfileProtocolOptionsCifsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-maximum"], _ = expandObjectFirewallProfileProtocolOptionsCifsTcpWindowMaximum(d, i["tcp_window_maximum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-minimum"], _ = expandObjectFirewallProfileProtocolOptionsCifsTcpWindowMinimum(d, i["tcp_window_minimum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-size"], _ = expandObjectFirewallProfileProtocolOptionsCifsTcpWindowSize(d, i["tcp_window_size"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-type"], _ = expandObjectFirewallProfileProtocolOptionsCifsTcpWindowType(d, i["tcp_window_type"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandObjectFirewallProfileProtocolOptionsCifsUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsDomainController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsFileFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "entries"
	if _, ok := d.GetOk(pre_append); ok {
		result["entries"], _ = expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntries(d, i["entries"], pre_append)
	} else {
		result["entries"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandObjectFirewallProfileProtocolOptionsCifsFileFilterLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallProfileProtocolOptionsCifsFileFilterStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["direction"], _ = expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["file-type"], _ = expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesFileType(d, i["file_type"], pre_append)
		} else {
			tmp["file-type"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter"], _ = expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesFilter(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesProtocol(d, i["protocol"], pre_append)
		} else {
			tmp["protocol"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesFileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsFileFilterEntriesProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsCifsFileFilterLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsFileFilterStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsCifsOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsCifsScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsServerCredentialType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsServerKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keytab"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["keytab"], _ = expandObjectFirewallProfileProtocolOptionsCifsServerKeytabKeytab(d, i["keytab"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password"], _ = expandObjectFirewallProfileProtocolOptionsCifsServerKeytabPassword(d, i["password"], pre_append)
		} else {
			tmp["password"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "principal"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["principal"], _ = expandObjectFirewallProfileProtocolOptionsCifsServerKeytabPrincipal(d, i["principal"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsServerKeytabKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsServerKeytabPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsCifsServerKeytabPrincipal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsTcpWindowMaximum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsTcpWindowMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsTcpWindowSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsTcpWindowType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallProfileProtocolOptionsDnsPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallProfileProtocolOptionsDnsStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsDnsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsDnsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFeatureSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := d.GetOk(pre_append); ok {
		result["comfort-amount"], _ = expandObjectFirewallProfileProtocolOptionsFtpComfortAmount(d, i["comfort_amount"], pre_append)
	}
	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["comfort-interval"], _ = expandObjectFirewallProfileProtocolOptionsFtpComfortInterval(d, i["comfort_interval"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandObjectFirewallProfileProtocolOptionsFtpInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectFirewallProfileProtocolOptionsFtpOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsFtpOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallProfileProtocolOptionsFtpPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandObjectFirewallProfileProtocolOptionsFtpScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-offloaded"], _ = expandObjectFirewallProfileProtocolOptionsFtpSslOffloaded(d, i["ssl_offloaded"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallProfileProtocolOptionsFtpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["stream-based-uncompressed-limit"], _ = expandObjectFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit(d, i["stream_based_uncompressed_limit"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-maximum"], _ = expandObjectFirewallProfileProtocolOptionsFtpTcpWindowMaximum(d, i["tcp_window_maximum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-minimum"], _ = expandObjectFirewallProfileProtocolOptionsFtpTcpWindowMinimum(d, i["tcp_window_minimum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-size"], _ = expandObjectFirewallProfileProtocolOptionsFtpTcpWindowSize(d, i["tcp_window_size"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-type"], _ = expandObjectFirewallProfileProtocolOptionsFtpTcpWindowType(d, i["tcp_window_type"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandObjectFirewallProfileProtocolOptionsFtpUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpComfortAmount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpComfortInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsFtpOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsFtpScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpSslOffloaded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpTcpWindowMaximum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpTcpWindowMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpTcpWindowSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpTcpWindowType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "block_page_status_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-page-status-code"], _ = expandObjectFirewallProfileProtocolOptionsHttpBlockPageStatusCode(d, i["block_page_status_code"], pre_append)
	}
	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := d.GetOk(pre_append); ok {
		result["comfort-amount"], _ = expandObjectFirewallProfileProtocolOptionsHttpComfortAmount(d, i["comfort_amount"], pre_append)
	}
	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["comfort-interval"], _ = expandObjectFirewallProfileProtocolOptionsHttpComfortInterval(d, i["comfort_interval"], pre_append)
	}
	pre_append = pre + ".0." + "fortinet_bar"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortinet-bar"], _ = expandObjectFirewallProfileProtocolOptionsHttpFortinetBar(d, i["fortinet_bar"], pre_append)
	}
	pre_append = pre + ".0." + "fortinet_bar_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortinet-bar-port"], _ = expandObjectFirewallProfileProtocolOptionsHttpFortinetBarPort(d, i["fortinet_bar_port"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandObjectFirewallProfileProtocolOptionsHttpInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectFirewallProfileProtocolOptionsHttpOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsHttpOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallProfileProtocolOptionsHttpPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "post_lang"
	if _, ok := d.GetOk(pre_append); ok {
		result["post-lang"], _ = expandObjectFirewallProfileProtocolOptionsHttpPostLang(d, i["post_lang"], pre_append)
	} else {
		result["post-lang"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandObjectFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "range_block"
	if _, ok := d.GetOk(pre_append); ok {
		result["range-block"], _ = expandObjectFirewallProfileProtocolOptionsHttpRangeBlock(d, i["range_block"], pre_append)
	}
	pre_append = pre + ".0." + "retry_count"
	if _, ok := d.GetOk(pre_append); ok {
		result["retry-count"], _ = expandObjectFirewallProfileProtocolOptionsHttpRetryCount(d, i["retry_count"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandObjectFirewallProfileProtocolOptionsHttpScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-offloaded"], _ = expandObjectFirewallProfileProtocolOptionsHttpSslOffloaded(d, i["ssl_offloaded"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallProfileProtocolOptionsHttpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["stream-based-uncompressed-limit"], _ = expandObjectFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit(d, i["stream_based_uncompressed_limit"], pre_append)
	}
	pre_append = pre + ".0." + "streaming_content_bypass"
	if _, ok := d.GetOk(pre_append); ok {
		result["streaming-content-bypass"], _ = expandObjectFirewallProfileProtocolOptionsHttpStreamingContentBypass(d, i["streaming_content_bypass"], pre_append)
	}
	pre_append = pre + ".0." + "strip_x_forwarded_for"
	if _, ok := d.GetOk(pre_append); ok {
		result["strip-x-forwarded-for"], _ = expandObjectFirewallProfileProtocolOptionsHttpStripXForwardedFor(d, i["strip_x_forwarded_for"], pre_append)
	}
	pre_append = pre + ".0." + "switching_protocols"
	if _, ok := d.GetOk(pre_append); ok {
		result["switching-protocols"], _ = expandObjectFirewallProfileProtocolOptionsHttpSwitchingProtocols(d, i["switching_protocols"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-maximum"], _ = expandObjectFirewallProfileProtocolOptionsHttpTcpWindowMaximum(d, i["tcp_window_maximum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-minimum"], _ = expandObjectFirewallProfileProtocolOptionsHttpTcpWindowMinimum(d, i["tcp_window_minimum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-size"], _ = expandObjectFirewallProfileProtocolOptionsHttpTcpWindowSize(d, i["tcp_window_size"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-type"], _ = expandObjectFirewallProfileProtocolOptionsHttpTcpWindowType(d, i["tcp_window_type"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_non_http"
	if _, ok := d.GetOk(pre_append); ok {
		result["tunnel-non-http"], _ = expandObjectFirewallProfileProtocolOptionsHttpTunnelNonHttp(d, i["tunnel_non_http"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandObjectFirewallProfileProtocolOptionsHttpUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "unknown_http_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["unknown-http-version"], _ = expandObjectFirewallProfileProtocolOptionsHttpUnknownHttpVersion(d, i["unknown_http_version"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpBlockPageStatusCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpComfortAmount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpComfortInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpFortinetBar(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpFortinetBarPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsHttpOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsHttpPostLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpRangeBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpRetryCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpSslOffloaded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpStreamingContentBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpStripXForwardedFor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpSwitchingProtocols(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpTcpWindowMaximum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpTcpWindowMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpTcpWindowSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpTcpWindowType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpTunnelNonHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsHttpUnknownHttpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsImap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandObjectFirewallProfileProtocolOptionsImapInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectFirewallProfileProtocolOptionsImapOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsImapOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallProfileProtocolOptionsImapPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandObjectFirewallProfileProtocolOptionsImapProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandObjectFirewallProfileProtocolOptionsImapScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-offloaded"], _ = expandObjectFirewallProfileProtocolOptionsImapSslOffloaded(d, i["ssl_offloaded"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallProfileProtocolOptionsImapStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandObjectFirewallProfileProtocolOptionsImapUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsImapInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsImapOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsImapOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsImapPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsImapProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsImapScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsImapSslOffloaded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsImapStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsImapUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsMailSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "signature"
	if _, ok := d.GetOk(pre_append); ok {
		result["signature"], _ = expandObjectFirewallProfileProtocolOptionsMailSignatureSignature(d, i["signature"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallProfileProtocolOptionsMailSignatureStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsMailSignatureSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsMailSignatureStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsMapi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectFirewallProfileProtocolOptionsMapiOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsMapiOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallProfileProtocolOptionsMapiPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandObjectFirewallProfileProtocolOptionsMapiScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallProfileProtocolOptionsMapiStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandObjectFirewallProfileProtocolOptionsMapiUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsMapiOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsMapiOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsMapiPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsMapiScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsMapiStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsMapiUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsNntp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandObjectFirewallProfileProtocolOptionsNntpInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectFirewallProfileProtocolOptionsNntpOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsNntpOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallProfileProtocolOptionsNntpPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandObjectFirewallProfileProtocolOptionsNntpProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandObjectFirewallProfileProtocolOptionsNntpScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallProfileProtocolOptionsNntpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandObjectFirewallProfileProtocolOptionsNntpUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsNntpInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsNntpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsNntpOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsNntpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsNntpProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsNntpScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsNntpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsNntpUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsOversizeLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsPop3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandObjectFirewallProfileProtocolOptionsPop3InspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectFirewallProfileProtocolOptionsPop3Options(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsPop3OversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallProfileProtocolOptionsPop3Ports(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandObjectFirewallProfileProtocolOptionsPop3ProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandObjectFirewallProfileProtocolOptionsPop3ScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-offloaded"], _ = expandObjectFirewallProfileProtocolOptionsPop3SslOffloaded(d, i["ssl_offloaded"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallProfileProtocolOptionsPop3Status(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandObjectFirewallProfileProtocolOptionsPop3UncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsPop3InspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsPop3Options(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsPop3OversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsPop3Ports(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsPop3ProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsPop3ScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsPop3SslOffloaded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsPop3Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsPop3UncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsRpcOverHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSmtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandObjectFirewallProfileProtocolOptionsSmtpInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectFirewallProfileProtocolOptionsSmtpOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsSmtpOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandObjectFirewallProfileProtocolOptionsSmtpPorts(d, i["ports"], pre_append)
	} else {
		result["ports"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandObjectFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandObjectFirewallProfileProtocolOptionsSmtpScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "server_busy"
	if _, ok := d.GetOk(pre_append); ok {
		result["server-busy"], _ = expandObjectFirewallProfileProtocolOptionsSmtpServerBusy(d, i["server_busy"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-offloaded"], _ = expandObjectFirewallProfileProtocolOptionsSmtpSslOffloaded(d, i["ssl_offloaded"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandObjectFirewallProfileProtocolOptionsSmtpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandObjectFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsSmtpInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSmtpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsSmtpOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSmtpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSmtpScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSmtpServerBusy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSmtpSslOffloaded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSmtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := d.GetOk(pre_append); ok {
		result["comfort-amount"], _ = expandObjectFirewallProfileProtocolOptionsSshComfortAmount(d, i["comfort_amount"], pre_append)
	}
	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["comfort-interval"], _ = expandObjectFirewallProfileProtocolOptionsSshComfortInterval(d, i["comfort_interval"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandObjectFirewallProfileProtocolOptionsSshOptions(d, i["options"], pre_append)
	} else {
		result["options"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsSshOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandObjectFirewallProfileProtocolOptionsSshScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-offloaded"], _ = expandObjectFirewallProfileProtocolOptionsSshSslOffloaded(d, i["ssl_offloaded"], pre_append)
	}
	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["stream-based-uncompressed-limit"], _ = expandObjectFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit(d, i["stream_based_uncompressed_limit"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-maximum"], _ = expandObjectFirewallProfileProtocolOptionsSshTcpWindowMaximum(d, i["tcp_window_maximum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-minimum"], _ = expandObjectFirewallProfileProtocolOptionsSshTcpWindowMinimum(d, i["tcp_window_minimum"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-size"], _ = expandObjectFirewallProfileProtocolOptionsSshTcpWindowSize(d, i["tcp_window_size"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-window-type"], _ = expandObjectFirewallProfileProtocolOptionsSshTcpWindowType(d, i["tcp_window_type"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandObjectFirewallProfileProtocolOptionsSshUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandObjectFirewallProfileProtocolOptionsSshUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}

	return result, nil
}

func expandObjectFirewallProfileProtocolOptionsSshComfortAmount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSshComfortInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSshOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsSshOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSshScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSshSslOffloaded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSshTcpWindowMaximum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSshTcpWindowMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSshTcpWindowSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSshTcpWindowType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSshUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSshUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsSwitchingProtocolsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallProfileProtocolOptions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cifs"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsCifs(d, v, "cifs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("dns"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsDns(d, v, "dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("ftp"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsFtp(d, v, "ftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("imap"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsImap(d, v, "imap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap"] = t
		}
	}

	if v, ok := d.GetOk("mail_signature"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsMailSignature(d, v, "mail_signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mail-signature"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsMapi(d, v, "mapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nntp"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsNntp(d, v, "nntp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nntp"] = t
		}
	}

	if v, ok := d.GetOk("oversize_log"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsOversizeLog(d, v, "oversize_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oversize-log"] = t
		}
	}

	if v, ok := d.GetOk("pop3"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsPop3(d, v, "pop3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("rpc_over_http"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsRpcOverHttp(d, v, "rpc_over_http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-over-http"] = t
		}
	}

	if v, ok := d.GetOk("smtp"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsSmtp(d, v, "smtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp"] = t
		}
	}

	if v, ok := d.GetOk("ssh"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsSsh(d, v, "ssh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh"] = t
		}
	}

	if v, ok := d.GetOk("switching_protocols_log"); ok {
		t, err := expandObjectFirewallProfileProtocolOptionsSwitchingProtocolsLog(d, v, "switching_protocols_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switching-protocols-log"] = t
		}
	}

	return &obj, nil
}
