package config

import (
	"github.com/alexfalkowski/go-service/client"
)

// Config for client.
type Config struct {
	client.Config `yaml:",inline" json:",inline" toml:",inline"`
}
