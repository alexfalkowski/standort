package cmd

import (
	"github.com/alexfalkowski/go-service/v2/module"
	"github.com/alexfalkowski/standort/v2/assets"
	v1 "github.com/alexfalkowski/standort/v2/internal/api/v1"
	v2 "github.com/alexfalkowski/standort/v2/internal/api/v2"
	"github.com/alexfalkowski/standort/v2/internal/config"
	"github.com/alexfalkowski/standort/v2/internal/health"
	"github.com/alexfalkowski/standort/v2/internal/location"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	module.Server,
	config.Module,
	health.Module,
	assets.Module,
	location.Module,
	v1.Module,
	v2.Module,
)
