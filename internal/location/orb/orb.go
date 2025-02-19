package orb

import (
	"embed"
	"errors"

	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/standort/internal/location/orb/provider"
	"github.com/alexfalkowski/standort/internal/location/orb/provider/rtree"
	tt "github.com/alexfalkowski/standort/internal/location/orb/provider/telemetry/tracer"
	"go.uber.org/fx"
)

// ErrNoProvider in the config.
var ErrNoProvider = errors.New("no provider configured")

// ProviderParams for orb.
type ProviderParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	FS        embed.FS
	Tracer    *tracer.Tracer
}

// NewProvider for orb.
func NewProvider(params ProviderParams) provider.Provider {
	var provider provider.Provider = rtree.NewProvider(params.FS)
	provider = tt.NewProvider(provider, params.Tracer)

	return provider
}
