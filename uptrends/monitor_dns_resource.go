package uptrends

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/wasfree/uptrends-go-sdk"
)

var (
	monitorDnsType = uptrends.MONITORTYPE_DNS
)

func ResourceMonitorDnsSchema() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages an Uptrends DNS Monitor.",
		CreateContext: monitorDnsCreate,
		ReadContext:   monitorDnsRead,
		UpdateContext: monitorDnsUpdate,
		DeleteContext: monitorDnsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: MergeSchema(MonitorGenericSchema, map[string]*schema.Schema{
			"port": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "The TCP Port for the dns Monitor, has to be between `1` and `65535`.",
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
			"dns_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The DNS Server that contains the entries you want to test.",
			},
			"dns_query": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The Type of DNS query you want to perform. Alloed values are `ARecord`, `CnameRecord`, `MxRecord`, `NsRecord`, `TxtRecord`, `SoaRecord`, `RootServer`, `AaaaRecord`, `SrvRecord`.",
				ValidateFunc: validation.StringInSlice([]string{"ARecord", "CnameRecord", "MxRecord", "NsRecord", "TxtRecord", "SoaRecord", "RootServer", "AaaaRecord", "SrvRecord"}, false),
			},
			"dns_test_value": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The value to test for.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"dns_expected_result": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The result you expect to get back from DNS query.",
				ValidateFunc: validation.StringIsNotEmpty,
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

func monitorDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	monitor, err := buildMonitorDnsStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, _, err := client.MonitorPostMonitor(auth).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*resp.MonitorGuid)

	return monitorDnsRead(ctx, d, meta)
}

func monitorDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	resp, _, err := client.MonitorGetMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return readMonitorDnsStruct(&resp, d)
}

func monitorDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	monitor, err := buildMonitorDnsStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.MonitorPatchMonitor(auth, id).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return monitorDnsRead(ctx, d, meta)
}

func monitorDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	_, err := client.MonitorDeleteMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildMonitorDnsStruct(d *schema.ResourceData) (*uptrends.Monitor, error) {
	monitor, err := buildMonitorGenericStruct(d)
	if err != nil {
		return nil, err
	}
	ipVersion, err := uptrends.NewIpVersionFromValue(d.Get("ip_version").(string))
	if err != nil {
		return nil, err
	}
	authType, err := uptrends.NewApiHttpAuthenticationTypeFromValue(d.Get("auth_type").(string))
	if err != nil {
		return nil, err
	}
	dnsQuery, err := uptrends.NewDnsQueryFromValue(d.Get("dns_query").(string))
	if err != nil {
		return nil, err
	}

	monitor.MonitorType = &monitorDnsType
	monitor.Port = Int32(int32(d.Get("port").(int)))
	monitor.IpVersion = ipVersion
	monitor.NativeIPv6Only = Bool(d.Get("native_ipv6_only").(bool))
	monitor.AuthenticationType = authType
	monitor.AlertOnLoadTimeLimit1 = Bool(d.Get("alert_on_load_time_limit_1").(bool))
	monitor.LoadTimeLimit1 = Int32(int32(d.Get("load_time_limit_1").(int)))
	monitor.AlertOnLoadTimeLimit2 = Bool(d.Get("alert_on_load_time_limit_2").(bool))
	monitor.LoadTimeLimit2 = Int32(int32(d.Get("load_time_limit_2").(int)))
	monitor.DnsServer = String(d.Get("dns_server").(string))
	monitor.DnsQuery = dnsQuery
	monitor.DnsTestValue = String(d.Get("dns_test_value").(string))
	monitor.DnsExpectedResult = String(d.Get("dns_expected_result").(string))

	return monitor, nil
}

func readMonitorDnsStruct(monitor *uptrends.Monitor, d *schema.ResourceData) diag.Diagnostics {
	if err := readMonitorGenericStruct(monitor, d); err != nil {
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
	if err := d.Set("dns_server", monitor.DnsServer); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("dns_query", monitor.DnsQuery); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("dns_test_value", monitor.DnsTestValue); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("dns_expected_result", monitor.DnsExpectedResult); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
