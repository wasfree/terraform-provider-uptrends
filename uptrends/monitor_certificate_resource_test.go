package uptrends

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type UptrendsMonitorCertificateResource struct{}

func TestAccUptrendsMonitorSSLCertificate_Basic(t *testing.T) {
	r := UptrendsMonitorCertificateResource{}
	randInt := acctest.RandIntRange(100000, 999999)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccUptrendsMonitorCertificateDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: r.basic(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorCertificateDestroyExists("uptrends_monitor_certificate.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "name", fmt.Sprintf("acctest-uptrends-monitor-certificate-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "url", "https://example.com/"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "mode", "Production"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "auth_type", "None"),
				),
			},
			{
				Config: r.update(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorCertificateDestroyExists("uptrends_monitor_certificate.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "name", fmt.Sprintf("acctest-uptrends-monitor-certificate-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "url", "https://example.org/"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "mode", "Staging"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "auth_type", "Basic"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "username", "user"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "password", "secret"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "cert_name", "www.example.org"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "cert_org", "Internet Corporation for Assigned Names and Numbers"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "cert_serial_number", "0F:BE:08:B0:85:4D:05:73:8A:B0:CC:E1:C9:AF:EE:C9"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "cert_fingerprint", "0A28A6EB176EA9CC596F4C73FD897EFBD32DCA2A"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "cert_issuer_company_name", "DigiCert Inc"),
					resource.TestCheckResourceAttr("uptrends_monitor_certificate.test", "cert_expiration_warning_days", "14"),
				),
			},
		},
	})
}

func testAccUptrendsMonitorCertificateDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*Uptrends).Client.MonitorApi
	auth := testAccProvider.Meta().(*Uptrends).AuthContext

	for _, r := range s.RootModule().Resources {
		if r.Type != "uptrends_monitor_certificate" {
			continue
		}

		if monitor, _, err := client.MonitorGetMonitor(auth, r.Primary.ID).Execute(); err == nil && monitor.MonitorGuid != nil {
			if err != nil {
				return fmt.Errorf("unable to get uptrends monitor response")
			}
			return fmt.Errorf("uptrends monitor still exists %s, %s", r.Primary.ID, *monitor.MonitorGuid)
		}

	}
	return nil
}

func testAccCheckUptrendsMonitorCertificateDestroyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Uptrends).Client.MonitorApi
		auth := testAccProvider.Meta().(*Uptrends).AuthContext

		time.Sleep(1 * time.Minute)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("monitor not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("monitor ID is unset")
		}

		resp, _, err := client.MonitorGetMonitor(auth, rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("unable to read uptrends monitor")
		}

		if found := resp.MonitorGuid; *found != rs.Primary.ID {
			return fmt.Errorf("uptrends monitor not found")
		}

		return nil
	}
}

func (r UptrendsMonitorCertificateResource) basic(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_certificate" "test" {
  name = "acctest-uptrends-monitor-certificate-%d"
  url  = "https://example.com/"
}
`, randInt)
}

func (r UptrendsMonitorCertificateResource) update(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_certificate" "test" {
  name                       = "acctest-uptrends-monitor-certificate-%d"
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
`, randInt)
}
