// Package grpc provides the Standort API v1 gRPC transport implementation.
//
// This package contains the concrete server that implements the generated
// `standort.v1.ServiceServer` interface.
//
// # Registration
//
// Use `Register` to register the server with a gRPC service registrar:
//
//   - `Register(registrar, server)` delegates to the generated
//     `v1.RegisterServiceServer` function.
//
// # Server construction
//
// Use `NewServer` to construct a `*Server` instance. The server delegates all
// lookups to the domain `internal/location.Location` service.
//
// # Error mapping
//
// The v1 gRPC transport maps any non-nil domain error to a gRPC
// `codes.NotFound` status (see `(*Server).error`). This ensures a consistent
// client-facing error contract for "no location found" (and other lookup
// failures), while responses may still include metadata.
//
// # Response metadata
//
// Handlers populate `resp.Meta` from request metadata using
// `meta.CamelStrings(ctx, strings.Empty)`. This propagates transport metadata to
// clients in a stable, camel-cased key format.
package grpc
