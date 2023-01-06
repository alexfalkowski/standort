package ip2location

// Config for ip2location.
type Config struct {
	Path string `yaml:"path" json:"path"`
}

// GetPath of config.
func (c *Config) GetPath() string {
	if c.Path != "" {
		return c.Path
	}

	return "/assets/ip2location.bin"
}
