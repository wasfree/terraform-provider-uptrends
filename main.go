package main

import (
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

  "github.com/wasfree/terraform-provider-uptends/uptrends"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: func() *schema.Provider {
      return uptrends.Provider()
    },
  })
}
