package tracer

import (
	"context"

	"github.com/alexfalkowski/go-service/v2/telemetry/tracer"
	"github.com/alexfalkowski/standort/internal/location/orb/provider"
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

// Search a lat lng and get country and continent.
func (p *Provider) Search(ctx context.Context, lat, lng float64) (string, string, error) {
	attrs := []attribute.KeyValue{
		attribute.Key("provider.lat").Float64(lat),
		attribute.Key("provider.lng").Float64(lng),
	}

	ctx, span := p.tracer.StartClient(ctx, operationName("search"), attrs...)
	defer span.End()

	country, continent, err := p.provider.Search(ctx, lat, lng)

	tracer.Meta(ctx, span)
	tracer.Error(err, span)

	return country, continent, err
}

func operationName(name string) string {
	return tracer.OperationName("orb", name)
}
