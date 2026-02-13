package location

import "github.com/alexfalkowski/go-service/v2/di"

// Module wires the transport-facing `location` package into the application's
// dependency injection graph.
//
// It registers `NewLocator`, which adapts the domain `internal/location.Location`
// service into a transport-friendly `Locator` that can:
//   - read fallback inputs from request metadata (IP address and geolocation), and
//   - record lookup/parsing failures into metadata attributes.
var Module = di.Module(
	di.Constructor(NewLocator),
)
