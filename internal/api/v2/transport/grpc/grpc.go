package grpc

import (
	"github.com/alexfalkowski/go-service/v2/net/grpc"
	"github.com/alexfalkowski/go-service/v2/net/grpc/codes"
	"github.com/alexfalkowski/go-service/v2/net/grpc/status"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/location"
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
