package ip

import (
	"errors"

	"github.com/alexfalkowski/standort/location/ip/provider"
	"github.com/alexfalkowski/standort/location/ip/provider/ip2location"
	"go.uber.org/fx"
)

// ErrNoProvider in the config.
var ErrNoProvider = errors.New("no provider configured")

// NewProvider for ip.
func NewProvider(lc fx.Lifecycle, cfg *Config) (provider.Provider, error) {
	if cfg.IsIP2location() {
		return ip2location.NewProvider(lc, &cfg.IP2Location)
	}

	return nil, ErrNoProvider
}
