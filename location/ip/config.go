package ip

import (
	"github.com/alexfalkowski/standort/location/ip/provider/geoip2"
	"github.com/alexfalkowski/standort/location/ip/provider/ip2location"
)

type Config struct {
	Kind        string             `yaml:"kind,omitempty" json:"kind,omitempty" toml:"kind,omitempty"`
	IP2Location ip2location.Config `yaml:"ip2location,omitempty" json:"ip2location,omitempty" toml:"ip2location,omitempty"`
	GeoIP2      geoip2.Config      `yaml:"geoip2,omitempty" json:"geoip2,omitempty" toml:"geoip2,omitempty"`
}

// IsIP2location configured.
func (c *Config) IsIP2location() bool {
	return c.Kind == "ip2location"
}

// IsGeoIP2 configured.
func (c *Config) IsGeoIP2() bool {
	return c.Kind == "geoip2"
}
