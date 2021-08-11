package uptrends

import (
	"context"

	uptrends "github.com/wasfree/uptrends-go-sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var (
	monitorHttpType = uptrends.MONITORTYPE_HTTP
)

func ResourceMonitorHttpSchema() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages an Uptrends Http Monitor.",
		CreateContext: monitorHttpCreate,
		ReadContext:   monitorHttpCreate,
		UpdateContext: monitorHttpCreate,
		DeleteContext: monitorHttpCreate,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: MergeSchema(MonitorGenericSchema, map[string]*schema.Schema{
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
			"url": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The full URL of the appropriate website, page or service that you want to monitor. The URL should include “http://” or “https://”. If relevant, please also include a port number if you are using a non-default port number, e.g. https://your-domain.com:8080/your-page. You can also use a fixed IP address as part of the URL instead of a host name, if your server listens to incoming requests without a host name.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"http_method": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "GET",
				Description:  "Specifies the HTTP methode for your monitor. Defaults to `GET`.",
				ValidateFunc: validation.StringInSlice([]string{"GET", "POST"}, false),
			},

			// "RequestHeaders": [
			// 	{
			// 	  "Name": "Content-Type",
			// 	  "Value": "Application/Json"
			// 	},
			// 	{
			// 	  "Name": "Foo",
			// 	  "Value": "Bar"
			// 	}
			//   ],
			// "request_headers": {
			// 	Type:         schema.TypeString,
			// 	Optional:     true,
			// 	Description:  "Specify HTTP Headers allowed are predefined and custom headers. Each header should appear on a separated line.",
			// 	ValidateFunc: validation.StringIsNotEmpty,
			// },
			"request_body": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "If `http_method` POST was specified it's possible post a form e.g. XML or JSON.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"expected_http_status_code": {
				Type:         schema.TypeBool,
				Optional:     true,
				Description:  "Check for specific HTTP status code, any other than the specified status code, will generate an error. Requires `expected_http_status_code_specified` to be enabled.",
				ValidateFunc: validation.IntInSlice(validHttpCodes),
			},
			"expected_http_status_code_specified": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable check for specific HTTP status code defined in `expected_http_status_code`.",
			},
			// Is this also for HTTP connections required?
			// "tls_version": {
			// 	Type: schema.TypeString,
			// 	Optional: true,
			// 	Default: "Tls12_Tls11_Tls10",
			// 	Description: "The TLS version the Monitor should support when setting up the secure HTTPS connection.",
			// },
			"user_agent": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "A string value that identifies which HTTP client is making the HTTP request. A browser typically sends a value that identifies the browser type and version.",
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
				Description: "Enable this option if you want to check that the response of the Server contains at least `min_bytes` of bytes",
			},
			"min_bytes": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Set threshold bytes that the response of the Server must at least contains. Required `alert_on_min_bytes` to be enabled.",
				ValidateFunc: validation.IntBetween(0, 20000000),
			},
			"match_pattern": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Insert a word or phrase here, if you want to verify that your web page contains this text.",
				ValidateFunc: validation.StringLenBetween(0, 512),
			},
		}),
	}
}

func monitorHttpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	_ = meta.(*Uptrends).Client.MonitorApi
	_ = meta.(*Uptrends).AuthContext

	return nil
}
