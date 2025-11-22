package orb

import (
	"embed"

	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/standort/v2/internal/location/orb/provider"
	"github.com/alexfalkowski/standort/v2/internal/location/orb/provider/rtree"
)

// ErrNoProvider in the config.
var ErrNoProvider = errors.New("no provider configured")

// ProviderParams for orb.
type ProviderParams struct {
	di.In
	Lifecycle di.Lifecycle
	FS        embed.FS
}

// NewProvider for orb.
func NewProvider(params ProviderParams) provider.Provider {
	return rtree.NewProvider(params.FS)
}
