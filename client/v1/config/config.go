package config

import (
	"time"

	"github.com/alexfalkowski/go-service/retry"
)

// Config for client.
type Config struct {
	Host      string        `yaml:"host" json:"host" toml:"host"`
	Timeout   time.Duration `yaml:"timeout" json:"timeout" toml:"timeout"`
	Retry     retry.Config  `yaml:"retry" json:"retry" toml:"retry"`
	UserAgent string        `yaml:"user_agent" json:"user_agent" toml:"user_agent"`
}
