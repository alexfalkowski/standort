package grpc

import (
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/location"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register server.
func Register(registrar grpc.ServiceRegistrar, server *Server) {
	v2.RegisterServiceServer(registrar, server)
}

// NewServer for gRPC.
func NewServer(service *location.Locator) *Server {
	return &Server{service: service}
}

// Server for gRPC.
type Server struct {
	v2.UnimplementedServiceServer
	service *location.Locator
}

func (s *Server) error(err error) error {
	if err == nil {
		return nil
	}

	return status.Error(codes.NotFound, err.Error())
}
