package cmd

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/go-service/v2/module"
	"github.com/alexfalkowski/standort/v2/assets"
	v1 "github.com/alexfalkowski/standort/v2/internal/api/v1"
	v2 "github.com/alexfalkowski/standort/v2/internal/api/v2"
	"github.com/alexfalkowski/standort/v2/internal/config"
	"github.com/alexfalkowski/standort/v2/internal/health"
	"github.com/alexfalkowski/standort/v2/internal/location"
)

// Module is the top-level dependency injection composition for the standort server.
//
// It composes framework/server wiring (`module.Server`) with standort's internal modules:
//   - `config.Module`: configuration loading/decoration for the service
//   - `health.Module`: health checks and HTTP/GRPC health observers
//   - `assets.Module`: embedded runtime assets (GeoJSON, GeoIP database)
//   - `location.Module`: domain location resolution (IP and point-in-polygon)
//   - `v1.Module`: API v1 transports (gRPC + HTTP route registration)
//   - `v2.Module`: API v2 transports (gRPC + HTTP route registration)
var Module = di.Module(
	module.Server,
	config.Module,
	health.Module,
	assets.Module,
	location.Module,
	v1.Module,
	v2.Module,
)
