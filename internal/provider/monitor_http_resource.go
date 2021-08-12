package uptrends

import (
	"context"

	uptrends "github.com/wasfree/uptrends-go-sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceMonitorHttpSchema() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages Uptrends HTTP and HTTPS Monitor.",
		CreateContext: monitorHttpCreate,
		ReadContext:   monitorHttpRead,
		UpdateContext: monitorHttpUpdate,
		DeleteContext: monitorHttpDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: MergeSchema(MonitorGenericSchema, map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The full URL of the appropriate website, page or service that you want to monitor. The URL should include “http://” or “https://”. If relevant, please also include a port number if you are using a non-default port number, e.g. https://your-domain.com:8080/your-page. You can also use a fixed IP address as part of the URL instead of a host name, if your server listens to incoming requests without a host name.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"type": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default: "Http",
				Description: "Select between `Http` and `Https` monitor type. Defaults to `Http`",
				ValidateFunc: validation.StringInSlice([]string{"Http", "Https"}, true),
			},
			"check_cert_errors": {
				Type: schema.TypeBool,
				Optional: true,
				Default:  true,
				Description: "An HTTPS check will only pass our checks if the SSL certificate does not cause any errors. Only set this option to `false` if you really want to ignore SSL certificate issues. This parameter takes only effect if `type` has been set to `Https`.",
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
			"http_method": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Get",
				Description:  "Specifies the HTTP methode for your monitor. Defaults to `GET`.",
				ValidateFunc: validation.StringInSlice([]string{"Get", "Post"}, false),
			},
			"request_headers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specify an HTTP header name used by your requests.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specify an HTTP header value used by your requests.",
						},
					},
				},
			},
			"request_body": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "When posting a form, fill in the form variables, every form variable has to be on separated line e.g. `foo=bar\r\nbar=foo\r\n`. Requires `http_method` to be set to `Post`.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"expected_http_status_code": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Check for specific HTTP status code, any other than the specified status code, will generate an error.",
				ValidateFunc: validation.IntInSlice(validHttpCodes),
			},
			"user_agent": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "chrome83",
				Description: "A string value that identifies which HTTP client is making the HTTP request. A browser typically sends a value that identifies the browser type and version.",
				StateFunc: func(v interface{}) string {
					return UserAgent(v.(string))
				},
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"auth_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "None",
				Description:  "Specify the username of the appropriate credentials here. Defaults to `None`.",
				ValidateFunc: validation.StringInSlice([]string{"None", "Basic", "NTLM", "Digest"}, false),
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
			"alert_on_min_bytes": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable this option if you want to check that the response of the Server contains at least `min_bytes` of bytes. Defaults to `false`",
			},
			"min_bytes": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Set threshold bytes that the response of the Server must at least contains. Required `alert_on_min_bytes` to be enabled.",
				ValidateFunc: validation.IntBetween(0, 20000000),
			},
			"match_pattern": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pattern": {
							Type:     schema.TypeString,
							Required: true,
							Description: "Insert a word or phrase if you want to verify that your web page contains that text.",
						},
						"is_positive": {
							Type:     schema.TypeBool,
							Required: true,
							Description: "Set this value to `true` if you want to verify that your web page contains `pattern`. Set to false if you want to verify that your web page does not contain that text.",
						},
					},
				},
			},
		}),
	}
}

func monitorHttpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	monitor, err := buildMonitorHttpStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, _, err := client.MonitorPostMonitor(auth).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*resp.MonitorGuid)

	return monitorHttpRead(ctx, d, meta)
}

func monitorHttpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	resp, _, err := client.MonitorGetMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return readMonitorHttpStruct(&resp, d)
}

func monitorHttpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	monitor, err := buildMonitorHttpStruct(d)
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

	return monitorHttpRead(ctx, d, meta)
}

func monitorHttpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	_, err := client.MonitorDeleteMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildMonitorHttpStruct(d *schema.ResourceData) (*uptrends.Monitor, error) {
	monitor, err := buildMonitorGenericStruct(d)
	if err != nil {
		return nil, err
	}
	ipVersion, err := uptrends.NewIpVersionFromValue(d.Get("ip_version").(string))
	if err != nil {
		return nil, err
	}
	httpMethod, err := uptrends.NewHttpMethodFromValue(d.Get("http_method").(string))
	if err != nil {
		return nil, err
	}
	authType, err := uptrends.NewApiHttpAuthenticationTypeFromValue(d.Get("auth_type").(string))
	if err != nil {
		return nil, err
	}

	monitorType, err := uptrends.NewMonitorTypeFromValue(d.Get("type").(string))
	if err != nil {
		return nil, err
	}

	monitor.MonitorType = monitorType
	monitor.IpVersion = ipVersion
	monitor.NativeIPv6Only = Bool(d.Get("native_ipv6_only").(bool))
	monitor.Url = String(d.Get("url").(string))
	monitor.HttpMethod = httpMethod
	monitor.RequestBody = String(d.Get("request_body").(string))
	monitor.AuthenticationType = authType
	monitor.UserAgent = String(UserAgent(d.Get("user_agent").(string)))
	monitor.Username = String(d.Get("username").(string))
	monitor.AlertOnLoadTimeLimit1 = Bool(d.Get("alert_on_load_time_limit_1").(bool))
	monitor.LoadTimeLimit1 = Int32(int32(d.Get("load_time_limit_1").(int)))
	monitor.AlertOnLoadTimeLimit2 = Bool(d.Get("alert_on_load_time_limit_2").(bool))
	monitor.LoadTimeLimit2 = Int32(int32(d.Get("load_time_limit_2").(int)))
	monitor.AlertOnMinimumBytes = Bool(d.Get("alert_on_min_bytes").(bool))
	monitor.MinimumBytes = Int32(int32(d.Get("min_bytes").(int)))

	if *monitorType == uptrends.MONITORTYPE_HTTPS {
		monitor.CheckCertificateErrors = Bool(d.Get("check_cert_errors").(bool))
	}

	// Optional without defaults
	if attr, ok := d.GetOk("expected_http_status_code"); ok {
		monitor.ExpectedHttpStatusCode = Int32(int32(attr.(int)))
	}

	if attr, ok := d.GetOk("password"); ok {
		monitor.Password = String(attr.(string))
	}

	if attr, ok := d.Get("request_headers").([]interface{}); ok {
		monitor.RequestHeaders = SliceInterfaceToSliceRequestHeader(attr)
	}

	if attr, ok := d.Get("match_pattern").([]interface{}); ok {
		monitor.MatchPatterns = SliceInterfaceToSlicePatternMatch(attr)
	}

	return monitor, nil
}

func readMonitorHttpStruct(monitor *uptrends.Monitor, d *schema.ResourceData) diag.Diagnostics {
	err := readMonitorGenericStruct(monitor, d)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ip_version", monitor.IpVersion); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("http_method", monitor.HttpMethod); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("auth_type", monitor.AuthenticationType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("native_ipv6_only", monitor.NativeIPv6Only); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("url", monitor.Url); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("expected_http_status_code", monitor.ExpectedHttpStatusCode); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("user_agent", monitor.UserAgent); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("auth_type", monitor.AuthenticationType); err != nil {
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
	if err := d.Set("alert_on_min_bytes", monitor.AlertOnMinimumBytes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("min_bytes", monitor.MinimumBytes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("request_headers", SliceRequestHeaderToSliceInterface(*monitor.RequestHeaders)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("match_pattern", SlicePatternMatchToSliceInterface(*monitor.MatchPatterns)); err != nil {
		return diag.FromErr(err)
	}
	if *monitor.MonitorType == uptrends.MONITORTYPE_HTTPS {
		if err := d.Set("check_cert_errors", monitor.CheckCertificateErrors); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}
