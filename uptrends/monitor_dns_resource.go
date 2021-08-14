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
		Schema: MergeSchema(MonitorGenericSchema, MonitorLoadTimeSchema, map[string]*schema.Schema{
			"port": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      53,
				Description:  "The TCP Port for the dns Monitor, has to be between `1` and `65535`. Defaults to `53`.",
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
				Description:  "The value to test for e.g. example.com.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"dns_expected_result": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The result you expect to get back from DNS query.",
				ValidateFunc: validation.StringIsNotEmpty,
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
	m := uptrends.NewMonitor()

	if err := buildMonitorGenericStruct(m, d); err != nil {
		return nil, err
	}
	if err := buildMonitorLoadTimeStruct(m, d); err != nil {
		return nil, err
	}
	ipVersion, err := uptrends.NewIpVersionFromValue(d.Get("ip_version").(string))
	if err != nil {
		return nil, err
	}
	dnsQuery, err := uptrends.NewDnsQueryFromValue(d.Get("dns_query").(string))
	if err != nil {
		return nil, err
	}

	m.MonitorType = &monitorDnsType
	m.Port = Int32(int32(d.Get("port").(int)))
	m.IpVersion = ipVersion
	m.NativeIPv6Only = Bool(d.Get("native_ipv6_only").(bool))
	m.DnsServer = String(d.Get("dns_server").(string))
	m.DnsQuery = dnsQuery
	m.DnsTestValue = String(d.Get("dns_test_value").(string))
	m.DnsExpectedResult = String(d.Get("dns_expected_result").(string))

	return m, nil
}

func readMonitorDnsStruct(m *uptrends.Monitor, d *schema.ResourceData) diag.Diagnostics {
	if err := readMonitorGenericStruct(m, d); err != nil {
		return diag.FromErr(err)
	}
	if err := readMonitorLoadTimeStruct(m, d); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("port", m.Port); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ip_version", m.IpVersion); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("native_ipv6_only", m.NativeIPv6Only); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("dns_server", m.DnsServer); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("dns_query", m.DnsQuery); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("dns_test_value", m.DnsTestValue); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("dns_expected_result", m.DnsExpectedResult); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
