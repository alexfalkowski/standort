package http

import (
	"github.com/alexfalkowski/go-service/v2/net/http/rpc"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/v2/transport/grpc"
)

// Register registers the v2 HTTP routes for the Standort API.
//
// This package uses go-service's HTTP↔gRPC RPC routing, mapping the generated
// full method name to the corresponding gRPC handler method:
//
//   - `standort.v2.Service/GetLocation` → `(*grpc.Server).GetLocation`
//
// The HTTP server and route shapes are provided by `rpc.Route`; this function
// only wires the route to the existing gRPC implementation.
func Register(server *grpc.Server) {
	rpc.Route(v2.Service_GetLocation_FullMethodName, server.GetLocation)
}
