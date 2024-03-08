package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/transport/grpc/telemetry/tracer"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
	v2c "github.com/alexfalkowski/standort/client/v2/config"
	"github.com/alexfalkowski/standort/transport/grpc"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ServiceClientParams for gRPC.
type ServiceClientParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	ClientConfig *v2c.Config
	Logger       *zap.Logger
	Tracer       tracer.Tracer
	Meter        metric.Meter
}

// NewServiceClient for gRPC.
func NewServiceClient(params ServiceClientParams) (v2.ServiceClient, error) {
	opts := grpc.ClientOpts{
		ClientConfig: params.ClientConfig.Config,
		Logger:       params.Logger,
		Tracer:       params.Tracer,
		Meter:        params.Meter,
	}

	conn, err := grpc.NewClient(context.Background(), opts)
	if err != nil {
		return nil, err
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

	return v2.NewServiceClient(conn), nil
}
