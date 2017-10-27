package main

import (
	"github.com/EvilSuperstars/terraform-provider-get/get"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: get.Provider,
	})
}
