package location

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/standort/v2/internal/location/country"
	"github.com/alexfalkowski/standort/v2/internal/location/ip"
	"github.com/alexfalkowski/standort/v2/internal/location/orb"
)

// Module wires the domain `location` service into the application's dependency injection graph.
//
// It registers constructors for the underlying lookup providers and the domain
// `*Location` service itself:
//
//   - `ip.NewProvider`: IP → ISO country code (via GeoIP database)
//   - `country.NewProvider`: ISO country code → country + continent (via gountries)
//   - `orb.NewProvider`: lat/lng → country + continent (via point-in-polygon R-tree)
//   - `New`: composes the above providers into the domain `*Location` service
//
// This module is intended to be composed into the top-level server module (see
// `internal/cmd.Module`).
var Module = di.Module(
	di.Constructor(ip.NewProvider),
	di.Constructor(country.NewProvider),
	di.Constructor(New),
	di.Constructor(orb.NewProvider),
)
