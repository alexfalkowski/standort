package http

import (
	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/standort/api/v1/transport/grpc"
)

// Register for HTTP.
func Register(server *grpc.Server) {
	rpc.Route("/v1/ip", server.GetLocationByIP)
	rpc.Route("/v1/coordinate", server.GetLocationByLatLng)
}
