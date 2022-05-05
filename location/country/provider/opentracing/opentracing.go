package opentracing

import (
	"context"

	"github.com/alexfalkowski/go-service/trace/opentracing"
	"github.com/alexfalkowski/standort/location/country/provider"
	otr "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

// Tracer for opentracing.
type Tracer otr.Tracer

// StartSpanFromContext for opentracing.
func StartSpanFromContext(ctx context.Context, tracer Tracer, operation, method string, opts ...otr.StartSpanOption) (context.Context, otr.Span) {
	return opentracing.StartSpanFromContext(ctx, tracer, "country", operation, method, opts...)
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

// GetByName a country and continent.
func (p *Provider) GetByName(ctx context.Context, name string) (string, string, error) {
	ctx, span := StartSpanFromContext(ctx, p.tracer, p.provider.String(), "by-name")
	defer span.Finish()

	span.SetTag("provider.name", name)

	country, continent, err := p.provider.GetByName(ctx, name)
	if err != nil {
		ext.Error.Set(span, true)
		span.LogFields(log.String("event", "error"), log.String("message", err.Error()))

		return "", "", err
	}

	return country, continent, nil
}

// String for opentracing.
func (p *Provider) String() string {
	return "opentracing"
}
