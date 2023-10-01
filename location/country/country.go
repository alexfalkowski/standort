package country

import (
	"github.com/alexfalkowski/standort/location/country/provider"
	"github.com/alexfalkowski/standort/location/country/provider/gountries"
	"github.com/alexfalkowski/standort/location/country/provider/telemetry/tracer"
)

// NewProvider for country.
func NewProvider(t tracer.Tracer) provider.Provider {
	var p provider.Provider = gountries.NewProvider()
	p = tracer.NewProvider(p, t)

	return p
}
