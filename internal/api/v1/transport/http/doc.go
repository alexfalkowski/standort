// Package http provides the Standort API v1 HTTP transport wiring.
//
// The v1 API is exposed over HTTP by routing RPC-style HTTP requests to a
// concrete transport server.
//
// # Routing
//
// `Register` maps the generated full method names to the corresponding HTTP
// server handlers:
//
//   - `standort.v1.Service/GetLocationByIP` → `Server.GetLocationByIP`
//   - `standort.v1.Service/GetLocationByLatLng` → `Server.GetLocationByLatLng`
//
// The route shapes (paths, verbs, encoding) are defined by the go-service RPC
// router (`rpc.Route`). This package is responsible only for wiring the routes
// to a `Server`.
//
// # Dependency injection
//
// The v1 API module (`internal/api/v1.Module`) registers this package's
// `NewServer` constructor and `Register` function into the application's
// dependency injection graph.
package http
