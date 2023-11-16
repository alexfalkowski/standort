package tracer

import (
	"context"

	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/go-service/version"
	"github.com/alexfalkowski/standort/location/country/provider"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

// Tracer for tracer.
type Tracer trace.Tracer

// NewTracer for tracer.
func NewTracer(lc fx.Lifecycle, cfg *tracer.Config, env env.Environment, ver version.Version) (Tracer, error) {
	return tracer.NewTracer(lc, "country", env, ver, cfg)
}

// Provider for tracer.
type Provider struct {
	provider provider.Provider
	tracer   Tracer
}

// NewProvider for tracer.
func NewProvider(provider provider.Provider, tracer Tracer) *Provider {
	return &Provider{provider: provider, tracer: tracer}
}

// GetByCode a country and continent.
func (p *Provider) GetByCode(ctx context.Context, name string) (string, string, error) {
	attrs := []attribute.KeyValue{
		attribute.Key("provider.name").String(name),
	}

	ctx, span := p.tracer.Start(ctx, "by-name", trace.WithSpanKind(trace.SpanKindClient), trace.WithAttributes(attrs...))
	defer span.End()

	country, continent, err := p.provider.GetByCode(ctx, name)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		return "", "", err
	}

	return country, continent, nil
}
