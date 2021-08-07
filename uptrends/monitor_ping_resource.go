package uptrends

import (
	"context"
	uptrends "github.com/wasfree/uptrends-go-sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceMonitorPingSchema() *schema.Resource {
	return &schema.Resource{
		CreateContext: monitorPingCreate,
		ReadContext:   monitorPingRead,
		UpdateContext: monitorPingUpdate,
		DeleteContext: monitorPingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Display name for the Monitor Ping resource.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"network_address": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The network address that should be used to connect to the server or service you want to monitor.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Production",
				Description:  "The monitor mode, either Development, Staging or Production.",
				ValidateFunc: validation.StringInSlice([]string{"Development", "Staging", "Production"}, false),
			},
			"is_active": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Indicates whether the monitor is actively running in the account.",
			},
			"generate_alert": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "When set to false, no alerts will be generated for this monitor in case of an error.",
			},
			"check_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      5,
				Description:  "Numeric value for the time interval between individual checks, in minutes. The maximum value is 240 (4 hours). The minimum value depends on the type of monitor.",
				ValidateFunc: validation.IntBetween(1, 240),
			},
			"notes": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Your notes for this monitor.",
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"selected_checkpoints": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"checkpoints": {
							Type:     schema.TypeSet,
							Optional: true,
							Description: "A checkpoint is a geographic location from which you can have your service uptime and performance checked periodically.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"regions": {
							Type:     schema.TypeSet,
							Optional: true,
							Description: "A Region contains one or more checkpoints, just define a region if all checkpoints in a region should be used.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
								// Todo normalize regionsid to region names
								// StateFunc: func(val interface{}) string {
								// 	return profileID(val.(string))
								// },
							},
						},
						"exclude_locations": {
							Type:     schema.TypeSet,
							Optional: true,
							Description: "Exclude locations e.g. single checkpoints if a entire region was specified.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
					},
				},
			},
			"primary_checkpoints_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "The recommended value is True. Only set this to False when you’re sure you want to execute your monitor on non-primary checkpoints.",
			},
			"is_locked": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "It specifies whether the monitor is currently locked for editing. This happens if the Support team is reviewing your monitor.",
			},
			"name_for_phone_alerts": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The value for the speech-friendly monitor name, if applicable. This is the monitor name we’ll use in text-to-speech phone alerting, provided that the ‘Use alternate monitor names’ option has been enabled in the phone alert integration.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"ip_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "IpV4",
				Description: "IpV4 or IpV6. Indicates which IP version should be used to connect to the server or network address you specify. If you choose IPv6, the monitor will only be executed on checkpoint locations that support IPv6.",
				ValidateFunc: validation.StringInSlice([]string{"IpV4", "IpV6"}, false),
			},
			"native_ipv6_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "True or False. This setting only applies when you select IpV6 for the IpVersion field. Set this value to true to only execute your monitor on checkpoint servers that support native IPv6 connectivity.",
			},
			"alert_on_load_time_limit_1": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"load_time_limit_1": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      2500,
				ValidateFunc: validation.IntBetween(1, 300000),
				Description:  "Time in ms", // Todo improve description
			},
			"alert_on_load_time_limit_2": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"load_time_limit_2": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      5000,
				ValidateFunc: validation.IntBetween(1, 300000),
				Description:  "Time in ms", // Todo improve description
			},
		},
	}
}

func expandSelectedCheckpoints(sc []interface{}) *uptrends.SelectedCheckpoints {
	if len(sc) == 0 {
		return &uptrends.SelectedCheckpoints{}
	}
	v := sc[0].(map[string]interface{})

	selectedCheckpoints := uptrends.SelectedCheckpoints{}
	if v["checkpoints"].(*schema.Set).Len() != 0 {
		c := SliceInterfaceToSliceInt32(v["checkpoints"].(*schema.Set).List())
		selectedCheckpoints.Checkpoints = c
	}

	if v["regions"].(*schema.Set).Len() != 0 {
		r := SliceInterfaceToSliceInt32(v["regions"].(*schema.Set).List())
		selectedCheckpoints.Regions = r
	}

	if v["exclude_locations"].(*schema.Set).Len() != 0 {
		el := SliceInterfaceToSliceInt32(v["exclude_locations"].(*schema.Set).List())
		selectedCheckpoints.Checkpoints = el
	}

	return &selectedCheckpoints
}

func flattenSelectedCheckpoints(sc *uptrends.SelectedCheckpoints) []interface{} {
	if sc == nil {
		return []interface{}{}
	}

	selectedCheckpoints := make(map[string]interface{})
	if v := sc.Checkpoints; v != nil {
		selectedCheckpoints["checkpoints"] = *v
	}

	if v := sc.Regions; v != nil {
		selectedCheckpoints["regions"] = *v
	}

	if v := sc.ExcludeLocations; v != nil {
		selectedCheckpoints["exclude_locations"] = *v
	}

	return []interface{}{selectedCheckpoints}
}

