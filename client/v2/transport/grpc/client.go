package grpc

import (
	"github.com/alexfalkowski/go-service/env"
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

	Lifecycle fx.Lifecycle
	Client    *v2c.Config
	Logger    *zap.Logger
	Tracer    trace.Tracer
	Meter     metric.Meter
	UserAgent env.UserAgent
}

// NewServiceClient for gRPC.
func NewServiceClient(params ServiceClientParams) (v2.ServiceClient, error) {
	opts := grpc.ClientOpts{
		Lifecycle: params.Lifecycle,
		Client:    params.Client.Config,
		Logger:    params.Logger,
		Tracer:    params.Tracer,
		Meter:     params.Meter,
		UserAgent: params.UserAgent,
	}
	conn, err := grpc.NewClient(opts)

	return v2.NewServiceClient(conn), err
}
