package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/v2/meta"
	v1 "github.com/alexfalkowski/standort/api/standort/v1"
)

// GetLocationByIP for gRPC.
func (s *Server) GetLocationByIP(ctx context.Context, req *v1.GetLocationByIPRequest) (*v1.GetLocationByIPResponse, error) {
	resp := &v1.GetLocationByIPResponse{}
	country, continent, err := s.location.GetByIP(ctx, req.GetIp())

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.Location = &v1.Location{Country: country, Continent: continent}

	return resp, s.error(err)
}

// GetLocationByLatLng for gRPC.
func (s *Server) GetLocationByLatLng(ctx context.Context, req *v1.GetLocationByLatLngRequest) (*v1.GetLocationByLatLngResponse, error) {
	resp := &v1.GetLocationByLatLngResponse{Location: &v1.Location{}}
	country, continent, err := s.location.GetByLatLng(ctx, req.GetLat(), req.GetLng())

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.Location = &v1.Location{Country: country, Continent: continent}

	return resp, s.error(err)
}
