package continent

import (
	"fmt"
)

// Config for continent.
type Config struct {
	AfricaPath       string `yaml:"africa_path" json:"africa_path" toml:"africa_path"`
	NorthAmericaPath string `yaml:"north_america_path" json:"north_america_path" toml:"north_america_path"`
	OceaniaPath      string `yaml:"oceania_path" json:"oceania_path" toml:"oceania_path"`
	AsiaPath         string `yaml:"asia_path" json:"asia_path" toml:"asia_path"`
	EuropePath       string `yaml:"europe_path" json:"europe_path" toml:"europe_path"`
	SouthAmericaPath string `yaml:"south_america_path" json:"south_america_path" toml:"south_america_path"`
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
