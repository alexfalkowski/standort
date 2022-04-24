package grpc

import (
	"context"
	"errors"
	"fmt"

	geouri "git.jlel.se/jlelse/go-geouri"
	meta "github.com/alexfalkowski/go-service/transport/grpc/meta"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
	"github.com/alexfalkowski/standort/location"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server for gRPC.
type Server struct {
	location *location.Location
	v2.UnimplementedServiceServer
}

// GetLocation for gRPC.
func (s *Server) GetLocation(ctx context.Context, req *v2.GetLocationRequest) (*v2.GetLocationResponse, error) {
	resp := &v2.GetLocationResponse{Locations: []*v2.Location{}}

	if ip := s.ip(ctx, req); ip != "" {
		country, continent, err := s.location.GetByIP(ctx, ip)
		if err != nil {
			if errors.Is(err, location.ErrInvalid) {
				return resp, status.Error(codes.InvalidArgument, err.Error())
			}

			if errors.Is(err, location.ErrNotFound) {
				return resp, status.Error(codes.NotFound, err.Error())
			}
		}

		resp.Locations = append(resp.Locations, &v2.Location{Country: country, Continent: continent, Type: v2.Location_TYPE_IP_UNSPECIFIED})
	}

	point, err := s.point(ctx, req)
	if err != nil {
		return resp, status.Error(codes.InvalidArgument, err.Error())
	}

	if point != nil {
		country, continent, err := s.location.GetByLatLng(ctx, point.Lat, point.Lng)
		if err != nil {
			if errors.Is(err, location.ErrInvalid) {
				return resp, status.Error(codes.InvalidArgument, err.Error())
			}

			if errors.Is(err, location.ErrNotFound) {
				return resp, status.Error(codes.NotFound, err.Error())
			}
		}

		resp.Locations = append(resp.Locations, &v2.Location{Country: country, Continent: continent, Type: v2.Location_TYPE_GEO})
	}

	return resp, nil
}

func (s *Server) ip(ctx context.Context, req *v2.GetLocationRequest) string {
	if req.Ip != "" {
		return req.Ip
	}

	md := meta.ExtractIncoming(ctx)

	values := md.Get("forwarded-for")
	if len(values) > 0 {
		return values[0]
	}

	return ""
}

func (s *Server) point(ctx context.Context, req *v2.GetLocationRequest) (*v2.Point, error) {
	if req.Point != nil {
		return req.Point, nil
	}

	md := meta.ExtractIncoming(ctx)

	values := md.Get("geolocation")
	if len(values) > 0 {
		geo, err := geouri.Parse(values[0])
		if err != nil {
			return nil, fmt.Errorf("geo uri: %w", err)
		}

		return &v2.Point{Lat: geo.Latitude, Lng: geo.Longitude}, nil
	}

	return nil, nil
}
