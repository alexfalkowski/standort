package cmd

import (
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/metrics"
	"github.com/alexfalkowski/go-service/trace"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/standort/config"
	"github.com/alexfalkowski/standort/location"
	"github.com/alexfalkowski/standort/location/ip"
	"github.com/alexfalkowski/standort/server/health"
	v1 "github.com/alexfalkowski/standort/server/v1"
	v2 "github.com/alexfalkowski/standort/server/v2"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, fx.Provide(NewVersion), config.Module, health.Module,
	logger.ZapModule, metrics.PrometheusModule,
	transport.GRPCServerModule, transport.GRPCJaegerModule,
	transport.HTTPServerModule, transport.HTTPJaegerModule,
	trace.JaegerOpenTracingModule, ip.ProviderJaegerModule,
	location.Module, v1.Module, v2.Module,
}
