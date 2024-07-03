package v2

import (
	"github.com/alexfalkowski/standort/server/location"
	"github.com/alexfalkowski/standort/server/v2/transport/grpc"
	"github.com/alexfalkowski/standort/server/v2/transport/http"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(location.NewLocator),
	fx.Provide(grpc.NewServer),
	fx.Invoke(grpc.Register),
	fx.Invoke(http.Register),
)
