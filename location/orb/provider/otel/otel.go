package otel

import (
	"context"

	"github.com/alexfalkowski/go-service/otel"
	"github.com/alexfalkowski/go-service/version"
	"github.com/alexfalkowski/standort/location/orb/provider"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

// Tracer for otel.
type Tracer trace.Tracer

// NewTracer for otel.
func NewTracer(lc fx.Lifecycle, cfg *otel.Config, version version.Version) (Tracer, error) {
	return otel.NewTracer(otel.TracerParams{Lifecycle: lc, Name: "orb", Config: cfg, Version: version})
}

// Provider for otel.
type Provider struct {
	provider provider.Provider
	tracer   Tracer
}

// NewProvider for otel.
func NewProvider(provider provider.Provider, tracer Tracer) *Provider {
	return &Provider{provider: provider, tracer: tracer}
}

// Search a lat lng and get country and continent.
func (p *Provider) Search(ctx context.Context, lat, lng float64) (string, string) {
	attrs := []attribute.KeyValue{
		attribute.Key("provider.lat").Float64(lat),
		attribute.Key("provider.lng").Float64(lng),
	}

	ctx, span := p.tracer.Start(ctx, "search", trace.WithSpanKind(trace.SpanKindClient), trace.WithAttributes(attrs...))
	defer span.End()

	return p.provider.Search(ctx, lat, lng)
}
