package ip

import (
	"embed"

	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/standort/v2/internal/location/ip/provider"
	"github.com/alexfalkowski/standort/v2/internal/location/ip/provider/geoip2"
)

// ProviderParams are dependency-injection inputs for constructing an IP lookup provider.
type ProviderParams struct {
	di.In

	// Lifecycle is provided by the DI framework and can be used to hook provider
	// startup/shutdown. It is currently unused by the default provider implementation,
	// but is kept here to allow providers to manage resources (files, network clients, etc.)
	// if needed.
	Lifecycle di.Lifecycle

	// FS is the embedded filesystem containing runtime assets.
	//
	// The default provider reads `geoip2.mmdb` from this filesystem.
	FS embed.FS
}

// NewProvider constructs the default IP lookup provider.
//
// The returned provider resolves an IP address to an ISO-3166 alpha-2 country code.
// The current implementation uses the embedded GeoIP2 database via `geoip2.NewProvider`.
func NewProvider(params ProviderParams) provider.Provider {
	return geoip2.NewProvider(params.FS)
}
