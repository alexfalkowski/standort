package location

import (
	"github.com/alexfalkowski/standort/location/country"
	"github.com/alexfalkowski/standort/location/ip"
	"github.com/alexfalkowski/standort/location/orb"
	"github.com/alexfalkowski/standort/location/orb/provider/opentracing/jaeger"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(ip.NewProvider),
		fx.Provide(country.NewQuery),
		fx.Provide(New),
		fx.Provide(jaeger.NewTracer),
		fx.Provide(orb.NewProvider),
	)
)
