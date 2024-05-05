package client

import (
	v1 "github.com/alexfalkowski/standort/client/v1/config"
	v2 "github.com/alexfalkowski/standort/client/v2/config"
)

// IsEnabled for client.
func IsEnabled(cfg *Config) bool {
	return cfg != nil
}

// Config for client.
type Config struct {
	V1 *v1.Config `yaml:"v1,omitempty" json:"v1,omitempty" toml:"v1,omitempty"`
	V2 *v2.Config `yaml:"v2,omitempty" json:"v2,omitempty" toml:"v2,omitempty"`
}
