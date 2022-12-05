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
		Description:   "Manages an Uptrends SSL Certificate Monitor.",
		CreateContext: monitorCertificateCreate,
		ReadContext:   monitorCertificateRead,
		UpdateContext: monitorCertificateUpdate,
		DeleteContext: monitorCertificateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: MergeSchema(MonitorGenericSchema, MonitorLoadTimeSchema, map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The full URL of the appropriate website, page or service that you want to monitor. The URL should include “https://”, please also include a port number if you are using a non-default port number, e.g. https://your-domain.com:8080/your-page. You can also use a fixed IP address as part of the URL instead of a host name, if your server listens to incoming requests without a host name.",
				ValidateFunc: validation.IsURLWithHTTPS,
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
			"cert_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Specify the common name of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_org": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Specify the organization of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_org_unit": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Specify the organizational unit of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_serial_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Specify the serial number of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_fingerprint": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Specify the fingerprint of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_issuer_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Specify the issuer common name of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_issuer_company_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Specify the issuer company name of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_issuer_org_unit": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Specify the issuer org unit of the certificate. Defaults to empty string to not check this field.",
			},
			"cert_expiration_warning_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     7,
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

	return readMonitorCertificateStruct(resp, d)
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
	m := uptrends.NewMonitor()

	if err := buildMonitorGenericStruct(m, d); err != nil {
		return nil, err
	}
	if err := buildMonitorLoadTimeStruct(m, d); err != nil {
		return nil, err
	}

	m.MonitorType = &monitorCertificateType
	m.Url = String(d.Get("url").(string))
	m.IpVersion = (*uptrends.IpVersion)(String(d.Get("ip_version").(string)))
	m.AuthenticationType = (*uptrends.ApiHttpAuthenticationType)(String(d.Get("auth_type").(string)))
	m.Username = String(d.Get("username").(string))
	m.CertificateName = String(d.Get("cert_name").(string))
	m.CertificateOrganization = String(d.Get("cert_org").(string))
	m.CertificateOrganizationalUnit = String(d.Get("cert_org_unit").(string))
	m.CertificateSerialNumber = String(d.Get("cert_serial_number").(string))
	m.CertificateFingerprint = String(d.Get("cert_fingerprint").(string))
	m.CertificateIssuerName = String(d.Get("cert_issuer_name").(string))
	m.CertificateIssuerCompanyName = String(d.Get("cert_issuer_company_name").(string))
	m.CertificateIssuerOrganizationalUnit = String(d.Get("cert_issuer_org_unit").(string))
	m.CertificateExpirationWarningDays = Int32(int32(d.Get("cert_expiration_warning_days").(int)))
	m.CheckCertificateErrors = Bool(d.Get("check_cert_errors").(bool))

	// Optional without defaults
	if attr, ok := d.GetOk("password"); ok {
		m.Password = String(attr.(string))
	}

	return m, nil
}

func readMonitorCertificateStruct(m *uptrends.Monitor, d *schema.ResourceData) diag.Diagnostics {
	if err := readMonitorGenericStruct(m, d); err != nil {
		return diag.FromErr(err)
	}
	if err := readMonitorLoadTimeStruct(m, d); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("url", m.Url); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ip_version", m.IpVersion); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("auth_type", m.AuthenticationType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("username", m.Username); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_name", m.CertificateName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_org", m.CertificateOrganization); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_org_unit", m.CertificateOrganizationalUnit); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_serial_number", m.CertificateSerialNumber); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_fingerprint", m.CertificateFingerprint); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_issuer_name", m.CertificateIssuerName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_issuer_company_name", m.CertificateIssuerCompanyName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_issuer_org_unit", m.CertificateIssuerOrganizationalUnit); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cert_expiration_warning_days", m.CertificateExpirationWarningDays); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("check_cert_errors", m.CheckCertificateErrors); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("check_cert_errors", m.CheckCertificateErrors); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
