package dynatrace

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DYNATRACE_API_KEY", nil),
				Sensitive:   true,
			},
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DYNATRACE_API_URL", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"dynatrace_application_detection_rule":       resourceApplicationDetectionRule(),
			"dynatrace_application_detection_rule_order": resourceApplicationDetectionRuleOrder(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIKey: data.Get("api_key").(string),
		APIURL: data.Get("api_url").(string),
	}

	log.Println("[INFO] Initializing Dynatrace client")
	client, err := config.Client()
	if err != nil {
		return nil, fmt.Errorf("Error initializing Dynatrace client: %s", err)
	}

	providerConfig := ProviderConfig{
		Client: client,
	}

	return &providerConfig, nil
}
