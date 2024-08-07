package location

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/standort/location/continent"
	country "github.com/alexfalkowski/standort/location/country/provider"
	ip "github.com/alexfalkowski/standort/location/ip/provider"
	orb "github.com/alexfalkowski/standort/location/orb/provider"
)

// ErrNotFound for location.
var ErrNotFound = errors.New("not found")

// IsErrNotFound for location.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// Location will find the country and continent by different criteria.
type Location struct {
	ipProvider      ip.Provider
	countryProvider country.Provider
	orbProvider     orb.Provider
}

// New location.
func New(ipProvider ip.Provider, countryProvider country.Provider, orbProvider orb.Provider) *Location {
	return &Location{ipProvider: ipProvider, countryProvider: countryProvider, orbProvider: orbProvider}
}

// GetByIP a country and continent, otherwise error.
func (l *Location) GetByIP(ctx context.Context, ip string) (string, string, error) {
	c, err := l.ipProvider.GetByIP(ctx, ip)
	if err != nil {
		meta.WithAttribute(ctx, "ipError", meta.Error(err))

		return "", "", fmt.Errorf("%s: %w", ip, ErrNotFound)
	}

	cou, con, err := l.countryProvider.GetByCode(ctx, c)
	if err != nil {
		meta.WithAttribute(ctx, "countryError", meta.Error(err))

		return "", "", fmt.Errorf("%s: %w", ip, ErrNotFound)
	}

	return cou, continent.Codes[con], nil
}

// GetByLatLng a country and continent, otherwise error.
func (l *Location) GetByLatLng(ctx context.Context, lat, lng float64) (string, string, error) {
	cou, con := l.orbProvider.Search(ctx, lat, lng)
	if cou == "" || con == "" {
		return "", "", fmt.Errorf("%f/%f: %w", lat, lng, ErrNotFound)
	}

	return cou, continent.Codes[con], nil
}
