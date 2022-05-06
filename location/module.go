package location

import (
	"github.com/alexfalkowski/standort/location/country"
	cotr "github.com/alexfalkowski/standort/location/country/provider/trace/opentracing"
	"github.com/alexfalkowski/standort/location/ip"
	iotr "github.com/alexfalkowski/standort/location/ip/provider/trace/opentracing"
	"github.com/alexfalkowski/standort/location/orb"
	ootr "github.com/alexfalkowski/standort/location/orb/provider/trace/opentracing"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(ip.NewProvider),
		fx.Provide(country.NewProvider),
		fx.Provide(New),
		fx.Provide(cotr.NewTracer),
		fx.Provide(orb.NewProvider),
		fx.Provide(ootr.NewTracer),
		fx.Provide(iotr.NewTracer),
	)
)
