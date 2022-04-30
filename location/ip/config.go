package ip

import (
	"github.com/alexfalkowski/standort/location/ip/provider/ip2location"
)

type Config struct {
	Type        string             `yaml:"type"`
	IP2Location ip2location.Config `yaml:"ip2location"`
}

// IsIP2location configured.
func (c *Config) IsIP2location() bool {
	return c.Type == "ip2location"
}
