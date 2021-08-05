package uptrends

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProvider *schema.Provider
var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
	testAccProvider = Provider()
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"uptrends": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("UPTRENDS_USERNAME"); err == "" {
		t.Fatal("UPTRENDS_USERNAME must be set for acceptance tests")
	}

	if err := os.Getenv("UPTRENDS_PASSWORD"); err == "" {
		t.Fatal("UPTRENDS_PASSWORD must be set for acceptance tests")
	}
}