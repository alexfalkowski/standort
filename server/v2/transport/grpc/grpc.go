package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/telemetry/tracer"
	"github.com/alexfalkowski/go-service/transport/http"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Lifecycle  fx.Lifecycle
	GRPCServer *grpc.Server
	HTTPServer *http.Server
	GRPCConfig *grpc.Config
	Logger     *zap.Logger
	Tracer     tracer.Tracer
	Meter      metric.Meter
	Server     v2.ServiceServer
}

// Register server.
func Register(params RegisterParams) error {
	ctx := context.Background()

	v2.RegisterServiceServer(params.GRPCServer.Server, params.Server)

	conn, err := grpc.NewClient(ctx, "127.0.0.1:"+params.GRPCConfig.Port,
		grpc.WithClientLogger(params.Logger), grpc.WithClientTracer(params.Tracer), grpc.WithClientMetrics(params.Meter),
		grpc.WithClientRetry(&params.GRPCConfig.Retry), grpc.WithClientUserAgent(params.GRPCConfig.UserAgent),
	)
	if err != nil {
		return err
	}

	if err := v2.RegisterServiceHandler(ctx, params.HTTPServer.Mux, conn); err != nil {
		return err
	}

	params.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			conn.ResetConnectBackoff()

			return nil
		},
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return nil
}
