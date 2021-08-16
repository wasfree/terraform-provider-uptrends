resource "uptrends_monitor_fpc" "example" {
  name           = "example"
  url            = "https://example.org"
  check_interval = 10
  notes          = "Managed by Terraform"

  primary_checkpoints_only   = false
  alert_on_load_time_limit_1 = true
  load_time_limit_1          = 3000
  alert_on_load_time_limit_2 = true
  load_time_limit_2          = 7000
  alert_on_min_bytes         = true
  min_bytes                  = 1024
  alert_on_max_bytes         = true
  max_bytes                  = 2048
  alert_on_max_element_bytes = true
  max_element_bytes          = 4096
  ignore_external_elements   = true
  alert_on_percentage_fail   = true
  browser_type               = "Firefox"

  browser_window_dimensions {
    is_mobile     = true
    width         = 1024
    height        = 768
    pixel_ratio   = 3
    mobile_device = "Apple iPhone X"
  }

  block_google_analytics = true
  block_uptrends_rum     = true
  block_urls             = ["example.org"]

  user_agent = "firefox77"
  auth_type  = "Basic"
  username   = "user"
  password   = "secret"

  request_headers {
    name  = "X-Test1-Header"
    value = "true"
  }

  match_pattern {
    pattern     = "example"
    is_positive = true
  }

  selected_checkpoints {
    regions           = ["Asia", "1005"]
    checkpoints       = ["Salzburg", "1"]
    exclude_locations = ["Vancouver"]
  }
}