package ip

import (
	"embed"

	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/standort/v2/internal/location/ip/provider"
	"github.com/alexfalkowski/standort/v2/internal/location/ip/provider/geoip2"
	"github.com/alexfalkowski/standort/v2/internal/location/ip/provider/telemetry/tracer"
)

// ProviderParams for ip.
type ProviderParams struct {
	di.In

	Lifecycle di.Lifecycle
	FS        embed.FS
	Tracer    *tracer.Tracer
}

// NewProvider for ip.
func NewProvider(params ProviderParams) provider.Provider {
	var provider provider.Provider = geoip2.NewProvider(params.FS)
	provider = tracer.NewProvider(provider, params.Tracer)

	return provider
}
