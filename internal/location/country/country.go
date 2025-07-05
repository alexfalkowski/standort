package country

import (
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider"
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider/gountries"
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider/telemetry/tracer"
)

// NewProvider for country.
func NewProvider(t *tracer.Tracer) provider.Provider {
	return tracer.NewProvider(gountries.NewProvider(), t)
}
