// Package v2 wires the Standort API v2 transport layer into the application.
//
// The v2 API is exposed over both gRPC and HTTP. The HTTP transport is built by
// routing RPC-style HTTP requests to transport-specific handlers.
//
// Compared to v1, v2 introduces a v2 response locator
// (`internal/api/v2/location.Locator`) backed by the transport-facing lookup
// service (`internal/api/location.Locator`) and an embedded lookup asset provider
// for operator-facing asset metadata.
//
// # Dependency injection
//
// This package exports `Module`, which registers:
//   - the transport-facing location adapter and v2 response locator,
//   - the v2 embedded lookup asset provider,
//   - the v2 gRPC server constructor and gRPC registration, and
//   - the v2 HTTP routing registration.
//
// The concrete transport implementations live under `internal/api/v2/transport/*`.
package v2
