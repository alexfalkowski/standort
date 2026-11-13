package grpc

import (
	"github.com/alexfalkowski/go-service/v3/context"
	"github.com/alexfalkowski/go-service/v3/net/grpc/codes"
	"github.com/alexfalkowski/go-service/v3/net/grpc/status"
	v1 "github.com/alexfalkowski/standort/v3/api/standort/v1"
)

// GetLocationByIP resolves a location from an IP address.
//
// The handler delegates response construction to the v1 locator.
//
// Response semantics:
//   - `resp.Meta` is populated from request metadata via `meta.Strings(ctx)`.
//   - `resp.Location` contains the resolved country/continent codes when the lookup succeeds.
//   - If the underlying lookup returns an error, no response body is returned and
//     the error is mapped to a gRPC `codes.NotFound` status.
func (s *Server) GetLocationByIP(ctx context.Context, req *v1.GetLocationByIPRequest) (*v1.GetLocationByIPResponse, error) {
	resp, err := s.locator.LocateByIP(ctx, req)
	if err != nil {
		return nil, status.SafeError(codes.NotFound, err)
	}

	return resp, nil
}

// GetLocationByLatLng resolves a location from a latitude/longitude coordinate.
//
// The handler delegates response construction to the v1 locator.
//
// Response semantics:
//   - `resp.Meta` is populated from request metadata via `meta.Strings(ctx)`.
//   - `resp.Location` contains the resolved country/continent codes when the lookup succeeds.
//   - If the underlying lookup returns an error, no response body is returned and
//     the error is mapped to a gRPC `codes.NotFound` status.
func (s *Server) GetLocationByLatLng(ctx context.Context, req *v1.GetLocationByLatLngRequest) (*v1.GetLocationByLatLngResponse, error) {
	resp, err := s.locator.LocateByLatLng(ctx, req)
	if err != nil {
		return nil, status.SafeError(codes.NotFound, err)
	}

	return resp, nil
}
