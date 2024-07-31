package main

import (
	"github.com/boopathi-007/terraform-provider-reactenv/reactenv"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"reactenv_file": reactenv.DataSourceEnv(),
		},
	}
}
