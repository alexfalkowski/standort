package location

import (
	"github.com/alexfalkowski/standort/location/country"
	ct "github.com/alexfalkowski/standort/location/country/provider/telemetry/tracer"
	"github.com/alexfalkowski/standort/location/ip"
	it "github.com/alexfalkowski/standort/location/ip/provider/telemetry/tracer"
	"github.com/alexfalkowski/standort/location/orb"
	ot "github.com/alexfalkowski/standort/location/orb/provider/telemetry/tracer"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(ip.NewProvider),
	fx.Provide(country.NewProvider),
	fx.Provide(New),
	fx.Provide(ct.NewTracer),
	fx.Provide(orb.NewProvider),
	fx.Provide(it.NewTracer),
	fx.Provide(ot.NewTracer),
)
