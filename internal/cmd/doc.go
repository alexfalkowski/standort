// Package cmd contains the standort CLI command wiring.
//
// This package composes the application's top-level dependency injection graph
// (see `Module`) and registers CLI commands that are exposed by the binary.
//
// # Server command
//
// `RegisterServer` registers the `server` command, which starts the standort
// service. The command uses go-service's server framework integration and the
// module graph rooted at `Module`, which typically includes:
//
//   - framework server wiring (`module.Server`)
//   - configuration (`internal/config`)
//   - health checks and observers (`internal/health`)
//   - embedded runtime assets (`assets`)
//   - the domain location service (`internal/location`)
//   - API transports for v1 and v2 (`internal/api/v1`, `internal/api/v2`)
//
// The command intentionally adds an empty input, deferring configuration loading
// to the defaults supported by the underlying framework (flags/env/config files
// as configured by go-service).
package cmd
