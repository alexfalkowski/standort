// Package location contains standort's domain logic for resolving a location
// (country and continent) from different kinds of inputs.
//
// The domain service exported by this package is `*Location`. It composes three
// provider interfaces that each handle a different concern:
//
//   - IP provider: resolves an IP address to an ISO-3166 alpha-2 country code.
//   - Country provider: resolves a country code to a country code and continent name.
//   - Orb (geospatial) provider: resolves latitude/longitude to a country code and
//     continent name (typically by point-in-polygon search).
//
// The domain service normalizes provider outputs into stable codes:
//
//   - Country is returned as an ISO-3166 alpha-2 code (for example "US").
//   - Continent is returned as a two-letter continent code (for example "NA"),
//     using the mapping in `internal/location/continent.Codes`.
//
// # APIs
//
// The main entry points are:
//
//   - `New`: constructs a `*Location` by composing provider implementations.
//   - `(*Location).GetByIP`: resolves (country, continent) from an IP address.
//   - `(*Location).GetByLatLng`: resolves (country, continent) from a lat/lng point.
//
// # Errors and not-found behavior
//
// Provider implementations may return sentinel "not found" errors (for example
// the R-tree provider returns `internal/location/orb/provider/rtree.ErrNotFound`
// when no polygon contains the point). This package does not interpret those
// sentinel errors; it returns them to callers, wrapping some errors with input
// context where appropriate (e.g. lat/lng formatting).
//
// Transport layers may choose to map these errors into transport-specific status
// codes (for example, gRPC `codes.NotFound`).
//
// # Dependency injection
//
// For applications using go-service DI, this package also exports `Module`,
// which registers the default provider constructors and the domain service
// constructor (`New`) into the application's dependency injection graph.
package location
