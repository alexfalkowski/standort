package gountries

import (
	"context"
	"fmt"

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
func (p *Provider) GetByCode(_ context.Context, code string) (string, string, error) {
	country, err := p.query.FindCountryByAlpha(code)

	return country.Codes.Alpha2, country.Continent, p.error(code, err)
}

func (p *Provider) error(code string, err error) error {
	if err != nil {
		return fmt.Errorf("%v: %w", code, errors.ErrNotFound)
	}

	return nil
}
