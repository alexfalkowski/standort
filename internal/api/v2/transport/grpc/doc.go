// Package grpc provides the Standort API v2 gRPC transport implementation.
//
// This package contains the concrete server that implements the generated
// `standort.v2.ServiceServer` interface.
//
// # Registration
//
// Use `Register` to register the server with a gRPC service registrar:
//
//   - `Register(registrar, server)` delegates to the generated
//     `v2.RegisterServiceServer` function.
//
// # Server construction
//
// Use `NewServer` to construct a `*Server` instance. The v2 server delegates
// request handling to the transport-facing `internal/api/location.Locator`
// service, which is responsible for applying transport-level behavior such as:
//
//   - reading fallback inputs from request metadata when request fields are empty
//   - attaching partial lookup failures as metadata attributes
//
// # Error mapping
//
// The v2 gRPC transport maps any non-nil service error to a gRPC
// `codes.NotFound` status (see `(*Server).error`). This provides a consistent
// client-facing error contract for "no location found" (and other lookup
// failures).
//
// # Response metadata
//
// Handlers populate `resp.Meta` from request metadata using
// `meta.CamelStrings(ctx, strings.Empty)`. This propagates transport metadata to
// clients in a stable, camel-cased key format.
package grpc
