package location

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/standort/v2/internal/location/country"
	"github.com/alexfalkowski/standort/v2/internal/location/ip"
	"github.com/alexfalkowski/standort/v2/internal/location/orb"
)

// Module for fx.
var Module = di.Module(
	di.Constructor(ip.NewProvider),
	di.Constructor(country.NewProvider),
	di.Constructor(New),
	di.Constructor(orb.NewProvider),
)