func buildMonitorPingStruct(d *schema.ResourceData) (*uptrends.Monitor, error) {
	monitorMode, err := uptrends.NewMonitorModeFromValue(d.Get("mode").(string))
	if err != nil {
		return nil, err
	}

	monitorType, err := uptrends.NewMonitorTypeFromValue("Ping")
	if err != nil {
		return nil, err
	}

	ipVersion, err := uptrends.NewIpVersionFromValue(d.Get("ip_version").(string))
	if err != nil {
		return nil, err
	}

	customFields := []uptrends.CustomField{}

	monitor := &uptrends.Monitor{
		Name:                      String(d.Get("name").(string)),
		MonitorMode:               monitorMode,
		MonitorType:               monitorType,
		NetworkAddress:            String(d.Get("network_address").(string)),
		CustomFields:              &customFields,
		IsActive:                  Bool(d.Get("is_active").(bool)),
		GenerateAlert:             Bool(d.Get("generate_alert").(bool)),
		CheckInterval:             Int32(int32(d.Get("check_interval").(int))),
		UsePrimaryCheckpointsOnly: Bool(d.Get("primary_checkpoints_only").(bool)),
		IpVersion:                 ipVersion,
		NativeIPv6Only:            Bool(d.Get("native_ipv6_only").(bool)),
		AlertOnLoadTimeLimit1:     Bool(d.Get("alert_on_load_time_limit_1").(bool)),
		LoadTimeLimit1:            Int32(int32(d.Get("load_time_limit_1").(int))),
		AlertOnLoadTimeLimit2:     Bool(d.Get("alert_on_load_time_limit_2").(bool)),
		LoadTimeLimit2:            Int32(int32(d.Get("load_time_limit_2").(int))),
		SelectedCheckpoints:       expandSelectedCheckpoints(d.Get("selected_checkpoints").([]interface{})),
	}

	if attr, ok := d.GetOk("notes"); ok {
		monitor.Notes = String(attr.(string))
	}

	if attr, ok := d.GetOk("name_for_phone_alerts"); ok {
		monitor.NameForPhoneAlerts = String(attr.(string))
	}

	return monitor, nil
}

func monitorPingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	monitor, err := buildMonitorPingStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, _, err := client.MonitorPostMonitor(auth).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*resp.MonitorGuid)

	return monitorPingRead(ctx, d, meta)
}

func readMonitorPingStruct(monitor uptrends.Monitor, d *schema.ResourceData) error {
	d.Set("name", monitor.Name)
	d.Set("is_active", monitor.IsActive)
	d.Set("network_address", monitor.NetworkAddress)
	d.Set("generate_alert", monitor.GenerateAlert)
	d.Set("check_interval", monitor.CheckInterval)
	d.Set("primary_checkpoints_only", monitor.UsePrimaryCheckpointsOnly)
	d.Set("notes", monitor.Notes)
	d.Set("name_for_phone_alerts", monitor.NameForPhoneAlerts)
	d.Set("ip_version", monitor.IpVersion)
	d.Set("native_ipv6_only", monitor.NativeIPv6Only)
	d.Set("alert_on_load_time_limit_1", monitor.AlertOnLoadTimeLimit1)
	d.Set("load_time_limit_1", monitor.LoadTimeLimit1)
	d.Set("alert_on_load_time_limit_2", monitor.AlertOnLoadTimeLimit2)
	d.Set("load_time_limit_2", monitor.LoadTimeLimit2)

	if sc := monitor.SelectedCheckpoints; *sc != (uptrends.SelectedCheckpoints{}) {
		d.Set("selected_checkpoints", flattenSelectedCheckpoints(monitor.SelectedCheckpoints))	
	}

	return nil
}

func monitorPingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	resp, _, err := client.MonitorGetMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	err = readMonitorPingStruct(resp, d)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func monitorPingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	monitor := uptrends.Monitor{}

	if d.HasChange("name") {
		monitor.Name = String(d.Get("name").(string))
	}

	if d.HasChange("is_active") {
		monitor.IsActive = Bool(d.Get("is_active").(bool))
	}

	if d.HasChange("network_address") {
		monitor.NetworkAddress = String(d.Get("network_address").(string))
	}

	if d.HasChange("generate_alert") {
		monitor.GenerateAlert = Bool(d.Get("generate_alert").(bool))
	}

	if d.HasChange("check_interval") {
		monitor.CheckInterval = Int32(int32(d.Get("check_interval").(int)))
	}

	if d.HasChange("primary_checkpoints_only") {
		monitor.UsePrimaryCheckpointsOnly = Bool(d.Get("primary_checkpoints_only").(bool))
	}

	if d.HasChange("notes") {
		monitor.Notes = String(d.Get("notes").(string))
	}

	if d.HasChange("name_for_phone_alerts") {
		monitor.NameForPhoneAlerts = String(d.Get("name_for_phone_alerts").(string))
	}

	if d.HasChange("ip_version") {
		monitor.NameForPhoneAlerts = String(d.Get("ip_version").(string))
	}

	if d.HasChange("native_ipv6_only") {
		monitor.NativeIPv6Only = Bool(d.Get("native_ipv6_only").(bool))
	}

	if d.HasChange("alert_on_load_time_limit_1") {
		monitor.AlertOnLoadTimeLimit1 = Bool(d.Get("alert_on_load_time_limit_1").(bool))
	}

	if d.HasChange("load_time_limit_1") {
		monitor.LoadTimeLimit1 = Int32(int32(d.Get("load_time_limit_1").(int)))
	}

	if d.HasChange("alert_on_load_time_limit_2") {
		monitor.AlertOnLoadTimeLimit2 = Bool(d.Get("alert_on_load_time_limit_2").(bool))
	}

	if d.HasChange("load_time_limit_2") {
		monitor.LoadTimeLimit2 = Int32(int32(d.Get("load_time_limit_2").(int)))
	}

	if d.HasChange("selected_checkpoints") {
		monitor.SelectedCheckpoints = expandSelectedCheckpoints(d.Get("selected_checkpoints").([]interface{}))
	}

	_, err := client.MonitorPatchMonitor(auth, id).Monitor(monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}


	return monitorPingRead(ctx, d, meta)
}

func monitorPingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	_, err := client.MonitorDeleteMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
