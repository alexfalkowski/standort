package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/standort/client"
	"github.com/alexfalkowski/standort/health"
	"github.com/alexfalkowski/standort/location"
)

// Config for the service.
type Config struct {
	Location      location.Config `yaml:"location,omitempty" json:"location,omitempty" toml:"location,omitempty"`
	Client        client.Config   `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	Health        health.Config   `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	config.Config `yaml:",inline" json:",inline" toml:",inline"`
}
