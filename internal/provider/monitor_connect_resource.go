package uptrends

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/wasfree/uptrends-go-sdk"
)

var (
	monitorConnectType = uptrends.MONITORTYPE_CONNECT
)

func ResourceMonitorConnectSchema() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages an Uptrends Connect Monitor.",
		CreateContext: monitorConnectCreate,
		ReadContext:   monitorConnectRead,
		UpdateContext: monitorConnectUpdate,
		DeleteContext: monitorConnectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: MergeSchema(MonitorGenericSchema, map[string]*schema.Schema{
			"port": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "The TCP Port for the Monitor.",
				ValidateFunc: validation.IntBetween(1, 65535),
			},
			"network_address": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The network address that should be used to connect to the server or service you want to monitor.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"ip_version": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "IpV4",
				Description:  "IpV4 or IpV6. Indicates which IP version should be used to connect to the server or network address you specify. If you choose IPv6, the monitor will only be executed on checkpoint locations that support IPv6. Defaults to `IpV4`.",
				ValidateFunc: validation.StringInSlice([]string{"IpV4", "IpV6"}, false),
			},
			"native_ipv6_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "True or False. This setting only applies when you select IpV6 for the IpVersion field. Set this value to true to only execute your monitor on checkpoint servers that support native IPv6 connectivity. Defaults to `false`.",
			},
			"alert_on_load_time_limit_1": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Set this value to true, if you want to receive alerts if your server response is slower than load_time_limit_1 threshold. Shows a yellow status in performance monitor. Defaults to `false`.",
			},
			"load_time_limit_1": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      2500,
				ValidateFunc: validation.IntBetween(1, 300000),
				Description:  "Set threshold time in ms for requires `alert_on_load_time_limit_1` to be enabled. Defaults to `2500`.",
			},
			"alert_on_load_time_limit_2": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Set this value to true, if you want to receive alerts if your server response is slower than load_time_limit_2. Shows a red status in performance monitor. Defaults to `false`.",
			},
			"load_time_limit_2": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      5000,
				ValidateFunc: validation.IntBetween(1, 300000),
				Description:  "Set threshold time in ms for requires `alert_on_load_time_limit_2` to be enabled. Defaults to `5000`.",
			},
		}),
	}
}

func monitorConnectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	monitor, err := buildMonitorConnectStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, _, err := client.MonitorPostMonitor(auth).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*resp.MonitorGuid)

	return monitorConnectRead(ctx, d, meta)
}

func monitorConnectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	resp, _, err := client.MonitorGetMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return readMonitorConnectStruct(&resp, d)
}

func monitorConnectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	monitor, err := buildMonitorConnectStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.MonitorPatchMonitor(auth, id).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return monitorConnectRead(ctx, d, meta)
}

func monitorConnectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	_, err := client.MonitorDeleteMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildMonitorConnectStruct(d *schema.ResourceData) (*uptrends.Monitor, error) {
	monitor, err := buildMonitorGenericStruct(d)
	if err != nil {
		return nil, err
	}
	ipVersion, err := uptrends.NewIpVersionFromValue(d.Get("ip_version").(string))
	if err != nil {
		return nil, err
	}
	monitor.MonitorType = &monitorConnectType
	monitor.NetworkAddress = String(d.Get("network_address").(string))
	monitor.Port = Int32(int32(d.Get("port").(int)))
	monitor.IpVersion = ipVersion
	monitor.NativeIPv6Only = Bool(d.Get("native_ipv6_only").(bool))
	monitor.AlertOnLoadTimeLimit1 = Bool(d.Get("alert_on_load_time_limit_1").(bool))
	monitor.LoadTimeLimit1 = Int32(int32(d.Get("load_time_limit_1").(int)))
	monitor.AlertOnLoadTimeLimit2 = Bool(d.Get("alert_on_load_time_limit_2").(bool))
	monitor.LoadTimeLimit2 = Int32(int32(d.Get("load_time_limit_2").(int)))

	return monitor, nil
}

func readMonitorConnectStruct(monitor *uptrends.Monitor, d *schema.ResourceData) diag.Diagnostics {
	err := readMonitorGenericStruct(monitor, d)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("network_address", monitor.NetworkAddress); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("port", monitor.Port); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ip_version", monitor.IpVersion); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("native_ipv6_only", monitor.NativeIPv6Only); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alert_on_load_time_limit_1", monitor.AlertOnLoadTimeLimit1); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("load_time_limit_1", monitor.LoadTimeLimit1); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alert_on_load_time_limit_2", monitor.AlertOnLoadTimeLimit2); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("load_time_limit_2", monitor.LoadTimeLimit2); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
