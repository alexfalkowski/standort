package location

import (
	"github.com/alexfalkowski/standort/v2/internal/location/country"
	"github.com/alexfalkowski/standort/v2/internal/location/ip"
	"github.com/alexfalkowski/standort/v2/internal/location/orb"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(ip.NewProvider),
	fx.Provide(country.NewProvider),
	fx.Provide(New),
	fx.Provide(orb.NewProvider),
)
