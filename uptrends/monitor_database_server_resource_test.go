package uptrends

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type UptrendsMonitorDatabaseServerResource struct{}

func TestAccUptrendsMonitorDatabaseServer_Basic(t *testing.T) {
	r := UptrendsMonitorDatabaseServerResource{}
	randInt := acctest.RandIntRange(100000, 999999)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccUptrendsMonitorDatabaseServerDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: r.basic(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorDatabaseServerDestroyExists("uptrends_monitor_database_server.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_database_server.test", "name", fmt.Sprintf("acctest-uptrends-monitor-database-server-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_database_server.test", "network_address", "example.com"),
					resource.TestCheckResourceAttr("uptrends_monitor_database_server.test", "mode", "Production"),
					resource.TestCheckResourceAttr("uptrends_monitor_database_server.test", "port", "3306"),
				),
			},
			{
				Config: r.update(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorDatabaseServerDestroyExists("uptrends_monitor_database_server.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_database_server.test", "name", fmt.Sprintf("acctest-uptrends-monitor-database-server-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_database_server.test", "network_address", "example.org"),
					resource.TestCheckResourceAttr("uptrends_monitor_database_server.test", "type", "MSSQL"),
					resource.TestCheckResourceAttr("uptrends_monitor_database_server.test", "mode", "Staging"),
					resource.TestCheckResourceAttr("uptrends_monitor_database_server.test", "port", "1433"),
					resource.TestCheckResourceAttr("uptrends_monitor_database_server.test", "username", "user"),
					resource.TestCheckResourceAttr("uptrends_monitor_database_server.test", "password", "secret"),
					resource.TestCheckResourceAttr("uptrends_monitor_database_server.test", "db_name", "exampledb"),
				),
			},
		},
	})
}

func testAccUptrendsMonitorDatabaseServerDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*Uptrends).Client.MonitorApi
	auth := testAccProvider.Meta().(*Uptrends).AuthContext

	for _, r := range s.RootModule().Resources {
		if r.Type != "uptrends_monitor_database_server" {
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

func testAccCheckUptrendsMonitorDatabaseServerDestroyExists(n string) resource.TestCheckFunc {
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

func (r UptrendsMonitorDatabaseServerResource) basic(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_database_server" "test" {
  name            = "acctest-uptrends-monitor-database-server-%d"
  network_address = "example.com"
  port            = 3306
}
`, randInt)
}

func (r UptrendsMonitorDatabaseServerResource) update(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_database_server" "test" {
  name                       = "acctest-uptrends-monitor-database-server-%d"
  mode                       = "Staging"
  network_address            = "example.org"
  type                       = "MSSQL"
  port                       = 1433
  generate_alert             = false
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
`, randInt)
}
