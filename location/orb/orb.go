package orb

import (
	"github.com/alexfalkowski/standort/location/continent"
	"github.com/alexfalkowski/standort/location/orb/provider"
	"github.com/alexfalkowski/standort/location/orb/provider/opentracing"
	"github.com/alexfalkowski/standort/location/orb/provider/rtree"
	"go.uber.org/fx"
)

// ProviderParams for orb.
type ProviderParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Config    *continent.Config
	Tracer    opentracing.Tracer
}

// NewProvider for orb.
func NewProvider(params ProviderParams) (provider.Provider, error) {
	var (
		p   provider.Provider
		err error
	)

	p, err = rtree.NewProvider(params.Config)
	if err != nil {
		return nil, err
	}

	p = opentracing.NewProvider(p, params.Tracer)

	return p, nil
}
