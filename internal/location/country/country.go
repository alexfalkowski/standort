package country

import (
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider"
	"github.com/alexfalkowski/standort/v2/internal/location/country/provider/gountries"
)

// NewProvider constructs the default country lookup provider.
//
// The returned provider resolves:
//   - ISO-3166 alpha-2 country codes (e.g. "US") to a country code and continent name.
//
// The current implementation uses the `gountries`-backed provider.
func NewProvider() provider.Provider {
	return gountries.NewProvider()
}
