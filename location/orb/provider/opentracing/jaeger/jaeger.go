package jaeger

import (
	"github.com/alexfalkowski/go-service/trace/opentracing/jaeger"
	"github.com/alexfalkowski/standort/location/orb/provider/opentracing"
	"go.uber.org/fx"
)

// NewTracer for jaeger.
func NewTracer(lc fx.Lifecycle, cfg *jaeger.Config) (opentracing.Tracer, error) {
	return jaeger.NewTracer(jaeger.TracerParams{Lifecycle: lc, Name: "orb", Config: cfg})
}
