package gountries

import (
	"context"

	"github.com/pariz/gountries"
)

// Provider for gountries.
type Provider struct {
	query *gountries.Query
}

// NewProvider for gountries.
func NewProvider() *Provider {
	return &Provider{query: gountries.New()}
}

// GetByCode a country and continent.
func (p *Provider) GetByCode(_ context.Context, name string) (string, string, error) {
	country, err := p.query.FindCountryByAlpha(name)
	if err != nil {
		return "", "", err
	}

	return country.Codes.Alpha2, country.Continent, nil
}
