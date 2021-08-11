package uptrends

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type UptrendsMonitorGroupResource struct{}

func TestAccUptrendsMonitorGroup_Basic(t *testing.T) {
	r := UptrendsMonitorGroupResource{}
	randInt := acctest.RandIntRange(100000, 999999)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccUptrendsMonitorGroupDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: r.basic(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorGroupDestroyExists("uptrends_monitor_group.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_group.test", "description", fmt.Sprintf("acctest-uptrends-monitor-group-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_group.test", "is_all", "false"),
				),
			},
			{
				Config: r.update(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsMonitorGroupDestroyExists("uptrends_monitor_group.test"),
					resource.TestCheckResourceAttr("uptrends_monitor_group.test", "description", fmt.Sprintf("acctest-uptrends-monitor-group-updated-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_monitor_group.test", "is_all", "false"),
				),
			},
		},
	})
}

func testAccUptrendsMonitorGroupDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*Uptrends).Client.MonitorGroupApi
	auth := testAccProvider.Meta().(*Uptrends).AuthContext

	for _, r := range s.RootModule().Resources {
		if r.Type != "uptrends_monitor_group" {
			continue
		}

		if monitorGroup, _, err := client.MonitorGroupGetMonitorGroup(auth, r.Primary.ID).Execute(); err == nil && monitorGroup.MonitorGroupGuid != nil {
			if err != nil {
				return fmt.Errorf("unable to get uptrends monitor group response: %s", err)
			}
			return fmt.Errorf("uptrends monitor still exists %s, %s", r.Primary.ID, *monitorGroup.MonitorGroupGuid)
		}

	}
	return nil
}

func testAccCheckUptrendsMonitorGroupDestroyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Uptrends).Client.MonitorGroupApi
		auth := testAccProvider.Meta().(*Uptrends).AuthContext

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("monitor group not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("monitor group ID is unset")
		}

		resp, _, err := client.MonitorGroupGetMonitorGroup(auth, rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("unable to read uptrends monitor group: %s", err)
		}

		if found := resp.MonitorGroupGuid; *found != rs.Primary.ID {
			return fmt.Errorf("uptrends monitor group not found")
		}

		return nil
	}
}

func (r UptrendsMonitorGroupResource) basic(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_group" "test" {
  description = "acctest-uptrends-monitor-group-%d"
}
`, randInt)
}

func (r UptrendsMonitorGroupResource) update(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_monitor_group" "test" {
  description = "acctest-uptrends-monitor-group-updated-%d"
}
`, randInt)
}
