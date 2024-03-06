package client

import (
	v1 "github.com/alexfalkowski/standort/client/v1/config"
	v2 "github.com/alexfalkowski/standort/client/v2/config"
)

// Config for client.
type Config struct {
	V1 v1.Config `yaml:"v1" json:"v1" toml:"v1"`
	V2 v2.Config `yaml:"v2" json:"v2" toml:"v2"`
}
