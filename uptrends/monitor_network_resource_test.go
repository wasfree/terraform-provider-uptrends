package uptrends

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type UptrendsMonitorNetworkResource struct{}

func TestAccUptrendsMonitorNetwork_Basic(t *testing.T) {
	r := UptrendsMonitorNetworkResource{}
	randInt := acctest.RandIntRange(100000, 999999)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccUptrendsMonitorNetworkDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: r.basic(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorNetworkDestroyExists("uptrends_monitor_network.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "name", fmt.Sprintf("acctest-uptrends-monitor-network-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "network_address", "8.8.8.8"),
				),
			},
			{
				Config: r.update(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorNetworkDestroyExists("uptrends_monitor_network.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "name", fmt.Sprintf("acctest-uptrends-monitor-network-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "network_address", "8.8.4.4"),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "mode", "Staging"),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "generate_alert", "false"),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "check_interval", "10"),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "notes", "Managed by Terraform"),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "alert_on_load_time_limit_1", "true"),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "load_time_limit_1", "3000"),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "alert_on_load_time_limit_2", "true"),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "load_time_limit_2", "7000"),
					resource.TestCheckResourceAttr("uptrends_monitor_network.test", "primary_checkpoints_only", "false"),
				),
			},
		},
	})
}

func testAccUptrendsMonitorNetworkDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*Uptrends).Client.MonitorApi
	auth := testAccProvider.Meta().(*Uptrends).AuthContext

	for _, r := range s.RootModule().Resources {
		if r.Type != "uptrends_monitor_network" {
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

func testAccCheckUptrendsMonitorNetworkDestroyExists(n string) resource.TestCheckFunc {
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

func (r UptrendsMonitorNetworkResource) basic(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_network" "test" {
  name            = "acctest-uptrends-monitor-network-%d"
  network_address = "8.8.8.8"
}
`, randInt)
}

func (r UptrendsMonitorNetworkResource) update(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_network" "test" {
  name                       = "acctest-uptrends-monitor-network-%d"
  network_address            = "8.8.4.4"
  mode                       = "Staging"
  generate_alert             = false
  check_interval             = 10
  notes                      = "Managed by Terraform"
  is_active                  = false
  primary_checkpoints_only   = false
  alert_on_load_time_limit_1 = true
  load_time_limit_1          = 3000
  alert_on_load_time_limit_2 = true
  load_time_limit_2          = 7000

  selected_checkpoints {
    regions           = ["Asia", "1005"]
    checkpoints       = ["Salzburg", "1"]
    exclude_locations = ["Vancouver"]
  }
}
`, randInt)
}
