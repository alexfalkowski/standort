// Package v2 wires the Standort API v2 transport layer into the application.
//
// The v2 API is exposed over both gRPC and HTTP. The HTTP transport is built by
// routing HTTP requests to the corresponding gRPC handler function using
// go-service's RPC routing.
//
// Compared to v1, v2 introduces a transport-facing lookup service
// (`internal/api/location.Locator`) that can resolve a location from multiple
// inputs (IP address and/or geographic point), including metadata fallbacks.
//
// # Dependency injection
//
// This package exports `Module`, which registers:
//   - the transport-facing location adapter (`location.NewLocator`),
//   - the v2 gRPC server constructor and gRPC registration, and
//   - the v2 HTTP routing registration that maps HTTP routes to the gRPC handler.
//
// The concrete transport implementations live under `internal/api/v2/transport/*`.
package v2
