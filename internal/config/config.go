package config

import (
	"github.com/alexfalkowski/go-service/v2/config"
	"github.com/alexfalkowski/standort/v2/internal/health"
)

// Config is the top-level service configuration structure.
//
// It composes standort-specific configuration (for example `Health`) with the
// shared go-service configuration via an embedded `*config.Config`.
type Config struct {
	// Health configures the health subsystem (checks, observers, durations/timeouts, etc.).
	Health *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`

	// Config is the embedded go-service base configuration.
	//
	// The `yaml:",inline"` / `json:",inline"` / `toml:",inline"` tags make the
	// embedded fields appear at the top level of the config file.
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

// decorateConfig adapts `*Config` into the embedded `*config.Config` for consumers
// that depend on the go-service base configuration type.
func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

// healthConfig extracts the health configuration from the top-level config.
func healthConfig(cfg *Config) *health.Config {
	return cfg.Health
}
