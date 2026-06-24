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
// response construction to `internal/api/v2/location.Locator`, then applies
// gRPC-specific behavior such as:
//
//   - mapping terminal lookup errors to gRPC `codes.NotFound`
//   - attaching terminal lookup diagnostics as gRPC trailers
//
// # Error mapping
//
// The v2 gRPC transport maps any non-nil service error to a gRPC
// `codes.NotFound` status. This provides a consistent client-facing error
// contract for "no location found" (and other lookup failures).
//
// # Response metadata
//
// Handlers populate `resp.Meta` from request metadata using
// `meta.CamelStrings(ctx, strings.Empty)`. This propagates transport metadata to
// clients in a stable, camel-cased key format. Lookup diagnostics are not written
// into response bodies; terminal lookup failures attach code-only diagnostics as
// gRPC trailers.
package grpc
