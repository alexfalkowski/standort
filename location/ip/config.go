package ip

// IsEnabled for ip.
func IsEnabled(cfg *Config) bool {
	return cfg != nil && cfg.Kind != ""
}

type Config struct {
	Kind string `yaml:"kind,omitempty" json:"kind,omitempty" toml:"kind,omitempty"`
}

// IsIP2location configured.
func (c *Config) IsIP2location() bool {
	return c.Kind == "ip2location"
}

// IsGeoIP2 configured.
func (c *Config) IsGeoIP2() bool {
	return c.Kind == "geoip2"
}
