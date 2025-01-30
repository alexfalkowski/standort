package grpc

import (
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/standort/api/location"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register server.
func Register(gs *grpc.Server, server *Server) {
	v2.RegisterServiceServer(gs.Server(), server)
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
