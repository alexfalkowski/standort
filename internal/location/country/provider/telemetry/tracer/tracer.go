package tracer

import (
	"context"

	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/standort/internal/location/country/provider"
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

// GetByCode a country and continent.
func (p *Provider) GetByCode(ctx context.Context, name string) (string, string, error) {
	attrs := []attribute.KeyValue{
		attribute.Key("provider.name").String(name),
	}

	ctx, span := p.tracer.Start(ctx, operationName("get"), trace.WithSpanKind(trace.SpanKindClient), trace.WithAttributes(attrs...))
	defer span.End()

	ctx = tracer.WithTraceID(ctx, span)
	country, continent, err := p.provider.GetByCode(ctx, name)
	tracer.Error(err, span)
	tracer.Meta(ctx, span)

	return country, continent, err
}

func operationName(name string) string {
	return tracer.OperationName("country", name)
}
