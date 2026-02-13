package v2

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/standort/v2/internal/api/location"
	"github.com/alexfalkowski/standort/v2/internal/api/v2/transport/grpc"
	"github.com/alexfalkowski/standort/v2/internal/api/v2/transport/http"
)

// Module wires the v2 API into the application's dependency injection graph.
//
// It registers:
//   - the transport-facing `location.Locator` constructor (`location.NewLocator`), which adapts the domain
//     `internal/location.Location` service to transport needs (metadata fallbacks, error attributes, etc.)
//   - the v2 gRPC server implementation (`grpc.NewServer`) and its gRPC registration (`grpc.Register`)
//   - the v2 HTTP routing registration (`http.Register`), which maps HTTP routes to the gRPC handler
//
// This module is intended to be included by the top-level server composition (see `internal/cmd.Module`).
var Module = di.Module(
	di.Constructor(location.NewLocator),
	di.Constructor(grpc.NewServer),
	di.Register(grpc.Register),
	di.Register(http.Register),
)
