package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/standort/assets"
	v1 "github.com/alexfalkowski/standort/internal/api/v1"
	v2 "github.com/alexfalkowski/standort/internal/api/v2"
	"github.com/alexfalkowski/standort/internal/config"
	"github.com/alexfalkowski/standort/internal/health"
	"github.com/alexfalkowski/standort/internal/location"
)

// RegisterServer for cmd.
func RegisterServer(command *cmd.Command) {
	flags := command.AddServer("server", "Start standort server",
		module.Module, debug.Module, feature.Module,
		telemetry.Module, transport.Module,
		health.Module, location.Module,
		assets.Module, config.Module,
		v1.Module, v2.Module, cmd.Module,
	)
	flags.AddInput("")
}
