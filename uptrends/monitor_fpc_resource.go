package uptrends

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/wasfree/uptrends-go-sdk"
)

var (
	monitorFullPageCheckType = uptrends.MONITORTYPE_FULL_PAGE_CHECK
)

func ResourceMonitorFullPageCheckSchema() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages Uptrends Full Page Check Monitor.",
		CreateContext: monitorFullPageCheckCreate,
		ReadContext:   monitorFullPageCheckRead,
		UpdateContext: monitorFullPageCheckUpdate,
		DeleteContext: monitorFullPageCheckDelete,
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
			"user_agent": {
				Type:        schema.TypeString,
				Optional:    true,
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
			"match_pattern": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pattern": {
							Type:         schema.TypeString,
							Required:     true,
							Description:  "Insert a word or phrase if you want to verify that your web page contains that text.",
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"is_positive": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Set this value to `true` if you want to verify that your web page contains `pattern`. Set to false if you want to verify that your web page does not contain that text.",
						},
					},
				},
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
			"alert_on_max_bytes": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable this option if you want to check that the response of the Server not exceed `max_bytes` of bytes. Defaults to `false`",
			},
			"max_bytes": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Set threshold bytes that the response of the Server must at least contains. Required `alert_on_max_bytes` to be enabled.",
				ValidateFunc: validation.IntBetween(0, 20000000),
			},
			"alert_on_max_element_bytes": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable this option if you want to check that any element on your page not exceeds `max_element_bytes` of bytes. Defaults to `false`",
			},
			"max_element_bytes": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Set threshold bytes that the response of the Server must at least contains. Required `alert_on_max_element_bytes` to be enabled.",
				ValidateFunc: validation.IntBetween(0, 20000000),
			},
			"ignore_external_elements": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable this option if you don't want to include external elements in any error condition that are related to download size.",
			},
			"alert_on_percentage_fail": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable this option if you want the time-, size-, or content-checks to generate alerts.",
			},
			"failed_object_percentage": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Specify a percentage treshold on which your want to get alerted 0% will report all element load failures",
				ValidateFunc: validation.IntBetween(0, 100),
			},
			"browser_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     uptrends.BROWSERTYPE_CHROME,
				Description: "The Type of Browser used to Load the Website. Allowed values are `Chrome`, `Firefox`, `IE`, `PhantomJS`, `PhantomJS20` and `Safari`. Defaults to `Chrome`.",
				ValidateFunc: validation.StringInSlice([]string{
					string(uptrends.BROWSERTYPE_CHROME),
					string(uptrends.BROWSERTYPE_FIREFOX),
					string(uptrends.BROWSERTYPE_IE),
					string(uptrends.BROWSERTYPE_PHANTOM_JS),
					string(uptrends.BROWSERTYPE_PHANTOM_JS20),
					string(uptrends.BROWSERTYPE_SAFARI)},
					false),
			},
			"browser_window_dimensions": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_mobile": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable for Mobile simulation.",
						},
						"width": {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "Set the width for resolution size.",
							ValidateFunc: validation.IntBetween(1, 1600),
						},
						"height": {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "Set the height for resolution size.",
							ValidateFunc: validation.IntBetween(1, 900),
						},
						"pixel_ratio": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The device pixel ratio is the ratio between physical pixels and logical pixels.",
						},
						"mobile_device": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Define the mobile device if `is_mobile` set to true.",
						},
					},
				},
			},
			"block_google_analytics": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enable this option if outgoing monitoring requests should block google analytics.",
			},
			"block_uptrends_rum": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enable this option if outgoing monitoring requests should block uptrends rum.",
			},
			"block_urls": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Set Urls which should be blocked by outgoing monitoring requests.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		}),
	}
}

func monitorFullPageCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	monitor, err := buildMonitorFullPageCheckStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, _, err := client.MonitorPostMonitor(auth).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*resp.MonitorGuid)

	return monitorFullPageCheckRead(ctx, d, meta)
}

func monitorFullPageCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	resp, _, err := client.MonitorGetMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return readMonitorFullPageCheckStruct(resp, d)
}

func monitorFullPageCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	monitor, err := buildMonitorFullPageCheckStruct(d)
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

	return monitorFullPageCheckRead(ctx, d, meta)
}

func monitorFullPageCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	_, err := client.MonitorDeleteMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildMonitorFullPageCheckStruct(d *schema.ResourceData) (*uptrends.Monitor, error) {
	m := uptrends.NewMonitor()

	if err := buildMonitorGenericStruct(m, d); err != nil {
		return nil, err
	}
	if err := buildMonitorLoadTimeStruct(m, d); err != nil {
		return nil, err
	}

	m.MonitorType = &monitorFullPageCheckType
	m.Url = String(d.Get("url").(string))
	m.UserAgent = String(UserAgent(d.Get("user_agent").(string)))
	m.AuthenticationType = (*uptrends.ApiHttpAuthenticationType)(String(d.Get("auth_type").(string)))
	m.Username = String(d.Get("username").(string))
	m.AlertOnMinimumBytes = Bool(d.Get("alert_on_min_bytes").(bool))
	m.MinimumBytes = Int32(int32(d.Get("min_bytes").(int)))
	m.AlertOnMaximumBytes = Bool(d.Get("alert_on_max_bytes").(bool))
	m.MaximumBytes = Int32(int32(d.Get("max_bytes").(int)))
	m.AlertOnMaximumSize = Bool(d.Get("alert_on_max_element_bytes").(bool))
	m.ElementMaximumSize = Int32(int32(d.Get("max_element_bytes").(int)))
	m.IgnoreExternalElements = Bool(d.Get("ignore_external_elements").(bool))
	m.AlertOnPercentageFail = Bool(d.Get("alert_on_percentage_fail").(bool))
	m.FailedObjectPercentage = Int32(20) //Int32(int32(d.Get("failed_object_percentage").(int)))
	m.BrowserType = (*uptrends.BrowserType)(String(d.Get("browser_type").(string)))
	m.BlockGoogleAnalytics = Bool(d.Get("block_google_analytics").(bool))
	m.BlockUptrendsRum = Bool(d.Get("block_uptrends_rum").(bool))
	m.BlockUrls = SliceInterfaceToSliceString(d.Get("block_urls").(*schema.Set).List())

	// Optional without defaults
	if attr, ok := d.GetOk("password"); ok {
		m.Password = String(attr.(string))
	}
	if attr, ok := d.Get("request_headers").([]interface{}); ok {
		m.RequestHeaders = expandRequestHeader(attr)
	}
	if attr, ok := d.Get("match_pattern").([]interface{}); ok {
		m.MatchPatterns = expandPatternMatch(attr)
	}
	if attr, ok := d.Get("browser_window_dimensions").([]interface{}); ok {
		m.BrowserWindowDimensions = expandBrowserWindowDimensions(attr)
	}

	return m, nil
}

func readMonitorFullPageCheckStruct(m *uptrends.Monitor, d *schema.ResourceData) diag.Diagnostics {
	if err := readMonitorGenericStruct(m, d); err != nil {
		return diag.FromErr(err)
	}
	if err := readMonitorLoadTimeStruct(m, d); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("url", m.Url); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("user_agent", m.UserAgent); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("auth_type", m.AuthenticationType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("username", m.Username); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alert_on_min_bytes", m.AlertOnMinimumBytes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("min_bytes", m.MinimumBytes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alert_on_max_bytes", m.AlertOnMaximumBytes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("max_bytes", m.MaximumBytes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alert_on_max_element_bytes", m.AlertOnMaximumSize); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("max_element_bytes", m.ElementMaximumSize); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ignore_external_elements", m.IgnoreExternalElements); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alert_on_percentage_fail", m.AlertOnPercentageFail); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("failed_object_percentage", m.FailedObjectPercentage); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("browser_type", m.BrowserType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("block_google_analytics", m.BlockGoogleAnalytics); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("block_uptrends_rum", m.BlockUptrendsRum); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("block_urls", SliceStringToSliceInterface(&m.BlockUrls)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("request_headers", flattenRequestHeader(&m.RequestHeaders)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("match_pattern", flattenPatternMatch(&m.MatchPatterns)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("browser_window_dimensions", flattenBrowserWindowDimensions(m.BrowserWindowDimensions)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
