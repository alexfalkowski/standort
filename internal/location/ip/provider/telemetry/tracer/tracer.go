package tracer

import (
	"context"

	"github.com/alexfalkowski/go-service/telemetry/tracer"
	"github.com/alexfalkowski/standort/internal/location/ip/provider"
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

// GetByIP a country.
func (p *Provider) GetByIP(ctx context.Context, ip string) (string, error) {
	attrs := []attribute.KeyValue{
		attribute.Key("provider.ip").String(ip),
	}

	ctx, span := p.tracer.StartClient(ctx, operationName("get"), attrs...)
	defer span.End()

	country, err := p.provider.GetByIP(ctx, ip)

	tracer.Error(err, span)
	tracer.Meta(ctx, span)

	return country, err
}

func operationName(name string) string {
	return tracer.OperationName("ip", name)
}
