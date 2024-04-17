package ip2location

// IsEnabled for ip2location.
func IsEnabled(c *Config) bool {
	return c != nil
}

// Config for ip2location.
type Config struct {
	Path string `yaml:"path,omitempty" json:"path,omitempty" toml:"path,omitempty"`
}

// GetPath of config.
func (c *Config) GetPath() string {
	if c.Path != "" {
		return c.Path
	}

	return "/assets/ip2location.bin"
}
