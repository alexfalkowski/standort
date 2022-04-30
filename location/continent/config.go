package continent

import (
	"fmt"
)

// Config for continent.
type Config struct {
	AfricaPath       string `yaml:"africa_path"`
	NorthAmericaPath string `yaml:"north_america_path"`
	OceaniaPath      string `yaml:"oceania_path"`
	AntarcticaPath   string `yaml:"antarctica_path"`
	AsiaPath         string `yaml:"asia_path"`
	EuropePath       string `yaml:"europe_path"`
	SouthAmericaPath string `yaml:"south_america_path"`
}

func (c *Config) GetAfricaPath() string {
	if c.AfricaPath != "" {
		return c.AfricaPath
	}

	return c.path("africa")
}

func (c *Config) GetNorthAmericaPath() string {
	if c.NorthAmericaPath != "" {
		return c.NorthAmericaPath
	}

	return c.path("north_america")
}

func (c *Config) GetOceaniaPath() string {
	if c.OceaniaPath != "" {
		return c.OceaniaPath
	}

	return c.path("oceania")
}

func (c *Config) GetAntarcticaPath() string {
	if c.AntarcticaPath != "" {
		return c.AntarcticaPath
	}

	return c.path("antarctica")
}

func (c *Config) GetAsiaPath() string {
	if c.AsiaPath != "" {
		return c.AsiaPath
	}

	return c.path("asia")
}

func (c *Config) GetEuropePath() string {
	if c.EuropePath != "" {
		return c.EuropePath
	}

	return c.path("europe")
}

func (c *Config) GetSouthAmericaPath() string {
	if c.SouthAmericaPath != "" {
		return c.SouthAmericaPath
	}

	return c.path("south_america")
}

func (c *Config) path(name string) string {
	return fmt.Sprintf("assets/%s.geojson", name)
}
