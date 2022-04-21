package location

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/standort/location/continent"
	"github.com/alexfalkowski/standort/location/ip"
	"github.com/ip2location/ip2location-go/v9"
	"github.com/pariz/gountries"
)

var (

	// ErrNotFound for location.
	ErrNotFound = errors.New("not found")
)

// Location will find the country and continent by different criteria.
type Location struct {
	db    *ip2location.DB
	query *gountries.Query
}

// New location.
func New(db *ip2location.DB, query *gountries.Query) *Location {
	return &Location{db: db, query: query}
}

// GetByIP a country and continent, otherwise error.
func (l *Location) GetByIP(ctx context.Context, ipa string) (string, string, error) {
	if err := ip.IsValid(ipa); err != nil {
		return "", "", fmt.Errorf("%s: %w", ipa, err)
	}

	rec, err := l.db.Get_all(ipa)
	if err != nil {
		meta.WithAttribute(ctx, "ip.error", err.Error())

		return "", "", fmt.Errorf("%s: %w", ipa, ErrNotFound)
	}

	country, err := l.query.FindCountryByName(rec.Country_long)
	if err != nil {
		meta.WithAttribute(ctx, "ip.error", err.Error())

		return "", "", fmt.Errorf("%s: %w", ipa, ErrNotFound)
	}

	return country.Codes.Alpha2, continent.Codes[country.Continent], nil
}
