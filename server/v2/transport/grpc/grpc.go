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
	v2.RegisterServiceServer(params.GRPCServer.Server, params.Server)

	opts := g.ClientOpts{
		Lifecycle:    params.Lifecycle,
		ClientConfig: params.ClientConfig.Config,
		Logger:       params.Logger,
		Tracer:       params.Tracer,
		Meter:        params.Meter,
	}

	conn, err := g.NewClient(opts)
	if err != nil {
		return err
	}

	return v2.RegisterServiceHandler(context.Background(), params.HTTPServer.Mux, conn)
}
