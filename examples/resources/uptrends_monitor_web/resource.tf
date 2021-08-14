resource "uptrends_monitor_web" "test" {
  name                     = "example"
  check_interval           = 10
  notes                    = "Managed by Terraform"
  primary_checkpoints_only = false

  url                       = "http://example.org/"
  match_pattern             = "Example Domain"
  user_agent                = "chrome83_android"
  auth_type                 = "Basic"
  username                  = "user"
  password                  = "secret"
  expected_http_status_code = 200

  alert_on_load_time_limit_1 = true
  load_time_limit_1          = 3000
  alert_on_load_time_limit_2 = true
  load_time_limit_2          = 7000
  alert_on_min_bytes         = true
  min_bytes                  = 1024

  selected_checkpoints {
    regions           = ["Asia", "1005"]
    checkpoints       = ["Salzburg", "1"]
    exclude_locations = ["Vancouver"]
  }
}