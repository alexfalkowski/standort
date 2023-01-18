package cmd

import (
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/metrics"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/standort/config"
	"github.com/alexfalkowski/standort/location"
	"github.com/alexfalkowski/standort/server/health"
	v1 "github.com/alexfalkowski/standort/server/v1"
	v2 "github.com/alexfalkowski/standort/server/v2"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, marshaller.Module, Module, config.Module, health.Module,
	logger.ZapModule, metrics.PrometheusModule, transport.Module,
	location.Module, v1.Module, v2.Module,
}
