package uptrends

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	v4 "github.com/wasfree/uptrends-go-sdk"
)

type Uptrends struct {
	Client      *v4.APIClient
	AuthContext context.Context
}

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:             schema.TypeString,
				Required:         true,
				DefaultFunc:      schema.EnvDefaultFunc("UPTRENDS_USERNAME", nil),
				Description:      "Username for the uptrends API. Can be specified with the UPTRENDS_USERNAME environment variable.",
				ValidateDiagFunc: stringIsNotEmpty,
			},
			"password": {
				Type:             schema.TypeString,
				Required:         true,
				Sensitive:        true,
				DefaultFunc:      schema.EnvDefaultFunc("UPTRENDS_PASSWORD", nil),
				Description:      "Password for the uptrends API. Can be specified with the UPTRENDS_PASSWORD environment variable.",
				ValidateDiagFunc: stringIsNotEmpty,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"uptrends_monitor_ping": ResourceMonitorPingSchema(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	config := v4.NewConfiguration()
	config.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	// enable debug for http client
	config.Debug = true
	client := v4.NewAPIClient(config)

	authContext := context.WithValue(
		context.Background(),
		v4.ContextBasicAuth,
		v4.BasicAuth{
			UserName: d.Get("username").(string),
			Password: d.Get("password").(string),
		},
	)

	return &Uptrends{
		Client:      client,
		AuthContext: authContext,
	}, diags
}

func stringIsNotEmpty(i interface{}, k cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type of %q to be string", k)
	}

	if v == "" {
		return diag.Errorf("expected %q to not be an empty string, got %v", k, i)
	}

	return nil
}
