package grpc

import (
	"github.com/alexfalkowski/go-service/v2/net/grpc"
	"github.com/alexfalkowski/go-service/v2/net/grpc/codes"
	"github.com/alexfalkowski/go-service/v2/net/grpc/status"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/location"
)

// Register registers the v2 Standort gRPC service implementation with the given registrar.
//
// This is a thin wrapper around the generated `v2.RegisterServiceServer`.
func Register(registrar grpc.ServiceRegistrar, server *Server) {
	v2.RegisterServiceServer(registrar, server)
}

// NewServer constructs a v2 gRPC `Server`.
//
// The returned server implements the generated `standort.v2.ServiceServer` and
// delegates request handling to the provided transport-facing `*location.Locator` service.
func NewServer(service *location.Locator) *Server {
	return &Server{service: service}
}

// Server implements the generated v2 gRPC service.
//
// It embeds `v2.UnimplementedServiceServer` for forward compatibility with newly-added RPC methods.
type Server struct {
	v2.UnimplementedServiceServer
	service *location.Locator
}

// error maps domain/service errors to transport errors.
//
// Current behavior: any non-nil error is returned to the client as a
// `codes.NotFound` gRPC status.
func (s *Server) error(err error) error {
	if err == nil {
		return nil
	}

	return status.Error(codes.NotFound, err.Error())
}
