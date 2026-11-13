package health

import "github.com/alexfalkowski/go-signal/v2"

// Module wires the health subsystem into the application's dependency injection graph.
//
// It registers health checks and exposes health endpoints/observers for both HTTP and gRPC:
//   - Service-level health registrations, including a noop checker and an online check.
//   - HTTP observers for `healthz`, `livez`, and `readyz` endpoints.
//   - gRPC observers for the v1 and v2 Standort service descriptors.
//
// This module is intended to be composed into the top-level server module (see `internal/cmd.Module`).
var Module = signal.Module(
	signal.Register(register),
	signal.Register(httpHealthObserver),
	signal.Register(httpLivenessObserver),
	signal.Register(httpReadinessObserver),
	signal.Register(grpcObserver),
)
