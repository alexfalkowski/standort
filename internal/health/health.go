package health

import (
	"github.com/alexfalkowski/go-health/v2/checker"
	"github.com/alexfalkowski/go-health/v2/server"
	"github.com/alexfalkowski/go-service/v2/env"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/health"
	v1 "github.com/alexfalkowski/standort/v2/api/standort/v1"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
)

func register(name env.Name, srv *server.Server, cfg *Config) {
	regs := health.Registrations{
		server.NewRegistration("noop", cfg.Duration.Duration(), checker.NewNoopChecker()),
		server.NewOnlineRegistration(cfg.Timeout.Duration(), cfg.Duration.Duration()),
	}

	srv.Register(name.String(), regs...)
	srv.Register(v1.Service_ServiceDesc.ServiceName, regs[0])
	srv.Register(v2.Service_ServiceDesc.ServiceName, regs[0])
}

func httpHealthObserver(name env.Name, server *server.Server) error {
	return server.Observe(name.String(), "healthz", "online")
}

func httpLivenessObserver(name env.Name, server *server.Server) error {
	return server.Observe(name.String(), "livez", "noop")
}

func httpReadinessObserver(name env.Name, server *server.Server) error {
	return server.Observe(name.String(), "readyz", "noop")
}

func grpcObserver(server *server.Server) error {
	return errors.Join(
		server.Observe(v1.Service_ServiceDesc.ServiceName, "grpc", "noop"),
		server.Observe(v2.Service_ServiceDesc.ServiceName, "grpc", "noop"),
	)
}
