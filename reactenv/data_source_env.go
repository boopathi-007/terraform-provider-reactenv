package reactenv

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joho/godotenv"
)

func DataSourceEnv() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEnvRead,
		Schema: map[string]*schema.Schema{
			"file_path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Path to the .env file",
			},
			"variables": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Key-value pairs from the .env file",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceEnvRead(d *schema.ResourceData, meta interface{}) error {
	filePath := d.Get("file_path").(string)

	envMap, err := godotenv.Read(filePath)
	if err != nil {
		return err
	}

	d.Set("variables", envMap)
	d.SetId(filePath)

	return nil
}
