package otel

import (
	"context"

	"github.com/alexfalkowski/go-service/otel"
	"github.com/alexfalkowski/go-service/version"
	"github.com/alexfalkowski/standort/location/ip/provider"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

// Tracer for otel.
type Tracer trace.Tracer

// NewTracer for otel.
func NewTracer(lc fx.Lifecycle, cfg *otel.Config, version version.Version) (Tracer, error) {
	return otel.NewTracer(otel.TracerParams{Lifecycle: lc, Name: "ip", Config: cfg, Version: version})
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

// GetByIP a country.
func (p *Provider) GetByIP(ctx context.Context, ip string) (string, error) {
	attrs := []attribute.KeyValue{
		attribute.Key("provider.ip").String(ip),
	}

	ctx, span := p.tracer.Start(ctx, "by-ip", trace.WithSpanKind(trace.SpanKindClient), trace.WithAttributes(attrs...))
	defer span.End()

	country, err := p.provider.GetByIP(ctx, ip)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		return "", err
	}

	return country, nil
}
