package cmd

import (
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/metrics"
	"github.com/alexfalkowski/go-service/trace"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/standort/config"
	"github.com/alexfalkowski/standort/server/health"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, config.Module, health.Module,
	logger.ZapModule, metrics.PrometheusModule,
	transport.HTTPServerModule, transport.GRPCServerModule,
	trace.JaegerOpenTracingModule,
}
