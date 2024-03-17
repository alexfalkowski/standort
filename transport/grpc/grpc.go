package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/client"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/standort/location/ip/provider/telemetry/tracer"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"
	g "google.golang.org/grpc"
)

// ClientOpts for gRPC.
type ClientOpts struct {
	Lifecycle    fx.Lifecycle
	ClientConfig client.Config
	Logger       *zap.Logger
	Tracer       tracer.Tracer
	Meter        metric.Meter
}

// NewClient for gRPC.
func NewClient(ctx context.Context, options ClientOpts) (*g.ClientConn, error) {
	opts := []grpc.ClientOption{
		grpc.WithClientLogger(options.Logger), grpc.WithClientTracer(options.Tracer),
		grpc.WithClientMetrics(options.Meter), grpc.WithClientRetry(options.ClientConfig.Retry),
		grpc.WithClientUserAgent(options.ClientConfig.UserAgent),
	}

	sec := options.ClientConfig.Security
	if sec != nil && sec.Enabled {
		sec, err := grpc.WithClientSecure(sec)
		if err != nil {
			return nil, err
		}

		opts = append(opts, sec)
	}

	conn, err := grpc.NewClient(ctx, options.ClientConfig.Host, opts...)
	if err != nil {
		return nil, err
	}

	options.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			conn.ResetConnectBackoff()

			return nil
		},
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return conn, nil
}
