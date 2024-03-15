package continent

import (
	"fmt"
)

// Config for continent.
type Config struct {
	AfricaPath       string `yaml:"africa_path,omitempty" json:"africa_path,omitempty" toml:"africa_path,omitempty"`
	NorthAmericaPath string `yaml:"north_america_path,omitempty" json:"north_america_path,omitempty" toml:"north_america_path,omitempty"`
	OceaniaPath      string `yaml:"oceania_path,omitempty" json:"oceania_path,omitempty" toml:"oceania_path,omitempty"`
	AsiaPath         string `yaml:"asia_path,omitempty" json:"asia_path,omitempty" toml:"asia_path,omitempty"`
	EuropePath       string `yaml:"europe_path,omitempty" json:"europe_path,omitempty" toml:"europe_path,omitempty"`
	SouthAmericaPath string `yaml:"south_america_path,omitempty" json:"south_america_path,omitempty" toml:"south_america_path,omitempty"`
}

// GetAfricaPath of geojson.
func (c *Config) GetAfricaPath() string {
	if c.AfricaPath != "" {
		return c.AfricaPath
	}

	return c.path("africa")
}

// GetNorthAmericaPath of geojson.
func (c *Config) GetNorthAmericaPath() string {
	if c.NorthAmericaPath != "" {
		return c.NorthAmericaPath
	}

	return c.path("north_america")
}

// GetOceaniaPath of geojson.
func (c *Config) GetOceaniaPath() string {
	if c.OceaniaPath != "" {
		return c.OceaniaPath
	}

	return c.path("oceania")
}

// GetAsiaPath of geojson.
func (c *Config) GetAsiaPath() string {
	if c.AsiaPath != "" {
		return c.AsiaPath
	}

	return c.path("asia")
}

// GetEuropePath of geojson.
func (c *Config) GetEuropePath() string {
	if c.EuropePath != "" {
		return c.EuropePath
	}

	return c.path("europe")
}

// GetSouthAmericaPath of geojson.
func (c *Config) GetSouthAmericaPath() string {
	if c.SouthAmericaPath != "" {
		return c.SouthAmericaPath
	}

	return c.path("south_america")
}

func (c *Config) path(name string) string {
	return fmt.Sprintf("/assets/%s.geojson", name)
}
