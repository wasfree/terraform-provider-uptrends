package uptrends

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/wasfree/uptrends-go-sdk"
)

func ResourceMonitorNetworkSchema() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages an Uptrends Ping and Connect Monitor.",
		CreateContext: monitorNetworkCreate,
		ReadContext:   monitorNetworkRead,
		UpdateContext: monitorNetworkUpdate,
		DeleteContext: monitorNetworkDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: MergeSchema(MonitorGenericSchema, MonitorLoadTimeSchema, map[string]*schema.Schema{
			"network_address": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The network address that should be used to connect to the server or service you want to monitor.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "Ping",
				Description:  "Select between `Ping` and `Connect` monitor type. Defaults to `Ping`",
				ValidateFunc: validation.StringInSlice([]string{"Ping", "Connect"}, true),
			},
			"port": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "The TCP Port for the connect Monitor. Only used by connect Monitor, has to be between `1` and `65535`.",
				ValidateFunc: validation.IntBetween(1, 65535),
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
		}),
	}
}

func monitorNetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	monitor, err := buildMonitorNetworkStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, _, err := client.MonitorPostMonitor(auth).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*resp.MonitorGuid)

	return monitorNetworkRead(ctx, d, meta)
}

func monitorNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	resp, _, err := client.MonitorGetMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return readMonitorNetworkStruct(&resp, d)
}

func monitorNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	monitor, err := buildMonitorNetworkStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.MonitorPatchMonitor(auth, id).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return monitorNetworkRead(ctx, d, meta)
}

func monitorNetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	_, err := client.MonitorDeleteMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildMonitorNetworkStruct(d *schema.ResourceData) (*uptrends.Monitor, error) {
	m := uptrends.NewMonitor()

	if err := buildMonitorGenericStruct(m, d); err != nil {
		return nil, err
	}
	if err := buildMonitorLoadTimeStruct(m, d); err != nil {
		return nil, err
	}

	m.MonitorType = (*uptrends.MonitorType)(String(d.Get("type").(string)))
	m.NetworkAddress = String(d.Get("network_address").(string))
	m.IpVersion = (*uptrends.IpVersion)(String(d.Get("ip_version").(string)))
	m.NativeIPv6Only = Bool(d.Get("native_ipv6_only").(bool))

	if *m.MonitorType == uptrends.MONITORTYPE_CONNECT {
		m.Port = Int32(int32(d.Get("port").(int)))
	}

	return m, nil
}

func readMonitorNetworkStruct(m *uptrends.Monitor, d *schema.ResourceData) diag.Diagnostics {
	if err := readMonitorGenericStruct(m, d); err != nil {
		return diag.FromErr(err)
	}
	if err := readMonitorLoadTimeStruct(m, d); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("network_address", m.NetworkAddress); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ip_version", m.IpVersion); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("native_ipv6_only", m.NativeIPv6Only); err != nil {
		return diag.FromErr(err)
	}
	if *m.MonitorType == uptrends.MONITORTYPE_CONNECT {
		if err := d.Set("port", m.Port); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}
