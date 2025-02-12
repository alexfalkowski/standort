package v1

import (
	"github.com/alexfalkowski/standort/internal/api/v1/transport/grpc"
	"github.com/alexfalkowski/standort/internal/api/v1/transport/http"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(grpc.NewServer),
	fx.Invoke(grpc.Register),
	fx.Invoke(http.Register),
)
