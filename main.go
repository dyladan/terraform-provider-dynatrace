package main

import (
	"github.com/dyladan/terraform-provider-dynatrace/dynatrace"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return dynatrace.Provider()
		},
	})
}
