package cmd

import (
	"github.com/alexfalkowski/go-service/compressor"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/standort/assets"
	"github.com/alexfalkowski/standort/config"
	"github.com/alexfalkowski/standort/location"
	"github.com/alexfalkowski/standort/server/health"
	v1 "github.com/alexfalkowski/standort/server/v1"
	v2 "github.com/alexfalkowski/standort/server/v2"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	runtime.Module, debug.Module, feature.Module,
	compressor.Module, marshaller.Module,
	telemetry.Module, metrics.Module,
	transport.Module, health.Module, location.Module,
	assets.Module, config.Module, v1.Module, v2.Module, Module,
}
