resource "uptrends_monitor_dns" "example" {
  name                       = "example"
  check_interval             = 10
  notes                      = "Managed by Terraform"
  primary_checkpoints_only   = false
  alert_on_load_time_limit_1 = true
  load_time_limit_1          = 3000
  alert_on_load_time_limit_2 = true
  load_time_limit_2          = 7000
  port = 53

  dns_server          = "8.8.4.4"
  dns_query           = "ARecord"
  dns_test_value      = "example.org"
  dns_expected_result = "93.184.216.34"

  selected_checkpoints {
    regions           = ["Asia", "1005"]
    checkpoints       = ["Salzburg", "1"]
    exclude_locations = ["Vancouver"]
  }
}