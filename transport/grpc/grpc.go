package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/go-service/transport/grpc"
	"github.com/alexfalkowski/standort/location/ip/provider/telemetry/tracer"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
	g "google.golang.org/grpc"
)

// ClientOpts for gRPC.
type ClientOpts struct {
	ClientConfig config.Client
	Logger       *zap.Logger
	Tracer       tracer.Tracer
	Meter        metric.Meter
}

// NewClient for gRPC.
func NewClient(ctx context.Context, options ClientOpts) (*g.ClientConn, error) {
	opts := []grpc.ClientOption{
		grpc.WithClientLogger(options.Logger), grpc.WithClientTracer(options.Tracer),
		grpc.WithClientMetrics(options.Meter), grpc.WithClientRetry(&options.ClientConfig.Retry),
		grpc.WithClientUserAgent(options.ClientConfig.UserAgent),
	}

	if options.ClientConfig.Security.IsEnabled() {
		sec, err := grpc.WithClientSecure(options.ClientConfig.Security)
		if err != nil {
			return nil, err
		}

		opts = append(opts, sec)
	}

	return grpc.NewClient(ctx, options.ClientConfig.Host, opts...)
}
