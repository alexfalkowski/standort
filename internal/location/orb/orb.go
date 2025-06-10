package orb

import (
	"embed"

	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/telemetry/tracer"
	"github.com/alexfalkowski/standort/v2/internal/location/orb/provider"
	"github.com/alexfalkowski/standort/v2/internal/location/orb/provider/rtree"
	tt "github.com/alexfalkowski/standort/v2/internal/location/orb/provider/telemetry/tracer"
)

// Tracer is an alias for the tracer.Tracer.
type Tracer = tracer.Tracer

// ErrNoProvider in the config.
var ErrNoProvider = errors.New("no provider configured")

// ProviderParams for orb.
type ProviderParams struct {
	di.In

	Lifecycle di.Lifecycle
	FS        embed.FS
	Tracer    *Tracer
}

// NewProvider for orb.
func NewProvider(params ProviderParams) provider.Provider {
	var provider provider.Provider = rtree.NewProvider(params.FS)
	provider = tt.NewProvider(provider, params.Tracer)

	return provider
}
