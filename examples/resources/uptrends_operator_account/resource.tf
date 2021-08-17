resource "uptrends_operator_account" "example" {
  full_name          = "example"
  email              = "uptrends@example.org"
  password           = "secret"
  backup_email       = "uptrends1@example.org"
  mobile_phone       = "+00000000000"
  sms_provider       = "SmsProviderInternational"
  is_on_duty         = true
  use_numeric_sender = true
  default_dashboard  = "AlertLog"
}