package config

import (
	"github.com/alexfalkowski/go-service/v2/config"
	"github.com/alexfalkowski/go-service/v2/di"
)

// Module wires standort's configuration into the application's dependency injection graph.
//
// It registers:
//   - `config.NewConfig[Config]` to load the top-level service configuration,
//   - `decorateConfig` to expose the embedded `*config.Config` to consumers, and
//   - `healthConfig` to provide `*health.Config` to the health module.
var Module = di.Module(
	di.Constructor(config.NewConfig[Config]),
	di.Decorate(decorateConfig),
	di.Constructor(healthConfig),
)
