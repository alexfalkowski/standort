package ip

import (
	"github.com/alexfalkowski/standort/location/ip/provider/geoip2"
	"github.com/alexfalkowski/standort/location/ip/provider/ip2location"
)

type Config struct {
	Kind        string             `yaml:"kind" json:"kind" toml:"kind"`
	IP2Location ip2location.Config `yaml:"ip2location" json:"ip2location" toml:"ip2location"`
	GeoIP2      geoip2.Config      `yaml:"geoip2" json:"geoip2" toml:"geoip2"`
}

// IsIP2location configured.
func (c *Config) IsIP2location() bool {
	return c.Kind == "ip2location"
}

// IsGeoIP2 configured.
func (c *Config) IsGeoIP2() bool {
	return c.Kind == "geoip2"
}
