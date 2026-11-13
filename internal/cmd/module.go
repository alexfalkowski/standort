package cmd

import (
	"github.com/alexfalkowski/go-service/v3/module"
	"github.com/alexfalkowski/go-signal/v2"
	"github.com/alexfalkowski/standort/v3/assets"
	v1 "github.com/alexfalkowski/standort/v3/internal/api/v1"
	v2 "github.com/alexfalkowski/standort/v3/internal/api/v2"
	"github.com/alexfalkowski/standort/v3/internal/config"
	"github.com/alexfalkowski/standort/v3/internal/health"
	"github.com/alexfalkowski/standort/v3/internal/location"
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
var Module = signal.Module(
	module.Server,
	config.Module,
	health.Module,
	assets.Module,
	location.Module,
	v1.Module,
	v2.Module,
)
