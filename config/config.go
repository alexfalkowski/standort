package config

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/standort/client"
	v1 "github.com/alexfalkowski/standort/client/v1/config"
	v2 "github.com/alexfalkowski/standort/client/v2/config"
	"github.com/alexfalkowski/standort/health"
)

// NewConfig for config.
func NewConfig(i *cmd.InputConfig) (*Config, error) {
	c := &Config{}

	return c, i.Unmarshal(c)
}

// IsEnabled for config.
func IsEnabled(cfg *Config) bool {
	return cfg != nil
}

// Config for the service.
type Config struct {
	Client         *client.Config `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	Health         *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	if !IsEnabled(cfg) {
		return nil
	}

	return cfg.Config
}

func v1Client(cfg *Config) *v1.Config {
	if !IsEnabled(cfg) || !client.IsEnabled(cfg.Client) {
		return nil
	}

	return cfg.Client.V1
}

func v2Client(cfg *Config) *v2.Config {
	if !IsEnabled(cfg) || !client.IsEnabled(cfg.Client) {
		return nil
	}

	return cfg.Client.V2
}

func healthConfig(cfg *Config) *health.Config {
	if !IsEnabled(cfg) {
		return nil
	}

	return cfg.Health
}
