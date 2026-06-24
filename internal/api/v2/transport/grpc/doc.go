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
//   - mapping oversized batch requests to gRPC `codes.InvalidArgument`
//   - attaching terminal lookup diagnostics as gRPC trailers
//
// # Error mapping
//
// Single lookup failures are mapped to a gRPC `codes.NotFound` status. Batch
// lookup entry failures are returned as per-entry `google.rpc.Status` values,
// while request-level validation failures such as oversized batches are mapped
// to `codes.InvalidArgument`.
//
// # Response metadata
//
// Handlers populate `resp.Meta` from request metadata using
// `meta.CamelStrings(ctx, strings.Empty)`. This propagates transport metadata to
// clients in a stable, camel-cased key format. Single lookup diagnostics are not
// written into response bodies; terminal lookup failures attach code-only
// diagnostics as gRPC trailers.
package grpc
