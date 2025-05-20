package country

import (
	"github.com/alexfalkowski/go-service/v2/telemetry/tracer"
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider"
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider/gountries"
	tt "github.com/alexfalkowski/standort/v2/internal/location/country/provider/telemetry/tracer"
)

// NewProvider for country.
func NewProvider(tracer *tracer.Tracer) provider.Provider {
	var provider provider.Provider = gountries.NewProvider()
	provider = tt.NewProvider(provider, tracer)

	return provider
}
