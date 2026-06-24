package location

import (
	"fmt"

	geouri "git.jlel.se/jlelse/go-geouri"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
	"github.com/alexfalkowski/standort/v2/internal/diagnostics"
	"github.com/alexfalkowski/standort/v2/internal/location"
)

// ErrNotFound is returned by `(*Locator).Locate` when no location could be
// derived from either the IP-based lookup or the geolocation (lat/lng) lookup.
var ErrNotFound = errors.New("not found")

// Kind describes which input produced a `Location`.
//
// This is transport-facing (API) information and is primarily useful for
// clients to understand whether a result came from IP-derived lookup or
// geolocation-derived lookup.
type Kind string

// Location is a transport-facing representation of a resolved country and
// continent.
//
// The `Country` and `Continent` fields contain ISO-3166 alpha-2 codes (for
// example, "US") and two-letter continent codes (for example, "NA")
// respectively, as produced by the underlying domain `internal/location`
// package.
//
// `Kind` indicates which lookup strategy produced this location.
type Location struct {
	Country   string
	Continent string
	Kind      Kind
}

// Locations contains successful lookup locations.
type Locations struct {
	IP  *Location
	GEO *Location
}

// Point is a latitude/longitude pair used for geolocation lookup.
//
// Latitude and longitude are expected to be finite values in degrees. Latitude
// must be in the range [-90, 90], and longitude must be in the range [-180, 180].
// Invalid points produce `ErrNotFound` when no lookup succeeds. For v2,
// the terminal error can carry a latitude/longitude diagnostic for transport
// metadata.
type Point struct {
	Lat float64
	Lng float64
}

// IP indicates the location was derived from an IP address.
const IP Kind = Kind("ip")

// GEO indicates the location was derived from a latitude/longitude point.
const GEO Kind = Kind("geo")

// NewLocator constructs a transport-facing `Locator`.
//
// It adapts the domain `internal/location.Location` service to transport needs,
// including reading fallback inputs from request metadata and collecting terminal
// lookup diagnostics for the transport layer.
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
// Metadata is treated as transport context that has already crossed the
// appropriate trust boundary. This package does not validate whether forwarded
// IP metadata is trustworthy; that responsibility belongs to the
// transport/framework/deployment layer that populates the metadata.
//
// If a lookup attempt fails, the diagnostic value is recorded internally and
// resolution continues with any other available inputs.
//
// If both lookups are unavailable or both fail to produce a result, `ErrNotFound`
// is returned with any diagnostic values collected from failed attempts.
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
// The return values are ordered as `(locations, err)`.
//
// Error handling semantics:
//
//   - If an IP lookup is attempted and fails, an IP diagnostic is recorded and
//     the method continues.
//   - If a point is derived from metadata and parsing fails,
//     a point diagnostic is recorded and the method continues.
//   - If a geolocation lookup is attempted and fails, a latitude/longitude
//     diagnostic is recorded and the method continues.
//   - If neither lookup produced a location, `ErrNotFound` is returned with the
//     collected code-only diagnostics.
//
// On partial success (one lookup succeeds and the other fails or is missing), the
// successful location is returned, collected diagnostics are discarded, and `err`
// is nil.
func (s *Locator) Locate(ctx context.Context, ip string, p *Point) (*Locations, error) {
	var (
		locations = &Locations{}
		diag      diagnostics.Values
	)

	if ip := s.ip(ctx, ip); !strings.IsEmpty(ip) {
		if country, continent, err := s.location.GetByIP(ctx, ip); err != nil {
			diag = diagnostics.IPError(diag, diagnostics.NotFound)
		} else {
			locations.IP = &Location{Country: country, Continent: continent, Kind: IP}
		}
	}

	p, err := s.point(ctx, p)
	if err != nil {
		diag = diagnostics.PointError(diag, diagnostics.InvalidGeoURI)
	} else if p != nil {
		if country, continent, err := s.location.GetByLatLng(ctx, p.Lat, p.Lng); err != nil {
			diag = diagnostics.LatLngError(diag, locationErrorCode(err))
		} else {
			locations.GEO = &Location{Country: country, Continent: continent, Kind: GEO}
		}
	}

	if locations.IP == nil && locations.GEO == nil {
		return nil, diagnostics.Error(ErrNotFound, diag)
	}

	return locations, nil
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

func locationErrorCode(err error) diagnostics.Code {
	if errors.Is(err, location.ErrInvalidPoint) {
		return diagnostics.InvalidPoint
	}

	return diagnostics.NotFound
}
