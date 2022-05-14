package grpc

import (
	"context"
	"fmt"

	sgrpc "github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/metrics/prometheus"
	"github.com/alexfalkowski/go-service/transport/grpc/trace/opentracing"
	shttp "github.com/alexfalkowski/go-service/transport/http"
	v1 "github.com/alexfalkowski/standort/api/standort/v1"
	"github.com/alexfalkowski/standort/location"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Lifecycle  fx.Lifecycle
	GRPCServer *grpc.Server
	HTTPServer *shttp.Server
	GRPCConfig *sgrpc.Config
	Logger     *zap.Logger
	Tracer     opentracing.Tracer
	Metrics    *prometheus.ClientMetrics
	Location   *location.Location
}

// Register server.
func Register(params RegisterParams) {
	server := NewServer(params.Location)

	v1.RegisterServiceServer(params.GRPCServer, server)

	var conn *grpc.ClientConn

	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			conn, _ = sgrpc.NewClient(
				sgrpc.ClientParams{Context: ctx, Host: fmt.Sprintf("127.0.0.1:%s", params.GRPCConfig.Port), Config: params.GRPCConfig},
				sgrpc.WithClientLogger(params.Logger), sgrpc.WithClientTracer(params.Tracer),
				sgrpc.WithClientDialOption(grpc.WithBlock()), sgrpc.WithClientMetrics(params.Metrics),
			)

			return v1.RegisterServiceHandler(ctx, params.HTTPServer.Mux, conn)
		},
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})
}

// NewServer for gRPC.
func NewServer(location *location.Location) v1.ServiceServer {
	return &Server{location: location}
}
