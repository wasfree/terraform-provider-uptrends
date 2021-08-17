package uptrends

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type UptrendsOperatorAccountResource struct{}

func TestAccUptrendsOperatorAccount_Basic(t *testing.T) {
	r := UptrendsOperatorAccountResource{}
	randInt := acctest.RandIntRange(100000, 999999)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccUptrendsOperatorAccountDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: r.basic(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsOperatorAccountDestroyExists("uptrends_operator_account.test"),
					resource.TestCheckResourceAttr("uptrends_operator_account.test", "full_name", fmt.Sprintf("acctest-uptrends-operator-account-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_operator_account.test", "email", "uptrends@example.com"),
					resource.TestCheckResourceAttr("uptrends_operator_account.test", "password", fmt.Sprintf("secret%d", randInt)),
				),
			},
			{
				Config: r.update(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUptrendsOperatorAccountDestroyExists("uptrends_operator_account.test"),
					resource.TestCheckResourceAttr("uptrends_operator_account.test", "full_name", fmt.Sprintf("acctest-uptrends-operator-account-%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_operator_account.test", "email", "uptrends@example.org"),
					resource.TestCheckResourceAttr("uptrends_operator_account.test", "password", fmt.Sprintf("secret%d", randInt)),
					resource.TestCheckResourceAttr("uptrends_operator_account.test", "backup_email", "uptrends1@example.org"),
					resource.TestCheckResourceAttr("uptrends_operator_account.test", "mobile_phone", "+00000000000"),
					resource.TestCheckResourceAttr("uptrends_operator_account.test", "sms_provider", "SmsProviderInternational"),
					resource.TestCheckResourceAttr("uptrends_operator_account.test", "is_on_duty", "true"),
					resource.TestCheckResourceAttr("uptrends_operator_account.test", "use_numeric_sender", "true"),
					resource.TestCheckResourceAttr("uptrends_operator_account.test", "default_dashboard", "AlertLog"),
				),
			},
		},
	})
}

func testAccUptrendsOperatorAccountDestroyCheck(s *terraform.State) error {
	client := testAccProvider.Meta().(*Uptrends).Client.OperatorApi
	auth := testAccProvider.Meta().(*Uptrends).AuthContext

	for _, r := range s.RootModule().Resources {
		if r.Type != "uptrends_operator_account" {
			continue
		}

		if operator, _, err := client.OperatorGetOperator(auth, r.Primary.ID).Execute(); err == nil && operator.OperatorGuid != nil {
			if err != nil {
				return fmt.Errorf("unable to get uptrends operator response")
			}
			return fmt.Errorf("uptrends operator still exists %s, %s", r.Primary.ID, *operator.OperatorGuid)
		}

	}
	return nil
}

func testAccCheckUptrendsOperatorAccountDestroyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Uptrends).Client.OperatorApi
		auth := testAccProvider.Meta().(*Uptrends).AuthContext

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("operator not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("operator ID is unset")
		}

		resp, _, err := client.OperatorGetOperator(auth, rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("unable to read uptrends operator")
		}

		if found := resp.OperatorGuid; *found != rs.Primary.ID {
			return fmt.Errorf("uptrends operator not found")
		}

		return nil
	}
}

func (r UptrendsOperatorAccountResource) basic(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_operator_account" "test" {
  full_name = "acctest-uptrends-operator-account-%d"
  email     = "uptrends@example.com"
  password  = "secret%d"
}
`, randInt, randInt)
}

func (r UptrendsOperatorAccountResource) update(randInt int) string {
	return fmt.Sprintf(`
resource "uptrends_operator_account" "test" {
  full_name          = "acctest-uptrends-operator-account-%d"
  email              = "uptrends@example.org"
  password           = "secret%d"
  backup_email       = "uptrends1@example.org"
  mobile_phone       = "+00000000000"
  sms_provider       = "SmsProviderInternational"
  is_on_duty         = true
  use_numeric_sender = true
  default_dashboard  = "AlertLog"
}
`, randInt, randInt)
}
