package grpc

import (
	"github.com/alexfalkowski/go-service/transport/grpc"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
	"github.com/alexfalkowski/standort/server/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register server.
func Register(gs *grpc.Server, server v2.ServiceServer) {
	v2.RegisterServiceServer(gs.Server(), server)
}

// NewServer for gRPC.
func NewServer(service *service.Service) v2.ServiceServer {
	return &Server{service: service}
}

// Server for gRPC.
type Server struct {
	v2.UnimplementedServiceServer
	service *service.Service
}

func (s *Server) error(err error) error {
	if service.IsNotFound(err) {
		return status.Error(codes.NotFound, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
