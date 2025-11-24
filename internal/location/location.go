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

// New location.
func New(ipProvider ip.Provider, countryProvider country.Provider, orbProvider orb.Provider) *Location {
	return &Location{ipProvider: ipProvider, countryProvider: countryProvider, orbProvider: orbProvider}
}

// Location will find the country and continent by different criteria.
type Location struct {
	ipProvider      ip.Provider
	countryProvider country.Provider
	orbProvider     orb.Provider
}

// GetByIP a country and continent, otherwise error.
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

// GetByLatLng a country and continent, otherwise error.
func (l *Location) GetByLatLng(ctx context.Context, lat, lng float64) (string, string, error) {
	cou, con, err := l.orbProvider.Search(ctx, lat, lng)
	if err != nil {
		return strings.Empty, strings.Empty, fmt.Errorf("%f/%f: %w", lat, lng, err)
	}

	return cou, continent.Codes[con], nil
}
