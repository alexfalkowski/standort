package http

import (
	"github.com/alexfalkowski/go-service/v2/net/http/rpc"
	v1 "github.com/alexfalkowski/standort/v2/api/standort/v1"
	"github.com/alexfalkowski/standort/v2/internal/api/v1/transport/grpc"
)

// Register for HTTP.
func Register(server *grpc.Server) {
	rpc.Route(v1.Service_GetLocationByIP_FullMethodName, server.GetLocationByIP)
	rpc.Route(v1.Service_GetLocationByLatLng_FullMethodName, server.GetLocationByLatLng)
}
