package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/go-service/transport/grpc/telemetry/tracer"
	"github.com/alexfalkowski/go-service/transport/http"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
	v2c "github.com/alexfalkowski/standort/client/v2/config"
	g "github.com/alexfalkowski/standort/transport/grpc"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RegisterParams for gRPC.
type RegisterParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	GRPCServer   *grpc.Server
	HTTPServer   *http.Server
	ClientConfig *v2c.Config
	Logger       *zap.Logger
	Tracer       tracer.Tracer
	Meter        metric.Meter
	Server       v2.ServiceServer
}

// Register server.
func Register(params RegisterParams) error {
	ctx := context.Background()

	v2.RegisterServiceServer(params.GRPCServer.Server, params.Server)

	opts := g.ClientOpts{
		ClientConfig: params.ClientConfig.Config,
		Logger:       params.Logger,
		Tracer:       params.Tracer,
		Meter:        params.Meter,
	}

	conn, err := g.NewClient(ctx, opts)
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
