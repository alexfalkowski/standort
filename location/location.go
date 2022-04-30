package location

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/standort/location/continent"
	"github.com/alexfalkowski/standort/location/ip/provider"
	"github.com/alexfalkowski/standort/location/orb"
	"github.com/pariz/gountries"
	"github.com/tidwall/rtree"
)

var (
	// ErrInvalid for location.
	ErrInvalid = errors.New("invalid")

	// ErrNotFound for location.
	ErrNotFound = errors.New("not found")
)

// Location will find the country and continent by different criteria.
type Location struct {
	provider provider.Provider
	query    *gountries.Query
	tree     *rtree.Generic[*orb.Node]
}

// New location.
func New(provider provider.Provider, query *gountries.Query, tree *rtree.Generic[*orb.Node]) *Location {
	return &Location{provider: provider, query: query, tree: tree}
}

// GetByIP a country and continent, otherwise error.
func (l *Location) GetByIP(ctx context.Context, ipa string) (string, string, error) {
	c, err := l.provider.GetByIP(ctx, ipa)
	if err != nil {
		meta.WithAttribute(ctx, "ip.error", err.Error())

		return "", "", fmt.Errorf("%s: %w", ipa, ErrNotFound)
	}

	country, err := l.query.FindCountryByName(c)
	if err != nil {
		meta.WithAttribute(ctx, "ip.error", err.Error())

		return "", "", fmt.Errorf("%s: %w", ipa, ErrNotFound)
	}

	return country.Codes.Alpha2, continent.Codes[country.Continent], nil
}

// GetByLatLng a country and continent, otherwise error.
func (l *Location) GetByLatLng(ctx context.Context, lat, lng float64) (string, string, error) {
	data := orb.SearchTree(l.tree, lat, lng)
	if data == nil {
		return "", "", fmt.Errorf("%f/%f: %w", lat, lng, ErrNotFound)
	}

	return data.Country, continent.Codes[data.Continent], nil
}
