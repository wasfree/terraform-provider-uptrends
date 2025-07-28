package uptrends

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	uptrendsdk "github.com/wasfree/uptrends-go-sdk"
)

func resourceAlertDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertDefinitionCreate,
		Read:   resourceAlertDefinitionRead,
		Update: resourceAlertDefinitionUpdate,
		Delete: resourceAlertDefinitionDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"monitors": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"monitor_groups": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"escalation_level": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold_error_count": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"threshold_minutes": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"reminders": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"reminder_delay": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"integrations": {
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Required: true,
						},
						"extra_email_addresses": {
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAlertDefinitionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Uptrends).Client.AlertDefinitionApi
	auth := meta.(*Uptrends).AuthContext

	// 1. Create the alert definition
	alertDef := uptrendsdk.AlertDefinition{
		AlertName: String(d.Get("name").(string)),
		IsActive:  Bool(d.Get("is_active").(bool)),
	}
	created, _, err := client.AlertDefinitionCreateAlertDefinition(auth).AlertDefinition(alertDef).Execute()
	if err != nil {
		return fmt.Errorf("failed to create alert definition: %w", err)
	}
	if created.AlertDefinitionGuid == nil {
		return fmt.Errorf("alert definition created but no GUID returned")
	}
	alertDefGuid := *created.AlertDefinitionGuid
	d.SetId(alertDefGuid)

	// 2. Add monitors
	if v, ok := d.GetOk("monitors"); ok {
		for _, m := range v.([]interface{}) {
			_, _, err := client.AlertDefinitionAddMonitorToAlertDefinition(auth, alertDefGuid, m.(string)).Execute()
			if err != nil {
				return fmt.Errorf("failed to add monitor %s: %w", m.(string), err)
			}
		}
	}
	// 3. Add monitor groups
	if v, ok := d.GetOk("monitor_groups"); ok {
		for _, mg := range v.([]interface{}) {
			_, _, err := client.AlertDefinitionAddMonitorGroupToAlertDefinition(auth, alertDefGuid, mg.(string)).Execute()
			if err != nil {
				return fmt.Errorf("failed to add monitor group %s: %w", mg.(string), err)
			}
		}
	}

	// 4. Add integrations to escalation levels using POST only (no need to fetch escalation levels)
	escLevels := d.Get("escalation_level").([]interface{})
	if len(escLevels) > 3 {
		return fmt.Errorf("Uptrends only supports 3 escalation levels (IDs 1, 2, 3) via the API. You specified %d.", len(escLevels))
	}
	for idx, escRaw := range escLevels {
		esc := escRaw.(map[string]interface{})
		escId := int32(idx + 1) // Escalation level IDs are 1-based (1,2,3)
		for _, integ := range esc["integrations"].([]interface{}) {
			integrationGuid := integ.(string)
			integrationBody := map[string]interface{}{
				"IntegrationGuid":    integrationGuid,
				"IsActive":           true,
				"SendOkAlerts":       true,
				"SendReminderAlerts": true,
			}
			// Marshal the body
			bodyBytes, err := json.Marshal(integrationBody)
			if err != nil {
				return err
			}
			// Build the request URL
			basePath := meta.(*Uptrends).Client.GetConfig().Servers[0].URL
			url := basePath + "/AlertDefinition/" + alertDefGuid + "/EscalationLevel/" + fmt.Sprintf("%d", escId) + "/Integration"
			// Create the HTTP request
			req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
			if err != nil {
				return err
			}
			// Set headers
			req.Header.Set("Content-Type", "application/json")
			// Set basic auth
			username := meta.(*Uptrends).AuthContext.Value(uptrendsdk.ContextBasicAuth).(uptrendsdk.BasicAuth).UserName
			password := meta.(*Uptrends).AuthContext.Value(uptrendsdk.ContextBasicAuth).(uptrendsdk.BasicAuth).Password
			req.SetBasicAuth(username, password)
			// Send the request
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			if resp.StatusCode >= 300 {
				var respBody bytes.Buffer
				respBody.ReadFrom(resp.Body)
				return fmt.Errorf("failed to add integration %s to escalation level %d: %s", integrationGuid, escId, respBody.String())
			}
		}
	}

	return resourceAlertDefinitionRead(d, meta)
}

func resourceAlertDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Uptrends).Client.AlertDefinitionApi
	auth := meta.(*Uptrends).AuthContext
	guid := d.Id()

	alertDef, _, err := client.AlertDefinitionGetSpecifiedAlertDefinitions(auth, guid).Execute()
	if err != nil {
		return fmt.Errorf("failed to get alert definition: %w", err)
	}
	if alertDef == nil || alertDef.AlertDefinitionGuid == nil {
		d.SetId("")
		return nil
	}
	if err := d.Set("name", alertDef.GetAlertName()); err != nil {
		return err
	}
	if err := d.Set("is_active", alertDef.GetIsActive()); err != nil {
		return err
	}

	monitors, _, err := client.AlertDefinitionGetAllMembers(auth, guid).Execute()
	if err != nil {
		return fmt.Errorf("failed to get alert definition members: %w", err)
	}
	monitorGuids := []string{}
	monitorGroupGuids := []string{}
	for _, m := range monitors {
		if m.MonitorGuid != nil {
			monitorGuids = append(monitorGuids, *m.MonitorGuid)
		}
		if m.MonitorGroupGuid != nil {
			monitorGroupGuids = append(monitorGroupGuids, *m.MonitorGroupGuid)
		}
	}
	if err := d.Set("monitors", monitorGuids); err != nil {
		return err
	}
	if err := d.Set("monitor_groups", monitorGroupGuids); err != nil {
		return err
	}

	// 3. Get escalation levels
	escLevels, _, err := client.AlertDefinitionGetAllEscalationLevels(auth, guid).Execute()
	if err != nil {
		return fmt.Errorf("failed to get escalation levels: %w", err)
	}
	escList := make([]map[string]interface{}, len(escLevels))
	for i, esc := range escLevels {
		escMap := map[string]interface{}{
			"threshold_error_count": esc.GetThresholdErrorCount(),
			"threshold_minutes":     esc.GetThresholdMinutes(),
			"reminders":             esc.GetNumberOfReminders(),
			"reminder_delay":        esc.GetReminderDelay(),
		}
		// 4. Get integrations for this escalation level
		integrations, _, err := client.AlertDefinitionGetAllEscalationLevelIntegrations(auth, guid, esc.GetId()).Execute()
		if err != nil {
			return fmt.Errorf("failed to get integrations for escalation level %d: %w", esc.GetId(), err)
		}
		integrationGuids := []string{}
		for _, integ := range integrations {
			integrationGuids = append(integrationGuids, integ.IntegrationGuid)
		}
		escMap["integrations"] = integrationGuids
		escList[i] = escMap
	}
	if err := d.Set("escalation_level", escList); err != nil {
		return err
	}

	return nil
}

func resourceAlertDefinitionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Uptrends).Client.AlertDefinitionApi
	auth := meta.(*Uptrends).AuthContext
	guid := d.Id()

	// 1. Update alert definition name and is_active
	alertDef := uptrendsdk.AlertDefinition{
		AlertDefinitionGuid: &guid,
		AlertName:           String(d.Get("name").(string)),
		IsActive:            Bool(d.Get("is_active").(bool)),
	}
	_, err := client.AlertDefinitionPutAlertDefinition(auth, guid).AlertDefinition(alertDef).Execute()
	if err != nil {
		return fmt.Errorf("failed to update alert definition: %w", err)
	}

	// 2. Remove all monitors and groups
	members, _, err := client.AlertDefinitionGetAllMembers(auth, guid).Execute()
	if err != nil {
		return fmt.Errorf("failed to get alert definition members: %w", err)
	}
	for _, m := range members {
		if m.MonitorGuid != nil {
			_, err := client.AlertDefinitionRemoveMonitorFromAlertDefinition(auth, guid, *m.MonitorGuid).Execute()
			if err != nil {
				return fmt.Errorf("failed to remove monitor %s: %w", *m.MonitorGuid, err)
			}
		}
		if m.MonitorGroupGuid != nil {
			_, err := client.AlertDefinitionRemoveMonitorGroupFromAlertDefinition(auth, guid, *m.MonitorGroupGuid).Execute()
			if err != nil {
				return fmt.Errorf("failed to remove monitor group %s: %w", *m.MonitorGroupGuid, err)
			}
		}
	}
	// 3. Re-add monitors
	if v, ok := d.GetOk("monitors"); ok {
		for _, m := range v.([]interface{}) {
			_, _, err := client.AlertDefinitionAddMonitorToAlertDefinition(auth, guid, m.(string)).Execute()
			if err != nil {
				return fmt.Errorf("failed to add monitor %s: %w", m.(string), err)
			}
		}
	}
	// 4. Re-add monitor groups
	if v, ok := d.GetOk("monitor_groups"); ok {
		for _, mg := range v.([]interface{}) {
			_, _, err := client.AlertDefinitionAddMonitorGroupToAlertDefinition(auth, guid, mg.(string)).Execute()
			if err != nil {
				return fmt.Errorf("failed to add monitor group %s: %w", mg.(string), err)
			}
		}
	}

	// 5. Remove all escalation levels (by setting to inactive)
	escLevels, _, err := client.AlertDefinitionGetAllEscalationLevels(auth, guid).Execute()
	if err != nil {
		return fmt.Errorf("failed to get escalation levels: %w", err)
	}
	for _, esc := range escLevels {
		// Set inactive
		esc.IsActive = Bool(false)
		_, err := client.AlertDefinitionPatchAlertDefinitionEscalation(auth, guid, esc.GetId()).EscalationLevel(esc).Execute()
		if err != nil {
			return fmt.Errorf("failed to deactivate escalation level %d: %w", esc.GetId(), err)
		}
		// Remove integrations for this escalation level
		integrations, _, err := client.AlertDefinitionGetAllEscalationLevelIntegrations(auth, guid, esc.GetId()).Execute()
		if err != nil {
			return fmt.Errorf("failed to get integrations for escalation level %d: %w", esc.GetId(), err)
		}
		for _, integ := range integrations {
			_, err := client.AlertDefinitionRemoveIntegrationFromEscalationLevel(auth, guid, esc.GetId(), integ.IntegrationGuid).Execute()
			if err != nil {
				return fmt.Errorf("failed to remove integration %s from escalation level %d: %w", integ.IntegrationGuid, esc.GetId(), err)
			}
		}
	}

	escBlocks := d.Get("escalation_level").([]interface{})
	for idx, escRaw := range escBlocks {
		esc := escRaw.(map[string]interface{})
		escLevel := uptrendsdk.EscalationLevel{
			ThresholdErrorCount: Int32(int32(esc["threshold_error_count"].(int))),
			ThresholdMinutes:    Int32(int32(esc["threshold_minutes"].(int))),
			NumberOfReminders:   Int32(int32(esc["reminders"].(int))),
			ReminderDelay:       Int32(int32(esc["reminder_delay"].(int))),
			IsActive:            Bool(true),
		}
		_, err := client.AlertDefinitionPatchAlertDefinitionEscalation(auth, guid, int32(idx)).EscalationLevel(escLevel).Execute()
		if err != nil {
			return fmt.Errorf("failed to add escalation level %d: %w", idx, err)
		}
		for _, integ := range esc["integrations"].([]interface{}) {
			integrationGuid := integ.(string)
			integration := uptrendsdk.EscalationLevelIntegration{
				IntegrationGuid: &integrationGuid,
				IsActive:        Bool(true),
			}
			// Optionally support extra_email_addresses for email integrations
			if v, ok := esc["extra_email_addresses"]; ok {
				if emails, ok := v.([]interface{}); ok && len(emails) > 0 {
					emailStrs := make([]string, len(emails))
					for i, e := range emails {
						emailStrs[i] = e.(string)
					}
					integration.ExtraEmailAddresses = emailStrs
					integration.ExtraEmailAddressesSpecified = Bool(true)
				}
			}
			_, _, err := client.AlertDefinitionAddIntegrationToEscalationLevel(auth, guid, int32(idx)).EscalationLevelIntegration(integration).Execute()
			if err != nil {
				return fmt.Errorf("failed to add integration %s to escalation level %d: %w", integrationGuid, idx, err)
			}
		}
	}

	return resourceAlertDefinitionRead(d, meta)
}

func resourceAlertDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Uptrends).Client.AlertDefinitionApi
	auth := meta.(*Uptrends).AuthContext
	guid := d.Id()

	_, err := client.AlertDefinitionDeleteAlertDefinition(auth, guid).Execute()
	if err != nil {
		return fmt.Errorf("failed to delete alert definition: %w", err)
	}
	d.SetId("")
	return nil
}
