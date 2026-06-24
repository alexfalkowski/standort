package grpc

import (
	"github.com/alexfalkowski/go-service/v2/net/grpc"
	v1 "github.com/alexfalkowski/standort/v2/api/standort/v1"
	"github.com/alexfalkowski/standort/v2/internal/api/v1/location"
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
// delegates response construction to the provided v1 locator.
func NewServer(locator *location.Locator) *Server {
	return &Server{locator: locator}
}

// Server implements the generated v1 gRPC service.
//
// It embeds `v1.UnimplementedServiceServer` for forward compatibility with
// newly-added RPC methods.
type Server struct {
	v1.UnimplementedServiceServer
	locator *location.Locator
}
