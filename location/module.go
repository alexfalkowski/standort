package location

import (
	"github.com/alexfalkowski/standort/location/country"
	cjaeger "github.com/alexfalkowski/standort/location/country/provider/opentracing/jaeger"
	"github.com/alexfalkowski/standort/location/ip"
	ijaeger "github.com/alexfalkowski/standort/location/ip/provider/trace/opentracing/jaeger"
	"github.com/alexfalkowski/standort/location/orb"
	ojaeger "github.com/alexfalkowski/standort/location/orb/provider/opentracing/jaeger"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(ip.NewProvider),
		fx.Provide(country.NewProvider),
		fx.Provide(New),
		fx.Provide(ojaeger.NewTracer),
		fx.Provide(orb.NewProvider),
		fx.Provide(cjaeger.NewTracer),
		fx.Provide(ijaeger.NewTracer),
	)
)
