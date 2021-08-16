package uptrends

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type UptrendsMonitorFullPageCheckResource struct{}

func TestAccUptrendsMonitorFullPageCheck_Basic(t *testing.T) {
	r := UptrendsMonitorFullPageCheckResource{}
	randInt := acctest.RandIntRange(100000, 999999)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccUptrendsMonitorFullPageCheckDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: r.basic(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorFullPageCheckDestroyExists("uptrends_monitor_fpc.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "name", fmt.Sprintf("acctest-uptrends-monitor-fpc-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "mode", "Production"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "url", "https://example.com"),
				),
			},
			{
				Config: r.update(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorFullPageCheckDestroyExists("uptrends_monitor_fpc.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "name", fmt.Sprintf("acctest-uptrends-monitor-fpc-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "mode", "Staging"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "alert_on_min_bytes", "true"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "min_bytes", "1024"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "alert_on_max_bytes", "true"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "max_bytes", "2048"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "alert_on_max_element_bytes", "true"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "max_element_bytes", "4096"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "ignore_external_elements", "true"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "browser_type", "Firefox"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "browser_window_dimensions.0.is_mobile", "true"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "browser_window_dimensions.0.width", "1024"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "browser_window_dimensions.0.height", "768"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "browser_window_dimensions.0.pixel_ratio", "3"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "browser_window_dimensions.0.mobile_device", "Apple iPhone X"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "block_google_analytics", "true"),
					resource.TestCheckResourceAttr("uptrends_monitor_fpc.test", "block_uptrends_rum", "true"),
				),
			},
		},
	})
}

func testAccUptrendsMonitorFullPageCheckDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*Uptrends).Client.MonitorApi
	auth := testAccProvider.Meta().(*Uptrends).AuthContext

	for _, r := range s.RootModule().Resources {
		if r.Type != "uptrends_monitor_fpc" {
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

func testAccCheckUptrendsMonitorFullPageCheckDestroyExists(n string) resource.TestCheckFunc {
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

func (r UptrendsMonitorFullPageCheckResource) basic(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_fpc" "test" {
  name = "acctest-uptrends-monitor-fpc-%d"
  url  = "https://example.com"
}
`, randInt)
}

func (r UptrendsMonitorFullPageCheckResource) update(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_fpc" "test" {
  name           = "acctest-uptrends-monitor-fpc-%d"
  url            = "https://example.org"
  mode           = "Staging"
  generate_alert = false
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
`, randInt)
}
