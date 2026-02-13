// Package v1 wires the Standort API v1 transport layer into the application.
//
// The v1 API is exposed over both gRPC and HTTP. The HTTP transport is built by
// routing HTTP requests to the corresponding gRPC handler functions using
// go-service's RPC routing.
//
// # Dependency injection
//
// This package exports `Module`, which registers:
//   - the v1 gRPC server constructor, and
//   - gRPC/HTTP registration functions that attach the service handlers to the
//     application's gRPC registrar and HTTP router.
//
// The concrete transport implementations live under `internal/api/v1/transport/*`.
package v1
