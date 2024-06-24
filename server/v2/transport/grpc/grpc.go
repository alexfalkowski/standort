package grpc

import (
	"github.com/alexfalkowski/go-service/transport/grpc"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
	"github.com/alexfalkowski/standort/location"
)

// Register server.
func Register(gs *grpc.Server, server v2.ServiceServer) {
	v2.RegisterServiceServer(gs.Server(), server)
}

// NewServer for gRPC.
func NewServer(location *location.Location) v2.ServiceServer {
	return &Server{location: location}
}

// Server for gRPC.
type Server struct {
	v2.UnimplementedServiceServer
	location *location.Location
}
