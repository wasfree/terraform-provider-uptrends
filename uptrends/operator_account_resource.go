package uptrends

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/wasfree/uptrends-go-sdk"
)

func ResourceOperatorAccountSchema() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages Uptrends Operator Account.",
		CreateContext: monitorOperatorAccountCreate,
		ReadContext:   monitorOperatorAccountRead,
		UpdateContext: monitorOperatorAccountUpdate,
		DeleteContext: monitorOperatorAccountDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"full_name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Set the full name of the operator account owner.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"email": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Required:     true,
				Description:  "Set the email of the operator account owner.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"password": {
				Type:         schema.TypeString,
				Required:     true,
				Sensitive:    true,
				Description:  "Specify the corresponding password value here.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"backup_email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Set the backup email of the operator account owner.",
			},
			"mobile_phone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the mobile number of the Operator Account owner. Has to start with (+) followed by country code.",
			},
			"is_account_admin": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the operator is account administrator or not.",
			},
			"is_on_duty": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Set to true if this account should be on-duty.",
			},
			"sms_provider": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     string(uptrends.SMSPROVIDER_USE_ACCOUNT_SETTING),
				Description: "Define the sms provider that will be used sending alerts.",
				ValidateFunc: validation.StringInSlice([]string{
					string(uptrends.SMSPROVIDER_SMS_PROVIDER_EUROPE),
					string(uptrends.SMSPROVIDER_SMS_PROVIDER_EUROPE2),
					string(uptrends.SMSPROVIDER_SMS_PROVIDER_INTERNATIONAL),
					string(uptrends.SMSPROVIDER_SMS_PROVIDER_USA),
					string(uptrends.SMSPROVIDER_USE_ACCOUNT_SETTING),
				}, false),
			},
			"use_numeric_sender": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Define if the sms sender should be numeric.",
			},
			"default_dashboard": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "UseAccountSpecifiedDashboard",
				Description:  "Specify the dashboard that will be displayed after login.",
				ValidateFunc: validation.StringInSlice(uptrendsDashboards, false),
			},
			"allow_native_login": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Should native login be enabled on the uptrends ui?",
			},
			"allow_single_signon": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Should single signon be enabled on the uptrends ui?",
			},
			"culture_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Set the locale which should be used for language, date/time and number.",
				ValidateFunc: validation.StringInSlice([]string{"de-DE", "fr-FR", "en-US", "en-GB", "nl-NL"}, false),
			},
			"time_zone_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Define the ID for timezone which should be used for this account.",
			},
			"outgoing_phone_number_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Define the ID for outgoing phone number which should be used by Phone provider.",
			},
		},
	}
}

func monitorOperatorAccountCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.OperatorApi
	auth := meta.(*Uptrends).AuthContext

	operator, err := buildOperatorAccountStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, _, err := client.OperatorCreateOperator(auth).UptrendsOperator(*operator).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*resp.OperatorGuid)

	return monitorOperatorAccountRead(ctx, d, meta)
}

func monitorOperatorAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.OperatorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	resp, _, err := client.OperatorGetOperator(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return readMonitorOperatorAccountStruct(resp, d)
}

func monitorOperatorAccountUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.OperatorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	operator, err := buildOperatorAccountStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("password") {
		operator.Password = String(d.Get("password").(string))
	}

	_, err = client.OperatorUpdateOperator(auth, id).UptrendsOperator(*operator).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return monitorWebRead(ctx, d, meta)
}

func monitorOperatorAccountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.OperatorApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	_, err := client.OperatorDeleteOperator(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func readMonitorOperatorAccountStruct(o *uptrends.Operator, d *schema.ResourceData) diag.Diagnostics {
	if err := d.Set("full_name", o.FullName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("email", o.Email); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("backup_email", o.BackupEmail); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mobile_phone", o.MobilePhone); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_account_admin", o.IsAccountAdministrator); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_on_duty", o.IsOnDuty); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("sms_provider", o.SmsProvider); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("use_numeric_sender", o.UseNumericSender); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_dashboard", o.DefaultDashboard); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allow_native_login", o.AllowNativeLogin); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allow_single_signon", o.AllowSingleSignon); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("culture_name", o.CultureName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("time_zone_id", o.TimeZoneId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("outgoing_phone_number_id", o.OutgoingPhoneNumberId); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildOperatorAccountStruct(d *schema.ResourceData) (*uptrends.Operator, error) {
	o := uptrends.NewOperator()

	o.FullName = String(d.Get("full_name").(string))
	o.Email = String(d.Get("email").(string))
	o.BackupEmail = String(d.Get("backup_email").(string))
	o.MobilePhone = String(d.Get("mobile_phone").(string))
	o.IsAccountAdministrator = Bool(d.Get("is_account_admin").(bool))
	o.IsOnDuty = Bool(d.Get("is_on_duty").(bool))
	o.SmsProvider = (*uptrends.SmsProvider)(String(d.Get("sms_provider").(string)))
	o.UseNumericSender = Bool(d.Get("use_numeric_sender").(bool))
	o.DefaultDashboard = String(d.Get("default_dashboard").(string))

	// Optional without defaults
	if attr, ok := d.GetOk("time_zone_id"); ok {
		o.TimeZoneId = Int32(int32(attr.(int)))
	} else {
		o.TimeZoneIdSpecified = Bool(true)
	}

	if attr, ok := d.GetOk("culture_name"); ok {
		o.CultureName = String(attr.(string))
	} else {
		o.CultureNameSpecified = Bool(true)
	}

	if attr, ok := d.GetOk("allow_single_signon"); ok {
		o.AllowSingleSignon = Bool(attr.(bool))
	} else {
		o.AllowSingleSignonSpecified = Bool(true)
	}

	if attr, ok := d.GetOk("allow_native_login"); ok {
		o.AllowNativeLogin = Bool(attr.(bool))
	} else {
		o.AllowNativeLoginSpecified = Bool(true)
	}

	if attr, ok := d.GetOk("outgoing_phone_number_id"); ok {
		o.OutgoingPhoneNumberId = Int32(int32(attr.(int)))
	} else {
		o.OutgoingPhoneNumberIdSpecified = Bool(true)
	}

	if attr, ok := d.GetOk("password"); ok {
		o.Password = String(attr.(string))
	}

	return o, nil
}
