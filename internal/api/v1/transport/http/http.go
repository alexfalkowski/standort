package http

import (
	"github.com/alexfalkowski/go-service/v2/net/http/rpc"
	v1 "github.com/alexfalkowski/standort/v2/api/standort/v1"
	"github.com/alexfalkowski/standort/v2/internal/api/v1/transport/grpc"
)

// Register registers the v1 HTTP routes for the Standort API.
//
// This package uses go-service's HTTP↔gRPC RPC routing, mapping the generated
// full method names to the corresponding gRPC handler methods:
//
//   - `standort.v1.Service/GetLocationByIP` → `(*grpc.Server).GetLocationByIP`
//   - `standort.v1.Service/GetLocationByLatLng` → `(*grpc.Server).GetLocationByLatLng`
//
// The HTTP server and route shapes are provided by `rpc.Route`; this function
// only wires the routes to the existing gRPC implementation.
func Register(server *grpc.Server) {
	rpc.Route(v1.Service_GetLocationByIP_FullMethodName, server.GetLocationByIP)
	rpc.Route(v1.Service_GetLocationByLatLng_FullMethodName, server.GetLocationByLatLng)
}
