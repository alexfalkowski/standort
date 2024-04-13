package grpc

import (
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
	v2c "github.com/alexfalkowski/standort/client/v2/config"
	"github.com/alexfalkowski/standort/transport/grpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ServiceClientParams for gRPC.
type ServiceClientParams struct {
	fx.In

	Lifecycle    fx.Lifecycle
	ClientConfig *v2c.Config
	Logger       *zap.Logger
	Tracer       trace.Tracer
	Meter        metric.Meter
}

// NewServiceClient for gRPC.
func NewServiceClient(params ServiceClientParams) (v2.ServiceClient, error) {
	opts := grpc.ClientOpts{
		Lifecycle:    params.Lifecycle,
		ClientConfig: params.ClientConfig.Config,
		Logger:       params.Logger,
		Tracer:       params.Tracer,
		Meter:        params.Meter,
	}

	conn, err := grpc.NewClient(opts)
	if err != nil {
		return nil, err
	}

	return v2.NewServiceClient(conn), nil
}
