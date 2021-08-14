package uptrends

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type UptrendsMonitorDnsResource struct{}

func TestAccUptrendsMonitorDns_Basic(t *testing.T) {
	r := UptrendsMonitorDnsResource{}
	randInt := acctest.RandIntRange(100000, 999999)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccUptrendsMonitorDnsDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: r.basic(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorDnsDestroyExists("uptrends_monitor_dns.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_dns.test", "name", fmt.Sprintf("acctest-uptrends-monitor-dns-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_dns.test", "mode", "Production"),
				),
			},
			{
				Config: r.update(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorDnsDestroyExists("uptrends_monitor_dns.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_dns.test", "name", fmt.Sprintf("acctest-uptrends-monitor-dns-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_dns.test", "mode", "Staging"),
				),
			},
		},
	})
}

func testAccUptrendsMonitorDnsDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*Uptrends).Client.MonitorApi
	auth := testAccProvider.Meta().(*Uptrends).AuthContext

	for _, r := range s.RootModule().Resources {
		if r.Type != "uptrends_monitor_dns" {
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

func testAccCheckUptrendsMonitorDnsDestroyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Uptrends).Client.MonitorApi
		auth := testAccProvider.Meta().(*Uptrends).AuthContext

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

func (r UptrendsMonitorDnsResource) basic(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_dns" "test" {
  name                = "acctest-uptrends-monitor-dns-%d"
  dns_server          = "8.8.4.4"
  dns_query           = "ARecord"
  dns_test_value      = "example.org"
  dns_expected_result = "93.184.216.34"
}
`, randInt)
}

func (r UptrendsMonitorDnsResource) update(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_dns" "test" {
  name                       = "acctest-uptrends-monitor-dns-%d"
  mode                       = "Staging"
  generate_alert             = false
  check_interval             = 10
  notes                      = "Managed by Terraform"
  primary_checkpoints_only   = false
  alert_on_load_time_limit_1 = true
  load_time_limit_1          = 3000
  alert_on_load_time_limit_2 = true
  load_time_limit_2          = 7000
  port                       = 53

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
`, randInt)
}
