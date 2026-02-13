package grpc

import (
	"github.com/alexfalkowski/go-service/v2/net/grpc"
	"github.com/alexfalkowski/go-service/v2/net/grpc/codes"
	"github.com/alexfalkowski/go-service/v2/net/grpc/status"
	v1 "github.com/alexfalkowski/standort/v2/api/standort/v1"
	"github.com/alexfalkowski/standort/v2/internal/location"
)

// Register registers the v1 Standort gRPC service implementation with the given registrar.
//
// This is a thin wrapper around the generated `v1.RegisterServiceServer`.
func Register(registrar grpc.ServiceRegistrar, server *Server) {
	v1.RegisterServiceServer(registrar, server)
}

// NewServer constructs a v1 gRPC `Server`.
//
// The returned server implements the generated `standort.v1.ServiceServer` and
// delegates location resolution to the provided domain `*location.Location` service.
func NewServer(location *location.Location) *Server {
	return &Server{location: location}
}

// Server implements the generated v1 gRPC service.
//
// It embeds `v1.UnimplementedServiceServer` for forward compatibility with
// newly-added RPC methods.
type Server struct {
	v1.UnimplementedServiceServer
	location *location.Location
}

// error maps domain errors to transport errors.
//
// Current behavior: any non-nil error is returned to the client as a
// `codes.NotFound` gRPC status.
func (s *Server) error(err error) error {
	if err == nil {
		return nil
	}

	return status.Error(codes.NotFound, err.Error())
}
