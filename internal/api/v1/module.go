package v1

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/standort/v2/internal/api/v1/transport/grpc"
	"github.com/alexfalkowski/standort/v2/internal/api/v1/transport/http"
)

// Module for fx.
var Module = di.Module(
	di.Constructor(grpc.NewServer),
	di.Register(grpc.Register),
	di.Register(http.Register),
)
