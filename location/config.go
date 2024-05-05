package location

import (
	"github.com/alexfalkowski/standort/location/continent"
	"github.com/alexfalkowski/standort/location/ip"
)

// IsEnabled for location.
func IsEnabled(cfg *Config) bool {
	return cfg != nil
}

// Config for location.
type Config struct {
	Continent *continent.Config `yaml:"continent,omitempty" json:"continent,omitempty" toml:"continent,omitempty"`
	IP        *ip.Config        `yaml:"ip,omitempty" json:"ip,omitempty" toml:"ip,omitempty"`
}
