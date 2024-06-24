package grpc

import (
	"github.com/alexfalkowski/go-service/transport/grpc"
	v1 "github.com/alexfalkowski/standort/api/standort/v1"
	"github.com/alexfalkowski/standort/location"
)

// Register server.
func Register(gs *grpc.Server, server v1.ServiceServer) {
	v1.RegisterServiceServer(gs.Server(), server)
}

// NewServer for gRPC.
func NewServer(location *location.Location) v1.ServiceServer {
	return &Server{location: location}
}

// Server for gRPC.
type Server struct {
	v1.UnimplementedServiceServer
	location *location.Location
}
