
resource "uptrends_alert_definition" "andrew-test" {
  name      = "example"
  is_active = true

  monitors = [
    "example-monitor-guid-1",
    "example-monitor-guid-2",
    "example-monitor-guid-3"
  ]


  escalation_level {
    threshold_error_count = 3
    threshold_minutes     = 5
    reminders             = 2
    reminder_delay        = 10
    integrations          = ["example-integration-guid-1", "example-integration-guid-2"]
  }
}
