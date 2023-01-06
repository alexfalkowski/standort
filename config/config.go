package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/standort/health"
	"github.com/alexfalkowski/standort/location"
)

// Config for the service.
type Config struct {
	Health        health.Config   `yaml:"health" json:"health"`
	Location      location.Config `yaml:"location" json:"location"`
	config.Config `yaml:",inline" json:",inline"`
}
