// Package assets provides access to standort's embedded runtime data files.
//
// The service embeds a small set of assets at build time and exposes them as an
// `embed.FS` so other components can load them using standard `io/fs` APIs.
//
// Embedded files
//
//   - earth.geojson: used to build the spatial index for point-in-polygon
//     lookups (latitude/longitude → country/continent).
//   - geoip2.mmdb: used to resolve IP addresses to ISO country codes.
//
// # Dependency injection
//
// The package also exports `Module`, which registers `NewFS` into the
// application's dependency injection graph so providers can depend on the
// embedded filesystem.
package assets
