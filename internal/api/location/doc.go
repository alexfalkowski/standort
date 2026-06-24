// Package location provides transport-facing location lookup helpers for the API layers.
//
// This package adapts the domain `internal/location` service to transport needs.
// It is used by API v2 directly (and is suitable for any transport) to:
//
//   - accept explicit lookup inputs (IP address and/or latitude/longitude),
//   - fall back to request metadata when inputs are omitted, and
//   - collect code-only terminal failure diagnostics for transports.
//
// # Metadata fallbacks
//
// When an input is not provided by the caller, `(*Locator).Locate` attempts to
// derive it from metadata:
//
//   - IP address: `meta.IPAddr(ctx).Value()`
//   - Geolocation: `meta.Geolocation(ctx)` parsed as a geo URI (RFC 5870-style)
//
// This package treats metadata as an already-established transport context. It
// does not decide whether forwarded IP headers or other metadata sources are
// trustworthy; that boundary belongs to the transport/framework/deployment layer
// that populates `meta.IPAddr`.
//
// # Failure reporting
//
// Lookup/parsing failures do not immediately fail the request. The locator keeps
// trying any other available inputs. If at least one lookup succeeds, the response
// is successful and failed-side diagnostics are discarded.
//
// When no lookup succeeds, `ErrNotFound` is returned with code-only diagnostics
// so v2 transports can surface why the terminal lookup failed without exposing
// provider error text. Diagnostic keys are owned by `internal/diagnostics`; this
// package only records failure categories such as IP lookup, geo URI parsing,
// and point lookup.
//
// v1 uses separate lookup methods and does not use this diagnostic transport
// contract.
//
// # Dependency injection
//
// The package exports `Module`, which registers `NewLocator` into the
// application's dependency injection graph.
package location
