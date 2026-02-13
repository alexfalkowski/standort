package assets

import "github.com/alexfalkowski/go-service/v2/di"

// Module wires the `assets` package into the application's dependency injection graph.
//
// It registers `NewFS` so other constructors can depend on the embedded asset
// filesystem (for example to read `earth.geojson` or `geoip2.mmdb`).
var Module = di.Module(
	di.Constructor(NewFS),
)
