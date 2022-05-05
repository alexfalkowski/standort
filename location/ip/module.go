package ip

import (
	"github.com/alexfalkowski/standort/location/ip/provider/trace/opentracing/jaeger"
	"go.uber.org/fx"
)

// ProviderJaegerModule for fx.
// nolint:gochecknoglobals
var ProviderJaegerModule = fx.Provide(jaeger.NewTracer)
