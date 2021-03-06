package geoip2

// Config for geoip2.
type Config struct {
	Path string `yaml:"path"`
}

// GetPath of config.
func (c *Config) GetPath() string {
	if c.Path != "" {
		return c.Path
	}

	return "/assets/geoip2.mmdb"
}
