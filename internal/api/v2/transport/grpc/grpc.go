package grpc

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/net/grpc"
	"github.com/alexfalkowski/go-service/v2/net/grpc/meta"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/v2/assets"
	"github.com/alexfalkowski/standort/v2/internal/api/v2/location"
	"github.com/alexfalkowski/standort/v2/internal/diagnostics"
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
// delegates response construction to the provided v2 locator.
func NewServer(locator *location.Locator, assets *assets.Repository) *Server {
	return &Server{locator: locator, assets: assets}
}

// Server implements the generated v2 gRPC service.
//
// It embeds `v2.UnimplementedServiceServer` for forward compatibility with newly-added RPC methods.
type Server struct {
	v2.UnimplementedServiceServer
	locator *location.Locator
	assets  *assets.Repository
}

func setFailureTrailer(ctx context.Context, values diagnostics.Values) {
	_ = grpc.SetTrailer(ctx, meta.New(values.Map()))
}
