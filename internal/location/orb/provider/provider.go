package provider

import "github.com/alexfalkowski/go-service/v2/context"

// Provider resolves a latitude/longitude coordinate to location codes.
//
// Implementations are expected to perform a geospatial search (for example,
// point-in-polygon) and return:
//   - an ISO-3166 alpha-2 country code (e.g. "US")
//   - a continent name (e.g. "North America")
//
// The domain `internal/location` service maps the returned continent name to a
// two-letter continent code via `internal/location/continent.Codes`.
type Provider interface {
	// Search resolves the country code and continent name for a latitude/longitude coordinate.
	//
	// Inputs are expressed in degrees.
	//
	// Returns:
	//   - countryCode: ISO-3166 alpha-2 country code (e.g. "US")
	//   - continent: continent name (e.g. "Europe")
	//   - err: any lookup/validation error (implementations may return a sentinel
	//     not-found error when no geometry contains the point)
	Search(ctx context.Context, lat, lng float64) (countryCode string, continent string, err error)
}
