package service

import (
	"context"
	"errors"
	"fmt"

	geouri "git.jlel.se/jlelse/go-geouri"
	"github.com/alexfalkowski/go-service/meta"
	tm "github.com/alexfalkowski/go-service/transport/meta"
	"github.com/alexfalkowski/standort/location"
)

// ErrNotFound for service.
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

	// Service for the different transports.
	Service struct {
		location *location.Location
	}
)

const (
	// IP kind.
	IP Kind = Kind("ip")

	// GEO kind.
	GEO Kind = Kind("geo")
)

// IsNotFound for service.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// NewService for the different transports.
func NewService(location *location.Location) *Service {
	return &Service{location: location}
}

// GetLocations from IP and a point, it is returned in that order.
func (s *Service) GetLocations(ctx context.Context, ip string, p *Point) (*Location, *Location, error) {
	var (
		ipLocation  *Location
		geoLocation *Location
	)

	if ip := s.ip(ctx, ip); ip != "" {
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

func (s *Service) ip(ctx context.Context, ip string) string {
	if ip != "" {
		return ip
	}

	return tm.IPAddr(ctx).Value()
}

func (s *Service) point(ctx context.Context, p *Point) (*Point, error) {
	if p != nil {
		return p, nil
	}

	l := tm.Geolocation(ctx).Value()
	if l == "" {
		return nil, nil //nolint:nilnil
	}

	geo, err := geouri.Parse(l)
	if err != nil {
		return nil, fmt.Errorf("geo uri: %w", err)
	}

	return &Point{Lat: geo.Latitude, Lng: geo.Longitude}, nil
}
