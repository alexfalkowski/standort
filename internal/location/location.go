package location

import (
	"context"
	"fmt"

	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/standort/internal/location/continent"
	country "github.com/alexfalkowski/standort/internal/location/country/provider"
	ip "github.com/alexfalkowski/standort/internal/location/ip/provider"
	orb "github.com/alexfalkowski/standort/internal/location/orb/provider"
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
//
//nolint:nonamedreturns
func (l *Location) GetByIP(ctx context.Context, ip string) (country string, cont string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s: %w", ip, runtime.ConvertRecover(r))
		}
	}()

	c, err := l.ipProvider.GetByIP(ctx, ip)
	runtime.Must(err)

	country, cont, err = l.countryProvider.GetByCode(ctx, c)
	runtime.Must(err)

	cont = continent.Codes[cont]

	return
}

// GetByLatLng a country and continent, otherwise error.
func (l *Location) GetByLatLng(ctx context.Context, lat, lng float64) (string, string, error) {
	cou, con, err := l.orbProvider.Search(ctx, lat, lng)
	if err != nil {
		return "", "", fmt.Errorf("%f/%f: %w", lat, lng, err)
	}

	return cou, continent.Codes[con], nil
}
