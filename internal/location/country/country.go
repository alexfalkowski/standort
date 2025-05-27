package country

import (
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider"
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider/gountries"
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider/telemetry/tracer"
)

// NewProvider for country.
func NewProvider(t *tracer.Tracer) provider.Provider {
	var provider provider.Provider = gountries.NewProvider()
	provider = tracer.NewProvider(provider, t)

	return provider
}
