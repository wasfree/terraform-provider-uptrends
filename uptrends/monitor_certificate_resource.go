package uptrends

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/wasfree/uptrends-go-sdk"
)

var (
	monitorCertificateType = uptrends.MONITORTYPE_CERTIFICATE
)

func ResourceMonitorCertificateSchema() *schema.Resource {
	return &schema.Resource{
		Description: "Manages an Uptrends SSL Certificate Monitor.",
		CreateContext: monitorCertificateCreate,
		ReadContext:   monitorCertificateRead,
		UpdateContext: monitorCertificateUpdate,
		DeleteContext: monitorCertificateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: MergeSchema(MonitorGenericSchema, map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The full URL of the appropriate website, page or service that you want to monitor. The URL should include “https://”, please also include a port number if you are using a non-default port number, e.g. https://your-domain.com:8080/your-page. You can also use a fixed IP address as part of the URL instead of a host name, if your server listens to incoming requests without a host name.",
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
			"cert_name": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				Description: "Specify the common name of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_org": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				Description: "Specify the organization of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_org_unit": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				Description: "Specify the organizational unit of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_serial_number": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				Description: "Specify the serial number of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_fingerprint": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				Description: "Specify the fingerprint of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_issuer_name": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				Description: "Specify the issuer common name of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_issuer_company_name": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				Description: "Specify the issuer company name of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_issuer_org_unit": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				Description: "Specify the issuer org unit of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_expiration_warning_days": {
				Type: schema.TypeInt,
				Optional: true,
				Default: 7,
				Description: "Specify the expiration warning days of the certificate. Defaults to `7` days.",
			},
			"check_cert_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "An HTTPS check will only pass our checks if the SSL certificate does not cause any errors. Only set this option to `false` if you really want to ignore SSL certificate issues. This parameter takes only effect if `type` has been set to `Https`.",
			},
		}),
	}
}

func monitorCertificateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	monitor, err := buildMonitorCertificateStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, _, err := client.MonitorPostMonitor(auth).Monitor(*monitor).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*resp.MonitorGuid)

	return monitorCertificateRead(ctx, d, meta)
}

func monitorCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	resp, _, err := client.MonitorGetMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return readMonitorCertificateStruct(&resp, d)
}

func monitorCertificateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	monitor, err := buildMonitorCertificateStruct(d)
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

	return monitorCertificateRead(ctx, d, meta)
}

func monitorCertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	_, err := client.MonitorDeleteMonitor(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildMonitorCertificateStruct(d *schema.ResourceData) (*uptrends.Monitor, error) {
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

	monitor.MonitorType = &monitorCertificateType
	monitor.Url = String(d.Get("url").(string))
	monitor.IpVersion = ipVersion
	monitor.NativeIPv6Only = Bool(d.Get("native_ipv6_only").(bool))
	monitor.AuthenticationType = authType
	monitor.Username = String(d.Get("username").(string))
	monitor.AlertOnLoadTimeLimit1 = Bool(d.Get("alert_on_load_time_limit_1").(bool))
	monitor.LoadTimeLimit1 = Int32(int32(d.Get("load_time_limit_1").(int)))
	monitor.AlertOnLoadTimeLimit2 = Bool(d.Get("alert_on_load_time_limit_2").(bool))
	monitor.LoadTimeLimit2 = Int32(int32(d.Get("load_time_limit_2").(int)))
	monitor.CertificateName = String(d.Get("cert_name").(string))
	monitor.CertificateOrganization = String(d.Get("cert_org").(string))
	monitor.CertificateOrganizationalUnit = String(d.Get("cert_org_unit").(string))
	monitor.CertificateSerialNumber = String(d.Get("cert_serial_number").(string))
	monitor.CertificateFingerprint = String(d.Get("cert_fingerprint").(string))
	monitor.CertificateIssuerName = String(d.Get("cert_issuer_name").(string))
	monitor.CertificateIssuerCompanyName = String(d.Get("cert_issuer_company_name").(string))
	monitor.CertificateIssuerOrganizationalUnit = String(d.Get("cert_issuer_org_unit").(string))
	monitor.CertificateExpirationWarningDays = Int32(int32(d.Get("cert_expiration_warning_days").(int)))
	monitor.CheckCertificateErrors = Bool(d.Get("check_cert_errors").(bool))

	// Optional without defaults
	if attr, ok := d.GetOk("password"); ok {
		monitor.Password = String(attr.(string))
	}

	return monitor, nil
}

func readMonitorCertificateStruct(monitor *uptrends.Monitor, d *schema.ResourceData) diag.Diagnostics {
	err := readMonitorGenericStruct(monitor, d)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("url", monitor.Url); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ip_version", monitor.IpVersion); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("auth_type", monitor.AuthenticationType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("native_ipv6_only", monitor.NativeIPv6Only); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("auth_type", monitor.AuthenticationType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("username", monitor.Username); err != nil {
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
	if err := d.Set("cert_name", monitor.CertificateName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_org", monitor.CertificateOrganization); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_org_unit", monitor.CertificateOrganizationalUnit); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_serial_number", monitor.CertificateSerialNumber); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_fingerprint", monitor.CertificateFingerprint); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_issuer_name", monitor.CertificateIssuerName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_issuer_company_name", monitor.CertificateIssuerCompanyName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_issuer_org_unit", monitor.CertificateIssuerOrganizationalUnit); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_expiration_warning_days", monitor.CertificateExpirationWarningDays); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("check_cert_errors", monitor.CheckCertificateErrors); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("check_cert_errors", monitor.CheckCertificateErrors); err != nil {
		return diag.FromErr(err)
	}

	return nil
}