package country

import (
	"github.com/alexfalkowski/standort/location/country/provider"
	"github.com/alexfalkowski/standort/location/country/provider/gountries"
	"github.com/alexfalkowski/standort/location/country/provider/opentracing"
)

// NewProvider for country.
func NewProvider(tracer opentracing.Tracer) provider.Provider {
	var p provider.Provider = gountries.NewProvider()
	p = opentracing.NewProvider(p, tracer)

	return p
}
