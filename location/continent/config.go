package continent

// IsEnabled for continent.
func IsEnabled(c *Config) bool {
	return c != nil
}

// Config for continent.
type Config struct {
	AfricaPath       string `yaml:"africa_path,omitempty" json:"africa_path,omitempty" toml:"africa_path,omitempty"`
	NorthAmericaPath string `yaml:"north_america_path,omitempty" json:"north_america_path,omitempty" toml:"north_america_path,omitempty"`
	OceaniaPath      string `yaml:"oceania_path,omitempty" json:"oceania_path,omitempty" toml:"oceania_path,omitempty"`
	AsiaPath         string `yaml:"asia_path,omitempty" json:"asia_path,omitempty" toml:"asia_path,omitempty"`
	EuropePath       string `yaml:"europe_path,omitempty" json:"europe_path,omitempty" toml:"europe_path,omitempty"`
	SouthAmericaPath string `yaml:"south_america_path,omitempty" json:"south_america_path,omitempty" toml:"south_america_path,omitempty"`
}
