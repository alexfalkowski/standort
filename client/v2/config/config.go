package config

import (
	"github.com/alexfalkowski/go-service/config"
)

// Config for client.
type Config struct {
	config.Client `yaml:",inline" json:",inline" toml:",inline"`
}
