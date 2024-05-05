package config

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/standort/client"
	v1 "github.com/alexfalkowski/standort/client/v1/config"
	v2 "github.com/alexfalkowski/standort/client/v2/config"
	"github.com/alexfalkowski/standort/health"
	"github.com/alexfalkowski/standort/location"
	"github.com/alexfalkowski/standort/location/continent"
	"github.com/alexfalkowski/standort/location/ip"
)

// NewConfig for config.
func NewConfig(i *cmd.InputConfig) (*Config, error) {
	c := &Config{}

	return c, i.Unmarshal(c)
}

// Config for the service.
type Config struct {
	Location       *location.Config `yaml:"location,omitempty" json:"location,omitempty" toml:"location,omitempty"`
	Client         *client.Config   `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	Health         *health.Config   `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func ipConfig(cfg *Config) *ip.Config {
	if !location.IsEnabled(cfg.Location) {
		return nil
	}

	return cfg.Location.IP
}

func continentConfig(cfg *Config) *continent.Config {
	if !location.IsEnabled(cfg.Location) {
		return nil
	}

	return cfg.Location.Continent
}

func v1Client(cfg *Config) *v1.Config {
	if !client.IsEnabled(cfg.Client) {
		return nil
	}

	return cfg.Client.V1
}

func v2Client(cfg *Config) *v2.Config {
	if !client.IsEnabled(cfg.Client) {
		return nil
	}

	return cfg.Client.V2
}

func healthConfig(cfg *Config) *health.Config {
	return cfg.Health
}
