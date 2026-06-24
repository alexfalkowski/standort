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
//   - `standort.v1.Service/GetLocationByIP` → `Server.GetLocationByIP`
//   - `standort.v1.Service/GetLocationByLatLng` → `Server.GetLocationByLatLng`
//
// The HTTP server and route shapes are provided by `rpc.Route`; this function
// only wires the routes to the v1 server.
func Register(server *Server) {
	rpc.Route(v1.Service_GetLocationByIP_FullMethodName, server.GetLocationByIP)
	rpc.Route(v1.Service_GetLocationByLatLng_FullMethodName, server.GetLocationByLatLng)
}

// NewServer constructs a v1 HTTP `Server`.
//
// The returned server delegates response construction to the provided v1 locator.
func NewServer(locator *location.Locator) *Server {
	return &Server{locator: locator}
}

// Server implements the v1 HTTP transport handlers.
type Server struct {
	locator *location.Locator
}
