package provider

import "github.com/alexfalkowski/go-service/v2/context"

// Provider resolves country and continent information for a given ISO country code.
//
// Implementations typically take an ISO-3166 alpha-2 (or alpha-3, depending on the
// implementation) country code and return:
//   - a normalized country code, and
//   - a continent name (for example, "Europe" or "North America").
//
// The domain `internal/location` service maps the returned continent name into a
// two-letter continent code via `internal/location/continent.Codes`.
type Provider interface {
	// GetByCode resolves a country code and continent name for the given ISO country code.
	//
	// Implementations should treat `code` as case-insensitive where possible.
	//
	// Returns:
	//   - country: a (typically alpha-2) country code (e.g. "US")
	//   - continent: a continent name (e.g. "North America")
	//   - err: any lookup/validation error
	GetByCode(ctx context.Context, code string) (country string, continent string, err error)
}
