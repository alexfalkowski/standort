package location

import (
	"github.com/alexfalkowski/standort/location/country"
	"github.com/alexfalkowski/standort/location/ip"
	"github.com/alexfalkowski/standort/location/orb"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(ip.NewDB),
		fx.Provide(country.NewQuery),
		fx.Provide(New),
		fx.Provide(orb.NewRTree),
	)
)
