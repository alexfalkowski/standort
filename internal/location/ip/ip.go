package ip

import (
	"embed"

	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/standort/v2/internal/location/ip/provider"
	"github.com/alexfalkowski/standort/v2/internal/location/ip/provider/geoip2"
)

// ProviderParams for ip.
type ProviderParams struct {
	di.In
	Lifecycle di.Lifecycle
	FS        embed.FS
}

// NewProvider for ip.
func NewProvider(params ProviderParams) provider.Provider {
	return geoip2.NewProvider(params.FS)
}
