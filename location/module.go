package location

import (
	"github.com/alexfalkowski/standort/location/ip"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(ip.NewDB),
	)
)
