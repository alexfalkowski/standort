package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/standort/client"
	"github.com/alexfalkowski/standort/health"
	"github.com/alexfalkowski/standort/location"
)

// Config for the service.
type Config struct {
	Health        health.Config   `yaml:"health" json:"health" toml:"health"`
	Client        client.Config   `yaml:"client" json:"client" toml:"client"`
	Location      location.Config `yaml:"location" json:"location" toml:"location"`
	config.Config `yaml:",inline" json:",inline" toml:",inline"`
}
