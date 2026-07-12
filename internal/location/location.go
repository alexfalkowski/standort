package location

import (
	"fmt"
	"math"
	"net"

	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/strings"
	"github.com/alexfalkowski/standort/v2/internal/location/continent"
	country "github.com/alexfalkowski/standort/v2/internal/location/country/provider"
	ip "github.com/alexfalkowski/standort/v2/internal/location/ip/provider"
	orb "github.com/alexfalkowski/standort/v2/internal/location/orb/provider"
)

// ErrInvalidIP is returned when an IP address is not a valid IPv4 or IPv6 address.
var ErrInvalidIP = errors.New("invalid ip")

// ErrInvalidPoint is returned when latitude or longitude is non-finite or outside
// the supported geographic coordinate range.
var ErrInvalidPoint = errors.New("invalid point")

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
// strings are empty. An input that is not valid IPv4 or IPv6 returns
// `ErrInvalidIP` before either provider is called; callers can classify that
// condition with `errors.Is`.
func (l *Location) GetByIP(ctx context.Context, ip string) (string, string, error) {
	if net.ParseIP(ip) == nil {
		return strings.Empty, strings.Empty, ErrInvalidIP
	}

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
// Non-finite values, latitudes outside [-90, 90], and longitudes outside
// [-180, 180] return an error wrapping `ErrInvalidPoint` before the provider is
// called. Provider errors are also wrapped with the input coordinate for
// context. Callers can classify the validation error with `errors.Is`. On
// error, both returned strings are empty.
func (l *Location) GetByLatLng(ctx context.Context, lat, lng float64) (string, string, error) {
	if !validPoint(lat, lng) {
		return strings.Empty, strings.Empty, fmt.Errorf("%f/%f: %w", lat, lng, ErrInvalidPoint)
	}

	cou, cont, err := l.orbProvider.Search(ctx, lat, lng)
	if err != nil {
		return strings.Empty, strings.Empty, fmt.Errorf("%f/%f: %w", lat, lng, err)
	}

	return cou, continent.Codes[cont], nil
}

func validPoint(lat, lng float64) bool {
	return !math.IsNaN(lat) && !math.IsNaN(lng) &&
		!math.IsInf(lat, 0) && !math.IsInf(lng, 0) &&
		lat >= -90 && lat <= 90 &&
		lng >= -180 && lng <= 180
}
