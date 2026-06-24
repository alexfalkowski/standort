package grpc

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/net/grpc/codes"
	"github.com/alexfalkowski/go-service/v2/net/grpc/status"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/diagnostics"
)

// GetLocation resolves a location using any available inputs (IP address and/or geo point).
//
// This handler delegates response construction to the v2 locator and maps terminal
// lookup errors onto gRPC.
//
// Response semantics:
//   - `resp.Meta` is populated from request metadata and does not contain lookup errors.
//   - `resp.Ip` is set when an IP-derived location could be resolved.
//   - `resp.Geo` is set when a geo-derived location could be resolved.
//   - if neither input produces a location, the resulting error is mapped to a gRPC
//     `codes.NotFound` status and no response body is returned.
//
// On terminal failure, safe diagnostics carried by the error are attached as gRPC
// trailers.
func (s *Server) GetLocation(ctx context.Context, req *v2.GetLocationRequest) (*v2.GetLocationResponse, error) {
	resp, err := s.locator.Locate(ctx, req)
	if err != nil {
		setFailureTrailer(ctx, diagnostics.FromError(err))
		return nil, status.SafeError(codes.NotFound, err)
	}

	return resp, nil
}
