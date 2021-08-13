resource "uptrends_monitor_certificate" "example" {
  name                       = "example"
  url                        = "https://example.org/"
  mode                       = "Staging"
  generate_alert             = false
  check_interval             = 10
  notes                      = "Managed by Terraform"
  primary_checkpoints_only   = false
  alert_on_load_time_limit_1 = true
  load_time_limit_1          = 3000
  alert_on_load_time_limit_2 = true
  load_time_limit_2          = 7000

  auth_type                 = "Basic"
  username                  = "user"
  password                  = "secret"

  cert_name = "www.example.org"
  cert_org = "Internet Corporation for Assigned Names and Numbers"
  cert_serial_number = "0F:BE:08:B0:85:4D:05:73:8A:B0:CC:E1:C9:AF:EE:C9"
  cert_fingerprint = "0A28A6EB176EA9CC596F4C73FD897EFBD32DCA2A"
  cert_issuer_name = "DigiCert TLS RSA SHA256 2020 CA1"
  cert_issuer_company_name = "DigiCert Inc"
  cert_expiration_warning_days = 14

  selected_checkpoints {
    regions           = ["Asia", "1005"]
    checkpoints       = ["Salzburg", "1"]
    exclude_locations = ["Vancouver"]
  }
}