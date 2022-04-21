package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/standort/health"
	"github.com/alexfalkowski/standort/ip"
)

// Config for the service.
type Config struct {
	Health        health.Config `yaml:"health"`
	IP            ip.Config     `yaml:"ip"`
	config.Config `yaml:",inline"`
}
