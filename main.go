package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/tlindsay42/terraform-provider-armor/armor"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: armor.Provider})
}
