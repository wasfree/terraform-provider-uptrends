package uptrends

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/wasfree/uptrends-go-sdk"
)

var MonitorLoadTimeSchema = map[string]*schema.Schema{
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
}

func buildMonitorLoadTimeStruct(m *uptrends.Monitor, d *schema.ResourceData) error {

	m.AlertOnLoadTimeLimit1 = Bool(d.Get("alert_on_load_time_limit_1").(bool))
	m.LoadTimeLimit1 = Int32(int32(d.Get("load_time_limit_1").(int)))
	m.AlertOnLoadTimeLimit2 = Bool(d.Get("alert_on_load_time_limit_2").(bool))
	m.LoadTimeLimit2 = Int32(int32(d.Get("load_time_limit_2").(int)))

	return nil
}

func readMonitorLoadTimeStruct(m *uptrends.Monitor, d *schema.ResourceData) error {
	if err := d.Set("alert_on_load_time_limit_1", m.AlertOnLoadTimeLimit1); err != nil {
		return err
	}
	if err := d.Set("load_time_limit_1", m.LoadTimeLimit1); err != nil {
		return err
	}
	if err := d.Set("alert_on_load_time_limit_2", m.AlertOnLoadTimeLimit2); err != nil {
		return err
	}
	if err := d.Set("load_time_limit_2", m.LoadTimeLimit2); err != nil {
		return err
	}

	return nil
}
