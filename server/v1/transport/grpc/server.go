package grpc

import (
	"context"
	"errors"

	v1 "github.com/alexfalkowski/standort/api/standort/v1"
	"github.com/alexfalkowski/standort/location"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server for gRPC.
type Server struct {
	location *location.Location
	v1.UnimplementedServiceServer
}

// GetLocationByIP for gRPC.
func (s *Server) GetLocationByIP(ctx context.Context, req *v1.GetLocationByIPRequest) (*v1.GetLocationByIPResponse, error) {
	resp := &v1.GetLocationByIPResponse{}

	country, continent, err := s.location.GetByIP(ctx, req.Ip)
	if err != nil {
		if errors.Is(err, location.ErrInvalid) {
			return resp, status.Error(codes.InvalidArgument, err.Error())
		}

		if errors.Is(err, location.ErrNotFound) {
			return resp, status.Error(codes.NotFound, err.Error())
		}
	}

	resp.Location = &v1.Location{
		Country:   country,
		Continent: continent,
	}

	return resp, nil
}
