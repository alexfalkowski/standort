package country

import (
	"github.com/alexfalkowski/standort/location/country/provider"
	"github.com/alexfalkowski/standort/location/country/provider/gountries"
	"github.com/alexfalkowski/standort/location/country/provider/otel"
)

// NewProvider for country.
func NewProvider(tracer otel.Tracer) provider.Provider {
	var p provider.Provider = gountries.NewProvider()
	p = otel.NewProvider(p, tracer)

	return p
}
