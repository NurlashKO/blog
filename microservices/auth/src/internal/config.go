package internal

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Config of the app.
	Config struct {
		Debug bool `envconfig:"debug"`

		Domain string `envconfig:"domain" default:"nurlashko.dev"`

		// SAML configuration
		SAML struct {
			// EntityID is the unique identifier for this SAML identity provider
			EntityID string `envconfig:"saml_entity_id" default:"https://auth.nurlashko.dev/saml/metadata"`
			// URL is the base URL for SAML endpoints
			URL string `envconfig:"saml_url" default:"https://auth.nurlashko.dev/saml"`
			// KeyFile is the path to the private key file for signing SAML assertions
			KeyFile string `envconfig:"saml_key_file" default:"saml.key"`
			// CertFile is the path to the certificate file for signing SAML assertions
			CertFile string `envconfig:"saml_cert_file" default:"saml.crt"`
			// ServiceProviders is a list of allowed service provider entity IDs
			ServiceProviders []string `envconfig:"saml_service_providers" default:"https://vpn.nurlashko.dev/saml/metadata"`
		}
	}
)

// ParseConfig environment variables and create Config.
func ParseConfig() (Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return cfg, err
}
