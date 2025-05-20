package cmd

import (
	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/go-service/v2/debug"
	"github.com/alexfalkowski/go-service/v2/feature"
	"github.com/alexfalkowski/go-service/v2/module"
	"github.com/alexfalkowski/go-service/v2/telemetry"
	"github.com/alexfalkowski/go-service/v2/transport"
	"github.com/alexfalkowski/standort/assets"
	v1 "github.com/alexfalkowski/standort/internal/api/v1"
	v2 "github.com/alexfalkowski/standort/internal/api/v2"
	"github.com/alexfalkowski/standort/internal/config"
	"github.com/alexfalkowski/standort/internal/health"
	"github.com/alexfalkowski/standort/internal/location"
)

// RegisterServer for cmd.
func RegisterServer(command cli.Commander) {
	cmd := command.AddServer("server", "Start standort server",
		module.Module, debug.Module, feature.Module,
		telemetry.Module, transport.Module,
		health.Module, location.Module,
		assets.Module, config.Module,
		v1.Module, v2.Module, cli.Module,
	)
	cmd.AddInput("")
}
