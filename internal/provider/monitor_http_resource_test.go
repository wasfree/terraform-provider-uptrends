package uptrends

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type UptrendsMonitorHttpResource struct{}

func TestAccUptrendsMonitorHttp_Basic(t *testing.T) {
	r := UptrendsMonitorHttpResource{}
	randInt := acctest.RandIntRange(100000, 999999)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccUptrendsMonitorHttpDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: r.basic(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorHttpDestroyExists("uptrends_monitor_http.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "name", fmt.Sprintf("acctest-uptrends-monitor-http-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "url", "http://example.com/"),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "mode", "Production"),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "user_agent", UserAgent("chrome83")),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "auth_type", "None"),
				),
			},
			{
				Config: r.update(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorHttpDestroyExists("uptrends_monitor_http.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "name", fmt.Sprintf("acctest-uptrends-monitor-http-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "url", "https://example.org/"),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "mode", "Staging"),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "expected_http_status_code", "200"),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "user_agent", UserAgent("chrome83_android")),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "auth_type", "Basic"),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "username", "user"),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "password", "secret"),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "alert_on_min_bytes", "true"),
					resource.TestCheckResourceAttr("uptrends_monitor_http.test", "min_bytes", "1024"),
				),
			},
		},
	})
}

func testAccUptrendsMonitorHttpDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*Uptrends).Client.MonitorApi
	auth := testAccProvider.Meta().(*Uptrends).AuthContext

	for _, r := range s.RootModule().Resources {
		if r.Type != "uptrends_monitor_http" {
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

func testAccCheckUptrendsMonitorHttpDestroyExists(n string) resource.TestCheckFunc {
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

func (r UptrendsMonitorHttpResource) basic(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_http" "test" {
  name = "acctest-uptrends-monitor-http-%d"
  url  = "http://example.com/"
}
`, randInt)
}

func (r UptrendsMonitorHttpResource) update(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_http" "test" {
  name                       = "acctest-uptrends-monitor-http-%d"
  type                       = "Http"
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
  alert_on_min_bytes         = true
  min_bytes                  = 1024

  expected_http_status_code = 200
  user_agent                = "chrome83_android"
  auth_type                 = "Basic"
  username                  = "user"
  password                  = "secret"

  request_headers {
    name  = "X-Test1-Header"
    value = "true"
  }

  request_headers {
    name  = "X-Test2-Header"
    value = "true"
  }

  match_pattern {
    pattern     = "example"
    is_positive = true
  }

  match_pattern {
    pattern     = "foo"
    is_positive = false
  }

  selected_checkpoints {
    regions           = ["Asia", "1005"]
    checkpoints       = ["Salzburg", "1"]
    exclude_locations = ["Vancouver"]
  }
}
`, randInt)
}