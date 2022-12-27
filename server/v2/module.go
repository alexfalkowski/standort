package v2

import (
	"github.com/alexfalkowski/standort/server/v2/transport/grpc"
	"go.uber.org/fx"
)

var (
	// Module for fx.
	Module = fx.Options(
		fx.Provide(grpc.NewServer),
		fx.Invoke(grpc.Register),
	)
)
