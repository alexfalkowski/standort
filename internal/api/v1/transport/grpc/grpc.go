package grpc

import (
	"github.com/alexfalkowski/go-service/transport/grpc"
	v1 "github.com/alexfalkowski/standort/api/standort/v1"
	"github.com/alexfalkowski/standort/internal/location"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register server.
func Register(gs *grpc.Server, server *Server) {
	v1.RegisterServiceServer(gs.ServiceRegistrar(), server)
}

// NewServer for gRPC.
func NewServer(location *location.Location) *Server {
	return &Server{location: location}
}

// Server for gRPC.
type Server struct {
	v1.UnimplementedServiceServer
	location *location.Location
}

func (s *Server) error(err error) error {
	if err == nil {
		return nil
	}

	return status.Error(codes.NotFound, err.Error())
}
