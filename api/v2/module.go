package v2

import (
	"github.com/alexfalkowski/standort/api/location"
	"github.com/alexfalkowski/standort/api/v2/transport/grpc"
	"github.com/alexfalkowski/standort/api/v2/transport/http"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(location.NewLocator),
	fx.Provide(grpc.NewServer),
	fx.Invoke(grpc.Register),
	fx.Provide(http.NewHandler),
	fx.Invoke(http.Register),
)
