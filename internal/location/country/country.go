package country

import (
	"github.com/alexfalkowski/standort/internal/location/country/provider"
	"github.com/alexfalkowski/standort/internal/location/country/provider/gountries"
	"github.com/alexfalkowski/standort/internal/location/country/provider/telemetry/tracer"
	"go.opentelemetry.io/otel/trace"
)

// NewProvider for country.
func NewProvider(t trace.Tracer) provider.Provider {
	var p provider.Provider = gountries.NewProvider()
	p = tracer.NewProvider(p, t)

	return p
}
