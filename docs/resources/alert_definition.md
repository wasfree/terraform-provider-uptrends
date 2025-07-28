# uptrends_alert_definition Resource

Provides an Uptrends Alert Definition resource. This resource allows you to manage alert definitions, escalation levels, and configure which integrations receive alerts.

## Example Usage

```hcl
resource "uptrends_alert_definition" "example" {
  name       = "Example Alert Definition"
  is_active  = true

  monitors = [
    "monitor-guid-1",
    "monitor-guid-2"
  ]

  monitor_groups = [
    "monitor-group-guid-1"
  ]

  escalation_level {
    threshold_error_count = 3
    threshold_minutes     = 5
    reminders            = 2
    reminder_delay       = 10
    integrations         = [
      "integration-guid-1",
      "integration-guid-2"
    ]
  }

  escalation_level {
    threshold_error_count = 5
    threshold_minutes     = 10
    reminders            = 1
    reminder_delay       = 15
    integrations         = [
      "integration-guid-3"
    ]
  }
}
```

## Argument Reference

- `name` (String, Required): The name of the alert definition.
- `is_active` (Boolean, Optional): Whether the alert definition is active. Defaults to `true`.
- `monitors` (List of String, Optional): List of monitor GUIDs to associate with this alert definition.
- `monitor_groups` (List of String, Optional): List of monitor group GUIDs to associate with this alert definition.
- `escalation_level` (Block, Required): One or more escalation levels. Each block supports:
  - `threshold_error_count` (Number, Required): Number of errors before escalation triggers.
  - `threshold_minutes` (Number, Required): Time window in minutes for error threshold.
  - `reminders` (Number, Optional): Number of reminders to send. Defaults to `0`.
  - `reminder_delay` (Number, Optional): Delay in minutes between reminders. Defaults to `0`.
  - `integrations` (List of String, Required): List of integration GUIDs to notify at this escalation level.

## Import

This resource does not currently support import.
