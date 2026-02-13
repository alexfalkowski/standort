package location

import (
	"fmt"

	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/strings"
	"github.com/alexfalkowski/standort/v2/internal/location/continent"
	country "github.com/alexfalkowski/standort/v2/internal/location/country/provider"
	ip "github.com/alexfalkowski/standort/v2/internal/location/ip/provider"
	orb "github.com/alexfalkowski/standort/v2/internal/location/orb/provider"
)

// New constructs a domain `Location` service.
//
// The service composes three providers:
//   - an IP provider (IP → ISO country code)
//   - a country provider (ISO country code → country + continent name)
//   - an orb provider (lat/lng → country + continent name, typically via point-in-polygon)
//
// The resulting API exposes higher-level methods that return a country code and
// a two-letter continent code.
func New(ipProvider ip.Provider, countryProvider country.Provider, orbProvider orb.Provider) *Location {
	return &Location{ipProvider: ipProvider, countryProvider: countryProvider, orbProvider: orbProvider}
}

// Location resolves country and continent information from different inputs.
//
// This is the domain service used by transports. It normalizes provider outputs
// into:
//   - ISO-3166 alpha-2 country codes (e.g. "US"), and
//   - two-letter continent codes (e.g. "NA").
//
// The continent code mapping is performed via `continent.Codes`.
type Location struct {
	ipProvider      ip.Provider
	countryProvider country.Provider
	orbProvider     orb.Provider
}

// GetByIP resolves a location from an IP address.
//
// It performs the following steps:
//  1. Resolve an ISO country code from the IP address via `ipProvider.GetByIP`.
//  2. Resolve the country code and continent name via `countryProvider.GetByCode`.
//  3. Map the continent name to a two-letter continent code via `continent.Codes`.
//
// It returns `(countryCode, continentCode, error)`. On error, both returned
// strings are empty.
func (l *Location) GetByIP(ctx context.Context, ip string) (string, string, error) {
	c, err := l.ipProvider.GetByIP(ctx, ip)
	if err != nil {
		return strings.Empty, strings.Empty, err
	}

	country, cont, err := l.countryProvider.GetByCode(ctx, c)
	if err != nil {
		return strings.Empty, strings.Empty, err
	}

	return country, continent.Codes[cont], nil
}

// GetByLatLng resolves a location from a latitude/longitude coordinate.
//
// It delegates to `orbProvider.Search` to perform the lookup (typically
// point-in-polygon). The returned continent name is mapped to a two-letter
// continent code via `continent.Codes`.
//
// Errors from the provider are wrapped with the input coordinate for context.
// On error, both returned strings are empty.
func (l *Location) GetByLatLng(ctx context.Context, lat, lng float64) (string, string, error) {
	cou, con, err := l.orbProvider.Search(ctx, lat, lng)
	if err != nil {
		return strings.Empty, strings.Empty, fmt.Errorf("%f/%f: %w", lat, lng, err)
	}

	return cou, continent.Codes[con], nil
}
