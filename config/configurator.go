package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/standort/health"
	"github.com/alexfalkowski/standort/location/continent"
	"github.com/alexfalkowski/standort/location/ip"
)

// NewConfigurator for config.
func NewConfigurator() config.Configurator {
	cfg := &Config{}

	return cfg
}

func healthConfig(cfg config.Configurator) *health.Config {
	return &cfg.(*Config).Health
}

func ipConfig(cfg config.Configurator) *ip.Config {
	return &cfg.(*Config).Location.IP
}

func continentConfig(cfg config.Configurator) *continent.Config {
	return &cfg.(*Config).Location.Continent
}
