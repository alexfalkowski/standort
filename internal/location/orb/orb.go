package orb

import (
	"embed"
	"errors"

	"github.com/alexfalkowski/standort/internal/location/orb/provider"
	"github.com/alexfalkowski/standort/internal/location/orb/provider/rtree"
	"github.com/alexfalkowski/standort/internal/location/orb/provider/telemetry/tracer"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

// ErrNoProvider in the config.
var ErrNoProvider = errors.New("no provider configured")

// ProviderParams for orb.
type ProviderParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	FS        embed.FS
	Tracer    trace.Tracer
}

// NewProvider for orb.
func NewProvider(params ProviderParams) provider.Provider {
	var p provider.Provider = rtree.NewProvider(params.FS)
	p = tracer.NewProvider(p, params.Tracer)

	return p
}
