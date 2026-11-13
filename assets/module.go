package assets

import "github.com/alexfalkowski/go-signal/v2"

// Module wires the `assets` package into the application's dependency injection graph.
//
// It registers `NewFS` so other constructors can depend on the embedded asset
// filesystem (for example to read `earth.geojson` or `geoip2.mmdb`).
var Module = signal.Module(
	signal.Constructor(NewFS),
)
