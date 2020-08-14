package main

import (
	"github.com/EvilSuperstars/terraform-provider-get/get"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: get.Provider,
	})
}
