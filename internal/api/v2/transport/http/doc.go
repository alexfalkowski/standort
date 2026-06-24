// Package http provides the Standort API v2 HTTP transport wiring.
//
// The v2 API is exposed over HTTP by routing RPC-style HTTP requests to a
// concrete transport server.
//
// # Routing
//
// Register maps the generated full method name to the corresponding HTTP server
// handler:
//
//   - `standort.v2.Service/GetLocation` → `Server.GetLocation`
//   - `standort.v2.Service/LookupLocations` → `Server.LookupLocations`
//
// The route shapes (paths, verbs, encoding) are defined by the go-service RPC
// router (`rpc.Route`). This package is responsible only for wiring the route to
// a `Server` and setting HTTP-specific diagnostics on terminal single-lookup
// failures.
//
// # Dependency injection
//
// The v2 API module (`internal/api/v2.Module`) registers this package's
// `NewServer` constructor and `Register` function into the application's
// dependency injection graph.
package http
