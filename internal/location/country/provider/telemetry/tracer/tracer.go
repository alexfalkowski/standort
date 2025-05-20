package tracer

import (
	"context"

	"github.com/alexfalkowski/go-service/v2/telemetry/tracer"
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider"
	"go.opentelemetry.io/otel/attribute"
)

// NewProvider for tracer.
func NewProvider(provider provider.Provider, tracer *tracer.Tracer) *Provider {
	return &Provider{provider: provider, tracer: tracer}
}

// Provider for tracer.
type Provider struct {
	provider provider.Provider
	tracer   *tracer.Tracer
}

// GetByCode a country and continent.
func (p *Provider) GetByCode(ctx context.Context, name string) (string, string, error) {
	attrs := []attribute.KeyValue{
		attribute.Key("provider.name").String(name),
	}

	ctx, span := p.tracer.StartClient(ctx, operationName("get"), attrs...)
	defer span.End()

	country, continent, err := p.provider.GetByCode(ctx, name)
	tracer.Error(err, span)
	tracer.Meta(ctx, span)

	return country, continent, err
}

func operationName(name string) string {
	return tracer.OperationName("country", name)
}
