package internal

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Config of the app.
	Config struct {
		Debug bool `envconfig:"debug"`

		Domain string `envconfig:"domain" default:"nurlashko.dev"`
	}
)

// ParseConfig environment variables and create Config.
func ParseConfig() (Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return cfg, err
}
