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
		UpdateContext: monitorPingCreate,
		DeleteContext: monitorPingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Display name for the Monitor.",
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
				ValidateFunc: validation.IntBetween(1, 240),
			},
			"notes": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The monitor mode, either Development, Staging or Production. See this article for more information.",
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
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"regions": {
							Type:     schema.TypeSet,
							Optional: true,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"exclude_locations": {
							Type:     schema.TypeSet,
							Optional: true,
							MinItems: 1,
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
			"locked": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "It specifies whether the monitor is currently locked for editing. This happens if the Support team is reviewing your monitor.",
			},
			"name_for_phone_alerts": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "It specifies whether the monitor is currently locked for editing. This happens if the Support team is reviewing your monitor.",
			},
			"ip_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "IpV4",
				Description: "IpV4 or IpV6. Indicates which IP version should be used to connect to the server or network address you specify. If you choose IPv6, the monitor will only be executed on checkpoint locations that support IPv6.",
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
				Default:  true,
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
				Default:  true,
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

	regions := []int32{1004}

	checkpoint := uptrends.SelectedCheckpoints{
		Regions: &regions,
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
		SelectedCheckpoints:       &checkpoint,
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

func readMonitorPingStruct(monitor *uptrends.Monitor, d *schema.ResourceData) error {
	d.Set("name", monitor.Name)
	d.Set("mode", monitor.MonitorMode)
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

	err = readMonitorPingStruct(&resp, d)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
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
