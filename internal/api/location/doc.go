// Package location provides transport-facing location lookup helpers for the API layers.
//
// This package adapts the domain `internal/location` service to transport needs.
// It is used by API v2 directly (and is suitable for any transport) to:
//
//   - accept explicit lookup inputs (IP address and/or latitude/longitude),
//   - fall back to request metadata when inputs are omitted, and
//   - record partial failures into request metadata attributes.
//
// # Metadata fallbacks
//
// When an input is not provided by the caller, `(*Locator).Locate` attempts to
// derive it from metadata:
//
//   - IP address: `meta.IPAddr(ctx).Value()`
//   - Geolocation: `meta.Geolocation(ctx)` parsed as a geo URI (RFC 5870-style)
//
// # Partial failure reporting
//
// Lookup/parsing failures do not immediately fail the request. Instead, they are
// attached to the request context as metadata attributes so they can be surfaced
// to clients via the transport layer:
//
//   - `locationIpError` for IP lookup failures
//   - `locationPointError` for geo URI parsing failures
//   - `locationLatLngError` for point lookup failures
//
// `ErrNotFound` is returned only when neither the IP-derived nor geo-derived
// lookup yields a location.
//
// # Dependency injection
//
// The package exports `Module`, which registers `NewLocator` into the
// application's dependency injection graph.
package location
