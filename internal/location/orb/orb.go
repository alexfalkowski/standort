package orb

import (
	"embed"

	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/standort/v2/internal/location/orb/provider"
	"github.com/alexfalkowski/standort/v2/internal/location/orb/provider/rtree"
)

// ErrNoProvider indicates no orb (lat/lng) provider was configured.
var ErrNoProvider = errors.New("no provider configured")

// ProviderParams are dependency-injection inputs for constructing an orb (lat/lng) lookup provider.
type ProviderParams struct {
	di.In

	// Lifecycle is provided by the DI framework and can be used to hook provider
	// startup/shutdown. It is currently unused by the default provider implementation,
	// but is kept here to allow providers to manage resources if needed.
	Lifecycle di.Lifecycle

	// FS is the embedded filesystem containing runtime assets.
	//
	// The default provider reads `earth.geojson` from this filesystem to build its
	// point-in-polygon index.
	FS embed.FS
}

// NewProvider constructs the default orb (lat/lng) lookup provider.
//
// The returned provider resolves a latitude/longitude coordinate into:
//   - an ISO-3166 alpha-2 country code, and
//   - a continent name (for example, "Europe").
//
// The current implementation uses an R-tree backed point-in-polygon provider
// (`rtree.NewProvider`) built from the embedded `earth.geojson` dataset.
func NewProvider(params ProviderParams) provider.Provider {
	return rtree.NewProvider(params.FS)
}
