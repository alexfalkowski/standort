package tracer

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/telemetry/attributes"
	"github.com/alexfalkowski/go-service/v2/telemetry/tracer"
	"github.com/alexfalkowski/standort/v2/internal/location/ip/provider"
)

// Tracer is an alias for the tracer.Tracer.
type Tracer = tracer.Tracer

// NewProvider for tracer.
func NewProvider(provider provider.Provider, tracer *Tracer) *Provider {
	return &Provider{provider: provider, tracer: tracer}
}

// Provider for tracer.
type Provider struct {
	provider provider.Provider
	tracer   *Tracer
}

// GetByIP a country.
func (p *Provider) GetByIP(ctx context.Context, ip string) (string, error) {
	ctx, span := p.tracer.StartClient(ctx, operationName("get"), attributes.String("provider.ip", ip))
	defer span.End()

	country, err := p.provider.GetByIP(ctx, ip)

	tracer.Error(err, span)
	tracer.Meta(ctx, span)

	return country, err
}

func operationName(name string) string {
	return tracer.OperationName("ip", name)
}
