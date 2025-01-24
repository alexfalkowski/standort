package gountries

import (
	"context"
	"strings"

	"github.com/alexfalkowski/standort/location/errors"
	"github.com/pariz/gountries"
)

// NewProvider for gountries.
func NewProvider() *Provider {
	return &Provider{query: gountries.New()}
}

// Provider for gountries.
type Provider struct {
	query *gountries.Query
}

// GetByCode a country and continent.
func (p *Provider) GetByCode(_ context.Context, name string) (string, string, error) {
	country, err := p.query.FindCountryByAlpha(name)
	if err != nil {
		return "", "", p.error(err)
	}

	return country.Codes.Alpha2, country.Continent, nil
}

func (p *Provider) error(err error) error {
	if strings.Contains(err.Error(), "find") {
		return errors.ErrNotFound
	}

	return err
}
