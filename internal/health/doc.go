// Package health wires standort's health subsystem into the application.
//
// standort uses go-health and go-service health integrations to expose both
// HTTP and gRPC health signals.
//
// # What this package does
//
//   - Registers health checks against the service name as well as the v1 and v2
//     gRPC service descriptors.
//   - Exposes HTTP observers for common Kubernetes-style endpoints:
//     `healthz`, `livez`, and `readyz`.
//   - Exposes gRPC observers for the Standort gRPC services.
//
// The concrete configuration for check intervals and timeouts is defined by
// `Config` (see `config.go`) and is expected to be provided by the application's
// configuration system.
//
// # Health semantics
//
// The HTTP `healthz` observer intentionally uses go-health's `online` checker.
// That checker verifies public internet reachability using its default URL set
// when no custom URLs are supplied. This is a product/operational signal for
// standort, not an accidental external dependency: `healthz` is expected to
// report unhealthy when the process cannot reach the public internet.
//
// Liveness and readiness are separate local signals. The HTTP `livez` and
// `readyz` observers, as well as the Standort gRPC service observers, use the
// local `noop` checker and do not require public egress.
//
// # Dependency injection
//
// This package exports `Module`, which registers the check registration and
// observer functions into the application's dependency injection graph.
// Typically, `internal/cmd.Module` composes this module into the overall server.
package health
