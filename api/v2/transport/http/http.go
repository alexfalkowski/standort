package http

import (
	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/standort/api/v2/transport/grpc"
)

// Register for HTTP.
func Register(server *grpc.Server) {
	rpc.Route("/v2/location", server.GetLocation)
}
