package location

import (
	"github.com/alexfalkowski/go-signal/v2"
	"github.com/alexfalkowski/standort/v3/internal/api/location"
)

// Module wires the v2 response locator into the application's dependency injection graph.
//
// It composes the lower-level transport-facing location module and registers the
// v2 `Locator`, which builds generated v2 responses from lookup results.
var Module = signal.Module(
	location.Module,
	signal.Constructor(NewLocator),
)
