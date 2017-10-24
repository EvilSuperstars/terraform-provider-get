package main

import (
	"github.com/ewbankkit/terraform-provider-go-getter/getter"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: getter.Provider,
	})
}
