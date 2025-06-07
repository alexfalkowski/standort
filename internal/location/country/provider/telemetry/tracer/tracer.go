package tracer

import (
	"context"

	"github.com/alexfalkowski/go-service/v2/telemetry/attributes"
	"github.com/alexfalkowski/go-service/v2/telemetry/tracer"
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider"
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

// GetByCode a country and continent.
func (p *Provider) GetByCode(ctx context.Context, name string) (string, string, error) {
	ctx, span := p.tracer.StartClient(ctx, operationName("get"), attributes.String("provider.name", name))
	defer span.End()

	country, continent, err := p.provider.GetByCode(ctx, name)
	tracer.Error(err, span)
	tracer.Meta(ctx, span)

	return country, continent, err
}

func operationName(name string) string {
	return tracer.OperationName("country", name)
}
