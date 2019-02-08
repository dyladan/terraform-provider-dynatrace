package dynatrace

import (
	"log"

	dt "github.com/dyladan/dynatrace-go-client/api"
)

// Config contains Dynatrace provider settings
type Config struct {
	APIKey string
	APIURL string
}

// Client returns a new client for accessing Dynatrace
func (c *Config) Client() (*dt.Client, error) {
	dtConfig := dt.Config{
		APIKey:  c.APIKey,
		BaseURL: c.APIURL,
	}

	client := dt.New(dtConfig)

	log.Printf("[INFO] New Dynatrace client configured")

	return &client, nil
}

// ProviderConfig for the custom provider
type ProviderConfig struct {
	Client *dt.Client
}
