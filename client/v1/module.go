package v1

import (
	"github.com/alexfalkowski/standort/client/v1/transport/grpc"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(grpc.NewServiceClient),
)