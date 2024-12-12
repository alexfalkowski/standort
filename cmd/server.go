package cmd

import (
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/sync"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
	v1 "github.com/alexfalkowski/standort/api/v1"
	v2 "github.com/alexfalkowski/standort/api/v2"
	"github.com/alexfalkowski/standort/assets"
	"github.com/alexfalkowski/standort/config"
	"github.com/alexfalkowski/standort/health"
	"github.com/alexfalkowski/standort/location"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	sync.Module, compress.Module, encoding.Module,
	runtime.Module, debug.Module, feature.Module,
	telemetry.Module, transport.Module,
	health.Module, location.Module,
	assets.Module, config.Module,
	v1.Module, v2.Module, Module,
}
