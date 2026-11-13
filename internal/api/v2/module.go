package v2

import (
	"github.com/alexfalkowski/go-signal/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/v2/assets"
	"github.com/alexfalkowski/standort/v2/internal/api/v2/location"
	"github.com/alexfalkowski/standort/v2/internal/api/v2/transport/grpc"
	"github.com/alexfalkowski/standort/v2/internal/api/v2/transport/http"
)

// Module wires the v2 API into the application's dependency injection graph.
//
// It registers:
//   - the v2 location module (`location.Module`), which composes the lower-level
//     transport-facing locator and v2 response locator
//   - embedded lookup asset metadata (`assets.Module`)
//   - the v2 gRPC server implementation (`grpc.NewServer`) and its gRPC registration (`grpc.Register`)
//   - the v2 HTTP server implementation (`http.NewServer`) and its HTTP routing registration (`http.Register`)
//
// This module is intended to be included by the top-level server composition (see `internal/cmd.Module`).
var Module = signal.Module(
	location.Module,
	assets.Module,
	signal.Constructor(grpc.NewServer),
	signal.Register(grpc.Register),
	signal.Constructor(http.NewServer),
	signal.Register(http.Register),
)
