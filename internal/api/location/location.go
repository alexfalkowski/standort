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

// ErrNotFound for location.
var ErrNotFound = errors.New("not found")

type (
	// Kind of location.
	Kind string

	// Location that are found.
	Location struct {
		Country   string
		Continent string
		Kind      Kind
	}

	// Point to look up.
	Point struct {
		Lat float64
		Lng float64
	}
)

const (
	// IP kind.
	IP Kind = Kind("ip")

	// GEO kind.
	GEO Kind = Kind("geo")
)

// NewLocator for the different transports.
func NewLocator(location *location.Location) *Locator {
	return &Locator{location: location}
}

// Locator for the different transports.
type Locator struct {
	location *location.Location
}

// Locate from IP and a point, it is returned in that order.
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
