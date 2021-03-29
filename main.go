package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/helmwave/terraform-provider/pkg/helmwave"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: helmwave.Provider,
	})
}
