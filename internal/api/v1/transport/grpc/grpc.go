package grpc

import (
	"github.com/alexfalkowski/go-service/v2/net/grpc"
	"github.com/alexfalkowski/go-service/v2/net/grpc/codes"
	"github.com/alexfalkowski/go-service/v2/net/grpc/status"
	v1 "github.com/alexfalkowski/standort/v2/api/standort/v1"
	"github.com/alexfalkowski/standort/v2/internal/location"
)

// Register server.
func Register(registrar grpc.ServiceRegistrar, server *Server) {
	v1.RegisterServiceServer(registrar, server)
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
