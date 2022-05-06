package opentracing

import (
	"context"

	"github.com/alexfalkowski/go-service/trace/opentracing"
	"github.com/alexfalkowski/standort/location/orb/provider"
	otr "github.com/opentracing/opentracing-go"
	"go.uber.org/fx"
)

// Tracer for opentracing.
type Tracer otr.Tracer

// StartSpanFromContext for opentracing.
func StartSpanFromContext(ctx context.Context, tracer Tracer, operation, method string, opts ...otr.StartSpanOption) (context.Context, otr.Span) {
	return opentracing.StartSpanFromContext(ctx, tracer, "orb", operation, method, opts...)
}

// NewTracer for opentracing.
func NewTracer(lc fx.Lifecycle, cfg *opentracing.Config) (Tracer, error) {
	return opentracing.NewTracer(opentracing.TracerParams{Lifecycle: lc, Name: "orb", Config: cfg})
}

// Provider for opentracing.
type Provider struct {
	provider provider.Provider
	tracer   Tracer
}

// NewProvider for opentracing.
func NewProvider(provider provider.Provider, tracer Tracer) *Provider {
	return &Provider{provider: provider, tracer: tracer}
}

// Search a lat lng and get country and continent.
func (p *Provider) Search(ctx context.Context, lat, lng float64) (string, string) {
	ctx, span := StartSpanFromContext(ctx, p.tracer, p.provider.String(), "search")
	defer span.Finish()

	span.SetTag("provider.lat", lat)
	span.SetTag("provider.lng", lng)

	return p.provider.Search(ctx, lat, lng)
}

// String for opentracing.
func (p *Provider) String() string {
	return "opentracing"
}
