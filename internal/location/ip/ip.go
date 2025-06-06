package ip

import (
	"embed"

	"github.com/alexfalkowski/standort/v2/internal/location/ip/provider"
	"github.com/alexfalkowski/standort/v2/internal/location/ip/provider/geoip2"
	"github.com/alexfalkowski/standort/v2/internal/location/ip/provider/telemetry/tracer"
	"go.uber.org/fx"
)

// ProviderParams for ip.
type ProviderParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	FS        embed.FS
	Tracer    *tracer.Tracer
}

// NewProvider for ip.
func NewProvider(params ProviderParams) provider.Provider {
	var provider provider.Provider = geoip2.NewProvider(params.FS)
	provider = tracer.NewProvider(provider, params.Tracer)

	return provider
}
