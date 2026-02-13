package health

import "github.com/alexfalkowski/go-service/v2/di"

// Module wires the health subsystem into the application's dependency injection graph.
//
// It registers health checks and exposes health endpoints/observers for both HTTP and gRPC:
//   - Service-level health registrations, including a noop checker and an online check.
//   - HTTP observers for `healthz`, `livez`, and `readyz` endpoints.
//   - gRPC observers for the v1 and v2 Standort service descriptors.
//
// This module is intended to be composed into the top-level server module (see `internal/cmd.Module`).
var Module = di.Module(
	di.Register(register),
	di.Register(httpHealthObserver),
	di.Register(httpLivenessObserver),
	di.Register(httpReadinessObserver),
	di.Register(grpcObserver),
)
