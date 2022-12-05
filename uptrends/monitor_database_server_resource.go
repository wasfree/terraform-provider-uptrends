package uptrends

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/wasfree/uptrends-go-sdk"
)

func ResourceMonitorDatabaseServerSchema() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages an Uptrends Database Server Monitor.",
		CreateContext: monitorDatabaseServerCreate,
		ReadContext:   monitorDatabaseServerRead,
		UpdateContext: monitorDatabaseServerUpdate,
		DeleteContext: monitorDatabaseServerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: MergeSchema(MonitorGenericSchema, MonitorLoadTimeSchema, map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "MySQL",
				Description: "Select between `MySQL` and `MSSQL` monitor type. Defaults to `MySQL`",
				ValidateFunc: validation.StringInSlice([]string{
					string(uptrends.MONITORTYPE_MSSQL),
					string(uptrends.MONITORTYPE_MY_SQL)},
					false),
			},
			"ip_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     uptrends.IPVERSION_IP_V4,
				Description: "IpV4 or IpV6. Indicates which IP version should be used to connect to the server or network address you specify. If you choose IPv6, the monitor will only be executed on checkpoint locations that support IPv6. Defaults to `IpV4`.",
				ValidateFunc: validation.StringInSlice([]string{
					string(uptrends.IPVERSION_IP_V4),
					string(uptrends.IPVERSION_IP_V6)},
					false),
			},
			"network_address": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The network address that should be used to connect to the server or service you want to monitor.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"port": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "The TCP Port for the dns Monitor, has to be between `1` and `65535`. Defaults to `53`.",
				ValidateFunc: validation.IntBetween(1, 65535),
			},
			"username": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Specify the username of the appropriate credentials here.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"password": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				Description:  "See the Username field. Specify the corresponding password value here.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"db_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "",
				Description:  "Optionally specify the name of the database you want to connect to.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
		}),
	}
}

func monitorDatabaseServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	monitor, err := buildMonitorDatabaseServerStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, _, err := client.MonitorPostMonitor(auth).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*resp.MonitorGuid)

	return monitorDatabaseServerRead(ctx, d, meta)
}

func monitorDatabaseServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	resp, _, err := client.MonitorGetMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return readMonitorDatabaseServerStruct(resp, d)
}

func monitorDatabaseServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	monitor, err := buildMonitorDatabaseServerStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("password") {
		monitor.Password = String(d.Get("password").(string))
	}

	_, err = client.MonitorPatchMonitor(auth, id).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return monitorDatabaseServerRead(ctx, d, meta)
}

func monitorDatabaseServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	_, err := client.MonitorDeleteMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildMonitorDatabaseServerStruct(d *schema.ResourceData) (*uptrends.Monitor, error) {
	m := uptrends.NewMonitor()

	if err := buildMonitorGenericStruct(m, d); err != nil {
		return nil, err
	}
	if err := buildMonitorLoadTimeStruct(m, d); err != nil {
		return nil, err
	}

	m.MonitorType = (*uptrends.MonitorType)(String(d.Get("type").(string)))
	m.Port = Int32(int32(d.Get("port").(int)))
	m.IpVersion = (*uptrends.IpVersion)(String(d.Get("ip_version").(string)))
	m.NetworkAddress = String(d.Get("network_address").(string))
	m.DatabaseName = String(d.Get("db_name").(string))

	if attr, ok := d.GetOk("username"); ok {
		m.Username = String(attr.(string))
	}
	if attr, ok := d.GetOk("password"); ok {
		m.Password = String(attr.(string))
	}

	return m, nil
}

func readMonitorDatabaseServerStruct(m *uptrends.Monitor, d *schema.ResourceData) diag.Diagnostics {
	if err := readMonitorGenericStruct(m, d); err != nil {
		return diag.FromErr(err)
	}
	if err := readMonitorLoadTimeStruct(m, d); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ip_version", m.IpVersion); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("network_address", m.NetworkAddress); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("port", m.Port); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("username", m.Username); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("db_name", m.DatabaseName); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
