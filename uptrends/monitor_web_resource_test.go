package uptrends

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type UptrendsMonitorWebResource struct{}

func TestAccUptrendsMonitorWeb_Basic(t *testing.T) {
	r := UptrendsMonitorWebResource{}
	randInt := acctest.RandIntRange(100000, 999999)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccUptrendsMonitorWebDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: r.basic(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorWebDestroyExists("uptrends_monitor_web.http"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "name", fmt.Sprintf("acctest-uptrends-monitor-web-http-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "url", "http://example.com/"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "mode", "Production"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "user_agent", UserAgent("chrome83")),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "auth_type", "None"),
					testAccCheckUptrendsMonitorWebDestroyExists("uptrends_monitor_web.https"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.https", "name", fmt.Sprintf("acctest-uptrends-monitor-web-https-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_web.https", "url", "https://example.com/"),
				),
			},
			{
				Config: r.update(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorWebDestroyExists("uptrends_monitor_web.http"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "name", fmt.Sprintf("acctest-uptrends-monitor-web-http-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "url", "http://example.org/"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "mode", "Staging"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "expected_http_status_code", "200"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "user_agent", UserAgent("chrome83_android")),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "auth_type", "Basic"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "username", "user"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "password", "secret"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "alert_on_min_bytes", "true"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.http", "min_bytes", "1024"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.https", "name", fmt.Sprintf("acctest-uptrends-monitor-web-https-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_web.https", "url", "https://example.org/"),
					resource.TestCheckResourceAttr("uptrends_monitor_web.https", "check_cert_errors", "false"),
				),
			},
		},
	})
}

func testAccUptrendsMonitorWebDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*Uptrends).Client.MonitorApi
	auth := testAccProvider.Meta().(*Uptrends).AuthContext

	for _, r := range s.RootModule().Resources {
		if r.Type != "uptrends_monitor_web" {
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

func testAccCheckUptrendsMonitorWebDestroyExists(n string) resource.TestCheckFunc {
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

func (r UptrendsMonitorWebResource) basic(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_web" "http" {
  name = "acctest-uptrends-monitor-web-http-%d"
  type = "Http"
  url  = "http://example.com/"
}

resource "uptrends_monitor_web" "https" {
  name = "acctest-uptrends-monitor-web-https-%d"
  type = "Https"
  url  = "https://example.com/"
}

resource "uptrends_monitor_web" "webservice_http" {
  name = "acctest-uptrends-monitor-webservice-http-%d"
  type = "WebserviceHttp"
  url  = "http://api.uptrends.com/v4/Checkpoint/Server/Ipv4"

  request_headers {
    name  = "accept"
    value = "application/json"
  }
}

resource "uptrends_monitor_web" "webservice_https" {
  name = "acctest-uptrends-monitor-webservice-https-%d"
  type = "WebserviceHttps"
  url  = "https://api.uptrends.com/v4/Checkpoint/Server/Ipv4"

  request_headers {
    name  = "accept"
    value = "application/json"
  }
}
`, randInt, randInt, randInt, randInt)
}

func (r UptrendsMonitorWebResource) update(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_web" "http" {
  name                       = "acctest-uptrends-monitor-web-http-%d"
  type                       = "Http"
  url                        = "http://example.org/"
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

resource "uptrends_monitor_web" "https" {
  name                       = "acctest-uptrends-monitor-web-https-%d"
  type                       = "Https"
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

  check_cert_errors         = false
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
`, randInt, randInt)
}
