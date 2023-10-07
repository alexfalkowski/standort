package grpc

import (
	"context"
	"errors"

	"github.com/alexfalkowski/go-service/meta"
	v1 "github.com/alexfalkowski/standort/api/standort/v1"
	"github.com/alexfalkowski/standort/location"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewServer for gRPC.
func NewServer(location *location.Location) v1.ServiceServer {
	return &Server{location: location}
}

// Server for gRPC.
type Server struct {
	location *location.Location
	v1.UnimplementedServiceServer
}

// GetLocationByIP for gRPC.
func (s *Server) GetLocationByIP(ctx context.Context, req *v1.GetLocationByIPRequest) (*v1.GetLocationByIPResponse, error) {
	resp := &v1.GetLocationByIPResponse{}

	country, continent, err := s.location.GetByIP(ctx, req.Ip)
	if err != nil && errors.Is(err, location.ErrNotFound) {
		return resp, status.Error(codes.NotFound, err.Error())
	}

	resp.Location = &v1.Location{Country: country, Continent: continent}
	resp.Meta = meta.Attributes(ctx)

	return resp, nil
}

// GetLocationByLatLng for gRPC.
func (s *Server) GetLocationByLatLng(ctx context.Context, req *v1.GetLocationByLatLngRequest) (*v1.GetLocationByLatLngResponse, error) {
	resp := &v1.GetLocationByLatLngResponse{Location: &v1.Location{}}

	country, continent, err := s.location.GetByLatLng(ctx, req.Lat, req.Lng)
	if err != nil && errors.Is(err, location.ErrNotFound) {
		return resp, status.Error(codes.NotFound, err.Error())
	}

	resp.Location = &v1.Location{Country: country, Continent: continent}
	resp.Meta = meta.Attributes(ctx)

	return resp, nil
}
