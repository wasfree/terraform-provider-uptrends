package uptrends

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/wasfree/uptrends-go-sdk"
)

func ResourceMonitorGroupSchema() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages an Uptrends Monitor Group.",
		CreateContext: monitorGroupCreate,
		ReadContext:   monitorGroupRead,
		UpdateContext: monitorGroupUpdate,
		DeleteContext: monitorGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Set an descriptive name for the Monitor Group resource.",
				ValidateFunc: validation.StringLenBetween(1, 128),
			},
			"is_all": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this group is the `All Monitors` system group.",
			},
		},
	}
}

func monitorGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorGroupApi
	auth := meta.(*Uptrends).AuthContext

	monitorGroup := uptrends.MonitorGroup{
		Description: String(d.Get("description").(string)),
		IsAll:       false,
	}

	resp, _, err := client.MonitorGroupCreateMonitorGroup(auth).MonitorGroup(monitorGroup).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*resp.MonitorGroupGuid)

	return monitorGroupRead(ctx, d, meta)
}

func monitorGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorGroupApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	resp, _, err := client.MonitorGroupGetMonitorGroup(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", resp.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_all", resp.IsAll); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func monitorGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorGroupApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	monitorGroup := uptrends.MonitorGroup{
		MonitorGroupGuid: &id,
		IsAll:            false,
	}

	if d.HasChange("description") {
		monitorGroup.Description = String(d.Get("description").(string))
	}

	_, err := client.MonitorGroupUpdateMonitorGroup(auth, id).MonitorGroup(monitorGroup).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return monitorGroupRead(ctx, d, meta)
}

func monitorGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Uptrends).Client.MonitorGroupApi
	auth := meta.(*Uptrends).AuthContext

	id := d.Id()

	_, err := client.MonitorGroupDeleteMonitorGroup(auth, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
