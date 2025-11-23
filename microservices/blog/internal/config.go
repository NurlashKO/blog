package internal

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Config of the app.
	Config struct {
		Debug       bool   `envconfig:"debug" default:"false"`
		DatabaseURI string `envconfig:"DATABASE_URI" default:"postgres://nurlashko:tmp@postgres.database:5432/blog?sslmode=disable"`
	}
)

// ParseConfig environment variables and create Config.
func ParseConfig() (Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return cfg, err
}
