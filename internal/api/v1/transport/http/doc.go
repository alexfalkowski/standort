// Package http provides the Standort API v1 HTTP transport wiring.
//
// The v1 API is implemented as gRPC handlers (see `internal/api/v1/transport/grpc`)
// and exposed over HTTP by routing HTTP requests to those gRPC handler functions
// using go-service's HTTP↔gRPC RPC routing.
//
// # Routing
//
// `Register` maps the generated gRPC full method names to the corresponding
// gRPC handler methods:
//
//   - `standort.v1.Service/GetLocationByIP` → `(*grpc.Server).GetLocationByIP`
//   - `standort.v1.Service/GetLocationByLatLng` → `(*grpc.Server).GetLocationByLatLng`
//
// The route shapes (paths, verbs, encoding) are defined by the go-service RPC
// router (`rpc.Route`). This package is responsible only for wiring the routes
// to the existing gRPC implementation.
//
// # Dependency injection
//
// The v1 API module (`internal/api/v1.Module`) registers this package's `Register`
// function into the application's dependency injection graph.
package http
