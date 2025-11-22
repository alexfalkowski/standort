package country

import (
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider"
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider/gountries"
)

// NewProvider for country.
func NewProvider() provider.Provider {
	return gountries.NewProvider()
}
