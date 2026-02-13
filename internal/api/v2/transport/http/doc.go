// Package http provides the Standort API v2 HTTP transport wiring.
//
// The v2 API is implemented as gRPC handlers (see `internal/api/v2/transport/grpc`)
// and exposed over HTTP by routing HTTP requests to those gRPC handler functions
// using go-service's HTTP↔gRPC RPC routing.
//
// # Routing
//
// Register maps the generated gRPC full method name to the corresponding gRPC
// handler method:
//
//   - `standort.v2.Service/GetLocation` → `(*grpc.Server).GetLocation`
//
// The route shapes (paths, verbs, encoding) are defined by the go-service RPC
// router (`rpc.Route`). This package is responsible only for wiring the route to
// the existing gRPC implementation.
//
// # Dependency injection
//
// The v2 API module (`internal/api/v2.Module`) registers this package's `Register`
// function into the application's dependency injection graph.
package http
