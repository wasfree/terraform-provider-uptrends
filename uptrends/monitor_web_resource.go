package uptrends

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/wasfree/uptrends-go-sdk"
)

func ResourceMonitorWebSchema() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages Uptrends HTTP, HTTPS, Webservice HTTP and Webservice HTTPS Monitor.",
		CreateContext: monitorWebCreate,
		ReadContext:   monitorWebRead,
		UpdateContext: monitorWebUpdate,
		DeleteContext: monitorWebDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: MergeSchema(MonitorGenericSchema, MonitorLoadTimeSchema, map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The full URL of the appropriate website, page or service that you want to monitor. The URL should include “http://” or “https://”. If relevant, please also include a port number if you are using a non-default port number, e.g. https://your-domain.com:8080/your-page. You can also use a fixed IP address as part of the URL instead of a host name, if your server listens to incoming requests without a host name.",
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     uptrends.MONITORTYPE_HTTP,
				Description: "Select between `Http`, `Https`, `WebserviceHttp` and `WebserviceHttps` monitor type. Defaults to `Http`",
				ValidateFunc: validation.StringInSlice([]string{
					string(uptrends.MONITORTYPE_HTTP),
					string(uptrends.MONITORTYPE_HTTPS),
					string(uptrends.MONITORTYPE_WEBSERVICE_HTTP),
					string(uptrends.MONITORTYPE_WEBSERVICE_HTTPS)},
					false),
			},
			"check_cert_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "An HTTPS check will only pass our checks if the SSL certificate does not cause any errors. Only set this option to `false` if you really want to ignore SSL certificate issues. This parameter takes only effect if `type` has been set to `Https`.",
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
			"native_ipv6_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "True or False. This setting only applies when you select IpV6 for the IpVersion field. Set this value to true to only execute your monitor on checkpoint servers that support native IPv6 connectivity. Defaults to `false`.",
			},
			"http_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     uptrends.HTTPMETHOD_GET,
				Description: "Specifies the HTTP methode for your monitor. Defaults to `GET`.",
				ValidateFunc: validation.StringInSlice([]string{
					string(uptrends.HTTPMETHOD_GET),
					string(uptrends.HTTPMETHOD_POST)},
					false),
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
				Type:        schema.TypeString,
				Optional:    true,
				Default:     uptrends.APIHTTPAUTHENTICATIONTYPE_NONE,
				Description: "Specify the username of the appropriate credentials here. Defaults to `None`.",
				ValidateFunc: validation.StringInSlice([]string{
					string(uptrends.APIHTTPAUTHENTICATIONTYPE_BASIC),
					string(uptrends.APIHTTPAUTHENTICATIONTYPE_DIGEST),
					string(uptrends.APIHTTPAUTHENTICATIONTYPE_NONE),
					string(uptrends.APIHTTPAUTHENTICATIONTYPE_NTLM)},
					false),
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
							Type:        schema.TypeString,
							Required:    true,
							Description: "Insert a word or phrase if you want to verify that your web page contains that text.",
						},
						"is_positive": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Set this value to `true` if you want to verify that your web page contains `pattern`. Set to false if you want to verify that your web page does not contain that text.",
						},
					},
				},
			},
		}),
	}
}

func monitorWebCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	monitor, err := buildMonitorWebStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, _, err := client.MonitorPostMonitor(auth).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*resp.MonitorGuid)

	return monitorWebRead(ctx, d, meta)
}

func monitorWebRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	resp, _, err := client.MonitorGetMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return readMonitorWebStruct(&resp, d)
}

func monitorWebUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	monitor, err := buildMonitorWebStruct(d)
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

	return monitorWebRead(ctx, d, meta)
}

func monitorWebDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	_, err := client.MonitorDeleteMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildMonitorWebStruct(d *schema.ResourceData) (*uptrends.Monitor, error) {
	m := uptrends.NewMonitor()

	if err := buildMonitorGenericStruct(m, d); err != nil {
		return nil, err
	}
	if err := buildMonitorLoadTimeStruct(m, d); err != nil {
		return nil, err
	}

	m.MonitorType = (*uptrends.MonitorType)(String(d.Get("type").(string)))
	m.IpVersion = (*uptrends.IpVersion)(String(d.Get("ip_version").(string)))
	m.NativeIPv6Only = Bool(d.Get("native_ipv6_only").(bool))
	m.Url = String(d.Get("url").(string))
	m.HttpMethod = (*uptrends.HttpMethod)(String(d.Get("http_method").(string)))
	m.RequestBody = String(d.Get("request_body").(string))
	m.AuthenticationType = (*uptrends.ApiHttpAuthenticationType)(String(d.Get("auth_type").(string)))
	m.UserAgent = String(UserAgent(d.Get("user_agent").(string)))
	m.Username = String(d.Get("username").(string))
	m.AlertOnMinimumBytes = Bool(d.Get("alert_on_min_bytes").(bool))
	m.MinimumBytes = Int32(int32(d.Get("min_bytes").(int)))

	// check_cert_errors should only be set if monitorType is Https
	if *m.MonitorType == uptrends.MONITORTYPE_HTTPS {
		m.CheckCertificateErrors = Bool(d.Get("check_cert_errors").(bool))
	}

	// Optional without defaults
	if attr, ok := d.GetOk("expected_http_status_code"); ok {
		m.ExpectedHttpStatusCode = Int32(int32(attr.(int)))
	}
	if attr, ok := d.GetOk("password"); ok {
		m.Password = String(attr.(string))
	}
	if attr, ok := d.Get("request_headers").([]interface{}); ok {
		m.RequestHeaders = expandRequestHeader(attr)
	}
	if attr, ok := d.Get("match_pattern").([]interface{}); ok {
		m.MatchPatterns = expandPatternMatch(attr)
	}

	return m, nil
}

func readMonitorWebStruct(m *uptrends.Monitor, d *schema.ResourceData) diag.Diagnostics {
	if err := readMonitorGenericStruct(m, d); err != nil {
		return diag.FromErr(err)
	}
	if err := readMonitorLoadTimeStruct(m, d); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ip_version", m.IpVersion); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("http_method", m.HttpMethod); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("auth_type", m.AuthenticationType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("username", m.Username); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("native_ipv6_only", m.NativeIPv6Only); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("url", m.Url); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("expected_http_status_code", m.ExpectedHttpStatusCode); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("user_agent", m.UserAgent); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alert_on_min_bytes", m.AlertOnMinimumBytes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("min_bytes", m.MinimumBytes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("request_headers", flattenRequestHeader(m.RequestHeaders)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("match_pattern", flattenPatternMatch(m.MatchPatterns)); err != nil {
		return diag.FromErr(err)
	}
	if *m.MonitorType == uptrends.MONITORTYPE_HTTPS {
		if err := d.Set("check_cert_errors", m.CheckCertificateErrors); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}
