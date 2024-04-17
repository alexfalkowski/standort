package orb

import (
	"errors"

	"github.com/alexfalkowski/standort/location/continent"
	"github.com/alexfalkowski/standort/location/orb/provider"
	"github.com/alexfalkowski/standort/location/orb/provider/rtree"
	"github.com/alexfalkowski/standort/location/orb/provider/telemetry/tracer"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

// ErrNoProvider in the config.
var ErrNoProvider = errors.New("no provider configured")

// ProviderParams for orb.
type ProviderParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Config    *continent.Config
	Tracer    trace.Tracer
}

// NewProvider for orb.
func NewProvider(params ProviderParams) (provider.Provider, error) {
	var (
		p   provider.Provider
		err error
	)

	if !continent.IsEnabled(params.Config) {
		return nil, ErrNoProvider
	}

	p, err = rtree.NewProvider(params.Config)
	if err != nil {
		return nil, err
	}

	p = tracer.NewProvider(p, params.Tracer)

	return p, nil
}
