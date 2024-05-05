package location

import (
	"github.com/alexfalkowski/standort/location/ip"
)

// IsEnabled for location.
func IsEnabled(cfg *Config) bool {
	return cfg != nil
}

// Config for location.
type Config struct {
	IP *ip.Config `yaml:"ip,omitempty" json:"ip,omitempty" toml:"ip,omitempty"`
}
