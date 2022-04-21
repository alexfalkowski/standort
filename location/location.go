package location

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/alexfalkowski/standort/location/continent"
	"github.com/ip2location/ip2location-go/v9"
	"github.com/pariz/gountries"
)

var (
	// ErrInvalidIP for location.
	ErrInvalidIP = errors.New("invalid ip address")

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
func (l *Location) GetByIP(ctx context.Context, ip string) (string, string, error) {
	if net.ParseIP(ip) == nil {
		return "", "", fmt.Errorf("%s: %w", ip, ErrInvalidIP)
	}

	rec, err := l.db.Get_all(ip)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", ip, ErrNotFound)
	}

	country, err := l.query.FindCountryByName(rec.Country_long)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", ip, ErrNotFound)
	}

	return country.Codes.Alpha2, continent.Codes[country.Continent], nil
}
