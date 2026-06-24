package location

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
	v1 "github.com/alexfalkowski/standort/v2/api/standort/v1"
	"github.com/alexfalkowski/standort/v2/internal/location"
)

// NewLocator constructs a v1 response locator.
func NewLocator(location *location.Location) *Locator {
	return &Locator{location: location}
}

// Locator resolves v1 requests and builds generated v1 responses.
type Locator struct {
	location *location.Location
}

// LocateByIP resolves a v1 IP request.
func (l *Locator) LocateByIP(ctx context.Context, req *v1.GetLocationByIPRequest) (*v1.GetLocationByIPResponse, error) {
	country, continent, err := l.location.GetByIP(ctx, req.GetIp())
	if err != nil {
		return nil, err
	}

	return &v1.GetLocationByIPResponse{
		Meta:     meta.CamelStrings(ctx, strings.Empty),
		Location: toLocation(country, continent),
	}, nil
}

// LocateByLatLng resolves a v1 latitude/longitude request.
func (l *Locator) LocateByLatLng(ctx context.Context, req *v1.GetLocationByLatLngRequest) (*v1.GetLocationByLatLngResponse, error) {
	country, continent, err := l.location.GetByLatLng(ctx, req.GetLat(), req.GetLng())
	if err != nil {
		return nil, err
	}

	return &v1.GetLocationByLatLngResponse{
		Meta:     meta.CamelStrings(ctx, strings.Empty),
		Location: toLocation(country, continent),
	}, nil
}

func toLocation(country, continent string) *v1.Location {
	return &v1.Location{Country: country, Continent: continent}
}
