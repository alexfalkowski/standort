package location

import (
	"fmt"

	geouri "git.jlel.se/jlelse/go-geouri"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
	"github.com/alexfalkowski/standort/v2/internal/location"
)

// ErrNotFound is returned by `(*Locator).Locate` when no location could be
// derived from either the IP-based lookup or the geolocation (lat/lng) lookup.
var ErrNotFound = errors.New("not found")

type (
	// Kind describes which input produced a `Location`.
	//
	// This is transport-facing (API) information and is primarily useful for
	// clients to understand whether a result came from IP-derived lookup or
	// geolocation-derived lookup.
	Kind string

	// Location is a transport-facing representation of a resolved country and
	// continent.
	//
	// The `Country` and `Continent` fields contain ISO-3166 alpha-2 codes (for
	// example, "US") and two-letter continent codes (for example, "NA")
	// respectively, as produced by the underlying domain `internal/location`
	// package.
	//
	// `Kind` indicates which lookup strategy produced this location.
	Location struct {
		Country   string
		Continent string
		Kind      Kind
	}

	// Point is a latitude/longitude pair used for geolocation lookup.
	//
	// Latitude and longitude are expected to be in degrees.
	Point struct {
		Lat float64
		Lng float64
	}
)

const (
	// IP indicates the location was derived from an IP address.
	IP Kind = Kind("ip")

	// GEO indicates the location was derived from a latitude/longitude point.
	GEO Kind = Kind("geo")
)

// NewLocator constructs a transport-facing `Locator`.
//
// It adapts the domain `internal/location.Location` service to transport needs,
// including reading fallback inputs from request metadata and recording partial
// lookup failures into metadata attributes.
func NewLocator(location *location.Location) *Locator {
	return &Locator{location: location}
}

// Locator provides a transport-friendly location lookup API.
//
// It can resolve a location from:
//
//   - an IP address (explicit parameter or request metadata), and/or
//   - a geographic point (explicit parameter or `Geolocation` metadata header).
//
// When an input is missing, the corresponding metadata value is used as a
// fallback:
//
//   - IP: `meta.IPAddr(ctx).Value()`
//   - geolocation: `meta.Geolocation(ctx)` parsed as a geo URI
//
// If a lookup attempt fails, the error is recorded as a metadata attribute and
// resolution continues with any other available inputs.
//
// If both lookups are unavailable or both fail to produce a result, `ErrNotFound`
// is returned.
type Locator struct {
	location *location.Location
}

// Locate attempts to resolve a location from an IP address and/or geographic point.
//
// Inputs are tried in the following order:
//
//  1. IP-based lookup (if an IP is available)
//  2. Geolocation-based lookup (if a point is available)
//
// The returned locations preserve that ordering: `(ipLocation, geoLocation, err)`.
//
// Error handling semantics:
//
//   - If an IP lookup is attempted and fails, the error is attached to the context
//     metadata under the `locationIpError` attribute and the method continues.
//   - If a point is derived from metadata and parsing fails, the error is attached
//     under `locationPointError` and the method continues.
//   - If a geolocation lookup is attempted and fails, the error is attached under
//     `locationLatLngError` and the method continues.
//   - If neither lookup produced a location, `ErrNotFound` is returned.
//
// On partial success (one lookup succeeds and the other fails or is missing), the
// successful location is returned and `err` is nil.
func (s *Locator) Locate(ctx context.Context, ip string, p *Point) (*Location, *Location, error) {
	var (
		ipLocation  *Location
		geoLocation *Location
	)

	if ip := s.ip(ctx, ip); !strings.IsEmpty(ip) {
		if country, continent, err := s.location.GetByIP(ctx, ip); err != nil {
			meta.WithAttribute(ctx, "locationIpError", meta.Error(err))
		} else {
			ipLocation = &Location{Country: country, Continent: continent, Kind: IP}
		}
	}

	p, err := s.point(ctx, p)
	if err != nil {
		meta.WithAttribute(ctx, "locationPointError", meta.Error(err))
	} else if p != nil {
		if country, continent, err := s.location.GetByLatLng(ctx, p.Lat, p.Lng); err != nil {
			meta.WithAttribute(ctx, "locationLatLngError", meta.Error(err))
		} else {
			geoLocation = &Location{Country: country, Continent: continent, Kind: GEO}
		}
	}

	if ipLocation == nil && geoLocation == nil {
		return nil, nil, ErrNotFound
	}

	return ipLocation, geoLocation, nil
}

func (s *Locator) ip(ctx context.Context, ip string) string {
	if !strings.IsEmpty(ip) {
		return ip
	}

	return meta.IPAddr(ctx).Value()
}

func (s *Locator) point(ctx context.Context, p *Point) (*Point, error) {
	if p != nil {
		return p, nil
	}

	loc := meta.Geolocation(ctx)
	if loc.IsEmpty() {
		return nil, nil //nolint:nilnil
	}

	geo, err := geouri.Parse(loc.Value())
	if err != nil {
		return nil, fmt.Errorf("geo uri: %w", err)
	}

	return &Point{Lat: geo.Latitude, Lng: geo.Longitude}, nil
}
