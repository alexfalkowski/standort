package location

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/location"
)

// NewLocator constructs a v2 response locator.
func NewLocator(locator *location.Locator) *Locator {
	return &Locator{locator: locator}
}

// Locator resolves v2 location requests and builds generated v2 responses.
type Locator struct {
	locator *location.Locator
}

// Locate resolves a v2 request.
func (l *Locator) Locate(ctx context.Context, req *v2.GetLocationRequest) (*v2.GetLocationResponse, error) {
	locations, err := l.locator.Locate(ctx, req.GetIp(), toPoint(req.GetPoint()))
	if err != nil {
		return nil, err
	}

	return &v2.GetLocationResponse{
		Meta: meta.CamelStrings(ctx, strings.Empty),
		Ip:   toLocation(locations.IP),
		Geo:  toLocation(locations.GEO),
	}, nil
}

func toPoint(p *v2.Point) *location.Point {
	if p == nil {
		return nil
	}

	return &location.Point{Lat: p.GetLat(), Lng: p.GetLng()}
}

func toLocation(l *location.Location) *v2.Location {
	if l == nil {
		return nil
	}

	return &v2.Location{Country: l.Country, Continent: l.Continent}
}
