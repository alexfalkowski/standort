package ip

import (
	"embed"

	"github.com/alexfalkowski/standort/location/ip/provider"
	"github.com/alexfalkowski/standort/location/ip/provider/geoip2"
	"github.com/alexfalkowski/standort/location/ip/provider/ip2location"
	"github.com/alexfalkowski/standort/location/ip/provider/telemetry/tracer"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

// ProviderParams for ip.
type ProviderParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Config    *Config
	FS        embed.FS
	Tracer    trace.Tracer
}

// NewProvider for ip.
func NewProvider(params ProviderParams) provider.Provider {
	provider := ipProvider(params.Lifecycle, params.Config, params.FS)
	provider = tracer.NewProvider(provider, params.Tracer)

	return provider
}

func ipProvider(lc fx.Lifecycle, cfg *Config, fs embed.FS) provider.Provider {
	var provider provider.Provider

	if !IsEnabled(cfg) || cfg.IsIP2location() {
		provider = ip2location.NewProvider(lc, fs)
	} else if cfg.IsGeoIP2() {
		provider = geoip2.NewProvider(fs)
	}

	return provider
}
