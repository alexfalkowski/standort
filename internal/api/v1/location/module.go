package location

import "github.com/alexfalkowski/go-service/v2/di"

// Module wires the v1 response locator into the application's dependency injection graph.
var Module = di.Module(
	di.Constructor(NewLocator),
)
