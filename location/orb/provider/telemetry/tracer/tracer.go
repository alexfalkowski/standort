package tracer

import (
	"context"

	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/standort/location/orb/provider"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Provider for tracer.
type Provider struct {
	provider provider.Provider
	tracer   trace.Tracer
}

// NewProvider for tracer.
func NewProvider(provider provider.Provider, tracer trace.Tracer) *Provider {
	return &Provider{provider: provider, tracer: tracer}
}

// Search a lat lng and get country and continent.
func (p *Provider) Search(ctx context.Context, lat, lng float64) (string, string, error) {
	attrs := []attribute.KeyValue{
		attribute.Key("provider.lat").Float64(lat),
		attribute.Key("provider.lng").Float64(lng),
	}

	ctx, span := p.tracer.Start(ctx, operationName("search"), trace.WithSpanKind(trace.SpanKindClient), trace.WithAttributes(attrs...))
	defer span.End()

	ctx = tracer.WithTraceID(ctx, span)
	country, continent, err := p.provider.Search(ctx, lat, lng)

	tracer.Meta(ctx, span)
	tracer.Error(err, span)

	return country, continent, err
}

func operationName(name string) string {
	return tracer.OperationName("orb", name)
}
