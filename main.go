package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
        "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/kerenbenzion/playtika_terraform_provider/drp"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() *schema.Provider {
                        return playtika.Provider()
                },
        })
}
