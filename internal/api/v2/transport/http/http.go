package http

import (
	"github.com/alexfalkowski/go-service/v2/net/http/rpc"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/v2/transport/grpc"
)

// Register for HTTP.
func Register(server *grpc.Server) {
	rpc.Route(v2.Service_GetLocation_FullMethodName, server.GetLocation)
}
