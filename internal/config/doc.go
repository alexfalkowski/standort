// Package config defines the standort service configuration model and provides
// dependency-injection wiring for loading and exposing configuration.
//
// The standort configuration composes standort-specific settings with the shared
// go-service configuration type. This allows the application to use the go-service
// framework configuration while still adding service-specific sections.
//
// # Configuration structure
//
// The package-level `Config` type embeds `*config.Config` from go-service and adds
// additional fields (for example `Health`). The embedded config is inlined for
// YAML/JSON/TOML serialization so that the go-service settings appear at the top
// level of the service configuration file.
//
// # Dependency injection
//
// This package exports `Module`, which registers constructors/decoration to:
//
//   - load the top-level `*Config` (via `config.NewConfig[Config]`),
//   - expose the embedded `*config.Config` to components that depend on the base
//     framework config type, and
//   - provide extracted sub-configs (for example `*health.Config`) to other
//     modules.
//
// This package is intended to be composed into the application’s top-level module
// graph (see `internal/cmd.Module`).
package config
