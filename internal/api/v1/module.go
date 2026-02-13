package v1

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/standort/v2/internal/api/v1/transport/grpc"
	"github.com/alexfalkowski/standort/v2/internal/api/v1/transport/http"
)

// Module wires the v1 API into the application's dependency injection graph.
//
// It registers:
//   - the v1 gRPC server implementation (`grpc.NewServer`) and its gRPC registration (`grpc.Register`)
//   - the v1 HTTP routing registration (`http.Register`), which maps HTTP routes to the gRPC handlers
//
// This module is intended to be included by the top-level server composition (see `internal/cmd.Module`).
var Module = di.Module(
	di.Constructor(grpc.NewServer),
	di.Register(grpc.Register),
	di.Register(http.Register),
)
