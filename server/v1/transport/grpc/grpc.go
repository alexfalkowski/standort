package grpc

import (
	"context"
	"fmt"

	sgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	shttp "github.com/alexfalkowski/go-service/transport/http"
	v1 "github.com/alexfalkowski/standort/api/standort/v1"
	"github.com/ip2location/ip2location-go/v9"
	"github.com/opentracing/opentracing-go"
	"github.com/pariz/gountries"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	GRPCServer *grpc.Server
	HTTPServer *shttp.Server
	GRPCConfig *sgrpc.Config
	Logger     *zap.Logger
	Tracer     opentracing.Tracer
	DB         *ip2location.DB
}

// Register server.
func Register(lc fx.Lifecycle, params RegisterParams) {
	server := NewServer(params.DB)

	v1.RegisterServiceServer(params.GRPCServer, server)

	var conn *grpc.ClientConn

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			conn, _ = sgrpc.NewClient(ctx, fmt.Sprintf("127.0.0.1:%s", params.GRPCConfig.Port),
				sgrpc.WithClientConfig(params.GRPCConfig), sgrpc.WithClientLogger(params.Logger),
				sgrpc.WithClientTracer(params.Tracer), sgrpc.WithClientDialOption(grpc.WithBlock()),
			)

			return v1.RegisterServiceHandler(ctx, params.HTTPServer.Mux, conn)
		},
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})
}

// NewServer for gRPC.
func NewServer(db *ip2location.DB) v1.ServiceServer {
	return &Server{db: db, query: gountries.New()}
}
