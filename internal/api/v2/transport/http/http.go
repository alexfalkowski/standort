package http

import (
	"github.com/alexfalkowski/go-service/net/http/rpc"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
	"github.com/alexfalkowski/standort/internal/api/v2/transport/grpc"
)

// Register for HTTP.
func Register(server *grpc.Server) {
	rpc.Route(v2.Service_GetLocation_FullMethodName, server.GetLocation)
}
