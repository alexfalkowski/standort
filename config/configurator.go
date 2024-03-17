package config

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
	v1 "github.com/alexfalkowski/standort/client/v1/config"
	v2 "github.com/alexfalkowski/standort/client/v2/config"
	"github.com/alexfalkowski/standort/health"
	"github.com/alexfalkowski/standort/location/continent"
	"github.com/alexfalkowski/standort/location/ip"
)

// NewConfigurator for config.
func NewConfigurator(i *cmd.InputConfig) (config.Configurator, error) {
	c := &Config{}

	return c, i.Unmarshal(c)
}

func ipConfig(cfg config.Configurator) *ip.Config {
	return cfg.(*Config).Location.IP
}

func continentConfig(cfg config.Configurator) *continent.Config {
	return cfg.(*Config).Location.Continent
}

func v1Client(cfg config.Configurator) *v1.Config {
	return cfg.(*Config).Client.V1
}

func v2Client(cfg config.Configurator) *v2.Config {
	return cfg.(*Config).Client.V2
}

func healthConfig(cfg config.Configurator) *health.Config {
	return cfg.(*Config).Health
}
