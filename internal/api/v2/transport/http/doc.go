// Package http provides the Standort API v2 HTTP transport wiring.
//
// The v2 API is exposed over HTTP by routing RPC-style HTTP requests to
// transport-specific handler functions.
//
// # Routing
//
// Register maps the generated full method name to the corresponding HTTP handler:
//
//   - `standort.v2.Service/GetLocation` → `getLocation`
//
// The route shapes (paths, verbs, encoding) are defined by the go-service RPC
// router (`rpc.Route`). This package is responsible only for wiring the route to
// `internal/api/v2/location.Locator` and setting HTTP-specific diagnostics on
// terminal lookup failures.
//
// # Dependency injection
//
// The v2 API module (`internal/api/v2.Module`) registers this package's `Register`
// function into the application's dependency injection graph.
package http
