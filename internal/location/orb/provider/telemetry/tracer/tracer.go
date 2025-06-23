package tracer

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/telemetry/attributes"
	"github.com/alexfalkowski/go-service/v2/telemetry/tracer"
	"github.com/alexfalkowski/standort/v2/internal/location/orb/provider"
)

// Tracer is an alias for the tracer.Tracer.
type Tracer = tracer.Tracer

// NewProvider for tracer.
func NewProvider(provider provider.Provider, tracer *Tracer) *Provider {
	return &Provider{provider: provider, tracer: tracer}
}

// Provider for tracer.
type Provider struct {
	provider provider.Provider
	tracer   *Tracer
}

// Search a lat lng and get country and continent.
func (p *Provider) Search(ctx context.Context, lat, lng float64) (string, string, error) {
	ctx, span := p.tracer.StartClient(ctx, operationName("search"),
		attributes.Float64("provider.lat", lat),
		attributes.Float64("provider.lng", lng))
	defer span.End()

	country, continent, err := p.provider.Search(ctx, lat, lng)

	tracer.Meta(ctx, span)
	tracer.Error(err, span)

	return country, continent, err
}

func operationName(name string) string {
	return tracer.OperationName("orb", name)
}
