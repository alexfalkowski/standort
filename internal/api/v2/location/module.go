package location

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/standort/v2/internal/api/location"
)

// Module wires the v2 response locator into the application's dependency injection graph.
//
// It composes the lower-level transport-facing location module and registers the
// v2 `Locator`, which builds generated v2 responses from lookup results.
var Module = di.Module(
	location.Module,
	di.Constructor(NewLocator),
)
