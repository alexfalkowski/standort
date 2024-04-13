package ip

import (
	"errors"

	"github.com/alexfalkowski/standort/location/ip/provider"
	"github.com/alexfalkowski/standort/location/ip/provider/geoip2"
	"github.com/alexfalkowski/standort/location/ip/provider/ip2location"
	"github.com/alexfalkowski/standort/location/ip/provider/telemetry/tracer"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

// ErrNoProvider in the config.
var ErrNoProvider = errors.New("no provider configured")

// ProviderParams for ip.
type ProviderParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Config    *Config
	Tracer    trace.Tracer
}

// NewProvider for ip.
func NewProvider(params ProviderParams) (provider.Provider, error) {
	var (
		provider provider.Provider
		err      error
	)

	if params.Config.IsIP2location() {
		provider, err = ip2location.NewProvider(params.Lifecycle, params.Config.IP2Location)
	}

	if params.Config.IsGeoIP2() {
		provider, err = geoip2.NewProvider(params.Config.GeoIP2)
	}

	if err != nil {
		return nil, err
	}

	if provider == nil {
		return nil, ErrNoProvider
	}

	provider = tracer.NewProvider(provider, params.Tracer)

	return provider, nil
}
