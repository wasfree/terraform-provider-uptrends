package uptrends

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/wasfree/uptrends-go-sdk"
)

// generic monitor fields - https://www.uptrends.com/support/kb/api/monitor-api-fields
var MonitorGenericSchema = map[string]*schema.Schema{
	"name": {
		Type:         schema.TypeString,
		Required:     true,
		Description:  "Display name for the Ping Monitor resource.",
		ValidateFunc: validation.StringLenBetween(1, 64),
	},
	"is_active": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Indicates whether the monitor is actively running in the account. Defaults to `true`.",
	},
	"generate_alert": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "When set to false, no alerts will be generated for this monitor in case of an error. Defaults to `true`.",
	},
	"check_interval": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      5,
		Description:  "Numeric value for the time interval between individual checks, in minutes. The maximum value is 240 (4 hours). The minimum value depends on the type of monitor. Defaults to `5`.",
		ValidateFunc: validation.IntBetween(1, 240),
	},
	"mode": {
		Type:         schema.TypeString,
		Optional:     true,
		Default:      "Production",
		Description:  "The monitor mode, either Development, Staging or Production. Defaults to `Production`.",
		ValidateFunc: validation.StringInSlice([]string{"Development", "Staging", "Production"}, false),
	},
	"notes": {
		Type:         schema.TypeString,
		Optional:     true,
		Description:  "Your notes for this monitor.",
		ValidateFunc: validation.StringIsNotEmpty,
	},
	"selected_checkpoints": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"checkpoints": {
					Type:        schema.TypeSet,
					Optional:    true,
					Description: "A checkpoint is a geographic location from which you can have your service uptime and performance checked periodically. Provide checkpoint ids or names.",
					Elem: &schema.Schema{
						Type: schema.TypeString,
						StateFunc: func(v interface{}) string {
							return CheckpointID(v.(string))
						},
					},
				},
				"regions": {
					Type:        schema.TypeSet,
					Optional:    true,
					Description: "A Region contains one or more checkpoints, just define a region if all checkpoints in a region should be used. Provide region id or name.",
					Elem: &schema.Schema{
						Type: schema.TypeString,
						StateFunc: func(v interface{}) string {
							return RegionID(v.(string))
						},
					},
				},
				"exclude_locations": {
					Type:        schema.TypeSet,
					Optional:    true,
					Description: "It is possible to keep an entire region of checkpoints (e.g. Canada) selected (with the benefit of automatically getting new checkpoints as they are added to that region) but have additional control over individual checkpoint locations that you want to skip. Provide checkpoint ids or names.",
					Elem: &schema.Schema{
						Type: schema.TypeString,
						StateFunc: func(v interface{}) string {
							return CheckpointID(v.(string))
						},
					},
				},
			},
		},
	},
	"primary_checkpoints_only": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Only set this to False when you’re sure you want to execute your monitor on non-primary checkpoints. Defaults to `true`.",
	},
	"is_locked": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "It specifies whether the monitor is currently locked for editing. This happens if the Support team is reviewing your monitor.",
	},
	"name_for_phone_alerts": {
		Type:         schema.TypeString,
		Optional:     true,
		Description:  "The value for the speech-friendly monitor name, if applicable. This is the monitor name we’ll use in text-to-speech phone alerting, provided that the ‘Use alternate monitor names’ option has been enabled in the phone alert integration.",
		ValidateFunc: validation.StringIsNotEmpty,
	},
}

func expandSelectedCheckpoints(input []interface{}) *uptrends.SelectedCheckpoints {
	if len(input) == 0 {
		return &uptrends.SelectedCheckpoints{}
	}
	v := input[0].(map[string]interface{})

	selectedCheckpoints := uptrends.SelectedCheckpoints{}
	if v["checkpoints"].(*schema.Set).Len() != 0 {
		selectedCheckpoints.Checkpoints = DeduplicateCheckpointIDs(v["checkpoints"].(*schema.Set).List())
	}

	if v["regions"].(*schema.Set).Len() != 0 {
		selectedCheckpoints.Regions = DeduplicateRegionIDs(v["regions"].(*schema.Set).List())
	}

	if v["exclude_locations"].(*schema.Set).Len() != 0 {
		selectedCheckpoints.ExcludeLocations = DeduplicateCheckpointIDs(v["exclude_locations"].(*schema.Set).List())
	}

	return &selectedCheckpoints
}

func flattenSelectedCheckpoints(input *uptrends.SelectedCheckpoints) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	selectedCheckpoints := make(map[string]interface{})
	if v := input.Checkpoints; v != nil {
		selectedCheckpoints["checkpoints"] = SliceInt32ToSliceString(*v)
	}

	if v := input.Regions; v != nil {
		selectedCheckpoints["regions"] = SliceInt32ToSliceString(*v)
	}

	if v := input.ExcludeLocations; v != nil {
		selectedCheckpoints["exclude_locations"] = SliceInt32ToSliceString(*v)
	}

	return []interface{}{selectedCheckpoints}
}

func buildMonitorGenericStruct(m *uptrends.Monitor, d *schema.ResourceData) error {

	monitorMode, err := uptrends.NewMonitorModeFromValue(d.Get("mode").(string))
	if err != nil {
		return err
	}
	m.Name = String(d.Get("name").(string))
	m.MonitorMode = monitorMode
	m.IsActive = Bool(d.Get("is_active").(bool))
	m.GenerateAlert = Bool(d.Get("generate_alert").(bool))
	m.CheckInterval = Int32(int32(d.Get("check_interval").(int)))
	m.UsePrimaryCheckpointsOnly = Bool(d.Get("primary_checkpoints_only").(bool))
	m.SelectedCheckpoints = expandSelectedCheckpoints(d.Get("selected_checkpoints").([]interface{}))

	if attr, ok := d.GetOk("notes"); ok {
		m.Notes = String(attr.(string))
	}

	if attr, ok := d.GetOk("name_for_phone_alerts"); ok {
		m.NameForPhoneAlerts = String(attr.(string))
	}

	return nil
}

func readMonitorGenericStruct(m *uptrends.Monitor, d *schema.ResourceData) error {
	if err := d.Set("name", m.Name); err != nil {
		return err
	}
	if err := d.Set("is_active", m.IsActive); err != nil {
		return err
	}
	if err := d.Set("is_locked", m.IsLocked); err != nil {
		return err
	}
	if err := d.Set("generate_alert", m.GenerateAlert); err != nil {
		return err
	}
	if err := d.Set("check_interval", m.CheckInterval); err != nil {
		return err
	}
	if err := d.Set("primary_checkpoints_only", m.UsePrimaryCheckpointsOnly); err != nil {
		return err
	}
	if err := d.Set("notes", m.Notes); err != nil {
		return err
	}
	if err := d.Set("name_for_phone_alerts", m.NameForPhoneAlerts); err != nil {
		return err
	}
	if sc := m.SelectedCheckpoints; *sc != (uptrends.SelectedCheckpoints{}) {
		if err := d.Set("selected_checkpoints", flattenSelectedCheckpoints(m.SelectedCheckpoints)); err != nil {
			return err
		}
	}

	return nil
}
