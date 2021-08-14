resource "uptrends_monitor_database_server" "example" {
  name                       = "example"
  network_address            = "example.org"
  type                       = "MSSQL"
  port                       = 1433
  check_interval             = 10
  notes                      = "Managed by Terraform"
  primary_checkpoints_only   = false
  alert_on_load_time_limit_1 = true
  load_time_limit_1          = 3000
  alert_on_load_time_limit_2 = true
  load_time_limit_2          = 7000

  username = "user"
  password = "secret"
  db_name  = "exampledb"

  selected_checkpoints {
    regions           = ["Asia", "1005"]
    checkpoints       = ["Salzburg", "1"]
    exclude_locations = ["Vancouver"]
  }
}