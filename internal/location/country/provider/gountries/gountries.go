package gountries

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/pariz/gountries"
)

// NewProvider constructs a country lookup provider backed by `github.com/pariz/gountries`.
//
// The returned provider implements `internal/location/country/provider.Provider` and
// resolves an input country code into:
//   - an ISO-3166 alpha-2 country code (e.g. "US"), and
//   - a continent name (e.g. "North America").
//
// Note: This provider returns the *continent name*, not the two-letter continent code.
// The mapping to two-letter codes is performed by the domain `internal/location` service
// via `internal/location/continent.Codes`.
func NewProvider() *Provider {
	return &Provider{query: gountries.New()}
}

// Provider implements country/continent lookup using a `gountries.Query`.
type Provider struct {
	query *gountries.Query
}

// GetByCode resolves a country code and continent name for the given country code.
//
// It delegates to `gountries.Query.FindCountryByAlpha`, which accepts common ISO alpha
// representations (typically alpha-2 or alpha-3).
//
// Returns:
//   - country: ISO-3166 alpha-2 country code (e.g. "US")
//   - continent: continent name (e.g. "Europe")
//   - err: any lookup error from gountries
func (p *Provider) GetByCode(_ context.Context, code string) (country string, continent string, err error) {
	c, err := p.query.FindCountryByAlpha(code)

	return c.Alpha2, c.Continent, err
}
