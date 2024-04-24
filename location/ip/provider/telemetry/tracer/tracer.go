package tracer

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	tm "github.com/alexfalkowski/go-service/transport/meta"
	"github.com/alexfalkowski/standort/location/ip/provider"
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

// GetByIP a country.
func (p *Provider) GetByIP(ctx context.Context, ip string) (string, error) {
	attrs := []attribute.KeyValue{
		attribute.Key("provider.ip").String(ip),
	}

	ctx, span := p.tracer.Start(ctx, operationName("get"), trace.WithSpanKind(trace.SpanKindClient), trace.WithAttributes(attrs...))
	defer span.End()

	ctx = tm.WithTraceID(ctx, meta.ToValuer(span.SpanContext().TraceID()))
	country, err := p.provider.GetByIP(ctx, ip)
	tracer.Error(err, span)
	tracer.Meta(ctx, span)

	return country, err
}

func operationName(name string) string {
	return tracer.OperationName("ip", name)
}
