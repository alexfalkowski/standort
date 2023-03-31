package location

import (
	"github.com/alexfalkowski/standort/location/country"
	cotel "github.com/alexfalkowski/standort/location/country/provider/otel"
	"github.com/alexfalkowski/standort/location/ip"
	iotel "github.com/alexfalkowski/standort/location/ip/provider/otel"
	"github.com/alexfalkowski/standort/location/orb"
	ootel "github.com/alexfalkowski/standort/location/orb/provider/otel"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(ip.NewProvider),
		fx.Provide(country.NewProvider),
		fx.Provide(New),
		fx.Provide(cotel.NewTracer),
		fx.Provide(orb.NewProvider),
		fx.Provide(iotel.NewTracer),
		fx.Provide(ootel.NewTracer),
	)
)
