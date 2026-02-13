package grpc

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
	v1 "github.com/alexfalkowski/standort/v2/api/standort/v1"
)

// GetLocationByIP resolves a location from an IP address.
//
// The handler delegates to the domain `(*location.Location).GetByIP` service.
//
// Response semantics:
//   - `resp.Meta` is populated from request metadata via `meta.CamelStrings(ctx, strings.Empty)`.
//   - `resp.Location` is always present and contains the resolved country/continent codes when the lookup succeeds.
//   - If the underlying lookup returns an error, it is mapped to a gRPC `codes.NotFound` status by `s.error`.
//
// Note: When an error occurs, `country` and `continent` may be empty strings.
func (s *Server) GetLocationByIP(ctx context.Context, req *v1.GetLocationByIPRequest) (*v1.GetLocationByIPResponse, error) {
	resp := &v1.GetLocationByIPResponse{}
	country, continent, err := s.location.GetByIP(ctx, req.GetIp())

	resp.Meta = meta.CamelStrings(ctx, strings.Empty)
	resp.Location = &v1.Location{Country: country, Continent: continent}

	return resp, s.error(err)
}

// GetLocationByLatLng resolves a location from a latitude/longitude coordinate.
//
// The handler delegates to the domain `(*location.Location).GetByLatLng` service.
//
// Response semantics:
//   - `resp.Meta` is populated from request metadata via `meta.CamelStrings(ctx, strings.Empty)`.
//   - `resp.Location` is always present and contains the resolved country/continent codes when the lookup succeeds.
//   - If the underlying lookup returns an error, it is mapped to a gRPC `codes.NotFound` status by `s.error`.
//
// Note: When an error occurs, `country` and `continent` may be empty strings.
func (s *Server) GetLocationByLatLng(ctx context.Context, req *v1.GetLocationByLatLngRequest) (*v1.GetLocationByLatLngResponse, error) {
	resp := &v1.GetLocationByLatLngResponse{Location: &v1.Location{}}
	country, continent, err := s.location.GetByLatLng(ctx, req.GetLat(), req.GetLng())

	resp.Meta = meta.CamelStrings(ctx, strings.Empty)
	resp.Location = &v1.Location{Country: country, Continent: continent}

	return resp, s.error(err)
}
