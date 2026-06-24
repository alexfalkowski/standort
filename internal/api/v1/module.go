package v1

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/standort/v2/internal/api/v1/location"
	"github.com/alexfalkowski/standort/v2/internal/api/v1/transport/grpc"
	"github.com/alexfalkowski/standort/v2/internal/api/v1/transport/http"
)

// Module wires the v1 API into the application's dependency injection graph.
//
// It registers:
//   - the v1 location module (`location.Module`), which builds generated v1 responses
//   - the v1 gRPC server implementation (`grpc.NewServer`) and its gRPC registration (`grpc.Register`)
//   - the v1 HTTP server implementation (`http.NewServer`) and its HTTP routing registration (`http.Register`)
//
// This module is intended to be included by the top-level server composition (see `internal/cmd.Module`).
var Module = di.Module(
	location.Module,
	di.Constructor(grpc.NewServer),
	di.Register(grpc.Register),
	di.Constructor(http.NewServer),
	di.Register(http.Register),
)
