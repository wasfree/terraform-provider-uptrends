package uptrends

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/wasfree/uptrends-go-sdk"
)

type Uptrends struct {
	Client      *uptrends.APIClient
	AuthContext context.Context
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:         schema.TypeString,
				Required:     true,
				DefaultFunc:  schema.EnvDefaultFunc("UPTRENDS_USERNAME", nil),
				Description:  "Username for the uptrends API. Can be specified with the UPTRENDS_USERNAME environment variable.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"password": {
				Type:         schema.TypeString,
				Required:     true,
				Sensitive:    true,
				DefaultFunc:  schema.EnvDefaultFunc("UPTRENDS_PASSWORD", nil),
				Description:  "Password for the uptrends API. Can be specified with the UPTRENDS_PASSWORD environment variable.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"uptrends_monitor_group":   ResourceMonitorGroupSchema(),
			"uptrends_monitor_web":     ResourceMonitorWebSchema(),
			"uptrends_monitor_network":    ResourceMonitorNetworkSchema(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	config := uptrends.NewConfiguration()
	config.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	// enable debug for http client
	config.Debug = true
	client := uptrends.NewAPIClient(config)

	authContext := context.WithValue(
		context.Background(),
		uptrends.ContextBasicAuth,
		uptrends.BasicAuth{
			UserName: d.Get("username").(string),
			Password: d.Get("password").(string),
		},
	)

	return &Uptrends{
		Client:      client,
		AuthContext: authContext,
	}, diags
}
