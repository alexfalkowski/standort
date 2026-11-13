package location

import "github.com/alexfalkowski/go-signal/v2"

// Module wires the v1 response locator into the application's dependency injection graph.
var Module = signal.Module(
	signal.Constructor(NewLocator),
)
