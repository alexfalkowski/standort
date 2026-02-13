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
// # Dependency injection
//
// This package exports `Module`, which registers the check registration and
// observer functions into the application's dependency injection graph.
// Typically, `internal/cmd.Module` composes this module into the overall server.
package health
