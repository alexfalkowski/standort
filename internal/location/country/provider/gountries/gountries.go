package gountries

import (
	"context"

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

	return country.Alpha2, country.Continent, err
}
