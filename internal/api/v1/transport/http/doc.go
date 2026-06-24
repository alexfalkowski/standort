// Package http provides the Standort API v1 HTTP transport wiring.
//
// The v1 API is exposed over HTTP by routing RPC-style HTTP requests to
// transport-specific handler functions.
//
// # Routing
//
// `Register` maps the generated full method names to the corresponding HTTP
// handlers:
//
//   - `standort.v1.Service/GetLocationByIP` → `getLocationByIP`
//   - `standort.v1.Service/GetLocationByLatLng` → `getLocationByLatLng`
//
// The route shapes (paths, verbs, encoding) are defined by the go-service RPC
// router (`rpc.Route`). This package is responsible only for wiring the routes
// to `internal/api/v1/location.Locator`.
//
// # Dependency injection
//
// The v1 API module (`internal/api/v1.Module`) registers this package's `Register`
// function into the application's dependency injection graph.
package http
