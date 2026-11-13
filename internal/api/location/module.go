package location

import "github.com/alexfalkowski/go-signal/v2"

// Module wires the transport-facing `location` package into the application's
// dependency injection graph.
//
// It registers `NewLocator`, which adapts the domain `internal/location.Location`
// service into a transport-friendly `Locator` that can:
//   - read fallback inputs from request metadata (IP address and geolocation), and
//   - collect lookup/parsing diagnostics for terminal lookup failures.
var Module = signal.Module(
	signal.Constructor(NewLocator),
)
