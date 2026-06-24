package http

import (
	"github.com/alexfalkowski/go-service/v2/net/http/rpc"
	v1 "github.com/alexfalkowski/standort/v2/api/standort/v1"
	"github.com/alexfalkowski/standort/v2/internal/api/v1/location"
)

// Register registers the v1 HTTP routes for the Standort API.
//
// This package uses go-service's RPC routing, mapping the generated full method
// names to the corresponding HTTP handlers:
//
//   - `standort.v1.Service/GetLocationByIP` → `getLocationByIP`
//   - `standort.v1.Service/GetLocationByLatLng` → `getLocationByLatLng`
//
// The HTTP server and route shapes are provided by `rpc.Route`; this function
// only wires the routes to the v1 locator.
func Register(locator *location.Locator) {
	rpc.Route(v1.Service_GetLocationByIP_FullMethodName, getLocationByIP(locator))
	rpc.Route(v1.Service_GetLocationByLatLng_FullMethodName, getLocationByLatLng(locator))
}
